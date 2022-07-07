//go:build !windows
// +build !windows

package ole

import "unsafe"

func (*IConnectionPoint) GetConnectionInterface(_ **GUID) int32 {
	return int32(0)
}

func (*IConnectionPoint) Advise(_ *IUnknown) (uint32, error) {
	return uint32(0), NewError(E_NOTIMPL)
}

func (*IConnectionPoint) Unadvise(_ uint32) error {
	return NewError(E_NOTIMPL)
}

func (*IConnectionPoint) EnumConnections(_ *unsafe.Pointer) (err error) {
	return NewError(E_NOTIMPL)
}
