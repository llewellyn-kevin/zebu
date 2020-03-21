package main

import(
	"fmt"
	"os"
	"strings"
)

func main() {
	namespace, action, args, _ := parseInput(os.Args[1:])
	orderedArgs, flagArgs := parseArgs(args)

	fmt.Println("NAMESPACE: " + namespace)
	fmt.Println("ACTION: " + action)
	fmt.Println(fmt.Sprintf("ARGS: %v", orderedArgs))
	for k, v := range(flagArgs) {
		fmt.Println(fmt.Sprintf(" -%v: %v", k, v))
	}
}

type IllegalInput struct {
	error string
}

func (i* IllegalInput) Error() string {
	return i.error
}

func newIllegalInput() IllegalInput {
	return IllegalInput{
		error: "IllegalInput Detected",
	}
}

func parseInput(input []string) (namespace string, action string, args []string, err error) {
	commandArr := strings.Split(input[0], ":")
	if len(commandArr) > 1 {
		namespace, action = commandArr[0], commandArr[1]
	} else {
		namespace, action = commandArr[1], "default"
	}
	args = input[1:]
	return
}

// parseArgs takes a slice of arguments and iterates through them to seperate them into 
// a two seperate data structures. Any argument that begins with a '-' will be treated
// as a flag, and add this argument to the flagArgs map with a key for whatever follows 
// the hyphen and value being the next argument in the slice. Any other argument is 
// appended to the orderedArgs slice.
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