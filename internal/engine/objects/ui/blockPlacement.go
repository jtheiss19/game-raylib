package objectsui

import (
	"encoding/json"
	"os"
	"reflect"

	"github.com/jtheiss19/game-raylib/internal/ecs"
	components3d "github.com/jtheiss19/game-raylib/internal/engine/components/3d"
	componentsui "github.com/jtheiss19/game-raylib/internal/engine/components/ui"
	systems3d "github.com/jtheiss19/game-raylib/internal/engine/systems/3d"
	"github.com/sirupsen/logrus"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func NewBlockPlacementUI(positionX, positionY float32, system *systems3d.PlayerControllerSystem) []ecs.Component {
	// Acutal Window
	windowComponent := componentsui.NewUIWindowComponent(192, 390)

	// Action Components
	toggleGroupComponent := componentsui.NewUIToggleGroupComponent(rl.NewVector2(20, 40), 74, 24, []string{
		"Individual",
		"Drag",
	})
	list := componentsui.NewUIListComponent(rl.NewVector2(20, 70), 152, 250, []string{
		"Craked Earth",
		"Ice",
		"Cobble",
		"Sand",
		"Ice",
		"Lava",
		"Snow",
		"Tile",
		"Carred Earth",
		"Long Brick",
		"Short Brick",
		"Thick Brick",
		"Miasma",
		"Dark Wood",
		"Fire Pit",
		"Dirt",
		"Mosaic Tile White",
		"Ice Sheet",
		"Dark Cobble",
		"Mosaic Tile Blue",
		"Grass",
		"Light Wood",
		"Cloth",
		"Fabric",
		"Clay Brick",
	})
	// Not rendering until we get confirmation that the input logic for the textbox doesn't bind to WASD movement
	// textBox := componentsui.NewUITextBoxComponent(rl.NewVector2(20, 330), 152, 20, "saveFile")
	saveButton := componentsui.NewUIButtonComponent(rl.NewVector2(20, 360), 74, 24, "Save")
	loadButton := componentsui.NewUIButtonComponent(rl.NewVector2(98, 360), 74, 24, "Load")

	// UI Rendering Abstraction
	uiComponent := componentsui.NewUIComponent(rl.NewVector2(positionX, positionY))

	// Drawing Func
	uiComponent.DrawFunc = func(position rl.Vector2) {
		if windowComponent.ShowWindow {
			// Window
			windowComponent.Render(position)

			// List
			toggleGroupComponent.Render(position)
			list.Render(position)
			//textBox.Render(position)
			saveButton.Render(position)
			loadButton.Render(position)
		}
	}

	uiComponent.UpdateFunc = func() {
		// Update Player Controller to use blockID
		system.BlockPlacementTextureIndex = int(list.ActiveOptionIndex) + 1

		// Save Button Function
		if saveButton.IsPressed {
			world := ecs.GetActiveWorld()
			comps := world.GetComponents(reflect.TypeOf(&components3d.Save3DComponent{}))
			newSaveFile := SaveFile{}
			for _, comp := range comps {
				id, _ := comp.GetComponentID()
				entity := world.GetEntity(id)

				ok := false
				var transform ecs.Component
				if transform, ok = entity[reflect.TypeOf(&components3d.Transformation3DComponent{})]; !ok {
					continue
				}
				var model ecs.Component
				if model, ok = entity[reflect.TypeOf(&components3d.Model3DComponent{})]; !ok {
					continue
				}

				transformCast := transform.(*components3d.Transformation3DComponent)
				newSaveFile.Transforms = append(newSaveFile.Transforms, transformCast)

				modelCast := model.(*components3d.Model3DComponent)
				prunedModel := &components3d.Model3DComponent{
					BaseComponent:       modelCast.BaseComponent,
					HashID:              modelCast.HashID,
					ModelDataLocation:   modelCast.ModelDataLocation,
					TextureDataLocation: modelCast.TextureDataLocation,
					TextureFrame:        modelCast.TextureFrame,
					FragmentShader:      modelCast.FragmentShader,
					VertexShader:        modelCast.VertexShader,
				}
				newSaveFile.Models = append(newSaveFile.Models, prunedModel)
			}

			bytes, _ := json.MarshalIndent(newSaveFile, "", "	")
			err := os.WriteFile("saveFile", bytes, 0644)
			if err != nil {
				logrus.Error("error saving data: ", err.Error())
			} else {
				logrus.Info("Saved Game")
			}
		}
		// Load Button Function
		if loadButton.IsPressed {
			bytes, err := os.ReadFile("saveFile")
			if err != nil {
				logrus.Error("error reading save data: ", err.Error())
				return
			}

			save := SaveFile{}
			err = json.Unmarshal(bytes, &save)
			if err != nil {
				logrus.Error("error loading save data: ", err.Error())
				return
			}

			world := ecs.GetActiveWorld()
			for _, comp := range save.Models {
				world.AddComponent(comp)
			}
			for _, comp := range save.Transforms {
				world.AddComponent(comp)
			}

		}
	}

	return []ecs.Component{
		uiComponent,
	}
}

type SaveFile struct {
	Models     []*components3d.Model3DComponent
	Transforms []*components3d.Transformation3DComponent
}
