package main

import (
	"math/rand"
	"rouge/internal/ecs"
	"rouge/internal/engine"
	objects2d "rouge/internal/engine/objects/2d"
	systems2d "rouge/internal/engine/systems/2d"
	"time"

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

	landCollision := systems2d.NewLandCollisionSystem()
	wrld.AddSystem(landCollision)

	pcs := systems2d.NewPlayerControllerSystem()
	wrld.AddSystem(pcs)

	// multiplayerSystem := multiplayer.NewNetworkingSystem(false)
	// wrld.AddSystem(multiplayerSystem)

	// Add objects to world
	wrld.AddEntity(objects2d.New2DPlayer(ecs.ID(uuid.New().String())))
	wrld.AddEntity(objects2d.NewBlock2d(0, 0))

	// Generate random land
	chunkWidth := 5
	chunkHeight := 5

	rand.Seed(time.Now().UnixNano())
	randomData := func() []int {
		data := make([]int, chunkWidth*chunkHeight)
		for i := 0; i < chunkWidth*chunkHeight; i++ {
			data[i] = rand.Intn(2)
		}
		return data
	}

	chunkSizeWidth := 3
	chunkSizeHeight := 3
	for width := -chunkSizeWidth; width < chunkSizeWidth; width++ {
		for height := 1; height <= chunkSizeHeight; height++ {
			wrld.AddEntity(
				objects2d.NewLand(
					chunkWidth,
					chunkHeight,
					float32(width*chunkWidth*10),
					float32(height*chunkHeight*10),
					randomData(),
				),
			)
		}
	}

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
