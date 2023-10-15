package componentsui

import (
	"github.com/jtheiss19/game-raylib/internal/ecs"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type UIWindowComponent struct {
	*ecs.BaseComponent

	WindowBox000Active bool
	Width              float32
	Height             float32
	Offset             rl.Vector2
}

func NewUIWindowComponent(offset rl.Vector2) *UIWindowComponent {
	return &UIWindowComponent{
		BaseComponent:      &ecs.BaseComponent{},
		WindowBox000Active: true,
		Width:              192,
		Height:             304,
		Offset:             offset,
	}
}

func (wc *UIWindowComponent) GetPosition(position rl.Vector2) rl.Rectangle {
	return rl.NewRectangle(
		position.X+wc.Offset.X,
		position.Y+wc.Offset.Y,
		wc.Width,
		wc.Height,
	)
}
