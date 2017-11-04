package ast

import (
	"bytes"

	"github.com/manhtai/goin/token"
)

///////////////////////////////////////////////////////////////////////////////
// Node, Statement & Expression interfaces
///////////////////////////////////////////////////////////////////////////////

// Node represent every node in our AST
type Node interface {
	TokenLiteral() string
	String() string
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

///////////////////////////////////////////////////////////////////////////////
// PROGRAM
///////////////////////////////////////////////////////////////////////////////

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

func (p *Program) String() string {
	var out bytes.Buffer
	for _, s := range p.Statements {
		out.WriteString(s.String())
	}
	return out.String()
}

///////////////////////////////////////////////////////////////////////////////
// LET
///////////////////////////////////////////////////////////////////////////////

// LetStatement is for let, e.g. let x = 5 + 5
type LetStatement struct {
	Token token.Token // the token.LET token
	Name  *Identifier
	Value Expression
}

// LetStatement implement Statement interface
func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }
func (ls *LetStatement) String() string {
	var out bytes.Buffer
	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")
	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}

	out.WriteString(";")
	return out.String()
}

///////////////////////////////////////////////////////////////////////////////
// IDENTIFIER
///////////////////////////////////////////////////////////////////////////////

// Identifier is for token.INDENT token to hold identifier, e.g. x
type Identifier struct {
	Token token.Token // the token.IDENT token Value string
	Value string
}

// Identifier implement Expression interface
func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
func (i *Identifier) String() string       { return i.Value }

///////////////////////////////////////////////////////////////////////////////
// RETURN
///////////////////////////////////////////////////////////////////////////////

// ReturnStatement is struct for return statement
type ReturnStatement struct {
	Token       token.Token
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode()       {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }
func (rs *ReturnStatement) String() string {
	var out bytes.Buffer
	out.WriteString(rs.TokenLiteral() + " ")
	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}
	out.WriteString(";")
	return out.String()
}

///////////////////////////////////////////////////////////////////////////////
// EXPRESSION
///////////////////////////////////////////////////////////////////////////////

// ExpressionStatement is a statement that contains only 1 expression
type ExpressionStatement struct {
	Token      token.Token // the first token of the expression
	Expression Expression
}

func (es *ExpressionStatement) statementNode()       {}
func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal }
func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}

///////////////////////////////////////////////////////////////////////////////
// PREFIX EXPRESSION
///////////////////////////////////////////////////////////////////////////////

type PrefixExpression struct {
	Token    token.Token
	Operator string
	Right    Expression
}

func (pe *PrefixExpression) expressionNode() {}

func (pe *PrefixExpression) TokenLiteral() string {
	return pe.Token.Literal
}

func (pe *PrefixExpression) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(pe.Operator)
	out.WriteString(pe.Right.String())
	out.WriteString(")")
	return out.String()
}

///////////////////////////////////////////////////////////////////////////////
// INTEGER
///////////////////////////////////////////////////////////////////////////////

type IntegerLiteral struct {
	Token token.Token
	Value int64
}

func (il *IntegerLiteral) expressionNode() {}
func (il *IntegerLiteral) TokenLiteral() string {
	return il.Token.Literal
}
func (il *IntegerLiteral) String() string {
	return il.Token.Literal
}
