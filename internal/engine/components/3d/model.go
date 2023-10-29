package components3d

import (
	"strconv"

	"github.com/jtheiss19/game-raylib/internal/ecs"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type TextureType string

const (
	CRATE_TEX TextureType = `assets\box\crate.jpg`
	GRASS_TEX TextureType = `assets\box\grass.jpg`
	IMAGE_TEX TextureType = `assets\box\image.jpg`
)

type ObjectType string

const (
	CRATE_OBJ ObjectType = `assets\box\Crate.obj`
)

type FragmentShader string

const (
	INSTANCED_FRAG FragmentShader = `assets\box\lighting.fs`
)

type VertexShader string

const (
	INSTANCED_VERT VertexShader = `assets\box\lighting_instancing.vs`
)

type Model3DComponent struct {
	*ecs.BaseComponent
	HashID              string
	ModelDataLocation   ObjectType
	TextureDataLocation TextureType
	TextureFrame        int
	FragmentShader      FragmentShader
	VertexShader        VertexShader
	Model               rl.Model
	LoadedModel         bool
}

func NewModel3DComponent(ModelDataLocation ObjectType, TextureDataLocation TextureType, FragmentShader FragmentShader, VertexShader VertexShader, textureBody int) *Model3DComponent {

	return &Model3DComponent{
		BaseComponent:       &ecs.BaseComponent{},
		HashID:              string(ModelDataLocation) + string(TextureDataLocation) + string(FragmentShader) + string(VertexShader) + strconv.Itoa(textureBody),
		ModelDataLocation:   ModelDataLocation,
		TextureDataLocation: TextureDataLocation,
		TextureFrame:        textureBody,
		FragmentShader:      FragmentShader,
		VertexShader:        VertexShader,
		Model:               rl.Model{},
		LoadedModel:         false,
	}
}
