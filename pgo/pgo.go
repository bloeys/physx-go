package pgo

/*
#cgo CFLAGS: -I physx-c
#cgo LDFLAGS: -L ./libs

// NOTE: If you change this update rigiddynamic.go as well
#cgo windows,amd64 LDFLAGS: -l physxc_checked_windows_amd64
#cgo windows,amd64,physx_release LDFLAGS: -l physxc_release_windows_amd64

#include <wrap.c>
#include <stdlib.h> //Needed for C.free

//simulation event callbacks forward declarations. Actual definitions MUST be in a go different file
void goOnContactCallback_cgo(void* pairHeader);
void goOnTriggerCallback_cgo(void* triggerPairs, CPxU32 count);
*/
import "C"
import (
	"runtime"
	"unsafe"

	"github.com/bloeys/gglm/gglm"
)

var (
	contactCallback func(ContactPairHeader) = func(ContactPairHeader) {}
	triggerCallback func([]TriggerPair)     = func([]TriggerPair) {}
)

type Foundation struct {
	cFoundation C.struct_CPxFoundation
}

func (f *Foundation) Release() {
	C.CPxFoundation_release(f.cFoundation)
}

func CreateFoundation() Foundation {

	f := Foundation{}
	f.cFoundation = C.CPxCreateFoundation()

	return f
}

type Pvd struct {
	cPvd C.struct_CPxPvd
}

func (p *Pvd) Connect(pvdTr PvdTransport, instFlag PvdInstrumentationFlag) bool {
	return bool(C.CPxPvd_connect(p.cPvd, pvdTr.cPvdTr, uint32(instFlag)))
}

func (p *Pvd) Release() {
	C.CPxPvd_release(p.cPvd)
}

func CreatePvd(f Foundation) *Pvd {

	p := &Pvd{}
	p.cPvd = C.CPxCreatePvd(f.cFoundation)

	return p
}

type PvdTransport struct {
	cPvdTr C.struct_CPxPvdTransport
}

func DefaultPvdSocketTransportCreate(host string, port, timeoutMillis int) PvdTransport {

	//This CString should NOT be freed because its stored internally. If this is freed connection to PVD will fail
	p := PvdTransport{}
	p.cPvdTr = C.CPxDefaultPvdSocketTransportCreate(C.CString(host), C.int(port), C.int(timeoutMillis))
	return p
}

type TolerancesScale struct {
	cTolScale C.struct_CPxTolerancesScale
}

func NewTolerancesScale(length, speed float32) TolerancesScale {
	return TolerancesScale{
		cTolScale: C.struct_CPxTolerancesScale{
			length: C.float(length),
			speed:  C.float(speed),
		},
	}
}

type Scene struct {
	cS C.struct_CPxScene
}

