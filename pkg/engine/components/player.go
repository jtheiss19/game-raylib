package components

import (
	"github.com/jtheiss19/game-raylib/pkg/ecs"
)

type PlayerComponent struct {
	*ecs.BaseComponent
	PlayerID ecs.ID
}

func NewPlayerComponent(playerID ecs.ID) *PlayerComponent {
	return &PlayerComponent{
		BaseComponent: &ecs.BaseComponent{},
		PlayerID:      playerID,
	}
}
