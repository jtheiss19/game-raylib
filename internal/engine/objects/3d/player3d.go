package objects3d

import (
	"rouge/internal/ecs"
	"rouge/internal/engine/components"
	components3d "rouge/internal/engine/components/3d"
)

func New3DPlayer(playerID ecs.ID) []ecs.Component {
	cameraComponent := components3d.NewCamera3DComponent()
	TransformationComponent := components3d.NewTransformation3DComponent()
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
