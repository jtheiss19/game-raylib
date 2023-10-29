package componentsui

import (
	"github.com/jtheiss19/game-raylib/internal/ecs"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type UIDropBoxComponent struct {
	*ecs.BaseComponent

	DropdownBox000EditMode bool
	DropdownBox000Active   *int32
	Width                  float32
	Height                 float32
	Offset                 rl.Vector2
}

func NewUIDropBoxComponent(offset rl.Vector2) *UIDropBoxComponent {
	return &UIDropBoxComponent{
		BaseComponent:          &ecs.BaseComponent{},
		DropdownBox000EditMode: false,
		DropdownBox000Active:   new(int32),
		Width:                  120,
		Height:                 24,
		Offset:                 offset,
	}
}

func (wc *UIDropBoxComponent) GetPosition(position rl.Vector2) rl.Rectangle {
	return rl.NewRectangle(
		position.X+wc.Offset.X,
		position.Y+wc.Offset.Y,
		wc.Width,
		wc.Height,
	)
}