func (s *Scene) GetScenePvdClient() PvdSceneClient {
	return PvdSceneClient{
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

func (s *Scene) Raycast(origin, unitDir *Vec3, distance float32) (bool, RaycastBuffer) {

	rb := RaycastBuffer{}
	ret := C.CPxScene_raycast(s.cS, &origin.cV, &unitDir.cV, C.float(distance), &rb.cRb)

	return bool(ret), rb
}

func (s *Scene) RaycastWithHitBuffer(origin, unitDir *Vec3, distance float32, rb RaycastBuffer, touchesToRead uint) bool {

	ret := C.CPxScene_raycastWithHitBuffer(s.cS, &origin.cV, &unitDir.cV, C.float(distance), rb.cRb, C.uint(touchesToRead))
	return bool(ret)
}

type RaycastBuffer struct {
	cRb *C.struct_CPxRaycastBuffer
}

func (rb *RaycastBuffer) HasBlock() bool {
	return bool(rb.cRb.hasBlock)
}

func (rb *RaycastBuffer) GetBlock() RaycastHit {
	return RaycastHit{
		cRh: &rb.cRb.block,
	}
}

func (rb *RaycastBuffer) GetnbTouches() int {
	return int(rb.cRb.nbTouches)
}

func (rb *RaycastBuffer) GetTouches() []RaycastHit {

	hits := make([]RaycastHit, rb.cRb.nbTouches)
	touches := unsafe.Slice(rb.cRb.touches, rb.cRb.nbTouches)
	for i := 0; i < len(hits); i++ {
		hits[i].cRh = &touches[i]
	}

	return hits
}

func (rb *RaycastBuffer) Release() {
	C.CPxRaycastBuffer_release(rb.cRb)
}

func NewRaycastBuffer(maxTouches uint32) RaycastBuffer {

	rb := RaycastBuffer{
		cRb: C.NewCPxRaycastBufferWithAlloc(C.uint(maxTouches)),
	}

	return rb
}

type RaycastHit struct {
	cRh *C.struct_CPxRaycastHit
}

func (rh *RaycastHit) GetActor() RigidActor {
	return RigidActor{
		cRa: rh.cRh.actor,
	}
}

func (rh *RaycastHit) GetShape() Shape {
	return Shape{
		cShape: rh.cRh.shape,
	}
}

func (rh *RaycastHit) GetDistance() float32 {
	return float32(rh.cRh.distance)
}

func (rh *RaycastHit) GetFaceIndex() uint {
	return uint(rh.cRh.faceIndex)
}

func (rh *RaycastHit) GetHitFlags() HitFlag {
	return HitFlag(rh.cRh.flags)
}

func (rh *RaycastHit) GetNormal() gglm.Vec3 {
	return gglm.Vec3{
		Data: [3]float32{
			float32(rh.cRh.normal.x),
			float32(rh.cRh.normal.y),
			float32(rh.cRh.normal.z),
		},
	}
}

func (rh *RaycastHit) GetPos() gglm.Vec3 {
	return gglm.Vec3{
		Data: [3]float32{
			float32(rh.cRh.position.x),
			float32(rh.cRh.position.y),
			float32(rh.cRh.position.z),
		},
	}
}

func (rh *RaycastHit) GetUV() (float32, float32) {
	return float32(rh.cRh.u), float32(rh.cRh.v)
}

type Physics struct {
	cPhysics C.struct_CPxPhysics
}

func (p *Physics) CreateScene(sd SceneDesc) Scene {
	return Scene{
		cS: C.CPxPhysics_createScene(p.cPhysics, sd.cSD),
	}
}

func (p *Physics) CreateMaterial(staticFriction, dynamicFriction, restitution float32) Material {
	return Material{
		cM: C.CPxPhysics_createMaterial(p.cPhysics, C.float(staticFriction), C.float(dynamicFriction), C.float(restitution)),
	}
}

func (p *Physics) CreateRigidDynamic(tr *Transform) RigidDynamic {
	return RigidDynamic{
		cRd: C.CPxPhysics_createRigidDynamic(p.cPhysics, &tr.cT),
	}
}

func (p *Physics) CreateRigidStatic(tr *Transform) RigidStatic {
	return RigidStatic{
		cRs: C.CPxPhysics_createRigidStatic(p.cPhysics, &tr.cT),
	}
}

func (p *Physics) Release() {
	C.CPxPhysics_release(p.cPhysics)
}

func CreatePhysics(f Foundation, ts TolerancesScale, trackOutstandingAllocations bool, pvd *Pvd) Physics {

	p := Physics{}
	if pvd != nil {
		p.cPhysics = C.CPxCreatePhysics(f.cFoundation, ts.cTolScale, C._Bool(trackOutstandingAllocations), &pvd.cPvd)
	} else {
		p.cPhysics = C.CPxCreatePhysics(f.cFoundation, ts.cTolScale, C._Bool(trackOutstandingAllocations), nil)
	}

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
		cT: C.CPxShape_getLocalPose(s.cShape),
	}
}

func (s *Shape) SetLocalPose(tr *Transform) {
	C.CPxShape_setLocalPose(s.cShape, &tr.cT)
}

func (s *Shape) GetSimulationFilterData() FilterData {
	return FilterData{
		cFilterData: C.CPxShape_getSimulationFilterData(s.cShape),
	}
}

func (s *Shape) SetSimulationFilterData(fd *FilterData) {
	C.CPxShape_setSimulationFilterData(s.cShape, &fd.cFilterData)
}

func CreateExclusiveShape(rigidActor RigidActor, geom Geometry, mat Material, shapeFlags ShapeFlags) Shape {
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
		cV: C.struct_CPxVec3{
			x: C.float(x),
			y: C.float(y),
			z: C.float(z),
		},
	}
}

