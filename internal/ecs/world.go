package ecs

import (
	"reflect"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

var (
	activeWorld *World
)

type World struct {
	systems      []system
	entityLookup map[ID]map[reflect.Type]Component
}

// NewWorld generates a new empty World and sets
// it as the active World
func NewWorld() *World {
	activeWorld = &World{
		systems:      []system{},
		entityLookup: map[ID]map[reflect.Type]Component{},
	}
	return activeWorld
}

// GetActiveWorld returns the active World
func GetActiveWorld() *World {
	return activeWorld
}

// UpdateSystems preforms an update on every system
// inside the World. Each system is updated with time
// dt.
func (wrld *World) UpdateSystems(dt float32) {
	for _, system := range wrld.systems {
		system.Update(dt)
	}
}

// AddComponent adds a new component to the World.
func (wrld *World) AddComponent(comp Component) {
	logrus.Trace("adding component")
	id, err := comp.GetComponentID()
	if err != nil {
		id = comp.generateComponentID()
	}

	if comps, ok := wrld.entityLookup[id]; ok {
		comps[reflect.TypeOf(comp)] = comp
		wrld.entityLookup[id] = comps
		wrld.checkCompsForNewSystemMatch(comps)
	} else {
		compsMap := map[reflect.Type]Component{}
		compsMap[reflect.TypeOf(comp)] = comp
		wrld.entityLookup[id] = compsMap
		wrld.checkCompsForNewSystemMatch(compsMap)
	}
}

// RemoveCompoent removes a compoent from the World
func (wrld *World) RemoveCompoent(comp Component) {
	logrus.Info("removing component")
	id, _ := comp.GetComponentID()

	if comps, ok := wrld.entityLookup[id]; ok {
		newComps := map[reflect.Type]Component{}
		for _, compItem := range comps {
			if reflect.TypeOf(compItem) != reflect.TypeOf(comp) {
				newComps[reflect.TypeOf(comp)] = comp
			}
		}
		wrld.entityLookup[id] = newComps

		for _, system := range wrld.systems {
			system.removeEntity(id)
		}
	}
}

// AddEntity adds a series of components to the World.
// each component is assigned a UUID from the first component
// linking them togeather.
func (wrld *World) AddEntity(comps []Component) ID {
	id := ID(uuid.New().String())
	logrus.Infof("Creating and adding new entity: %v", id.String())
	for _, comp := range comps {
		comp.setComponentID(id)
	}

	compsMap := map[reflect.Type]Component{}
	for _, comp := range comps {
		compsMap[reflect.TypeOf(comp)] = comp
	}

	wrld.entityLookup[id] = compsMap

	wrld.checkCompsForNewSystemMatch(compsMap)

	return id
}

// Get Entity gets all the components in the World that
// share a UUID
func (wrld *World) GetEntity(entityID ID) map[reflect.Type]Component {
	if compsMap, ok := wrld.entityLookup[entityID]; !ok {
		return map[reflect.Type]Component{}
	} else {
		// Create the target map
		targetMap := map[reflect.Type]Component{}

		// Copy from the original map to the target map
		for key, value := range compsMap {
			targetMap[key] = value
		}
		return targetMap
	}
}

// RemoveEntity removes all components that share the UUID
func (wrld *World) RemoveEntity(id ID) {
	logrus.Info("removing entity")
	for _, system := range wrld.systems {
		system.removeEntity(id)
	}
}

// AddSystem adds a system to the World after initilizing it
func (wrld *World) AddSystem(system system) {
	logrus.Info("trying to add system")
	system.setRequiredComponents(system.GetRequiredComponents())
	system.calculateBaseComponents()

	system.Initilizer()

	wrld.systems = append(wrld.systems, system)
	logrus.Info("added system")
}

func (wrld *World) checkCompsForNewSystemMatch(comps map[reflect.Type]Component) {
	logrus.Trace("checking for system matches")
	for _, system := range wrld.systems {
		system.addEntity(comps)
	}
}
