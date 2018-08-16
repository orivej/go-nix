//line nix.y:2
package parser

import __yyfmt__ "fmt"

//line nix.y:2
//line nix.y:5
type yySymType struct {
	yys   int
	token int
	node  *Node
}

const assert_ = 57346
const if_ = 57347
const then = 57348
const else_ = 57349
const let = 57350
const in = 57351
const with = 57352
const or_ = 57353
const rec = 57354
const inherit = 57355
const ellipsis = 57356
const interp = 57357
const space = 57358
const comment = 57359
const ii = 57360
const uri = 57361
const path = 57362
const float = 57363
const int_ = 57364
const id = 57365
const text = 57366
const argID = 57367
const argBracket = 57368
const impl = 57369
const or = 57370
const and = 57371
const eq = 57372
const neq = 57373
const leq = 57374
const geq = 57375
const update = 57376
const concat = 57377
const negate = 57378

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"assert_",
	"if_",
	"then",
	"else_",
	"let",
	"in",
	"with",
	"or_",
	"rec",
	"inherit",
	"ellipsis",
	"interp",
	"space",
	"comment",
	"ii",
	"uri",
	"path",
	"float",
	"int_",
	"id",
	"text",
	"argID",
	"argBracket",
	"':'",
	"'@'",
	"','",
	"';'",
	"'\"'",
	"'.'",
	"'('",
	"')'",
	"'['",
	"']'",
	"'{'",
	"'}'",
	"'='",
	"impl",
	"or",
	"and",
	"eq",
	"neq",
	"'<'",
	"'>'",
	"leq",
	"geq",
	"update",
	"'!'",
	"'+'",
	"'-'",
	"'*'",
	"'/'",
	"concat",
	"'?'",
	"negate",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 81,
	43, 0,
	44, 0,
	-2, 43,
	-1, 82,
	43, 0,
	44, 0,
	-2, 44,
	-1, 85,
	40, 0,
	-2, 47,
}

const yyPrivate = 57344

const yyLast = 370

