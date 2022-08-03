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
	isServer        bool
	playerID        ecs.ID
	timeSinceUpdate float32
	connections     []*gob.Encoder
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
	Player  []*RequirePlayers
}

type RequirePlayers struct {
	Player *components.PlayerComponent
}

type RequireNetworking struct {
	Network *components.NetworkComponent
}

func (ts *NetworkingSystem) GetRequiredComponents() interface{} {
	return &RequiredNetworkingSystemComps{
		Network: []*RequireNetworking{{Network: &components.NetworkComponent{}}},
		Player:  []*RequirePlayers{{Player: &components.PlayerComponent{}}},
	}
}

// Functionality
func (ts *NetworkingSystem) Update(dt float32) {
	networkedEntities, ok := ts.TrackedEntities.(*RequiredNetworkingSystemComps)
	if !ok {
		logrus.Error("could not update system, bad tracked entities")
		return
	}

	// Contstantly push entities the player owns via the player component.
	if !ts.isServer {
		for _, reqComp := range networkedEntities.Player {
			if reqComp.Player.PlayerID == ts.playerID {
				id, _ := reqComp.Player.GetComponentID()
				entity := ecs.GetActiveWorld().GetEntity(id)
				for _, comp := range entity {
					packet := network.CreatePacket(string(COMPONENT_DATA), comp)
					err := ts.connections[0].Encode(packet)
					if err != nil {
						logrus.Error(err)
					}
				}
			}
		}
	}

	// Periodically have the client request world updates
	ts.timeSinceUpdate += dt
	if ts.timeSinceUpdate > 5000 && !ts.isServer && len(ts.connections) > 0 {
		ts.timeSinceUpdate = 0
		logrus.Info("Requesting World State")
		packet := network.CreatePacket(string(WORLD_UPDATE), "request regular world update")
		err := ts.connections[0].Encode(packet)
		if err != nil {
			logrus.Error(err)
		}
	}

}

func (ts *NetworkingSystem) Initilizer() {
	if ts.isServer {
		network.HandleTCPFunc = ts.serverTCPhandler
		go network.ListenTCP()
	} else {
		network.HandleTCPFunc = ts.clientTCPhandler
		err := network.StartTCPConnection()
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
	WORLD_UPDATE   MessageType = "request for world data"
)

// Server Handler
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

		// Add connection to pool
		ts.connections = append(ts.connections, enc)

		comps := objects.NewPlayer(newPlayersID)
		ecs.GetActiveWorld().AddEntity(comps)

	case WORLD_UPDATE:
		logrus.Info("Recieved World Update Request")
		networkedEntities, ok := ts.TrackedEntities.(*RequiredNetworkingSystemComps)
		if !ok {
			logrus.Error("could not update system, bad tracked entities")
			return
		}

		for _, entity := range networkedEntities.Network {
			compID, _ := entity.Network.GetComponentID()
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
	case COMPONENT_DATA:
		newComp := packet.Data
		logrus.Debugf("I got a new comp: %v", newComp)
		newComponent, ok := newComp.(ecs.Component)
		if !ok {
			break
		}
		switch newComponent := newComponent.(type) {
		default:
			createComponent(newComponent)
		}
	}
}

// Client Handler
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
		default:
			createComponent(newComponent)
		}

	case PLAYER_DATA:
		// Add connection to pool
		ts.connections = append(ts.connections, enc)

		id := packet.Data
		logrus.Infof("setting playerID: %v", id)
		if convertedID, ok := id.(ecs.ID); ok {
			ts.playerID = convertedID
		}

		packet := network.CreatePacket(string(WORLD_UPDATE), "request regular world update")
		err := enc.Encode(packet)
		if err != nil {
			logrus.Error(err)
		}

	default:
		packet := network.CreatePacket(string(JOIN_PACKET), "Joseph")
		err := enc.Encode(packet)
		if err != nil {
			logrus.Error(err)
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
