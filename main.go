package main

import(
	"fmt"
	"os"
	"strings"
)

func main() {
	var(
		namespace, action string 
		args []string
	)

	command := os.Args[1]
	args = os.Args[2:]

	commandArr := strings.Split(command, ":")
	if len(commandArr) > 1 {
		namespace, action = commandArr[0], commandArr[1]
	} else {
		namespace, action = commandArr[1], commandArr[1]
	}

	fmt.Println("NAMESPACE: " + namespace)
	fmt.Println("ACTION: " + action)
	fmt.Println(fmt.Sprintf("ARGS: %v", args))
}