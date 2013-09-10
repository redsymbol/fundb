package main

import (
	"fundb/shell"
	"fundb/constants"
)

func main() {
	shell.Start("127.0.0.1", constants.DEFAULT_PORT)
}