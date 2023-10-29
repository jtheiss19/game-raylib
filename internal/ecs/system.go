package ecs

import (
	"reflect"

	"github.com/sirupsen/logrus"
)

type system interface {
	// Update is used to update the system. dt is
	// the time since the last update, although this
	// value is not enforced by the library
	Update(dt float32)

	// Initilizer is ran at startup to init the system. Runs after
	// polling of initial components in world.
	Initilizer()

	// GetRequiredComponents returns a struct of arrays of structs
	// to create a definition for ecs to use to add and
	// remove entities to your system. See the example system
	// for how to use and querry data.
	GetRequiredComponents() interface{}

	// all of the rest are internal functions NOT to be called
	// by the user.
	setRequiredComponents(i interface{})
	getBaseComponents() [][]reflect.Type
	calculateBaseComponents() [][]reflect.Type
	addEntity(comps map[reflect.Type]Component)
	removeEntity(id ID)
}

// BaseSystem is a default struct which provide functionality
// and default behavior to structs to make them match the system
// interface for use in the ecs.
type BaseSystem struct {
	TrackedEntities      interface{}
	baseComponentsType   [][]reflect.Type
	baseComponentsStruct interface{}
	foundIDs             map[ID][]int
}

// User Overrides
func (bs *BaseSystem) Update(dt float32)                  {}
func (bs *BaseSystem) Initilizer()                        {}
func (bs *BaseSystem) GetRequiredComponents() interface{} { return nil }

// Boiler Plate
func (bs *BaseSystem) setRequiredComponents(i interface{}) {
	bs.baseComponentsStruct = i
}

func (bs *BaseSystem) getBaseComponents() [][]reflect.Type {
	return bs.baseComponentsType
}

func (bs *BaseSystem) calculateBaseComponents() [][]reflect.Type {
	bs.foundIDs = map[ID][]int{}
	if bs.baseComponentsStruct == nil {
		logrus.Trace("can not calculate base componets, type is <nil>")
		return nil
	}

	reqComponentsStruct := reflect.ValueOf(bs.baseComponentsStruct).Elem()

	returnTypes := [][]reflect.Type{}
	componentInterface := reflect.TypeOf((*Component)(nil)).Elem()

	for requiredCompTypes := 0; requiredCompTypes < reqComponentsStruct.NumField(); requiredCompTypes++ {
		returnTypesElem := []reflect.Type{}
		currentType := reqComponentsStruct.Field(requiredCompTypes)

		for reqCompPos := 0; reqCompPos < currentType.Index(0).Elem().NumField(); reqCompPos++ {
			reqCompType := currentType.Index(0).Type().Elem().Field(reqCompPos).Type

			if reqCompType.Implements(componentInterface) {
				returnTypesElem = append(returnTypesElem, reqCompType)
			} else {
				logrus.Error("listed system required component does not implement component interface")
			}
		}
		returnTypes = append(returnTypes, returnTypesElem)
	}

	bs.baseComponentsType = returnTypes
	logrus.Trace("created required components struct for system: ", returnTypes)

	trackingStruct := reflect.New(reflect.TypeOf(bs.baseComponentsStruct).Elem())
	bs.TrackedEntities = trackingStruct.Interface()

	return returnTypes
}

func (bs *BaseSystem) addEntity(comps map[reflect.Type]Component) {
	logrus.Trace("trying to add entity to system")

	for reqStructTypeIndex, reqCompStruct := range bs.getBaseComponents() {
		// Test if already exists
		id := ID("")
		for _, comp := range comps {
			id, _ = comp.GetComponentID()
			break
		}
		if array, ok := bs.foundIDs[id]; ok {
			found := false
			for _, index := range array {
				if index == reqStructTypeIndex {
					found = true
					break
				}
			}
			if found { // Found Entity Already Exists
				continue
			}
		}

		matchedComps := []Component{}
		for _, reqComp := range reqCompStruct {
			if matchedComp, ok := comps[reqComp]; !ok {
				logrus.Trace("comp does not contain all requied comps")
				break
			} else {
				matchedComps = append(matchedComps, matchedComp)
			}
		}

		// If we found a match for each requirement
		if len(matchedComps) == len(reqCompStruct) {
			logrus.Trace("found system match with entity")

			outerStruct := reflect.ValueOf(bs.TrackedEntities).Elem() // [][][]type All Tracked Requirements
			innerStruct := outerStruct.Field(reqStructTypeIndex)      // [][]type A type of tracked Comps
			reqFieldElem := innerStruct

			reqFieldType := innerStruct.Type().Elem()
			newReqFieldEntry := reflect.New(reqFieldType.Elem())
			Fill(newReqFieldEntry, comps)

			reqFieldElem.Set(reflect.Append(reqFieldElem, newReqFieldEntry))

			// add to found list
			id, _ := matchedComps[0].GetComponentID()
			if array, ok := bs.foundIDs[id]; ok {
				bs.foundIDs[id] = append(array, reqStructTypeIndex)
			} else {
				bs.foundIDs[id] = []int{reqStructTypeIndex}
			}

			logrus.Tracef("added entity of to system for requirement: %v with id: %v", reqFieldType, id)
		}
	}
}

func (bs *BaseSystem) removeEntity(id ID) {
	logrus.Trace("removing entity from system")
}
