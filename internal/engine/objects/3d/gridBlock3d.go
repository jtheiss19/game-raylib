package objects3d

import (
	"rouge/internal/ecs"
	"rouge/internal/engine/components"
	components3d "rouge/internal/engine/components/3d"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func NewGridBlock3d(x, y, z int) []ecs.Component {
	TransformationComponent := components3d.NewGridTransformation3DComponent()
	TransformationComponent.Position = rl.NewVector3(float32(x), float32(y), float32(z))
	NetworkComponent := components.NewNetworkComponent()

	return []ecs.Component{
		TransformationComponent,
		NetworkComponent,
	}
}
