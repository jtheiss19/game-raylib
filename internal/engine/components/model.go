package components

import (
	"github.com/jtheiss19/game-raylib/internal/ecs"
)

type ModelComponent struct {
	*ecs.BaseComponent
}

func NewModelComponent() *ModelComponent {
	return &ModelComponent{
		BaseComponent: &ecs.BaseComponent{},
	}
}
