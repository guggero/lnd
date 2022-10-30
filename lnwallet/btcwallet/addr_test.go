package btcwallet

import (
	"flag"
	"testing"
	"time"

	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/btcec/v2/schnorr"
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/btcutil/hdkeychain"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcwallet/waddrmgr"
	"github.com/btcsuite/btcwallet/walletdb"
	"github.com/stretchr/testify/require"
)

const (
	tempDir         = "/tmp"
	defaultNumAddrs = 1000

	randomRootKey = "xprv9s21ZrQH143K2FFHJTCGi23PU5ZZejrr9QdA7kiJoT41tK6t" +
		"k5YDjhBZH5GHsn5zcug4EdbRMbWDXiXp6c38wjLEymP2WJKGcLv5LCfNo9B"

	firstAddrStr = "bc1pd375vg87r4m53e4cldej6pe2qgtpknkeyfcctap79k7cnpj74" +
		"x5s8xt0kj"
)

var (
	mainNetParams = &chaincfg.MainNetParams
	password      = []byte("recovery")

	taprootAddrSchema = waddrmgr.ScopeAddrSchema{
		ExternalAddrType: waddrmgr.TaprootPubKey,
		InternalAddrType: waddrmgr.TaprootPubKey,
	}

	bip86KeyScope = waddrmgr.KeyScope{
		Purpose: 86,
		Coin:    mainNetParams.HDCoinType,
	}

	HardenedKeyStart = uint32(hdkeychain.HardenedKeyStart)

	testDir  = flag.String("dir", tempDir, "test directory")
	numAddrs = flag.Uint64(
		"addrs", defaultNumAddrs, "number of addresses to derive",
	)
)

func TestWalletAddrDerivation(t *testing.T) {
	dir := t.TempDir()
	if testDir != nil && *testDir != tempDir {
		dir = *testDir
	}
	loaderOpt := LoaderWithLocalWalletDB(dir, false, time.Second)
	loader, err := NewWalletLoader(mainNetParams, 0, loaderOpt)
	require.NoError(t, err)

	extendedKey, err := hdkeychain.NewKeyFromString(randomRootKey)
	require.NoError(t, err)

	wallet, err := loader.CreateNewWalletExtendedKey(
		password, password, extendedKey, time.Now(),
	)

	err = wallet.Unlock(password, nil)
	require.NoError(t, err)

	defer wallet.Stop()

	db := wallet.Database()

	// Because we might add new "default" key scopes over time, they are
	// created correctly for new wallets. Existing wallets don't
	// automatically add them, we need to do that manually now.
	for _, scope := range LndDefaultKeyScopes {
		_, err := wallet.Manager.FetchScopedKeyManager(scope)
		if waddrmgr.IsError(err, waddrmgr.ErrScopeNotFound) {
			// The default scope wasn't found, that probably means
			// it was added recently and older wallets don't know it
			// yet. Let's add it now.
			addrSchema := waddrmgr.ScopeAddrMap[scope]
			err := walletdb.Update(
				db, func(tx walletdb.ReadWriteTx) error {
					addrmgrNs := tx.ReadWriteBucket(
						waddrmgrNamespaceKey,
					)

					_, err := wallet.Manager.NewScopedKeyManager(
						addrmgrNs, scope, addrSchema,
					)
					return err
				},
			)
			require.NoError(t, err)
		}
	}

	scope, err := wallet.Manager.FetchScopedKeyManager(bip86KeyScope)
	if err != nil {
		// If the scope hasn't yet been created (it wouldn't been
		// loaded by default if it was), then we'll manually create the
		// scope for the first time ourselves.
		err := walletdb.Update(db, func(tx walletdb.ReadWriteTx) error {
			addrmgrNs := tx.ReadWriteBucket(waddrmgrNamespaceKey)

			scope, err = wallet.Manager.NewScopedKeyManager(
				addrmgrNs, bip86KeyScope, taprootAddrSchema,
			)
			return err
		})
		require.NoError(t, err)
	}

	maxNumAddrs := int(*numAddrs)
	for i := 0; i < maxNumAddrs; i++ {
		walletAddr := DeriveAddress(t, db, scope)
		walletAddrStr := walletAddr.Address().EncodeAddress()

		if i == 0 {
			require.Equal(t, firstAddrStr, walletAddrStr)
		}

		path := []uint32{
			HardenedKeyStart + bip86KeyScope.Purpose,
			HardenedKeyStart + bip86KeyScope.Coin,
			HardenedKeyStart + 0,
			0,
			uint32(i),
		}

		walletAddrLoaded := LoadAddress(
			t, db, scope, walletAddr.Address(),
		)
		walletAddrLoadedStr := walletAddrLoaded.Address().EncodeAddress()

		if walletAddrLoadedStr != walletAddrStr {
			t.Logf("Loaded wallet addr (%v) disagrees with "+
				"stored addr (%v)!", walletAddrLoadedStr,
				walletAddrStr)
		}

		addrBIP32 := DeriveAddr(t, extendedKey, path, false)
		addrBIP32Str := addrBIP32.EncodeAddress()
		addrQuirk := DeriveAddr(t, extendedKey, path, true)
		addrQuirkStr := addrQuirk.EncodeAddress()

		if addrQuirkStr != walletAddrStr {
			t.Logf("Wallet addr (%v) disagrees with derived addr "+
				"(%v)! BIP32: %v", walletAddrStr, addrQuirkStr,
				addrBIP32Str)
		}

		if addrQuirkStr != addrBIP32Str {
			t.Logf("BIP32 (%v) disagrees with derived addr (%v).",
				addrBIP32, addrQuirkStr)
		}

		if i != 0 && i%10000 == 0 {
			t.Logf("Derived %d addresses", i)
		}
	}
}

