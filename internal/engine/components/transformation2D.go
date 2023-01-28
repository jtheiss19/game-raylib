package components

import (
	"rouge/internal/ecs"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Transformation2DComponent struct {
	*ecs.BaseComponent
	Position rl.Vector3
	Scale    rl.Vector3
	Rotation rl.Vector3
	Forward  rl.Vector3
}

func NewTransformationComponent() *Transformation2DComponent {

	return &Transformation2DComponent{
		BaseComponent: &ecs.BaseComponent{},
		Position:      rl.Vector3{X: 0, Y: 10.0, Z: 10.0},
		Scale:         rl.Vector3{X: 1, Y: 1, Z: 1},
		Rotation:      rl.Vector3{X: 0, Y: 0, Z: 0},
		Forward:       rl.Vector3{X: 0, Y: 0, Z: 0},
	}
}
