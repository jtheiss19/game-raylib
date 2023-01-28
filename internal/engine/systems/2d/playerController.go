package systems3d

import (
	"rouge/internal/ecs"
	"rouge/internal/engine/components"
	components2d "rouge/internal/engine/components/2d"

	"github.com/sirupsen/logrus"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type PlayerControllerSystem struct {
	*ecs.BaseSystem
}

func NewPlayerControllerSystem() *PlayerControllerSystem {
	return &PlayerControllerSystem{
		BaseSystem: &ecs.BaseSystem{},
	}
}

// Comps
type RequiredPlayerControllerComps struct {
	Player []*RequirePlayer
}

type RequirePlayer struct {
	Input          *components.InputComponent
	Camera         *components2d.Camera2DComponent
	Transformation *components2d.Transformation2DComponent
	Player         *components.PlayerComponent
}

func (ts *PlayerControllerSystem) GetRequiredComponents() interface{} {
	return &RequiredPlayerControllerComps{
		Player: []*RequirePlayer{{}},
	}
}

// Functionality
func (ts *PlayerControllerSystem) Update(dt float32) {
	entities, ok := ts.TrackedEntities.(*RequiredPlayerControllerComps)
	if !ok {
		logrus.Error("could not update system, bad tracked entities")
		return
	}

	for _, player := range entities.Player {
		// calculate input
		player.Input.CalculateInput()
		keys := player.Input.Keys

		movementVector := rl.Vector3{
			X: 0,
			Y: 0,
			Z: 0,
		}

		// resolve movement
		if keys[components.MOVE_FORWARD] {
			movementVector.Y -= 1
		}
		if keys[components.MOVE_BACKWARDS] {
			movementVector.Y += 1
		}
		if keys[components.MOVE_RIGHT] {
			movementVector.X += 1
		}
		if keys[components.MOVE_LEFT] {
			movementVector.X -= 1
		}

		// normalize 2D movement
		movementVector = rl.Vector3Normalize(movementVector)

		player.Transformation.Position.X += movementVector.X
		player.Transformation.Position.Y += movementVector.Y

		camera := player.Camera.Camera
		camera.Target = player.Transformation.Position

	}
}

func (ts *PlayerControllerSystem) Initilizer() {
}
