package wand

import (
	"context"
	"fmt"
	"strings"
)

type Context interface {
	// All the benefits of a tradition context.Context
	context.Context
	// ParseArgs Does something with inputs provided when calling the app
	ParseArgs([]string)
	// Args returns arguments parsed by your ParseArgs method
	Args() map[string]string
	// Flags returns flags parsed by your ParseArgs method
	Flags() map[string]string
}

// DefaultContext adds simple organizing of arguments and flags
type DefaultContext struct {
	context.Context
	flags map[string]string
	args  map[string]string
}

func NewDefaultContext(ctx context.Context) *DefaultContext {
	return &DefaultContext{
		Context: ctx,
		flags:   make(map[string]string),
		args:    make(map[string]string),
	}
}

func isFlag(i string) bool {
	if len(i) < 2 {
		return false
	}
	if string(i[0]) != string("-") {
		return false
	}
	return true
}

// parseFlag
func parseFlag(flag string) (string, string) {
	if !strings.Contains(flag, "=") {
		return strings.Trim(flag, "-"), ""
	}
	s := strings.Split(flag, "=")
	return strings.Trim(s[0], "-"), s[1]
}

// ParseArgs fullfills the wand.Context interface, along with
// sorting and cleaning up the arguments passed to the called binary
// os.Args
func (ctx *DefaultContext) ParseArgs(inputs []string) {
	if len(inputs) < 1 {
		fmt.Println("no commands given")
		return
	}
	for i, input := range inputs {
		switch {
		case isFlag(input):
			flag, flagArg := parseFlag(input)
			ctx.flags[flag] = flagArg
		case i+1 < len(inputs):
			ctx.args[input] = inputs[i+1]
		default:
			ctx.args[input] = ""
		}
	}
}

// Args fullfills the getter method requirements by wand.Context interface
func (ctx *DefaultContext) Args() map[string]string {
	return ctx.args
}

// Flags fullfills the getter method requirements by wand.Context interface
func (ctx *DefaultContext) Flags() map[string]string {
	return ctx.flags
}
