package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// 7.2.3 Identifiers
	IDENTIFIER = "IDENTIFIER"

	// 7.2.4  Keywords
	ABSTRACT    = "abstract"
	ANY         = "any"
	ALIAS       = "alias"
	ATTRIBUTE   = "attribute"
	BITFIELD    = "bitfield"
	BITMASK     = "bitmask"
	BITSET      = "bitset"
	BOOLEAN     = "boolean"
	CASE        = "case"
	CHAR        = "char"
	COMPONENT   = "component"
	CONNECTOR   = "connector"
	CONST       = "const"
	CONSUMES    = "consumes"
	CONTEXT     = "context"
	CUSTOM      = "custom"
	DEFAULT     = "default"
	DOUBLE      = "double"
	EXCEPTION   = "exception"
	EMITS       = "emits"
	ENUM        = "enum"
	EVENTTYPE   = "eventtype"
	FACTORY     = "factory"
	FALSE       = "FALSE"
	FINDER      = "finder"
	FIXED       = "fixed"
	FLOAT       = "float"
	GETRAISES   = "getraises"
	HOME        = "home"
	IMPORT      = "import"
	IN          = "in"
	INOUT       = "inout"
	INTERFACE   = "interface"
	LOCAL       = "local"
	LONG        = "long"
	MANAGES     = "manages"
	MAP         = "map"
	MIRRORPORT  = "mirrorport"
	MODULE      = "module"
	MULTIPLE    = "multiple"
	NATIVE      = "native"
	OBJECT      = "object"
	OCTET       = "octet"
	ONEWAY      = "oneway"
	OUT         = "out"
	PRIMARYKEY  = "primarykey"
	PRIVATE     = "private"
	PORT        = "port"
	PORTTYPE    = "porttype"
	PROVIDES    = "provides"
	PUBLIC      = "public"
	PUBLISHES   = "publishes"
	RAISES      = "raises"
	READONLY    = "readonly"
	SETRAISES   = "setraises"
	SEQUENCE    = "sequence"
	SHORT       = "short"
	STRING      = "string"
	STRUCT      = "struct"
	SUPPORTS    = "supports"
	SWITCH      = "switch"
	TRUE        = "TRUE"
	TRUNCATABLE = "trucatable"
	TYPEDEF     = "typedef"
	TYPEID      = "typeid"
	TYPENAME    = "typename"
	TYPEPREFIX  = "typeprefix"
	UNSIGNED    = "unsigned"
	UNION       = "union"
	USES        = "uses"
	VALUEBASE   = "ValueBase"
	VALUETYPE   = "valuetype"
	VOID        = "void"
	WCHAR       = "wchar"
	WSTRING     = "wstring"
	INT8        = "int8"
	UINT8       = "uint8"
	INT16       = "int16"
	INT32       = "int32"
	INT64       = "int64"
	UINT16      = "uint16"
	UINT32      = "unit32"
	UINT64      = "unit64"

	// 7.2.6.1 Integer Literal
	INTEGER = "INTEGER"

	// 7.2.6.2 Character Literals
	CHAR_LITERAL  = "CHAR_LITERAL"
	WCHAR_LITERAL = "WCHAR_LITERAL"

	// Delimiters
	SEMICOLON = ";"
	LCURLY    = "{"
	RCURLY    = "}"
	COLON     = ":"
	COMMA     = ","
	LPAREN    = "("
	RPAREN    = ")"
	LT        = "<"
	GT        = ">"
	SCOPESEP  = "::"

	// Operators
	EQUALS   = "="
	PLUS     = "+"
	MINUS    = "-"
	LBRACKET = "["
	RBRACKET = "]"
	OR       = "|"
	XOR      = "^"
	AND      = "&"
	MUL      = "*"
	DIV      = "/"
	MODULO   = "%"
	NOT      = "~"
	LSHIFT   = "<<"
	RSHIFT   = ">>"
)

var Keywords = map[string]TokenType{
	"abstract":    ABSTRACT,
	"any":         ANY,
	"alias":       ALIAS,
	"attribute":   ATTRIBUTE,
	"bitfield":    BITFIELD,
	"bitmask":     BITMASK,
	"bitset":      BITSET,
	"boolean":     BOOLEAN,
	"case":        CASE,
	"char":        CHAR,
	"component":   COMPONENT,
	"connector":   CONNECTOR,
	"const":       CONST,
	"consumes":    CONSUMES,
	"context":     CONTEXT,
	"custom":      CUSTOM,
	"default":     DEFAULT,
	"double":      DOUBLE,
	"exception":   EXCEPTION,
	"emits":       EMITS,
	"enum":        ENUM,
	"eventtype":   EVENTTYPE,
	"factory":     FACTORY,
	"FALSE":       FALSE,
	"finder":      FINDER,
	"fixed":       FIXED,
	"float":       FLOAT,
	"getraises":   GETRAISES,
	"home":        HOME,
	"import":      IMPORT,
	"in":          IN,
	"inout":       INOUT,
	"interface":   INTERFACE,
	"local":       LOCAL,
	"long":        LONG,
	"manages":     MANAGES,
	"map":         MAP,
	"mirrorport":  MIRRORPORT,
	"module":      MODULE,
	"multiple":    MULTIPLE,
	"native":      NATIVE,
	"Object":      OBJECT,
	"octet":       OCTET,
	"oneway":      ONEWAY,
	"out":         OUT,
	"primarykey":  PRIMARYKEY,
	"private":     PRIVATE,
	"port":        PORT,
	"porttype":    PORTTYPE,
	"provides":    PROVIDES,
	"public":      PUBLIC,
	"publishes":   PUBLISHES,
	"raises":      RAISES,
	"readonly":    READONLY,
	"setraises":   SETRAISES,
	"sequence":    SEQUENCE,
	"short":       SHORT,
	"string":      STRING,
	"struct":      STRUCT,
	"supports":    SUPPORTS,
	"switch":      SWITCH,
	"TRUE":        TRUE,
	"truncatable": TRUNCATABLE,
	"typedef":     TYPEDEF,
	"typeid":      TYPEID,
	"typename":    TYPENAME,
	"typeprefix":  TYPEPREFIX,
	"unsigned":    UNSIGNED,
	"union":       UNION,
	"uses":        USES,
	"ValueBase":   VALUEBASE,
	"valuetype":   VALUETYPE,
	"void":        VOID,
	"wchar":       WCHAR,
	"wstring":     WSTRING,
	"int8":        INT8,
	"uint8":       UINT8,
	"int16":       INT16,
	"int32":       INT32,
	"int64":       INT64,
	"uint16":      UINT16,
	"uint32":      UINT32,
	"uint64":      UINT64,
}
