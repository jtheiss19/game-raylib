package systems3d

import (
	"github.com/jtheiss19/game-raylib/internal/ecs"
	components3d "github.com/jtheiss19/game-raylib/internal/engine/components/3d"

	"github.com/sirupsen/logrus"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type TextureType string

const (
	CRATE_TEX TextureType = `assets\box\crate.jpg`
	GRASS_TEX TextureType = `assets\box\grass.jpg`
)

type ObjectType string

const (
	CRATE_OBJ TextureType = `assets\box\Crate.obj`
)

type FragmentShader string

const (
	INSTANCED_FRAG FragmentShader = `assets\box\lighting.fs`
)

type VertexShader string

const (
	INSTANCED_VERT VertexShader = `assets\box\lighting_instancing.vs`
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
			loadedModel := rl.LoadModel(modelData.ModelComp.ModelDataLocation)

			// Load Texture
			loadedTexture := rl.LoadTexture(modelData.ModelComp.TextureDataLocation)
			loadedModel.Materials.Maps.Texture = loadedTexture

			// Load Shader
			shader := rl.LoadShader(modelData.ModelComp.VertexShader, modelData.ModelComp.FragmentShader)
			shader.UpdateLocation(rl.LocMatrixMvp, rl.GetShaderLocation(shader, "mvp"))
			shader.UpdateLocation(rl.LocVectorView, rl.GetShaderLocation(shader, "viewPos"))
			shader.UpdateLocation(rl.LocMatrixModel, rl.GetShaderLocationAttrib(shader, "instanceTransform"))
			loadedModel.Materials.Shader = shader

			// Create Records
			modelData.ModelComp.Model = loadedModel
			modelData.ModelComp.LoadedModel = true
			ts.loadedModels[modelData.ModelComp.HashID] = &loadedModel
		} else if !modelData.ModelComp.LoadedModel { // Model already loaded in map but not set on component
			modelData.ModelComp.Model = *loadedMapModel
		}
	}

}
