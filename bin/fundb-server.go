package main

import (
	"fundb/server"
	"fundb/constants"
)

func main() {
	server.Start(constants.DEFAULT_PORT)
}