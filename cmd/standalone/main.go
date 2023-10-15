package main

import (
	"rouge/internal/ecs"
	"rouge/internal/engine"
	objects3d "rouge/internal/engine/objects/3d"

	systems3d "rouge/internal/engine/systems/3d"

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
	gridRenderer := systems3d.NewGridRenderingSystem()
	wrld.AddSystem(gridRenderer)
	pcs := systems3d.NewGridPlayerControllerSystem()
	wrld.AddSystem(pcs)

	// Add objects to world
	wrld.AddEntity(objects3d.New3DPlayer(ecs.ID(uuid.New().String())))
	wrld.AddEntity(objects3d.NewGridBlock3d(10, 0, 0))

	// GameLoop
	for !rl.WindowShouldClose() {
		delay := float32(1 / (rl.GetFPS() + 1) * 1000)
		if delay > 10000 {
			wrld.UpdateSystems(0)
		} else {
			wrld.UpdateSystems(delay)
		}
	}

	rl.CloseWindow()

}
