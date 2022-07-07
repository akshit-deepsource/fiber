//go:build !windows
// +build !windows

package ole

func getIDsOfName(_ *IDispatch, _ []string) ([]int32, error) {
	return []int32{}, NewError(E_NOTIMPL)
}

func getTypeInfoCount(_ *IDispatch) (uint32, error) {
	return uint32(0), NewError(E_NOTIMPL)
}

func getTypeInfo(_ *IDispatch) (*ITypeInfo, error) {
	return nil, NewError(E_NOTIMPL)
}

func invoke(_ *IDispatch, _ int32, _ int16, _ ...interface{}) (*VARIANT, error) {
	return nil, NewError(E_NOTIMPL)
}
