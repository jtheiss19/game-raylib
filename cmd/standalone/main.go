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
	modelManager := systems3d.NewModelLoadingSystem()
	wrld.AddSystem(modelManager)

	// Add objects to world
	wrld.AddEntity(objects3d.New3DPlayer(ecs.ID(uuid.New().String()), 0, 1, 0))
	wrld.AddEntity(objects3d.NewGridBlock3d(5, 1, -1, objects3d.CRATE))
	wrld.AddEntity(objects3d.NewGridBlock3d(5, 1, 0, objects3d.CRATE))
	wrld.AddEntity(objects3d.NewGridBlock3d(5, 1, 1, objects3d.GRASS))

	// GameLoop
	for !rl.WindowShouldClose() {
		delay := 1 / float32((rl.GetFPS())+1) * 1000
		if delay > 10000 {
			// The time between frames has gotten so high in ms the game
			// needs to preform only inplace logic updates to help reduce
			// the load. Rendering functions shouldn't run during this step
			// for example
			wrld.UpdateSystems(0)
		} else {
			wrld.UpdateSystems(delay)
		}
	}

	rl.CloseWindow()

}
