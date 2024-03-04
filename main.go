package main

import (
	"dreamberd-compiler/ast"
	"dreamberd-compiler/common"
	"fmt"
)

func main() {
	// read the file
	file, err := common.ReadFile("test/hello_world.dreamberd")
	if err != nil {
		panic(err)
	}

	// tokenize the file
	tokens := ast.Tokenize(file)
	for i := range tokens {
		fmt.Println(tokens[i].ToString())
	}
}
