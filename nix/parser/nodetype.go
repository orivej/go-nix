package parser

type NodeType uint16

const (
	ApplyNode NodeType = iota
	ArgNode
	ArgSetNode
	AssertNode
	AttrPathNode
	BindNode
	BindsNode
	FloatNode
	FunctionNode
	IDNode
	IStringNode
	IfNode
	InheritFromNode
	InheritListNode
	InheritNode
	IntNode
	InterpNode
	LetNode
	ListNode
	ParensNode
	PathNode
	RecSetNode
	SelectNode
	SelectOrNode
	SetNode
	StringNode
	TextNode
	URINode
	WithNode

	OpNode
)

var nodeName = map[NodeType]string{
	ApplyNode:       "apply",
	ArgNode:         "arg",
	ArgSetNode:      "argset",
	AssertNode:      "assert",
	AttrPathNode:    "attrpath",
	BindNode:        "bind",
	BindsNode:       "binds",
	FloatNode:       "float",
	FunctionNode:    "function",
	IDNode:          "id",
	IStringNode:     "istring",
	IfNode:          "if",
	InheritFromNode: "inheritfrom",
	InheritListNode: "inheritlist",
	InheritNode:     "inherit",
	IntNode:         "int",
	InterpNode:      "interp",
	LetNode:         "let",
	ListNode:        "list",
	ParensNode:      "parens",
	PathNode:        "path",
	RecSetNode:      "recset",
	SelectNode:      "select",
	SelectOrNode:    "selector",
	SetNode:         "set",
	StringNode:      "string",
	TextNode:        "text",
	URINode:         "uri",
	WithNode:        "with",

	OpNode + negate: "-",
	OpNode + '?':    "?",
	OpNode + concat: "++",
	OpNode + '/':    "/",
	OpNode + '*':    "*",
	OpNode + '-':    "-",
	OpNode + '+':    "+",
	OpNode + '!':    "!",
	OpNode + update: "//",
	OpNode + geq:    ">=",
	OpNode + leq:    "<=",
	OpNode + '>':    ">",
	OpNode + '<':    "<",
	OpNode + neq:    "!=",
	OpNode + eq:     "==",
	OpNode + and:    "&&",
	OpNode + or:     "||",
	OpNode + impl:   "->",
}

func (nt NodeType) String() string {
	if s, ok := nodeName[nt]; ok {
		return s
	}
	panic("unknown node type")
}
