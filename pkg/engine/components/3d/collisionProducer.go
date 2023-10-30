package components3d

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/jtheiss19/game-raylib/pkg/ecs"
)

type CollisionProducer3DComponent struct {
	*ecs.BaseComponent
	Ray       rl.Ray
	Collision rl.RayCollision
}

func NewCollisionProducer3DComponent() *CollisionProducer3DComponent {
	return &CollisionProducer3DComponent{
		BaseComponent: &ecs.BaseComponent{},
		Ray: rl.Ray{
			Position: rl.Vector3{
				X: 0,
				Y: 0,
				Z: 0,
			},
			Direction: rl.Vector3{
				X: 1,
				Y: 0,
				Z: 0,
			},
		},
		Collision: rl.RayCollision{},
	}
}
