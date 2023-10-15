package main

import (
	"fmt"
	"time"

	"github.com/jtheiss19/game-raylib/internal/ecs"
	objects2d "github.com/jtheiss19/game-raylib/internal/engine/objects/2d"
	"github.com/jtheiss19/game-raylib/internal/engine/systems/multiplayer"
)

func main() {

	wrld := ecs.NewWorld()
	multiplayerSystem := multiplayer.NewNetworkingSystem(true)
	wrld.AddSystem(multiplayerSystem)

	wrld.AddEntity(objects2d.NewBlock2d(50, 0))
	wrld.AddEntity(objects2d.NewBlock2d(0, 50))
	wrld.AddEntity(objects2d.NewBlock2d(0, 0))

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
