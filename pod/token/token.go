package token

import "fmt"

// Token is the token used in pod job.
type Token struct {
	Type    Type
	Content string
	Pos     Pos
}

//go:generate stringer -type=Type
type Type int

const (
	// Identifier
	IDENT Type = iota

	// Type
	STRING     // "abc"
	HEREDOC    // <<EOF\nabc\nEOF
	INTEGER    // 123, 0x123, 0o123, 0b110
	FLOAT      // 123.123
	BOOL       // True
	STATEMENT  // `create_card xxx`

	// Numeric
	ADD      // +
	SUB      // -
	STAR     // *
	SLASH    // /
	PERCENT  // %
	POWER    // ^

	// Logic
	EQUAL     // ==
	NOTEQUAL  // !=
	GT        // >
	LT        // <
	GTE       // >=
	LTE       // <=
	AND       // &&
	OR        // ||

	// Operate
	LPAREN    // (
	RPAREN    // )
	LBRACKET  // [
	RBRACKET  // ]
	LBRACE    // {
	RBRACE    // }
	COMMA     // ,
	PERIOD    // .
	QUOTE     // "

	ASSIGN   // =
	EOF      // EOF
	INVALID  // INVALID
)

// String will output current token's content.
func (t *Token) String() string {
	switch t.Type {
	case EOF:
		return "end of string"
	case INVALID:
		return fmt.Sprintf("invalid sequence %q", t.Content)
	case STRING:
		return fmt.Sprintf("string %q", t.Content)
	case HEREDOC:
		return fmt.Sprintf("heredoc %q", t.Content)
	case INTEGER:
		return fmt.Sprintf("integer %s", t.Content)
	case FLOAT:
		return fmt.Sprintf("float %s", t.Content)
	case STATEMENT:
		return fmt.Sprintf("statement `%s`", t.Content)
	default:
		// The remaining token types have content that
		// speaks for itself.
		return fmt.Sprintf("%s", t.Content)
	}
}
