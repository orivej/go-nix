
//line machine.rl:1
// compile-command: "ragel -Z -G2 machine.rl"

package parser

import "fmt"

const maxstack = 64


//line machine.rl:99



//line machine.go:17
const expr_start int = 22
const expr_first_final int = 22
const expr_error int = 0

const expr_en_qstring int = 79
const expr_en_istring int = 82
const expr_en_expr int = 22


//line machine.rl:102

func lexData(data []byte, r *lexResult) (err error) {
	var cs, act, ts, te, top int
	var stack [maxstack]int
	p, pe := 0, len(data)
	eof := pe

//line machine.go:35
	{
	cs = expr_start
	top = 0
	ts = 0
	te = 0
	act = 0
	}

//line machine.rl:109
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


//line machine.go:106
	{
	if p == pe {
		goto _test_eof
	}
	goto _resume

_again:
	switch cs {
	case 22:
		goto st22
	case 23:
		goto st23
	case 1:
		goto st1
	case 24:
		goto st24
	case 2:
		goto st2
	case 3:
		goto st3
	case 25:
		goto st25
	case 26:
		goto st26
	case 27:
		goto st27
	case 28:
		goto st28
	case 29:
		goto st29
	case 30:
		goto st30
	case 31:
		goto st31
	case 32:
		goto st32
	case 33:
		goto st33
	case 4:
		goto st4
	case 5:
		goto st5
	case 34:
		goto st34
	case 35:
		goto st35
	case 36:
		goto st36
	case 37:
		goto st37
	case 6:
		goto st6
	case 38:
		goto st38
	case 7:
		goto st7
	case 8:
		goto st8
	case 39:
		goto st39
	case 40:
		goto st40
	case 9:
		goto st9
	case 10:
		goto st10
	case 41:
		goto st41
	case 42:
		goto st42
	case 43:
		goto st43
	case 44:
		goto st44
	case 45:
		goto st45
	case 11:
		goto st11
	case 12:
		goto st12
	case 46:
		goto st46
	case 47:
		goto st47
	case 48:
		goto st48
	case 13:
		goto st13
	case 14:
		goto st14
	case 49:
		goto st49
	case 50:
		goto st50
	case 51:
		goto st51
	case 52:
		goto st52
	case 53:
		goto st53
	case 54:
		goto st54
	case 55:
		goto st55
	case 56:
		goto st56
	case 57:
		goto st57
	case 58:
		goto st58
	case 59:
		goto st59
	case 60:
		goto st60
	case 61:
		goto st61
	case 62:
		goto st62
	case 63:
		goto st63
	case 64:
		goto st64
	case 65:
		goto st65
	case 66:
		goto st66
	case 67:
		goto st67
	case 68:
		goto st68
	case 69:
		goto st69
	case 70:
		goto st70
	case 71:
		goto st71
	case 72:
		goto st72
	case 73:
		goto st73
	case 74:
		goto st74
	case 75:
		goto st75
	case 76:
		goto st76
	case 15:
		goto st15
	case 77:
		goto st77
	case 78:
		goto st78
	case 79:
		goto st79
	case 80:
		goto st80
	case 16:
		goto st16
	case 17:
		goto st17
	case 81:
		goto st81
	case 82:
		goto st82
	case 83:
		goto st83
	case 18:
		goto st18
	case 19:
		goto st19
	case 20:
		goto st20
	case 21:
		goto st21
	case 84:
		goto st84
	case 85:
		goto st85
	case 86:
		goto st86
	case 0:
		goto st0
	}

	if p++; p == pe {
		goto _test_eof
	}
_resume:
	switch cs {
	case 22:
		goto st_case_22
	case 23:
		goto st_case_23
	case 1:
		goto st_case_1
	case 24:
		goto st_case_24
	case 2:
		goto st_case_2
	case 3:
		goto st_case_3
	case 25:
		goto st_case_25
	case 26:
		goto st_case_26
	case 27:
		goto st_case_27
	case 28:
		goto st_case_28
	case 29:
		goto st_case_29
	case 30:
		goto st_case_30
	case 31:
		goto st_case_31
	case 32:
		goto st_case_32
	case 33:
		goto st_case_33
	case 4:
		goto st_case_4
	case 5:
		goto st_case_5
	case 34:
		goto st_case_34
	case 35:
		goto st_case_35
	case 36:
		goto st_case_36
	case 37:
		goto st_case_37
	case 6:
		goto st_case_6
	case 38:
		goto st_case_38
	case 7:
		goto st_case_7
	case 8:
		goto st_case_8
	case 39:
		goto st_case_39
	case 40:
		goto st_case_40
	case 9:
		goto st_case_9
	case 10:
		goto st_case_10
	case 41:
		goto st_case_41
	case 42:
		goto st_case_42
	case 43:
		goto st_case_43
	case 44:
		goto st_case_44
	case 45:
		goto st_case_45
	case 11:
		goto st_case_11
	case 12:
		goto st_case_12
	case 46:
		goto st_case_46
	case 47:
		goto st_case_47
	case 48:
		goto st_case_48
	case 13:
		goto st_case_13
	case 14:
		goto st_case_14
	case 49:
		goto st_case_49
	case 50:
		goto st_case_50
	case 51:
		goto st_case_51
	case 52:
		goto st_case_52
	case 53:
		goto st_case_53
	case 54:
		goto st_case_54
	case 55:
		goto st_case_55
	case 56:
		goto st_case_56
	case 57:
		goto st_case_57
	case 58:
		goto st_case_58
	case 59:
		goto st_case_59
	case 60:
		goto st_case_60
	case 61:
		goto st_case_61
	case 62:
		goto st_case_62
	case 63:
		goto st_case_63
	case 64:
		goto st_case_64
	case 65:
		goto st_case_65
	case 66:
		goto st_case_66
	case 67:
		goto st_case_67
	case 68:
		goto st_case_68
	case 69:
		goto st_case_69
	case 70:
		goto st_case_70
	case 71:
		goto st_case_71
	case 72:
		goto st_case_72
	case 73:
		goto st_case_73
	case 74:
		goto st_case_74
	case 75:
		goto st_case_75
	case 76:
		goto st_case_76
	case 15:
		goto st_case_15
	case 77:
		goto st_case_77
	case 78:
		goto st_case_78
	case 79:
		goto st_case_79
	case 80:
		goto st_case_80
	case 16:
		goto st_case_16
	case 17:
		goto st_case_17
	case 81:
		goto st_case_81
	case 82:
		goto st_case_82
	case 83:
		goto st_case_83
	case 18:
		goto st_case_18
	case 19:
		goto st_case_19
	case 20:
		goto st_case_20
	case 21:
		goto st_case_21
	case 84:
		goto st_case_84
	case 85:
		goto st_case_85
	case 86:
		goto st_case_86
	case 0:
		goto st_case_0
	}
	goto st_out
tr0:
//line machine.rl:95
p = (te) - 1
{ tok(int(data[ts])) }
	goto st22
tr2:
//line machine.rl:69
p = (te) - 1
{ tok(float) }
	goto st22
tr5:
//line NONE:1
	switch act {
	case 9:
	{p = (te) - 1
 tok(assert_) }
	case 10:
	{p = (te) - 1
 tok(else_) }
	case 11:
	{p = (te) - 1
 tok(if_) }
	case 12:
	{p = (te) - 1
 tok(in) }
	case 13:
	{p = (te) - 1
 tok(inherit) }
	case 14:
	{p = (te) - 1
 tok(let) }
	case 15:
	{p = (te) - 1
 tok(or_) }
	case 16:
	{p = (te) - 1
 tok(rec) }
	case 17:
	{p = (te) - 1
 tok(then) }
	case 18:
	{p = (te) - 1
 tok(with) }
	case 27:
	{p = (te) - 1
 tok(float) }
	case 28:
	{p = (te) - 1
 tok(int_) }
	case 29:
	{p = (te) - 1
 tok(id) }
	case 30:
	{p = (te) - 1
 tok(ellipsis) }
	case 39:
	{p = (te) - 1
 tok(concat) }
	case 48:
	{p = (te) - 1
 tok(int(data[ts])) }
	}
	
	goto st22
tr14:
//line machine.rl:64
te = p+1
{ tokcomment(comment); addLines() }
	goto st22
tr17:
//line machine.rl:68
te = p+1
{ tok(path) }
	goto st22
tr31:
//line machine.rl:95
te = p+1
{ tok(int(data[ts])) }
	goto st22
tr33:
//line machine.rl:62
te = p+1
{ r.file.AddLine(ts) }
	goto st22
tr35:
//line machine.rl:84
te = p+1
{ tokenter('"', '"'); { if nostack() { return }; {stack[top] = 22; top++; goto st79 }} }
	goto st22
tr40:
//line machine.rl:88
te = p+1
{ tokenter('(', ')'); { if nostack() { return }; {stack[top] = 22; top++; goto st22 }} }
	goto st22
tr41:
//line machine.rl:91
te = p+1
{ if !tokleave(int(data[ts])) { return }; {top--; cs = stack[top];goto _again } }
	goto st22
tr48:
//line machine.rl:93
te = p+1
{ if !tokarg() { return }; }
	goto st22
tr53:
//line machine.rl:89
te = p+1
{ tokenter('[', ']'); { if nostack() { return }; {stack[top] = 22; top++; goto st22 }} }
	goto st22
tr63:
//line machine.rl:90
te = p+1
{ tokenter('{', '}'); { if nostack() { return }; {stack[top] = 22; top++; goto st22 }} }
	goto st22
tr66:
//line machine.rl:95
te = p
p--
{ tok(int(data[ts])) }
	goto st22
tr68:
//line machine.rl:69
te = p
p--
{ tok(float) }
	goto st22
tr70:
//line machine.rl:61
te = p
p--

	goto st22
tr71:
//line machine.rl:78
te = p+1
{ tok(neq) }
	goto st22
tr72:
//line machine.rl:63
te = p
p--
{ tokcomment(comment) }
	goto st22
tr73:
//line machine.rl:87
te = p+1
{ tokenter(interp, '}'); { if nostack() { return }; {stack[top] = 22; top++; goto st22 }} }
	goto st22
tr74:
//line machine.rl:76
te = p+1
{ tok(and) }
	goto st22
tr75:
//line machine.rl:85
te = p+1
{ tokenter(ii, ii); { if nostack() { return }; {stack[top] = 22; top++; goto st82 }}  }
	goto st22
tr77:
//line machine.rl:66
te = p
p--
{ tok(path) }
	goto st22
tr79:
//line machine.rl:74
te = p+1
{ tok(impl) }
	goto st22
tr83:
//line machine.rl:81
te = p+1
{ tok(update) }
	goto st22
tr84:
//line machine.rl:70
te = p
p--
{ tok(int_) }
	goto st22
tr86:
//line machine.rl:71
te = p
p--
{ tok(id) }
	goto st22
tr87:
//line machine.rl:79
te = p+1
{ tok(leq) }
	goto st22
tr88:
//line machine.rl:77
te = p+1
{ tok(eq) }
	goto st22
tr89:
//line machine.rl:80
te = p+1
{ tok(geq) }
	goto st22
tr90:
//line machine.rl:65
te = p
p--
{ tok(uri) }
	goto st22
tr101:
//line machine.rl:53
te = p
p--
{ tok(in) }
	goto st22
tr118:
//line machine.rl:75
te = p+1
{ tok(or) }
	goto st22
tr120:
//line machine.rl:67
te = p
p--
{ tok(path) }
	goto st22
	st22:
//line NONE:1
ts = 0

		if p++; p == pe {
			goto _test_eof22
		}
	st_case_22:
//line NONE:1
ts = p

//line machine.go:707
		switch data[p] {
		case 0:
			goto tr30
		case 9:
			goto st26
		case 10:
			goto tr33
		case 13:
			goto st26
		case 32:
			goto st26
		case 33:
			goto st27
		case 34:
			goto tr35
		case 35:
			goto st28
		case 36:
			goto st29
		case 38:
			goto st30
		case 39:
			goto st31
		case 40:
			goto tr40
		case 41:
			goto tr41
		case 43:
			goto tr42
		case 45:
			goto tr43
		case 46:
			goto tr44
		case 47:
			goto tr45
		case 48:
			goto tr46
		case 58:
			goto tr48
		case 60:
			goto tr49
		case 61:
			goto st46
		case 62:
			goto st47
		case 64:
			goto tr48
		case 91:
			goto tr53
		case 93:
			goto tr41
		case 95:
			goto tr54
		case 97:
			goto tr55
		case 101:
			goto tr56
		case 105:
			goto tr57
		case 108:
			goto tr58
		case 111:
			goto tr59
		case 114:
			goto tr60
		case 116:
			goto tr61
		case 119:
			goto tr62
		case 123:
			goto tr63
		case 124:
			goto st75
		case 125:
			goto tr41
		case 126:
			goto tr65
		}
		switch {
		case data[p] < 65:
			if 49 <= data[p] && data[p] <= 57 {
				goto tr47
			}
		case data[p] > 90:
			if 98 <= data[p] && data[p] <= 122 {
				goto tr52
			}
		default:
			goto tr52
		}
		goto tr31
tr30:
//line NONE:1
te = p+1

	goto st23
	st23:
		if p++; p == pe {
			goto _test_eof23
		}
	st_case_23:
//line machine.go:809
		if data[p] == 46 {
			goto st1
		}
		goto tr66
	st1:
		if p++; p == pe {
			goto _test_eof1
		}
	st_case_1:
		if 48 <= data[p] && data[p] <= 57 {
			goto tr1
		}
		goto tr0
tr1:
//line NONE:1
te = p+1

	goto st24
	st24:
		if p++; p == pe {
			goto _test_eof24
		}
	st_case_24:
//line machine.go:833
		switch data[p] {
		case 69:
			goto st2
		case 101:
			goto st2
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr1
		}
		goto tr68
	st2:
		if p++; p == pe {
			goto _test_eof2
		}
	st_case_2:
		switch data[p] {
		case 43:
			goto st3
		case 45:
			goto st3
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st25
		}
		goto tr2
	st3:
		if p++; p == pe {
			goto _test_eof3
		}
	st_case_3:
		if 48 <= data[p] && data[p] <= 57 {
			goto st25
		}
		goto tr2
	st25:
		if p++; p == pe {
			goto _test_eof25
		}
	st_case_25:
		if 48 <= data[p] && data[p] <= 57 {
			goto st25
		}
		goto tr68
	st26:
		if p++; p == pe {
			goto _test_eof26
		}
	st_case_26:
		switch data[p] {
		case 9:
			goto st26
		case 13:
			goto st26
		case 32:
			goto st26
		}
		goto tr70
	st27:
		if p++; p == pe {
			goto _test_eof27
		}
	st_case_27:
		if data[p] == 61 {
			goto tr71
		}
		goto tr66
	st28:
		if p++; p == pe {
			goto _test_eof28
		}
	st_case_28:
		switch data[p] {
		case 10:
			goto tr72
		case 13:
			goto tr72
		}
		goto st28
	st29:
		if p++; p == pe {
			goto _test_eof29
		}
	st_case_29:
		if data[p] == 123 {
			goto tr73
		}
		goto tr66
	st30:
		if p++; p == pe {
			goto _test_eof30
		}
	st_case_30:
		if data[p] == 38 {
			goto tr74
		}
		goto tr66
	st31:
		if p++; p == pe {
			goto _test_eof31
		}
	st_case_31:
		if data[p] == 39 {
			goto tr75
		}
		goto tr66
tr42:
//line NONE:1
te = p+1

//line machine.rl:95
act = 48;
	goto st32
	st32:
		if p++; p == pe {
			goto _test_eof32
		}
	st_case_32:
//line machine.go:951
		switch data[p] {
		case 43:
			goto tr76
		case 47:
			goto st5
		case 95:
			goto st4
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto st4
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st4
			}
		default:
			goto st4
		}
		goto tr66
tr9:
//line NONE:1
te = p+1

//line machine.rl:73
act = 30;
	goto st33
tr76:
//line NONE:1
te = p+1

//line machine.rl:82
act = 39;
	goto st33
	st33:
		if p++; p == pe {
			goto _test_eof33
		}
	st_case_33:
//line machine.go:992
		switch data[p] {
		case 43:
			goto st4
		case 47:
			goto st5
		case 95:
			goto st4
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto st4
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st4
			}
		default:
			goto st4
		}
		goto tr5
	st4:
		if p++; p == pe {
			goto _test_eof4
		}
	st_case_4:
		switch data[p] {
		case 43:
			goto st4
		case 47:
			goto st5
		case 95:
			goto st4
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto st4
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st4
			}
		default:
			goto st4
		}
		goto tr5
	st5:
		if p++; p == pe {
			goto _test_eof5
		}
	st_case_5:
		switch data[p] {
		case 43:
			goto st34
		case 95:
			goto st34
		}
		switch {
		case data[p] < 48:
			if 45 <= data[p] && data[p] <= 46 {
				goto st34
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st34
				}
			case data[p] >= 65:
				goto st34
			}
		default:
			goto st34
		}
		goto tr5
	st34:
		if p++; p == pe {
			goto _test_eof34
		}
	st_case_34:
		switch data[p] {
		case 43:
			goto st34
		case 47:
			goto st35
		case 95:
			goto st34
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto st34
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st34
			}
		default:
			goto st34
		}
		goto tr77
	st35:
		if p++; p == pe {
			goto _test_eof35
		}
	st_case_35:
		switch data[p] {
		case 43:
			goto st34
		case 95:
			goto st34
		}
		switch {
		case data[p] < 48:
			if 45 <= data[p] && data[p] <= 46 {
				goto st34
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st34
				}
			case data[p] >= 65:
				goto st34
			}
		default:
			goto st34
		}
		goto tr77
