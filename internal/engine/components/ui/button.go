package componentsui

import (
	"rouge/internal/ecs"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type UIButtonComponent struct {
	*ecs.BaseComponent

	Button001Pressed bool
	Width            float32
	Height           float32
	Offset           rl.Vector2
}

func NewUIButtonComponent(offset rl.Vector2) *UIButtonComponent {
	return &UIButtonComponent{
		BaseComponent:    &ecs.BaseComponent{},
		Button001Pressed: false,
		Width:            120,
		Height:           24,
		Offset:           offset,
	}
}

func (wc *UIButtonComponent) GetPosition(position rl.Vector2) rl.Rectangle {
	return rl.NewRectangle(
		position.X+wc.Offset.X,
		position.Y+wc.Offset.Y,
		wc.Width,
		wc.Height,
	)
}
