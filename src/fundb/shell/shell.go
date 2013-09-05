package shell

import (
	"os"
	"fmt"
	"bufio"
)

const PROMPT = "> "

// Starts a command-line shell
func Start() {
	stdin := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(PROMPT)
		line, err := stdin.ReadString('\n')
		if err != nil {
			panic(err)
		}
		fmt.Println("you said: ", line)
	}
}