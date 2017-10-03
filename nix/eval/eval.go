package eval

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/orivej/e"
	p "github.com/orivej/go-nix/nix/parser"
)

type Scope struct {
	Binds   Set
	LowPrio bool
	Parent  *Scope
}

func (scope *Scope) Subscope(binds Set, lowPrio bool) *Scope {
	return &Scope{Binds: binds, LowPrio: lowPrio, Parent: scope}
}

type Expression struct {
	Value  Value
	Scope  *Scope
	Parser *p.Parser
	Node   *p.Node
	Sym    Sym
}

func (x *Expression) String() string {
	return ValueString(x.Value)
}

func (x *Expression) WithNode(n *p.Node) *Expression {
	y := *x
	y.Node = n
	return &y
}

func (x *Expression) WithScoped(n *p.Node, scope *Scope) *Expression {
	y := *x
	y.Scope = scope
	y.Node = n
	return &y
}

func (x *Expression) Eval() Value {
	if x.Value == nil {
		x.Value = x.force()
		x.Scope = nil
	}
	return x.Value
}

func (x *Expression) tokenString(i int) string {
	return x.Parser.TokenString(x.Node.Tokens[i])
}

func (x *Expression) force() (val Value) {
	var err error
	pr := x.Parser
	n := x.Node
	nt := x.Node.Type
	if x.Sym != 0 {
		nt = p.IDNode
	}
	switch nt {
	default:
		panic(fmt.Sprintln("unsupported node type:", n.Type))
	case p.URINode:
		return URI(x.tokenString(0))
	case p.PathNode:
		return Path(x.tokenString(0))
	case p.FloatNode:
		val, err = strconv.ParseFloat(x.tokenString(0), 64)
		noerr(err)
	case p.IntNode:
		val, err = strconv.Atoi(x.tokenString(0))
		noerr(err)

	case p.StringNode, p.IStringNode:
		parts := make([]string, len(n.Nodes))
		for i, c := range n.Nodes {
			switch c.Type {
			default:
				panic(fmt.Sprintln("unsupported string part type:", c.Type))
			case p.TextNode:
				parts[i] = pr.TokenString(c.Tokens[0])
			case p.InterpNode:
				y := x.WithNode(c.Nodes[0])
				parts[i] = InterpString(y.Eval())
			}
		}
		return strings.Join(parts, "")

	case p.ParensNode:
		return x.WithNode(n.Nodes[0]).Eval()

	case p.ListNode:
		parts := make(List, len(n.Nodes))
		for i, c := range n.Nodes {
			parts[i] = x.WithNode(c)
		}
		return parts

	case p.SetNode, p.RecSetNode:
		set := make(Set, len(n.Nodes)) // Inheriting makes it larger than this.
		scope := x.Scope
		if nt == p.RecSetNode {
			scope = scope.Subscope(set, false)
		}
		for _, c := range n.Nodes {
			switch c.Type {
			case p.BindNode:
				attrpath := x.WithScoped(c.Nodes[0], scope).evalAttrpath()
				set.Bind(attrpath, x.WithScoped(c.Nodes[1], scope))
			case p.InheritNode:
				for _, interpid := range c.Nodes[0].Nodes {
					y := x.WithNode(interpid)
					y.Sym = Intern(y.attrString())
					set.Bind1(y.Sym, y)
				}
			case p.InheritFromNode:
				// This is not as lazy as it can be.
				from := x.WithScoped(c.Nodes[0], scope)
				for _, interpid := range c.Nodes[1].Nodes {
					sym := Intern(x.WithNode(interpid).attrString())
					set.Bind1(sym, from.Select1(sym))
				}
			}
		}
		return set

	case p.SelectNode, p.SelectOrNode:
		attrpath := x.WithNode(n.Nodes[1]).evalAttrpath()
		var or *Expression
		if nt == p.SelectOrNode {
			or = x.WithNode(n.Nodes[2])
		}
		return x.WithNode(n.Nodes[0]).Select(attrpath, or).Eval()
	}

	return
}

func (x *Expression) Select1(sym Sym) *Expression {
	return x.Eval().(Set)[sym]
}

func (x *Expression) Select(syms []Sym, or *Expression) *Expression {
	for _, sym := range syms {
		val := x.Eval()
		if set, ok := val.(Set); ok {
			if y, ok := set[sym]; ok {
				x = y
			} else if or != nil {
				return or
			} else {
				throw(fmt.Errorf("%v does not contain %v", y, sym))
			}
		} else {
			throw(fmt.Errorf("%v is not a set", val))
		}
	}
	return x
}

func (x *Expression) evalAttrpath() []Sym {
	attrs := make([]Sym, len(x.Node.Nodes))
	for i, c := range x.Node.Nodes {
		y := x.WithNode(c)
		switch c.Type {
		case p.IDNode:
			attrs[i] = Intern(y.tokenString(0))
		case p.StringNode:
			attrs[i] = Intern(InterpString(y.Eval()))
		case p.InterpNode:
			attrs[i] = Intern(InterpString(y.WithNode(y.Node.Nodes[0]).Eval()))
		default:
			panic(fmt.Errorf("unsupported attrpath node %v", c.Type))
		}
	}
	return attrs
}

func (x *Expression) attrString() string {
	switch x.Node.Type {
	case p.IDNode:
		return x.tokenString(0)
	case p.StringNode:
		return x.Eval().(string)
	case p.InterpNode:
		return InterpString(x.WithNode(x.Node.Nodes[0]).Eval())
	default:
		panic(fmt.Errorf("unsupported attr type %v", x.Node.Type))
	}
}

func ParseResult(pr *p.Parser) Value {
	x := Expression{Parser: pr, Node: pr.Result}
	return x.Eval()
}

func throw(err error) {
	if err != nil {
		panic(err)
	}
}

func noerr(err error) {
	e.Exit(err)
}
