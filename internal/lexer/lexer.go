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

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	switch l.ch {
	case ';':
		tok = newToken(token.SEMICOLON, ";")
	case '{':
		tok = newToken(token.LCURLY, "{")
	case '}':
		tok = newToken(token.RCURLY, "}")
	case ':':
		tok = newToken(token.COLON, ":")
	case ',':
		tok = newToken(token.COMMA, ",")
	case '=':
		tok = newToken(token.EQUALS, "=")
	case '+':
		tok = newToken(token.PLUS, "+")
	case '-':
		tok = newToken(token.MINUS, "-")
	case '(':
		tok = newToken(token.LPAREN, "(")
	case ')':
		tok = newToken(token.RPAREN, ")")
	case '<':
		tok = newToken(token.LT, "<")
	case '>':
		tok = newToken(token.GT, ">")
	case '[':
		tok = newToken(token.LBRACKET, "[")
	case ']':
		tok = newToken(token.RBRACKET, "]")
	case '\'':
		tok = newToken(token.SINGLEQUOTE, "'")
	case '"':
		tok = newToken(token.DOUBLEQUOTE, "\"")
	case '\\':
		tok = newToken(token.BACKSLASH, "\\")
	case '|':
		tok = newToken(token.VERTBAR, "|")
	case '^':
		tok = newToken(token.CARET, "^")
	case '&':
		tok = newToken(token.AMPERSAND, "&")
	case '*':
		tok = newToken(token.ASTERISK, "*")
	case '/':
		tok = newToken(token.SLASH, "/")
	case '%':
		tok = newToken(token.PERCENT, "%")
	case '~':
		tok = newToken(token.TILDE, "~")
	case '@':
		tok = newToken(token.AT, "@")
	case 0:
		tok = newToken(token.EOF, "")
	default:
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
