package components3d

import (
	"rouge/internal/ecs"
)

type Chunk3DComponent struct {
	*ecs.BaseComponent
	Width, Length int
	Data          []int
}

func NewChunk3DComponent(Width, Length int, Data []int) *Chunk3DComponent {
	return &Chunk3DComponent{
		BaseComponent: &ecs.BaseComponent{},
		Width:         Width,
		Length:        Length,
		Data:          Data,
	}
}
