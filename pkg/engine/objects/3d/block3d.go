package objects3d

import (
	"github.com/jtheiss19/game-raylib/pkg/ecs"
	"github.com/jtheiss19/game-raylib/pkg/engine/components"
	components3d "github.com/jtheiss19/game-raylib/pkg/engine/components/3d"
	"github.com/sirupsen/logrus"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func NewBlock3d(x, y, z float32, blockType components3d.TextureType, objectType components3d.ObjectType, textureFrame int) []ecs.Component {
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
		objectType,
		blockType,
		"",
		"",
		textureFrame,
	)
	NetworkComponent := components.NewNetworkComponent()
	Collision := components3d.NewCollisionReceiver3DComponent()
	minBounding := rl.Vector3Subtract(TransformationComponent.Position, rl.Vector3{X: 0.5, Y: 0.5, Z: 0.5})
	maxBounding := rl.Vector3Add(TransformationComponent.Position, rl.Vector3{X: 0.5, Y: 0.5, Z: 0.5})
	Collision.BoundingBox = rl.NewBoundingBox(minBounding, maxBounding)
	saveComp := components3d.NewSave3DComponent()

	return []ecs.Component{
		TransformationComponent,
		ModelComponent,
		NetworkComponent,
		Collision,
		saveComp,
	}
}
