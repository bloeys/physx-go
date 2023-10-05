package main

import (
	"fmt"
	"unsafe"

	"github.com/bloeys/physx-go/pgo"
)

func contactHandler(cph pgo.ContactPairHeader) {

	// ra1 := cph.GetRigidActors()[0]
	// ra2 := cph.GetRigidActors()[1]
	// fmt.Printf("Collision! User data 1: %v; User data 2: %v\n", ra1.GetUserData(), ra2.GetUserData())

	// pairs := cph.GetPairs()
	// for i := 0; i < len(pairs); i++ {

	// 	points := pairs[i].GetContactPoints()
	// 	for j := 0; j < pairs[i].GetContactPointCount(); j++ {
	// 		pos := points[j].GetPos()
	// 		fmt.Println("Contact at pos:", pos.String())
	// 	}
	// }
}

func main() {

	f := pgo.CreateFoundation()
	fmt.Println("foundation:", f)

	var pvd *pgo.Pvd
	if pgo.PvdSupported {
		pvdTr := pgo.DefaultPvdSocketTransportCreate("127.0.0.1", 5425, 100000)
		fmt.Println("Pvd transport:", pvdTr)

		pvd = pgo.CreatePvd(f)
		fmt.Println("Pvd:", pvd)
		fmt.Println("connected to PVD:", pvd.Connect(pvdTr, pgo.PvdInstrumentationFlag_eALL))
	}

	ts := pgo.NewTolerancesScale(1, 9.81)
	p := pgo.CreatePhysics(f, ts, false, pvd)
	fmt.Println("Physics:", p)

	sd := pgo.NewSceneDesc(ts)
	sd.SetGravity(pgo.NewVec3(0, -9.8, 0))

	defaultCpuDispatcher := pgo.DefaultCpuDispatcherCreate(2, nil)
	sd.SetCpuDispatcher(defaultCpuDispatcher.ToCpuDispatcher())
	sd.SetOnContactCallback(contactHandler)

	scene := p.CreateScene(sd)
	fmt.Println("Scene:", scene)

	if pgo.PvdSupported {
		scenePvdClient := scene.GetScenePvdClient()
		fmt.Println("ScenePvdClient:", scenePvdClient)

		scenePvdClient.SetScenePvdFlag(pgo.PvdSceneFlag_eTRANSMIT_CONSTRAINTS, true)
		scenePvdClient.SetScenePvdFlag(pgo.PvdSceneFlag_eTRANSMIT_CONTACTS, true)
		scenePvdClient.SetScenePvdFlag(pgo.PvdSceneFlag_eTRANSMIT_SCENEQUERIES, true)
	}

	//Add plane
	pMat := p.CreateMaterial(0.5, 0.5, 0.6)
	groundPlane := pgo.CreatePlane(p, pgo.NewPlane(0, 1, 0, 0), pMat)
	scene.AddActor(groundPlane.ToActor())

	//W0/W1 are filter groups the shape belongs to, and W2/W3 are a filter group mask
	fd := pgo.NewFilterData(1, 1, 1, 1)

	//Add box1
	v := pgo.NewVec3(0, 10, 0)
	q := pgo.NewQuat(0, 0, 1, 0)
	tr := pgo.NewTransform(v, q)

	qID := pgo.NewQuat(0, 0, 1, 0)
	shapeOffset := pgo.NewTransform(v, qID)

	box := pgo.NewBoxGeometry(0.5, 0.5, 0.5)
	dynBox := pgo.CreateDynamic(p, tr, box.ToGeometry(), pMat, 1, shapeOffset)

	ra := dynBox.ToRigidActor()
	ra.SetSimFilterData(&fd)
	scene.AddActor(dynBox.ToActor())

	//Add box2
	v = pgo.NewVec3(0.5, 12, 0)
	tr2 := pgo.NewTransform(v, qID)
	dynBox2 := pgo.CreateDynamic(p, tr2, box.ToGeometry(), pMat, 1, shapeOffset)

	ra = dynBox2.ToRigidActor()
	ra.SetSimFilterData(&fd)
	scene.AddActor(dynBox2.ToActor())

	//Add sphere
	v = pgo.NewVec3(0, 16, 0)
	tr3 := pgo.NewTransform(v, qID)
	dynSphere := pgo.CreateDynamic(p, tr3, pgo.NewSphereGeometry(3).ToGeometry(), pMat, 1, shapeOffset)

	ra = dynSphere.ToRigidActor()
	ra.SetSimFilterData(&fd)
	scene.AddActor(dynSphere.ToActor())

	//Add capsule
	v = pgo.NewVec3(0, 20, 0)
	tr4 := pgo.NewTransform(v, qID)
	dynCapsule := pgo.CreateDynamic(p, tr4, pgo.NewCapsuleGeometry(0.25, 0.5).ToGeometry(), pMat, 1, shapeOffset)
	ra = dynCapsule.ToRigidActor()
	ra.SetSimFilterData(&fd)
	scene.AddActor(dynCapsule.ToActor())

	//Add compound shape
	dynComp := p.CreateRigidDynamic(pgo.NewTransform(pgo.NewVec3(2.5, 35, 0), qID))

	pgo.CreateExclusiveShape(dynComp.ToRigidActor(), pgo.NewBoxGeometry(10, 0.1, 0.1).ToGeometry(), pMat, pgo.ShapeFlags_eSCENE_QUERY_SHAPE|pgo.ShapeFlags_eSIMULATION_SHAPE|pgo.ShapeFlags_eVISUALIZATION)

	someShape := pgo.CreateExclusiveShape(dynComp.ToRigidActor(), pgo.NewSphereGeometry(2).ToGeometry(), pMat, pgo.ShapeFlags_eSCENE_QUERY_SHAPE|pgo.ShapeFlags_eSIMULATION_SHAPE|pgo.ShapeFlags_eVISUALIZATION)
	someShape.SetLocalPose(pgo.NewTransform(pgo.NewVec3(5, 0, 0), qID))

	someShape = pgo.CreateExclusiveShape(dynComp.ToRigidActor(), pgo.NewSphereGeometry(2).ToGeometry(), pMat, pgo.ShapeFlags_eSCENE_QUERY_SHAPE|pgo.ShapeFlags_eSIMULATION_SHAPE|pgo.ShapeFlags_eVISUALIZATION)
	someShape.SetLocalPose(pgo.NewTransform(pgo.NewVec3(-5, 0, 0), qID))

	ra = dynComp.ToRigidActor()
	ra.SetSimFilterData(&fd)
	scene.AddActor(dynComp.ToActor())

	//Make some changes and print info
	dynSphere.SetMass(1)
	dynCapsule.SetMass(1)
	fmt.Println("Box 1 mass:", dynBox.GetMass())
	fmt.Println("Box 2 mass:", dynBox2.GetMass())
	fmt.Println("Sphere mass:", dynSphere.GetMass())
	fmt.Println("Capsule mass:", dynCapsule.GetMass())

	fmt.Println("Capsule linear damping A:", dynCapsule.GetLinearDamping())
	dynCapsule.SetLinearDamping(0.05)
	fmt.Println("Capsule linear damping B:", dynCapsule.GetLinearDamping())

	//Run simulation
	// r := bufio.NewReader(os.Stdin)
	raycastBuffer := pgo.NewRaycastBuffer(1)
	defer raycastBuffer.Release()

	// Example of correct usage of user data
	x := new(int64)
	*x = 1095
	ra.SetUserData(unsafe.Pointer(x))
	z := (*int64)(ra.GetUserData())
	fmt.Println("User data:", *z)

	// The rigid actor might get garbage collected after this point (as its no longer used), but that will now cause
	// a memory leak and a crash since the user data got pinned (i.e. GC will not move or free it) when we used SetUserData above, but can no longer be unpinned as the runtime.Pinner object will get garbage collected with the rigid actor.
	//
	// To solve this we have 2 options:
	//   1. The pinner is unpinned before it is garbage collected by calling ClearUserData
	//   2. The object holding the active pinner (the rigid actor on which SetUserData was used) must remain alive (e.g. by pinning it, putting it in file scope, in a long lived object etc)
	ra.ClearUserData()

	scene.SetScratchBuffer(4)
	for {
		scene.Collide(1 / 50.0)
		scene.FetchCollision(true)
		scene.Advance()
		scene.FetchResults(true)

		scene.RaycastWithHitBuffer(pgo.NewVec3(0, 0, 0), pgo.NewVec3(0, 1, 0), 9, raycastBuffer, 1)
		if raycastBuffer.HasBlock() {
			// block := raycastBuffer.GetBlock()
			// d := block.GetDistance()
			// pos := block.GetPos()
			// fmt.Printf("Raycast hit at dist (%v) and post %v\n", d, pos.String())
		}
		// fmt.Printf("\nRaycast hit: %v\n", rHit)
		// fmt.Println("Press enter...")
		// r.ReadBytes('\n')
	}
}
