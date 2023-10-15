package components3d

import "github.com/jtheiss19/game-raylib/internal/network"

func init() {
	network.RegisterType(Camera3DComponent{})
	network.RegisterType(Transformation3DComponent{})
}
