package main

import (
	"encoding/json"
	"io/ioutil"
)

// Manifest is a type used to hold a manifest file in memory.
type Manifest struct {
	buffer 		[]byte
	Namespaces	Namespaces
}

// Namespaces has a list of Namespace that will be unmarshelled from a JSON.
type Namespaces struct {
	List []Namespace `json:"namespaces"`
}

// Namespace is each element in the Namespaces array.
type Namespace struct {
	Name 	string 		`json:"name"`
	Actions []Action 	`json:"actions"`
}

// Action is each action is the Action array in the Namespace model.
type Action struct {
	Name 	string `json:"name"`
	Action 	string `json:"action"`
}

// GetManifest returns a new instance of Manifest struct with the contents of the given
// json file unmarshelled into Namespaces.
func GetManifest(file string) (new Manifest, err error) {
	new.buffer, err = ioutil.ReadFile(file)
	if err == nil {
		err = json.Unmarshal(new.buffer, &new.Namespaces)
	}
	return
}

// FindNamespace searches the manifest array for a namespace with a given name and 
// returns that namespace when found. Returns a NoSuchNamespace error if that namespace
// is not found.
func (m Manifest) FindNamespace(name string) (Namespace, error) {
	for _, namespace := range(m.Namespaces.List) {
		if namespace.Name == name {
			return namespace, nil
		}
	}
	return Namespace{}, NoSuchNamespaceError(name)
}

// FindAction searches through all the elements in Namespace until it finds the action 
// with the given name, and returns the command from the action. Returns a NoSuchAction
// error if the action is not found in the namespace.
func (n Namespace) FindAction(name string) (string, error) {
	for _, action := range(n.Actions) {
		if action.Name == name {
			return action.Action, nil
		}
	}
	return "", NoSuchActionError(name, n.Name)
}

// Add args takes an action string and appends an array of args, so args can be passed 
// from the args zebu recieved
func AddArgs(action string, args []string) (command string) {
	command = action
	for _, arg := range(args) {
		command = command + " " + arg
	}
	return
}