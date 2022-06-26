package components

import (
	"rouge/internal/ecs"
)

type NetworkComponent struct {
	*ecs.BaseComponent
}

func NewNetworkComponent() *NetworkComponent {
	return &NetworkComponent{
		BaseComponent: &ecs.BaseComponent{},
	}
}