type CpuDispatcher struct {
	cCpuDisp C.struct_CPxCpuDispatcher
}

type DefaultCpuDispatcher struct {
	cDefCpuDisp C.struct_CPxDefaultCpuDispatcher
}

func (d *DefaultCpuDispatcher) ToCpuDispatcher() CpuDispatcher {
	return CpuDispatcher{cCpuDisp: C.CPxDefaultCpuDispatcher_toCPxCpuDispatcher(d.cDefCpuDisp)}
}

// DefaultCpuDispatcherCreate sets the number of threads used by physX.
// If affinityMasksPerThread is nil/zero then default masks are used, otherwise the size of the array
// must match the number of threads
func DefaultCpuDispatcherCreate(numThreads uint32, affinityMasksPerThread []uint32) DefaultCpuDispatcher {

	if len(affinityMasksPerThread) == 0 {
		return DefaultCpuDispatcher{
			cDefCpuDisp: C.CPxDefaultCpuDispatcherCreate(C.uint(numThreads), nil),
		}
	}

	arr := make([]C.uint, len(affinityMasksPerThread))
	for i := 0; i < len(arr); i++ {
		arr[i] = C.uint(affinityMasksPerThread[i])
	}

	return DefaultCpuDispatcher{
		cDefCpuDisp: C.CPxDefaultCpuDispatcherCreate(C.uint(numThreads), &arr[0]),
	}
}

type SceneDesc struct {
	cSD C.struct_CPxSceneDesc
}

func (sd *SceneDesc) SetGravity(v *Vec3) {
	C.CPxSceneDesc_set_gravity(sd.cSD, v.cV)
}

func (sd *SceneDesc) SetCpuDispatcher(cd CpuDispatcher) {
	C.CPxSceneDesc_set_cpuDispatcher(sd.cSD, cd.cCpuDisp)
}

// SetOnContactCallback sets the GLOBAL contact callback handler. We currently only supports 1 contact callback handler.
// Setting a contact callback handler overrides the previous one. Only the most recent one gets called.
func (sd *SceneDesc) SetOnContactCallback(cb func(ContactPairHeader)) {
	contactCallback = cb
	C.CPxSceneDesc_set_onContactCallback(sd.cSD, (C.CPxOnContactCallback)(unsafe.Pointer(C.goOnContactCallback_cgo)))
}

// SetOnTriggerCallback sets the GLOBAL trigger callback handler. We currently only supports 1 trugger callback handler.
// Setting a handler overrides the previous one. Only the most recent one gets called.
func (sd *SceneDesc) SetOnTriggerCallback(cb func([]TriggerPair)) {
	triggerCallback = cb
	C.CPxSceneDesc_set_onTriggerCallback(sd.cSD, (C.CPxOnTriggerCallback)(unsafe.Pointer(C.goOnTriggerCallback_cgo)))
}

func NewSceneDesc(ts TolerancesScale) SceneDesc {
	return SceneDesc{
		cSD: C.NewCPxSceneDesc(ts.cTolScale),
	}
}

type ContactPairHeader struct {
	cCPH *C.struct_CPxContactPairHeader
}

