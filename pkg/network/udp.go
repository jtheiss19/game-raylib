package network

import (
	"bytes"
	"encoding/gob"
	"net"

	"github.com/sirupsen/logrus"
)

func handleUDPRequest(conn net.Conn, packet *Packet) {
	logrus.Warn("using default UDP request handler")
}

func ListenUDP() {
	dst, err := net.ResolveUDPAddr("udp", host+udpPort)
	if err != nil {
		logrus.Error("UDP Bad creation")
		logrus.Error(err)
	}
	conn, err := net.ListenUDP("udp", dst)
	if err != nil {
		logrus.Error("UDP Bad Connection")
		logrus.Error(err)
	}
	defer conn.Close()

	logrus.Info("starting UDP Listener")
	for {
		p, err := ReadPacketUDP(conn)
		if err != nil {
			logrus.Error(err)
			continue
		}

		// Handle
		logrus.Debug("handling udp")
		go HandleUDPFunc(conn, p)
	}
}

func SendUDP(data interface{}, typeOfRequest string) error {
	// Conenct
	dst, err := net.ResolveUDPAddr("udp", host+udpPort)
	if err != nil {
		return err
	}
	conn, err := net.ListenPacket("udp", ":0")
	if err != nil {
		return err
	}
	defer conn.Close()

	SendPacketUDP(conn, dst, data, typeOfRequest)

	return nil
}

func ReadPacketUDP(conn *net.UDPConn) (*Packet, error) {
	// Read
	buf := make([]byte, 4096)
	n, _, err := conn.ReadFromUDP(buf[:])
	logrus.Debug("just read from UDP connection")
	if err != nil {
		logrus.Error("UDP Bad Connection Read")
		logrus.Error(err)
		return nil, err
	}

	// Decode
	logrus.Debug("going to decode buffer from UDP read")
	dec := gob.NewDecoder(bytes.NewReader(buf[:n]))
	p, err := readPacket(dec)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	return p, nil
}

func SendPacketUDP(conn net.PacketConn, dst *net.UDPAddr, data interface{}, typeOfRequest string) error {
	// Create Packet
	packet := &Packet{
		Type: typeOfRequest,
		Data: data,
	}

	// Encode
	var buf bytes.Buffer
	encoder := gob.NewEncoder(&buf)
	err := encoder.Encode(packet)
	if err != nil {
		return err
	}

	// Send
	_, err = conn.WriteTo(buf.Bytes(), dst)
	if err != nil {
		return err
	}

	return nil
}
