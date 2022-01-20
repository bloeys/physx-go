package pgo

/*
#cgo CPPFLAGS: -I physx-c
#cgo LDFLAGS: -L ./libs -l physx-c

#include <wrap.cxx>
#include <stdlib.h> //Needed for C.free
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

func (s *Scene) AddActor(a *Actor) {
	C.CPxScene_addActor(s.cS, a.cA)
}

// void CPxScene_simulate(CSTRUCT CPxScene*, CPxReal elapsedTime);
func (s *Scene) Simulate(elapsedTime float32) {
	C.CPxScene_simulate(s.cS, C.float(elapsedTime))
}

// bool CPxScene_fetchResults(struct CPxScene*, bool block, CPxU32* errorState);
func (s *Scene) FetchResults(block bool) (bool, uint32) {

	var errState uint32
	b := C.CPxScene_fetchResults(s.cS, C._Bool(block), (*C.uint)(&errState))
	return bool(b), errState
}

type Physics struct {
	cPhysics *C.struct_CPxPhysics
}

func (p *Physics) CreateScene(sd *SceneDesc) *Scene {
	return &Scene{
		cS: C.CPxPhysics_createScene(p.cPhysics, sd.cSD),
	}
}

func (p *Physics) CreateMaterial(staticFriction, dynamicFriction, restitution float32) *Material {
	return &Material{
		cM: C.CPxPhysics_createMaterial(p.cPhysics, C.float(staticFriction), C.float(dynamicFriction), C.float(restitution)),
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

type Material struct {
	cM *C.struct_CPxMaterial
}

type Plane struct {
	cP *C.struct_CPxPlane
}

func NewPlane(nx, ny, nz, distance float32) *Plane {
	return &Plane{
		cP: C.NewCPxPlane(C.float(nx), C.float(ny), C.float(nz), C.float(distance)),
	}
}

type Quat struct {
	cQ C.struct_CPxQuat
}

// CPxAPI CPxInline CSTRUCT CPxQuat NewCPxQuat(float angleRads, float x, float y, float z);
func NewQuat(angleRads, x, y, z float32) *Quat {
	return &Quat{
		cQ: C.NewCPxQuat(C.float(angleRads), C.float(x), C.float(y), C.float(z)),
	}
}

type Transform struct {
	cT C.struct_CPxTransform
}

// struct CPxTransform NewCPxTransform(struct CPxVec3*, struct CPxQuat*);
func NewTransform(v *Vec3, q *Quat) *Transform {
	return &Transform{
		cT: C.NewCPxTransform(&v.cV, &q.cQ),
	}
}

type Geometry struct {
	cG C.struct_CPxGeometry
}

type SphereGeometry struct {
	cSg C.struct_CPxSphereGeometry
}

// struct CPxGeometry CPxSphereGeometry_toCPxGeometry(struct CPxSphereGeometry*);
func (sg *SphereGeometry) ToGeometry() *Geometry {
	return &Geometry{
		cG: C.CPxSphereGeometry_toCPxGeometry(&sg.cSg),
	}
}

// struct CPxSphereGeometry NewCPxSphereGeometry(CPxReal radius);
func NewSphereGeometry(radius float32) *SphereGeometry {
	return &SphereGeometry{
		cSg: C.NewCPxSphereGeometry(C.float(radius)),
	}
}

type BoxGeometry struct {
	cBg C.struct_CPxBoxGeometry
}

func (bg *BoxGeometry) ToGeometry() *Geometry {
	return &Geometry{
		cG: C.CPxBoxGeometry_toCPxGeometry(&bg.cBg),
	}
}

func NewBoxGeometry(hx, hy, hz float32) *BoxGeometry {
	return &BoxGeometry{
		cBg: C.NewCPxBoxGeometry(C.float(hx), C.float(hy), C.float(hz)),
	}
}

type Actor struct {
	cA C.struct_CPxActor
}

type RigidStatic struct {
	cRs *C.struct_CPxRigidStatic
}

func (rs *RigidStatic) ToActor() *Actor {
	return &Actor{
		cA: C.CPxRigidStatic_toCPxActor(rs.cRs),
	}
}

func CreatePlane(p *Physics, plane *Plane, mat *Material) *RigidStatic {
	return &RigidStatic{
		cRs: C.CPxCreatePlane(p.cPhysics, plane.cP, mat.cM),
	}
}

type RigidDynamic struct {
	cRd *C.struct_CPxRigidDynamic
}

func (rs *RigidDynamic) ToActor() *Actor {
	return &Actor{
		cA: C.CPxRigidDynamic_toCPxActor(rs.cRd),
	}
}

func CreateDynamic(p *Physics, t *Transform, g *Geometry, m *Material, density float32, shapeOffset *Transform) *RigidDynamic {
	return &RigidDynamic{
		cRd: C.CPxCreateDynamic(p.cPhysics, &t.cT, g.cG, m.cM, C.float(density), &shapeOffset.cT),
	}
}
