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
bool CPxPvd_connect(struct CPxPvd*, struct CPxPvdTransport*, enum CPxPvdInstrumentationFlag);
void CPxPvd_release(struct CPxPvd*);

struct CPxPvdTransport* CPxDefaultPvdSocketTransportCreate(const char* address, int port, int timeoutMillis);
void CPxPvdTransport_release(struct CPxPvdTransport* cppt);

struct CPxTolerancesScale NewCPxTolerancesScale(CPxReal length, CPxReal speed);

struct CPxPhysics* CPxCreatePhysics(struct CPxFoundation* cfoundation, struct CPxTolerancesScale cscale, bool trackOutstandingAllocations, struct CPxPvd* cpvd);
struct CPxScene* CPxPhysics_createScene(struct CPxPhysics*, struct CPxSceneDesc*);
void CPxPhysics_release(struct CPxPhysics*);

struct CPxPvdSceneClient* CPxScene_getScenePvdClient(struct CPxScene*);

void CPxPvdSceneClient_setScenePvdFlag(struct CPxPvdSceneClient* c, enum CPxPvdSceneFlag flag, bool value);

struct CPxVec3 NewCPxVec3(float x, float y, float z);

struct CPxDefaultCpuDispatcher* CPxDefaultCpuDispatcherCreate(CPxU32 numThreads, CPxU32 affinityMasks);
struct CPxCpuDispatcher* CPxDefaultCpuDispatcher_toCPxCpuDispatcher(struct CPxDefaultCpuDispatcher* cdcd);

struct CPxSceneDesc* NewCPxSceneDesc(struct CPxTolerancesScale);
void CPxSceneDesc_set_gravity(struct CPxSceneDesc*, struct CPxVec3);
void CPxSceneDesc_set_cpuDispatcher(struct CPxSceneDesc*, struct CPxCpuDispatcher*);
*/
import "C"

type PvdInstrumentationFlag uint32

const (
	PvdInstrumentationFlag_eDEBUG   PvdInstrumentationFlag = 1 << 0
	PvdInstrumentationFlag_ePROFILE PvdInstrumentationFlag = 1 << 1
	PvdInstrumentationFlag_eMEMORY  PvdInstrumentationFlag = 1 << 2
	PvdInstrumentationFlag_eALL     PvdInstrumentationFlag = (PvdInstrumentationFlag_eDEBUG | PvdInstrumentationFlag_ePROFILE | PvdInstrumentationFlag_eMEMORY)
)

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

func (p *Pvd) Connect(pvdTr *PvdTransport, instFlag PvdInstrumentationFlag) bool {
	return bool(C.CPxPvd_connect(p.cPvd, pvdTr.cPvdTr, uint32(instFlag)))
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

	//This CString should NOT be freed because its stored internally. If this is freed connection to PVD will fail
	p := &PvdTransport{}
	p.cPvdTr = C.CPxDefaultPvdSocketTransportCreate(C.CString(host), C.int(port), C.int(timeoutMillis))
	return p
}

type TolerancesScale struct {
	cTolScale C.struct_CPxTolerancesScale
}

func NewTolerancesScale(length, speed float32) *TolerancesScale {

	ts := &TolerancesScale{}
	ts.cTolScale = C.NewCPxTolerancesScale(C.float(length), C.float(speed))
	return ts
}

type Scene struct {
	cS *C.struct_CPxScene
}

func (s *Scene) GetScenePvdClient() *PvdSceneClient {
	return &PvdSceneClient{
		cPvdSceneClient: C.CPxScene_getScenePvdClient(s.cS),
	}
}

type Physics struct {
	cPhysics *C.struct_CPxPhysics
}

func (p *Physics) CreateScene(sd *SceneDesc) *Scene {
	return &Scene{
		cS: C.CPxPhysics_createScene(p.cPhysics, sd.cSD),
	}
}

func CreatePhysics(f *Foundation, ts *TolerancesScale, trackOutstandingAllocations bool, pvd *Pvd) *Physics {

	p := &Physics{}
	p.cPhysics = C.CPxCreatePhysics(f.cFoundation, ts.cTolScale, C._Bool(trackOutstandingAllocations), pvd.cPvd)

	return p
}

func (p *Physics) Release() {
	C.CPxPhysics_release(p.cPhysics)
}

type Vec3 struct {
	cV C.struct_CPxVec3
}

func NewVec3(x, y, z float32) *Vec3 {
	return &Vec3{
		cV: C.NewCPxVec3(C.float(x), C.float(y), C.float(z)),
	}
}

type CpuDispatcher struct {
	cCpuDisp *C.struct_CPxCpuDispatcher
}

type DefaultCpuDispatcher struct {
	cDefCpuDisp *C.struct_CPxDefaultCpuDispatcher
}

// struct CPxCpuDispatcher* CPxDefaultCpuDispatcher_toCPxCpuDispatcher(struct CPxDefaultCpuDispatcher* cdcd);
func (d *DefaultCpuDispatcher) ToCpuDispatcher() *CpuDispatcher {
	return &CpuDispatcher{cCpuDisp: (*C.struct_CPxCpuDispatcher)(d.cDefCpuDisp)}
}

func DefaultCpuDispatcherCreate(numThreads, affinityMasks uint32) *DefaultCpuDispatcher {
	return &DefaultCpuDispatcher{
		cDefCpuDisp: C.CPxDefaultCpuDispatcherCreate(C.uint(numThreads), C.uint(affinityMasks)),
	}
}

type SceneDesc struct {
	cSD *C.struct_CPxSceneDesc
}

func (sd *SceneDesc) SetGravity(v *Vec3) {
	C.CPxSceneDesc_set_gravity(sd.cSD, v.cV)
}

func (sd *SceneDesc) SetCpuDispatcher(cd *CpuDispatcher) {
	C.CPxSceneDesc_set_cpuDispatcher(sd.cSD, cd.cCpuDisp)
}

//struct CPxSceneDesc* NewCPxSceneDesc(struct CPxTolerancesScale);
func NewSceneDesc(ts *TolerancesScale) *SceneDesc {
	return &SceneDesc{
		cSD: C.NewCPxSceneDesc(ts.cTolScale),
	}
}

type PvdSceneFlag uint32

const (
	PvdSceneFlag_eTRANSMIT_CONTACTS     PvdSceneFlag = (1 << 0) //Transmits contact stream to PVD.
	PvdSceneFlag_eTRANSMIT_SCENEQUERIES PvdSceneFlag = (1 << 1) //Transmits scene query stream to PVD.
	PvdSceneFlag_eTRANSMIT_CONSTRAINTS  PvdSceneFlag = (1 << 2) //Transmits constraints visualize stream to PVD.
)

type PvdSceneClient struct {
	cPvdSceneClient *C.struct_CPxPvdSceneClient
}

func (p *PvdSceneClient) SetScenePvdFlag(flag PvdSceneFlag, value bool) {
	C.CPxPvdSceneClient_setScenePvdFlag(p.cPvdSceneClient, uint32(flag), C._Bool(value))
}
