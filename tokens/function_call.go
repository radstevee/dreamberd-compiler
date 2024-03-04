package tokens

import (
	"dreamberd-compiler/common"
	"fmt"
)

func FunctionCall(c *common.FunctionContext) {
	if c.Name == "print" {
		msg := ""
		for i := range c.Arguments {
			msg += fmt.Sprint(c.Arguments[i])
		}
		fmt.Println(msg)
	}
}
