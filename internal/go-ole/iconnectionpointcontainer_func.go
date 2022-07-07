//go:build !windows
// +build !windows

package ole

func (v *IConnectionPointContainer) EnumConnectionPoints(_ interface{}) error {
	return NewError(E_NOTIMPL)
}

func (v *IConnectionPointContainer) FindConnectionPoint(_ *GUID, _ **IConnectionPoint) error {
	return NewError(E_NOTIMPL)
}
