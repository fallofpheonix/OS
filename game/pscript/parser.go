package pscript

import (
	"fmt"
	"strconv"

	phxmath "github.com/fallofpheonix/phoenix/foundation/math"
	"github.com/fallofpheonix/phoenix/game/engine"
)

type Parser struct {
	l         *Lexer
	curToken  Token
	peekToken Token

	variables map[string]interface{}
}

func NewParser(l *Lexer) *Parser {
	p := &Parser{
		l:         l,
		variables: make(map[string]interface{}),
	}

	// Read two tokens, so curToken and peekToken are both set
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) Parse() ([]engine.Instruction, error) {
	instructions := []engine.Instruction{}

	for p.curToken.Type != TokenEOF {
		instr, err := p.parseStatement()
		if err != nil {
			return nil, err
		}
		if instr != nil {
			instructions = append(instructions, *instr)
		}
		p.nextToken()
	}

	return instructions, nil
}

func (p *Parser) parseStatement() (*engine.Instruction, error) {
	switch p.curToken.Type {
	case TokenIdent:
		if p.peekToken.Type == TokenAssign {
			return p.parseAssignment()
		}
		return p.parseFunctionCall()
	default:
		// Simplified for now, real parsing needs more structure
		if p.curToken.Literal == "move" {
			return p.parseMove()
		} else if p.curToken.Literal == "verify" {
			return p.parseVerify()
		}
		return nil, nil
	}
}

func (p *Parser) parseAssignment() (*engine.Instruction, error) {
	varName := p.curToken.Literal
	p.nextToken() // IDENT
	p.nextToken() // ASSIGN

	val, err := p.parseExpression()
	if err != nil {
		return nil, err
	}
	p.variables[varName] = val
	return nil, nil
}

func (p *Parser) parseExpression() (interface{}, error) {
	switch p.curToken.Type {
	case TokenNumber:
		v, _ := strconv.ParseInt(p.curToken.Literal, 10, 64)
		return phxmath.NewFixedPoint(v), nil
	case TokenString:
		return p.curToken.Literal, nil
	case TokenIdent:
		val, ok := p.variables[p.curToken.Literal]
		if !ok {
			return nil, fmt.Errorf("undefined variable: %s", p.curToken.Literal)
		}
		return val, nil
	default:
		return nil, fmt.Errorf("invalid expression: %s", p.curToken.Literal)
	}
}

func (p *Parser) parseMove() (*engine.Instruction, error) {
	p.nextToken() // MOVE
	if p.curToken.Type != TokenLParen {
		return nil, fmt.Errorf("expected (")
	}
	p.nextToken()

	entityID, err := p.parseExpression()
	if err != nil {
		return nil, err
	}

	if p.curToken.Type != TokenComma {
		return nil, fmt.Errorf("expected ,")
	}
	p.nextToken()

	pos, err := p.parseExpression()
	if err != nil {
		return nil, err
	}

	if p.curToken.Type != TokenRParen {
		return nil, fmt.Errorf("expected )")
	}

	return &engine.Instruction{
		Op:   engine.OpPush, // Placeholder, real compiler would push args and call
		Args: []interface{}{entityID, pos},
	}, nil
}

func (p *Parser) parseVerify() (*engine.Instruction, error) {
	p.nextToken() // VERIFY
	if p.curToken.Type != TokenLParen {
		return nil, fmt.Errorf("expected (")
	}
	p.nextToken()

	entityID, err := p.parseExpression()
	if err != nil {
		return nil, err
	}

	if p.curToken.Type != TokenRParen {
		return nil, fmt.Errorf("expected )")
	}

	return &engine.Instruction{
		Op:   engine.OpPush,
		Args: []interface{}{entityID},
	}, nil
}

func (p *Parser) parseFunctionCall() (*engine.Instruction, error) {
	return nil, nil
}
