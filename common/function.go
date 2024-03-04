package common

import "fmt"

type Function struct {
	Name      string
	Arguments []string
	Execute   func(*FunctionContext)
}

var (
	Print Function = Function{
		Name:      "print",
		Arguments: []string{"msg"},
		Execute: func(c *FunctionContext) {
			msg := ""
			for i := range c.Arguments {
				msg += fmt.Sprint(c.Arguments[i])
			}
			fmt.Println(msg)
		},
	}

	AllFunctions = []Function{
		Print,
	}
)
