package systems3d

import (
	"math"

	"github.com/jtheiss19/game-raylib/internal/ecs"
	"github.com/jtheiss19/game-raylib/internal/engine/components"
	components3d "github.com/jtheiss19/game-raylib/internal/engine/components/3d"

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

var (
	gridControllerFrameTime float32 = 0
)

// Functionality
func (ts *GridPlayerControllerSystem) Update(dt float32) {
	gridControllerFrameTime += dt

	entities, ok := ts.TrackedEntities.(*RequiredGridPlayerControllerComps)
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
		} else if keys[components.MOVE_BACKWARDS] {
			movementVector.X -= 1
		} else if keys[components.MOVE_RIGHT] {
			movementVector.Z += 1
		} else if keys[components.MOVE_LEFT] {
			movementVector.Z -= 1
		}

		// normalize 2D movement
		movementVector = rl.Vector3Normalize(movementVector)
		lookVector := rl.Vector3{}

		forwardVec := mouseRay.Direction

		if math.Abs(float64(forwardVec.X)) < math.Abs(float64(forwardVec.Z)) {
			lookVector.Z = float32(math.Abs(float64(forwardVec.Z)) / float64(forwardVec.Z))
		} else {
			lookVector.X = float32(math.Abs(float64(forwardVec.X)) / float64(forwardVec.X))
		}
		rightVec := rl.Vector3CrossProduct(lookVector, rl.Vector3{X: 0, Y: 1, Z: 0})

		forwardMovement := rl.Vector3Multiply(lookVector, movementVector.X)
		rightMovement := rl.Vector3Multiply(rightVec, movementVector.Z)

		totalMovement := rl.Vector3Add(forwardMovement, rightMovement)

		if gridControllerFrameTime > 300 && rl.Vector3Length(totalMovement) > 0 {
			player.Transformation.Position.X += totalMovement.X
			player.Transformation.Position.Z += totalMovement.Z
			gridControllerFrameTime = 0
		}
	}
}

func (ts *GridPlayerControllerSystem) Initilizer() {
}
