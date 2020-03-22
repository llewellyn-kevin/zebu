package main

import(
	"fmt"
	"os"
	"strings"
)

func main() {
	namespaceName, actionName, args, err := parseInput(os.Args[1:])
	if err != nil {
		panic(err)
	}

	manifest, err := GetManifest("manifest.json")
	if err != nil {
		panic(err)
	}

	namespace, err := manifest.FindNamespace(namespaceName)
	if err != nil {
		panic(err)
	}

	action, err := namespace.FindAction(actionName)
	if err != nil {
		panic(err)
	}

	command := AddArgs(action, args)

	fmt.Println(command)
}

// A struct used to handle all error messages for when the program does not understand the user's
// command input
type IllegalInput struct {
	error string
}

// Error returns a string containing the error message contained in the IllegalInput struct
func (i IllegalInput) Error() string {
	return i.error
}

// emptyInputError returns an instance of the IllegalInput struct with an error message indicating
// that no command has been given.
func emptyInputError() IllegalInput {
	return IllegalInput{
		error: "Error parsing command: no command given. Type 'zebu help' to see a list of commands.",
	}
}

// emptyNamespaceError returns an instance of the IllegalInput struct with an error message indicating
// that no namespace was given before the `:` in the command.
func emptyNamespaceError() IllegalInput {
	return IllegalInput{
		error: "Error parsing command: no namespace provided. If a command has a `:`, please ensure it is preceded by a namespace.",
	}
}

// emptyActionError returns an instance of the IllegalInput struct with an error message indicating
// that no command was given after the `:` in the command.
func emptyActionError() IllegalInput {
	return IllegalInput{
		error: "Error parsing command: no command provided. If a command has a `:`, please ensure it is followed by a command.",
	}
}

// NoSuchNamespaceError returns an instance of the IllegalInput struct with an error message indicating
// that no namespace with the given name was found.
func NoSuchNamespaceError(namespace string) IllegalInput {
	return IllegalInput{
		error: fmt.Sprintf("Error parsing command: no namespace with identifier `%v` found in manifest", namespace),
	}
}

// NoSuchNameActionError returns an instance of the IllegalInput struct with an error message indicating
// that no action with the given name was found in the given namespace.
func NoSuchActionError(action, namespace string) IllegalInput {
	return IllegalInput{
		error: fmt.Sprintf("Error parsing command: no action with identifier `%v` found in namespace `%v`", action, namespace),
	}
}

// parseInput takes a slice of command line arguments that can be retrieved from os.Args. Note that
// this slice is not meant to include os.Args[0] which is the command used to execute the code, so
// this function should usually be called by: "parseInput[1:]". 
//
// parseInput expects the user to enter their command in the format: "[namespace]:[action] [args...]"
// where namespace indicates where zebu should look for the action, action specifies what zebu should
// do, and args is a list of arguments to provide any additional information. args should then be 
// passed into the parseArgs function so all the arguments are divided by arguments with flags and 
// arguments without flags. 
func parseInput(input []string) (namespace string, action string, args []string, err error) {
	// Check that the input has at list one field
	if len(input) == 0 {
		err = emptyInputError()
		return 
	}

	// Parse first field into namespace and action
	commandArr := strings.Split(input[0], ":")
	if len(commandArr) > 1 {
		namespace, action = commandArr[0], commandArr[1]
	} else {
		namespace, action = commandArr[0], "default"
	}

	// Check that if colon is present, both a namespace and action are present
	if namespace == "" {
		err = emptyNamespaceError()
	} else if action == "" {
		err = emptyActionError()
	}

	// Put the remainder of the fields into args
	args = input[1:]

	return
}

// parseArgs takes a slice of arguments and iterates through them to seperate them into 
// a two seperate data structures. Any argument that begins with a '-' will be treated
// as a flag, and add this argument to the flagArgs map with a key for whatever follows 
// the hyphen and value being the next argument in the slice. Any other argument is 
// appended to the orderedArgs slice.
func ParseArgs(args []string) (orderedArgs []string, flagArgs map[string]string) {
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