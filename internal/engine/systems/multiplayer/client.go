package multiplayer

import (
	"encoding/gob"
	"rouge/internal/ecs"
	"rouge/internal/engine/components"
	"rouge/internal/network"

	"github.com/sirupsen/logrus"
)

func (ts *NetworkingSystem) updateClient(dt float32) {

	networkedEntities, ok := ts.TrackedEntities.(*RequiredNetworkingSystemComps)
	if !ok {
		logrus.Error("could not update system, bad tracked entities")
		return
	}

	// Contstantly push entities the player owns via the player component.
	for _, reqComp := range networkedEntities.Player {
		if reqComp.Player.PlayerID == ts.playerID {
			id, _ := reqComp.Player.GetComponentID()
			entity := ecs.GetActiveWorld().GetEntity(id)
			for _, comp := range entity {
				packet := network.CreatePacket(string(COMPONENT_DATA), comp)
				err := ts.connections[ts.playerID].Encode(packet)
				if err != nil {
					logrus.Error(err)
				}
			}
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

		id := packet.Data
		logrus.Infof("setting playerID: %v", id)
		if convertedID, ok := id.(ecs.ID); ok {
			ts.playerID = convertedID
			ts.connections[convertedID] = enc
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