tr43:
//line NONE:1
te = p+1

//line machine.rl:95
act = 48;
	goto st36
	st36:
		if p++; p == pe {
			goto _test_eof36
		}
	st_case_36:
//line machine.go:1136
		switch data[p] {
		case 43:
			goto st4
		case 47:
			goto st5
		case 62:
			goto tr79
		case 95:
			goto st4
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto st4
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st4
			}
		default:
			goto st4
		}
		goto tr66
tr44:
//line NONE:1
te = p+1

//line machine.rl:95
act = 48;
	goto st37
	st37:
		if p++; p == pe {
			goto _test_eof37
		}
	st_case_37:
//line machine.go:1172
		switch data[p] {
		case 43:
			goto st4
		case 45:
			goto st4
		case 46:
			goto st6
		case 47:
			goto st5
		case 95:
			goto st4
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr81
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st4
			}
		default:
			goto st4
		}
		goto tr66
	st6:
		if p++; p == pe {
			goto _test_eof6
		}
	st_case_6:
		switch data[p] {
		case 43:
			goto st4
		case 46:
			goto tr9
		case 47:
			goto st5
		case 95:
			goto st4
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto st4
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st4
			}
		default:
			goto st4
		}
		goto tr0
tr81:
//line NONE:1
te = p+1

