package parser

import (
	"errors"
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

	errors error
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l: l,
	}

	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) Errors() []string {
	out := []string{}

	if errs, ok := p.errors.(interface{ Unwrap() []error }); ok {
		for _, err := range errs.Unwrap() {
			out = append(out, err.Error())
		}
	} else {
		out = append(out, p.errors.Error())
	}

	return out
}

func (p *Parser) handleErr(err error) {
	p.errors = errors.Join(p.errors, err)
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() (*ast.Program, error) {
	program := &ast.Program{
		Expressions: []ast.Expression{},
	}

	for !p.curTokenIs(token.EOF) {
		exp, err := p.parseExpression()
		if err != nil {
			// run through the rest of the program to find all errors
			p.handleErr(err)
		}

		program.Expressions = append(program.Expressions, exp)
	}

	// return all errors found after parsing the whole program
	if p.errors != nil {
		return nil, p.errors
	}

	return program, nil
}

func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}

func (p *Parser) expectCurToken(t token.TokenType) error {
	if p.curTokenIs(t) {
		p.nextToken()
		return nil
	}

	return fmt.Errorf("expected current token to be %s, got %s instead", t, p.curToken.Type)
}

func (p *Parser) expectTokens(tokens ...token.TokenType) error {
	for _, t := range tokens {
		if !p.curTokenIs(t) {
			return fmt.Errorf("expected current token to be %s, got %s instead", t, p.curToken.Type)
		}
		p.nextToken()
	}

	return nil
}

func (p *Parser) parseExpression() (ast.Expression, error) {
	defer p.nextToken()

	if p.curTokenIs(token.LBRACE) {
		return p.parseObject()
	}
	return p.parseAtom()
}

func (p *Parser) parseObject() (ast.Expression, error) {
	if err := p.expectCurToken(token.LBRACE); err != nil {
		return nil, err
	}

	// parse main key
	if err := p.expectCurToken(token.DOUBLE_QUOTE); err != nil {
		return nil, err
	}
	switch p.curToken.Type {
	case token.COMMAND:
		return p.parseCommand()
	case token.IF:
		return p.parseIfExpression()
	case token.SET:
		return p.parseSetExpression()
	default:
		return nil, fmt.Errorf("unexpected token type %s", p.curToken.Type)
	}
}

func (p *Parser) parseCommand() (*ast.CommandObject, error) {
	commandToken := p.curToken

	// skip to symbol
	if err := p.expectTokens(
		token.COMMAND,
		token.DOUBLE_QUOTE,
		token.COLON,
		token.LBRACE,
		token.DOUBLE_QUOTE,
		token.SYMBOLKEY,
		token.DOUBLE_QUOTE,
		token.COLON,
		token.DOUBLE_QUOTE,
	); err != nil {
		return nil, err
	}

	// parse symbol
	symbol, err := p.parseExpression()
	if err != nil {
		return nil, err
	}

	// skip to args
	if err := p.expectTokens(
		token.DOUBLE_QUOTE,
		token.COMMA,
		token.DOUBLE_QUOTE,
		token.ARGS,
		token.DOUBLE_QUOTE,
		token.COLON,
	); err != nil {
		return nil, err
	}

	// parse args
	args, err := p.parseArgs()
	if err != nil {
		return nil, err
	}

	// skip to end of object
	if err := p.expectTokens(token.RBRACE); err != nil {
		return nil, err
	}

	return &ast.CommandObject{
		Token:  commandToken,
		Symbol: symbol,
		Args:   args,
	}, nil
}

func (p *Parser) parseArgs() ([]ast.Expression, error) {
	args := []ast.Expression{}

	if err := p.expectCurToken(token.LBRACKET); err != nil {
		return nil, err
	}

	for !p.curTokenIs(token.RBRACKET) {
		arg, err := p.parseExpression()
		if err != nil {
			return nil, err
		}

		args = append(args, arg)

		if p.curTokenIs(token.COMMA) {
			p.nextToken()
		}
	}

	if err := p.expectCurToken(token.RBRACKET); err != nil {
		return nil, err
	}

	return args, nil
}

