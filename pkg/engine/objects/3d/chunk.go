package objects3d

import (
	"github.com/jtheiss19/game-raylib/pkg/ecs"
	"github.com/jtheiss19/game-raylib/pkg/engine/components"
	components3d "github.com/jtheiss19/game-raylib/pkg/engine/components/3d"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func NewChunk(x, y, z float32, length, width, height int, blockType components3d.TextureType, TextureFrame int) []ecs.Component {
	TransformationComponent := components3d.NewChunkTransformation3DComponent()
	TransformationComponent.Position = rl.NewVector3(x, y, z)
	TransformationComponent.Scale.X = 0.5
	TransformationComponent.Scale.Y = 0.5
	TransformationComponent.Scale.Z = 0.5
	NetworkComponent := components.NewNetworkComponent()
	ModelComponent := components3d.NewModel3DComponent(
		components3d.CUBE_OBJ,
		blockType,
		components3d.INSTANCED_FRAG,
		components3d.INSTANCED_VERT,
		TextureFrame,
	)

	data := []components3d.Block{}
	for x := 0; x < length; x++ {
		for z := 0; z < width; z++ {
			for y := 0; y < height; y++ {
				data = append(data, components3d.Block{
					RelativePosition: rl.Vector3{
						X: float32(x),
						Y: float32(y),
						Z: float32(z),
					},
					BlockType: 1,
				})
			}
		}
	}
	ChunkComponent := components3d.NewChunk3DComponent(10, 10, data)

	return []ecs.Component{
		TransformationComponent,
		ChunkComponent,
		ModelComponent,
		NetworkComponent,
	}
}
