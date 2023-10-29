package systemsui

import (
	"github.com/jtheiss19/game-raylib/internal/ecs"
	componentsui "github.com/jtheiss19/game-raylib/internal/engine/components/ui"

	"github.com/sirupsen/logrus"
)

type UIRenderingSystem struct {
	*ecs.BaseSystem
}

func NewUIRenderingSystem() *UIRenderingSystem {
	return &UIRenderingSystem{
		BaseSystem: &ecs.BaseSystem{},
	}
}

// Comps
type RequiredUIRenderingSystemComps struct {
	UI []*RequireUI
}

type RequireUI struct {
	UIComponent *componentsui.UIComponent
}

func (ts *UIRenderingSystem) GetRequiredComponents() interface{} {
	return &RequiredUIRenderingSystemComps{
		UI: []*RequireUI{{}},
	}
}

// Functionality
func (ts *UIRenderingSystem) Update(dt float32) {
	entities, ok := ts.TrackedEntities.(*RequiredUIRenderingSystemComps)
	if !ok {
		logrus.Error("could not update system, bad tracked entities")
		return
	}

	if len(entities.UI) > 0 {
		for _, uiElement := range entities.UI {
			// Actual rendering happens in the DRAWFUNC of the UI comoponent which is set in the object creation func
			uiElement.UIComponent.Draw()
		}
	}
}

func (ts *UIRenderingSystem) Initilizer() {
}
