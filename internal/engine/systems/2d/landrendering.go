package systems2d

import (
	"rouge/internal/ecs"
	"rouge/internal/engine/components"
	components2d "rouge/internal/engine/components/2d"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/sirupsen/logrus"
)

type LandRenderingSystem struct {
	*ecs.BaseSystem
}

func NewLandRenderingSystem() *LandRenderingSystem {
	return &LandRenderingSystem{
		BaseSystem: &ecs.BaseSystem{},
	}
}

// Comps
type RequiredLandRenderingSystemComps struct {
	Land    []*RequireLand
	Cameras []*RequireLandPlayer
}

type RequireLand struct {
	LandComponent      *components2d.LandComponent
	TransformComponent *components2d.Transformation2DComponent
}

type RequireLandPlayer struct {
	Input          *components.InputComponent
	Camera         *components2d.Camera2DComponent
	Transformation *components2d.Transformation2DComponent
	Player         *components.PlayerComponent
}

func (ts *LandRenderingSystem) GetRequiredComponents() interface{} {
	return &RequiredLandRenderingSystemComps{
		Land:    []*RequireLand{{}},
		Cameras: []*RequireLandPlayer{{}},
	}
}

// Functionality
func (ts *LandRenderingSystem) Update(dt float32) {
	entities, ok := ts.TrackedEntities.(*RequiredLandRenderingSystemComps)
	if !ok {
		logrus.Error("could not update system, bad tracked entities")
		return
	}

	if len(entities.Cameras) > 0 {

		camera := entities.Cameras[0].Camera.Camera

		// Draw
		rl.BeginMode2D(*camera)
		for _, entity := range entities.Land {
			for i, value := range entity.LandComponent.Data {
				if value != 0 {
					xOffset := float32(rl.GetScreenWidth())/2 - 5 + entity.TransformComponent.Position.X
					yOffset := float32(rl.GetScreenHeight())/2 - 5 + entity.TransformComponent.Position.Y
					rl.DrawTextureRec(
						demoTexture,
						rl.Rectangle{
							X:      0,
							Y:      0,
							Width:  10,
							Height: 10,
						},
						rl.Vector2{
							X: xOffset + float32(i%entity.LandComponent.Width*10),
							Y: yOffset + float32(i/entity.LandComponent.Width*10),
						},
						rl.White,
					)
				}
			}
		}
		rl.EndMode2D()
	}
}

func (ts *LandRenderingSystem) Initilizer() {
}
