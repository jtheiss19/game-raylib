package componentsui

import (
	"github.com/jtheiss19/game-raylib/internal/ecs"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type UIComponent struct {
	*ecs.BaseComponent
	Position   rl.Vector2
	DrawFunc   func(position rl.Vector2)
	UpdateFunc func()
}

func NewUIComponent(position rl.Vector2) *UIComponent {
	return &UIComponent{
		BaseComponent: &ecs.BaseComponent{},
		Position:      position,
		DrawFunc: func(position rl.Vector2) {
		},
		UpdateFunc: func() {
		},
	}
}

func (bc *UIComponent) Draw() {
	bc.DrawFunc(bc.Position)
}

func (bc *UIComponent) Update() {
	bc.UpdateFunc()
}
