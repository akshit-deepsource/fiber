//go:build !windows
// +build !windows

package oleutil

import ole "github.com/gofiber/fiber/v2/internal/go-ole"

// ConnectObject creates a connection point between two services for communication.
func ConnectObject(_ *ole.IDispatch, _ *ole.GUID, _ interface{}) (uint32, error) {
	return 0, ole.NewError(ole.E_NOTIMPL)
}
