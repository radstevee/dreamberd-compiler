package ast

import (
	"dreamberd-compiler/common"
	"strings"
)

func Tokenize(line string) []common.Token {
	keywords := make(map[string]common.TokenType)
	for _, tokenType := range common.AllTokenTypes {
		keywords[tokenType.Keyword] = tokenType
	}

	// split the line down to words
	var words []string
	for i := 0; i < len(line); i++ {
		if line[i] == '"' {
			// handle string literals
			end := strings.Index(line[i:], "\"")
			if end == -1 {
				panic("unterminated string literal")
			}
			// what the actual fuck is `line[i:i+end+1]`??
			words = append(words, line[i:i+end+1])
			i += end
		} else if line[i] == ' ' || line[i] == '(' || line[i] == ')' || line[i] == '!' {
			words = append(words, strings.TrimSpace(line[:i]))
			words = append(words, string(line[i]))
			line = line[i+1:]
			i = 0
		}
	}

	// create tokens from the words
	var tokens []common.Token
	for i := range words {
		if keyword, ok := keywords[words[i]]; ok {
			tokens = append(tokens, common.Token{
				Line:   0,
				Column: i,
				Type:   keyword,
			})
		} else if words[i] != "" {
			// check if it is supposed to be a string literal
			if words[i][0] == '"' && words[i][len(words[i])-1] == '"' {
				tokens = append(tokens, common.Token{
					Line:   0,
					Column: i,
					Type:   common.TokenType{Name: "Literal", Keyword: words[i]},
				})
				continue
			}

			// check if it is supposed to be a function call
			if i+1 < len(words) && words[i+1] == "(" {
				tokens = append(tokens, common.Token{
					Line:   0,
					Column: i,
					Type:   common.TokenType{Name: "FunctionCall", Keyword: words[i]},
				})
				continue
			}
		}
	}

	return tokens
}
