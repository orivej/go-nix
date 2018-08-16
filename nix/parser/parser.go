package parser

//go:generate goyacc nix.y
//go:generate sed "/yyS :=/a\\\tp := yylex.(*Parser)" -i y.go

import (
	"fmt"
	"strings"
)

type Node struct {
	Type   NodeType
	Tokens []int
	Nodes  []*Node

	preallocatedNodes [2]*Node
}

const nodesBlock = 1024

type Parser struct {
	*lexResult
	prev   int
	last   int
	errors []string

	nodes    []Node
	nextNode int

	Result *Node
}

func newParser(lr *lexResult) *Parser {
	return &Parser{lexResult: lr, prev: -1, last: len(lr.tokens) - 1}
}

func (p *Parser) Lex(lval *yySymType) int {
	if p.prev == p.last {
		return 0
	}
	p.prev++
	lval.token = p.prev
	return p.tokens[p.prev].sym
}

func (p *Parser) Error(s string) {
	s = p.tokens[p.prev].String() + ": " + s
	p.errors = append(p.errors, s)
}

func (p *Parser) NewNode(t NodeType, tokens ...int) *Node {
	if p.nextNode == len(p.nodes) {
		p.nodes = make([]Node, nodesBlock)
		p.nextNode = 0
	}
	n := &p.nodes[p.nextNode]
	p.nextNode++
	n.Type = t
	n.Tokens = tokens
	n.Nodes = n.preallocatedNodes[:0]
	return n
}

func (p *Parser) Op(t NodeType, tokens ...int) *Node { return p.NewNode(OpNode+t, tokens...) }

func (n *Node) SetType(t NodeType) *Node { n.Type = t; return n }

func (n *Node) T(tokens ...int) *Node   { n.Tokens = append(n.Tokens, tokens...); return n }
func (n *Node) T1(t0 int) *Node         { n.Tokens = append(n.Tokens, t0); return n }
func (n *Node) T2(t0, t1 int) *Node     { n.Tokens = append(n.Tokens, t0, t1); return n }
func (n *Node) T3(t0, t1, t2 int) *Node { n.Tokens = append(n.Tokens, t0, t1, t2); return n }

func (n *Node) N(nodes ...*Node) *Node    { n.Nodes = append(n.Nodes, nodes...); return n }
func (n *Node) N1(n0 *Node) *Node         { n.Nodes = append(n.Nodes, n0); return n }
func (n *Node) N2(n0, n1 *Node) *Node     { n.Nodes = append(n.Nodes, n0, n1); return n }
func (n *Node) N3(n0, n1, n2 *Node) *Node { n.Nodes = append(n.Nodes, n0, n1, n2); return n }

func (p *Parser) LispResult() string {
	return p.Result.lispFormat(p)
}

func (n *Node) lispFormat(p *Parser) string {
	s := n.Type.String()
	if s == "" {
		panic("unknown node type")
	}
	for _, child := range n.Nodes {
		s += " " + child.lispFormat(p)
	}
	switch n.Type {
	case URINode, PathNode, FloatNode, IntNode:
		s += " " + p.TokenString(n.Tokens[0])
	case IDNode:
		s += " |" + p.TokenString(n.Tokens[0]) + "|"
	case TextNode:
		s += ` "` + p.TokenString(n.Tokens[0]) + `"`
	}
	return "(" + s + ")"
}

func parse(lr *lexResult) (p *Parser, err error) {
	p = newParser(lr)
	yyErrorVerbose = true
	yyParse(p)
	if len(p.errors) == 0 {
		return
	}
	path := p.file.Name()
	err = fmt.Errorf("%s: %s", path, strings.Join(p.errors, "; "))
	return
}

func ParseFile(path string) (p *Parser, err error) {
	lr, err := lexFile(path)
	if err != nil {
		return
	}
	return parse(lr)
}

func ParseString(s string) (p *Parser, err error) {
	lr, err := lex([]byte(s), "")
	if err != nil {
		return
	}
	return parse(lr)
}