func (p *Parser) parseIfExpression() (*ast.IfExpression, error) {
	ifToken := p.curToken

	// skip to condition
	if err := p.expectTokens(
		token.IF,
		token.DOUBLE_QUOTE,
		token.COLON,
		token.LBRACE,
		token.DOUBLE_QUOTE,
		token.COND,
		token.DOUBLE_QUOTE,
		token.COLON,
	); err != nil {
		return nil, err
	}

	// parse condition
	condition, err := p.parseExpression()
	if err != nil {
		return nil, err
	}

	// skip to consequence
	if err := p.expectTokens(
		token.COMMA,
		token.DOUBLE_QUOTE,
		token.CONSEQ,
		token.DOUBLE_QUOTE,
		token.COLON,
	); err != nil {
		return nil, err
	}

	// parse consequence
	consequence, err := p.parseExpression()
	if err != nil {
		return nil, err
	}

	// parse alternative
	if p.curTokenIs(token.COMMA) {
		if err := p.expectTokens(
			token.COMMA,
			token.DOUBLE_QUOTE,
			token.ALT,
			token.DOUBLE_QUOTE,
			token.COLON); err != nil {
			return nil, err
		}

		alternative, err := p.parseExpression()
		if err != nil {
			return nil, err
		}

		if err := p.expectTokens(token.RBRACE, token.RBRACE); err != nil {
			return nil, err
		}

		return &ast.IfExpression{
			Token:       ifToken,
			Condition:   condition,
			Consequence: consequence,
			Alternative: alternative,
		}, nil
	}

	if err := p.expectTokens(token.RBRACE, token.RBRACE); err != nil {
		return nil, err
	}

	return &ast.IfExpression{
		Token:       ifToken,
		Condition:   condition,
		Consequence: consequence,
		Alternative: nil,
	}, nil
}

func (p *Parser) parseSetExpression() (*ast.SetExpression, error) {
	setToken := p.curToken

	// skip to var
	if err := p.expectTokens(
		token.SET,
		token.DOUBLE_QUOTE,
		token.COLON,
		token.LBRACE,
		token.DOUBLE_QUOTE,
		token.VAR,
		token.DOUBLE_QUOTE,
		token.COLON,
		token.DOUBLE_QUOTE,
		token.DOLLAR,
	); err != nil {
		return nil, err
	}

	// parse var
	variable, err := p.parseExpression()
	if err != nil {
		return nil, err
	}
	varName, ok := variable.(*ast.Symbol)
	if !ok {
		return nil, fmt.Errorf("expected symbol, got %T", variable)
	}

	// skip to val
	if err := p.expectTokens(
		token.DOUBLE_QUOTE,
		token.COMMA,
		token.DOUBLE_QUOTE,
		token.VAL,
		token.DOUBLE_QUOTE,
		token.COLON,
	); err != nil {
		return nil, err
	}

	// parse val
	value, err := p.parseExpression()
	if err != nil {
		return nil, err
	}

	if err := p.expectTokens(token.RBRACE, token.RBRACE); err != nil {
		return nil, err
	}

	return &ast.SetExpression{
		Token: setToken,
		Name:  varName,
		Value: value,
	}, nil
}

func (p *Parser) parseAtom() (ast.Expression, error) {
	switch p.curToken.Type {
	case token.MINUS, token.EXCLAM:
		return p.parsePrefixAtom()
	case token.INT:
		return p.parseIntegerLiteral()
	case token.TRUE:
		return &ast.Boolean{Token: p.curToken, Value: true}, nil
	case token.FALSE:
		return &ast.Boolean{Token: p.curToken, Value: false}, nil
	case token.SYMBOL:
		return &ast.Symbol{Token: p.curToken, Value: p.curToken.Literal}, nil
	default:
		return nil, fmt.Errorf("unexpected token type %s", p.curToken.Type)
	}
}

func (p *Parser) parsePrefixAtom() (*ast.PrefixAtom, error) {
	pa := &ast.PrefixAtom{
		Token:    p.curToken,
		Operator: p.curToken.Literal,
	}

	p.nextToken()

	right, err := p.parseAtom()
	if err != nil {
		return nil, err
	}
	pa.Right = right

	return pa, nil
}

func (p *Parser) parseIntegerLiteral() (*ast.IntegerLiteral, error) {
	intValue, err := strconv.ParseInt(p.curToken.Literal, 0, 64)
	if err != nil {
		return nil, fmt.Errorf("could not parse %q as integer", p.curToken.Literal)
	}

	return &ast.IntegerLiteral{Token: p.curToken, Value: intValue}, nil
}
