package server

import (
	"net"
	"fmt"
	"io/ioutil"
)

func handleConnection(conn net.Conn) {
	bytes, err := ioutil.ReadAll(conn)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}

func Start(port int) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		panic(err)
	}
	fmt.Println("Listening on port: ", port)
	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		go handleConnection(conn)
	}
}