package componentsui

import (
	"strings"

	"github.com/jtheiss19/game-raylib/pkg/ecs"

	gui "github.com/gen2brain/raylib-go/raygui"
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

func NewUIListComponent(offset rl.Vector2, width, height float32, values []string) *UIListComponent {
	return &UIListComponent{
		BaseComponent:     &ecs.BaseComponent{},
		Values:            strings.Join(values, ";"),
		ScrollIndex:       new(int32),
		ActiveOptionIndex: 0,
		Width:             width,
		Height:            height,
		Offset:            offset,
	}
}

func (wc *UIListComponent) getPosition(position rl.Vector2) rl.Rectangle {
	return rl.NewRectangle(
		position.X+wc.Offset.X,
		position.Y+wc.Offset.Y,
		wc.Width,
		wc.Height,
	)
}

func (wc *UIListComponent) Render(position rl.Vector2) {
	wc.ActiveOptionIndex = gui.ListView(wc.getPosition(position), wc.Values, wc.ScrollIndex, wc.ActiveOptionIndex)
}
