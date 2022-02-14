package pgo

/*
#cgo CFLAGS: -I physx-c
#cgo LDFLAGS: -L ./libs -l physx-c

#include <wrap.c>
#include <stdlib.h> //Needed for C.free

//simulation event callbacks forward declarations. Actual definitions MUST be in a go different file
void goOnContactCallback_cgo(void* pairHeader);
*/
import "C"
import "unsafe"

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

func (s *Scene) AddActor(a Actor) {
	C.CPxScene_addActor(s.cS, a.cA)
}

func (s *Scene) Simulate(elapsedTime float32) {
	C.CPxScene_simulate(s.cS, C.float(elapsedTime))
}

func (s *Scene) Collide(elapsedTime float32) {
	C.CPxScene_collide(s.cS, C.float(elapsedTime))
}

func (s *Scene) FetchCollision(block bool) bool {
	return bool(C.CPxScene_fetchCollision(s.cS, C._Bool(block)))
}

func (s *Scene) Advance() {
	C.CPxScene_advance(s.cS)
}

func (s *Scene) FetchResults(block bool) (bool, uint32) {

	var errState uint32
	b := C.CPxScene_fetchResults(s.cS, C._Bool(block), (*C.uint)(&errState))
	return bool(b), errState
}

func (s *Scene) SetScratchBuffer(multiplesOf16k uint32) {
	C.CPxScene_setScratchBuffer(s.cS, C.uint(multiplesOf16k))
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

func (p *Physics) CreateRigidDynamic(tr *Transform) *RigidDynamic {
	return &RigidDynamic{
		cRd: C.CPxPhysics_createRigidDynamic(p.cPhysics, &tr.cT),
	}
}

func (p *Physics) CreateRigidStatic(tr *Transform) *RigidStatic {
	return &RigidStatic{
		cRs: C.CPxPhysics_createRigidStatic(p.cPhysics, &tr.cT),
	}
}

func (p *Physics) Release() {
	C.CPxPhysics_release(p.cPhysics)
}

func CreatePhysics(f *Foundation, ts *TolerancesScale, trackOutstandingAllocations bool, pvd *Pvd) *Physics {

	p := &Physics{}
	p.cPhysics = C.CPxCreatePhysics(f.cFoundation, ts.cTolScale, C._Bool(trackOutstandingAllocations), pvd.cPvd)

	return p
}

type FilterData struct {
	cFilterData C.struct_CPxFilterData
}

func NewFilterData(w0, w1, w2, w3 uint32) FilterData {
	return FilterData{
		cFilterData: C.struct_CPxFilterData{
			word0: C.uint(w0),
			word1: C.uint(w1),
			word2: C.uint(w2),
			word3: C.uint(w3),
		},
	}
}

type Shape struct {
	cShape C.struct_CPxShape
}

func (s *Shape) GetLocalPose() *Transform {
	return &Transform{
		cT: C.CPxShape_getLocalPose(&s.cShape),
	}
}

func (s *Shape) SetLocalPose(tr *Transform) {
	C.CPxShape_setLocalPose(&s.cShape, &tr.cT)
}

func (s *Shape) GetSimulationFilterData() FilterData {
	return FilterData{
		cFilterData: C.CPxShape_getSimulationFilterData(&s.cShape),
	}
}

func (s *Shape) SetSimulationFilterData(fd *FilterData) {
	C.CPxShape_setSimulationFilterData(&s.cShape, &fd.cFilterData)
}

func CreateExclusiveShape(rigidActor RigidActor, geom *Geometry, mat *Material, shapeFlags ShapeFlags) Shape {
	return Shape{
		cShape: C.createExclusiveShape(rigidActor.cRa, geom.cG, mat.cM, uint32(shapeFlags)),
	}
}

type Vec3 struct {
	cV C.struct_CPxVec3
}

func (v *Vec3) X() float32 {
	return float32(v.cV.x)
}

func (v *Vec3) Y() float32 {
	return float32(v.cV.y)
}

func (v *Vec3) Z() float32 {
	return float32(v.cV.z)
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

//export goOnContactCallback
func goOnContactCallback(p unsafe.Pointer) {
	println("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
}

func (sd *SceneDesc) SetOnContactCallback() {
	C.CPxSceneDesc_set_onContactCallback(sd.cSD, (C.CPxonContactCallback)(unsafe.Pointer(C.goOnContactCallback_cgo)))
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

type CapsuleGeometry struct {
	cCg C.struct_CPxCapsuleGeometry
}

func (bg *CapsuleGeometry) ToGeometry() *Geometry {
	return &Geometry{
		cG: C.CPxCapsuleGeometry_toCPxGeometry(&bg.cCg),
	}
}

func NewCapsuleGeometry(radius, halfHeight float32) *CapsuleGeometry {
	return &CapsuleGeometry{
		cCg: C.NewCPxCapsuleGeometry(C.float(radius), C.float(halfHeight)),
	}
}

type Actor struct {
	cA C.struct_CPxActor
}

type RigidActor struct {
	cRa C.struct_CPxRigidActor
}

func (ra *RigidActor) SetSimFilterData(fd *FilterData) {
	C.CPxRigidActor_setSimFilterData(&ra.cRa, &fd.cFilterData)
}

// CPxAPI void CPxRigidActor_setSimFilterData(CSTRUCT CPxRigidActor* cra, CSTRUCT CPxFilterData* cfd);

type RigidStatic struct {
	cRs *C.struct_CPxRigidStatic
}

func (rs *RigidStatic) ToActor() Actor {
	return Actor{
		cA: C.CPxRigidStatic_toCPxActor(rs.cRs),
	}
}

func (rs *RigidStatic) ToRigidActor() RigidActor {
	return RigidActor{
		cRa: C.CPxRigidStatic_toCPxRigidActor(rs.cRs),
	}
}

func CreatePlane(p *Physics, plane *Plane, mat *Material) *RigidStatic {
	return &RigidStatic{
		cRs: C.CPxCreatePlane(p.cPhysics, plane.cP, mat.cM),
	}
}
