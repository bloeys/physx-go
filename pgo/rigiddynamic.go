package pgo

/*
#cgo CFLAGS: -I physx-c
#cgo LDFLAGS: -L ./libs

// NOTE: If you change this update pgo.go as well
#cgo windows,amd64 LDFLAGS: -l physxc_checked_windows_amd64
#cgo windows,amd64,physx_release LDFLAGS: -l physxc_release_windows_amd64

#include <wrap.c>
#include <stdlib.h> //Needed for C.free
*/
import "C"

type RigidDynamic struct {
	cRd *C.struct_CPxRigidDynamic
}

func (rd *RigidDynamic) AddForce(force *Vec3, fmode ForceMode, autoAwake bool) {
	C.CPxRigidDynamic_addForce(rd.cRd, &force.cV, uint32(fmode), C._Bool(autoAwake))
}

func (rd *RigidDynamic) AddTorque(torque *Vec3, fmode ForceMode, autoAwake bool) {
	C.CPxRigidDynamic_addTorque(rd.cRd, &torque.cV, uint32(fmode), C._Bool(autoAwake))
}

func (rd *RigidDynamic) SetLinearDamping(damping float32) {
	C.CPxRigidDynamic_setLinearDamping(rd.cRd, C.float(damping))
}

func (rd *RigidDynamic) SetAngularDamping(damping float32) {
	C.CPxRigidDynamic_setAngularDamping(rd.cRd, C.float(damping))
}

func (rd *RigidDynamic) GetLinearDamping() float32 {
	return float32(C.CPxRigidDynamic_getLinearDamping(rd.cRd))
}

func (rd *RigidDynamic) GetAngularDamping() float32 {
	return float32(C.CPxRigidDynamic_getAngularDamping(rd.cRd))
}

func (rd *RigidDynamic) GetLinearVelocity() Vec3 {
	return Vec3{
		cV: C.CPxRigidDynamic_getLinearVelocity(rd.cRd),
	}
}

func (rd *RigidDynamic) GetAngularVelocity() Vec3 {
	return Vec3{
		cV: C.CPxRigidDynamic_getAngularVelocity(rd.cRd),
	}
}

func (rd *RigidDynamic) SetMass(mass float32) {
	C.CPxRigidDynamic_setMass(rd.cRd, C.float(mass))
}

func (rd *RigidDynamic) GetMass() float32 {
	return float32(C.CPxRigidDynamic_getMass(rd.cRd))
}

func (rd *RigidDynamic) SetRigidBodyFlag(flag RigidbodyFlags, value bool) {
	C.CPxRigidDynamic_setRigidBodyFlag(rd.cRd, uint32(flag), C._Bool(value))
}

func (rd *RigidDynamic) SetRigidBodyFlags(flags RigidbodyFlags) {
	C.CPxRigidDynamic_setRigidBodyFlags(rd.cRd, uint32(flags))
}

func (rd *RigidDynamic) GetRigidBodyFlags() RigidbodyFlags {
	return RigidbodyFlags(C.CPxRigidDynamic_getRigidBodyFlags(rd.cRd))
}

func (rd *RigidDynamic) SetRigidDynamicLockFlag(flag RigidDynamicLockFlags, value bool) {
	C.CPxRigidDynamic_setRigidDynamicLockFlag(rd.cRd, uint32(flag), C._Bool(value))
}

func (rd *RigidDynamic) SetRigidDynamicLockFlags(flags RigidDynamicLockFlags) {
	C.CPxRigidDynamic_setRigidDynamicLockFlags(rd.cRd, uint32(flags))
}

func (rd *RigidDynamic) GetRigidDynamicLockFlags() RigidDynamicLockFlags {
	return RigidDynamicLockFlags(C.CPxRigidDynamic_getRigidDynamicLockFlags(rd.cRd))
}

func (rd *RigidDynamic) PutToSleep() {
	C.CPxRigidDynamic_putToSleep(rd.cRd)
}

func (rd *RigidDynamic) GetGlobalPose() *Transform {
	return &Transform{
		cT: C.CPxRigidDynamic_getGlobalPose(rd.cRd),
	}
}

func (rd *RigidDynamic) SetGlobalPose(tr *Transform, autoAwake bool) {
	C.CPxRigidDynamic_setGlobalPose(rd.cRd, &tr.cT, C._Bool(autoAwake))
}

func (rd *RigidDynamic) GetCMassLocalPose() *Transform {
	return &Transform{
		cT: C.CPxRigidDynamic_getCMassLocalPose(rd.cRd),
	}
}

func (rd *RigidDynamic) SetCMassLocalPose(tr *Transform) {
	C.CPxRigidDynamic_setCMassLocalPose(rd.cRd, &tr.cT)
}

func (rd *RigidDynamic) ToActor() Actor {
	return Actor{
		cA: C.CPxRigidDynamic_toCPxActor(rd.cRd),
	}
}

func (rd *RigidDynamic) ToRigidActor() RigidActor {
	return RigidActor{
		cRa: C.CPxRigidDynamic_toCPxRigidActor(rd.cRd),
	}
}

func CreateDynamic(p *Physics, t *Transform, g *Geometry, m *Material, density float32, shapeOffset *Transform) *RigidDynamic {
	return &RigidDynamic{
		cRd: C.CPxCreateDynamic(p.cPhysics, &t.cT, g.cG, m.cM, C.float(density), &shapeOffset.cT),
	}
}
