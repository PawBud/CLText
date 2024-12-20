package main

import (
	"log"
	"net"
)

func main() {
	s := newServer()
	go s.run()


	// assign the port to a random available port instead of using a static port
	listener, err := net.Listen("tcp", ":0")
	addr := listener.Addr().String()
	if err != nil {
		log.Fatalf("unable to start server: %s", err.Error())
	}

	defer listener.Close()
	log.Printf(" ===================== Welcome to CLText =====================\n")
	log.Printf("This is the admin server")
	log.Printf("Started a server on :%s", addr)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("unable to accept connection: %s", err.Error())
			continue
		}

		go s.newClient(conn)
	}
}
