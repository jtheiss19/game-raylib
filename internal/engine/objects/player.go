package objects

import (
	"rouge/internal/ecs"
	"rouge/internal/engine/components"
)

func NewPlayer(playerID ecs.ID) []ecs.Component {
	cameraComponent := components.NewCameraComponent()
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
