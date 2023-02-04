package components2d

import "rouge/internal/ecs"

type CollisionComponent struct {
	*ecs.BaseComponent
}

func NewCollision2DComponent() *CollisionComponent {
	return &CollisionComponent{
		BaseComponent: &ecs.BaseComponent{},
	}
}
