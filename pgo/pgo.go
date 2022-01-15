package pgo

/*
#cgo CPPFLAGS: -I physx-c
#cgo LDFLAGS: -L ./libs -l physx-c

#include <wrap.cxx>
#include <stdlib.h> //Needed for C.free

//Functions
struct CPxFoundation* NewCPxFoundation();
void FreeCPxFoundation(struct CPxFoundation*);
*/
import "C"

func Test() {
	x := C.NewCPxFoundation()
	println("Result:", x)
	C.FreeCPxFoundation(x)
}
