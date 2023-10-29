package objectsui

import (
	"github.com/jtheiss19/game-raylib/internal/ecs"
	componentsui "github.com/jtheiss19/game-raylib/internal/engine/components/ui"

	gui "github.com/gen2brain/raylib-go/raygui"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func NewBlockPlacementUI(positionX, positionY float32) []ecs.Component {
	// Acutal Window
	windowComponent := componentsui.NewUIWindowComponent(rl.NewVector2(0, 0))

	// Action Components
	List := componentsui.NewUIListComponent(rl.NewVector2(20, 40), []string{
		"Grass",
		"Sand",
		"Brick",
		"Floor",
		"Water",
		"Tile",
		"Stone",
	})

	// UI Rendering Abstraction
	uiComponent := componentsui.NewUIComponent(rl.NewVector2(positionX, positionY))

	// Drawing Func
	uiComponent.DrawFunc = func(position rl.Vector2) {
		if windowComponent.WindowBox000Active {
			// Window
			windowComponent.WindowBox000Active = !gui.WindowBox(windowComponent.GetPosition(position), "Block Picker")

			// List
			List.ActiveOptionIndex = gui.ListView(List.GetPosition(position), List.Values, List.ScrollIndex, List.ActiveOptionIndex)
		}
	}

	return []ecs.Component{
		uiComponent,
	}
}
