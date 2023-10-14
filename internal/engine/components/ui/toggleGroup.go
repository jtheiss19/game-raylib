package componentsui

import (
	"rouge/internal/ecs"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type UIToggleGroupComponent struct {
	*ecs.BaseComponent

	ToggleGroup006Active int32
	Width                float32
	Height               float32
	Offset               rl.Vector2
}

func NewUIToggleGroupComponent(offset rl.Vector2) *UIToggleGroupComponent {
	return &UIToggleGroupComponent{
		BaseComponent:        &ecs.BaseComponent{},
		ToggleGroup006Active: 0,
		Width:                40,
		Height:               24,
		Offset:               offset,
	}
}

func (wc *UIToggleGroupComponent) GetPosition(position rl.Vector2) rl.Rectangle {
	return rl.NewRectangle(
		position.X+wc.Offset.X,
		position.Y+wc.Offset.Y,
		wc.Width,
		wc.Height,
	)
}
