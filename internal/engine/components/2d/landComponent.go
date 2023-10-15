package components2d

import (
	"github.com/jtheiss19/game-raylib/internal/ecs"
)

type LandComponent struct {
	*ecs.BaseComponent
	Width, Height int
	Data          []int
}

func NewLandComponent(width, height int, data []int) *LandComponent {
	return &LandComponent{
		BaseComponent: &ecs.BaseComponent{},
		Width:         width,
		Height:        height,
		Data:          data,
	}
}

func (lc *LandComponent) SetData(x, y int, value int) {

}

func (lc *LandComponent) GetData(x, y int) int {
	return 0
}
