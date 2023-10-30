package componentsui

import (
	"strings"

	"github.com/jtheiss19/game-raylib/pkg/ecs"

	gui "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type UIToggleGroupComponent struct {
	*ecs.BaseComponent

	ToggleGroup006Active int32
	Values               string
	Width                float32
	Height               float32
	Offset               rl.Vector2
}

func NewUIToggleGroupComponent(offset rl.Vector2, width, height float32, Values []string) *UIToggleGroupComponent {
	return &UIToggleGroupComponent{
		BaseComponent:        &ecs.BaseComponent{},
		ToggleGroup006Active: 0,
		Values:               strings.Join(Values, ";"),
		Width:                width,
		Height:               height,
		Offset:               offset,
	}
}

func (wc *UIToggleGroupComponent) getPosition(position rl.Vector2) rl.Rectangle {
	return rl.NewRectangle(
		position.X+wc.Offset.X,
		position.Y+wc.Offset.Y,
		wc.Width,
		wc.Height,
	)
}

func (wc *UIToggleGroupComponent) Render(position rl.Vector2) {
	wc.ToggleGroup006Active = gui.ToggleGroup(wc.getPosition(position), wc.Values, wc.ToggleGroup006Active)
}