func (cph *ContactPairHeader) GetRigidActors() [2]RigidActor {
	return [2]RigidActor{
		{
			cRa: cph.cCPH.actors[0],
		},
		{
			cRa: cph.cCPH.actors[1],
		},
	}
}

func (cph *ContactPairHeader) GetFlags() ContactPairHeaderFlag {
	return ContactPairHeaderFlag(cph.cCPH.flags)
}

func (cph *ContactPairHeader) GetnbPairs() int {
	return int(cph.cCPH.nbPairs)
}

func (cph *ContactPairHeader) GetPairs() []ContactPair {

	contactPairs := make([]ContactPair, cph.cCPH.nbPairs)
	cPairs := unsafe.Slice(cph.cCPH.pairs, cph.cCPH.nbPairs)
	for i := 0; i < len(contactPairs); i++ {
		contactPairs[i].cCp = &cPairs[i]
	}

	return contactPairs
}

type ContactPair struct {
	cCp *C.struct_CPxContactPair
}

func (cp *ContactPair) GetFlags() ContactPairFlag {
	return ContactPairFlag(cp.cCp.flags)
}

func (cp *ContactPair) GetEvents() PairFlags {
	return PairFlags(cp.cCp.events)
}

func (cp *ContactPair) GetPatchCount() int {
	return int(cp.cCp.patchCount)
}

func (cp *ContactPair) GetContactPointCount() int {
	return int(cp.cCp.contactCount)
}

func (cp *ContactPair) GetContactPoints() []ContactPairPoint {

	ccps := make([]ContactPairPoint, cp.cCp.contactCount)
	extractedPoints := unsafe.Slice(cp.cCp.extractedContactPoints, cp.cCp.contactCount)
	for i := 0; i < len(extractedPoints); i++ {
		ccps[i].cCpp = &extractedPoints[i]
	}

	return ccps
}

type ContactPairPoint struct {
	cCpp *C.struct_CPxContactPairPoint
}

func (cpp *ContactPairPoint) GetPos() gglm.Vec3 {
	return gglm.Vec3{Data: [3]float32{
		float32(cpp.cCpp.position.x),
		float32(cpp.cCpp.position.y),
		float32(cpp.cCpp.position.z),
	}}
}

func (cpp *ContactPairPoint) GetImpulse() gglm.Vec3 {
	return gglm.Vec3{Data: [3]float32{
		float32(cpp.cCpp.impulse.x),
		float32(cpp.cCpp.impulse.y),
		float32(cpp.cCpp.impulse.z),
	}}
}

func (cpp *ContactPairPoint) GetNormal() gglm.Vec3 {
	return gglm.Vec3{Data: [3]float32{
		float32(cpp.cCpp.normal.x),
		float32(cpp.cCpp.normal.y),
		float32(cpp.cCpp.normal.z),
	}}
}

func (cpp *ContactPairPoint) GetSeparation() float32 {
	return float32(cpp.cCpp.separation)
}

func (cpp *ContactPairPoint) GetInternalFaceIndices() (float32, float32) {
	return float32(cpp.cCpp.internalFaceIndex0), float32(cpp.cCpp.internalFaceIndex1)
}

type TriggerPair struct {
	cTp *C.struct_CPxTriggerPair
}

func (tp *TriggerPair) TriggerShape() Shape {
	return Shape{
		cShape: tp.cTp.triggerShape,
	}
}

func (tp *TriggerPair) TriggerActor() RigidActor {
	return RigidActor{
		cRa: tp.cTp.triggerActor,
	}
}

func (tp *TriggerPair) OtherShape() Shape {
	return Shape{
		cShape: tp.cTp.otherShape,
	}
}

func (tp *TriggerPair) OtherActor() RigidActor {
	return RigidActor{
		cRa: tp.cTp.otherActor,
	}
}

func (tp *TriggerPair) Status() PairFlags {
	return PairFlags(tp.cTp.status)
}

func (tp *TriggerPair) Flags() TriggerPairFlag {
	return TriggerPairFlag(tp.cTp.flags)
}

