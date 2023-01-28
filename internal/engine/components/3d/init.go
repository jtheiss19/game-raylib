package components3d

import "rouge/internal/network"

func init() {
	network.RegisterType(Camera3DComponent{})
	network.RegisterType(Transformation3DComponent{})
}
