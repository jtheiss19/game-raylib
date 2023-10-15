package components3d

import (
	"rouge/internal/ecs"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type GridTransformation3DComponent struct {
	*ecs.BaseComponent
	Position rl.Vector3
	Scale    rl.Vector3
	Rotation rl.Vector3
	Forward  rl.Vector3
}

func NewGridTransformation3DComponent() *GridTransformation3DComponent {

	return &GridTransformation3DComponent{
		BaseComponent: &ecs.BaseComponent{},
		Position:      rl.Vector3{X: 0, Y: 0, Z: 0},
		Scale:         rl.Vector3{X: 1, Y: 1, Z: 1},
		Rotation:      rl.Vector3{X: 0, Y: 0, Z: 0},
		Forward:       rl.Vector3{X: 0, Y: 0, Z: 0},
	}
}
