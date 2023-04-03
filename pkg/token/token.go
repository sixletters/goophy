package token

// Can be optimized to use an int or a byte.
type TokenType string
type Token struct {
	Type    TokenType
	Literal string
}

func NewToken(tokenType TokenType, ch byte) Token {
	return Token{
		Type:    tokenType,
		Literal: string(ch),
	}
}

func NewTokenWithStr(tokenType TokenType, ch string) Token {
	return Token{
		Type:    tokenType,
		Literal: ch,
	}
}

const (
	ILLEGAL TokenType = "ILLEGAL" // Token/Char that we dont know about
	EOF     TokenType = "EOF"     // end of file

	// Identifiers + literals
	IDENT TokenType = "IDENT"
	INT   TokenType = "INT"
	GO    TokenType = "GO"

	// Operators
	ASSIGN   TokenType = "="
	PLUS     TokenType = "+"
	MINUS    TokenType = "-"
	BANG     TokenType = "!"
	ASTERISK TokenType = "*"
	SLASH    TokenType = "/"
	LT       TokenType = "<"
	GT       TokenType = ">"

	// Delimiters
	COMMA     TokenType = ","
	SEMICOLON TokenType = ";"

	LPAREN TokenType = "("
	RPAREN TokenType = ")"
	LBRACE TokenType = "{"
	RBRACE TokenType = "}"

	// Keywords
	FUNCTION TokenType = "FUNCTION"
	LET      TokenType = "LET"
	TRUE     TokenType = "TRUE"
	FALSE    TokenType = "FALSE"
	IF       TokenType = "IF"
	ELSE     TokenType = "ELSE"
	RETURN   TokenType = "RETURN"

	EQ     TokenType = "=="
	NOT_EQ TokenType = "!="
)

var Keywords = map[string]TokenType{
	"func":   FUNCTION,
	"var":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
	"go":     GO,
}

func LookUpIdentifier(ident string) TokenType {
	if tok, ok := Keywords[ident]; ok {
		return tok
	}
	return IDENT
}
