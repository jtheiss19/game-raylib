package objects3d

import (
	"github.com/jtheiss19/game-raylib/internal/ecs"
	"github.com/jtheiss19/game-raylib/internal/engine/components"
	components3d "github.com/jtheiss19/game-raylib/internal/engine/components/3d"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func NewChunk(x, y, z float32) []ecs.Component {
	TransformationComponent := components3d.NewTransformation3DComponent()
	TransformationComponent.Position = rl.NewVector3(x, y, z)
	TransformationComponent.Scale.X = 0.5
	TransformationComponent.Scale.Y = 0.5
	TransformationComponent.Scale.Z = 0.5
	NetworkComponent := components.NewNetworkComponent()
	ModelComponent := components3d.NewModel3DComponent(`assets\box\Crate.obj`, string(CRATE))

	length := 10
	width := 10
	data := []components3d.Block{}
	for x := 0; x < length; x++ {
		for z := 0; z < width; z++ {
			data = append(data, components3d.Block{
				RelativePosition: rl.Vector3{
					X: float32(x),
					Y: 0,
					Z: float32(z),
				},
				BlockType: 1,
			})
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
