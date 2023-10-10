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

//export goOnTriggerCallback
func goOnTriggerCallback(p unsafe.Pointer, count uint32) {

	// @PERF
	triggerPairs := make([]TriggerPair, count)
	tPairs := unsafe.Slice((*C.struct_CPxTriggerPair)(p), count)

	for i := 0; i < len(triggerPairs); i++ {
		triggerPairs[i].cTp = &tPairs[i]
	}

	triggerCallback(triggerPairs)
}
