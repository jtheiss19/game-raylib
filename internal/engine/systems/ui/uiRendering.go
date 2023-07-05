package systemsui

import (
	"fmt"
	"rouge/internal/ecs"
	componentsui "rouge/internal/engine/components/ui"

	"github.com/sirupsen/logrus"

	gui "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	demoTexture         = rl.Texture2D{}
	value       float32 = 0.4
)

func loadDemoTextures() {
	demoTexture = rl.LoadTexture("assets/box/crate_1.jpg")

}

type UIRenderingSystem struct {
	*ecs.BaseSystem
}

func NewUIRenderingSystem() *UIRenderingSystem {
	loadDemoTextures()
	return &UIRenderingSystem{
		BaseSystem: &ecs.BaseSystem{},
	}
}

// Comps
type RequiredUIRenderingSystemComps struct {
	UI []*RequireUI
}

type RequireUI struct {
	UIComponent *componentsui.UIBoxComponent
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

		for range entities.UI {
			value = gui.SliderBar(rl.NewRectangle(50, 150, 100, 40), "Click", "clicker", value, 0, 1)
			if value > 0.9 {
				fmt.Println(value)
			}
			rl.DrawTextureEx(
				demoTexture,
				rl.Vector2{
					X: 0,
					Y: 0,
				},
				0,
				0.5,
				rl.White,
			)
		}
	}
}

func (ts *UIRenderingSystem) Initilizer() {
}
