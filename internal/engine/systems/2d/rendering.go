package systems3d

import (
	"rouge/internal/ecs"
	"rouge/internal/engine/components"
	components2d "rouge/internal/engine/components/2d"

	"github.com/sirupsen/logrus"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type RenderingSystem struct {
	*ecs.BaseSystem
}

func NewRenderingSystem() *RenderingSystem {
	loadDemoTextures()
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
	Transformation *components2d.Transformation2DComponent
	Model          *components.ModelComponent
}

type RequireCamera struct {
	Player         *components.PlayerComponent
	Camera         *components2d.Camera2DComponent
	Transformation *components2d.Transformation2DComponent
}

func (ts *RenderingSystem) GetRequiredComponents() interface{} {
	return &RequiredRenderingSystemComps{
		Model: []*RequireModel{{
			Transformation: &components2d.Transformation2DComponent{},
			Model:          &components.ModelComponent{},
		}},
		Camera: []*RequireCamera{{
			Camera:         &components2d.Camera2DComponent{},
			Transformation: &components2d.Transformation2DComponent{},
		}},
	}
}

var (
	demoTexture = rl.Texture2D{}
	rec         = rl.Rectangle{}
)

func loadDemoTextures() {
	demoTexture = rl.LoadTexture("assets/box/crate_1.jpg")

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
		rl.BeginMode2D(*camera)

		for _, entity := range entities.Model {
			rl.DrawTextureRec(
				demoTexture,
				rl.Rectangle{
					X:      0,
					Y:      0,
					Width:  10,
					Height: 10,
				},
				rl.Vector2{
					X: float32(rl.GetScreenWidth())/2 - 5 + entity.Transformation.Position.X,
					Y: float32(rl.GetScreenHeight())/2 - 5 + entity.Transformation.Position.Y,
				},
				rl.White,
			)
		}

		rl.EndMode2D()

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
