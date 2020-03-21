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

	orderedArgs, flagArgs := parseArgs(args)

	fmt.Println("NAMESPACE: " + namespace)
	fmt.Println("ACTION: " + action)
	fmt.Println(fmt.Sprintf("ARGS: %v", orderedArgs))
	for k, v := range(flagArgs) {
		fmt.Println(fmt.Sprintf(" -%v: %v", k, v))
	}
}

func parseArgs(args []string) (orderedArgs []string, flagArgs map[string]string) {
	flagArgs = make(map[string]string)

	for i := 0; i < len(args); i++ {
		if strings.Index(args[i], "-") == 0 {
			flagArgs[args[i]] = args[i+1]
			i = i + 1
		} else {
			orderedArgs = append(orderedArgs, args[i])
		}
	}

	return
}