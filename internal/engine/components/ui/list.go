package componentsui

import (
	"strings"

	"github.com/jtheiss19/game-raylib/internal/ecs"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type UIListComponent struct {
	*ecs.BaseComponent

	Values            string
	ScrollIndex       *int32
	ActiveOptionIndex int32
	Width             float32
	Height            float32
	Offset            rl.Vector2
}

func NewUIListComponent(offset rl.Vector2, values []string) *UIListComponent {
	return &UIListComponent{
		BaseComponent:     &ecs.BaseComponent{},
		Values:            strings.Join(values, ";"),
		ScrollIndex:       new(int32),
		ActiveOptionIndex: 0,
		Width:             152,
		Height:            250,
		Offset:            offset,
	}
}

func (wc *UIListComponent) GetPosition(position rl.Vector2) rl.Rectangle {
	return rl.NewRectangle(
		position.X+wc.Offset.X,
		position.Y+wc.Offset.Y,
		wc.Width,
		wc.Height,
	)
}
