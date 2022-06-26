package components

import (
	"rouge/internal/ecs"
)

type ModelComponent struct {
	*ecs.BaseComponent
}

func NewModelComponent() *ModelComponent {
	return &ModelComponent{
		BaseComponent: &ecs.BaseComponent{},
	}
}
