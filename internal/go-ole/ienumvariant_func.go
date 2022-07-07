//go:build !windows
// +build !windows

package ole

func (*IEnumVARIANT) Clone() (*IEnumVARIANT, error) {
	return nil, NewError(E_NOTIMPL)
}

func (*IEnumVARIANT) Reset() error {
	return NewError(E_NOTIMPL)
}

func (*IEnumVARIANT) Skip(_ uint) error {
	return NewError(E_NOTIMPL)
}

func (*IEnumVARIANT) Next(_ uint) (VARIANT, uint, error) {
	return NewVariant(VT_NULL, int64(0)), 0, NewError(E_NOTIMPL)
}
