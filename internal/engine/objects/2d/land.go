package objects2d

import (
	"rouge/internal/ecs"
	components2d "rouge/internal/engine/components/2d"
)

func NewLand(width, height int, positionX, positionY float32, data []int) []ecs.Component {
	LandComponent := components2d.NewLandComponent(width, height, data)
	TransformationComponent := components2d.NewTransformation2DComponent()
	TransformationComponent.Position.X = positionX
	TransformationComponent.Position.Y = positionY

	return []ecs.Component{
		LandComponent,
		TransformationComponent,
	}
}
