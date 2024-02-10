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

func TestKeyword(t *testing.T) {
	input := `abstract any alias attribute bitfield
	          bitmask bitset boolean case char
	          component connector const consumes context
	          custom default double exception emits
	          enum eventtype factory FALSE finder
	          fixed float getraises home import
	          in inout interface local long
	          manages map mirrorport module multiple
	          native Object octet oneway out
	          primarykey private port porttype provides
	          public publishes raises readonly setraises
	          sequence short string struct supports 
	          switch TRUE truncatable typedef typeid
	          typename typeprefix unsigned union uses
	          ValueBase valuetype void wchar wstring
	          int8 uint8 int16 int32 int64
	          uint16 uint32 uint64`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ABSTRACT, "abstract"},
		{token.ANY, "any"},
		{token.ALIAS, "alias"},
		{token.ATTRIBUTE, "attribute"},
		{token.BITFIELD, "bitfield"},
		{token.BITMASK, "bitmask"},
		{token.BITSET, "bitset"},
		{token.BOOLEAN, "boolean"},
		{token.CASE, "case"},
		{token.CHAR, "char"},
		{token.COMPONENT, "component"},
		{token.CONNECTOR, "connector"},
		{token.CONST, "const"},
		{token.CONSUMES, "consumes"},
		{token.CONTEXT, "context"},
		{token.CUSTOM, "custom"},
		{token.DEFAULT, "default"},
		{token.DOUBLE, "double"},
		{token.EXCEPTION, "exception"},
		{token.EMITS, "emits"},
		{token.ENUM, "enum"},
		{token.EVENTTYPE, "eventtype"},
		{token.FACTORY, "factory"},
		{token.FALSE, "FALSE"},
		{token.FINDER, "finder"},
		{token.FIXED, "fixed"},
		{token.FLOAT, "float"},
		{token.GETRAISES, "getraises"},
		{token.HOME, "home"},
		{token.IMPORT, "import"},
		{token.IN, "in"},
		{token.INOUT, "inout"},
		{token.INTERFACE, "interface"},
		{token.LOCAL, "local"},
		{token.LONG, "long"},
		{token.MANAGES, "manages"},
		{token.MAP, "map"},
		{token.MIRRORPORT, "mirrorport"},
		{token.MODULE, "module"},
		{token.MULTIPLE, "multiple"},
		{token.NATIVE, "native"},
		{token.OBJECT, "Object"},
		{token.OCTET, "octet"},
		{token.ONEWAY, "oneway"},
		{token.OUT, "out"},
		{token.PRIMARYKEY, "primarykey"},
		{token.PRIVATE, "private"},
		{token.PORT, "port"},
		{token.PORTTYPE, "porttype"},
		{token.PROVIDES, "provides"},
		{token.PUBLIC, "public"},
		{token.PUBLISHES, "publishes"},
		{token.RAISES, "raises"},
		{token.READONLY, "readonly"},
		{token.SETRAISES, "setraises"},
		{token.SEQUENCE, "sequence"},
		{token.SHORT, "short"},
		{token.STRING, "string"},
		{token.STRUCT, "struct"},
		{token.SUPPORTS, "supports"},
		{token.SWITCH, "switch"},
		{token.TRUE, "TRUE"},
		{token.TRUNCATABLE, "truncatable"},
		{token.TYPEDEF, "typedef"},
		{token.TYPEID, "typeid"},
		{token.TYPENAME, "typename"},
		{token.TYPEPREFIX, "typeprefix"},
		{token.UNSIGNED, "unsigned"},
		{token.UNION, "union"},
		{token.USES, "uses"},
		{token.VALUEBASE, "ValueBase"},
		{token.VALUETYPE, "valuetype"},
		{token.VOID, "void"},
		{token.WCHAR, "wchar"},
		{token.WSTRING, "wstring"},
		{token.INT8, "int8"},
		{token.UINT8, "uint8"},
		{token.INT16, "int16"},
		{token.INT32, "int32"},
		{token.INT64, "int64"},
		{token.UINT16, "uint16"},
		{token.UINT32, "uint32"},
		{token.UINT64, "uint64"},
	}

	l := New(input)
	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf(
				"Keyword Test[%d] - Wrong Type expected=%+q, actual=%+q",
				i,
				tt.expectedType,
				tok.Type,
			)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf(
				"Keyword Test[%d] - Wrong Literal expected=%+q, actual=%+q",
				i,
				tt.expectedLiteral,
				tok.Literal,
			)
		}
	}
}

func TestIdentifier(t *testing.T) {
	tests := []struct {
		input           string
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{"a", token.IDENTIFIER, "a"},
		{"x8", token.IDENTIFIER, "x8"},
		{"_abstract", token.IDENTIFIER, "abstract"},
		{"_Y", token.IDENTIFIER, "Y"},
		{"__bad", token.ILLEGAL, ""},
		{"_0notgood", token.ILLEGAL, ""},
	}

	for _, test := range tests {
		l := New(test.input)
		tok := l.NextToken()
		assertType(t, test.input, test.expectedType, tok.Type)
		assertLiteral(t, test.input, test.expectedLiteral, tok.Literal)
	}
}

func assertType(t *testing.T, input string, expected token.TokenType, actual token.TokenType) {
	if expected != actual {
		t.Fatalf("assertion failed '%s', expected=%+q, actual=%+q", input, expected, actual)
	}
}

func assertLiteral(t *testing.T, input string, expected string, actual string) {
	if expected != actual {
		t.Fatalf("assertion failed '%s', expected='%s', actual='%s'", input, expected, actual)
	}
}
