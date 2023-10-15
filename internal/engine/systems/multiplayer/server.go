package multiplayer

import (
	"encoding/gob"

	"github.com/jtheiss19/game-raylib/internal/ecs"
	objects2d "github.com/jtheiss19/game-raylib/internal/engine/objects/2d"
	"github.com/jtheiss19/game-raylib/internal/network"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

func (ts *NetworkingSystem) updateServer(dt float32) {

	for playerID, enc := range ts.connections {
		networkedEntities, ok := ts.TrackedEntities.(*RequiredNetworkingSystemComps)
		if !ok {
			logrus.Error("could not update system, bad tracked entities")
			return
		}

	out:
		for _, entity := range networkedEntities.Network {
			compID, _ := entity.Network.GetComponentID()
			if playerID == compID {
				continue
			}
			entity := ecs.GetActiveWorld().GetEntity(compID)
			for _, comp := range entity {
				logrus.Debugf("sending Data: %v", comp)

				packet := network.CreatePacket(string(COMPONENT_DATA), comp)
				err := enc.Encode(packet)
				if err != nil {
					logrus.Error(err)
					delete(ts.connections, playerID)
					break out
				}
			}
		}
	}

}

// Server Handler
func (ts *NetworkingSystem) serverTCPhandler(enc *gob.Encoder, packet *network.Packet) {

	switch MessageType(packet.Type) {
	case JOIN_PACKET:
		logrus.Infof("%s is joining the game", packet.Data)
		logrus.Infof("sending %s basic starting data", packet.Data)

		// Add connection to pool
		serverJoinHandler(enc, ts)

	case WORLD_UPDATE:
		serverWorldDataHandler(ts, enc)

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

func serverWorldDataHandler(ts *NetworkingSystem, enc *gob.Encoder) {
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

			logrus.Debugf("sending Data: %v", comp)

			packet := network.CreatePacket(string(COMPONENT_DATA), comp)
			err := enc.Encode(packet)
			if err != nil {
				logrus.Error(err)
			}
		}
	}
}

func serverJoinHandler(enc *gob.Encoder, ts *NetworkingSystem) {
	newPlayersID := ecs.ID(uuid.New().String())
	playerDataPacket := network.CreatePacket(string(PLAYER_DATA), newPlayersID)
	err := enc.Encode(playerDataPacket)
	if err != nil {
		logrus.Error(err)
	}

	comps := objects2d.New2DPlayer(newPlayersID)
	playerID := ecs.GetActiveWorld().AddEntity(comps)

	ts.connections[playerID] = enc
}
