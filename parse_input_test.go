package main

import (
	"strings"
	"testing"
)

type ParseInputResult struct {
	namespace	string
	action		string
	args		[]string
	err			error
}

func TestParseInput(t *testing.T) {
	cases := []struct {
		input			[]string
		expectedOutput	ParseInputResult
	}{
		{ // test no input
			nil, 
			ParseInputResult{"", "", nil, emptyInputError()},
		},
		{ // test empty namespace
			makeArray(":seed"), 
			ParseInputResult{"", "seed", nil, emptyNamespaceError()},
		},
		{ // test empty command
			makeArray("db: FakeUsers"),
			ParseInputResult{"db", "", []string{"FakeUsers"}, emptyActionError()},
		},
		{ // test valid namespace only
			makeArray("init"),
			ParseInputResult{"init", "default", nil, nil},
		}, 
		{ // test valid namespace and args only
			makeArray("init some args"),
			ParseInputResult{"init", "default", []string{"some", "args"}, nil},
		},
		{ // test full valid command
			makeArray("make:controller User -c crud"),
			ParseInputResult{"make", "controller", []string{"User", "-c", "crud"}, nil},
		},
	}

	for _, c := range cases {
		namespace, action, args, err := parseInput(c.input)

		if namespace != c.expectedOutput.namespace {
			t.Errorf(
				"incorrect output for namespace given input `%v`: expected `%v` but got `%v`", 
				c.input, 
				c.expectedOutput.namespace, 
				namespace,
			)
		} else if action != c.expectedOutput.action {
			t.Errorf(
				"incorrect output for action given input `%v`: expected `%v` but got `%v`", 
				c.input, 
				c.expectedOutput.action, 
				action,
			)
		} else if !isSliceEqual(args, c.expectedOutput.args) {
			t.Errorf(
				"incorrect output for args given input `%v`: expected `%v` but got `%v`", 
				c.input, 
				c.expectedOutput.args, 
				args,
			)
		} else if err != c.expectedOutput.err {
			t.Errorf(
				"incorrect output for err given input `%v`: expected `%v` but got `%v`", 
				c.input, 
				c.expectedOutput.err, 
				err,
			)
		}
	}
}

// isSliceEqual tells whether a and b contain the same elements.
// A nil argument is equivalent to an empty slice.
func isSliceEqual(a, b []string) bool {
    if len(a) != len(b) {
        return false
    }
    for i, v := range a {
        if v != b[i] {
            return false
        }
    }
    return true
}

// makeArray translates a string that represents user input and converts it into a slice
func makeArray(s string) []string {
	return strings.Split(s, " ")
}