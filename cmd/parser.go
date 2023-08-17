package main

type Parser struct {
	l *Lexer

	curToken  Token
	peekToken Token
}

func NewParser(l *Lexer) *Parser {
	p := &Parser{l: l}
	p.NextToken()
	p.NextToken()
	return p
}

func (p *Parser) NextToken() {
	p.peekToken = p.curToken
	p.curToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *Program {
	return nil
}
