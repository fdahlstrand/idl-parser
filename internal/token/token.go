package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// 7.2.5 - Punctation
	SEMICOLON   = ";"
	LCURLY      = "{"
	RCURLY      = "}"
	COLON       = ":"
	COMMA       = ","
	EQUALS      = "="
	PLUS        = "+"
	MINUS       = "-"
	LPAREN      = "("
	RPAREN      = ")"
	LT          = "<"
	GT          = ">"
	LBRACKET    = "["
	RBRACKET    = "]"
	SINGLEQUOTE = "'"
	DOUBLEQUOTE = "\""
	BACKSLASH   = "\\"
	VERTBAR     = "|"
	CARET       = "^"
	AMPERSAND   = "&"
	ASTERISK    = "*"
	SLASH       = "/"
	PERCENT     = "%"
	TILDE       = "~"
	AT          = "@"
)
