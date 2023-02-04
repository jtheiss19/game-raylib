package main

import (
	"rouge/internal/ecs"
	"rouge/internal/engine"
	objects2d "rouge/internal/engine/objects/2d"
	systems2d "rouge/internal/engine/systems/2d"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/google/uuid"
)

func main() {
	// Create actual usable screen
	engine.SetupScreen()

	// Create world
	wrld := ecs.NewWorld()
	engine.BootstrapWorldRenderRaylib(wrld)

	// Create and add systems
	entityRenderer := systems2d.NewRenderingSystem()
	wrld.AddSystem(entityRenderer)

	landRenderer := systems2d.NewLandRenderingSystem()
	wrld.AddSystem(landRenderer)

	pcs := systems2d.NewPlayerControllerSystem()
	wrld.AddSystem(pcs)

	// multiplayerSystem := multiplayer.NewNetworkingSystem(false)
	// wrld.AddSystem(multiplayerSystem)

	// Add objects to world
	wrld.AddEntity(objects2d.New2DPlayer(ecs.ID(uuid.New().String())))
	wrld.AddEntity(objects2d.NewLand(5, 10, 0, 0))

	// GameLoop
	for !rl.WindowShouldClose() {
		delay := 1 / rl.GetFPS() * 1000
		if delay > 10000 {
			wrld.UpdateSystems(0)
		} else {
			wrld.UpdateSystems(delay)
		}
	}

	rl.CloseWindow()

}
