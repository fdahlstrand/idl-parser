package lexer

import (
	"fmt"
	"strconv"
	"unicode/utf8"

	"github.com/fdahlstrand/idl-parser/internal/token"
)

type Lexer struct {
	input   string
	readPos int
	pos     [3]int
	ch      [3]rune
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.advance(3)
	return l
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	// == SEPARATORS & OPERATORS ============================================
	if l.ch[0] == ';' {
		tok = newToken(token.SEMICOLON, string(l.ch[0]))
	} else if l.ch[0] == '{' {
		tok = newToken(token.LCURLY, string(l.ch[0]))
	} else if l.ch[0] == '}' {
		tok = newToken(token.RCURLY, string(l.ch[0]))
	} else if l.ch[0] == ',' {
		tok = newToken(token.COMMA, string(l.ch[0]))
	} else if l.ch[0] == '=' {
		tok = newToken(token.EQUALS, string(l.ch[0]))
	} else if l.ch[0] == '+' {
		tok = newToken(token.PLUS, string(l.ch[0]))
	} else if l.ch[0] == '-' {
		tok = newToken(token.MINUS, string(l.ch[0]))
	} else if l.ch[0] == '(' {
		tok = newToken(token.LPAREN, string(l.ch[0]))
	} else if l.ch[0] == ')' {
		tok = newToken(token.RPAREN, string(l.ch[0]))
	} else if l.ch[0] == '[' {
		tok = newToken(token.LBRACKET, string(l.ch[0]))
	} else if l.ch[0] == ']' {
		tok = newToken(token.RBRACKET, string(l.ch[0]))
	} else if l.ch[0] == '|' {
		tok = newToken(token.OR, string(l.ch[0]))
	} else if l.ch[0] == '^' {
		tok = newToken(token.XOR, string(l.ch[0]))
	} else if l.ch[0] == '&' {
		tok = newToken(token.AND, string(l.ch[0]))
	} else if l.ch[0] == '*' {
		tok = newToken(token.MUL, string(l.ch[0]))
	} else if l.ch[0] == '/' {
		tok = newToken(token.DIV, string(l.ch[0]))
	} else if l.ch[0] == '%' {
		tok = newToken(token.MODULO, string(l.ch[0]))
	} else if l.ch[0] == '~' {
		tok = newToken(token.NOT, string(l.ch[0]))
	} else if l.ch[0] == ':' && l.ch[1] == ':' {
		l.advance(1)
		tok = newToken(token.SCOPESEP, "::")
	} else if l.ch[0] == ':' {
		tok = newToken(token.COLON, ":")
	} else if l.ch[0] == '<' && l.ch[1] == '<' {
		l.advance(1)
		tok = newToken(token.LSHIFT, "<<")
	} else if l.ch[0] == '>' && l.ch[1] == '>' {
		l.advance(1)
		tok = newToken(token.RSHIFT, ">>")
	} else if l.ch[0] == '<' {
		tok = newToken(token.LT, "<")
	} else if l.ch[0] == '>' {
		tok = newToken(token.GT, ">")

		// == CHARACTER LITERALS ============================================
	} else if l.ch[0] == '\'' || (l.ch[0] == 'L' && l.ch[1] == '\'') {
		wide := l.ch[0] == 'L'

		if wide {
			l.advance(2)
		} else {
			l.advance(1)
		}

		ch_lit, err := l.readCharLiteral(wide)
		if l.ch[0] == '\'' {
			l.advance(1)
			if wide {
				tok = newToken(token.WCHAR_LITERAL, ch_lit)
			} else {
				tok = newToken(token.CHAR_LITERAL, ch_lit)
			}
		} else {
			if err == nil {
				err = fmt.Errorf("Syntax Error: Character literal not terminated")
			}
			tok = newToken(token.ILLEGAL, err.Error())
		}

		// == IDENTIFIERS & KEYWORDS ========================================
	} else if isAlpha(l.ch[0]) {
		ident := l.readIdentifier()
		if tt, ok := token.Keywords[ident]; ok {
			tok = newToken(tt, ident)
		} else {
			tok = newToken(token.IDENTIFIER, ident)
		}
	} else if l.ch[0] == '_' {
		l.advance(1)
		ident := l.readIdentifier()
		if ident != "" {
			tok = newToken(token.IDENTIFIER, ident)
		} else {
			tok = newToken(token.ILLEGAL, "")
		}
		// == INTEGER LITERALS ==============================================
	} else if l.ch[0] == '0' && isOctalDigit(l.ch[1]) {
		oct := l.readOctalInteger()
		tok = newToken(token.INTEGER, oct)
	} else if l.ch[0] == '0' && (l.ch[1] == 'x' || l.ch[1] == 'X') && isHexDigit(l.ch[2]) {
		prefix := "0" + string(l.ch[1])
		l.advance(2)
		hex := l.readHexInteger()
		tok = newToken(token.INTEGER, prefix+hex)
	} else if isDigit(l.ch[0]) {
		dec := l.readDecimalInteger()
		tok = newToken(token.INTEGER, dec)

		// == END OF FILE ===================================================
	} else if l.ch[0] == 0 {
		tok = newToken(token.EOF, "")

		// == SYNTAX ERROR ==================================================
	} else {
		tok = newToken(token.ILLEGAL, string(l.ch[0]))
	}

	l.advance(1)
	return tok
}

