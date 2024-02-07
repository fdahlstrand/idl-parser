package lexer

import (
	"unicode/utf8"

	"github.com/fdahlstrand/idl-parser/internal/token"
)

type Lexer struct {
	input string
	pos   int
	ch    rune
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.nextRune()
	return l
}

var punctation = map[rune]token.TokenType{
	';':  token.SEMICOLON,
	'{':  token.LCURLY,
	'}':  token.RCURLY,
	':':  token.COLON,
	',':  token.COMMA,
	'=':  token.EQUALS,
	'+':  token.PLUS,
	'-':  token.MINUS,
	'(':  token.LPAREN,
	')':  token.RPAREN,
	'<':  token.LT,
	'>':  token.GT,
	'[':  token.LBRACKET,
	']':  token.RBRACKET,
	'\'': token.SINGLEQUOTE,
	'"':  token.DOUBLEQUOTE,
	'\\': token.BACKSLASH,
	'|':  token.VERTBAR,
	'^':  token.CARET,
	'&':  token.AMPERSAND,
	'*':  token.ASTERISK,
	'/':  token.SLASH,
	'%':  token.PERCENT,
	'~':  token.TILDE,
	'@':  token.AT,
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	if tt, ok := punctation[l.ch]; ok {
		tok = newToken(tt, string(l.ch))
	} else if l.ch == 0 {
		tok = newToken(token.EOF, "")
	} else {
		tok = newToken(token.ILLEGAL, string(l.ch))
	}

	l.nextRune()
	return tok
}

func (l *Lexer) nextRune() {
	if l.pos < len(l.input) {
		ch, w := utf8.DecodeRuneInString(l.input[l.pos:])
		l.ch = ch
		l.pos += w
	} else {
		l.ch = 0
	}
}

func newToken(typ token.TokenType, lit string) token.Token {
	return token.Token{Type: typ, Literal: lit}
}
