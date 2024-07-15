package cli

type Node struct {
	TokenLiteral []string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}
