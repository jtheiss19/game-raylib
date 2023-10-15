package components2d

import "github.com/jtheiss19/game-raylib/internal/network"

func init() {
	network.RegisterType(Camera2DComponent{})
	network.RegisterType(Transformation2DComponent{})
}
