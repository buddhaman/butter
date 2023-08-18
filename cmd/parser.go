package main

import (
	"fmt"
)

const (
	_ int = iota
	LOWEST
	EQUALS
	LESSGREATER
	SUM
	PRODUCT
	PREFIX
	CALL
)

type (
	prefixParserFn func() Expression
	infixParserFn  func(Expression) Expression
)

type Parser struct {
	l *Lexer

	curToken  Token
	peekToken Token

	errors []string

	prefixParseFns map[TokenType]prefixParserFn
	infixParserFns map[TokenType]infixParserFn
}

func (p *Parser) registerPrefix(tokenType TokenType, fn prefixParserFn) {
	p.prefixParseFns[tokenType] = fn
}

func (p *Parser) registerInfix(tokenType TokenType, fn infixParserFn) {
	p.infixParserFns[tokenType] = fn
}

func NewParser(l *Lexer) *Parser {
	p := &Parser{
		l:      l,
		errors: []string{},
	}
	p.NextToken()
	p.NextToken()

	p.prefixParseFns = make(map[TokenType]prefixParserFn)
	p.registerPrefix(IDENTIFIER, p.parseIdentifier)

	return p
}

func (p *Parser) parseIdentifier() Expression {
	return &Identifier{Token: p.curToken, Value: p.curToken.Literal}
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) peekError(t TokenType) {
	msg := fmt.Sprintf("Expected next token to be %s, got %s (%s) instead.",
		TokStr(t), TokStr(p.peekToken.Type), p.peekToken.Literal)
	p.errors = append(p.errors, msg)
}

func (p *Parser) NextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) peekTokenIs(t TokenType) bool {
	return p.peekToken.Type == t
}

func (p *Parser) curTokenIs(t TokenType) bool {
	return p.curToken.Type == t
}

func (p *Parser) expectPeek(t TokenType) bool {
	if p.peekTokenIs(t) {
		p.NextToken()
		return true
	} else {
		p.peekError(t)
		return false
	}
}

func (p *Parser) parseLetStatement() *LetStatement {
	stmt := &LetStatement{Token: p.curToken}

	if !p.expectPeek(IDENTIFIER) {
		return nil
	}

	stmt.Name = &Identifier{Token: p.curToken, Value: p.curToken.Literal}

	if !p.expectPeek(ASSIGN) {
		return nil
	}

	for !p.curTokenIs(SEMICOLON) {
		p.NextToken()
	}

	return stmt
}

func (p *Parser) parseReturnStatement() *ReturnStatement {
	stmt := &ReturnStatement{Token: p.curToken}

	p.NextToken()

	for !p.curTokenIs(SEMICOLON) {
		p.NextToken()
	}

	return stmt
}

func (p *Parser) parseExpressionStatement() *ExpressionStatement {
	stmt := &ExpressionStatement{Token: p.curToken}

	stmt.Expression = p.parseExpression(LOWEST)

	if p.peekTokenIs(SEMICOLON) {
		p.NextToken()
	}

	return stmt
}

func (p *Parser) parseExpression(precedence int) Expression {
	prefix := p.prefixParseFns[p.curToken.Type]
	if prefix == nil {
		return nil
	}
	leftExp := prefix()
	return leftExp
}

func (p *Parser) parseStatement() Statement {
	switch p.curToken.Type {
	case LET:
		return p.parseLetStatement()
	case RETURN:
		return p.parseReturnStatement()
	default:
		return p.parseExpressionStatement()
	}
}

func (p *Parser) ParseProgram() *Program {
	program := &Program{}
	program.Statements = []Statement{}

	for p.curToken.Type != EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.NextToken()
	}
	return program
}
