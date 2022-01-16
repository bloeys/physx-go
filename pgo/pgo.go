package pgo

/*
#cgo CPPFLAGS: -I physx-c
#cgo LDFLAGS: -L ./libs -l physx-c

#include <wrap.cxx>
#include <stdlib.h> //Needed for C.free

//Functions
struct CPxFoundation* CPxCreateFoundation();
void CPxFoundation_release(struct CPxFoundation*);

struct CPxPvd* CPxCreatePvd(struct CPxFoundation*);
void CPxPvd_release(struct CPxPvd*);

struct CPxPvdTransport* CPxDefaultPvdSocketTransportCreate(const char* address, int port, int timeoutMillis);
void CPxPvdTransport_release(struct CPxPvdTransport* cppt);
*/
import "C"
import "unsafe"

type Foundation struct {
	cFoundation *C.struct_CPxFoundation
}

func (f *Foundation) Release() {
	C.CPxFoundation_release(f.cFoundation)
}

func CreateFoundation() *Foundation {

	f := &Foundation{}
	f.cFoundation = C.CPxCreateFoundation()

	return f
}

type Pvd struct {
	cPvd *C.struct_CPxPvd
}

func (p *Pvd) Release() {
	C.CPxPvd_release(p.cPvd)
}

func CreatePvd(f *Foundation) *Pvd {

	p := &Pvd{}
	p.cPvd = C.CPxCreatePvd(f.cFoundation)

	return p
}

type PvdTransport struct {
	cPvdTr *C.struct_CPxPvdTransport
}

func (p *PvdTransport) Release() {
	C.CPxPvdTransport_release(p.cPvdTr)
}

func DefaultPvdSocketTransportCreate(host string, port, timeoutMillis int) *PvdTransport {

	hostCStr := C.CString(host)
	defer C.free(unsafe.Pointer(hostCStr))

	p := &PvdTransport{}
	p.cPvdTr = C.CPxDefaultPvdSocketTransportCreate(hostCStr, C.int(port), C.int(timeoutMillis))

	return p
}
