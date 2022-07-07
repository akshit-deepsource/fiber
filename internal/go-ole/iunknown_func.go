//go:build !windows
// +build !windows

package ole

func reflectQueryInterface(_ interface{}, _ uintptr, _ *GUID, _ interface{}) (err error) {
	return NewError(E_NOTIMPL)
}

func queryInterface(_ *IUnknown, _ *GUID) (disp *IDispatch, err error) {
	return nil, NewError(E_NOTIMPL)
}

func addRef(_ *IUnknown) int32 {
	return 0
}

func release(_ *IUnknown) int32 {
	return 0
}
