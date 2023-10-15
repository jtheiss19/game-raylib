package components

import (
	"github.com/jtheiss19/game-raylib/internal/ecs"
)

type NetworkComponent struct {
	*ecs.BaseComponent
}

func NewNetworkComponent() *NetworkComponent {
	return &NetworkComponent{
		BaseComponent: &ecs.BaseComponent{},
	}
}
