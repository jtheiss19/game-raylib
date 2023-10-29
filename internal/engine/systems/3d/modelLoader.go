package systems3d

import (
	"github.com/jtheiss19/game-raylib/internal/ecs"
	components3d "github.com/jtheiss19/game-raylib/internal/engine/components/3d"

	"github.com/sirupsen/logrus"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type ModelLoadingSystem struct {
	*ecs.BaseSystem

	loadedModels map[string]*rl.Model
}

func NewModelLoadingSystem() *ModelLoadingSystem {
	return &ModelLoadingSystem{
		BaseSystem:   &ecs.BaseSystem{},
		loadedModels: map[string]*rl.Model{},
	}
}

// Comps
type RequiredModelLoadingSystemComps struct {
	Models []*RequiredModelData
}

type RequiredModelData struct {
	ModelComp *components3d.Model3DComponent
}

func (ts *ModelLoadingSystem) GetRequiredComponents() interface{} {
	return &RequiredModelLoadingSystemComps{
		Models: []*RequiredModelData{{
			ModelComp: &components3d.Model3DComponent{},
		}},
	}
}

// Functionality
func (ts *ModelLoadingSystem) Update(dt float32) {
	entities, ok := ts.TrackedEntities.(*RequiredModelLoadingSystemComps)
	if !ok {
		logrus.Error("could not update system, bad tracked entities")
		return
	}

	for _, modelData := range entities.Models {
		if modelData.ModelComp.LoadedModel { // Model already loaded in map and set on component, no work to do
			continue
		}
		loadedMapModel, ok := ts.loadedModels[modelData.ModelComp.HashID]
		if !ok { // Model Not loaded yet, load from disk
			// Load Model
			loadedModel := rl.LoadModel(string(modelData.ModelComp.ModelDataLocation))

			// Load Texture
			loadedImage := rl.LoadImage(string(modelData.ModelComp.TextureDataLocation))
			// Handle Sprite Sheets
			if modelData.ModelComp.TextureFrame != 0 {
				rl.ImageCrop(loadedImage, rl.Rectangle{
					X:      0,
					Y:      0,
					Width:  float32(loadedImage.Width),
					Height: float32(loadedImage.Height),
				})
			}
			loadedModel.Materials.Maps.Texture = rl.LoadTextureFromImage(loadedImage)

			// Load Shader
			if modelData.ModelComp.VertexShader != "" && modelData.ModelComp.FragmentShader != "" {
				shader := rl.LoadShader(string(modelData.ModelComp.VertexShader), string(modelData.ModelComp.FragmentShader))
				shader.UpdateLocation(rl.LocMatrixMvp, rl.GetShaderLocation(shader, "mvp"))
				shader.UpdateLocation(rl.LocVectorView, rl.GetShaderLocation(shader, "viewPos"))
				shader.UpdateLocation(rl.LocMatrixModel, rl.GetShaderLocationAttrib(shader, "instanceTransform"))
				loadedModel.Materials.Shader = shader
			}

			// Create Records
			modelData.ModelComp.Model = loadedModel
			modelData.ModelComp.LoadedModel = true
			ts.loadedModels[modelData.ModelComp.HashID] = &loadedModel
		} else if !modelData.ModelComp.LoadedModel { // Model already loaded in map but not set on component
			modelData.ModelComp.Model = *loadedMapModel
		}
	}

}
