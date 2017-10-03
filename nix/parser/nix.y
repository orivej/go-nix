%{
package parser
%}

%union {
token int
node *Node
}

%type <node>
Expression Interp String ID Atom Select Apply Op InterpID AttrPath List Binds InheritList Function ArgSet Arg

%token <token>
assert_ if_ then else_ let in with or_ rec inherit ellipsis interp space comment ii
uri path float int_ id text argID argBracket
':' '@' ',' ';' '"' '.' '(' ')' '[' ']' '{' '}' '='

%nonassoc <token> impl
%left <token> or
%left <token> and
%nonassoc <token> eq neq
%left <token> '<' '>' leq geq
%right <token> update
%nonassoc <token> '!'
%left <token> '+' '-'
%left <token> '*' '/'
%right <token> concat
%nonassoc <token> '?'
%nonassoc <token> negate

/*
p := yylex.(*Parser)
*/

%%

Main
: Expression
{ p.Result = $1 }
;

Expression
: Op
| assert_ Expression ';' Expression
{ $$ = p.NewNode(AssertNode, $1, $3).N2($2, $4) }
| if_ Expression then Expression else_ Expression
{ $$ = p.NewNode(IfNode, $1, $3, $5).N3($2, $4, $6) }
| let Binds in Expression
{ $$ = p.NewNode(LetNode, $1, $3).N2($2, $4) }
| with Expression ';' Expression
{ $$ = p.NewNode(WithNode, $1, $3).N2($2, $4) }
| Function
;

Interp
: interp Expression '}'
{ $$ = p.NewNode(InterpNode, $1, $3).N1($2) }
;

String
:
{ $$ = p.NewNode(StringNode) }
| String text
{ $$ = $1.N1(p.NewNode(TextNode, $2)) }
| String Interp
{ $$ = $1.N1($2) }
;

ID
: id
{ $$ = p.NewNode(IDNode, $1) }
;

Atom
: uri
{ $$ = p.NewNode(URINode, $1) }
| path
{ $$ = p.NewNode(PathNode, $1) }
| float
{ $$ = p.NewNode(FloatNode, $1) }
| int_
{ $$ = p.NewNode(IntNode, $1) }
| ID
| '"' String '"'
{ $$ = $2.T2($1, $3) }
| ii String ii
{ $$ = $2.T2($1, $3).SetType(IStringNode) }
| '(' Expression ')'
{ $$ = p.NewNode(ParensNode, $1, $3).N1($2) }
| '[' List ']'
{ $$ = $2.T2($1, $3) }
| '{' Binds '}'
{ $$ = $2.T2($1, $3).SetType(SetNode) }
| rec '{' Binds '}'
{ $$ = $3.T3($1, $2, $4).SetType(RecSetNode) }
;

Select
: Atom
| Atom '.' AttrPath
{ $$ = p.NewNode(SelectNode, $2).N2($1, $3) }
| Atom '.' AttrPath or_ Select
{ $$ = p.NewNode(SelectOrNode, $2, $4).N3($1, $3, $5) }
;

Apply
: Select
| Apply Select
{ $$ = p.NewNode(ApplyNode).N2($1, $2) }
;

Op
: Apply
| '-' Op %prec negate
{ $$ = p.Op(negate, $1).N1($2) }
| Op '?' AttrPath
{ $$ = p.Op('?', $2).N2($1, $3) }
| Op concat Op
{ $$ = p.Op(concat, $2).N2($1, $3) }
| Op '/' Op
{ $$ = p.Op('/', $2).N2($1, $3) }
| Op '*' Op
{ $$ = p.Op('*', $2).N2($1, $3) }
| Op '-' Op
{ $$ = p.Op('-', $2).N2($1, $3) }
| Op '+' Op
{ $$ = p.Op('+', $2).N2($1, $3) }
| '!' Op
{ $$ = p.Op('!', $1).N1($2) }
| Op update Op
{ $$ = p.Op(update, $2).N2($1, $3) }
| Op geq Op
{ $$ = p.Op(geq, $2).N2($1, $3) }
| Op leq Op
{ $$ = p.Op(leq, $2).N2($1, $3) }
| Op '>' Op
{ $$ = p.Op('>', $2).N2($1, $3) }
| Op '<' Op
{ $$ = p.Op('<', $2).N2($1, $3) }
| Op neq Op
{ $$ = p.Op(neq, $2).N2($1, $3) }
| Op eq Op
{ $$ = p.Op(eq, $2).N2($1, $3) }
| Op and Op
{ $$ = p.Op(and, $2).N2($1, $3) }
| Op or Op
{ $$ = p.Op(or, $2).N2($1, $3) }
| Op impl Op
{ $$ = p.Op(impl, $2).N2($1, $3) }
;

InterpID
: ID
| or_
{ $$ = p.NewNode(IDNode, $1) }
| Interp
| '"' String '"'
{ $$ = $2.T2($1, $3) }
;

AttrPath
: InterpID
{ $$ = p.NewNode(AttrPathNode).N1($1) }
| AttrPath '.' InterpID
{ $$ = $1.T1($2).N1($3) }
;

List
:
{ $$ = p.NewNode(ListNode) }
| List Select
{ $$ = $1.N1($2) }
;

Binds
:
{ $$ = p.NewNode(BindsNode) }
| Binds AttrPath '=' Expression ';'
{ $$ = $1.N1(p.NewNode(BindNode, $3, $5).N2($2, $4)) }
| Binds inherit InheritList ';'
{ $$ = $1.N1(p.NewNode(InheritNode, $2, $4).N1($3)) }
| Binds inherit '(' Expression ')' InheritList ';'
{ $$ = $1.N1(p.NewNode(InheritFromNode, $2, $3, $5, $7).N2($4, $6)) }
;

InheritList
:
{ $$ = p.NewNode(InheritListNode) }
| InheritList InterpID
{ $$ = $1.N1($2) }
;


Function
: argID ':' Expression
{ $$ = p.NewNode(FunctionNode, $2).N2(p.NewNode(IDNode, $1), $3) }
| argBracket ArgSet '}' ':' Expression
{ $$ = p.NewNode(FunctionNode, $4).N2($2.T2($1, $3), $5) }
| argID '@' argBracket ArgSet '}' ':' Expression
{ $$ = p.NewNode(FunctionNode, $2, $6).N3(p.NewNode(IDNode, $1), $4.T2($3, $5), $7) }
| argBracket ArgSet '}' '@' argID ':' Expression
{ $$ = p.NewNode(FunctionNode, $4, $6).N3(p.NewNode(IDNode, $5), $2.T2($1, $3), $7) }
;

ArgSet
:
{ $$ = p.NewNode(ArgSetNode) }
| Arg
{ $$ = p.NewNode(ArgSetNode).N1($1) }
| ellipsis
{ $$ = p.NewNode(ArgSetNode).N1(p.NewNode(ArgNode, $1)) }
| Arg ',' ArgSet
{ $$ = $3.N1($1.T1($2)) }
;

Arg
: ID
{ $$ = p.NewNode(ArgNode).N1($1) }
| ID '?' Expression
{ $$ = p.NewNode(ArgNode, $2).N2($1, $3) }
;
