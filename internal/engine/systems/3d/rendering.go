package systems3d

import (
	"github.com/jtheiss19/game-raylib/internal/ecs"
	"github.com/jtheiss19/game-raylib/internal/engine/components"
	components3d "github.com/jtheiss19/game-raylib/internal/engine/components/3d"

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
	Model  []*Require3DModel
	Camera []*RequireCamera
}

type Require3DModel struct {
	Transformation *components3d.Transformation3DComponent
	ModelComp      *components3d.Model3DComponent
}

type RequireCamera struct {
	Player         *components.PlayerComponent
	Camera         *components3d.Camera3DComponent
	Transformation *components3d.Transformation3DComponent
}

func (ts *RenderingSystem) GetRequiredComponents() interface{} {
	return &RequiredRenderingSystemComps{
		Model: []*Require3DModel{{
			Transformation: &components3d.Transformation3DComponent{},
			ModelComp:      &components3d.Model3DComponent{},
		}},
		Camera: []*RequireCamera{{
			Camera:         &components3d.Camera3DComponent{},
			Transformation: &components3d.Transformation3DComponent{},
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
		rl.BeginMode3D(*camera)

		for _, entity := range entities.Model {
			rl.DrawModel(entity.ModelComp.Model, entity.Transformation.Position, entity.Transformation.Scale.X, rl.White)
		}

		rl.EndMode3D()
	}
}
