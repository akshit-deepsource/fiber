//go:build !windows
// +build !windows

package ole

func getClassInfo(_ *IProvideClassInfo) (tinfo *ITypeInfo, err error) {
	return nil, NewError(E_NOTIMPL)
}
