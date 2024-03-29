package components3d

import (
	"github.com/jtheiss19/game-raylib/pkg/ecs"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type ChunkTransformation3DComponent struct {
	*ecs.BaseComponent
	Position rl.Vector3
	Scale    rl.Vector3
	Rotation rl.Vector3
	Forward  rl.Vector3
}

func NewChunkTransformation3DComponent() *ChunkTransformation3DComponent {

	return &ChunkTransformation3DComponent{
		BaseComponent: &ecs.BaseComponent{},
		Position:      rl.Vector3{X: 0, Y: 0, Z: 0},
		Scale:         rl.Vector3{X: 1, Y: 1, Z: 1},
		Rotation:      rl.Vector3{X: 0, Y: 0, Z: 0},
		Forward:       rl.Vector3{X: 0, Y: 0, Z: 0},
	}
}
