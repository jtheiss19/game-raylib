package components2d

import (
	"rouge/internal/ecs"
)

type LandComponent struct {
	*ecs.BaseComponent
	width, height int
	data          []int
}

func NewLandComponent(width, height int) *LandComponent {
	return &LandComponent{
		BaseComponent: &ecs.BaseComponent{},
		width:         width,
		height:        height,
		data:          make([]int, width*height),
	}
}

func (lc *LandComponent) SetData(x, y int, value int) {

}

func (lc *LandComponent) GetData(x, y int) int {
	return 0
}