type PvdSceneFlag uint32

const (
	PvdSceneFlag_eTRANSMIT_CONTACTS     PvdSceneFlag = (1 << 0) //Transmits contact stream to PVD.
	PvdSceneFlag_eTRANSMIT_SCENEQUERIES PvdSceneFlag = (1 << 1) //Transmits scene query stream to PVD.
	PvdSceneFlag_eTRANSMIT_CONSTRAINTS  PvdSceneFlag = (1 << 2) //Transmits constraints visualize stream to PVD.
)

type PvdSceneClient struct {
	cPvdSceneClient C.struct_CPxPvdSceneClient
}

func (p *PvdSceneClient) SetScenePvdFlag(flag PvdSceneFlag, value bool) {
	C.CPxPvdSceneClient_setScenePvdFlag(p.cPvdSceneClient, uint32(flag), C._Bool(value))
}

type Material struct {
	cM C.struct_CPxMaterial
}

type Plane struct {
	cP C.struct_CPxPlane
}

func NewPlane(nx, ny, nz, distance float32) *Plane {
	//If we don't keep a space between return and func definition this crashes?????
	return &Plane{
		cP: C.NewCPxPlane(C.float(nx), C.float(ny), C.float(nz), C.float(distance)),
	}
}

type Quat struct {
	cQ C.struct_CPxQuat
}

func NewQuat(angleRads, x, y, z float32) *Quat {
	return &Quat{
		cQ: C.NewCPxQuat(C.float(angleRads), C.float(x), C.float(y), C.float(z)),
	}
}

type Transform struct {
	cT C.struct_CPxTransform
}

func (t *Transform) Pos() Vec3 {
	return Vec3{cV: t.cT.p}
}

func (t *Transform) PosX() float32 {
	return float32(t.cT.p.x)
}

func (t *Transform) PosY() float32 {
	return float32(t.cT.p.y)
}

func (t *Transform) PosZ() float32 {
	return float32(t.cT.p.z)
}

func (t *Transform) SetPos(v *Vec3) {
	t.cT.p = v.cV
}

func (t *Transform) Rot() Quat {
	return Quat{cQ: t.cT.q}
}

func (t *Transform) RotX() float32 {
	return float32(t.cT.q.x)
}

func (t *Transform) RotY() float32 {
	return float32(t.cT.q.y)
}

func (t *Transform) RotZ() float32 {
	return float32(t.cT.q.z)
}

func (t *Transform) RotW() float32 {
	return float32(t.cT.q.w)
}

