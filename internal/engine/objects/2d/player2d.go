package objects2d

import (
	"github.com/jtheiss19/game-raylib/internal/ecs"
	"github.com/jtheiss19/game-raylib/internal/engine/components"
	components2d "github.com/jtheiss19/game-raylib/internal/engine/components/2d"
)

func New2DPlayer(playerID ecs.ID) []ecs.Component {
	cameraComponent := components2d.NewCamera2DComponent()
	TransformationComponent := components2d.NewTransformation2DComponent()
	InputComponent := components.NewInputComponent()
	NetworkComponent := components.NewNetworkComponent()
	ModelComponent := components.NewModelComponent()
	playerComponent := components.NewPlayerComponent(playerID)
	collisionComponent := components2d.NewCollision2DComponent()

	return []ecs.Component{
		cameraComponent,
		TransformationComponent,
		InputComponent,
		NetworkComponent,
		playerComponent,
		ModelComponent,
		collisionComponent,
	}
}
