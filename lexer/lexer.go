package lexer

import "github.com/Infi-Knight/vonur/token"

// Lexer with abilities to peek and read tokens.
type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // points to the next character in the input
	ch           byte // current char under examination
}

// NewLexer creates and initialises a Lexer and returns a reference to it
func NewLexer(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

// readChar gives us next character in the input string and advances
// out position to the next character. Our lexer does not support
// full range of Unicode code points
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		// “ASCII code for the "NULL" character and signifies either
		// "we haven't read anything yet" or "end of file" for us. ”
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
}

// readIdentifier reads in an identifier and advances our lexer's positions
// until it encounters a non-letter-character.
func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

// NextToken returns the next token from the input sequence
func (l *Lexer) NextToken() token.Token {
	var tok token.Token
	l.skipWhiteSpace()

	switch l.ch {
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
		} else if isDigit(l.ch) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
		return tok
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = "EOF"
	}
	l.readChar()
	return tok
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

// isLetter helper function just checks whether the given argument is a letter
// In our case it contains the check ch == '_', which means that we'll treat
// _ as a letter and allow it in identifiers and keywords.
// That means we can use variable names like foo_bar.
func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

// whitespace is not significant in Monkey so we need to skip it
func (l *Lexer) skipWhiteSpace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

// For now, Monkey only supports integer
func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}
