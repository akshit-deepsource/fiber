//go:build !windows
// +build !windows

package ole

func (*IInspectable) GetIids() ([]*GUID, error) {
	return []*GUID{}, NewError(E_NOTIMPL)
}

func (*IInspectable) GetRuntimeClassName() (string, error) {
	return "", NewError(E_NOTIMPL)
}

func (*IInspectable) GetTrustLevel() (uint32, error) {
	return uint32(0), NewError(E_NOTIMPL)
}
