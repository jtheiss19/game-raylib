package systems3d

import (
	"rouge/internal/ecs"
	"rouge/internal/engine/components"
	components3d "rouge/internal/engine/components/3d"

	"github.com/sirupsen/logrus"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type GridPlayerControllerSystem struct {
	*ecs.BaseSystem
}

func NewGridPlayerControllerSystem() *GridPlayerControllerSystem {
	return &GridPlayerControllerSystem{
		BaseSystem: &ecs.BaseSystem{},
	}
}

// Comps
type RequiredGridPlayerControllerComps struct {
	Player []*RequireGridPlayer
}

type RequireGridPlayer struct {
	Input          *components.InputComponent
	Camera         *components3d.Camera3DComponent
	Transformation *components3d.Transformation3DComponent
	Player         *components.PlayerComponent
}

func (ts *GridPlayerControllerSystem) GetRequiredComponents() interface{} {
	return &RequiredGridPlayerControllerComps{
		Player: []*RequireGridPlayer{{
			Input:          &components.InputComponent{},
			Transformation: &components3d.Transformation3DComponent{},
		}},
	}
}

// Functionality
func (ts *GridPlayerControllerSystem) Update(dt float32) {
	entities, ok := ts.TrackedEntities.(*RequiredGridPlayerControllerComps)
	if !ok {
		logrus.Error("could not update system, bad tracked entities")
		return
	}

	for _, player := range entities.Player {
		// Mouse movement
		camera := player.Camera.Camera

		screenSize := rl.Vector2{float32(rl.GetScreenWidth()) / 2, float32(rl.GetScreenHeight()) / 2}
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

		// normalize 2D movement
		movementVector = rl.Vector3Normalize(movementVector)

		forwardVec := mouseRay.Direction
		rightVec := rl.Vector3CrossProduct(forwardVec, rl.Vector3{0, 1, 0})

		forwardMovement := rl.Vector3Multiply(forwardVec, movementVector.X)
		rightMovement := rl.Vector3Multiply(rightVec, movementVector.Z)

		totalMovement := rl.Vector3Add(forwardMovement, rightMovement)

		player.Transformation.Position.X += rl.Vector3DotProduct(totalMovement, rl.Vector3{1, 0, 0}) * 0.25
		player.Transformation.Position.Z += rl.Vector3DotProduct(totalMovement, rl.Vector3{0, 0, 1}) * 0.25

	}
}

func (ts *GridPlayerControllerSystem) Initilizer() {
}
