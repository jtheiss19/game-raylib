package main

import (
	"github.com/jtheiss19/game-raylib/pkg/ecs"
	"github.com/jtheiss19/game-raylib/pkg/engine"
	components3d "github.com/jtheiss19/game-raylib/pkg/engine/components/3d"
	objects3d "github.com/jtheiss19/game-raylib/pkg/engine/objects/3d"

	systems3d "github.com/jtheiss19/game-raylib/pkg/engine/systems/3d"

	ui "github.com/jtheiss19/game-raylib/pkg/engine/systems/ui"

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
	uiManager := ui.NewUIRenderingSystem()
	world.AddSystem(uiManager)
	collisionSystem := systems3d.NewCollisionSystem()
	world.AddSystem(collisionSystem)

	// Add objects to world
	world.AddEntity(objects3d.New3DPlayer(ecs.ID(uuid.New().String()), 0, 3, 0))

	// Draw Floor
	world.AddEntity(objects3d.NewChunk(0, -256, 16, 16, 16, 256, components3d.IMAGE_TEX, 4))
	world.AddEntity(objects3d.NewChunk(0, -256, 0, 16, 16, 256, components3d.IMAGE_TEX, 4))
	world.AddEntity(objects3d.NewChunk(0, -256, -16, 16, 16, 256, components3d.IMAGE_TEX, 4))
	// world.AddEntity(objects3d.NewChunk(16, -256, 0, 16, 16, 256, components3d.IMAGE_TEX, 4))

	//world.AddEntity(objects3d.NewBlock3d(6, 1, 0, components3d.IMAGE_TEX, components3d.CUBE_OBJ, 4))

	// Draw UI
	//world.AddEntity(objectsui.NewBlockPlacementUI(20, 40, pcs))

	// GameLoop
	engine.RunWorld(world)
}
