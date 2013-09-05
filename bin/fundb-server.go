package main

import "fundb/server"

const PORT = 8042

func main() {
	server.Start(PORT)
}