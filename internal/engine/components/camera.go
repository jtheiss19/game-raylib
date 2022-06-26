package components

import (
	"rouge/internal/ecs"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type CameraComponent struct {
	*ecs.BaseComponent
	Camera *rl.Camera3D
}

func NewCameraComponent() *CameraComponent {
	camera := rl.Camera3D{}
	camera.Position = rl.NewVector3(4.0, 2.0, 4.0)
	camera.Target = rl.NewVector3(0.0, 1.8, 0.0)
	camera.Up = rl.NewVector3(0.0, 1.0, 0.0)
	camera.Fovy = 60.0
	camera.Projection = rl.CameraPerspective

	// rl.SetCameraMode(camera, rl.CameraFirstPerson)

	return &CameraComponent{
		BaseComponent: &ecs.BaseComponent{},
		Camera:        &camera,
	}
}
