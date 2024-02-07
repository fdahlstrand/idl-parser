package lexer

import (
	"testing"

	"github.com/fdahlstrand/idl-parser/internal/token"
)

func TestPunctation(t *testing.T) {
	input := ";{}:,=+-()<>[]'\"\\|^&*/%~@"

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.SEMICOLON, ";"},
		{token.LCURLY, "{"},
		{token.RCURLY, "}"},
		{token.COLON, ":"},
		{token.COMMA, ","},
		{token.EQUALS, "="},
		{token.PLUS, "+"},
		{token.MINUS, "-"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LT, "<"},
		{token.GT, ">"},
		{token.LBRACKET, "["},
		{token.RBRACKET, "]"},
		{token.SINGLEQUOTE, "'"},
		{token.DOUBLEQUOTE, "\""},
		{token.BACKSLASH, "\\"},
		{token.VERTBAR, "|"},
		{token.CARET, "^"},
		{token.AMPERSAND, "&"},
		{token.ASTERISK, "*"},
		{token.SLASH, "/"},
		{token.PERCENT, "%"},
		{token.TILDE, "~"},
		{token.AT, "@"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf(
				"Punctation Test[%d] - Wrong Type. expected=%+q, actual=%+q",
				i,
				tt.expectedType,
				tok.Type,
			)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf(
				"Punctation Test[%d] - Wrong Literal. expected=%+q, actual=%+q",
				i,
				tt.expectedLiteral,
				tok.Literal,
			)
		}
	}
}
