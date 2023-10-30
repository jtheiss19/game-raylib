package componentsui

import (
	"github.com/jtheiss19/game-raylib/internal/ecs"

	gui "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type UITextBoxComponent struct {
	*ecs.BaseComponent

	Value             string
	ScrollIndex       *int32
	ActiveOptionIndex int32
	Width             float32
	Height            float32
	Offset            rl.Vector2
}

func NewUITextBoxComponent(offset rl.Vector2, width, height float32, defaultValue string) *UITextBoxComponent {
	return &UITextBoxComponent{
		BaseComponent:     &ecs.BaseComponent{},
		Value:             defaultValue,
		ActiveOptionIndex: 0,
		Width:             width,
		Height:            height,
		Offset:            offset,
	}
}

func (wc *UITextBoxComponent) getPosition(position rl.Vector2) rl.Rectangle {
	return rl.NewRectangle(
		position.X+wc.Offset.X,
		position.Y+wc.Offset.Y,
		wc.Width,
		wc.Height,
	)
}

func (wc *UITextBoxComponent) Render(position rl.Vector2) {
	gui.TextBox(wc.getPosition(position), &wc.Value, 128, true)
}
