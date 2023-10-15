package example

import (
	"github.com/jtheiss19/game-raylib/internal/ecs"

	"github.com/sirupsen/logrus"
)

type exampleSystem struct {
	*ecs.BaseSystem
}

func NewExampleSystem() *exampleSystem {
	return &exampleSystem{
		BaseSystem: &ecs.BaseSystem{},
	}
}

// Comps
type RequiredExampleComps struct {
	TypeOne []*RequireTypeOne
}

type RequireTypeOne struct {
	Somestuff *ExampleStruct
}

func (ts *exampleSystem) GetRequiredComponents() interface{} {
	return &RequiredExampleComps{
		TypeOne: []*RequireTypeOne{{
			Somestuff: &ExampleStruct{},
		}},
	}
}

// Functionality
func (ts *exampleSystem) Update(dt float32) {
	entities, ok := ts.TrackedEntities.(*RequiredExampleComps)
	if !ok {
		logrus.Error("could not update system, bad tracked entities")
		return
	}
	logrus.Info(entities.TypeOne[0].Somestuff)
}

func (ts *exampleSystem) Initilizer() {
	// Some initing Code
	logrus.Info("initing example system")
}
