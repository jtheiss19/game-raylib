package main

import (
	"github.com/jtheiss19/game-raylib/internal/ecs"
	"github.com/jtheiss19/game-raylib/internal/engine"
	components3d "github.com/jtheiss19/game-raylib/internal/engine/components/3d"
	objects3d "github.com/jtheiss19/game-raylib/internal/engine/objects/3d"

	systems3d "github.com/jtheiss19/game-raylib/internal/engine/systems/3d"

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
	chunkRenderer := systems3d.NewChunkRenderingSystem()
	wrld.AddSystem(chunkRenderer)
	Renderer := systems3d.NewRenderingSystem()
	wrld.AddSystem(Renderer)
	pcs := systems3d.NewGridPlayerControllerSystem()
	wrld.AddSystem(pcs)
	modelManager := systems3d.NewModelLoadingSystem()
	wrld.AddSystem(modelManager)

	// Add objects to world
	wrld.AddEntity(objects3d.New3DPlayer(ecs.ID(uuid.New().String()), 0, 1, 0))
	wrld.AddEntity(objects3d.NewBlock3d(5, 1, -1, components3d.CRATE_TEX, 1))
	wrld.AddEntity(objects3d.NewBlock3d(5, 1, 0, components3d.CRATE_TEX, 1))
	wrld.AddEntity(objects3d.NewBlock3d(5, 1, 1, components3d.IMAGE_TEX, 9))
	wrld.AddEntity(objects3d.NewChunk(0, 0, 0))

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
