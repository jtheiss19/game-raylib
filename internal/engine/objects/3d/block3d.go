package objects3d

import (
	"github.com/jtheiss19/game-raylib/internal/ecs"
	"github.com/jtheiss19/game-raylib/internal/engine/components"
	components3d "github.com/jtheiss19/game-raylib/internal/engine/components/3d"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type BlockType string

var (
	CRATE BlockType = `assets\box\crate.jpg`
	GRASS BlockType = `assets\box\grass.jpg`
)

func NewBlock3d(x, y, z float32, blockType BlockType) []ecs.Component {
	TransformationComponent := components3d.NewTransformation3DComponent()
	TransformationComponent.Position = rl.NewVector3(x, y, z)
	ModelComponent := components3d.NewModel3DComponent(`assets\box\Crate1.obj`, string(blockType))
	NetworkComponent := components.NewNetworkComponent()

	return []ecs.Component{
		TransformationComponent,
		ModelComponent,
		NetworkComponent,
	}
}
