package parser

import (
	"fmt"
	"strconv"

	"github.com/JunNishimura/jsop/ast"
	"github.com/JunNishimura/jsop/lexer"
	"github.com/JunNishimura/jsop/token"
)

type Parser struct {
	l *lexer.Lexer

	curToken  token.Token
	peekToken token.Token

	errors []string
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l:      l,
		errors: []string{},
	}

	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{
		Expressions: []ast.Expression{},
	}

	for !p.curTokenIs(token.EOF) {
		exp := p.parseExpression()
		if exp != nil {
			program.Expressions = append(program.Expressions, exp)
		}
		p.nextToken()
	}

	return program
}

func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}

func (p *Parser) parseExpression() ast.Expression {
	if p.curTokenIs(token.LBRACKET) {
		// return p.parseObject()
	}
	return p.parseAtom()
}

func (p *Parser) parseAtom() ast.Expression {
	switch p.curToken.Type {
	case token.MINUS:
		return p.parsePrefixAtom()
	case token.INT:
		return p.parseIntegerLiteral()
	default:
		return nil
	}
}

func (p *Parser) parsePrefixAtom() *ast.PrefixAtom {
	pa := &ast.PrefixAtom{
		Token:    p.curToken,
		Operator: p.curToken.Literal,
	}

	p.nextToken()

	pa.Right = p.parseAtom()

	return pa
}

func (p *Parser) parseIntegerLiteral() *ast.IntegerLiteral {
	intValue, err := strconv.ParseInt(p.curToken.Literal, 0, 64)
	if err != nil {
		msg := fmt.Sprintf("could not parse %q as integer", p.curToken.Literal)
		p.errors = append(p.errors, msg)
		return nil
	}

	return &ast.IntegerLiteral{Token: p.curToken, Value: intValue}
}
