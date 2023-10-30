package systems3d

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/jtheiss19/game-raylib/pkg/ecs"
	components3d "github.com/jtheiss19/game-raylib/pkg/engine/components/3d"

	"github.com/sirupsen/logrus"
)

type CollisionSystem struct {
	*ecs.BaseSystem
}

func NewCollisionSystem() *CollisionSystem {
	return &CollisionSystem{
		BaseSystem: &ecs.BaseSystem{},
	}
}

// Comps
type RequiredCollisionSystemComps struct {
	Recievers []*RequiredCollisionRecievers
	Producers []*RequiredCollisionProducers
}

type RequiredCollisionRecievers struct {
	Collidables *components3d.CollisionReceiver3DComponent
}

type RequiredCollisionProducers struct {
	Collidables *components3d.CollisionProducer3DComponent
}

func (ts *CollisionSystem) GetRequiredComponents() interface{} {
	return &RequiredCollisionSystemComps{
		Recievers: []*RequiredCollisionRecievers{{}},
		Producers: []*RequiredCollisionProducers{{}},
	}
}

// Functionality
func (ts *CollisionSystem) Update(dt float32) {
	entities, ok := ts.TrackedEntities.(*RequiredCollisionSystemComps)
	if !ok {
		logrus.Error("could not update system, bad tracked entities")
		return
	}

	for _, producer := range entities.Producers {
		var closestCollision float32 = 10000.0
		producer.Collidables.Collision.Hit = false
		for _, receiver := range entities.Recievers {
			collision := rl.GetRayCollisionBox(producer.Collidables.Ray, receiver.Collidables.BoundingBox)
			if collision.Hit && closestCollision > collision.Distance {
				producer.Collidables.Collision = collision
			}
		}
	}
}

func (ts *CollisionSystem) Initilizer() {
}
