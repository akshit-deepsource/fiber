//go:build !windows
// +build !windows

package ole

// errstr converts error code to string.
func errstr(_ int) string {
	return ""
}
