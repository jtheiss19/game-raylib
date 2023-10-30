package components3d

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/jtheiss19/game-raylib/pkg/ecs"
)

type CollisionReceiver3DComponent struct {
	*ecs.BaseComponent
	BoundingBox rl.BoundingBox
}

func NewCollisionReceiver3DComponent() *CollisionReceiver3DComponent {
	return &CollisionReceiver3DComponent{
		BaseComponent: &ecs.BaseComponent{},
		BoundingBox:   rl.BoundingBox{},
	}
}