func (t *Transform) SetRot(r *Quat) {
	t.cT.q = r.cQ
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

func (sg *SphereGeometry) GetRadius() float32 {
	return float32(sg.cSg.radius)
}

func (sg *SphereGeometry) SetRadius(r float32) {
	sg.cSg.radius = C.float(r)
}

// struct CPxGeometry CPxSphereGeometry_toCPxGeometry(struct CPxSphereGeometry*);
func (sg SphereGeometry) ToGeometry() Geometry {
	return Geometry{
		cG: C.CPxSphereGeometry_toCPxGeometry(&sg.cSg),
	}
}

// struct CPxSphereGeometry NewCPxSphereGeometry(CPxReal radius);
func NewSphereGeometry(radius float32) SphereGeometry {
	return SphereGeometry{
		cSg: C.struct_CPxSphereGeometry{
			radius: C.float(radius),
		},
	}
}

type BoxGeometry struct {
	cBg C.struct_CPxBoxGeometry
}

// GetExtents returns the extents of each dimension of the box.
// An extent is half the length total length.
//
// For example, a cube of size 1x1x1 would have an extent of 0.5 on each side
func (bg *BoxGeometry) GetExtents() (ex, ey, ez float32) {
	return float32(bg.cBg.hx), float32(bg.cBg.hy), float32(bg.cBg.hz)
}

func (bg *BoxGeometry) SetExtents(ex, ey, ez float32) {
	bg.cBg.hx = C.float(ex)
	bg.cBg.hy = C.float(ey)
	bg.cBg.hz = C.float(ez)
}

func (bg BoxGeometry) ToGeometry() Geometry {
	return Geometry{
		cG: C.CPxBoxGeometry_toCPxGeometry(&bg.cBg),
	}
}

func NewBoxGeometry(hx, hy, hz float32) BoxGeometry {
	return BoxGeometry{
		cBg: C.struct_CPxBoxGeometry{
			hx: C.float(hx),
			hy: C.float(hy),
			hz: C.float(hz),
		},
	}
}

type CapsuleGeometry struct {
	cCg C.struct_CPxCapsuleGeometry
}

func (bg *CapsuleGeometry) GetParams() (radius, extent float32) {
	return float32(bg.cCg.radius), float32(bg.cCg.halfHeight)
}

func (bg *CapsuleGeometry) SetParams(radius, extent float32) {
	bg.cCg.radius = C.float(radius)
	bg.cCg.halfHeight = C.float(extent)
}

func (bg CapsuleGeometry) ToGeometry() Geometry {
	return Geometry{
		cG: C.CPxCapsuleGeometry_toCPxGeometry(&bg.cCg),
	}
}

func NewCapsuleGeometry(radius, halfHeight float32) CapsuleGeometry {
	return CapsuleGeometry{
		cCg: C.struct_CPxCapsuleGeometry{
			radius:     C.float(radius),
			halfHeight: C.float(halfHeight),
		},
	}
}

type Actor struct {
	cA C.struct_CPxActor
}

type RigidActor struct {
	cRa    C.struct_CPxRigidActor
	pinner runtime.Pinner
}

func (ra *RigidActor) SetSimFilterData(fd *FilterData) {
	C.CPxRigidActor_setSimFilterData(ra.cRa, fd.cFilterData)
}

// SetUserData sets the void* field on the rigid actor which can be used for any purpose.
// For example, it can be used to store an id or pointer that ties this rigid actor to some other object
//
// Note-1: The passed pointer will be stored in C and as such needs to be pinned, which this function will do.
// You can refer to this for notes on pinning and pointer rules: Refer to: https://pkg.go.dev/cmd/cgo#hdr-Passing_pointers
//
// Note-2: Since this RigidActor object is the one that pinned the user data, it MUST be kept alive at least until ClearUserData is used, at which point the data is unpinned and cleared.
// If this RigidActor object gets garabage collected before clear, the pinner will detect its getting collected with stuff still pinned (which is a leak) and will panic.
func (ra *RigidActor) SetUserData(userData unsafe.Pointer) {

	// Note: Do NOT use interfaces here, as we need to ensure the original value
	// pointed to is pinned, not the pointer to the interface (i.e. pinning the interface).
	// Better avoid crazy to debug issues

	// User data is a Go pointer stored in C/C++ code, and as such MUST be pinned
	// before that is done, and must be unpinned when no longer stored in C/C++.
	//
	// Here we assume every write is of a different object, and so we always unpin before storing
	// the new object.
	//
	// Refer to: https://pkg.go.dev/cmd/cgo#hdr-Passing_pointers
	ra.pinner.Unpin()
	ra.pinner.Pin(userData)

	C.CPxRigidActor_set_userData(ra.cRa, userData)
}

func (ra *RigidActor) GetUserData() unsafe.Pointer {
	return C.CPxRigidActor_get_userData(ra.cRa)
}

func (ra *RigidActor) ClearUserData() {
	ra.pinner.Unpin()
	C.CPxRigidActor_set_userData(ra.cRa, nil)
}

type RigidStatic struct {
	cRs C.struct_CPxRigidStatic
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

func CreatePlane(p Physics, plane *Plane, mat Material) RigidStatic {
	return RigidStatic{
		cRs: C.CPxCreatePlane(p.cPhysics, &plane.cP, mat.cM),
	}
}
