package common

import "fmt"

type Token struct {
	Line   int
	Column int
	Type   TokenType
}

func (t Token) ToString() string {
	return fmt.Sprintf("Line: %d, Column: %d, Type: %s, Keyword: %s", t.Line, t.Column, t.Type.Name, t.Type.Keyword)
}

type TokenType struct {
	Name    string
	Keyword string
}

var (
	DeclarationVar = TokenType{
		Name:    "DeclarationVar",
		Keyword: "var",
	}
	DeclarationFunc = TokenType{
		Name:    "DeclarationFunc",
		Keyword: "function",
	}
	FunctionCallStart = TokenType{
		Name:    "FunctionCallStart",
		Keyword: "(",
	}
	FunctionCallEnd = TokenType{
		Name:    "FunctionCallEnd",
		Keyword: ")",
	}
	EndOfLine = TokenType{
		Name:    "EndOfLine",
		Keyword: "!",
	}
	StringStartEnd = TokenType{
		Name:    "StringStartEnd",
		Keyword: "\"",
	}

	AllTokenTypes = []TokenType{
		DeclarationVar,
		DeclarationFunc,
		FunctionCallStart,
		FunctionCallEnd,
		EndOfLine,
		StringStartEnd,
	}
)
