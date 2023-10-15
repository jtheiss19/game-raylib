package components3d

import (
	"github.com/jtheiss19/game-raylib/internal/ecs"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Model3DComponent struct {
	*ecs.BaseComponent
	HashID              string
	ModelDataLocation   string
	TextureDataLocation string
	Model               rl.Model
	LoadedModel         bool
}

func NewModel3DComponent(ModelDataLocation, TextureDataLocation string) *Model3DComponent {

	return &Model3DComponent{
		BaseComponent:       &ecs.BaseComponent{},
		HashID:              ModelDataLocation + TextureDataLocation,
		ModelDataLocation:   ModelDataLocation,
		TextureDataLocation: TextureDataLocation,
		Model:               rl.Model{},
		LoadedModel:         false,
	}
}
