package main

import (
	"fmt"
	"rouge/internal/ecs"
	"rouge/internal/engine/objects"
	"rouge/internal/engine/systems/multiplayer"
	"time"
)

func main() {

	wrld := ecs.NewWorld()
	multiplayerSystem := multiplayer.NewNetworkingSystem(true)
	wrld.AddSystem(multiplayerSystem)

	wrld.AddEntity(objects.NewBlock(10, 0, 0))
	wrld.AddEntity(objects.NewBlock(0, 10, 0))
	wrld.AddEntity(objects.NewBlock(0, 0, 10))

	step := time.Millisecond * 16

	next := time.Now().Add(step)
	for {
		wrld.UpdateSystems(16.0)

		if time.Until(next) < -step {
			delaycnt := float32(-time.Until(next).Milliseconds()) / float32(step.Milliseconds())
			fmt.Printf("ERROR, SERVER CAN'T KEEP UP BY %v STEPS\n", delaycnt)
			next = next.Add(step * time.Duration(delaycnt+1))
		} else {
			time.Sleep(time.Until(next))
			next = next.Add(step)
		}
	}
}
