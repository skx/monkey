package token

// Type is a string
type Type string

// Token struct represent the lexer token
type Token struct {
	Type    Type
	Literal string
}

// pre-defined Type
const (
	EOF             = "EOF"
	IDENT           = "IDENT"
	INT             = "INT"
	FLOAT           = "FLOAT"
	ASSIGN          = "="
	PLUS            = "+"
	PLUS_PLUS       = "++"
	PLUS_EQUALS     = "+="
	AND             = "&&"
	OR              = "||"
	MOD             = "%"
	COMMA           = ","
	BACKTICK        = "`"
	SEMICOLON       = ";"
	MINUS           = "-"
	MINUS_MINUS     = "--"
	MINUS_EQUALS    = "-="
	BANG            = "!"
	ASTERISK        = "*"
	ASTERISK_EQUALS = "*="
	POW             = "**"
	SLASH           = "/"
	SLASH_EQUALS    = "/="
	LT              = "<"
	LT_EQUALS       = "<="
	GT              = ">"
	GT_EQUALS       = ">="
	LPAREN          = "("
	RPAREN          = ")"
	LBRACE          = "{"
	RBRACE          = "}"
	FUNCTION        = "FUNCTION"
	DEFINE_FUNCTION = "DEFINE_FUNCTION"
	LET             = "LET"
	CONST           = "CONST"
	TRUE            = "TRUE"
	FALSE           = "FALSE"
	IF              = "IF"
	ELSE            = "ELSE"
	RETURN          = "RETURN"
	FOR             = "FOR"
	EQ              = "=="
	NOT_EQ          = "!="
	STRING          = "STRING"
	REGEXP          = "REGEXP"
	LBRACKET        = "["
	RBRACKET        = "]"
	COLON           = ":"
	PERIOD          = "."
)

// reversed keywords
var keywords = map[string]Type{
	"const":    CONST,
	"else":     ELSE,
	"false":    FALSE,
	"fn":       FUNCTION,
	"for":      FOR,
	"function": DEFINE_FUNCTION,
	"if":       IF,
	"let":      LET,
	"return":   RETURN,
	"true":     TRUE,
}

// LookupIdentifier used to determinate whether identifier is keyword nor not
func LookupIdentifier(identifier string) Type {
	if tok, ok := keywords[identifier]; ok {
		return tok
	}
	return IDENT
}
