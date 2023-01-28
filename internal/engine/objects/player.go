package objects

import (
	"rouge/internal/ecs"
	"rouge/internal/engine/components"
)

func New3DPlayer(playerID ecs.ID) []ecs.Component {
	cameraComponent := components.NewCamera3DComponent()
	TransformationComponent := components.NewTransformationComponent()
	InputComponent := components.NewInputComponent()
	NetworkComponent := components.NewNetworkComponent()
	playerComponent := components.NewPlayerComponent(playerID)

	return []ecs.Component{
		cameraComponent,
		TransformationComponent,
		InputComponent,
		NetworkComponent,
		playerComponent,
	}
}
