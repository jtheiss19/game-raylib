package components2d

import (
	"rouge/internal/ecs"
)

type LandComponent struct {
	*ecs.BaseComponent
}

func NewLandComponent() *LandComponent {
	return &LandComponent{
		BaseComponent: &ecs.BaseComponent{},
	}
}
