/*
Package token implements the data structures and types defined by the Monkey
programming language.
*/
package token

// TokenType we will use string to store the type of our tokens.
type TokenType string

// Token a token has a type and a value.
type Token struct {
	Type    TokenType
	Literal string
}

// ILLEGAL represents a token/character we don't know about.
// EOF tells the parser where to stop.
const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"
)

// IDENTIFIER identifiers e.g foo, bar, add.
// INT integer literals e.g 456
const (
	IDENTIFIER = "IDENTIFIER"
	INT        = "INT"
)

// Operators
const (
	ASSIGN = "="
	PLUS   = "+"
)

// Delimiters
const (
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"
)

// Keywords
const (
	FUNCTION = "FUNCTION"
	LET      = "LET"
)

// keywords: map of keywords in our language and their token type
var keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
}

// LookupIdent checks the keywords table to see whether
// the given identifier is infact a language keyword
func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENTIFIER
}
