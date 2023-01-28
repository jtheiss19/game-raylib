package components

import (
	"rouge/internal/network"
)

func init() {
	network.RegisterType(InputComponent{})
	network.RegisterType(NetworkComponent{})
	network.RegisterType(ModelComponent{})
	network.RegisterType(PlayerComponent{})
}
