package objects2d

import (
	"rouge/internal/ecs"
	"rouge/internal/engine/components"
	components2d "rouge/internal/engine/components/2d"
)

func New2DPlayer(playerID ecs.ID) []ecs.Component {
	cameraComponent := components2d.NewCamera2DComponent()
	TransformationComponent := components2d.NewTransformation2DComponent()
	InputComponent := components.NewInputComponent()
	NetworkComponent := components.NewNetworkComponent()
	ModelComponent := components.NewModelComponent()
	playerComponent := components.NewPlayerComponent(playerID)

	return []ecs.Component{
		cameraComponent,
		TransformationComponent,
		InputComponent,
		NetworkComponent,
		playerComponent,
		ModelComponent,
	}
}
