#!/bin/bash

CGO_ENABLED=0 GOOS=linux GOARCH=arm GOARM=7 go test -v -c -o /tmp/wallet-test-armv7 ./lnwallet/btcwallet
CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go test -v -c -o /tmp/wallet-test-arm64 ./lnwallet/btcwallet
CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go test -v -c -o /tmp/wallet-test-darwin-arm64 ./lnwallet/btcwallet
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go test -v -c -o /tmp/wallet-test-amd64 ./lnwallet/btcwallet
scp /tmp/wallet-test-armv7 dockerhost:/home/docker/mounts/sites/www.guggero.org/wallet-test-armv7
scp /tmp/wallet-test-arm64 dockerhost:/home/docker/mounts/sites/www.guggero.org/wallet-test-arm64
scp /tmp/wallet-test-darwin-arm64 dockerhost:/home/docker/mounts/sites/www.guggero.org/wallet-test-darwin-arm64
scp /tmp/wallet-test-amd64 dockerhost:/home/docker/mounts/sites/www.guggero.org/wallet-test-amd64
scp ./lnwallet/btcwallet/addr_test.go dockerhost:/home/docker/mounts/sites/www.guggero.org/addr_test.go
