package lexer

import (
	"testing"

	"github.com/fdahlstrand/idl-parser/internal/token"
)

var tests = []struct {
	input           string
	expectedType    token.TokenType
	expectedLiteral string
}{
	// Delimiters
	{";", token.SEMICOLON, ";"},
	{"{", token.LCURLY, "{"},
	{"}", token.RCURLY, "}"},
	{"::", token.SCOPESEP, "::"},
	{"(", token.LPAREN, "("},
	{")", token.RPAREN, ")"},
	{"<", token.LT, "<"},
	{">", token.GT, ">"},
	{",", token.COMMA, ","},
	{":", token.COLON, ":"},
	{"[", token.LBRACKET, "["},
	{"]", token.RBRACKET, "]"},

	// Operators
	{"=", token.EQUALS, "="},
	{"|", token.OR, "|"},
	{"^", token.XOR, "^"},
	{"&", token.AND, "&"},
	{">>", token.RSHIFT, ">>"},
	{"<<", token.LSHIFT, "<<"},
	{"+", token.PLUS, "+"},
	{"-", token.MINUS, "-"},
	{"*", token.MUL, "*"},
	{"/", token.DIV, "/"},
	{"%", token.MODULO, "%"},
	{"~", token.NOT, "~"},

	// Identifiers
	{"a", token.IDENTIFIER, "a"},
	{"x8", token.IDENTIFIER, "x8"},
	{"_abstract", token.IDENTIFIER, "abstract"},
	{"_Y", token.IDENTIFIER, "Y"},
	{"__bad", token.ILLEGAL, ""},
	{"_0notgood", token.ILLEGAL, ""},

	// Integers
	{"0", token.INTEGER, "0"},
	{"122", token.INTEGER, "122"},
	{"012", token.INTEGER, "012"},
	{"0xAB", token.INTEGER, "0xAB"},
	{"0X1a2f", token.INTEGER, "0X1a2f"},
	{"01a2f", token.INTEGER, "01"},
	{"9a2f", token.INTEGER, "9"},
	{"0X", token.INTEGER, "0"},

	// Keywords
	{"abstract", token.ABSTRACT, "abstract"},
	{"any", token.ANY, "any"},
	{"alias", token.ALIAS, "alias"},
	{"attribute", token.ATTRIBUTE, "attribute"},
	{"bitfield", token.BITFIELD, "bitfield"},
	{"bitmask", token.BITMASK, "bitmask"},
	{"bitset", token.BITSET, "bitset"},
	{"boolean", token.BOOLEAN, "boolean"},
	{"case", token.CASE, "case"},
	{"char", token.CHAR, "char"},
	{"component", token.COMPONENT, "component"},
	{"connector", token.CONNECTOR, "connector"},
	{"const", token.CONST, "const"},
	{"consumes", token.CONSUMES, "consumes"},
	{"context", token.CONTEXT, "context"},
	{"custom", token.CUSTOM, "custom"},
	{"default", token.DEFAULT, "default"},
	{"double", token.DOUBLE, "double"},
	{"exception", token.EXCEPTION, "exception"},
	{"emits", token.EMITS, "emits"},
	{"enum", token.ENUM, "enum"},
	{"eventtype", token.EVENTTYPE, "eventtype"},
	{"factory", token.FACTORY, "factory"},
	{"FALSE", token.FALSE, "FALSE"},
	{"finder", token.FINDER, "finder"},
	{"fixed", token.FIXED, "fixed"},
	{"float", token.FLOAT, "float"},
	{"getraises", token.GETRAISES, "getraises"},
	{"home", token.HOME, "home"},
	{"import", token.IMPORT, "import"},
	{"in", token.IN, "in"},
	{"inout", token.INOUT, "inout"},
	{"interface", token.INTERFACE, "interface"},
	{"local", token.LOCAL, "local"},
	{"long", token.LONG, "long"},
	{"manages", token.MANAGES, "manages"},
	{"map", token.MAP, "map"},
	{"mirrorport", token.MIRRORPORT, "mirrorport"},
	{"module", token.MODULE, "module"},
	{"multiple", token.MULTIPLE, "multiple"},
	{"native", token.NATIVE, "native"},
	{"Object", token.OBJECT, "Object"},
	{"octet", token.OCTET, "octet"},
	{"oneway", token.ONEWAY, "oneway"},
	{"out", token.OUT, "out"},
	{"primarykey", token.PRIMARYKEY, "primarykey"},
	{"private", token.PRIVATE, "private"},
	{"port", token.PORT, "port"},
	{"porttype", token.PORTTYPE, "porttype"},
	{"provides", token.PROVIDES, "provides"},
	{"public", token.PUBLIC, "public"},
	{"publishes", token.PUBLISHES, "publishes"},
	{"raises", token.RAISES, "raises"},
	{"readonly", token.READONLY, "readonly"},
	{"setraises", token.SETRAISES, "setraises"},
	{"sequence", token.SEQUENCE, "sequence"},
	{"short", token.SHORT, "short"},
	{"string", token.STRING, "string"},
	{"struct", token.STRUCT, "struct"},
	{"supports", token.SUPPORTS, "supports"},
	{"switch", token.SWITCH, "switch"},
	{"TRUE", token.TRUE, "TRUE"},
	{"truncatable", token.TRUNCATABLE, "truncatable"},
	{"typedef", token.TYPEDEF, "typedef"},
	{"typeid", token.TYPEID, "typeid"},
	{"typename", token.TYPENAME, "typename"},
	{"typeprefix", token.TYPEPREFIX, "typeprefix"},
	{"unsigned", token.UNSIGNED, "unsigned"},
	{"union", token.UNION, "union"},
	{"uses", token.USES, "uses"},
	{"ValueBase", token.VALUEBASE, "ValueBase"},
	{"valuetype", token.VALUETYPE, "valuetype"},
	{"void", token.VOID, "void"},
	{"wchar", token.WCHAR, "wchar"},
	{"wstring", token.WSTRING, "wstring"},
	{"int8", token.INT8, "int8"},
	{"uint8", token.UINT8, "uint8"},
	{"int16", token.INT16, "int16"},
	{"int32", token.INT32, "int32"},
	{"int64", token.INT64, "int64"},
	{"uint16", token.UINT16, "uint16"},
	{"uint32", token.UINT32, "uint32"},
	{"uint64", token.UINT64, "uint64"},
}

func TestSingleToken(t *testing.T) {
	for _, test := range tests {
		l := New(test.input)
		tok := l.NextToken()
		assertToken(
			t,
			test.input,
			token.Token{Type: test.expectedType, Literal: test.expectedLiteral},
			tok,
		)
	}
}

func TestWhitespace(t *testing.T) {
	input := "module\tfoo\n   {\r\n\013}"

	expectedTokens := []token.Token{
		{Type: token.MODULE, Literal: "module"},
		{Type: token.IDENTIFIER, Literal: "foo"},
		{Type: token.LCURLY, Literal: "{"},
		{Type: token.RCURLY, Literal: "}"},
		{Type: token.EOF, Literal: ""},
	}

	l := New(input)
	for _, et := range expectedTokens {
		tok := l.NextToken()
		assertToken(t, input, et, tok)
	}
}

func assertToken(t *testing.T, input string, expected token.Token, actual token.Token) {
	if expected != actual {
		t.Fatalf(
			"token type assertion failed '%s', expected=%+q, actual=%+q",
			input,
			expected,
			actual,
		)
	}
}