//line machine.rl:69
act = 27;
	goto st38
	st38:
		if p++; p == pe {
			goto _test_eof38
		}
	st_case_38:
//line machine.go:1238
		switch data[p] {
		case 43:
			goto st4
		case 47:
			goto st5
		case 69:
			goto st7
		case 95:
			goto st4
		case 101:
			goto st7
		}
		switch {
		case data[p] < 48:
			if 45 <= data[p] && data[p] <= 46 {
				goto st4
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st4
				}
			case data[p] >= 65:
				goto st4
			}
		default:
			goto tr81
		}
		goto tr68
	st7:
		if p++; p == pe {
			goto _test_eof7
		}
	st_case_7:
		switch data[p] {
		case 43:
			goto st8
		case 45:
			goto st8
		case 46:
			goto st4
		case 47:
			goto st5
		case 95:
			goto st4
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr11
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st4
			}
		default:
			goto st4
		}
		goto tr2
	st8:
		if p++; p == pe {
			goto _test_eof8
		}
	st_case_8:
		switch data[p] {
		case 43:
			goto st4
		case 47:
			goto st5
		case 95:
			goto st4
		}
		switch {
		case data[p] < 48:
			if 45 <= data[p] && data[p] <= 46 {
				goto st4
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st4
				}
			case data[p] >= 65:
				goto st4
			}
		default:
			goto tr11
		}
		goto tr2
tr11:
//line NONE:1
te = p+1

//line machine.rl:69
act = 27;
	goto st39
	st39:
		if p++; p == pe {
			goto _test_eof39
		}
	st_case_39:
//line machine.go:1342
		switch data[p] {
		case 43:
			goto st4
		case 47:
			goto st5
		case 95:
			goto st4
		}
		switch {
		case data[p] < 48:
			if 45 <= data[p] && data[p] <= 46 {
				goto st4
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st4
				}
			case data[p] >= 65:
				goto st4
			}
		default:
			goto tr11
		}
		goto tr68