func DeriveAddress(t *testing.T, db walletdb.DB,
	scope *waddrmgr.ScopedKeyManager) waddrmgr.ManagedAddress {

	var firstAddr waddrmgr.ManagedAddress
	err := walletdb.Update(db, func(tx walletdb.ReadWriteTx) error {
		addrmgrNs := tx.ReadWriteBucket(waddrmgrNamespaceKey)
		addrs, err := scope.NextExternalAddresses(addrmgrNs, 0, 1)
		if err != nil {
			return err
		}

		firstAddr = addrs[0]

		return nil
	})
	require.NoError(t, err)
	require.NotNil(t, firstAddr)

	return firstAddr
}

func LoadAddress(t *testing.T, db walletdb.DB, scope *waddrmgr.ScopedKeyManager,
	addr btcutil.Address) waddrmgr.ManagedAddress {

	var (
		firstAddr waddrmgr.ManagedAddress
		err       error
	)
	dbErr := walletdb.Update(db, func(tx walletdb.ReadWriteTx) error {
		addrmgrNs := tx.ReadWriteBucket(waddrmgrNamespaceKey)

		firstAddr, err = scope.Address(addrmgrNs, addr)
		return err
	})
	require.NoError(t, dbErr)
	require.NotNil(t, firstAddr)

	return firstAddr
}

func DeriveAddr(t *testing.T, key *hdkeychain.ExtendedKey, path []uint32,
	useFix bool) *btcutil.AddressTaproot {

	var currentKey = key
	for idx, pathPart := range path {
		derivedKey, err := currentKey.DeriveNonStandard(pathPart)
		require.NoError(t, err)

		if !useFix {
			currentKey = derivedKey
			continue
		}

		// There's this special case in lnd's wallet (btcwallet) where
		// the coin type and account keys are always serialized as a
		// string and encrypted, which actually fixes the key padding
		// issue that makes the difference between DeriveNonStandard and
		// Derive. To replicate lnd's behavior exactly, we need to
		// serialize and de-serialize the extended key at the coin type
		// and account level (depth = 2 or depth = 3). This does not
		// apply to the default account (id = 0) because that is always
		// derived directly.
		depth := derivedKey.Depth()
		keyID := pathPart - hdkeychain.HardenedKeyStart
		nextID := uint32(0)
		if depth == 2 && len(path) > 2 {
			nextID = path[idx+1] - hdkeychain.HardenedKeyStart
		}
		if (depth == 2 && nextID != 0) || (depth == 3 && keyID != 0) {
			currentKey, err = hdkeychain.NewKeyFromString(
				derivedKey.String(),
			)
			require.NoError(t, err)
		} else {
			currentKey = derivedKey
		}
	}

	pubKey, err := currentKey.ECPubKey()
	require.NoError(t, err)

	p2trAddr, err := P2TRAddr(pubKey, mainNetParams)
	require.NoError(t, err)

	return p2trAddr
}

func P2TRAddr(pubKey *btcec.PublicKey,
	params *chaincfg.Params) (*btcutil.AddressTaproot, error) {

	taprootKey := txscript.ComputeTaprootKeyNoScript(pubKey)
	return btcutil.NewAddressTaproot(
		schnorr.SerializePubKey(taprootKey), params,
	)
}
