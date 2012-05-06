/*

This file is a modified excerpt from the GNU Bison Manual examples originally found here:
http://www.gnu.org/software/bison/manual/html_node/Infix-Calc.html#Infix-Calc

The Copyright License for the GNU Bison Manual can be found in the "fdl-1.3" file.

*/

/* Infix notation calculator.  */

%{

// +build ignore

package main

import (
    "bufio"
    "fmt"
    "math"
    "os"
)

type CmmNode struct{
	
}

%}

%union{
    intVal int
	strVal string
	charVal char
	symVal string
	floatVal float
//	cmmNode *CmmNode
}
%token SECTION SPAN IMPORT EXPORT CONST TYPEDEF INVARIANT PRAGMA 
%token TARGET BYTEORDER MEMSIZE LITTLE BIG POINTERSIZE WORDSIZE
%token ALIGN STACKDATA IF SWITCH SPAN JUMP RETURN CONTINUATION
%token GOTO CUT TO CASE ALSO ABORTS NEVER RETURNS READS WRITES 
%token TARGETS FOREIGN
%token	<intVal> INT
%token  <symVal> NAME
%token  <charVal> CHAR
%token  <strVal> STRING TYPE
%token  <floatVal> FLOAT
%left	'-' '+'
%left	'*' '/'
%left	NEG     /* negation--unary minus */
%right	'^'     /* exponentiation */

%type  <cmmnode>	NUM, exp

%% /* The grammar follows.  */

unit: '{' toplevel '}' {}
;

toplevel: SECTION STRING '{' sectionszo '}'
		| decl
		| procedure
;

section: decl
		| procedure
		| daturn
		| SPAN expr expr '{' sectionszo '}'
;

sections:
		| section
		| sections section
;

decl: IMPORT importszo ';'
		| EXPORT exportszo ';'
		| CONST constdec ';'
		| TYPEDEF typedec ';'
		| INVARIANT registers ';'
		| registers ';'
		| PRAGMA NAME '{' prama '}'
		| TARGET targettype ';'
;

importszo: 
		| import
		| importzo ',' imports
;

constdec: NAME '=' expr 
		| type NAME '=' expr
;

typedec: type nameszo
;

names: NAME
		| names ',' NAME
;

nameszo:
		| names
;


targettype: MEMSIZE int
		| BYTEORDER '(' byteorderType ')'
		| POINTERSIZE int
		| WORDSIZE int
;

byteorderType: LITTLE
		| BIG
;

import: NAME
		| STRING AS NAME
;

export: NAME
		| NAME AS STRING
;

datum: NAME ':'
		| ALIGN int ';'
		| type ';'
		| type size ';'
		| type init ';'
		| type size init ';'
;

datumzo:
		| datumzo datum
;

init: '{' expr exprszo '}'
		| STRING
		| string16
;

exprszo:
		| exprszo expr
;

registers: kindopt regs
;

reg: type NAME
		| type NAME '=' STRING
;

regs: reg
		| regs ',' reg
;

kindopt:
		| kind
;

hardwarestring: STRING
;

hardwarestringopt:
		| hardwarestring
;

size: '[' expropt ']'
;

expropt: 
		| expr
;

body: decl
		| stackdecl
		| stmt
		| body body
;

procedure: convopt NAME '(' formalsopt ')' '{' body '}'
;

convopt: 
		| conv
;

formalsopt:
		| formals
;

formal: kindopt invariantopt type NAME
;

invariantopt: 
		| INVARIANT
;

actual: kindopt expr
;

kind: STRING
;

formals: formal
		| formals ',' formal
;

actuals: actual
		| actauls ',' actual
;

actualzo:
		| actuals
;

actualzoparn:'(' ')'
		| '(' actualzo ')'
;
			
stackdecl: STACKDATA '{' datanumzo '}'
;

ifstmt: IF expr '{' body '}' elsestmtopt
;

elsestmtopt: 
		| ELSE '{' body '}'
;

swtichstmt: SWITCH expr '{' armzo '}'
;

spanstmt: SPAN expr expr '{' body '}'
;

assigstmt: lvalues = rvalues ';'

lvalues: lvalue
		| lvalues lvalue
;

rvalue: expr
;

rvalues: rvalue
		| rvalues rvalue
;


stmt: ';'
		| ifstmt
		| switchstmt
		| spanstmt
		| assigstmt
		| NAME '=' '%''%' NAME '(' actualzo')' flowzo ';'
		| kindednameseqopt convopt expr '(' actualzo ')' targetsopt floworaliaszo ';'
		| convopt JUMP expr actualzoparn targetsopt ';'
		| convopt RETURN exprbracksopt actualzoparn ';'
		| NAME ':'
		| CONTINUATION NAME '(' kindednamesopt ')' ':'
		| GOTO expr targetsopt ';'
		| CUT TO expr '(' actualopt ')' flowzo ';'
;

kindedname: kindopt NAME

kindednames: kindedname
		| kindednamezo ',' kindedname
;

kindednameszo:
		| kindednames
;

kindednameseqopt:
		| kindednames '='
;

arm: CASE ranges ':' '{' body '}'
;

armzo:
		| arms arm
;

range: expr '.''.' expr
;

ranges: range
		| ranges ',' range
;

rangezo:
		| ranges
;

lvalue: NAME
		| type '['expr assertionsopt ']'
;

flow: ALSO flowctrl TO names
		| ALSO ABORTS
		| NEVER RETURNS
;

flowctrl: CUTS
		| UNWIND
		| RETURNS
;

flows: flow
		| flows flow
;

flowzo :
		| flows
;

alias: aliastype names
;

aliastype: READS
		| WRITES
;

targets: TARGETS names
;

targetsopt:
		| targets
;

expr: INT typeopt
		| FLOAT typeopt
		| '\'' CHAR '\'' typeopt
		| NAME
		| type '[' expr assertionsopt ']'
		| '(' expr ')'
		| expr op expr
		| NEGSUB expr
		| '%' NAME actualzoparn
;

typeopt: 
		| TYPE
;

type: bitsn
		| NAME
;

string16: UNICODE '(' STRING ')'
;

conv: FOREIGN STRING

assertions: ALIGNED INT innamesopt
		| IN names alignedopt
;

innamesopt:
		| IN names
;

alignedopt: ALIGNED INT
;

floworalias: flow
		| alias
;

floworaliaszo:
		| floworalias
		| floworaliaszo floworalias
;

op: '\'' NAME '\'' | PLUS | SUB | MUL | DIV | MOD | BITAND 
		| BITOR | SFTL | SFTR | EQ | NEQ | GT | LT| GEQ | LEQ
;
