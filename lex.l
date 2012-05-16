%{
	package main
	import(
		"bufio"
		"log"
		"strconv"
	)
	
	type yylexer struct{
		src *bufio.Reader
		buf []byte
		empty bool
		current byte
	}
	
	func newLexer(src *bufio.Reader) (y *yylexer){
		y = &yylexer{src: src}
		if b, err := src.ReadByte(); err == nil{
			y.current = b
		}
		return
	}
	
	func (y *yylexer) getc() byte{
		if y.current != 0 {
			y.buf = append(y.buf, y.current)
		}
		y.current = 0
		if b, err := y.src.ReadByte(); err == nil {
			y.current = b
		}
		return y.current
	}
	
	func (y *yylexer) Error(e string) {
		log.Fatal(e)
		return
	}
	
	func (y *yylexer) Lex(lval yySymType) int {
		var err error
		c := y.current
		if y.empty {
			c, y.empty = y.getc(), false
		}
%}

%yyc c
%yyn c = y.getc()

decdigit 	[0-9]
hexdigit 	[0-9a-fA-F]
octdigit 	[0-7]
dec 	{decdigit}+
hex 	0[xX]{hexdigit}+
oct 	0{octdigit}+
string 	[.]*
type 	::bits{dec}
exp 	[eE]([\+\-]){dec}
signedint 	[-]{dec}
unsignedint 	{dec}|{oct}|({hex}[uU])|(0[uU])
intlit		{signedint}|{unsignedint}{type}?
stringlit 	\"{string}\"
charlit 	[\'][.|[\\.]][\']
name 	[\.\$\_0-9\@a-zA-Z]+
float 	({dec}[\.]{dec}({exp})?|{dec}{exp}){type}?
comment 	\%[.]*

%%


{intlit}
	lval.intVal, err = strconv.ParseInt(yytext,0 , length(yytext)) 
		if err != nil {
			log.Fatal(err)
		}
		return INT

{charlit}
		lval.charVal = yytext[0]
		return CHAR

{stringlit}
		fmt.Sscanf(yytext,"\"%s\"",lval.strVal)
		return STRING

{comment} 

{name}
		lval.symVal = yytext
		return NAME



section		return SECTION
span		return SPAN
import		return IMPORT 
export		return EXPORT 
const		return CONST
typedef		return TYPEDEF
invariant	return INVARIANT
pragma		return PRAGMA
target		return TARGET
byteorder	return BYTEORDER
memsize		return MEMSIZE
little		return LITTLE
big			return BIG
pointersize	return POINTERSIZE
wordsize	return WORDSIZE
align		return ALIGN
stackdata   return STACKDATA
if			return IF
switch		return SWITCH
jump		return JUMP 
return		return RETURN
continuation	return CONTINUATION
goto		return GOTO
cut			return CUT 
also		return ALSO 
aborts		return ABORTS
never		return NEVER
returns		return RETURNS
reads		return READS
writes		return WRITES	
foreign		return foreign

%%
	y.empty = true
	return int(c)
}
