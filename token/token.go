package token

// TokenType string ...
type TokenType string

// Token ほげほげ
type Token struct {
	Type    TokenType
	Literal string
}

const (
	// ILLEGAL ...
	ILLEGAL = "ILLEGAL"
	// EOF ...
	EOF = "EOF"

	// 識別子 + リテラル

	// IDENT ...
	IDENT = "IDENT" // add, foobar, x, y,
	// INT ...
	INT = "INT"

	// ASSIGN ...
	ASSIGN = "="
	// PLUS ...
	PLUS = "+"

	// デリミタ

	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"

	LBRACE = "{"
	RBRACE = "}"

	// キーワード

	FUNCTION = "FUNCTION"
	LET      = "LET"
)

var keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
}

// LookupIdent ...
func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
