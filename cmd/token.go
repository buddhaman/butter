package main

type TokenType int

const (
	ILLEGAL TokenType = iota
	EOF
	IDENTIFIER
	INT
	ASSIGN
	PLUS
	COMMA
	SEMICOLON
	LPAREN
	RPAREN
	LBRACE
	RBRACE
	FUNCTION
	LET
)

var tokenTypeNames = map[TokenType]string{
	ILLEGAL:    "ILLEGAL",
	EOF:        "EOF",
	IDENTIFIER: "IDENTIFIER",
	INT:        "INT",
	ASSIGN:     "ASSIGN",
	PLUS:       "PLUS",
	COMMA:      "COMMA",
	SEMICOLON:  "SEMICOLON",
	LPAREN:     "LPAREN",
	RPAREN:     "RPAREN",
	LBRACE:     "LBRACE",
	RBRACE:     "RBRACE",
	FUNCTION:   "FUNCTION",
	LET:        "LET",
}

func TokStr(t TokenType) string {
	return tokenTypeNames[t]
}

var keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
}

func LookupIdentifier(identifier string) TokenType {
	if tok, ok := keywords[identifier]; ok {
		return tok
	}
	return IDENTIFIER
}

type Token struct {
	Type    TokenType
	Literal string
}
