package lexer

import (
	"unicode/utf8"

	"github.com/fdahlstrand/idl-parser/internal/token"
)

type Lexer struct {
	input   string
	pos     int
	readPos int
	ch      rune
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readRune()
	return l
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	case ';':
		tok = newToken(token.SEMICOLON, string(l.ch))
	case '{':
		tok = newToken(token.LCURLY, string(l.ch))
	case '}':
		tok = newToken(token.RCURLY, string(l.ch))
	case ',':
		tok = newToken(token.COMMA, string(l.ch))
	case '=':
		tok = newToken(token.EQUALS, string(l.ch))
	case '+':
		tok = newToken(token.PLUS, string(l.ch))
	case '-':
		tok = newToken(token.MINUS, string(l.ch))
	case '(':
		tok = newToken(token.LPAREN, string(l.ch))
	case ')':
		tok = newToken(token.RPAREN, string(l.ch))
	case '[':
		tok = newToken(token.LBRACKET, string(l.ch))
	case ']':
		tok = newToken(token.RBRACKET, string(l.ch))
	case '|':
		tok = newToken(token.OR, string(l.ch))
	case '^':
		tok = newToken(token.XOR, string(l.ch))
	case '&':
		tok = newToken(token.AND, string(l.ch))
	case '*':
		tok = newToken(token.MUL, string(l.ch))
	case '/':
		tok = newToken(token.DIV, string(l.ch))
	case '%':
		tok = newToken(token.MODULO, string(l.ch))
	case '~':
		tok = newToken(token.NOT, string(l.ch))
	case ':':
		l.readRune()
		if l.ch == ':' {
			tok = newToken(token.SCOPESEP, "::")
		} else {
			return newToken(token.COLON, ":")
		}
	case '<':
		l.readRune()
		if l.ch == '<' {
			tok = newToken(token.LSHIFT, "<<")
		} else {
			return newToken(token.LT, "<")
		}
	case '>':
		l.readRune()
		if l.ch == '>' {
			tok = newToken(token.RSHIFT, ">>")
		} else {
			return newToken(token.GT, ">")
		}
	default:
		if isAlpha(l.ch) {
			ident := l.readIdentifier()
			if tt, ok := token.Keywords[ident]; ok {
				tok = newToken(tt, ident)
			} else {
				tok = newToken(token.IDENTIFIER, ident)
			}
		} else if l.ch == '_' {
			l.readRune()
			ident := l.readIdentifier()
			if ident != "" {
				tok = newToken(token.IDENTIFIER, ident)
			} else {
				tok = newToken(token.ILLEGAL, "")
			}
		} else if l.ch == '0' {
			l.readRune()
			if isOctalDigit(l.ch) {
				oct := l.readOctalInteger()
				tok = newToken(token.INTEGER, "0"+oct)
			} else if l.ch == 'x' || l.ch == 'X' {
				prefix := "0" + string(l.ch)
				l.readRune()

				if isHexDigit(l.ch) {
					hex := l.readHexInteger()
					tok = newToken(token.INTEGER, prefix+hex)
				} else {
					tok = newToken(token.ILLEGAL, prefix)
				}
			} else {
				return newToken(token.INTEGER, "0")
			}
		} else if isDigit(l.ch) {
			dec := l.readDecimalInteger()
			tok = newToken(token.INTEGER, dec)
		} else if l.ch == 0 {
			tok = newToken(token.EOF, "")
		} else {
			tok = newToken(token.ILLEGAL, string(l.ch))
		}
	}

	l.readRune()
	return tok
}

func (l *Lexer) skipWhitespace() {
	for isWhitespace(l.ch) {
		l.readRune()
	}
}

func (l *Lexer) readIdentifier() string {
	pos := l.pos

	if isAlpha(l.ch) {
		l.readRune()
		for isAlpha(l.ch) || isDigit(l.ch) || l.ch == '_' {
			l.readRune()
		}
	}

	return l.input[pos:l.pos]
}

func (l *Lexer) readDecimalInteger() string {
	pos := l.pos

	for isDigit(l.ch) {
		l.readRune()
	}

	return l.input[pos:l.pos]
}

func (l *Lexer) readOctalInteger() string {
	pos := l.pos

	for isOctalDigit(l.ch) {
		l.readRune()
	}

	return l.input[pos:l.pos]
}

func (l *Lexer) readHexInteger() string {
	pos := l.pos

	for isHexDigit(l.ch) {
		l.readRune()
	}

	return l.input[pos:l.pos]
}

func (l *Lexer) readRune() {
	var w int = 0
	if l.readPos < len(l.input) {
		var ch rune
		ch, w = utf8.DecodeRuneInString(l.input[l.readPos:])
		l.ch = ch
	} else {
		l.ch = 0
	}
	l.pos = l.readPos
	l.readPos += w
}

func isWhitespace(ch rune) bool {
	return ch == ' ' || ch == '\t' || ch == '\n' || ch == '\r' || ch == '\013'
}

func isAlpha(ch rune) bool {
	return ('a' <= ch && ch <= 'z') || ('A' <= ch && ch <= 'Z')
}

func isDigit(ch rune) bool {
	return ('0' <= ch && ch <= '9')
}

func isOctalDigit(ch rune) bool {
	return ('0' <= ch && ch <= '7')
}

func isHexDigit(ch rune) bool {
	return ('0' <= ch && ch <= '9') || ('A' <= ch && ch <= 'F') || ('a' <= ch && ch <= 'f')
}

func newToken(typ token.TokenType, lit string) token.Token {
	return token.Token{Type: typ, Literal: lit}
}
