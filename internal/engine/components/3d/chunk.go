package components3d

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/jtheiss19/game-raylib/internal/ecs"
)

type Block struct {
	RelativePosition rl.Vector3
	BlockType        int
}

type Chunk3DComponent struct {
	*ecs.BaseComponent
	Width, Length int
	Data          []Block
}

func NewChunk3DComponent(Width, Length int, Data []Block) *Chunk3DComponent {
	return &Chunk3DComponent{
		BaseComponent: &ecs.BaseComponent{},
		Width:         Width,
		Length:        Length,
		Data:          Data,
	}
}
