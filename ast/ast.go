package ast

import "github.com/huderlem/poryscript/token"

// Node is an interface that represents a node in a Poryscript AST.
type Node interface {
	TokenLiteral() string
}

// Statement is an interface that represents a statement node in a Poryscript AST.
type Statement interface {
	Node
	statementNode()
}

// Text holds a label and value for some script text.
type Text struct {
	Name  string
	Value string
}

// Program represents the root-level Node in any Poryscript AST.
type Program struct {
	TopLevelStatements []Statement
	Texts              []Text
}

// TokenLiteral returns a string representation of the Program node.
func (p *Program) TokenLiteral() string {
	if len(p.TopLevelStatements) > 0 {
		return p.TopLevelStatements[0].TokenLiteral()
	}
	return ""
}

// ScriptStatement is a Poryscript script statement. Script statements define
// the block of a script's execution.
type ScriptStatement struct {
	Token token.Token
	Name  *Identifier
	Body  *BlockStatement
}

func (ss *ScriptStatement) statementNode() {}

// TokenLiteral returns a string representation of the script statement.
func (ss *ScriptStatement) TokenLiteral() string { return ss.Token.Literal }

// BlockStatement is a Poryscript block, which can hold many statements and blocks inside.
// It is defined by curly braces.
type BlockStatement struct {
	Token      token.Token
	Statements []Statement
}

func (bs *BlockStatement) statementNode() {}

// TokenLiteral returns a string representation of the block statement.
func (bs *BlockStatement) TokenLiteral() string { return bs.Token.Literal }

// CommandStatement is a Poryscript command statement. Command statements map directly to
// original engine script commands.
type CommandStatement struct {
	Token token.Token
	Name  *Identifier
	Args  []string
}

func (cs *CommandStatement) statementNode() {}

// TokenLiteral returns a string representation of the command statement.
func (cs *CommandStatement) TokenLiteral() string { return cs.Token.Literal }

// Identifier represents a Poryscript identifier.
type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) expressionNode() {}

// TokenLiteral returns a string representation of the identifier.
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

// RawStatement is a Poryscript raw statement. Raw statements are directly
// included into the target bytecode script.
type RawStatement struct {
	Token    token.Token
	Name     *Identifier
	Value    string
	IsGlobal bool
}

func (rs *RawStatement) statementNode() {}

// TokenLiteral returns a string representation of the raw statement.
func (rs *RawStatement) TokenLiteral() string { return rs.Token.Literal }

// ConditionExpression is the condition for an if statement.
type ConditionExpression struct {
	Type            token.Type
	Operand         string
	Operator        token.Type
	ComparisonValue string
	Body            *BlockStatement
}

// IfStatement is an if statement in Poryscript.
type IfStatement struct {
	Token            token.Token
	Consequence      *ConditionExpression
	ElifConsequences []*ConditionExpression
	ElseConsequence  *BlockStatement
}

func (rs *IfStatement) statementNode() {}

// TokenLiteral returns a string representation of the if statement.
func (ie *IfStatement) TokenLiteral() string { return ie.Token.Literal }