package healthcheck

// AvailableDiskSpaceRatio returns ratio of available disk space to total
// capacity.
func AvailableDiskSpaceRatio(path string) (float64, error) {
	return 1, nil
}

// AvailableDiskSpace returns the available disk space in bytes of the given
// file system.
func AvailableDiskSpace(path string) (uint64, error) {
	// 10 MB for local storage in most browsers.
	return 10 * 1024 * 1024, nil
}
