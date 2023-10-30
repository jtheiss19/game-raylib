package engine

import (
	"github.com/jtheiss19/game-raylib/pkg/ecs"
	"github.com/sirupsen/logrus"

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

func RunWorld(world *ecs.World) {
	for !rl.WindowShouldClose() {
		delay := 1 / float32((rl.GetFPS())+1) * 1000
		if delay > 10000 {
			// The time between frames has gotten so high in ms the game
			// needs to preform only inplace logic updates to help reduce
			// the load. Rendering functions shouldn't run during this step
			// for example
			world.UpdateSystems(0)
			logrus.Error("World Can't Keep Up")
		} else {
			world.UpdateSystems(delay)
		}
	}

	rl.CloseWindow()
}

func SetupScreen() {
	screenWidth := int32(800)
	screenHeight := int32(450)

	rl.SetTraceLog(rl.LogWarning)
	rl.InitWindow(screenWidth, screenHeight, "game")
	rl.DisableCursor()

	rl.SetTargetFPS(60)
}
