package components2d

import (
	"rouge/internal/ecs"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Camera2DComponent struct {
	*ecs.BaseComponent
	Camera *rl.Camera2D
}

func NewCamera2DComponent() *Camera2DComponent {
	camera := rl.Camera2D{}
	camera.Offset = rl.NewVector2(0, 0)
	camera.Target = rl.NewVector2(0, 0)
	camera.Rotation = 0
	camera.Zoom = 1

	// rl.SetCameraMode(camera, rl.CameraFirstPerson)

	return &Camera2DComponent{
		BaseComponent: &ecs.BaseComponent{},
		Camera:        &camera,
	}
}
