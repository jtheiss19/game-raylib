package componentsui

import (
	"github.com/jtheiss19/game-raylib/internal/ecs"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type UISliderComponent struct {
	*ecs.BaseComponent

	SliderBar002Value float32
	Width             float32
	Height            float32
	Offset            rl.Vector2
}

func NewUISliderComponent(offset rl.Vector2) *UISliderComponent {
	return &UISliderComponent{
		BaseComponent:     &ecs.BaseComponent{},
		SliderBar002Value: 0,
		Width:             120,
		Height:            16,
		Offset:            offset,
	}
}

func (wc *UISliderComponent) GetPosition(position rl.Vector2) rl.Rectangle {
	return rl.NewRectangle(
		position.X+wc.Offset.X,
		position.Y+wc.Offset.Y,
		wc.Width,
		wc.Height,
	)
}
