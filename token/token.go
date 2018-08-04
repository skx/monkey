package token

// TokenType is a string
type TokenType string

// Token struct represent the lexer token
type Token struct {
	Type    TokenType
	Literal string
}

// pre-defined TokenType
const (
	ILLEGAL         = "ILLEGAL"
	EOF             = "EOF"
	IDENT           = "IDENT"
	INT             = "INT"
	FLOAT           = "FLOAT"
	ASSIGN          = "="
	PLUS            = "+"
	PLUS_PLUS       = "++"
	PLUS_EQUALS     = "+="
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
	TRUE            = "TRUE"
	FALSE           = "FALSE"
	IF              = "IF"
	ELSE            = "ELSE"
	RETURN          = "RETURN"
	FOR             = "FOR"
	EQ              = "=="
	NOT_EQ          = "!="
	STRING          = "STRING"
	LBRACKET        = "["
	RBRACKET        = "]"
	COLON           = ":"
)

// reversed keywords
var keywords = map[string]TokenType{
	"fn":       FUNCTION,
	"function": DEFINE_FUNCTION,
	"let":      LET,
	"true":     TRUE,
	"false":    FALSE,
	"if":       IF,
	"else":     ELSE,
	"return":   RETURN,
	"for":      FOR,
}

// LookupIdentifier used to determinate whether identifier is keyword nor not
func LookupIdentifier(identifier string) TokenType {
	if tok, ok := keywords[identifier]; ok {
		return tok
	}
	return IDENT
}
