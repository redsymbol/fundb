package main

import (
	"fundb/client"
	"fundb/shell"
	"fundb/constants"
)

func main() {
	cl := client.Client{Port:constants.DEFAULT_PORT}
	shell.Start(cl)
}