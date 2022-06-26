package network

import (
	"bytes"
	"encoding/gob"
	"errors"
	"net"

	"github.com/sirupsen/logrus"
)

type Packet struct {
	Type string
	Data interface{}
}

var (
	host    = "localhost:"
	tcpPort = "8081"
	udpPort = "8080"

	HandleUDPFunc = handleUDPRequest
	HandleTCPFunc = handleTCPRequest
)

func init() {
	RegisterType(Packet{})
}

func RegisterType(thingToRegister interface{}) {
	gob.Register(thingToRegister)
}

func CreatePacket(dataType string, data interface{}) *Packet {
	return &Packet{
		Type: dataType,
		Data: data,
	}
}

func readPacket(dec *gob.Decoder) (*Packet, error) {
	// Decode
	pack := &Packet{}
	err := dec.Decode(&pack)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	// Verify
	err = pack.verifyPacket()
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	return pack, nil
}

func (pack *Packet) verifyPacket() error {
	if pack.Type == "" {
		return errors.New("no message type")
	}

	return nil
}

func (pack *Packet) SendPacket(conn net.Conn) error {
	// Encode
	var buf bytes.Buffer
	encoder := gob.NewEncoder(&buf)
	err := encoder.Encode(pack)
	if err != nil {
		return err
	}

	// Send
	_, err = conn.Write(buf.Bytes())
	if err != nil {
		return err
	}

	return nil
}
