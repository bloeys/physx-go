package main

import (
	"github.com/bloeys/physx-go/pgo"
)

func main() {
	f := pgo.CreateFoundation()
	defer f.Release()
	println("foundation:", f)

	pvd := pgo.CreatePvd(f)
	// defer pvd.Release()
	println("Pvd:", pvd)

	pvdTr := pgo.DefaultPvdSocketTransportCreate("127.0.0.1", 9876, 500)
	// defer pvdTr.Release()
	println("Pvd transport:", pvdTr)

	// for {
	// 	time.Sleep(1 / 60 * time.Second)
	// }
}
