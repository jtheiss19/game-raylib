package components2d

import (
	"rouge/internal/ecs"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Transformation2DComponent struct {
	*ecs.BaseComponent
	Position rl.Vector2
	Scale    rl.Vector2
	Rotation float32
}

func NewTransformation2DComponent() *Transformation2DComponent {

	return &Transformation2DComponent{
		BaseComponent: &ecs.BaseComponent{},
		Position:      rl.Vector2{X: 0, Y: 0},
		Scale:         rl.Vector2{X: 1, Y: 1},
		Rotation:      0,
	}
}
