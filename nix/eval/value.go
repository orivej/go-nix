package eval

import (
	"fmt"
	"strconv"
	"strings"
)

// Value is one of: int, float64, string, URI, Path, List, Set
type Value interface{}

type URI string
type Path string
type List []*Expression
type Set map[Sym]*Expression

func (set Set) Bind1(sym Sym, x *Expression) {
	if _, ok := set[sym]; ok {
		throw(fmt.Errorf("%v is already defined", sym))
	}
	set[sym] = x
}

func (set Set) Bind(syms []Sym, x *Expression) {
	last := len(syms) - 1
	for _, sym := range syms[:last] {
		if subset, ok := set[sym]; ok {
			set = subset.Value.(Set)
		} else {
			subset := Set{}
			set[sym] = &Expression{Value: subset}
			set = subset
		}
	}
	set.Bind1(syms[last], x)
}

func ValueString(val Value) string {
	switch v := val.(type) {
	case nil:
		return "..."
	case int:
		return strconv.Itoa(v)
	case float64:
		return strconv.FormatFloat(v, 'g', -1, 64)
	case string:
		return v
	case URI:
		return string(v)
	case Path:
		return string(v)
	case List:
		last := len(v) + 1
		parts := make([]string, last+1)
		parts[0], parts[last] = "[", "]"
		for i, x := range v {
			parts[i+1] = x.String()
		}
		return strings.Join(parts, " ")
	case Set:
		last := len(v) + 1
		parts := make([]string, last+1)
		parts[0], parts[last] = "{", "}"
		i := 1
		for sym, x := range v {
			parts[i] = fmt.Sprintf("%s = %s;", sym, x)
			i++
		}
		return strings.Join(parts, " ")
	default:
		panic(fmt.Errorf("can not coerce %v to a string", val))
	}
}

func InterpString(val Value) string {
	switch v := val.(type) {
	case string:
		return v
	case URI:
		return string(v)
	case Path:
		return string(v)
	default:
		panic(fmt.Errorf("can not coerce %v to a string", val))
	}
}
