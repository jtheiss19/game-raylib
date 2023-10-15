package multiplayer

import (
	"encoding/gob"
	"reflect"

	"github.com/jtheiss19/game-raylib/internal/ecs"
	"github.com/jtheiss19/game-raylib/internal/engine/components"
	components2d "github.com/jtheiss19/game-raylib/internal/engine/components/2d"
	"github.com/jtheiss19/game-raylib/internal/network"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type NetworkingSystem struct {
	*ecs.BaseSystem
	isServer        bool
	playerID        ecs.ID
	timeSinceUpdate float32
	connections     map[ecs.ID]*gob.Encoder
}

func NewNetworkingSystem(isServer bool) *NetworkingSystem {
	return &NetworkingSystem{
		BaseSystem:      &ecs.BaseSystem{},
		isServer:        isServer,
		playerID:        "",
		timeSinceUpdate: 0,
		connections:     map[ecs.ID]*gob.Encoder{},
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
	if !ts.isServer {
		ts.updateClient(dt)
	} else {
		ts.updateServer(dt)
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
	network.RegisterType(ecs.BaseComponent{})
	network.RegisterType(components2d.CollisionComponent{})
	network.RegisterType(ecs.ID(uuid.Nil.String()))
}

type MessageType string

const (
	JOIN_PACKET    MessageType = "join" // Data should be player name
	COMPONENT_DATA MessageType = "new comp"
	PLAYER_DATA    MessageType = "player data"
	WORLD_UPDATE   MessageType = "request for world data"
)

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
