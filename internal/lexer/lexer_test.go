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
	{"Abstract", token.ILLEGAL, "Syntax Error: 'Abstract' collides with keyword 'abstract'"},
	{"vAlUeBaSe", token.ILLEGAL, "Syntax Error: 'vAlUeBaSe' collides with keyword 'ValueBase'"},

	// Integer Literals
	{"0", token.INTEGER, "0"},
	{"122", token.INTEGER, "122"},
	{"012", token.INTEGER, "012"},
	{"0xAB", token.INTEGER, "0xAB"},
	{"0X1a2f", token.INTEGER, "0X1a2f"},
	{"01a2f", token.INTEGER, "01"},
	{"9a2f", token.INTEGER, "9"},
	{"0X", token.INTEGER, "0"},

	// Chracter Literals
	{"'a'", token.CHAR_LITERAL, "a"},
	{"''", token.CHAR_LITERAL, ""},
	{"'a", token.ILLEGAL, "Syntax Error: Character literal not terminated"},
	{"'\\n'", token.CHAR_LITERAL, "\n"},
	{"'\\t'", token.CHAR_LITERAL, "\t"},
	{"'\\v'", token.CHAR_LITERAL, "\v"},
	{"'\\b'", token.CHAR_LITERAL, "\b"},
	{"'\\r'", token.CHAR_LITERAL, "\r"},
	{"'\\f'", token.CHAR_LITERAL, "\f"},
	{"'\\a'", token.CHAR_LITERAL, "\a"},
	{"'\\\\'", token.CHAR_LITERAL, "\\"},
	{"'\\?'", token.CHAR_LITERAL, "?"},
	{"'\\''", token.CHAR_LITERAL, "'"},
	{"'\\\"'", token.CHAR_LITERAL, "\""},
	{"'\\m'", token.ILLEGAL, "Syntax Error: Unknown escape sequence '\\m'"},
	{"'\\xa'", token.CHAR_LITERAL, "\n"},
	{"'\\x41'", token.CHAR_LITERAL, "A"},
	{"'\\x413'", token.ILLEGAL, "Syntax Error: Character literal not terminated"},
	{"'\\xz'", token.ILLEGAL, "Syntax Error: Illegal character 'z' in escape sequence"},
	{"'\\0'", token.CHAR_LITERAL, "\x00"},
	{"'\\42'", token.CHAR_LITERAL, "\""},
	{"'\\172'", token.CHAR_LITERAL, "z"},
	{"'\\4135'", token.ILLEGAL, "Syntax Error: Character literal not terminated"},
	{"'\\42m'", token.ILLEGAL, "Syntax Error: Character literal not terminated"},
	{"L'A'", token.WCHAR_LITERAL, "A"},
	{"L'\\ud'", token.WCHAR_LITERAL, "\r"},
	{"L'\\u58'", token.WCHAR_LITERAL, "X"},
	{"L'\\u14B'", token.WCHAR_LITERAL, "\u014b"},
	{"L'\\u2713'", token.WCHAR_LITERAL, "✓"},
	{"L'\\u42432'", token.ILLEGAL, "Syntax Error: Character literal not terminated"},
	{"L'\\u42U'", token.ILLEGAL, "Syntax Error: Character literal not terminated"},
	{"L'\\up'", token.ILLEGAL, "Syntax Error: Illegal character 'p' in escape sequence"},
	{"'\\u14B'", token.ILLEGAL, "Syntax Error: Unknown escape sequence '\\u'"},

	// String Literals
	{"\"ABC\"", token.STRING_LITERAL, "ABC"},
	{"\"\"", token.STRING_LITERAL, ""},
	{"\"", token.ILLEGAL, "Syntax Error: String literal not terminated"},
	{"\"\\n\"", token.STRING_LITERAL, "\n"},
	{"\"\\n\\t\\v\\b\\r\\f\\\\\\?\\'\\\"\"", token.STRING_LITERAL, "\n\t\v\b\r\f\\?'\""},
	{"\"\\p\"", token.ILLEGAL, "Syntax Error: Unknown escape sequence '\\p'"},
	{"\"\\x9\\x4d\\7\\60\\112\"", token.STRING_LITERAL, "\tM\a0J"},
	{"\"\\xag\"", token.STRING_LITERAL, "\ng"},
	{"\"\\15r&\"", token.STRING_LITERAL, "\rr&"},
	{"\"ABC\\0DEF\"", token.ILLEGAL, "Syntax Error: (null) character not allowed in strings"},
	{"L\"xyz\"", token.WSTRING_LITERAL, "xyz"},
	{"L\"\\uC\\u2A\\u11e\\u2713\\u27145\"", token.WSTRING_LITERAL, "\f*Ğ✓\u27145"},
	{"\"AAA\\u14Bmno\"", token.ILLEGAL, "Syntax Error: Unknown escape sequence '\\u'"},
	{"\"AB\nC\"", token.ILLEGAL, "Syntax Error: String literal not terminated"},

	// Floating Point Literals
	{"1.0", token.FLOATING_PT_LITERAL, "1.0"},
	{"1.", token.FLOATING_PT_LITERAL, "1."},
	{".1", token.FLOATING_PT_LITERAL, ".1"},
	{"1.0e1", token.FLOATING_PT_LITERAL, "1.0e1"},
	{"1.0e-2", token.FLOATING_PT_LITERAL, "1.0e-2"},
	{"1.0e+3", token.FLOATING_PT_LITERAL, "1.0e+3"},
	{"2.14E+3", token.FLOATING_PT_LITERAL, "2.14E+3"},
	{"3.56e", token.ILLEGAL, "Syntax Error: Missing exponent"},
	{".0e1", token.FLOATING_PT_LITERAL, ".0e1"},
	{".0e-2", token.FLOATING_PT_LITERAL, ".0e-2"},
	{".0e+3", token.FLOATING_PT_LITERAL, ".0e+3"},
	{".14E+3", token.FLOATING_PT_LITERAL, ".14E+3"},
	{"2e2", token.FLOATING_PT_LITERAL, "2e2"},
	{"3.e4", token.FLOATING_PT_LITERAL, "3.e4"},
	{".56e", token.ILLEGAL, "Syntax Error: Missing exponent"},
	{"e2", token.IDENTIFIER, "e2"},

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
