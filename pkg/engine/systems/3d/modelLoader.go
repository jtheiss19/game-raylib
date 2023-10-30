package systems3d

import (
	"fmt"
	"math"
	"unsafe"

	"github.com/jtheiss19/game-raylib/pkg/ecs"
	components3d "github.com/jtheiss19/game-raylib/pkg/engine/components/3d"

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
			loadedImage := rl.LoadImage(modelData.ModelComp.TextureDataLocation.TexturePath)
			// Handle Sprite Sheets
			textureWidth := float32(loadedImage.Width) / float32(modelData.ModelComp.TextureDataLocation.TextureFrameWidth)
			textureHeight := float32(loadedImage.Height) / float32(modelData.ModelComp.TextureDataLocation.TextureFrameHeight)
			xStart := modelData.ModelComp.TextureFrame % modelData.ModelComp.TextureDataLocation.TextureFrameWidth
			yStart := math.Floor(float64(modelData.ModelComp.TextureFrame) / float64(modelData.ModelComp.TextureDataLocation.TextureFrameWidth))
			rl.ImageCrop(loadedImage, rl.Rectangle{
				X:      textureWidth * float32(xStart),
				Y:      textureHeight * float32(yStart),
				Width:  textureWidth,
				Height: textureHeight,
			})
			generatedTexture := rl.LoadTextureFromImage(loadedImage)
			rl.SetModelMeshMaterial(&loadedModel, 0, 0)
			rl.SetMaterialTexture(loadedModel.Materials, rl.MapDiffuse, generatedTexture)

			// Load Shader
			if modelData.ModelComp.VertexShader != "" && modelData.ModelComp.FragmentShader != "" {
				shader := rl.LoadShader(string(modelData.ModelComp.VertexShader), string(modelData.ModelComp.FragmentShader))
				shader.UpdateLocation(rl.LocMatrixMvp, rl.GetShaderLocation(shader, "mvp"))
				shader.UpdateLocation(rl.LocVectorView, rl.GetShaderLocation(shader, "viewPos"))
				shader.UpdateLocation(rl.LocMatrixModel, rl.GetShaderLocationAttrib(shader, "instanceTransform"))

				ambientLoc := rl.GetShaderLocation(shader, "ambient")
				rl.SetShaderValue(shader, ambientLoc, []float32{0.2, 0.2, 0.2, 1.0}, rl.ShaderUniformVec4)

				lightColor := rl.Gray
				NewLight(LightTypeDirectional, rl.NewVector3(13, 50.0, 13), rl.Vector3Zero(), lightColor, shader)

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

type LightType int32

const (
	LightTypeDirectional LightType = iota
	LightTypePoint
)

type Light struct {
	shader    rl.Shader
	lightType LightType
	position  rl.Vector3
	target    rl.Vector3
	color     rl.Color
	enabled   int32
	// shader locations
	enabledLoc int32
	typeLoc    int32
	posLoc     int32
	targetLoc  int32
	colorLoc   int32
}

const maxLightsCount = 4

var lightCount = 0

func NewLight(
	lightType LightType,
	position, target rl.Vector3,
	color rl.Color,
	shader rl.Shader) Light {
	light := Light{
		shader: shader,
	}
	if lightCount < maxLightsCount {
		light.enabled = 1
		light.lightType = lightType
		light.position = position
		light.target = target
		light.color = color
		light.enabledLoc = rl.GetShaderLocation(shader, fmt.Sprintf("lights[%d].enabled", lightCount))
		light.typeLoc = rl.GetShaderLocation(shader, fmt.Sprintf("lights[%d].type", lightCount))
		light.posLoc = rl.GetShaderLocation(shader, fmt.Sprintf("lights[%d].position", lightCount))
		light.targetLoc = rl.GetShaderLocation(shader, fmt.Sprintf("lights[%d].target", lightCount))
		light.colorLoc = rl.GetShaderLocation(shader, fmt.Sprintf("lights[%d].color", lightCount))
		light.UpdateValues()
		lightCount++
	}
	return light
}

func (lt *Light) UpdateValues() {
	// Send to shader light enabled state and type
	rl.SetShaderValue(lt.shader, lt.enabledLoc, unsafe.Slice((*float32)(unsafe.Pointer(&lt.enabled)), 4), rl.ShaderUniformInt)
	rl.SetShaderValue(lt.shader, lt.typeLoc, unsafe.Slice((*float32)(unsafe.Pointer(&lt.lightType)), 4), rl.ShaderUniformInt)

	// Send to shader light position values
	rl.SetShaderValue(lt.shader, lt.posLoc, []float32{lt.position.X, lt.position.Y, lt.position.Z}, rl.ShaderUniformVec3)

	// Send to shader light target target values
	rl.SetShaderValue(lt.shader, lt.targetLoc, []float32{lt.target.X, lt.target.Y, lt.target.Z}, rl.ShaderUniformVec3)

	// Send to shader light color values
	rl.SetShaderValue(lt.shader, lt.colorLoc,
		[]float32{float32(lt.color.R) / 255, float32(lt.color.G) / 255, float32(lt.color.B) / 255, float32(lt.color.A) / 255},
		rl.ShaderUniformVec4)
}
