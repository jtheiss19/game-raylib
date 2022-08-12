package main

import (
	"rouge/internal/ecs"
	"rouge/internal/engine/systems"
	"rouge/internal/engine/systems/multiplayer"

	rl "github.com/gen2brain/raylib-go/raylib"
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

	newRenderer := systems.NewRenderingSystem()
	wrld.AddSystem(newRenderer)

	multiplayerSystem := multiplayer.NewNetworkingSystem(false)
	wrld.AddSystem(multiplayerSystem)

	pcs := systems.NewPlayerControllerSystem()
	wrld.AddSystem(pcs)

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
