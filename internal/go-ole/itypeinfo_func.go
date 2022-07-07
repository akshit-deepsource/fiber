//go:build !windows
// +build !windows

package ole

func (*ITypeInfo) GetTypeAttr() (*TYPEATTR, error) {
	return nil, NewError(E_NOTIMPL)
}
