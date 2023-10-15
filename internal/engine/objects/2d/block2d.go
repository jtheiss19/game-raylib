package objects2d

import (
	"github.com/jtheiss19/game-raylib/internal/ecs"
	"github.com/jtheiss19/game-raylib/internal/engine/components"
	components2d "github.com/jtheiss19/game-raylib/internal/engine/components/2d"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func NewBlock2d(x, y float32) []ecs.Component {
	TransformationComponent := components2d.NewTransformation2DComponent()
	TransformationComponent.Position = rl.NewVector2(x, y)
	ModelComponent := components.NewModelComponent()
	NetworkComponent := components.NewNetworkComponent()

	return []ecs.Component{
		TransformationComponent,
		NetworkComponent,
		ModelComponent,
	}
}
