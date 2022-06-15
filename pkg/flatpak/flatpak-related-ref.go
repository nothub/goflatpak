// Code generated by girgen. DO NOT EDIT.

package flatpak

import (
	"runtime"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <flatpak.h>
// #include <glib-object.h>
import "C"

// glib.Type values for flatpak-related-ref.go.
var GTypeRelatedRef = externglib.Type(C.flatpak_related_ref_get_type())

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: GTypeRelatedRef, F: marshalRelatedRef},
	})
}

// RelatedRefOverrider contains methods that are overridable.
type RelatedRefOverrider interface {
}

type RelatedRef struct {
	_ [0]func() // equal guard
	Ref
}

var (
	_ externglib.Objector = (*RelatedRef)(nil)
)

func classInitRelatedReffer(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapRelatedRef(obj *externglib.Object) *RelatedRef {
	return &RelatedRef{
		Ref: Ref{
			Object: obj,
		},
	}
}

func marshalRelatedRef(p uintptr) (interface{}, error) {
	return wrapRelatedRef(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// Subpaths returns the subpaths that should be installed/updated for the ref.
// This returns NULL if all files should be installed.
//
// The function returns the following values:
//
//    - utf8s: strv, or NULL.
//
func (self *RelatedRef) Subpaths() []string {
	var _arg0 *C.FlatpakRelatedRef // out
	var _cret **C.char             // in

	_arg0 = (*C.FlatpakRelatedRef)(unsafe.Pointer(externglib.InternObject(self).Native()))

	_cret = C.flatpak_related_ref_get_subpaths(_arg0)
	runtime.KeepAlive(self)

	var _utf8s []string // out

	{
		var i int
		var z *C.char
		for p := _cret; *p != z; p = &unsafe.Slice(p, 2)[1] {
			i++
		}

		src := unsafe.Slice(_cret, i)
		_utf8s = make([]string, i)
		for i := range src {
			_utf8s[i] = C.GoString((*C.gchar)(unsafe.Pointer(src[i])))
		}
	}

	return _utf8s
}

// ShouldAutoprune returns whether to delete when pruning unused refs.
//
// The function returns the following values:
//
//    - ok: TRUE if the ref should be considered unused when pruning.
//
func (self *RelatedRef) ShouldAutoprune() bool {
	var _arg0 *C.FlatpakRelatedRef // out
	var _cret C.gboolean           // in

	_arg0 = (*C.FlatpakRelatedRef)(unsafe.Pointer(externglib.InternObject(self).Native()))

	_cret = C.flatpak_related_ref_should_autoprune(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ShouldDelete returns whether to auto-delete the ref with the main ref.
//
// The function returns the following values:
//
//    - ok: TRUE if the ref should be deleted with the main ref.
//
func (self *RelatedRef) ShouldDelete() bool {
	var _arg0 *C.FlatpakRelatedRef // out
	var _cret C.gboolean           // in

	_arg0 = (*C.FlatpakRelatedRef)(unsafe.Pointer(externglib.InternObject(self).Native()))

	_cret = C.flatpak_related_ref_should_delete(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ShouldDownload returns whether to auto-download the ref with the main ref.
//
// The function returns the following values:
//
//    - ok: TRUE if the ref should be downloaded with the main ref.
//
func (self *RelatedRef) ShouldDownload() bool {
	var _arg0 *C.FlatpakRelatedRef // out
	var _cret C.gboolean           // in

	_arg0 = (*C.FlatpakRelatedRef)(unsafe.Pointer(externglib.InternObject(self).Native()))

	_cret = C.flatpak_related_ref_should_download(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}
