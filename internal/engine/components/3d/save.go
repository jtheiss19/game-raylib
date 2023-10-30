package components3d

import (
	"github.com/jtheiss19/game-raylib/internal/ecs"
)

type Save3DComponent struct {
	*ecs.BaseComponent
}

func NewSave3DComponent() *Save3DComponent {
	return &Save3DComponent{
		BaseComponent: &ecs.BaseComponent{},
	}
}
