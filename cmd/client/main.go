package main

import (
	"rouge/internal/ecs"
	"rouge/internal/engine/systems"

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

	multiplayerSystem := systems.NewNetworkingSystem(false)
	wrld.AddSystem(multiplayerSystem)

	pcs := systems.NewPlayerControllerSystem()
	wrld.AddSystem(pcs)

	for !rl.WindowShouldClose() {
		wrld.UpdateSystems(0)
	}

	rl.CloseWindow()

}
