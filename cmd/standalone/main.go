package main

import (
	"github.com/jtheiss19/game-raylib/internal/ecs"
	"github.com/jtheiss19/game-raylib/internal/engine"
	components3d "github.com/jtheiss19/game-raylib/internal/engine/components/3d"
	objects3d "github.com/jtheiss19/game-raylib/internal/engine/objects/3d"

	systems3d "github.com/jtheiss19/game-raylib/internal/engine/systems/3d"

	"github.com/google/uuid"
)

func main() {
	// Create actual usable screen
	engine.SetupScreen()

	// Create world
	world := ecs.NewWorld()
	engine.BootstrapWorldRenderRaylib(world)

	// Create and add systems
	chunkRenderer := systems3d.NewChunkRenderingSystem()
	world.AddSystem(chunkRenderer)
	Renderer := systems3d.NewRenderingSystem()
	world.AddSystem(Renderer)
	pcs := systems3d.NewPlayerControllerSystem()
	world.AddSystem(pcs)
	modelManager := systems3d.NewModelLoadingSystem()
	world.AddSystem(modelManager)

	// Add objects to world
	world.AddEntity(objects3d.New3DPlayer(ecs.ID(uuid.New().String()), 0, 1, 0))
	world.AddEntity(objects3d.NewBlock3d(5, 1, -1, components3d.CRATE_TEX, 1))
	world.AddEntity(objects3d.NewBlock3d(5, 1, 0, components3d.CRATE_TEX, 1))
	world.AddEntity(objects3d.NewBlock3d(5, 1, 1, components3d.IMAGE_TEX, 9))
	world.AddEntity(objects3d.NewChunk(0, 0, 0))

	// GameLoop
	engine.RunWorld(world)

}
