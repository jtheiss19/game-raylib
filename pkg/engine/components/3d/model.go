package components3d

import (
	"strconv"

	"github.com/jtheiss19/game-raylib/pkg/ecs"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type TextureType struct {
	TexturePath        string
	TextureFrameWidth  int
	TextureFrameHeight int
}

var (
	CRATE_TEX = TextureType{
		TexturePath:        `assets\box\crate.jpg`,
		TextureFrameWidth:  1,
		TextureFrameHeight: 1,
	}
	GRASS_TEX = TextureType{
		TexturePath:        `assets\box\grass.jpg`,
		TextureFrameWidth:  1,
		TextureFrameHeight: 1,
	}
	IMAGE_TEX = TextureType{
		TexturePath:        `assets\box\image.jpg`,
		TextureFrameWidth:  5,
		TextureFrameHeight: 5,
	}
)

type ObjectType string

const (
	CUBE_OBJ ObjectType = `assets\box\cube.glb`
	RAMP_OBJ ObjectType = `assets\box\ramp.glb`
	SLAB_OBJ ObjectType = `assets\box\slab.glb`
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

func NewModel3DComponent(ModelDataLocation ObjectType, Texture TextureType, FragmentShader FragmentShader, VertexShader VertexShader, textureFrame int) *Model3DComponent {

	return &Model3DComponent{
		BaseComponent:       &ecs.BaseComponent{},
		HashID:              string(ModelDataLocation) + string(Texture.TexturePath) + string(FragmentShader) + string(VertexShader) + strconv.Itoa(textureFrame),
		ModelDataLocation:   ModelDataLocation,
		TextureDataLocation: Texture,
		TextureFrame:        textureFrame - 1,
		FragmentShader:      FragmentShader,
		VertexShader:        VertexShader,
		Model:               rl.Model{},
		LoadedModel:         false,
	}
}
