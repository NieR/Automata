package token

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToken_String(t *testing.T) {
	cases := []struct {
		input  Token
		output string
	}{
		{
			Token{
				Type:    EOF,
				Content: "",
			},
			"end of string",
		},
		{
			Token{
				Type:    INVALID,
				Content: "forbar",
			},
			`invalid sequence "forbar"`,
		},
		{
			Token{
				Type:    STRING,
				Content: "hello",
			},
			`string "hello"`,
		},
		{
			Token{
				Type:    HEREDOC,
				Content: "long string",
			},
			`heredoc "long string"`,
		},
		{
			Token{
				Type:    INTEGER,
				Content: "123",
			},
			"integer 123",
		},
		{
			Token{
				Type:    FLOAT,
				Content: "1.23",
			},
			"float 1.23",
		},
		{
			Token{
				Type:    STATEMENT,
				Content: "have light",
			},
			"statement `have light`",
		},
		{
			Token{
				Type:    EQUAL,
				Content: "==",
			},
			"==",
		},
	}

	for _, v := range cases {
		assert.Equal(t, v.output, v.input.String())
	}
}
