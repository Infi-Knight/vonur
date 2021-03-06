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

// IDENT identifiers e.g foo, bar, add.
// INT integer literals e.g 456
const (
	IDENT = "IDENT"
	INT   = "INT"
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