tr45:
//line NONE:1
te = p+1

	goto st40
	st40:
		if p++; p == pe {
			goto _test_eof40
		}
	st_case_40:
//line machine.go:1379
		switch data[p] {
		case 42:
			goto st9
		case 43:
			goto st34
		case 47:
			goto tr83
		case 95:
			goto st34
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto st34
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st34
			}
		default:
			goto st34
		}
		goto tr66
	st9:
		if p++; p == pe {
			goto _test_eof9
		}
	st_case_9:
		if data[p] == 42 {
			goto st10
		}
		goto st9
	st10:
		if p++; p == pe {
			goto _test_eof10
		}
	st_case_10:
		switch data[p] {
		case 42:
			goto st10
		case 47:
			goto tr14
		}
		goto st9
tr46:
//line NONE:1
te = p+1

//line machine.rl:70
act = 28;
	goto st41
	st41:
		if p++; p == pe {
			goto _test_eof41
		}
	st_case_41:
//line machine.go:1436
		switch data[p] {
		case 39:
			goto st42
		case 43:
			goto st4
		case 45:
			goto tr54
		case 46:
			goto st4
		case 47:
			goto st5
		case 95:
			goto tr54
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr46
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr54
			}
		default:
			goto tr54
		}
		goto tr84
	st42:
		if p++; p == pe {
			goto _test_eof42
		}
	st_case_42:
		switch data[p] {
		case 39:
			goto st42
		case 45:
			goto st42
		case 95:
			goto st42
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st42
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st42
			}
		default:
			goto st42
		}
		goto tr86
tr54:
//line NONE:1
te = p+1

//line machine.rl:71
act = 29;
	goto st43
	st43:
		if p++; p == pe {
			goto _test_eof43
		}
	st_case_43:
//line machine.go:1502
		switch data[p] {
		case 39:
			goto st42
		case 43:
			goto st4
		case 46:
			goto st4
		case 47:
			goto st5
		case 95:
			goto tr54
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr54
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr54
			}
		default:
			goto tr54
		}
		goto tr86
tr47:
//line NONE:1
te = p+1

//line machine.rl:70
act = 28;
	goto st44
	st44:
		if p++; p == pe {
			goto _test_eof44
		}
	st_case_44:
//line machine.go:1540
		switch data[p] {
		case 39:
			goto st42
		case 43:
			goto st4
		case 45:
			goto tr54
		case 46:
			goto tr81
		case 47:
			goto st5
		case 95:
			goto tr54
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr47
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr54
			}
		default:
			goto tr54
		}
		goto tr84
tr49:
//line NONE:1
te = p+1

	goto st45
	st45:
		if p++; p == pe {
			goto _test_eof45
		}
	st_case_45:
//line machine.go:1578
		switch data[p] {
		case 43:
			goto st11
		case 61:
			goto tr87
		case 95:
			goto st11
		}
		switch {
		case data[p] < 48:
			if 45 <= data[p] && data[p] <= 46 {
				goto st11
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st11
				}
			case data[p] >= 65:
				goto st11
			}
		default:
			goto st11
		}
		goto tr66
	st11:
		if p++; p == pe {
			goto _test_eof11
		}
	st_case_11:
		switch data[p] {
		case 43:
			goto st11
		case 47:
			goto st12
		case 62:
			goto tr17
		case 95:
			goto st11
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto st11
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st11
			}
		default:
			goto st11
		}
		goto tr0
	st12:
		if p++; p == pe {
			goto _test_eof12
		}
	st_case_12:
		switch data[p] {
		case 43:
			goto st11
		case 95:
			goto st11
		}
		switch {
		case data[p] < 48:
			if 45 <= data[p] && data[p] <= 46 {
				goto st11
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st11
				}
			case data[p] >= 65:
				goto st11
			}
		default:
			goto st11
		}
		goto tr0
	st46:
		if p++; p == pe {
			goto _test_eof46
		}
	st_case_46:
		if data[p] == 61 {
			goto tr88
		}
		goto tr66
	st47:
		if p++; p == pe {
			goto _test_eof47
		}
	st_case_47:
		if data[p] == 61 {
			goto tr89
		}
		goto tr66
tr52:
//line NONE:1
te = p+1

//line machine.rl:71
act = 29;
	goto st48
tr95:
//line NONE:1
te = p+1

//line machine.rl:50
act = 9;
	goto st48
tr98:
//line NONE:1
te = p+1

//line machine.rl:51
act = 10;
	goto st48
tr99:
//line NONE:1
te = p+1

//line machine.rl:52
act = 11;
	goto st48
tr106:
//line NONE:1
te = p+1

//line machine.rl:54
act = 13;
	goto st48
tr108:
//line NONE:1
te = p+1

//line machine.rl:55
act = 14;
	goto st48
tr109:
//line NONE:1
te = p+1

//line machine.rl:56
act = 15;
	goto st48
tr111:
//line NONE:1
te = p+1

//line machine.rl:57
act = 16;
	goto st48
tr114:
//line NONE:1
te = p+1

//line machine.rl:58
act = 17;
	goto st48
tr117:
//line NONE:1
te = p+1

//line machine.rl:59
act = 18;
	goto st48
	st48:
		if p++; p == pe {
			goto _test_eof48
		}
	st_case_48:
//line machine.go:1755
		switch data[p] {
		case 39:
			goto st42
		case 43:
			goto st13
		case 46:
			goto st13
		case 47:
			goto st5
		case 58:
			goto st14
		case 95:
			goto tr54
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr52
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr52
			}
		default:
			goto tr52
		}
		goto tr5
	st13:
		if p++; p == pe {
			goto _test_eof13
		}
	st_case_13:
		switch data[p] {
		case 43:
			goto st13
		case 47:
			goto st5
		case 58:
			goto st14
		case 95:
			goto st4
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto st13
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st13
			}
		default:
			goto st13
		}
		goto tr5
	st14:
		if p++; p == pe {
			goto _test_eof14
		}
	st_case_14:
		switch data[p] {
		case 33:
			goto st49
		case 61:
			goto st49
		case 95:
			goto st49
		case 126:
			goto st49
		}
		switch {
		case data[p] < 42:
			if 36 <= data[p] && data[p] <= 39 {
				goto st49
			}
		case data[p] > 58:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st49
				}
			case data[p] >= 63:
				goto st49
			}
		default:
			goto st49
		}
		goto tr5
	st49:
		if p++; p == pe {
			goto _test_eof49
		}
	st_case_49:
		switch data[p] {
		case 33:
			goto st49
		case 61:
			goto st49
		case 95:
			goto st49
		case 126:
			goto st49
		}
		switch {
		case data[p] < 42:
			if 36 <= data[p] && data[p] <= 39 {
				goto st49
			}
		case data[p] > 58:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st49
				}
			case data[p] >= 63:
				goto st49
			}
		default:
			goto st49
		}
		goto tr90
