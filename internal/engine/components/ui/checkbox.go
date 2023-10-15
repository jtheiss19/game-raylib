package componentsui

import (
	"github.com/jtheiss19/game-raylib/internal/ecs"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type UICheckBoxComponent struct {
	*ecs.BaseComponent
	CheckBoxEx005Checked bool
	Width                float32
	Height               float32
	Offset               rl.Vector2
}

func NewUICheckBoxComponent(offset rl.Vector2) *UICheckBoxComponent {
	return &UICheckBoxComponent{
		BaseComponent:        &ecs.BaseComponent{},
		CheckBoxEx005Checked: false,
		Width:                24,
		Height:               24,
		Offset:               offset,
	}
}

func (wc *UICheckBoxComponent) GetPosition(position rl.Vector2) rl.Rectangle {
	return rl.NewRectangle(
		position.X+wc.Offset.X,
		position.Y+wc.Offset.Y,
		wc.Width,
		wc.Height,
	)
}
