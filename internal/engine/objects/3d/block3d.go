package objects3d

import (
	"github.com/jtheiss19/game-raylib/internal/ecs"
	"github.com/jtheiss19/game-raylib/internal/engine/components"
	components3d "github.com/jtheiss19/game-raylib/internal/engine/components/3d"
	systems3d "github.com/jtheiss19/game-raylib/internal/engine/systems/3d"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func NewBlock3d(x, y, z float32, blockType systems3d.TextureType) []ecs.Component {
	TransformationComponent := components3d.NewTransformation3DComponent()
	TransformationComponent.Position = rl.NewVector3(x, y, z)
	ModelComponent := components3d.NewModel3DComponent(
		string(systems3d.CRATE_OBJ),
		string(blockType),
		string(systems3d.INSTANCED_FRAG),
		string(systems3d.INSTANCED_VERT),
	)
	NetworkComponent := components.NewNetworkComponent()

	return []ecs.Component{
		TransformationComponent,
		ModelComponent,
		NetworkComponent,
	}
}
