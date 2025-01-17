// Code generated by girgen. DO NOT EDIT.

package flatpak

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
)

// #include <stdlib.h>
// #include <flatpak.h>
// #include <glib-object.h>
import "C"

// glib.Type values for flatpak-instance.go.
var GTypeInstance = externglib.Type(C.flatpak_instance_get_type())

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: GTypeInstance, F: marshalInstance},
	})
}

// InstanceOverrider contains methods that are overridable.
type InstanceOverrider interface {
}

type Instance struct {
	_ [0]func() // equal guard
	*externglib.Object
}

var (
	_ externglib.Objector = (*Instance)(nil)
)

func classInitInstancer(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapInstance(obj *externglib.Object) *Instance {
	return &Instance{
		Object: obj,
	}
}

func marshalInstance(p uintptr) (interface{}, error) {
	return wrapInstance(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// App gets the application ID of the application running in the instance.
//
// Note that this may return NULL for sandboxes that don't have an application.
//
// The function returns the following values:
//
//   - utf8 (optional): application ID or NULL.
//
func (self *Instance) App() string {
	var _arg0 *C.FlatpakInstance // out
	var _cret *C.char            // in

	_arg0 = (*C.FlatpakInstance)(unsafe.Pointer(externglib.InternObject(self).Native()))

	_cret = C.flatpak_instance_get_app(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// Arch gets the architecture of the application running in the instance.
//
// The function returns the following values:
//
//   - utf8: architecture.
//
func (self *Instance) Arch() string {
	var _arg0 *C.FlatpakInstance // out
	var _cret *C.char            // in

	_arg0 = (*C.FlatpakInstance)(unsafe.Pointer(externglib.InternObject(self).Native()))

	_cret = C.flatpak_instance_get_arch(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// Branch gets the branch of the application running in the instance.
//
// The function returns the following values:
//
//   - utf8: architecture.
//
func (self *Instance) Branch() string {
	var _arg0 *C.FlatpakInstance // out
	var _cret *C.char            // in

	_arg0 = (*C.FlatpakInstance)(unsafe.Pointer(externglib.InternObject(self).Native()))

	_cret = C.flatpak_instance_get_branch(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// ChildPid gets the PID of the application process in the sandbox.
//
// See flatpak_instance_get_pid().
//
// Note that this function may return 0 immediately after launching a sandbox,
// for a short amount of time.
//
// The function returns the following values:
//
//   - gint: application process PID.
//
func (self *Instance) ChildPid() int {
	var _arg0 *C.FlatpakInstance // out
	var _cret C.int              // in

	_arg0 = (*C.FlatpakInstance)(unsafe.Pointer(externglib.InternObject(self).Native()))

	_cret = C.flatpak_instance_get_child_pid(_arg0)
	runtime.KeepAlive(self)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// Commit gets the commit of the application running in the instance.
//
// The function returns the following values:
//
//   - utf8: commit.
//
func (self *Instance) Commit() string {
	var _arg0 *C.FlatpakInstance // out
	var _cret *C.char            // in

	_arg0 = (*C.FlatpakInstance)(unsafe.Pointer(externglib.InternObject(self).Native()))

	_cret = C.flatpak_instance_get_commit(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// ID gets the instance ID. The ID is used by Flatpak for bookkeeping purposes
// and has no further relevance.
//
// The function returns the following values:
//
//   - utf8: instance ID.
//
func (self *Instance) ID() string {
	var _arg0 *C.FlatpakInstance // out
	var _cret *C.char            // in

	_arg0 = (*C.FlatpakInstance)(unsafe.Pointer(externglib.InternObject(self).Native()))

	_cret = C.flatpak_instance_get_id(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// Info gets a keyfile that holds information about the running sandbox.
//
// This file is available as /.flatpak-info inside the sandbox as well.
//
// The most important data in the keyfile is available with separate getters,
// but there may be more information in the keyfile.
//
// The function returns the following values:
//
//   - keyFile: flatpak-info keyfile.
//
func (self *Instance) Info() *glib.KeyFile {
	var _arg0 *C.FlatpakInstance // out
	var _cret *C.GKeyFile        // in

	_arg0 = (*C.FlatpakInstance)(unsafe.Pointer(externglib.InternObject(self).Native()))

	_cret = C.flatpak_instance_get_info(_arg0)
	runtime.KeepAlive(self)

	var _keyFile *glib.KeyFile // out

	_keyFile = (*glib.KeyFile)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_keyFile)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.free(intern.C)
		},
	)

	return _keyFile
}

// Pid gets the PID of the outermost process in the sandbox. This is not the
// application process itself, but a bubblewrap 'babysitter' process.
//
// See flatpak_instance_get_child_pid().
//
// The function returns the following values:
//
//   - gint: outermost process PID.
//
func (self *Instance) Pid() int {
	var _arg0 *C.FlatpakInstance // out
	var _cret C.int              // in

	_arg0 = (*C.FlatpakInstance)(unsafe.Pointer(externglib.InternObject(self).Native()))

	_cret = C.flatpak_instance_get_pid(_arg0)
	runtime.KeepAlive(self)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// Runtime gets the ref of the runtime used in the instance.
//
// The function returns the following values:
//
//   - utf8: runtime ref.
//
func (self *Instance) Runtime() string {
	var _arg0 *C.FlatpakInstance // out
	var _cret *C.char            // in

	_arg0 = (*C.FlatpakInstance)(unsafe.Pointer(externglib.InternObject(self).Native()))

	_cret = C.flatpak_instance_get_runtime(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// RuntimeCommit gets the commit of the runtime used in the instance.
//
// The function returns the following values:
//
//   - utf8: runtime commit.
//
func (self *Instance) RuntimeCommit() string {
	var _arg0 *C.FlatpakInstance // out
	var _cret *C.char            // in

	_arg0 = (*C.FlatpakInstance)(unsafe.Pointer(externglib.InternObject(self).Native()))

	_cret = C.flatpak_instance_get_runtime_commit(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// IsRunning finds out if the sandbox represented by self is still running.
//
// The function returns the following values:
//
//   - ok: TRUE if the sandbox is still running.
//
func (self *Instance) IsRunning() bool {
	var _arg0 *C.FlatpakInstance // out
	var _cret C.gboolean         // in

	_arg0 = (*C.FlatpakInstance)(unsafe.Pointer(externglib.InternObject(self).Native()))

	_cret = C.flatpak_instance_is_running(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}
