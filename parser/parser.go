package parser

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/JunNishimura/jsop/ast"
	"github.com/JunNishimura/jsop/lexer"
	"github.com/JunNishimura/jsop/token"
)

type Parser struct {
	l *lexer.Lexer

	curToken  token.Token
	peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l: l,
	}

	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() (ast.Expression, error) {
	if p.curTokenIs(token.EOF) {
		return nil, nil
	}

	exp, err := p.parseExpression()
	if err != nil {
		return nil, err
	}

	if err := p.expectCurToken(token.EOF); err != nil {
		return nil, err
	}

	return exp, nil
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

func (p *Parser) expectQuotedToken(t token.TokenType) (token.Token, error) {
	if err := p.expectCurToken(token.DOUBLE_QUOTE); err != nil {
		return token.Token{}, err
	}

	if !p.curTokenIs(t) {
		return token.Token{}, fmt.Errorf("expected %s, got %s instead", t, p.curToken.Type)
	}
	ret := p.curToken
	p.nextToken()

	if err := p.expectCurToken(token.DOUBLE_QUOTE); err != nil {
		return token.Token{}, err
	}

	return ret, nil
}

func (p *Parser) parseExpression() (ast.Expression, error) {
	if p.curTokenIs(token.LBRACE) {
		return p.parseObject()
	}
	return p.parseAtom()
}

func (p *Parser) parseObject() (ast.Object, error) {
	if !p.curTokenIs(token.LBRACE) {
		return nil, fmt.Errorf("expected LBRACE, got %s instead", p.curToken.Type)
	}
	object := &ast.KeyValueObject{
		Token: p.curToken,
		KV:    []*ast.KeyValuePair{},
	}
	p.nextToken()

	if p.curTokenIs(token.RBRACE) {
		p.nextToken()
		return object, nil
	}

	// parse key value pairs
	for {
		kvPair, err := p.parseKeyValuePair()
		if err != nil {
			return nil, err
		}
		object.KV = append(object.KV, kvPair)

		if !p.curTokenIs(token.COMMA) {
			break
		}
		p.nextToken()
	}

	if err := p.expectCurToken(token.RBRACE); err != nil {
		return nil, err
	}

	return object, nil
}

func (p *Parser) parseKeyValuePair() (*ast.KeyValuePair, error) {
	key, err := p.parseKey()
	if err != nil {
		return nil, err
	}

	if err := p.expectCurToken(token.COLON); err != nil {
		return nil, err
	}

	value, err := p.parseExpression()
	if err != nil {
		return nil, err
	}

	return &ast.KeyValuePair{Key: key, Value: value}, nil
}

func (p *Parser) parseKey() (*ast.StringLiteral, error) {
	keyToken, err := p.expectQuotedToken(token.STRING)
	if err != nil {
		return nil, err
	}

	return &ast.StringLiteral{
		Token: keyToken,
		Value: strings.ToLower(keyToken.Literal),
	}, nil
}

func (p *Parser) parseAtom() (ast.Expression, error) {
	switch p.curToken.Type {
	case token.MINUS:
		return p.parsePrefixAtom()
	case token.INT:
		return p.parseIntegerLiteral()
	case token.TRUE, token.FALSE:
		return p.parseBoolean()
	case token.DOUBLE_QUOTE:
		return p.parseDoubleQuotedString()
	case token.LBRACKET:
		return p.parseArray()
	default:
		err := fmt.Errorf("unexpected token type %s", p.curToken.Type)
		p.nextToken()
		return nil, err
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
	if !p.curTokenIs(token.INT) {
		return nil, fmt.Errorf("expected integer, got %s instead", p.curToken.Type)
	}

	intValue, err := strconv.ParseInt(p.curToken.Literal, 0, 64)
	if err != nil {
		return nil, fmt.Errorf("could not parse %q as integer", p.curToken.Literal)
	}

	result := &ast.IntegerLiteral{Token: p.curToken, Value: intValue}

	p.nextToken()

	return result, nil
}

func (p *Parser) parseBoolean() (*ast.Boolean, error) {
	if !p.curTokenIs(token.TRUE) && !p.curTokenIs(token.FALSE) {
		return nil, fmt.Errorf("expected boolean, got %s instead", p.curToken.Type)
	}

	var result *ast.Boolean

	switch p.curToken.Type {
	case token.TRUE:
		result = &ast.Boolean{Token: p.curToken, Value: true}
	case token.FALSE:
		result = &ast.Boolean{Token: p.curToken, Value: false}
	}

	p.nextToken()

	return result, nil
}

func (p *Parser) parseDoubleQuotedString() (ast.Expression, error) {
	if err := p.expectCurToken(token.DOUBLE_QUOTE); err != nil {
		return nil, err
	}

	strLitVal := strings.TrimSpace(p.curToken.Literal)
	res := &ast.StringLiteral{Token: p.curToken, Value: strLitVal}
	p.nextToken()

	if err := p.expectCurToken(token.DOUBLE_QUOTE); err != nil {
		return nil, err
	}

	return res, nil
}

func (p *Parser) parseArray() (*ast.Array, error) {
	if !p.curTokenIs(token.LBRACKET) {
		return nil, fmt.Errorf("expected LBRACKET, got %s instead", p.curToken.Type)
	}
	array := &ast.Array{
		Token:    p.curToken,
		Elements: []ast.Expression{},
	}
	p.nextToken()

	// empty array
	if p.curTokenIs(token.RBRACKET) {
		p.nextToken()
		return array, nil
	}

	for {
		element, err := p.parseExpression()
		if err != nil {
			return nil, err
		}
		array.Elements = append(array.Elements, element)

		if !p.curTokenIs(token.COMMA) {
			break
		}
		p.nextToken()
	}

	if err := p.expectCurToken(token.RBRACKET); err != nil {
		return nil, err
	}

	return array, nil
}
