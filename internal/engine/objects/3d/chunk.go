package objects3d

import (
	"rouge/internal/ecs"
	"rouge/internal/engine/components"
	components3d "rouge/internal/engine/components/3d"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func NewChunk(x, y, z float32) []ecs.Component {
	TransformationComponent := components3d.NewTransformation3DComponent()
	TransformationComponent.Position = rl.NewVector3(x, y, z)
	TransformationComponent.Scale.X = 0.5
	NetworkComponent := components.NewNetworkComponent()
	ModelComponent := components3d.NewModel3DComponent(`assets\box\Crate.obj`, string(GRASS))

	length := 10
	width := 10
	data := make([]int, length*width)
	for i := 0; i < length; i++ {
		for x := 0; x < width; x++ {
			if i == 0 {
				data[i*length+x] = 1
			} else {
				data[i*length+x] = 0
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
