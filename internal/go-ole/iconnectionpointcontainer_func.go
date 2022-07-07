//go:build !windows
// +build !windows

package ole

func (*IConnectionPointContainer) EnumConnectionPoints(_ interface{}) error {
	return NewError(E_NOTIMPL)
}

func (*IConnectionPointContainer) FindConnectionPoint(_ *GUID, _ **IConnectionPoint) error {
	return NewError(E_NOTIMPL)
}
