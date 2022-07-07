//go:build !windows
// +build !windows

package ole

import (
	"unsafe"
)

// safeArrayAccessData returns raw array pointer.
//
// AKA: SafeArrayAccessData in Windows API.
func safeArrayAccessData(_ *SafeArray) (uintptr, error) {
	return uintptr(0), NewError(E_NOTIMPL)
}

// safeArrayUnaccessData releases raw array.
//
// AKA: SafeArrayUnaccessData in Windows API.
func safeArrayUnaccessData(_ *SafeArray) error {
	return NewError(E_NOTIMPL)
}

// safeArrayAllocData allocates SafeArray.
//
// AKA: SafeArrayAllocData in Windows API.
func safeArrayAllocData(_ *SafeArray) error {
	return NewError(E_NOTIMPL)
}

// safeArrayAllocDescriptor allocates SafeArray.
//
// AKA: SafeArrayAllocDescriptor in Windows API.
func safeArrayAllocDescriptor(_ uint32) (*SafeArray, error) {
	return nil, NewError(E_NOTIMPL)
}

// safeArrayAllocDescriptorEx allocates SafeArray.
//
// AKA: SafeArrayAllocDescriptorEx in Windows API.
func safeArrayAllocDescriptorEx(_ VT, _ uint32) (*SafeArray, error) {
	return nil, NewError(E_NOTIMPL)
}

// safeArrayCopy returns copy of SafeArray.
//
// AKA: SafeArrayCopy in Windows API.
func safeArrayCopy(_ *SafeArray) (*SafeArray, error) {
	return nil, NewError(E_NOTIMPL)
}

// safeArrayCopyData duplicates SafeArray into another SafeArray object.
//
// AKA: SafeArrayCopyData in Windows API.
func safeArrayCopyData(_ *SafeArray, _ *SafeArray) error {
	return NewError(E_NOTIMPL)
}

// safeArrayCreate creates SafeArray.
//
// AKA: SafeArrayCreate in Windows API.
func safeArrayCreate(_ VT, _ uint32, _ *SafeArrayBound) (*SafeArray, error) {
	return nil, NewError(E_NOTIMPL)
}

// safeArrayCreateEx creates SafeArray.
//
// AKA: SafeArrayCreateEx in Windows API.
func safeArrayCreateEx(_ VT, _ uint32, _ *SafeArrayBound, _ uintptr) (*SafeArray, error) {
	return nil, NewError(E_NOTIMPL)
}

// safeArrayCreateVector creates SafeArray.
//
// AKA: SafeArrayCreateVector in Windows API.
func safeArrayCreateVector(_ VT, _ int32, _ uint32) (*SafeArray, error) {
	return nil, NewError(E_NOTIMPL)
}

// safeArrayCreateVectorEx creates SafeArray.
//
// AKA: SafeArrayCreateVectorEx in Windows API.
func safeArrayCreateVectorEx(_ VT, _ int32, _ uint32, _ uintptr) (*SafeArray, error) {
	return nil, NewError(E_NOTIMPL)
}

// safeArrayDestroy destroys SafeArray object.
//
// AKA: SafeArrayDestroy in Windows API.
func safeArrayDestroy(_ *SafeArray) error {
	return NewError(E_NOTIMPL)
}

// safeArrayDestroyData destroys SafeArray object.
//
// AKA: SafeArrayDestroyData in Windows API.
func safeArrayDestroyData(_ *SafeArray) error {
	return NewError(E_NOTIMPL)
}

// safeArrayDestroyDescriptor destroys SafeArray object.
//
// AKA: SafeArrayDestroyDescriptor in Windows API.
func safeArrayDestroyDescriptor(_ *SafeArray) error {
	return NewError(E_NOTIMPL)
}

// safeArrayGetDim is the amount of dimensions in the SafeArray.
//
// SafeArrays may have multiple dimensions. Meaning, it could be
// multidimensional array.
//
// AKA: SafeArrayGetDim in Windows API.
func safeArrayGetDim(_ *SafeArray) (*uint32, error) {
	u := uint32(0)
	return &u, NewError(E_NOTIMPL)
}

// safeArrayGetElementSize is the element size in bytes.
//
// AKA: SafeArrayGetElemsize in Windows API.
func safeArrayGetElementSize(_ *SafeArray) (*uint32, error) {
	u := uint32(0)
	return &u, NewError(E_NOTIMPL)
}

// safeArrayGetElement retrieves element at given index.
func safeArrayGetElement(_ *SafeArray, _ int32, _ unsafe.Pointer) error {
	return NewError(E_NOTIMPL)
}

// safeArrayGetElement retrieves element at given index and converts to string.
func safeArrayGetElementString(_ *SafeArray, _ int32) (string, error) {
	return "", NewError(E_NOTIMPL)
}

// safeArrayGetIID is the InterfaceID of the elements in the SafeArray.
//
// AKA: SafeArrayGetIID in Windows API.
func safeArrayGetIID(_ *SafeArray) (*GUID, error) {
	return nil, NewError(E_NOTIMPL)
}

// safeArrayGetLBound returns lower bounds of SafeArray.
//
// SafeArrays may have multiple dimensions. Meaning, it could be
// multidimensional array.
//
// AKA: SafeArrayGetLBound in Windows API.
func safeArrayGetLBound(_ *SafeArray, _ uint32) (int32, error) {
	return int32(0), NewError(E_NOTIMPL)
}

// safeArrayGetUBound returns upper bounds of SafeArray.
//
// SafeArrays may have multiple dimensions. Meaning, it could be
// multidimensional array.
//
// AKA: SafeArrayGetUBound in Windows API.
func safeArrayGetUBound(_ *SafeArray, _ uint32) (int32, error) {
	return int32(0), NewError(E_NOTIMPL)
}

// safeArrayGetVartype returns data type of SafeArray.
//
// AKA: SafeArrayGetVartype in Windows API.
func safeArrayGetVartype(_ *SafeArray) (uint16, error) {
	return uint16(0), NewError(E_NOTIMPL)
}

// safeArrayLock locks SafeArray for reading to modify SafeArray.
//
// This must be called during some calls to ensure that another process does not
// read or write to the SafeArray during editing.
//
// AKA: SafeArrayLock in Windows API.
func safeArrayLock(_ *SafeArray) error {
	return NewError(E_NOTIMPL)
}

// safeArrayUnlock unlocks SafeArray for reading.
//
// AKA: SafeArrayUnlock in Windows API.
func safeArrayUnlock(_ *SafeArray) error {
	return NewError(E_NOTIMPL)
}

// safeArrayPutElement stores the data element at the specified location in the
// array.
//
// AKA: SafeArrayPutElement in Windows API.
func safeArrayPutElement(_ *SafeArray, _ int64, _ uintptr) error {
	return NewError(E_NOTIMPL)
}

// safeArrayGetRecordInfo accesses IRecordInfo info for custom types.
//
// AKA: SafeArrayGetRecordInfo in Windows API.
//
// XXX: Must implement IRecordInfo interface for this to return.
func safeArrayGetRecordInfo(_ *SafeArray) (interface{}, error) {
	return nil, NewError(E_NOTIMPL)
}

// safeArraySetRecordInfo mutates IRecordInfo info for custom types.
//
// AKA: SafeArraySetRecordInfo in Windows API.
//
// XXX: Must implement IRecordInfo interface for this to return.
func safeArraySetRecordInfo(_ *SafeArray, _ interface{}) error {
	return NewError(E_NOTIMPL)
}
