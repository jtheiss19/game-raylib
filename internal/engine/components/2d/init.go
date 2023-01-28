package components2d

import "rouge/internal/network"

func init() {
	network.RegisterType(Camera2DComponent{})
	network.RegisterType(Transformation2DComponent{})
}
