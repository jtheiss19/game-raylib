package main

import (
	"rouge/internal/ecs"
	objects2d "rouge/internal/engine/objects/2d"
	systems2d "rouge/internal/engine/systems/2d"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/google/uuid"
)

func setupScreen() {
	screenWidth := int32(800)
	screenHeight := int32(450)

	rl.InitWindow(screenWidth, screenHeight, "game")

	rl.SetTargetFPS(60)
}

func main() {
	setupScreen()

	wrld := ecs.NewWorld()

	newRenderer := systems2d.NewRenderingSystem()
	wrld.AddSystem(newRenderer)

	pcs := systems2d.NewPlayerControllerSystem()
	wrld.AddSystem(pcs)

	// multiplayerSystem := multiplayer.NewNetworkingSystem(false)
	// wrld.AddSystem(multiplayerSystem)

	wrld.AddEntity(objects2d.NewBlock2d(50, 0))
	wrld.AddEntity(objects2d.NewBlock2d(0, 50))
	wrld.AddEntity(objects2d.NewBlock2d(0, 0))

	wrld.AddEntity(objects2d.New2DPlayer(ecs.ID(uuid.New().String())))

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
