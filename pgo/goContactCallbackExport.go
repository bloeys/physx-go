package pgo

/*
#include "wrap.c"
*/
import "C"
import "unsafe"

//export goOnContactCallback
func goOnContactCallback(p unsafe.Pointer) {
	contactCallback(ContactPairHeader{cCPH: (*C.struct_CPxContactPairHeader)(p)})
}
