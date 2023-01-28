package systems2d

import (
	"fmt"
	"rouge/internal/ecs"
	components2d "rouge/internal/engine/components/2d"

	"github.com/sirupsen/logrus"
)

type LandRenderingSystem struct {
	*ecs.BaseSystem
}

func NewLandRenderingSystem() *LandRenderingSystem {
	return &LandRenderingSystem{
		BaseSystem: &ecs.BaseSystem{},
	}
}

// Comps
type RequiredLandRenderingSystemComps struct {
	Land []*RequireLand
}

type RequireLand struct {
	LandComponent *components2d.LandComponent
}

func (ts *LandRenderingSystem) GetRequiredComponents() interface{} {
	return &RequiredLandRenderingSystemComps{
		Land: []*RequireLand{{}},
	}
}

// Functionality
func (ts *LandRenderingSystem) Update(dt float32) {
	entities, ok := ts.TrackedEntities.(*RequiredLandRenderingSystemComps)
	if !ok {
		logrus.Error("could not update system, bad tracked entities")
		return
	}

	fmt.Println(entities)
}

func (ts *LandRenderingSystem) Initilizer() {
}
