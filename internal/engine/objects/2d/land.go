package objects2d

import (
	"rouge/internal/ecs"
	components2d "rouge/internal/engine/components/2d"
)

func NewLand(x, y float32) []ecs.Component {
	LandComponent := components2d.NewLandComponent()

	return []ecs.Component{
		LandComponent,
	}
}
