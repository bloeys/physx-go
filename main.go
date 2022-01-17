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

	pMat := p.CreateMaterial(0.5, 0.5, 0.6)
	groundPlane := pgo.CreatePlane(p, pgo.NewPlane(0, 1, 0, 1), pMat)
	s.AddActor(groundPlane.ToActor())

	for {
		s.Simulate(1 / 60.0)
		a, b := s.FetchResults(true)
		println("a:", a, "; b:", b)
	}

	p.Release()
	pvd.Release()
	pvdTr.Release()
}