tr55:
//line NONE:1
te = p+1

//line machine.rl:71
act = 29;
	goto st50
	st50:
		if p++; p == pe {
			goto _test_eof50
		}
	st_case_50:
//line machine.go:1889
		switch data[p] {
		case 39:
			goto st42
		case 43:
			goto st13
		case 46:
			goto st13
		case 47:
			goto st5
		case 58:
			goto st14
		case 95:
			goto tr54
		case 115:
			goto tr91
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr52
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr52
			}
		default:
			goto tr52
		}
		goto tr86
tr91:
//line NONE:1
te = p+1

//line machine.rl:71
act = 29;
	goto st51
	st51:
		if p++; p == pe {
			goto _test_eof51
		}
	st_case_51:
//line machine.go:1931
		switch data[p] {
		case 39:
			goto st42
		case 43:
			goto st13
		case 46:
			goto st13
		case 47:
			goto st5
		case 58:
			goto st14
		case 95:
			goto tr54
		case 115:
			goto tr92
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr52
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr52
			}
		default:
			goto tr52
		}
		goto tr86
tr92:
//line NONE:1
te = p+1

//line machine.rl:71
act = 29;
	goto st52
	st52:
		if p++; p == pe {
			goto _test_eof52
		}
	st_case_52:
//line machine.go:1973
		switch data[p] {
		case 39:
			goto st42
		case 43:
			goto st13
		case 46:
			goto st13
		case 47:
			goto st5
		case 58:
			goto st14
		case 95:
			goto tr54
		case 101:
			goto tr93
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr52
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr52
			}
		default:
			goto tr52
		}
		goto tr86
tr93:
//line NONE:1
te = p+1

//line machine.rl:71
act = 29;
	goto st53
	st53:
		if p++; p == pe {
			goto _test_eof53
		}
	st_case_53:
//line machine.go:2015
		switch data[p] {
		case 39:
			goto st42
		case 43:
			goto st13
		case 46:
			goto st13
		case 47:
			goto st5
		case 58:
			goto st14
		case 95:
			goto tr54
		case 114:
			goto tr94
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr52
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr52
			}
		default:
			goto tr52
		}
		goto tr86
tr94:
//line NONE:1
te = p+1

//line machine.rl:71
act = 29;
	goto st54
	st54:
		if p++; p == pe {
			goto _test_eof54
		}
	st_case_54:
//line machine.go:2057
		switch data[p] {
		case 39:
			goto st42
		case 43:
			goto st13
		case 46:
			goto st13
		case 47:
			goto st5
		case 58:
			goto st14
		case 95:
			goto tr54
		case 116:
			goto tr95
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr52
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr52
			}
		default:
			goto tr52
		}
		goto tr86
tr56:
//line NONE:1
te = p+1

//line machine.rl:71
act = 29;
	goto st55
	st55:
		if p++; p == pe {
			goto _test_eof55
		}
	st_case_55:
//line machine.go:2099
		switch data[p] {
		case 39:
			goto st42
		case 43:
			goto st13
		case 46:
			goto st13
		case 47:
			goto st5
		case 58:
			goto st14
		case 95:
			goto tr54
		case 108:
			goto tr96
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr52
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr52
			}
		default:
			goto tr52
		}
		goto tr86
tr96:
//line NONE:1
te = p+1

//line machine.rl:71
act = 29;
	goto st56
	st56:
		if p++; p == pe {
			goto _test_eof56
		}
	st_case_56:
//line machine.go:2141
		switch data[p] {
		case 39:
			goto st42
		case 43:
			goto st13
		case 46:
			goto st13
		case 47:
			goto st5
		case 58:
			goto st14
		case 95:
			goto tr54
		case 115:
			goto tr97
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr52
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr52
			}
		default:
			goto tr52
		}
		goto tr86
tr97:
//line NONE:1
te = p+1

//line machine.rl:71
act = 29;
	goto st57
	st57:
		if p++; p == pe {
			goto _test_eof57
		}
	st_case_57:
//line machine.go:2183
		switch data[p] {
		case 39:
			goto st42
		case 43:
			goto st13
		case 46:
			goto st13
		case 47:
			goto st5
		case 58:
			goto st14
		case 95:
			goto tr54
		case 101:
			goto tr98
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr52
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr52
			}
		default:
			goto tr52
		}
		goto tr86
tr57:
//line NONE:1
te = p+1

//line machine.rl:71
act = 29;
	goto st58
	st58:
		if p++; p == pe {
			goto _test_eof58
		}
	st_case_58:
//line machine.go:2225
		switch data[p] {
		case 39:
			goto st42
		case 43:
			goto st13
		case 46:
			goto st13
		case 47:
			goto st5
		case 58:
			goto st14
		case 95:
			goto tr54
		case 102:
			goto tr99
		case 110:
			goto tr100
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr52
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr52
			}
		default:
			goto tr52
		}
		goto tr86
tr100:
//line NONE:1
te = p+1

//line machine.rl:53
act = 12;
	goto st59
	st59:
		if p++; p == pe {
			goto _test_eof59
		}
	st_case_59:
//line machine.go:2269
		switch data[p] {
		case 39:
			goto st42
		case 43:
			goto st13
		case 46:
			goto st13
		case 47:
			goto st5
		case 58:
			goto st14
		case 95:
			goto tr54
		case 104:
			goto tr102
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr52
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr52
			}
		default:
			goto tr52
		}
		goto tr101
tr102:
//line NONE:1
te = p+1

//line machine.rl:71
act = 29;
	goto st60
	st60:
		if p++; p == pe {
			goto _test_eof60
		}
	st_case_60:
//line machine.go:2311
		switch data[p] {
		case 39:
			goto st42
		case 43:
			goto st13
		case 46:
			goto st13
		case 47:
			goto st5
		case 58:
			goto st14
		case 95:
			goto tr54
		case 101:
			goto tr103
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr52
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr52
			}
		default:
			goto tr52
		}
		goto tr86
tr103:
//line NONE:1
te = p+1

//line machine.rl:71
act = 29;
	goto st61
	st61:
		if p++; p == pe {
			goto _test_eof61
		}
	st_case_61:
//line machine.go:2353
		switch data[p] {
		case 39:
			goto st42
		case 43:
			goto st13
		case 46:
			goto st13
		case 47:
			goto st5
		case 58:
			goto st14
		case 95:
			goto tr54
		case 114:
			goto tr104
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr52
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr52
			}
		default:
			goto tr52
		}
		goto tr86
tr104:
//line NONE:1
te = p+1

//line machine.rl:71
act = 29;
	goto st62
	st62:
		if p++; p == pe {
			goto _test_eof62
		}
	st_case_62:
//line machine.go:2395
		switch data[p] {
		case 39:
			goto st42
		case 43:
			goto st13
		case 46:
			goto st13
		case 47:
			goto st5
		case 58:
			goto st14
		case 95:
			goto tr54
		case 105:
			goto tr105
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr52
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr52
			}
		default:
			goto tr52
		}
		goto tr86
tr105:
//line NONE:1
te = p+1

//line machine.rl:71
act = 29;
	goto st63
	st63:
		if p++; p == pe {
			goto _test_eof63
		}
	st_case_63:
//line machine.go:2437
		switch data[p] {
		case 39:
			goto st42
		case 43:
			goto st13
		case 46:
			goto st13
		case 47:
			goto st5
		case 58:
			goto st14
		case 95:
			goto tr54
		case 116:
			goto tr106
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr52
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr52
			}
		default:
			goto tr52
		}
		goto tr86
tr58:
//line NONE:1
te = p+1

//line machine.rl:71
act = 29;
	goto st64
	st64:
		if p++; p == pe {
			goto _test_eof64
		}
	st_case_64:
//line machine.go:2479
		switch data[p] {
		case 39:
			goto st42
		case 43:
			goto st13
		case 46:
			goto st13
		case 47:
			goto st5
		case 58:
			goto st14
		case 95:
			goto tr54
		case 101:
			goto tr107
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr52
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr52
			}
		default:
			goto tr52
		}
		goto tr86
tr107:
//line NONE:1
te = p+1

//line machine.rl:71
act = 29;
	goto st65
	st65:
		if p++; p == pe {
			goto _test_eof65
		}
	st_case_65:
//line machine.go:2521
		switch data[p] {
		case 39:
			goto st42
		case 43:
			goto st13
		case 46:
			goto st13
		case 47:
			goto st5
		case 58:
			goto st14
		case 95:
			goto tr54
		case 116:
			goto tr108
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr52
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr52
			}
		default:
			goto tr52
		}
		goto tr86
tr59:
//line NONE:1
te = p+1

//line machine.rl:71
act = 29;
	goto st66
	st66:
		if p++; p == pe {
			goto _test_eof66
		}
	st_case_66:
//line machine.go:2563
		switch data[p] {
		case 39:
			goto st42
		case 43:
			goto st13
		case 46:
			goto st13
		case 47:
			goto st5
		case 58:
			goto st14
		case 95:
			goto tr54
		case 114:
			goto tr109
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr52
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr52
			}
		default:
			goto tr52
		}
		goto tr86
tr60:
//line NONE:1
te = p+1

//line machine.rl:71
act = 29;
	goto st67
	st67:
		if p++; p == pe {
			goto _test_eof67
		}
	st_case_67:
//line machine.go:2605
		switch data[p] {
		case 39:
			goto st42
		case 43:
			goto st13
		case 46:
			goto st13
		case 47:
			goto st5
		case 58:
			goto st14
		case 95:
			goto tr54
		case 101:
			goto tr110
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr52
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr52
			}
		default:
			goto tr52
		}
		goto tr86
tr110:
//line NONE:1
te = p+1

//line machine.rl:71
act = 29;
	goto st68
	st68:
		if p++; p == pe {
			goto _test_eof68
		}
	st_case_68:
//line machine.go:2647
		switch data[p] {
		case 39:
			goto st42
		case 43:
			goto st13
		case 46:
			goto st13
		case 47:
			goto st5
		case 58:
			goto st14
		case 95:
			goto tr54
		case 99:
			goto tr111
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr52
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr52
			}
		default:
			goto tr52
		}
		goto tr86
tr61:
//line NONE:1
te = p+1

//line machine.rl:71
act = 29;
	goto st69
	st69:
		if p++; p == pe {
			goto _test_eof69
		}
	st_case_69:
//line machine.go:2689
		switch data[p] {
		case 39:
			goto st42
		case 43:
			goto st13
		case 46:
			goto st13
		case 47:
			goto st5
		case 58:
			goto st14
		case 95:
			goto tr54
		case 104:
			goto tr112
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr52
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr52
			}
		default:
			goto tr52
		}
		goto tr86
tr112:
//line NONE:1
te = p+1

//line machine.rl:71
act = 29;
	goto st70
	st70:
		if p++; p == pe {
			goto _test_eof70
		}
	st_case_70:
//line machine.go:2731
		switch data[p] {
		case 39:
			goto st42
		case 43:
			goto st13
		case 46:
			goto st13
		case 47:
			goto st5
		case 58:
			goto st14
		case 95:
			goto tr54
		case 101:
			goto tr113
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr52
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr52
			}
		default:
			goto tr52
		}
		goto tr86
tr113:
//line NONE:1
te = p+1

//line machine.rl:71
act = 29;
	goto st71
	st71:
		if p++; p == pe {
			goto _test_eof71
		}
	st_case_71:
//line machine.go:2773
		switch data[p] {
		case 39:
			goto st42
		case 43:
			goto st13
		case 46:
			goto st13
		case 47:
			goto st5
		case 58:
			goto st14
		case 95:
			goto tr54
		case 110:
			goto tr114
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr52
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr52
			}
		default:
			goto tr52
		}
		goto tr86
tr62:
//line NONE:1
te = p+1

//line machine.rl:71
act = 29;
	goto st72
	st72:
		if p++; p == pe {
			goto _test_eof72
		}
	st_case_72:
//line machine.go:2815
		switch data[p] {
		case 39:
			goto st42
		case 43:
			goto st13
		case 46:
			goto st13
		case 47:
			goto st5
		case 58:
			goto st14
		case 95:
			goto tr54
		case 105:
			goto tr115
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr52
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr52
			}
		default:
			goto tr52
		}
		goto tr86
tr115:
//line NONE:1
te = p+1

//line machine.rl:71
act = 29;
	goto st73
	st73:
		if p++; p == pe {
			goto _test_eof73
		}
	st_case_73:
//line machine.go:2857
		switch data[p] {
		case 39:
			goto st42
		case 43:
			goto st13
		case 46:
			goto st13
		case 47:
			goto st5
		case 58:
			goto st14
		case 95:
			goto tr54
		case 116:
			goto tr116
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr52
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr52
			}
		default:
			goto tr52
		}
		goto tr86
tr116:
//line NONE:1
te = p+1

//line machine.rl:71
act = 29;
	goto st74
	st74:
		if p++; p == pe {
			goto _test_eof74
		}
	st_case_74:
//line machine.go:2899
		switch data[p] {
		case 39:
			goto st42
		case 43:
			goto st13
		case 46:
			goto st13
		case 47:
			goto st5
		case 58:
			goto st14
		case 95:
			goto tr54
		case 104:
			goto tr117
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto tr52
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr52
			}
		default:
			goto tr52
		}
		goto tr86
	st75:
		if p++; p == pe {
			goto _test_eof75
		}
	st_case_75:
		if data[p] == 124 {
			goto tr118
		}
		goto tr66
tr65:
//line NONE:1
te = p+1

	goto st76
	st76:
		if p++; p == pe {
			goto _test_eof76
		}
	st_case_76:
//line machine.go:2948
		if data[p] == 47 {
			goto st15
		}
		goto tr66
	st15:
		if p++; p == pe {
			goto _test_eof15
		}
	st_case_15:
		switch data[p] {
		case 43:
			goto st77
		case 95:
			goto st77
		}
		switch {
		case data[p] < 48:
			if 45 <= data[p] && data[p] <= 46 {
				goto st77
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st77
				}
			case data[p] >= 65:
				goto st77
			}
		default:
			goto st77
		}
		goto tr0
	st77:
		if p++; p == pe {
			goto _test_eof77
		}
	st_case_77:
		switch data[p] {
		case 43:
			goto st77
		case 47:
			goto st78
		case 95:
			goto st77
		}
		switch {
		case data[p] < 65:
			if 45 <= data[p] && data[p] <= 57 {
				goto st77
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st77
			}
		default:
			goto st77
		}
		goto tr120
	st78:
		if p++; p == pe {
			goto _test_eof78
		}
	st_case_78:
		switch data[p] {
		case 43:
			goto st77
		case 95:
			goto st77
		}
		switch {
		case data[p] < 48:
			if 45 <= data[p] && data[p] <= 46 {
				goto st77
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st77
				}
			case data[p] >= 65:
				goto st77
			}
		default:
			goto st77
		}
		goto tr120
tr22:
//line machine.rl:32
p = (te) - 1
{ tok(text); addLines() }
	goto st79
tr24:
//line NONE:1
	switch act {
	case 0:
	{{goto st0 }}
	case 3:
	{p = (te) - 1
 tok(text); addLines() }
	}
	
	goto st79
tr122:
//line machine.rl:30
te = p+1
{ if !tokleave('"') { return }; {top--; cs = stack[top];goto _again } }
	goto st79
tr125:
//line machine.rl:32
te = p
p--
{ tok(text); addLines() }
	goto st79
tr127:
//line machine.rl:33
te = p
p--
{ tok(text) }
	goto st79
tr128:
//line machine.rl:31
te = p+1
{ tokenter(interp, '}'); { if nostack() { return }; {stack[top] = 79; top++; goto st22 }} }
	goto st79
	st79:
//line NONE:1
ts = 0

//line NONE:1
act = 0

		if p++; p == pe {
			goto _test_eof79
		}
	st_case_79:
//line NONE:1
ts = p

//line machine.go:3089
		switch data[p] {
		case 34:
			goto tr122
		case 36:
			goto st81
		case 92:
			goto st17
		}
		goto tr23
tr23:
//line NONE:1
te = p+1

//line machine.rl:32
act = 3;
	goto st80
	st80:
		if p++; p == pe {
			goto _test_eof80
		}
	st_case_80:
//line machine.go:3111
		switch data[p] {
		case 34:
			goto tr125
		case 36:
			goto st16
		case 92:
			goto st17
		}
		goto tr23
	st16:
		if p++; p == pe {
			goto _test_eof16
		}
	st_case_16:
		switch data[p] {
		case 34:
			goto tr22
		case 123:
			goto tr22
		}
		goto tr23
	st17:
		if p++; p == pe {
			goto _test_eof17
		}
	st_case_17:
		goto tr23
	st81:
		if p++; p == pe {
			goto _test_eof81
		}
	st_case_81:
		switch data[p] {
		case 34:
			goto tr127
		case 123:
			goto tr128
		}
		goto tr23
tr25:
//line machine.rl:42
p = (te) - 1
{ tok(text); addLines() }
	goto st82
tr29:
//line NONE:1
	switch act {
	case 5:
	{p = (te) - 1
 if !tokleave(ii) { return }; {top--; cs = stack[top];goto _again } }
	case 7:
	{p = (te) - 1
 tok(text); addLines() }
	}
	
	goto st82
tr131:
//line machine.rl:42
te = p
p--
{ tok(text); addLines() }
	goto st82
tr134:
//line machine.rl:43
te = p
p--
{ tok(text) }
	goto st82
tr135:
//line machine.rl:41
te = p+1
{ tokenter(interp, '}'); { if nostack() { return }; {stack[top] = 82; top++; goto st22 }} }
	goto st82
tr137:
//line machine.rl:40
te = p
p--
{ if !tokleave(ii) { return }; {top--; cs = stack[top];goto _again } }
	goto st82
	st82:
//line NONE:1
ts = 0

		if p++; p == pe {
			goto _test_eof82
		}
	st_case_82:
//line NONE:1
ts = p

//line machine.go:3202
		switch data[p] {
		case 36:
			goto st84
		case 39:
			goto st85
		}
		goto tr26
tr26:
//line NONE:1
te = p+1

//line machine.rl:42
act = 7;
	goto st83
	st83:
		if p++; p == pe {
			goto _test_eof83
		}
	st_case_83:
//line machine.go:3222
		switch data[p] {
		case 36:
			goto st18
		case 39:
			goto st19
		}
		goto tr26
	st18:
		if p++; p == pe {
			goto _test_eof18
		}
	st_case_18:
		switch data[p] {
		case 39:
			goto tr25
		case 123:
			goto tr25
		}
		goto tr26
	st19:
		if p++; p == pe {
			goto _test_eof19
		}
	st_case_19:
		switch data[p] {
		case 36:
			goto tr25
		case 39:
			goto st20
		}
		goto tr26
	st20:
		if p++; p == pe {
			goto _test_eof20
		}
	st_case_20:
		switch data[p] {
		case 36:
			goto tr26
		case 39:
			goto tr26
		case 92:
			goto st21
		}
		goto tr25
	st21:
		if p++; p == pe {
			goto _test_eof21
		}
	st_case_21:
		goto tr26
	st84:
		if p++; p == pe {
			goto _test_eof84
		}
	st_case_84:
		switch data[p] {
		case 39:
			goto tr134
		case 123:
			goto tr135
		}
		goto tr26
	st85:
		if p++; p == pe {
			goto _test_eof85
		}
	st_case_85:
		switch data[p] {
		case 36:
			goto tr134
		case 39:
			goto tr136
		}
		goto tr26
tr136:
//line NONE:1
te = p+1

//line machine.rl:40
act = 5;
	goto st86
	st86:
		if p++; p == pe {
			goto _test_eof86
		}
	st_case_86:
//line machine.go:3310
		switch data[p] {
		case 36:
			goto tr26
		case 39:
			goto tr26
		case 92:
			goto st21
		}
		goto tr137
