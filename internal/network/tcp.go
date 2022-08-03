package network

import (
	"encoding/gob"
	"net"

	"github.com/sirupsen/logrus"
)

func handleTCPRequest(enc *gob.Encoder, packet *Packet) {
	logrus.Warn("using default TCP request handler")
}

func StartTCPConnection() error {
	// Connect
	conn, err := net.Dial("tcp", host+tcpPort)
	if err != nil {
		return err
	}

	go tcpHandleConnection(conn)

	return nil
}

func ListenTCP() {
	// Create Listener
	listener, err := net.Listen("tcp", host+tcpPort)
	if err != nil {
		logrus.Error("TCP Bad Listener Start")
		logrus.Error(err)
	}
	defer listener.Close()

	// Start Listener
	logrus.Info("starting TCP Listener")
	for {
		conn, err := listener.Accept()
		if err != nil {
			logrus.Error("TCP Bad Connection Acception")
			logrus.Error(err)
			continue
		}
		logrus.Info("TCP new connection")

		// Handle Every New Connection
		go tcpHandleConnection(conn)
	}
}

func tcpHandleConnection(conn net.Conn) {
	defer conn.Close()
	dec := gob.NewDecoder(conn)
	enc := gob.NewEncoder(conn)

	// Triggers Default Path
	HandleTCPFunc(enc, &Packet{Type: "Default"})

	for {
		// Read
		p, err := readPacket(dec)
		if err != nil {
			logrus.Warn(err)
			break
		}

		// Handle
		go HandleTCPFunc(enc, p)
	}
	logrus.Info("closing connection")
}
