package parser

import (
	"github.com/manhtai/goin/ast"
	"github.com/manhtai/goin/lexer"
	"github.com/manhtai/goin/token"
)

// Parser contains a Lexer and two pointer for traveling Lexer tokens
type Parser struct {
	l         *lexer.Lexer
	curToken  token.Token
	peekToken token.Token
}

// New construct Parser from Lexer
func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}
	// Read two tokens, so curToken and peekToken are both set
	p.nextToken()
	p.nextToken()
	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

// ParseProgram will parse a Program and return a Parser
func (p *Parser) ParseProgram() *ast.Program {
	return nil
}
