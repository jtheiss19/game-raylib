package objectsui

import (
	"rouge/internal/ecs"
	componentsui "rouge/internal/engine/components/ui"

	gui "github.com/gen2brain/raylib-go/raygui"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func NewWindowUI(positionX, positionY float32) []ecs.Component {
	windowComponent := componentsui.NewUIWindowComponent(rl.NewVector2(0, 0))

	checkBoxComponent := componentsui.NewUICheckBoxComponent(rl.NewVector2(32, 160))
	toggleGroupComponent := componentsui.NewUIToggleGroupComponent(rl.NewVector2(32, 208))
	sliderComponent := componentsui.NewUISliderComponent(rl.NewVector2(32, 248))
	buttonComponent := componentsui.NewUIButtonComponent(rl.NewVector2(32, 272))

	uiComponent := componentsui.NewUIComponent(rl.NewVector2(positionX, positionY))

	uiComponent.DrawFunc = func(position rl.Vector2) {
		if windowComponent.WindowBox000Active {
			windowComponent.WindowBox000Active = !gui.WindowBox(windowComponent.GetPosition(position), "SAMPLE TEXT")
			sliderComponent.SliderBar002Value = gui.SliderBar(sliderComponent.GetPosition(position), "", "", sliderComponent.SliderBar002Value, 0, 100)
			checkBoxComponent.CheckBoxEx005Checked = gui.CheckBox(checkBoxComponent.GetPosition(position), "SAMPLE TEXT", checkBoxComponent.CheckBoxEx005Checked)
			toggleGroupComponent.ToggleGroup006Active = gui.ToggleGroup(toggleGroupComponent.GetPosition(position), "ONE;TWO;THREE", toggleGroupComponent.ToggleGroup006Active)
			buttonComponent.Button001Pressed = gui.Button(buttonComponent.GetPosition(position), gui.IconText(5, "test"))
		}
	}

	return []ecs.Component{
		uiComponent,
	}
}