func (l *Lexer) skipWhitespace() {
	for isWhitespace(l.ch[0]) {
		l.advance(1)
	}
}

func (l *Lexer) readIdentifier() string {
	pos := l.pos[0]

	if isAlpha(l.ch[0]) {
		l.advance(1)
		for isAlpha(l.ch[0]) || isDigit(l.ch[0]) || l.ch[0] == '_' {
			l.advance(1)
		}
	}

	return l.input[pos:l.pos[0]]
}

func (l *Lexer) readDecimalInteger() string {
	pos := l.pos[0]

	for isDigit(l.ch[0]) {
		l.advance(1)
	}

	return l.input[pos:l.pos[0]]
}

func (l *Lexer) readOctalInteger() string {
	pos := l.pos[0]

	for isOctalDigit(l.ch[0]) {
		l.advance(1)
	}

	return l.input[pos:l.pos[0]]
}

func (l *Lexer) readHexInteger() string {
	pos := l.pos[0]

	for isHexDigit(l.ch[0]) {
		l.advance(1)
	}

	return l.input[pos:l.pos[0]]
}

func (l *Lexer) readCharLiteral(wide bool) (lit string, err error) {
	var ch_lit string
	if l.ch[0] == '\\' {
		l.advance(1)
		ch_lit, err = l.readEscapeCharacter(wide)
		if err != nil {
			return "", err
		}
	} else if l.ch[0] != '\'' {
		ch_lit = string(l.ch[0])
		l.advance(1)
	}
	return ch_lit, nil
}

func (l *Lexer) readEscapeCharacter(wide bool) (lit string, err error) {
	escapes := map[rune]string{
		'n':  "\n",
		't':  "\t",
		'v':  "\v",
		'b':  "\b",
		'r':  "\r",
		'f':  "\f",
		'a':  "\a",
		'\\': "\\",
		'?':  "?",
		'\'': "'",
		'"':  "\"",
	}

	if lit, ok := escapes[l.ch[0]]; ok {
		l.advance(1)
		return lit, nil
	} else if l.ch[0] == 'x' {
		l.advance(1)
		if isHexDigit(l.ch[0]) {
			pos := l.pos[0]
			for n := 0; n < 2 && isHexDigit(l.ch[0]); n++ {
				l.advance(1)
			}
			code, _ := strconv.ParseUint(l.input[pos:l.pos[0]], 16, 16)
			return string(rune(code)), nil
		} else {
			return "", fmt.Errorf("Syntax Error: Illegal character '%c' in escape sequence", l.ch[0])
		}
	} else if wide && l.ch[0] == 'u' {
		l.advance(1)
		if isHexDigit(l.ch[0]) {
			pos := l.pos[0]
			for n := 0; n < 4 && isHexDigit(l.ch[0]); n++ {
				l.advance(1)
			}
			code, _ := strconv.ParseUint(l.input[pos:l.pos[0]], 16, 16)
			return string(rune(code)), nil
		} else {
			return "", fmt.Errorf("Syntax Error: Illegal character '%c' in escape sequence", l.ch[0])
		}
	} else if isOctalDigit(l.ch[0]) {
		pos := l.pos[0]
		for n := 0; n < 3 && isOctalDigit(l.ch[0]); n++ {
			l.advance(1)
		}
		code, _ := strconv.ParseUint(l.input[pos:l.pos[0]], 8, 16)
		return string(rune(code)), nil
	} else {
		return "", fmt.Errorf("Syntax Error: Unknown escape sequence '\\%c'", l.ch[0])
	}
}

func (l *Lexer) advance(n int) {
	for i := 0; i < n; i++ {
		var w int = 0
		l.ch[0] = l.ch[1]
		l.ch[1] = l.ch[2]
		l.pos[0] = l.pos[1]
		l.pos[1] = l.pos[2]
		if l.readPos < len(l.input) {
			var ch rune
			ch, w = utf8.DecodeRuneInString(l.input[l.readPos:])
			l.ch[2] = ch
		} else {
			l.ch[2] = 0
		}
		l.pos[2] = l.readPos
		l.readPos += w
	}
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
