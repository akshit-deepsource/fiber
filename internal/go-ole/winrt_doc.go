//go:build !windows
// +build !windows

package ole

// RoInitialize
func RoInitialize(_ uint32) (err error) {
	return NewError(E_NOTIMPL)
}

// RoActivateInstance
func RoActivateInstance(_ string) (ins *IInspectable, err error) {
	return nil, NewError(E_NOTIMPL)
}

// RoGetActivationFactory
func RoGetActivationFactory(_ string, _ *GUID) (ins *IInspectable, err error) {
	return nil, NewError(E_NOTIMPL)
}

// HString is handle string for pointers.
type HString uintptr

// NewHString returns a new HString for Go string.
func NewHString(_ string) (hstring HString, err error) {
	return HString(uintptr(0)), NewError(E_NOTIMPL)
}

// DeleteHString deletes HString.
func DeleteHString(_ HString) (err error) {
	return NewError(E_NOTIMPL)
}

// String returns Go string value of HString.
func (HString) String() string {
	return ""
}
