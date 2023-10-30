package main

import (
	"github.com/jtheiss19/game-raylib/pkg/ecs"
	"github.com/jtheiss19/game-raylib/pkg/engine"
	components3d "github.com/jtheiss19/game-raylib/pkg/engine/components/3d"
	objects3d "github.com/jtheiss19/game-raylib/pkg/engine/objects/3d"

	systems3d "github.com/jtheiss19/game-raylib/pkg/engine/systems/3d"

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
	pcs := systems3d.NewGridPlayerControllerSystem()
	world.AddSystem(pcs)
	modelManager := systems3d.NewModelLoadingSystem()
	world.AddSystem(modelManager)

	// Add objects to world
	world.AddEntity(objects3d.New3DPlayer(ecs.ID(uuid.New().String()), 0, 2, 0))

	// Draw House
	DrawHouse(world)

	// Draw Floor
	world.AddEntity(objects3d.NewChunk(0, 0, -5, 10, 10, 10, components3d.IMAGE_TEX, 4))

	// GameLoop
	engine.RunWorld(world)
}

func DrawHouse(world *ecs.World) {
	// Draw Row 1
	world.AddEntity(objects3d.NewBlock3d(5, 1, -1, components3d.IMAGE_TEX, components3d.CUBE_OBJ, 12))
	world.AddEntity(objects3d.NewBlock3d(5, 1, -2, components3d.IMAGE_TEX, components3d.CUBE_OBJ, 12))
	world.AddEntity(objects3d.NewBlock3d(5, 1, 1, components3d.IMAGE_TEX, components3d.CUBE_OBJ, 12))
	world.AddEntity(objects3d.NewBlock3d(5, 1, 2, components3d.IMAGE_TEX, components3d.CUBE_OBJ, 12))

	world.AddEntity(objects3d.NewBlock3d(5, 1, 3, components3d.IMAGE_TEX, components3d.CUBE_OBJ, 12))
	world.AddEntity(objects3d.NewBlock3d(6, 1, 3, components3d.IMAGE_TEX, components3d.CUBE_OBJ, 12))
	world.AddEntity(objects3d.NewBlock3d(7, 1, 3, components3d.IMAGE_TEX, components3d.CUBE_OBJ, 12))
	world.AddEntity(objects3d.NewBlock3d(8, 1, 3, components3d.IMAGE_TEX, components3d.CUBE_OBJ, 12))

	world.AddEntity(objects3d.NewBlock3d(5, 1, -3, components3d.IMAGE_TEX, components3d.CUBE_OBJ, 12))
	world.AddEntity(objects3d.NewBlock3d(6, 1, -3, components3d.IMAGE_TEX, components3d.CUBE_OBJ, 12))
	world.AddEntity(objects3d.NewBlock3d(7, 1, -3, components3d.IMAGE_TEX, components3d.CUBE_OBJ, 12))
	world.AddEntity(objects3d.NewBlock3d(8, 1, -3, components3d.IMAGE_TEX, components3d.CUBE_OBJ, 12))

	world.AddEntity(objects3d.NewBlock3d(8, 1, -2, components3d.IMAGE_TEX, components3d.CUBE_OBJ, 12))
	world.AddEntity(objects3d.NewBlock3d(8, 1, -1, components3d.IMAGE_TEX, components3d.CUBE_OBJ, 12))
	world.AddEntity(objects3d.NewBlock3d(8, 1, -0, components3d.IMAGE_TEX, components3d.CUBE_OBJ, 12))
	world.AddEntity(objects3d.NewBlock3d(8, 1, 1, components3d.IMAGE_TEX, components3d.CUBE_OBJ, 12))
	world.AddEntity(objects3d.NewBlock3d(8, 1, 2, components3d.IMAGE_TEX, components3d.CUBE_OBJ, 12))

	// Draw Row 2
	world.AddEntity(objects3d.NewBlock3d(5, 2, -1, components3d.IMAGE_TEX, components3d.CUBE_OBJ, 12))
	world.AddEntity(objects3d.NewBlock3d(5, 2, -2, components3d.IMAGE_TEX, components3d.CUBE_OBJ, 12))
	world.AddEntity(objects3d.NewBlock3d(5, 2, 1, components3d.IMAGE_TEX, components3d.CUBE_OBJ, 12))
	world.AddEntity(objects3d.NewBlock3d(5, 2, 2, components3d.IMAGE_TEX, components3d.CUBE_OBJ, 12))

	world.AddEntity(objects3d.NewBlock3d(5, 2, 3, components3d.IMAGE_TEX, components3d.CUBE_OBJ, 12))
	world.AddEntity(objects3d.NewBlock3d(6, 2, 3, components3d.IMAGE_TEX, components3d.CUBE_OBJ, 12))
	world.AddEntity(objects3d.NewBlock3d(7, 2, 3, components3d.IMAGE_TEX, components3d.CUBE_OBJ, 12))
	world.AddEntity(objects3d.NewBlock3d(8, 2, 3, components3d.IMAGE_TEX, components3d.CUBE_OBJ, 12))

	world.AddEntity(objects3d.NewBlock3d(5, 2, -3, components3d.IMAGE_TEX, components3d.CUBE_OBJ, 12))
	world.AddEntity(objects3d.NewBlock3d(6, 2, -3, components3d.IMAGE_TEX, components3d.CUBE_OBJ, 12))
	world.AddEntity(objects3d.NewBlock3d(7, 2, -3, components3d.IMAGE_TEX, components3d.CUBE_OBJ, 12))
	world.AddEntity(objects3d.NewBlock3d(8, 2, -3, components3d.IMAGE_TEX, components3d.CUBE_OBJ, 12))

	world.AddEntity(objects3d.NewBlock3d(8, 2, -2, components3d.IMAGE_TEX, components3d.CUBE_OBJ, 12))
	world.AddEntity(objects3d.NewBlock3d(8, 2, -1, components3d.IMAGE_TEX, components3d.CUBE_OBJ, 12))
	world.AddEntity(objects3d.NewBlock3d(8, 2, -0, components3d.IMAGE_TEX, components3d.CUBE_OBJ, 12))
	world.AddEntity(objects3d.NewBlock3d(8, 2, 1, components3d.IMAGE_TEX, components3d.CUBE_OBJ, 12))
	world.AddEntity(objects3d.NewBlock3d(8, 2, 2, components3d.IMAGE_TEX, components3d.CUBE_OBJ, 12))

	// Draw Row 3
	world.AddEntity(objects3d.NewBlock3d(5, 3, -3, components3d.IMAGE_TEX, components3d.CUBE_OBJ, 14))
	world.AddEntity(objects3d.NewBlock3d(5, 3, -2, components3d.IMAGE_TEX, components3d.CUBE_OBJ, 14))
	world.AddEntity(objects3d.NewBlock3d(5, 3, -1, components3d.IMAGE_TEX, components3d.CUBE_OBJ, 14))
	world.AddEntity(objects3d.NewBlock3d(5, 3, -0, components3d.IMAGE_TEX, components3d.CUBE_OBJ, 14))
	world.AddEntity(objects3d.NewBlock3d(5, 3, 1, components3d.IMAGE_TEX, components3d.CUBE_OBJ, 14))
	world.AddEntity(objects3d.NewBlock3d(5, 3, 2, components3d.IMAGE_TEX, components3d.CUBE_OBJ, 14))
	world.AddEntity(objects3d.NewBlock3d(5, 3, 3, components3d.IMAGE_TEX, components3d.CUBE_OBJ, 14))

	world.AddEntity(objects3d.NewBlock3d(6, 3, -3, components3d.IMAGE_TEX, components3d.CUBE_OBJ, 14))
	world.AddEntity(objects3d.NewBlock3d(6, 3, -2, components3d.IMAGE_TEX, components3d.CUBE_OBJ, 14))
	world.AddEntity(objects3d.NewBlock3d(6, 3, -1, components3d.IMAGE_TEX, components3d.CUBE_OBJ, 14))
	world.AddEntity(objects3d.NewBlock3d(6, 3, -0, components3d.IMAGE_TEX, components3d.CUBE_OBJ, 14))
	world.AddEntity(objects3d.NewBlock3d(6, 3, 1, components3d.IMAGE_TEX, components3d.CUBE_OBJ, 14))
	world.AddEntity(objects3d.NewBlock3d(6, 3, 2, components3d.IMAGE_TEX, components3d.CUBE_OBJ, 14))
	world.AddEntity(objects3d.NewBlock3d(6, 3, 3, components3d.IMAGE_TEX, components3d.CUBE_OBJ, 14))

	world.AddEntity(objects3d.NewBlock3d(7, 3, -3, components3d.IMAGE_TEX, components3d.CUBE_OBJ, 14))
	world.AddEntity(objects3d.NewBlock3d(7, 3, -2, components3d.IMAGE_TEX, components3d.CUBE_OBJ, 14))
	world.AddEntity(objects3d.NewBlock3d(7, 3, -1, components3d.IMAGE_TEX, components3d.CUBE_OBJ, 14))
	world.AddEntity(objects3d.NewBlock3d(7, 3, -0, components3d.IMAGE_TEX, components3d.CUBE_OBJ, 14))
	world.AddEntity(objects3d.NewBlock3d(7, 3, 1, components3d.IMAGE_TEX, components3d.CUBE_OBJ, 14))
	world.AddEntity(objects3d.NewBlock3d(7, 3, 2, components3d.IMAGE_TEX, components3d.CUBE_OBJ, 14))
	world.AddEntity(objects3d.NewBlock3d(7, 3, 3, components3d.IMAGE_TEX, components3d.CUBE_OBJ, 14))

	world.AddEntity(objects3d.NewBlock3d(8, 3, -3, components3d.IMAGE_TEX, components3d.CUBE_OBJ, 14))
	world.AddEntity(objects3d.NewBlock3d(8, 3, -2, components3d.IMAGE_TEX, components3d.CUBE_OBJ, 14))
	world.AddEntity(objects3d.NewBlock3d(8, 3, -1, components3d.IMAGE_TEX, components3d.CUBE_OBJ, 14))
	world.AddEntity(objects3d.NewBlock3d(8, 3, -0, components3d.IMAGE_TEX, components3d.CUBE_OBJ, 14))
	world.AddEntity(objects3d.NewBlock3d(8, 3, 1, components3d.IMAGE_TEX, components3d.CUBE_OBJ, 14))
	world.AddEntity(objects3d.NewBlock3d(8, 3, 2, components3d.IMAGE_TEX, components3d.CUBE_OBJ, 14))
	world.AddEntity(objects3d.NewBlock3d(8, 3, 3, components3d.IMAGE_TEX, components3d.CUBE_OBJ, 14))

}