var yyAct = [...]int{

	2, 14, 20, 65, 114, 44, 45, 53, 47, 29,
	28, 48, 96, 68, 132, 34, 56, 33, 32, 31,
	30, 29, 28, 126, 60, 31, 30, 29, 28, 46,
	58, 66, 43, 42, 41, 40, 39, 38, 37, 36,
	35, 34, 94, 33, 32, 31, 30, 29, 28, 66,
	89, 107, 92, 59, 63, 62, 138, 102, 113, 115,
	66, 137, 107, 104, 57, 66, 91, 86, 95, 122,
	140, 109, 99, 99, 38, 37, 36, 35, 34, 64,
	33, 32, 31, 30, 29, 28, 139, 110, 111, 112,
	107, 93, 116, 106, 134, 67, 56, 121, 56, 70,
	108, 117, 127, 120, 70, 87, 1, 27, 97, 66,
	66, 124, 55, 98, 128, 69, 131, 66, 130, 133,
	125, 27, 99, 54, 135, 8, 70, 61, 136, 33,
	32, 31, 30, 29, 28, 98, 118, 119, 9, 15,
	142, 143, 100, 141, 66, 130, 4, 5, 51, 52,
	6, 0, 7, 0, 26, 0, 0, 0, 0, 0,
	22, 16, 17, 18, 19, 27, 67, 12, 13, 0,
	70, 70, 0, 21, 101, 23, 0, 24, 27, 25,
	98, 0, 0, 0, 0, 144, 69, 0, 0, 0,
	0, 0, 11, 0, 10, 42, 41, 40, 39, 38,
	37, 36, 35, 34, 0, 33, 32, 31, 30, 29,
	28, 41, 40, 39, 38, 37, 36, 35, 34, 0,
	33, 32, 31, 30, 29, 28, 40, 39, 38, 37,
	36, 35, 34, 0, 33, 32, 31, 30, 29, 28,
	26, 0, 0, 0, 0, 0, 22, 16, 17, 18,
	19, 27, 67, 0, 90, 3, 70, 0, 0, 21,
	0, 23, 0, 24, 27, 25, 49, 50, 0, 0,
	0, 0, 69, 0, 0, 0, 0, 0, 11, 123,
	10, 0, 0, 0, 0, 71, 72, 73, 74, 75,
	76, 77, 78, 79, 80, 81, 82, 83, 84, 85,
	26, 0, 0, 0, 0, 0, 22, 16, 17, 18,
	19, 27, 0, 0, 0, 0, 0, 0, 0, 21,
	26, 23, 0, 24, 103, 25, 22, 16, 17, 18,
	19, 27, 67, 0, 90, 0, 70, 0, 67, 21,
	0, 23, 70, 24, 27, 25, 0, 88, 0, 67,
	27, 90, 69, 70, 0, 0, 0, 129, 69, 105,
	0, 27, 0, 0, 0, 0, 0, 0, 0, 69,
}
var yyPact = [...]int{

	142, -1000, -1000, -8, 142, 142, -1000, 142, -1000, 308,
	228, 228, 121, 98, -1000, 32, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, 142, -1000, -1000, 17, -1000, 84, 228,
	228, 228, 228, 228, 228, 228, 228, 228, 228, 228,
	228, 228, 228, 228, 37, 99, 338, 36, -1000, -1000,
	78, 142, 65, 4, 39, -1000, -44, 84, 111, 156,
	23, 288, 321, -1000, 30, -1000, -1000, -1000, -1000, -1000,
	142, -46, -46, -46, -28, -28, -34, -34, -34, -34,
	-34, 29, 29, 183, 169, 154, 142, 142, 142, 19,
	26, 142, -1000, 98, 109, 98, 142, 58, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, 241, 84, 89, -15,
	-1000, 95, -1000, 142, 327, 142, -1000, -24, 142, 69,
	-1000, -1000, 308, -1000, -1000, -1000, -1000, 142, 31, -1000,
	-1000, 22, 59, -1000, 43, -1000, -1000, -1000, -1000, 142,
	142, 155, -1000, -1000, -1000,
}
var yyPgo = [...]int{

	0, 0, 13, 30, 2, 139, 1, 138, 255, 3,
	50, 127, 29, 4, 125, 7, 123, 106,
}
var yyR1 = [...]int{

	0, 17, 1, 1, 1, 1, 1, 1, 2, 3,
	3, 3, 4, 5, 5, 5, 5, 5, 5, 5,
	5, 5, 5, 5, 6, 6, 6, 7, 7, 8,
	8, 8, 8, 8, 8, 8, 8, 8, 8, 8,
	8, 8, 8, 8, 8, 8, 8, 8, 9, 9,
	9, 9, 10, 10, 11, 11, 12, 12, 12, 12,
	13, 13, 14, 14, 14, 14, 15, 15, 15, 15,
	16, 16,
}
var yyR2 = [...]int{

	0, 1, 1, 4, 6, 4, 4, 1, 3, 0,
	2, 2, 1, 1, 1, 1, 1, 1, 3, 3,
	3, 3, 3, 4, 1, 3, 5, 1, 2, 1,
	2, 3, 3, 3, 3, 3, 3, 2, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 1, 1,
	1, 3, 1, 3, 0, 2, 0, 5, 4, 7,
	0, 2, 3, 5, 7, 7, 0, 1, 1, 3,
	1, 3,
}
var yyChk = [...]int{

	-1000, -17, -1, -8, 4, 5, 8, 10, -14, -7,
	52, 50, 25, 26, -6, -5, 19, 20, 21, 22,
	-4, 31, 18, 33, 35, 37, 12, 23, 56, 55,
	54, 53, 52, 51, 49, 48, 47, 46, 45, 44,
	43, 42, 41, 40, -1, -1, -12, -1, -6, -8,
	-8, 27, 28, -15, -16, 14, -4, 32, -3, -3,
	-1, -11, -12, 37, -10, -9, -4, 11, -2, 31,
	15, -8, -8, -8, -8, -8, -8, -8, -8, -8,
	-8, -8, -8, -8, -8, -8, 30, 6, 9, -10,
	13, 30, -1, 26, 38, 29, 56, -10, 24, -2,
	31, 18, 34, 36, -6, 38, -12, 32, -3, -1,
	-1, -1, -1, 39, -13, 33, -1, -15, 27, 28,
	-15, -1, 11, 38, -9, 31, 38, 7, -1, 30,
	-9, -1, 38, -1, 25, -6, -1, 30, 34, 27,
	27, -13, -1, -1, 30,
}
var yyDef = [...]int{

	0, -2, 1, 2, 0, 0, 56, 0, 7, 29,
	0, 0, 0, 66, 27, 24, 13, 14, 15, 16,
	17, 9, 9, 0, 54, 56, 0, 12, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 28, 30,
	37, 0, 0, 0, 67, 68, 70, 0, 0, 0,
	0, 0, 0, 56, 31, 52, 48, 49, 50, 9,
	0, 32, 33, 34, 35, 36, 38, 39, 40, 41,
	42, -2, -2, 45, 46, -2, 0, 0, 0, 0,
	60, 0, 62, 66, 0, 66, 0, 25, 10, 11,
	18, 19, 20, 21, 55, 22, 0, 0, 0, 0,
	3, 0, 5, 0, 0, 0, 6, 0, 0, 0,
	69, 71, 0, 23, 53, 51, 8, 0, 0, 58,
	61, 0, 0, 63, 0, 26, 4, 57, 60, 0,
	0, 0, 64, 65, 59,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 50, 31, 3, 3, 3, 3, 3,
	33, 34, 53, 51, 29, 52, 32, 54, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 27, 30,
	45, 39, 46, 56, 28, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 35, 3, 36, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 37, 3, 38,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 40, 41, 42, 43, 44,
	47, 48, 49, 55, 57,
}
var yyTok3 = [...]int{
	0,
}

var yyErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	yyDebug        = 0
	yyErrorVerbose = false
)

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyParser interface {
	Parse(yyLexer) int
	Lookahead() int
}

type yyParserImpl struct {
	lval  yySymType
	stack [yyInitialStackSize]yySymType
	char  int
}

func (p *yyParserImpl) Lookahead() int {
	return p.char
}

func yyNewParser() yyParser {
	return &yyParserImpl{}
}

const yyFlag = -1000

func yyTokname(c int) string {
	if c >= 1 && c-1 < len(yyToknames) {
		if yyToknames[c-1] != "" {
			return yyToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yyErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !yyErrorVerbose {
		return "syntax error"
	}

	for _, e := range yyErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + yyTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := yyPact[state]
	for tok := TOKSTART; tok-1 < len(yyToknames); tok++ {
		if n := base + tok; n >= 0 && n < yyLast && yyChk[yyAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if yyDef[state] == -2 {
		i := 0
		for yyExca[i] != -1 || yyExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; yyExca[i] >= 0; i += 2 {
			tok := yyExca[i]
			if tok < TOKSTART || yyExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if yyExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += yyTokname(tok)
	}
	return res
}

func yylex1(lex yyLexer, lval *yySymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		token = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			token = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		token = yyTok3[i+0]
		if token == char {
			token = yyTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(token), uint(char))
	}
	return char, token
}

func yyParse(yylex yyLexer) int {
	return yyNewParser().Parse(yylex)
}

func (yyrcvr *yyParserImpl) Parse(yylex yyLexer) int {
	var yyn int
	var yyVAL yySymType
	var yyDollar []yySymType
	_ = yyDollar // silence set and not used
	yyS := yyrcvr.stack[:]
	p := yylex.(*Parser)

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yyrcvr.char = -1
	yytoken := -1 // yyrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		yystate = -1
		yyrcvr.char = -1
		yytoken = -1
	}()
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yytoken), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = yyPact[yystate]
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yyrcvr.char < 0 {
		yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
	}
	yyn += yytoken
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yytoken { /* valid shift */
		yyrcvr.char = -1
		yytoken = -1
		yyVAL = yyrcvr.lval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = yyDef[yystate]
	if yyn == -2 {
		if yyrcvr.char < 0 {
			yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && yyExca[xi+1] == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = yyExca[xi+0]
			if yyn < 0 || yyn == yytoken {
				break
			}
		}
		yyn = yyExca[xi+1]
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error(yyErrorMessage(yystate, yytoken))
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yytoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = yyPact[yyS[yyp].yys] + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = yyAct[yyn] /* simulate a shift of "error" */
					if yyChk[yystate] == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yytoken))
			}
			if yytoken == yyEofCode {
				goto ret1
			}
			yyrcvr.char = -1
			yytoken = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= yyR2[yyn]
	// yyp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = yyR1[yyn]
	yyg := yyPgo[yyn]
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = yyAct[yyg]
	} else {
		yystate = yyAct[yyj]
		if yyChk[yystate] != -yyn {
			yystate = yyAct[yyg]
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 1:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line nix.y:39
		{
			p.Result = yyDollar[1].node
		}
	case 3:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line nix.y:45
		{
			yyVAL.node = p.NewNode(AssertNode, yyDollar[1].token, yyDollar[3].token).N2(yyDollar[2].node, yyDollar[4].node)
		}
	case 4:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line nix.y:47
		{
			yyVAL.node = p.NewNode(IfNode, yyDollar[1].token, yyDollar[3].token, yyDollar[5].token).N3(yyDollar[2].node, yyDollar[4].node, yyDollar[6].node)
		}
	case 5:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line nix.y:49
		{
			yyVAL.node = p.NewNode(LetNode, yyDollar[1].token, yyDollar[3].token).N2(yyDollar[2].node, yyDollar[4].node)
		}
	case 6:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line nix.y:51
		{
			yyVAL.node = p.NewNode(WithNode, yyDollar[1].token, yyDollar[3].token).N2(yyDollar[2].node, yyDollar[4].node)
		}
	case 8:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line nix.y:57
		{
			yyVAL.node = p.NewNode(InterpNode, yyDollar[1].token, yyDollar[3].token).N1(yyDollar[2].node)
		}
	case 9:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line nix.y:62
		{
			yyVAL.node = p.NewNode(StringNode)
		}
	case 10:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line nix.y:64
		{
			yyVAL.node = yyDollar[1].node.N1(p.NewNode(TextNode, yyDollar[2].token))
		}
	case 11:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line nix.y:66
		{
			yyVAL.node = yyDollar[1].node.N1(yyDollar[2].node)
		}
	case 12:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line nix.y:71
		{
			yyVAL.node = p.NewNode(IDNode, yyDollar[1].token)
		}
	case 13:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line nix.y:76
		{
			yyVAL.node = p.NewNode(URINode, yyDollar[1].token)
		}
	case 14:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line nix.y:78
		{
			yyVAL.node = p.NewNode(PathNode, yyDollar[1].token)
		}
	case 15:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line nix.y:80
		{
			yyVAL.node = p.NewNode(FloatNode, yyDollar[1].token)
		}
	case 16:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line nix.y:82
		{
			yyVAL.node = p.NewNode(IntNode, yyDollar[1].token)
		}
	case 18:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line nix.y:85
		{
			yyVAL.node = yyDollar[2].node.T2(yyDollar[1].token, yyDollar[3].token)
		}
	case 19:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line nix.y:87
		{
			yyVAL.node = yyDollar[2].node.T2(yyDollar[1].token, yyDollar[3].token).SetType(IStringNode)
		}
	case 20:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line nix.y:89
		{
			yyVAL.node = p.NewNode(ParensNode, yyDollar[1].token, yyDollar[3].token).N1(yyDollar[2].node)
		}
	case 21:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line nix.y:91
		{
			yyVAL.node = yyDollar[2].node.T2(yyDollar[1].token, yyDollar[3].token)
		}
	case 22:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line nix.y:93
		{
			yyVAL.node = yyDollar[2].node.T2(yyDollar[1].token, yyDollar[3].token).SetType(SetNode)
		}
	case 23:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line nix.y:95
		{
			yyVAL.node = yyDollar[3].node.T3(yyDollar[1].token, yyDollar[2].token, yyDollar[4].token).SetType(RecSetNode)
		}
	case 25:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line nix.y:101
		{
			yyVAL.node = p.NewNode(SelectNode, yyDollar[2].token).N2(yyDollar[1].node, yyDollar[3].node)
		}
	case 26:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line nix.y:103
		{
			yyVAL.node = p.NewNode(SelectOrNode, yyDollar[2].token, yyDollar[4].token).N3(yyDollar[1].node, yyDollar[3].node, yyDollar[5].node)
		}
	case 28:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line nix.y:109
		{
			yyVAL.node = p.NewNode(ApplyNode).N2(yyDollar[1].node, yyDollar[2].node)
		}
	case 30:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line nix.y:115
		{
			yyVAL.node = p.Op(negate, yyDollar[1].token).N1(yyDollar[2].node)
		}
	case 31:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line nix.y:117
		{
			yyVAL.node = p.Op('?', yyDollar[2].token).N2(yyDollar[1].node, yyDollar[3].node)
		}
	case 32:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line nix.y:119
		{
			yyVAL.node = p.Op(concat, yyDollar[2].token).N2(yyDollar[1].node, yyDollar[3].node)
		}
	case 33:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line nix.y:121
		{
			yyVAL.node = p.Op('/', yyDollar[2].token).N2(yyDollar[1].node, yyDollar[3].node)
		}
	case 34:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line nix.y:123
		{
			yyVAL.node = p.Op('*', yyDollar[2].token).N2(yyDollar[1].node, yyDollar[3].node)
		}
	case 35:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line nix.y:125
		{
			yyVAL.node = p.Op('-', yyDollar[2].token).N2(yyDollar[1].node, yyDollar[3].node)
		}
	case 36:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line nix.y:127
		{
			yyVAL.node = p.Op('+', yyDollar[2].token).N2(yyDollar[1].node, yyDollar[3].node)
		}
	case 37:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line nix.y:129
		{
			yyVAL.node = p.Op('!', yyDollar[1].token).N1(yyDollar[2].node)
		}
	case 38:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line nix.y:131
		{
			yyVAL.node = p.Op(update, yyDollar[2].token).N2(yyDollar[1].node, yyDollar[3].node)
		}
	case 39:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line nix.y:133
		{
			yyVAL.node = p.Op(geq, yyDollar[2].token).N2(yyDollar[1].node, yyDollar[3].node)
		}
	case 40:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line nix.y:135
		{
			yyVAL.node = p.Op(leq, yyDollar[2].token).N2(yyDollar[1].node, yyDollar[3].node)
		}
	case 41:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line nix.y:137
		{
			yyVAL.node = p.Op('>', yyDollar[2].token).N2(yyDollar[1].node, yyDollar[3].node)
		}
	case 42:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line nix.y:139
		{
			yyVAL.node = p.Op('<', yyDollar[2].token).N2(yyDollar[1].node, yyDollar[3].node)
		}
	case 43:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line nix.y:141
		{
			yyVAL.node = p.Op(neq, yyDollar[2].token).N2(yyDollar[1].node, yyDollar[3].node)
		}
	case 44:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line nix.y:143
		{
			yyVAL.node = p.Op(eq, yyDollar[2].token).N2(yyDollar[1].node, yyDollar[3].node)
		}
	case 45:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line nix.y:145
		{
			yyVAL.node = p.Op(and, yyDollar[2].token).N2(yyDollar[1].node, yyDollar[3].node)
		}
	case 46:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line nix.y:147
		{
			yyVAL.node = p.Op(or, yyDollar[2].token).N2(yyDollar[1].node, yyDollar[3].node)
		}
	case 47:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line nix.y:149
		{
			yyVAL.node = p.Op(impl, yyDollar[2].token).N2(yyDollar[1].node, yyDollar[3].node)
		}
	case 49:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line nix.y:155
		{
			yyVAL.node = p.NewNode(IDNode, yyDollar[1].token)
		}
	case 51:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line nix.y:158
		{
			yyVAL.node = yyDollar[2].node.T2(yyDollar[1].token, yyDollar[3].token)
		}
	case 52:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line nix.y:163
		{
			yyVAL.node = p.NewNode(AttrPathNode).N1(yyDollar[1].node)
		}
	case 53:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line nix.y:165
		{
			yyVAL.node = yyDollar[1].node.T1(yyDollar[2].token).N1(yyDollar[3].node)
		}
	case 54:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line nix.y:170
		{
			yyVAL.node = p.NewNode(ListNode)
		}
	case 55:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line nix.y:172
		{
			yyVAL.node = yyDollar[1].node.N1(yyDollar[2].node)
		}
	case 56:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line nix.y:177
		{
			yyVAL.node = p.NewNode(BindsNode)
		}
	case 57:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line nix.y:179
		{
			yyVAL.node = yyDollar[1].node.N1(p.NewNode(BindNode, yyDollar[3].token, yyDollar[5].token).N2(yyDollar[2].node, yyDollar[4].node))
		}
	case 58:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line nix.y:181
		{
			yyVAL.node = yyDollar[1].node.N1(p.NewNode(InheritNode, yyDollar[2].token, yyDollar[4].token).N1(yyDollar[3].node))
		}
	case 59:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line nix.y:183
		{
			yyVAL.node = yyDollar[1].node.N1(p.NewNode(InheritFromNode, yyDollar[2].token, yyDollar[3].token, yyDollar[5].token, yyDollar[7].token).N2(yyDollar[4].node, yyDollar[6].node))
		}
	case 60:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line nix.y:188
		{
			yyVAL.node = p.NewNode(InheritListNode)
		}
	case 61:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line nix.y:190
		{
			yyVAL.node = yyDollar[1].node.N1(yyDollar[2].node)
		}
	case 62:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line nix.y:196
		{
			yyVAL.node = p.NewNode(FunctionNode, yyDollar[2].token).N2(p.NewNode(IDNode, yyDollar[1].token), yyDollar[3].node)
		}
	case 63:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line nix.y:198
		{
			yyVAL.node = p.NewNode(FunctionNode, yyDollar[4].token).N2(yyDollar[2].node.T2(yyDollar[1].token, yyDollar[3].token), yyDollar[5].node)
		}
	case 64:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line nix.y:200
		{
			yyVAL.node = p.NewNode(FunctionNode, yyDollar[2].token, yyDollar[6].token).N3(p.NewNode(IDNode, yyDollar[1].token), yyDollar[4].node.T2(yyDollar[3].token, yyDollar[5].token), yyDollar[7].node)
		}
	case 65:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line nix.y:202
		{
			yyVAL.node = p.NewNode(FunctionNode, yyDollar[4].token, yyDollar[6].token).N3(p.NewNode(IDNode, yyDollar[5].token), yyDollar[2].node.T2(yyDollar[1].token, yyDollar[3].token), yyDollar[7].node)
		}
	case 66:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line nix.y:207
		{
			yyVAL.node = p.NewNode(ArgSetNode)
		}
	case 67:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line nix.y:209
		{
			yyVAL.node = p.NewNode(ArgSetNode).N1(yyDollar[1].node)
		}
	case 68:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line nix.y:211
		{
			yyVAL.node = p.NewNode(ArgSetNode).N1(p.NewNode(ArgNode, yyDollar[1].token))
		}
	case 69:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line nix.y:213
		{
			yyVAL.node = yyDollar[3].node.N1(yyDollar[1].node.T1(yyDollar[2].token))
		}
	case 70:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line nix.y:218
		{
			yyVAL.node = p.NewNode(ArgNode).N1(yyDollar[1].node)
		}
	case 71:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line nix.y:220
		{
			yyVAL.node = p.NewNode(ArgNode, yyDollar[2].token).N2(yyDollar[1].node, yyDollar[3].node)
		}
	}
	goto yystack /* stack new state and value */
}
