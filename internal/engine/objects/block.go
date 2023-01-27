package objects

import (
	"rouge/internal/ecs"
	"rouge/internal/engine/components"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func NewBlock(x, y, z float32) []ecs.Component {
	TransformationComponent := components.NewTransformationComponent()
	TransformationComponent.Position = rl.NewVector3(x, y, z)
	NetworkComponent := components.NewNetworkComponent()

	return []ecs.Component{
		TransformationComponent,
		NetworkComponent,
	}
}
