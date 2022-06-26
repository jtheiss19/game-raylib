package systems

import (
	"rouge/internal/ecs"
	"rouge/internal/engine/components"

	"github.com/sirupsen/logrus"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type RenderingSystem struct {
	*ecs.BaseSystem
}

func NewRenderingSystem() *RenderingSystem {
	return &RenderingSystem{
		BaseSystem: &ecs.BaseSystem{},
	}
}

// Comps
type RequiredRenderingSystemComps struct {
	Model  []*RequireModel
	Camera []*RequireCamera
}

type RequireModel struct {
	Transformation *components.TransformationComponent
}

type RequireCamera struct {
	Player         *components.PlayerComponent
	Camera         *components.CameraComponent
	Transformation *components.TransformationComponent
}

func (ts *RenderingSystem) GetRequiredComponents() interface{} {
	return &RequiredRenderingSystemComps{
		Model: []*RequireModel{{
			Transformation: &components.TransformationComponent{},
		}},
		Camera: []*RequireCamera{{
			Camera:         &components.CameraComponent{},
			Transformation: &components.TransformationComponent{},
		}},
	}
}

// Functionality
func (ts *RenderingSystem) Update(dt float32) {
	entities, ok := ts.TrackedEntities.(*RequiredRenderingSystemComps)
	if !ok {
		logrus.Error("could not update system, bad tracked entities")
		return
	}

	if len(entities.Camera) > 0 {

		camera := entities.Camera[0].Camera.Camera

		// Draw
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		rl.BeginMode3D(*camera)

		rl.DrawPlane(rl.NewVector3(0.0, 0.0, 0.0), rl.NewVector2(32.0, 32.0), rl.LightGray) // Draw ground
		rl.DrawCube(rl.NewVector3(-16.0, 2.5, 0.0), 1.0, 5.0, 32.0, rl.Blue)                // Draw a blue wall
		rl.DrawCube(rl.NewVector3(16.0, 2.5, 0.0), 1.0, 5.0, 32.0, rl.Lime)                 // Draw a green wall
		rl.DrawCube(rl.NewVector3(0.0, 2.5, 16.0), 32.0, 5.0, 1.0, rl.Gold)                 // Draw a yellow wall

		// Draw some cubes around
		for i := 0; i < maxColumns; i++ {
			rl.DrawCube(positions[i], 2.0, heights[i], 2.0, colors[i])
			rl.DrawCubeWires(positions[i], 2.0, heights[i], 2.0, rl.Maroon)
		}

		for _, entity := range entities.Model {
			rl.DrawCubeV(entity.Transformation.Position, entity.Transformation.Scale, rl.Green)
		}

		rl.EndMode3D()

		rl.DrawRectangle(10, 10, 220, 70, rl.Fade(rl.SkyBlue, 0.5))
		rl.DrawRectangleLines(10, 10, 220, 70, rl.Blue)

		rl.DrawFPS(10, 10)

		rl.EndDrawing()

	}

}

var (
	maxColumns = 5
	heights    []float32
	positions  []rl.Vector3
	colors     []rl.Color
)

func (ts *RenderingSystem) Initilizer() {
	// Generates some random columns
	heights = make([]float32, maxColumns)
	positions = make([]rl.Vector3, maxColumns)
	colors = make([]rl.Color, maxColumns)

	for i := 0; i < maxColumns; i++ {
		heights[i] = float32(rl.GetRandomValue(1, 12))
		positions[i] = rl.NewVector3(float32(rl.GetRandomValue(-15, 15)), heights[i]/2, float32(rl.GetRandomValue(-15, 15)))
		colors[i] = rl.NewColor(uint8(rl.GetRandomValue(20, 255)), uint8(rl.GetRandomValue(10, 55)), 30, 255)
	}
}
