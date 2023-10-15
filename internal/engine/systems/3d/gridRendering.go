package systems3d

import (
	"rouge/internal/ecs"
	"rouge/internal/engine/components"
	components3d "rouge/internal/engine/components/3d"

	"github.com/sirupsen/logrus"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type GridRenderingSystem struct {
	*ecs.BaseSystem
}

func NewGridRenderingSystem() *GridRenderingSystem {

	return &GridRenderingSystem{
		BaseSystem: &ecs.BaseSystem{},
	}
}

// Comps
type RequiredGridRenderingSystemComps struct {
	Model  []*RequireGridModel
	Camera []*RequireGridCamera
}

type RequireGridModel struct {
	Transformation *components3d.GridTransformation3DComponent
	ModelComp      *components3d.Model3DComponent
}

type RequireGridCamera struct {
	Player         *components.PlayerComponent
	Camera         *components3d.Camera3DComponent
	Transformation *components3d.Transformation3DComponent
}

func (ts *GridRenderingSystem) GetRequiredComponents() interface{} {
	return &RequiredGridRenderingSystemComps{
		Model: []*RequireGridModel{{
			Transformation: &components3d.GridTransformation3DComponent{},
			ModelComp:      &components3d.Model3DComponent{},
		}},
		Camera: []*RequireGridCamera{{
			Camera:         &components3d.Camera3DComponent{},
			Transformation: &components3d.Transformation3DComponent{},
		}},
	}
}

// Functionality
func (ts *GridRenderingSystem) Update(dt float32) {
	entities, ok := ts.TrackedEntities.(*RequiredGridRenderingSystemComps)
	if !ok {
		logrus.Error("could not update system, bad tracked entities")
		return
	}

	if len(entities.Camera) > 0 {

		camera := entities.Camera[0].Camera.Camera

		// Draw
		rl.BeginMode3D(*camera)

		for _, entity := range entities.Model {
			rl.DrawModel(entity.ModelComp.Model, entity.Transformation.Position, entity.Transformation.Scale.X, rl.White)
		}

		rl.DrawGrid(100, 2)

		rl.EndMode3D()

		rl.DrawRectangle(10, 10, 220, 70, rl.Fade(rl.SkyBlue, 0.5))
		rl.DrawRectangleLines(10, 10, 220, 70, rl.Blue)

		rl.DrawFPS(10, 10)

	}
}
