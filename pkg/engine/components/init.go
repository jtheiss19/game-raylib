package components

import (
	"github.com/jtheiss19/game-raylib/pkg/network"
)

func init() {
	network.RegisterType(InputComponent{})
	network.RegisterType(NetworkComponent{})
	network.RegisterType(ModelComponent{})
	network.RegisterType(PlayerComponent{})
}
