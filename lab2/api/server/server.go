package server

import (
	"CN_lab2/lab2/pkg"
	"fmt"
	"net"
)

func RunServer() {
	ip, port := pkg.CheckConnection()
	listener, err := net.Listen("tcp", ip+":"+port)
	if err != nil {
		_ = fmt.Errorf("runServer: error on listening %s:%s : %w", ip, port, err)
		return
	}
	defer listener.Close()

	fmt.Printf("runServer: listening on %s:%s\n", ip, port)

	for {
		conn, err := listener.Accept()
		if err != nil {
			_ = fmt.Errorf("runServer: error on accepting connection: %w", err)
			return
		}
		go handleConnection(conn)
	}

}

func handleConnection(conn net.Conn) {
	input := make([]byte, 4096)

}
