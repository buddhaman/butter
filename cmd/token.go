package main

type TokenType int

const (
	ILLEGAL TokenType = iota
	EOF
	IDENTIFIER
	INT
	ASSIGN
	BANG
	SLASH
	ASTERISK
	LT
	GT
	PLUS
	COMMA
	SEMICOLON
	LPAREN
	RPAREN
	LBRACE
	RBRACE
	FUNCTION
	LET
	TRUE
	FALSE
	IF
	ELSE
	RETURN
)

var tokenTypeNames = map[TokenType]string{
	ILLEGAL:    "ILLEGAL",
	EOF:        "EOF",
	IDENTIFIER: "IDENTIFIER",
	INT:        "INT",
	ASSIGN:     "ASSIGN",
	BANG:       "BANG",
	SLASH:      "SLASH",
	ASTERISK:   "ASTERISK",
	LT:         "LT",
	GT:         "GT",
	PLUS:       "PLUS",
	COMMA:      "COMMA",
	SEMICOLON:  "SEMICOLON",
	LPAREN:     "LPAREN",
	RPAREN:     "RPAREN",
	LBRACE:     "LBRACE",
	RBRACE:     "RBRACE",
	FUNCTION:   "FUNCTION",
	LET:        "LET",
	TRUE:       "TRUE",
	FALSE:      "FALSE",
	IF:         "IF",
	ELSE:       "ELSE",
	RETURN:     "RETURN",
}

func TokStr(t TokenType) string {
	if name, ok := tokenTypeNames[t]; ok {
		return name
	} else {
		return "Unknown token"
	}
}

var keywords = map[string]TokenType{
	"fn":    FUNCTION,
	"let":   LET,
	"false": TRUE,
	"true":  FALSE,
	"if":    IF,
	"else":  ELSE,
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
