package componentsui

import (
	"github.com/jtheiss19/game-raylib/pkg/ecs"

	gui "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type UIWindowComponent struct {
	*ecs.BaseComponent

	ShowWindow bool
	Width      float32
	Height     float32
}

func NewUIWindowComponent(width, height float32) *UIWindowComponent {
	return &UIWindowComponent{
		BaseComponent: &ecs.BaseComponent{},
		ShowWindow:    true,
		Width:         width,
		Height:        height,
	}
}

func (wc *UIWindowComponent) getPosition(position rl.Vector2) rl.Rectangle {
	return rl.NewRectangle(
		position.X,
		position.Y,
		wc.Width,
		wc.Height,
	)
}

func (wc *UIWindowComponent) Render(position rl.Vector2) {
	wc.ShowWindow = !gui.WindowBox(wc.getPosition(position), "Block Picker")
}
