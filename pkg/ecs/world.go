package ecs

import (
	"reflect"
	"sync"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

var (
	mu          sync.Mutex
	activeWorld *World
)

type World struct {
	systems         []system
	entityLookup    map[ID]map[reflect.Type]Component
	StartUpdateFunc func()
	EndUpdateFunc   func()
}

// NewWorld generates a new empty World and sets
// it as the active World
func NewWorld() *World {
	activeWorld = &World{
		systems:      []system{},
		entityLookup: map[ID]map[reflect.Type]Component{},
		StartUpdateFunc: func() {
		},
		EndUpdateFunc: func() {
		},
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
	wrld.StartUpdateFunc()
	for _, system := range wrld.systems {
		system.Update(dt)
	}
	wrld.EndUpdateFunc()
}

// AddComponent adds a new component to the World.
func (wrld *World) AddComponent(comp Component) {
	logrus.Trace("adding component")
	id, err := comp.GetComponentID()
	if err != nil {
		id = comp.generateComponentID()
	}

	mu.Lock()
	if comps, ok := wrld.entityLookup[id]; ok {
		if foundComp, ok := comps[reflect.TypeOf(comp)]; ok {
			v := reflect.ValueOf(foundComp).Elem()
			v.Set(reflect.ValueOf(comp).Elem())
		} else {
			comps[reflect.TypeOf(comp)] = comp
		}

		wrld.entityLookup[id] = comps
		wrld.checkCompsForNewSystemMatch(comps)
	} else {
		compsMap := map[reflect.Type]Component{}
		compsMap[reflect.TypeOf(comp)] = comp
		wrld.entityLookup[id] = compsMap
		wrld.checkCompsForNewSystemMatch(compsMap)
	}
	mu.Unlock()
}

// RemoveCompoent removes a compoent from the World
func (wrld *World) RemoveCompoent(comp Component) {
	logrus.Trace("removing component")
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
	logrus.Tracef("Creating and adding new entity: %v", id.String())
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
	mu.Lock()
	defer mu.Unlock()
	if compsMap, ok := wrld.entityLookup[entityID]; !ok {
		return map[reflect.Type]Component{}
	} else {
		return compsMap
	}
}

// GetComponents gets all the components in the World that
// match your input type
func (wrld *World) GetComponents(lookupType reflect.Type) []Component {
	mu.Lock()
	defer mu.Unlock()
	returnList := []Component{}
	for _, entities := range wrld.entityLookup {
		if comp, ok := entities[lookupType]; ok {
			returnList = append(returnList, comp)
		}
	}
	return returnList
}

// RemoveEntity removes all components that share the UUID
func (wrld *World) RemoveEntity(id ID) {
	logrus.Trace("removing entity")
	for _, system := range wrld.systems {
		system.removeEntity(id)
	}
}

// AddSystem adds a system to the World after initilizing it
func (wrld *World) AddSystem(system system) {
	logrus.Trace("trying to add system")
	system.setRequiredComponents(system.GetRequiredComponents())
	system.calculateBaseComponents()

	system.Initilizer()

	wrld.systems = append(wrld.systems, system)
	logrus.Trace("added system")
}

func (wrld *World) checkCompsForNewSystemMatch(comps map[reflect.Type]Component) {
	logrus.Trace("checking for system matches")
	for _, system := range wrld.systems {
		system.addEntity(comps)
	}
}
