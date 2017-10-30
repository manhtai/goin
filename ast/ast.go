package ast

import "github.com/manhtai/goin/token"

// Node represent every node in our AST
type Node interface {
	TokenLiteral() string
}

// Statement represent a statement, e.g. let x = 5
type Statement interface {
	Node
	statementNode()
}

// Expression represent an expression, e.g. add(5, 7)
type Expression interface {
	Node
	expressionNode()
}

// Program is root of all Node, contains many statements
type Program struct {
	Statements []Statement
}

// TokenLiteral is used for debugging and testing only
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}

// LetStatement is for let, e.g. let x = 5 + 5
type LetStatement struct {
	Token token.Token // the token.LET token
	Name  *Identifier
	Value Expression
}

// LetStatement implement Statement interface
func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

// Identifier is for token.INDENT token to hold identifier, e.g. x
type Identifier struct {
	Token token.Token // the token.IDENT token Value string
	Value string
}

// Identifier implement Expression interface
func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

// ReturnStatement is struct for return statement
type ReturnStatement struct {
	Token       token.Token
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode()       {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }
