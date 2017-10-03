// compile-command: "ragel -Z -G2 machine.rl"

package parser

import "fmt"

const maxstack = 64

%%{

machine expr;


prepush { if nostack() { return }; }

Space    = [ \t\r\n]+;
Comment1 = "#" [^\r\n]*;
Comment2 = "/*" any* :>> "*/";
URI      = [a-zA-Z] [a-zA-Z0-9.+\-]* ":" [a-zA-Z0-9%/?:@&=+$,_.!~*'\-]+;
Path     = [a-zA-Z0-9._+\-]* ([/] [a-zA-Z0-9._+\-]+)+ [/]?;
HPath    = "~" ([/] [a-zA-Z0-9._+\-]+)+ [/]?;
SPath    = "<" [a-zA-Z0-9._+\-]+ ([/] [a-zA-Z0-9._+\-]+)* ">";
ID       = [a-zA-Z0-9_] [a-zA-Z0-9_'\-]*;
Float    = (([1-9] [0-9]* [.] [0-9]*) | (0? [.] [0-9]+)) ([Ee] [+\-]? [0-9]+)?;
Int      = [0-9]+;


qstring := |*

["] => { if !tokleave('"') { return }; fret; };
"${" => { tokenter(interp, '}'); fcall expr; };
( [^$"\\] | /[$][^{"]/ | /\\./ )+ => { tok(text); addLines() };
"$" => { tok(text) };

*|;


istring := |*

"''" => { if !tokleave(ii) { return }; fret; };
"${" => { tokenter(interp, '}'); fcall expr; };
( [^'$] | "'''" | "''$" | /''\\./ | /'[^'$]/ | /[$][^{']/ )+ => { tok(text); addLines() };
['$] => { tok(text) };

*|;


expr := |*

"assert"  => { tok(assert_) };
"else"    => { tok(else_) };
"if"      => { tok(if_) };
"in"      => { tok(in) };
"inherit" => { tok(inherit) };
"let"     => { tok(let) };
"or"      => { tok(or_) };
"rec"     => { tok(rec) };
"then"    => { tok(then) };
"with"    => { tok(with) };

[ \t\r]+;
"\n"     => { r.file.AddLine(ts) };
Comment1 => { tokcomment(comment) };
Comment2 => { tokcomment(comment); addLines() };
URI      => { tok(uri) };
Path     => { tok(path) };
HPath    => { tok(path) };
SPath    => { tok(path) };
Float    => { tok(float) };
Int      => { tok(int_) };
ID       => { tok(id) };

"..." => { tok(ellipsis) };
"->"  => { tok(impl) };
"||"  => { tok(or) };
"&&"  => { tok(and) };
"=="  => { tok(eq) };
"!="  => { tok(neq) };
"<="  => { tok(leq) };
">="  => { tok(geq) };
"//"  => { tok(update) };
"++"  => { tok(concat) };

["]  => { tokenter('"', '"'); fcall qstring; };
"''" => { tokenter(ii, ii); fcall istring;  };

"${"   => { tokenter(interp, '}'); fcall expr; };
"(" => { tokenter('(', ')'); fcall expr; };
"[" => { tokenter('[', ']'); fcall expr; };
"{" => { tokenter('{', '}'); fcall expr; };
[}\])] => { if !tokleave(int(data[ts])) { return }; fret; };

[@:] => { if !tokarg() { return }; };

any => { tok(int(data[ts])) };

*|;

}%%

%% write data;

func lexData(data []byte, r *lexResult) (err error) {
	var cs, act, ts, te, top int
	var stack [maxstack]int
	p, pe := 0, len(data)
	eof := pe
%% write init;
	_, _, _, _, _ = expr_first_final, expr_error, expr_en_qstring, expr_en_istring, expr_en_expr
	if r.file == nil {
		r.file = fileset.AddFile("(string)", -1, len(data))
	}
	if r.data == nil {
		r.data = data
	} else {
		r.data = append(r.data, data...)
	}
	nostack := func() bool {
		if top != maxstack {
			return false
		}
		err = r.Errorf("exceeds recursion limit")
		return true
	}
	tok := func(sym int) { r.tokens = append(r.tokens, lexerToken{sym: sym, pos: ts, end: te}) }
	tokcomment := func(sym int) { r.comments = append(r.comments, lexerToken{sym: sym, pos: ts, end: te}) }
	var backrefs Backrefs
	tokenter := func(sym, fin int) { backrefs.Push(len(r.tokens), fin); tok(sym); }
	tokleave := func(sym int) bool {
		tok(sym)
		if top == 0 || len(backrefs) == 0 {
			err = r.Errorf("does not close anything")
			return false
		}
		iprev, prevsym := backrefs.Pop()
		if prevsym != sym {
			err = r.Errorf("does not close %v", r.tokens[iprev])
			return false
		}
		r.tokens[len(r.tokens)-1].prev = iprev
		return true
	}
	tokarg := func() bool {
		tok(int(data[ts]))
		if len(r.tokens) == 1 {
			err = r.Errorf("does not follow anything")
			return false
		}
		prev := &r.tokens[len(r.tokens)-2]
		switch prev.sym {
		case id:
			prev.sym = argID
		case '}':
			r.tokens[prev.prev].sym = argBracket
		default:
			err = r.Errorf("does not follow an argument of a function")
			return false
		}
		return true
	}
	addLines := func() {
		for i := ts; i < te; i++ {
			if data[i] == '\n' {
				r.file.AddLine(i)
			}
		}
	}

%% write exec;

	if p != eof {
		err = r.Errorf("precedes the token that failed to lex")
	} else if len(backrefs) != 0 {
		iprev, _ := backrefs.Pop()
		prev := r.tokens[iprev]
		err = fmt.Errorf("%s%s is not terminated", r.At(prev.pos), symString(prev.sym))
	}
	return
}
