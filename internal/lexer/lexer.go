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

var keywords = map[string]token.TokenType{
	"abstract":    token.ABSTRACT,
	"any":         token.ANY,
	"alias":       token.ALIAS,
	"attribute":   token.ATTRIBUTE,
	"bitfield":    token.BITFIELD,
	"bitmask":     token.BITMASK,
	"bitset":      token.BITSET,
	"boolean":     token.BOOLEAN,
	"case":        token.CASE,
	"char":        token.CHAR,
	"component":   token.COMPONENT,
	"connector":   token.CONNECTOR,
	"const":       token.CONST,
	"consumes":    token.CONSUMES,
	"context":     token.CONTEXT,
	"custom":      token.CUSTOM,
	"default":     token.DEFAULT,
	"double":      token.DOUBLE,
	"exception":   token.EXCEPTION,
	"emits":       token.EMITS,
	"enum":        token.ENUM,
	"eventtype":   token.EVENTTYPE,
	"factory":     token.FACTORY,
	"FALSE":       token.FALSE,
	"finder":      token.FINDER,
	"fixed":       token.FIXED,
	"float":       token.FLOAT,
	"getraises":   token.GETRAISES,
	"home":        token.HOME,
	"import":      token.IMPORT,
	"in":          token.IN,
	"inout":       token.INOUT,
	"interface":   token.INTERFACE,
	"local":       token.LOCAL,
	"long":        token.LONG,
	"manages":     token.MANAGES,
	"map":         token.MAP,
	"mirrorport":  token.MIRRORPORT,
	"module":      token.MODULE,
	"multiple":    token.MULTIPLE,
	"native":      token.NATIVE,
	"Object":      token.OBJECT,
	"octet":       token.OCTET,
	"oneway":      token.ONEWAY,
	"out":         token.OUT,
	"primarykey":  token.PRIMARYKEY,
	"private":     token.PRIVATE,
	"port":        token.PORT,
	"porttype":    token.PORTTYPE,
	"provides":    token.PROVIDES,
	"public":      token.PUBLIC,
	"publishes":   token.PUBLISHES,
	"raises":      token.RAISES,
	"readonly":    token.READONLY,
	"setraises":   token.SETRAISES,
	"sequence":    token.SEQUENCE,
	"short":       token.SHORT,
	"string":      token.STRING,
	"struct":      token.STRUCT,
	"supports":    token.SUPPORTS,
	"switch":      token.SWITCH,
	"TRUE":        token.TRUE,
	"truncatable": token.TRUNCATABLE,
	"typedef":     token.TYPEDEF,
	"typeid":      token.TYPEID,
	"typename":    token.TYPENAME,
	"typeprefix":  token.TYPEPREFIX,
	"unsigned":    token.UNSIGNED,
	"union":       token.UNION,
	"uses":        token.USES,
	"ValueBase":   token.VALUEBASE,
	"valuetype":   token.VALUETYPE,
	"void":        token.VOID,
	"wchar":       token.WCHAR,
	"wstring":     token.WSTRING,
	"int8":        token.INT8,
	"uint8":       token.UINT8,
	"int16":       token.INT16,
	"int32":       token.INT32,
	"int64":       token.INT64,
	"uint16":      token.UINT16,
	"uint32":      token.UINT32,
	"uint64":      token.UINT64,
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	if tt, ok := punctation[l.ch]; ok {
		tok = newToken(tt, string(l.ch))
	} else if isStartOfIdent(l.ch) {
		kw := l.readKeyword()
		if tt, ok := keywords[kw]; ok {
			tok = newToken(tt, kw)
		} else {
			println(kw)
			tok = newToken(token.ILLEGAL, kw)
		}
	} else if l.ch == 0 {
		tok = newToken(token.EOF, "")
	} else {
		tok = newToken(token.ILLEGAL, string(l.ch))
	}

	l.readRune()
	return tok
}

func (l *Lexer) readKeyword() string {
	pos := l.pos
	for isIdentChar(l.ch) {
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

func isStartOfIdent(ch rune) bool {
	return ('a' <= ch && ch <= 'z') || ('A' <= ch && ch <= 'Z') || (ch == '_')
}

func isIdentChar(ch rune) bool {
	return isStartOfIdent(ch) || ('0' <= ch && ch <= '9')
}

func newToken(typ token.TokenType, lit string) token.Token {
	return token.Token{Type: typ, Literal: lit}
}