st_case_0:
	st0:
		cs = 0
		goto _out
	st_out:
	_test_eof22: cs = 22; goto _test_eof
	_test_eof23: cs = 23; goto _test_eof
	_test_eof1: cs = 1; goto _test_eof
	_test_eof24: cs = 24; goto _test_eof
	_test_eof2: cs = 2; goto _test_eof
	_test_eof3: cs = 3; goto _test_eof
	_test_eof25: cs = 25; goto _test_eof
	_test_eof26: cs = 26; goto _test_eof
	_test_eof27: cs = 27; goto _test_eof
	_test_eof28: cs = 28; goto _test_eof
	_test_eof29: cs = 29; goto _test_eof
	_test_eof30: cs = 30; goto _test_eof
	_test_eof31: cs = 31; goto _test_eof
	_test_eof32: cs = 32; goto _test_eof
	_test_eof33: cs = 33; goto _test_eof
	_test_eof4: cs = 4; goto _test_eof
	_test_eof5: cs = 5; goto _test_eof
	_test_eof34: cs = 34; goto _test_eof
	_test_eof35: cs = 35; goto _test_eof
	_test_eof36: cs = 36; goto _test_eof
	_test_eof37: cs = 37; goto _test_eof
	_test_eof6: cs = 6; goto _test_eof
	_test_eof38: cs = 38; goto _test_eof
	_test_eof7: cs = 7; goto _test_eof
	_test_eof8: cs = 8; goto _test_eof
	_test_eof39: cs = 39; goto _test_eof
	_test_eof40: cs = 40; goto _test_eof
	_test_eof9: cs = 9; goto _test_eof
	_test_eof10: cs = 10; goto _test_eof
	_test_eof41: cs = 41; goto _test_eof
	_test_eof42: cs = 42; goto _test_eof
	_test_eof43: cs = 43; goto _test_eof
	_test_eof44: cs = 44; goto _test_eof
	_test_eof45: cs = 45; goto _test_eof
	_test_eof11: cs = 11; goto _test_eof
	_test_eof12: cs = 12; goto _test_eof
	_test_eof46: cs = 46; goto _test_eof
	_test_eof47: cs = 47; goto _test_eof
	_test_eof48: cs = 48; goto _test_eof
	_test_eof13: cs = 13; goto _test_eof
	_test_eof14: cs = 14; goto _test_eof
	_test_eof49: cs = 49; goto _test_eof
	_test_eof50: cs = 50; goto _test_eof
	_test_eof51: cs = 51; goto _test_eof
	_test_eof52: cs = 52; goto _test_eof
	_test_eof53: cs = 53; goto _test_eof
	_test_eof54: cs = 54; goto _test_eof
	_test_eof55: cs = 55; goto _test_eof
	_test_eof56: cs = 56; goto _test_eof
	_test_eof57: cs = 57; goto _test_eof
	_test_eof58: cs = 58; goto _test_eof
	_test_eof59: cs = 59; goto _test_eof
	_test_eof60: cs = 60; goto _test_eof
	_test_eof61: cs = 61; goto _test_eof
	_test_eof62: cs = 62; goto _test_eof
	_test_eof63: cs = 63; goto _test_eof
	_test_eof64: cs = 64; goto _test_eof
	_test_eof65: cs = 65; goto _test_eof
	_test_eof66: cs = 66; goto _test_eof
	_test_eof67: cs = 67; goto _test_eof
	_test_eof68: cs = 68; goto _test_eof
	_test_eof69: cs = 69; goto _test_eof
	_test_eof70: cs = 70; goto _test_eof
	_test_eof71: cs = 71; goto _test_eof
	_test_eof72: cs = 72; goto _test_eof
	_test_eof73: cs = 73; goto _test_eof
	_test_eof74: cs = 74; goto _test_eof
	_test_eof75: cs = 75; goto _test_eof
	_test_eof76: cs = 76; goto _test_eof
	_test_eof15: cs = 15; goto _test_eof
	_test_eof77: cs = 77; goto _test_eof
	_test_eof78: cs = 78; goto _test_eof
	_test_eof79: cs = 79; goto _test_eof
	_test_eof80: cs = 80; goto _test_eof
	_test_eof16: cs = 16; goto _test_eof
	_test_eof17: cs = 17; goto _test_eof
	_test_eof81: cs = 81; goto _test_eof
	_test_eof82: cs = 82; goto _test_eof
	_test_eof83: cs = 83; goto _test_eof
	_test_eof18: cs = 18; goto _test_eof
	_test_eof19: cs = 19; goto _test_eof
	_test_eof20: cs = 20; goto _test_eof
	_test_eof21: cs = 21; goto _test_eof
	_test_eof84: cs = 84; goto _test_eof
	_test_eof85: cs = 85; goto _test_eof
	_test_eof86: cs = 86; goto _test_eof

	_test_eof: {}
	if p == eof {
		switch cs {
		case 23:
			goto tr66
		case 1:
			goto tr0
		case 24:
			goto tr68
		case 2:
			goto tr2
		case 3:
			goto tr2
		case 25:
			goto tr68
		case 26:
			goto tr70
		case 27:
			goto tr66
		case 28:
			goto tr72
		case 29:
			goto tr66
		case 30:
			goto tr66
		case 31:
			goto tr66
		case 32:
			goto tr66
		case 33:
			goto tr5
		case 4:
			goto tr5
		case 5:
			goto tr5
		case 34:
			goto tr77
		case 35:
			goto tr77
		case 36:
			goto tr66
		case 37:
			goto tr66
		case 6:
			goto tr0
		case 38:
			goto tr68
		case 7:
			goto tr2
		case 8:
			goto tr2
		case 39:
			goto tr68
		case 40:
			goto tr66
		case 9:
			goto tr0
		case 10:
			goto tr0
		case 41:
			goto tr84
		case 42:
			goto tr86
		case 43:
			goto tr86
		case 44:
			goto tr84
		case 45:
			goto tr66
		case 11:
			goto tr0
		case 12:
			goto tr0
		case 46:
			goto tr66
		case 47:
			goto tr66
		case 48:
			goto tr5
		case 13:
			goto tr5
		case 14:
			goto tr5
		case 49:
			goto tr90
		case 50:
			goto tr86
		case 51:
			goto tr86
		case 52:
			goto tr86
		case 53:
			goto tr86
		case 54:
			goto tr86
		case 55:
			goto tr86
		case 56:
			goto tr86
		case 57:
			goto tr86
		case 58:
			goto tr86
		case 59:
			goto tr101
		case 60:
			goto tr86
		case 61:
			goto tr86
		case 62:
			goto tr86
		case 63:
			goto tr86
		case 64:
			goto tr86
		case 65:
			goto tr86
		case 66:
			goto tr86
		case 67:
			goto tr86
		case 68:
			goto tr86
		case 69:
			goto tr86
		case 70:
			goto tr86
		case 71:
			goto tr86
		case 72:
			goto tr86
		case 73:
			goto tr86
		case 74:
			goto tr86
		case 75:
			goto tr66
		case 76:
			goto tr66
		case 15:
			goto tr0
		case 77:
			goto tr120
		case 78:
			goto tr120
		case 80:
			goto tr125
		case 16:
			goto tr22
		case 17:
			goto tr24
		case 81:
			goto tr127
		case 83:
			goto tr131
		case 18:
			goto tr25
		case 19:
			goto tr25
		case 20:
			goto tr25
		case 21:
			goto tr29
		case 84:
			goto tr134
		case 85:
			goto tr134
		case 86:
			goto tr137
		}
	}

	_out: {}
	}

//line machine.rl:170

	if p != eof {
		err = r.Errorf("precedes the token that failed to lex")
	} else if len(backrefs) != 0 {
		iprev, _ := backrefs.Pop()
		prev := r.tokens[iprev]
		err = fmt.Errorf("%s%s is not terminated", r.At(prev.pos), symString(prev.sym))
	}
	return
}
