package systems

import (
	"encoding/gob"
	"fmt"
	"reflect"
	"rouge/internal/ecs"
	"rouge/internal/engine/components"
	"rouge/internal/engine/objects"
	"rouge/internal/network"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type NetworkingSystem struct {
	*ecs.BaseSystem
	isServer bool
	playerID ecs.ID
}

func NewNetworkingSystem(isServer bool) *NetworkingSystem {
	return &NetworkingSystem{
		BaseSystem: &ecs.BaseSystem{},
		isServer:   isServer,
	}
}

// Comps
type RequiredNetworkingSystemComps struct {
	Network []*RequireNetworking
}

type RequireNetworking struct {
	Network *components.NetworkComponent
}

func (ts *NetworkingSystem) GetRequiredComponents() interface{} {
	return &RequiredNetworkingSystemComps{
		Network: []*RequireNetworking{{
			Network: &components.NetworkComponent{},
		}},
	}
}

// Functionality
func (ts *NetworkingSystem) Update(dt float32) {
	_, ok := ts.TrackedEntities.(*RequiredNetworkingSystemComps)
	if !ok {
		logrus.Error("could not update system, bad tracked entities")
		return
	}
}

func (ts *NetworkingSystem) Initilizer() {
	if ts.isServer {
		network.HandleTCPFunc = ts.serverTCPhandler
		go network.ListenTCP()
	} else {
		network.HandleTCPFunc = ts.clientTCPhandler
		conn, err := network.StartTCPConnection()
		if err != nil {
			logrus.Fatal(err)
		}
		err = network.CreatePacket(string(JOIN_PACKET), "Joseph").SendPacket(conn)
		if err != nil {
			logrus.Fatal(err)
		}
	}
}

func init() {
	network.RegisterType(components.CameraComponent{})
	network.RegisterType(components.InputComponent{})
	network.RegisterType(components.TransformationComponent{})
	network.RegisterType(components.NetworkComponent{})
	network.RegisterType(components.ModelComponent{})
	network.RegisterType(components.PlayerComponent{})
	network.RegisterType(ecs.BaseComponent{})
	network.RegisterType(ecs.ID(uuid.Nil.String()))
}

type MessageType string

const (
	JOIN_PACKET    MessageType = "join" // Data should be player name
	COMPONENT_DATA MessageType = "new comp"
	PLAYER_DATA    MessageType = "player data"
)

func (ts *NetworkingSystem) serverTCPhandler(enc *gob.Encoder, packet *network.Packet) {

	switch MessageType(packet.Type) {
	case JOIN_PACKET:
		logrus.Infof("%s is joining the game", packet.Data)
		logrus.Infof("sending %s basic starting data", packet.Data)

		newPlayersID := ecs.ID(uuid.New().String())
		playerDataPacket := network.CreatePacket(string(PLAYER_DATA), newPlayersID)
		err := enc.Encode(playerDataPacket)
		if err != nil {
			logrus.Error(err)
		}

		comps := objects.NewPlayer(newPlayersID)
		ecs.GetActiveWorld().AddEntity(comps)

		entities, ok := ts.TrackedEntities.(*RequiredNetworkingSystemComps)
		if !ok {
			logrus.Error("could not update system, bad tracked entities")
			return
		}
		for _, comp := range entities.Network {
			compID := comp.Network.EntityID
			entity := ecs.GetActiveWorld().GetEntity(compID)
			for _, comp := range entity {
				if testComp, ok := comp.(*components.TransformationComponent); ok {
					fmt.Println(testComp)
				}
				logrus.Debugf("sending Data: %v", comp)

				packet := network.CreatePacket(string(COMPONENT_DATA), comp)
				err := enc.Encode(packet)
				if err != nil {
					logrus.Error(err)
				}
			}
		}

	}
}

func (ts *NetworkingSystem) clientTCPhandler(enc *gob.Encoder, packet *network.Packet) {
	switch MessageType(packet.Type) {
	case COMPONENT_DATA:
		newComp := packet.Data
		logrus.Debugf("I got a new comp: %v", newComp)
		newComponent, ok := newComp.(ecs.Component)
		if !ok {
			break
		}
		switch newComponent := newComponent.(type) {
		case components.PlayerComponent:
			if newComponent.PlayerID != ts.playerID {
				break
			}
			ecs.GetActiveWorld().AddComponent(&newComponent)
		case components.TransformationComponent:
			fmt.Println(newComponent)
			ecs.GetActiveWorld().AddComponent(&newComponent)
		default:
			createComponent(newComponent)
		}

	case PLAYER_DATA:
		id := packet.Data
		logrus.Infof("setting playerID: %v", id)
		if convertedID, ok := id.(ecs.ID); ok {
			ts.playerID = convertedID
		}
	}
}

func createComponent(newComponent ecs.Component) {
	val := reflect.ValueOf(newComponent)
	vp := reflect.New(val.Type())
	vp.Elem().Set(val)
	newComponent = vp.Interface().(ecs.Component)
	logrus.Tracef("Data After Handling: %v", newComponent)
	ecs.GetActiveWorld().AddComponent(newComponent)
	id, _ := newComponent.GetComponentID()
	logrus.Debugf("adding %v component id: %v", val.Type(), id.String())
}
