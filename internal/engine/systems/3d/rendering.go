package systems3d

import (
	"rouge/internal/ecs"
	"rouge/internal/engine/components"
	components3d "rouge/internal/engine/components/3d"

	"github.com/sirupsen/logrus"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type RenderingSystem struct {
	*ecs.BaseSystem
}

func NewRenderingSystem() *RenderingSystem {
	LoadDefaults()

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
	Transformation *components3d.Transformation3DComponent
}

type RequireCamera struct {
	Player         *components.PlayerComponent
	Camera         *components3d.Camera3DComponent
	Transformation *components3d.Transformation3DComponent
}

func (ts *RenderingSystem) GetRequiredComponents() interface{} {
	return &RequiredRenderingSystemComps{
		Model: []*RequireModel{{
			Transformation: &components3d.Transformation3DComponent{},
		}},
		Camera: []*RequireCamera{{
			Camera:         &components3d.Camera3DComponent{},
			Transformation: &components3d.Transformation3DComponent{},
		}},
	}
}

var (
	demoModel = rl.Model{}
	texture2D = rl.Texture2D{}
)

func LoadDefaults() {
	demoModel = rl.LoadModel(`assets\box\Crate1.obj`)
	texture2D = rl.LoadTexture(`assets\box\crate_1.jpg`)
	demoModel.Materials.Maps.Texture = texture2D
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
		rl.BeginMode3D(*camera)

		for _, entity := range entities.Model {
			rl.DrawModel(demoModel, entity.Transformation.Position, entity.Transformation.Scale.X, rl.White)
		}

		rl.EndMode3D()

		rl.DrawRectangle(10, 10, 220, 70, rl.Fade(rl.SkyBlue, 0.5))
		rl.DrawRectangleLines(10, 10, 220, 70, rl.Blue)

		rl.DrawFPS(10, 10)

	}
}
