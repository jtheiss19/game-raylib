package objectsui

import (
	"rouge/internal/ecs"
	componentsui "rouge/internal/engine/components/ui"
)

func NewBoxUI() []ecs.Component {
	uiComponent := componentsui.NewUIBoxComponent()

	return []ecs.Component{
		uiComponent,
	}
}
