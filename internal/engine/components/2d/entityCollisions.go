package components2d

import "github.com/jtheiss19/game-raylib/internal/ecs"

type CollisionComponent struct {
	*ecs.BaseComponent
}

func NewCollision2DComponent() *CollisionComponent {
	return &CollisionComponent{
		BaseComponent: &ecs.BaseComponent{},
	}
}
