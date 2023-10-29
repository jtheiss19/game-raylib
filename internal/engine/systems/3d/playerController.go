package systems3d

import (
	"github.com/jtheiss19/game-raylib/internal/ecs"
	"github.com/jtheiss19/game-raylib/internal/engine/components"
	components3d "github.com/jtheiss19/game-raylib/internal/engine/components/3d"

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
	Camera         *components3d.Camera3DComponent
	Transformation *components3d.Transformation3DComponent
	Player         *components.PlayerComponent
}

func (ts *PlayerControllerSystem) GetRequiredComponents() interface{} {
	return &RequiredPlayerControllerComps{
		Player: []*RequirePlayer{{
			Input:          &components.InputComponent{},
			Transformation: &components3d.Transformation3DComponent{},
		}},
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
		// Mouse movement
		camera := player.Camera.Camera

		screenSize := rl.Vector2{X: float32(rl.GetScreenWidth()) / 2, Y: float32(rl.GetScreenHeight()) / 2}
		mousePosVec := rl.Vector2Add(rl.GetMouseDelta(), screenSize)
		mouseRay := rl.GetMouseRay(mousePosVec, *camera)

		camera.Target.X = mouseRay.Direction.X + player.Transformation.Position.X
		camera.Target.Y = mouseRay.Direction.Y + player.Transformation.Position.Y
		camera.Target.Z = mouseRay.Direction.Z + player.Transformation.Position.Z

		camera.Position = player.Transformation.Position

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
			movementVector.X += 1
		}
		if keys[components.MOVE_BACKWARDS] {
			movementVector.X -= 1
		}
		if keys[components.MOVE_RIGHT] {
			movementVector.Z += 1
		}
		if keys[components.MOVE_LEFT] {
			movementVector.Z -= 1
		}
		if keys[components.MOVE_UP] {
			movementVector.Y += 1
		}
		if keys[components.MOVE_DOWN] {
			movementVector.Y -= 1
		}

		// normalize 2D movement
		movementVector = rl.Vector3Normalize(movementVector)

		forwardVec := mouseRay.Direction
		forwardVec.Y = 0
		rightVec := rl.Vector3CrossProduct(forwardVec, rl.Vector3{X: 0, Y: 1, Z: 0})
		upVec := rl.Vector3CrossProduct(rightVec, forwardVec)

		forwardMovement := rl.Vector3Multiply(forwardVec, movementVector.X)
		rightMovement := rl.Vector3Multiply(rightVec, movementVector.Z)
		upMovement := rl.Vector3Multiply(upVec, movementVector.Y)

		totalMovement := rl.Vector3Add(rl.Vector3Add(forwardMovement, rightMovement), upMovement)

		player.Transformation.Position.X += rl.Vector3DotProduct(totalMovement, rl.Vector3{X: 1, Y: 0, Z: 0}) * 0.25
		player.Transformation.Position.Z += rl.Vector3DotProduct(totalMovement, rl.Vector3{X: 0, Y: 0, Z: 1}) * 0.25
		player.Transformation.Position.Y += rl.Vector3DotProduct(totalMovement, rl.Vector3{X: 0, Y: 1, Z: 0}) * 0.25

	}
}

func (ts *PlayerControllerSystem) Initilizer() {
}
