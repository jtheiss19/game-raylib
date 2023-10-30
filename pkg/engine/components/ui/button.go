package componentsui

import (
	"github.com/jtheiss19/game-raylib/pkg/ecs"

	gui "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type UIButtonComponent struct {
	*ecs.BaseComponent

	IsPressed bool
	Width     float32
	Height    float32
	Offset    rl.Vector2
	Text      string
}

func NewUIButtonComponent(offset rl.Vector2, width, height float32, defaultText string) *UIButtonComponent {
	return &UIButtonComponent{
		BaseComponent: &ecs.BaseComponent{},
		IsPressed:     false,
		Width:         width,
		Height:        height,
		Offset:        offset,
		Text:          defaultText,
	}
}

func (wc *UIButtonComponent) getPosition(position rl.Vector2) rl.Rectangle {
	return rl.NewRectangle(
		position.X+wc.Offset.X,
		position.Y+wc.Offset.Y,
		wc.Width,
		wc.Height,
	)
}

func (wc *UIButtonComponent) Render(position rl.Vector2) {
	wc.IsPressed = gui.Button(wc.getPosition(position), gui.IconText(5, wc.Text))
}
