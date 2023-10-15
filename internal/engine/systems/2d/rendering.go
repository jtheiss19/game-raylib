package systems2d

import (
	"github.com/jtheiss19/game-raylib/internal/ecs"
	"github.com/jtheiss19/game-raylib/internal/engine/components"
	components2d "github.com/jtheiss19/game-raylib/internal/engine/components/2d"

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

		rl.DrawRectangle(10, 10, 80, 20, rl.Fade(rl.SkyBlue, 0.5))
		rl.DrawRectangleLines(10, 10, 80, 20, rl.Blue)
		rl.DrawFPS(10, 10)
	}
}

func (ts *RenderingSystem) Initilizer() {

}
