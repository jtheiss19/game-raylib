package components3d

import (
	"rouge/internal/ecs"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Camera3DComponent struct {
	*ecs.BaseComponent
	Camera *rl.Camera3D
}

func NewCamera3DComponent() *Camera3DComponent {
	camera := rl.Camera3D{}
	camera.Position = rl.NewVector3(0, 0, 0)
	camera.Target = rl.NewVector3(1, 0, 0.0) //Left, Forwards,
	camera.Up = rl.NewVector3(0, 1, 0)
	camera.Fovy = 90
	camera.Projection = rl.CameraPerspective

	// rl.SetCameraMode(camera, rl.CameraFirstPerson)

	return &Camera3DComponent{
		BaseComponent: &ecs.BaseComponent{},
		Camera:        &camera,
	}
}
