package objects3d

import (
	"github.com/jtheiss19/game-raylib/internal/ecs"
	"github.com/jtheiss19/game-raylib/internal/engine/components"
	components3d "github.com/jtheiss19/game-raylib/internal/engine/components/3d"
	"github.com/sirupsen/logrus"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func NewBlock3d(x, y, z float32, blockType components3d.TextureType, textureFrame int) []ecs.Component {
	TransformationComponent := components3d.NewTransformation3DComponent()
	TransformationComponent.Position = rl.NewVector3(x, y, z)
	TransformationComponent.Scale.X = 0.5
	TransformationComponent.Scale.Y = 0.5
	TransformationComponent.Scale.Z = 0.5
	if textureFrame > blockType.TextureFrameHeight*blockType.TextureFrameWidth || textureFrame < 1 {
		logrus.Errorf("Block 3d attempted to be created with textureFrame %v outside spec %v", textureFrame, blockType.TextureFrameHeight*blockType.TextureFrameWidth)
		return []ecs.Component{}
	}

	ModelComponent := components3d.NewModel3DComponent(
		components3d.CRATE_OBJ,
		blockType,
		"",
		"",
		textureFrame,
	)
	NetworkComponent := components.NewNetworkComponent()

	return []ecs.Component{
		TransformationComponent,
		ModelComponent,
		NetworkComponent,
	}
}
