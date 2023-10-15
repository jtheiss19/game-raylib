package objects3d

import (
	"rouge/internal/ecs"
	"rouge/internal/engine/components"
	components3d "rouge/internal/engine/components/3d"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type GridBlockType string

var (
	CRATE GridBlockType = `assets\box\crate_1.jpg`
	GRASS GridBlockType = `assets\box\grass.jpg`
)

func NewGridBlock3d(x, y, z int, blockType GridBlockType) []ecs.Component {
	TransformationComponent := components3d.NewGridTransformation3DComponent()
	TransformationComponent.Position = rl.NewVector3(float32(x), float32(y), float32(z))
	TransformationComponent.Scale.X = 0.5
	ModelComponent := components3d.NewModel3DComponent(`assets\box\Crate1.obj`, string(blockType))
	NetworkComponent := components.NewNetworkComponent()

	return []ecs.Component{
		TransformationComponent,
		ModelComponent,
		NetworkComponent,
	}
}
