package objects3d

import (
	"github.com/jtheiss19/game-raylib/internal/ecs"
	"github.com/jtheiss19/game-raylib/internal/engine/components"
	components3d "github.com/jtheiss19/game-raylib/internal/engine/components/3d"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func New3DPlayer(playerID ecs.ID, x, y, z int) []ecs.Component {
	cameraComponent := components3d.NewCamera3DComponent()
	TransformationComponent := components3d.NewTransformation3DComponent()
	TransformationComponent.Position = rl.NewVector3(float32(x), float32(y), float32(z))
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
