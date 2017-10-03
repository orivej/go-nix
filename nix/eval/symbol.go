package eval

type Sym int

type Symtab struct {
	names []string
	syms  map[string]Sym
}

func NewSymtab() *Symtab {
	return &Symtab{names: []string{""}, syms: map[string]Sym{"": 0}}
}

func (st *Symtab) Intern(name string) Sym {
	if sym, ok := st.syms[name]; ok {
		return sym
	}
	sym := Sym(len(st.names))
	st.names = append(st.names, name)
	st.syms[name] = sym
	return sym
}

func (st *Symtab) Name(sym Sym) string {
	return st.names[sym]
}

var globalSymtab = NewSymtab()

func Intern(name string) Sym {
	return globalSymtab.Intern(name)
}

func (sym Sym) String() string {
	return globalSymtab.Name(sym)
}
