package shell

import (
	"os"
	"fmt"
	"bufio"
	"io"
	"net"
)

const PROMPT = "> "

// Starts a command-line shell
func Start(host string, port int) {
	stdin := bufio.NewReader(os.Stdin)
	for {
		conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", host, port))
		if err != nil {
			panic(err)
		}
		fmt.Print(PROMPT)
		line, err := stdin.ReadString('\n')
		if io.EOF == err {
			os.Exit(0)
		} else if err != nil {
			panic(err)
		}
		_, err = conn.Write([]byte(line))
		if err != nil {
			panic(err)
		}
		err = conn.Close()
		if err != nil {
			panic(err)
		}
	}
}