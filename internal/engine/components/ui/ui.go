package componentsui

import (
	"rouge/internal/ecs"
)

type UIBoxComponent struct {
	*ecs.BaseComponent
}

func NewUIBoxComponent() *UIBoxComponent {
	return &UIBoxComponent{
		BaseComponent: &ecs.BaseComponent{},
	}
}
