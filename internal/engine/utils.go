package engine

import (
	"rouge/internal/ecs"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func BootstrapWorldRenderRaylib(world *ecs.World) {
	world.StartUpdateFunc = func() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
	}

	world.EndUpdateFunc = func() {
		rl.EndDrawing()
	}
}

func SetupScreen() {
	screenWidth := int32(800)
	screenHeight := int32(450)

	rl.InitWindow(screenWidth, screenHeight, "game")
	rl.DisableCursor()

	rl.SetTargetFPS(60)
}
