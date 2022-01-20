package main

import (
	"github.com/bloeys/physx-go/pgo"
)

func main() {
	f := pgo.CreateFoundation()
	println("foundation:", f)

	pvdTr := pgo.DefaultPvdSocketTransportCreate("127.0.0.1", 5425, 100000)
	println("Pvd transport:", pvdTr)

	pvd := pgo.CreatePvd(f)
	println("Pvd:", pvd)
	println("connect:", pvd.Connect(pvdTr, pgo.PvdInstrumentationFlag_eALL))

	ts := pgo.NewTolerancesScale(1, 9.81)
	p := pgo.CreatePhysics(f, ts, false, pvd)
	println("Physics:", p)

	sd := pgo.NewSceneDesc(ts)
	sd.SetGravity(pgo.NewVec3(0, -9.8, 0))
	sd.SetCpuDispatcher(pgo.DefaultCpuDispatcherCreate(2, 0).ToCpuDispatcher())

	s := p.CreateScene(sd)
	println("Scene:", s)

	scenePvdClient := s.GetScenePvdClient()
	println("ScenePvdClient:", scenePvdClient)

	scenePvdClient.SetScenePvdFlag(pgo.PvdSceneFlag_eTRANSMIT_CONSTRAINTS, true)
	scenePvdClient.SetScenePvdFlag(pgo.PvdSceneFlag_eTRANSMIT_CONTACTS, true)
	scenePvdClient.SetScenePvdFlag(pgo.PvdSceneFlag_eTRANSMIT_SCENEQUERIES, true)

	//Add plane
	pMat := p.CreateMaterial(0.5, 0.5, 0.6)
	groundPlane := pgo.CreatePlane(p, pgo.NewPlane(0, 1, 0, 0), pMat)
	s.AddActor(groundPlane.ToActor())

	//Add dynamic box
	v := pgo.NewVec3(0, 10, 0)
	q := pgo.NewQuat(0, 0, 1, 0)
	tr := pgo.NewTransform(v, q)

	qID := pgo.NewQuat(0, 0, 1, 0)
	shapeOffset := pgo.NewTransform(v, qID)

	box := pgo.NewBoxGeometry(0.5, 0.5, 0.5)
	dynBox := pgo.CreateDynamic(p, tr, box.ToGeometry(), pMat, 10, shapeOffset)
	s.AddActor(dynBox.ToActor())

	v = pgo.NewVec3(0.5, 12, 0)
	tr2 := pgo.NewTransform(v, qID)
	dynBox2 := pgo.CreateDynamic(p, tr2, box.ToGeometry(), pMat, 10, shapeOffset)
	s.AddActor(dynBox2.ToActor())

	for {
		s.Simulate(1 / 60.0)
		a, b := s.FetchResults(true)
		println("a:", a, "; b:", b)
	}

	p.Release()
	pvd.Release()
	pvdTr.Release()
}
