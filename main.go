package main

import (
	"fmt"

	"github.com/bloeys/physx-go/pgo"
)

func contactHandler(cph pgo.ContactPairHeader) {

	// pairs := cph.GetPairs()
	// for i := 0; i < len(pairs); i++ {

	// 	points := pairs[i].GetContactPoints()
	// 	for j := 0; j < pairs[i].GetContactPointCount(); j++ {
	// 		pos := points[j].GetPos()
	// 		println("Contact at pos:", pos.String())
	// 	}
	// }
}

func main() {

	const enablePvd = true

	f := pgo.CreateFoundation()
	println("foundation:", f)

	var pvd *pgo.Pvd
	if enablePvd {
		pvdTr := pgo.DefaultPvdSocketTransportCreate("127.0.0.1", 5425, 100000)
		println("Pvd transport:", pvdTr)

		pvd = pgo.CreatePvd(f)
		println("Pvd:", pvd)
		println("connected to PVD:", pvd.Connect(pvdTr, pgo.PvdInstrumentationFlag_eALL))
	}

	ts := pgo.NewTolerancesScale(1, 9.81)
	p := pgo.CreatePhysics(f, ts, false, pvd)
	println("Physics:", p)

	sd := pgo.NewSceneDesc(ts)
	sd.SetGravity(pgo.NewVec3(0, -9.8, 0))
	sd.SetCpuDispatcher(pgo.DefaultCpuDispatcherCreate(2, nil).ToCpuDispatcher())
	sd.SetOnContactCallback(contactHandler)

	scene := p.CreateScene(sd)
	println("Scene:", scene)

	if enablePvd {
		scenePvdClient := scene.GetScenePvdClient()
		println("ScenePvdClient:", scenePvdClient)

		scenePvdClient.SetScenePvdFlag(pgo.PvdSceneFlag_eTRANSMIT_CONSTRAINTS, true)
		scenePvdClient.SetScenePvdFlag(pgo.PvdSceneFlag_eTRANSMIT_CONTACTS, true)
		scenePvdClient.SetScenePvdFlag(pgo.PvdSceneFlag_eTRANSMIT_SCENEQUERIES, true)
		scenePvdClient.Release()
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
	println("Box 1 mass:", dynBox.GetMass())
	println("Box 2 mass:", dynBox2.GetMass())
	println("Sphere mass:", dynSphere.GetMass())
	println("Capsule mass:", dynCapsule.GetMass())

	println("Capsule linear damping A:", dynCapsule.GetLinearDamping())
	dynCapsule.SetLinearDamping(0.05)
	println("Capsule linear damping B:", dynCapsule.GetLinearDamping())

	//Run simulation
	// r := bufio.NewReader(os.Stdin)
	raycastBuffer := pgo.NewRaycastBuffer(1)
	defer raycastBuffer.Release()

	scene.SetScratchBuffer(4)
	for {
		scene.Collide(1 / 50.0)
		scene.FetchCollision(true)
		scene.Advance()
		scene.FetchResults(true)

		scene.RaycastWithHitBuffer(pgo.NewVec3(0, 0, 0), pgo.NewVec3(0, 1, 0), 9, raycastBuffer, 1)
		if raycastBuffer.HasBlock() {
			block := raycastBuffer.GetBlock()
			d := block.GetDistance()
			pos := block.GetPos()
			fmt.Printf("Raycast hit at dist (%v) and post %v\n", d, pos.String())
		}
		// fmt.Printf("\nRaycast hit: %v\n", rHit)
		// println("Press enter...")
		// r.ReadBytes('\n')
	}
}
