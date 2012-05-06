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
	}
	
	func (y *yylexer) Lex(lval (yySymType) int {
		var err error
		c := y.current
		if y.empty {
			c, y.empty = y.getc(), false
		}


yystate0:


goto yystart1

goto yystate1 // silence unused label error
yystate1:
 c = y.getc()
yystart1:
switch {
default:
goto yyabort
case  c == '"':
goto yystate2
case  c == '$' ||  c == '.' ||  c >= '@' &&  c <= 'Z' ||  c == '_' ||  c == 'd' ||  c == 'h' ||  c == 'k' ||  c == 'o' ||  c == 'q' ||  c == 'u' ||  c == 'v' ||  c >= 'x' &&  c <= 'z':
goto yystate5
case  c == '%':
goto yystate7
case  c == '-':
goto yystate13
case  c == '0':
goto yystate16
case  c == '\'':
goto yystate8
case  c == 'a':
goto yystate27
case  c == 'b':
goto yystate39
case  c == 'c':
goto yystate50
case  c == 'e':
goto yystate66
case  c == 'f':
goto yystate72
case  c == 'g':
goto yystate79
case  c == 'i':
goto yystate83
case  c == 'j':
goto yystate98
case  c == 'l':
goto yystate102
case  c == 'm':
goto yystate108
case  c == 'n':
goto yystate115
case  c == 'p':
goto yystate120
case  c == 'r':
goto yystate136
case  c == 's':
goto yystate146
case  c == 't':
goto yystate169
case  c == 'w':
goto yystate181
case  c >= '1' &&  c <= '9':
goto yystate18
}

yystate2:
 c = y.getc()
switch {
default:
goto yyabort
case  c == '"':
goto yystate3
case  c == '.':
goto yystate2
}

yystate3:
 c = y.getc()
switch {
default:
goto yyabort
case  c == '\r':
goto yystate4
}

yystate4:
 c = y.getc()
goto yyrule3

yystate5:
 c = y.getc()
switch {
default:
goto yyabort
case  c == '$' ||  c == '.' ||  c >= '0' &&  c <= '9' ||  c >= '@' &&  c <= 'Z' ||  c == '_' ||  c >= 'a' &&  c <= 'z':
goto yystate5
case  c == '\r':
goto yystate6
}

yystate6:
 c = y.getc()
goto yyrule5

yystate7:
 c = y.getc()
switch {
default:
goto yyrule4
case  c == '.':
goto yystate7
}

yystate8:
 c = y.getc()
switch {
default:
goto yyabort
case  c == '.' ||  c == '[' ||  c == '\\' ||  c == '|':
goto yystate9
}

yystate9:
 c = y.getc()
switch {
default:
goto yyabort
case  c == ']':
goto yystate10
}

yystate10:
 c = y.getc()
switch {
default:
goto yyabort
case  c == '\'':
goto yystate11
}

yystate11:
 c = y.getc()
switch {
default:
goto yyabort
case  c == '\r':
goto yystate12
}

yystate12:
 c = y.getc()
goto yyrule2

yystate13:
 c = y.getc()
switch {
default:
goto yyabort
case  c >= '0' &&  c <= '9':
goto yystate14
}

yystate14:
 c = y.getc()
switch {
default:
goto yyabort
case  c == '\r':
goto yystate15
case  c >= '0' &&  c <= '9':
goto yystate14
}

yystate15:
 c = y.getc()
goto yyrule1

yystate16:
 c = y.getc()
switch {
default:
goto yyabort
case  c == '$' ||  c == '.' ||  c >= '@' &&  c <= 'T' ||  c == 'V' ||  c == 'W' ||  c == 'Y' ||  c == 'Z' ||  c == '_' ||  c >= 'a' &&  c <= 't' ||  c == 'v' ||  c == 'w' ||  c == 'y' ||  c == 'z':
goto yystate5
case  c == ':':
goto yystate19
case  c == 'U' ||  c == 'u':
goto yystate24
case  c == 'X' ||  c == 'x':
goto yystate25
case  c == '\r':
goto yystate17
case  c >= '0' &&  c <= '9':
goto yystate18
}

yystate17:
 c = y.getc()
goto yyrule1

yystate18:
 c = y.getc()
switch {
default:
goto yyabort
case  c == '$' ||  c == '.' ||  c >= '@' &&  c <= 'Z' ||  c == '_' ||  c >= 'a' &&  c <= 'z':
goto yystate5
case  c == ':':
goto yystate19
case  c == '\r':
goto yystate17
case  c >= '0' &&  c <= '9':
goto yystate18
}

yystate19:
 c = y.getc()
switch {
default:
goto yyabort
case  c == ':':
goto yystate20
}

yystate20:
 c = y.getc()
switch {
default:
goto yyabort
case  c == 'b':
goto yystate21
}

yystate21:
 c = y.getc()
switch {
default:
goto yyabort
case  c == 'i':
goto yystate22
}

yystate22:
 c = y.getc()
switch {
default:
goto yyabort
case  c == 't':
goto yystate23
}

yystate23:
 c = y.getc()
switch {
default:
goto yyabort
case  c == 's':
goto yystate13
}

yystate24:
 c = y.getc()
switch {
default:
goto yyabort
case  c == '$' ||  c == '.' ||  c >= '0' &&  c <= '9' ||  c >= '@' &&  c <= 'Z' ||  c == '_' ||  c >= 'a' &&  c <= 'z':
goto yystate5
case  c == ':':
goto yystate19
case  c == '\r':
goto yystate17
}

yystate25:
 c = y.getc()
switch {
default:
goto yyabort
case  c == '$' ||  c == '.' ||  c == '@' ||  c >= 'G' &&  c <= 'Z' ||  c == '_' ||  c >= 'g' &&  c <= 'z':
goto yystate5
case  c == '\r':
goto yystate6
case  c >= '0' &&  c <= '9' ||  c >= 'A' &&  c <= 'F' ||  c >= 'a' &&  c <= 'f':
goto yystate26
}

yystate26:
 c = y.getc()
switch {
default:
goto yyabort
case  c == '$' ||  c == '.' ||  c == '@' ||  c >= 'G' &&  c <= 'T' ||  c >= 'V' &&  c <= 'Z' ||  c == '_' ||  c >= 'g' &&  c <= 't' ||  c >= 'v' &&  c <= 'z':
goto yystate5
case  c == 'U' ||  c == 'u':
goto yystate24
case  c == '\r':
goto yystate6
case  c >= '0' &&  c <= '9' ||  c >= 'A' &&  c <= 'F' ||  c >= 'a' &&  c <= 'f':
goto yystate26
}

yystate27:
 c = y.getc()
switch {
default:
goto yyabort
case  c == '$' ||  c == '.' ||  c >= '0' &&  c <= '9' ||  c >= '@' &&  c <= 'Z' ||  c == '_' ||  c == 'a' ||  c >= 'c' &&  c <= 'k' ||  c >= 'm' &&  c <= 'z':
goto yystate5
case  c == '\r':
goto yystate6
case  c == 'b':
goto yystate28
case  c == 'l':
goto yystate33
}

yystate28:
 c = y.getc()
switch {
default:
goto yyabort
case  c == '$' ||  c == '.' ||  c >= '0' &&  c <= '9' ||  c >= '@' &&  c <= 'Z' ||  c == '_' ||  c >= 'a' &&  c <= 'n' ||  c >= 'p' &&  c <= 'z':
goto yystate5
case  c == '\r':
goto yystate6
case  c == 'o':
goto yystate29
}

yystate29:
 c = y.getc()
switch {
default:
goto yyabort
case  c == '$' ||  c == '.' ||  c >= '0' &&  c <= '9' ||  c >= '@' &&  c <= 'Z' ||  c == '_' ||  c >= 'a' &&  c <= 'q' ||  c >= 's' &&  c <= 'z':
goto yystate5
case  c == '\r':
goto yystate6
case  c == 'r':
goto yystate30
}

yystate30:
 c = y.getc()
switch {
default:
goto yyabort
case  c == '$' ||  c == '.' ||  c >= '0' &&  c <= '9' ||  c >= '@' &&  c <= 'Z' ||  c == '_' ||  c >= 'a' &&  c <= 's' ||  c >= 'u' &&  c <= 'z':
goto yystate5
case  c == '\r':
goto yystate6
case  c == 't':
goto yystate31
}

yystate31:
 c = y.getc()
switch {
default:
goto yyabort
case  c == '$' ||  c == '.' ||  c >= '0' &&  c <= '9' ||  c >= '@' &&  c <= 'Z' ||  c == '_' ||  c >= 'a' &&  c <= 'r' ||  c >= 't' &&  c <= 'z':
goto yystate5
case  c == '\r':
goto yystate6
case  c == 's':
goto yystate32
}

yystate32:
 c = y.getc()
switch {
default:
goto yyrule31
case  c == '$' ||  c == '.' ||  c >= '0' &&  c <= '9' ||  c >= '@' &&  c <= 'Z' ||  c == '_' ||  c >= 'a' &&  c <= 'z':
goto yystate5
case  c == '\r':
goto yystate6
}

yystate33:
 c = y.getc()
switch {
default:
goto yyabort
case  c == '$' ||  c == '.' ||  c >= '0' &&  c <= '9' ||  c >= '@' &&  c <= 'Z' ||  c == '_' ||  c >= 'a' &&  c <= 'h' ||  c >= 'j' &&  c <= 'r' ||  c >= 't' &&  c <= 'z':
goto yystate5
case  c == '\r':
goto yystate6
case  c == 'i':
goto yystate34
case  c == 's':
goto yystate37
}

yystate34:
 c = y.getc()
switch {
default:
goto yyabort
case  c == '$' ||  c == '.' ||  c >= '0' &&  c <= '9' ||  c >= '@' &&  c <= 'Z' ||  c == '_' ||  c >= 'a' &&  c <= 'f' ||  c >= 'h' &&  c <= 'z':
goto yystate5
case  c == '\r':
goto yystate6
case  c == 'g':
goto yystate35
}

yystate35:
 c = y.getc()
switch {
default:
goto yyabort
case  c == '$' ||  c == '.' ||  c >= '0' &&  c <= '9' ||  c >= '@' &&  c <= 'Z' ||  c == '_' ||  c >= 'a' &&  c <= 'm' ||  c >= 'o' &&  c <= 'z':
goto yystate5
case  c == '\r':
goto yystate6
case  c == 'n':
goto yystate36
}

yystate36:
 c = y.getc()
switch {
default:
goto yyrule21
case  c == '$' ||  c == '.' ||  c >= '0' &&  c <= '9' ||  c >= '@' &&  c <= 'Z' ||  c == '_' ||  c >= 'a' &&  c <= 'z':
goto yystate5
case  c == '\r':
goto yystate6
}

yystate37:
 c = y.getc()
switch {
default:
goto yyabort
case  c == '$' ||  c == '.' ||  c >= '0' &&  c <= '9' ||  c >= '@' &&  c <= 'Z' ||  c == '_' ||  c >= 'a' &&  c <= 'n' ||  c >= 'p' &&  c <= 'z':
goto yystate5
case  c == '\r':
goto yystate6
case  c == 'o':
goto yystate38
}

yystate38:
 c = y.getc()
switch {
default:
goto yyrule30
case  c == '$' ||  c == '.' ||  c >= '0' &&  c <= '9' ||  c >= '@' &&  c <= 'Z' ||  c == '_' ||  c >= 'a' &&  c <= 'z':
goto yystate5
case  c == '\r':
goto yystate6
}

yystate39:
 c = y.getc()
switch {
default:
goto yyabort
case  c == '$' ||  c == '.' ||  c >= '0' &&  c <= '9' ||  c >= '@' &&  c <= 'Z' ||  c == '_' ||  c >= 'a' &&  c <= 'h' ||  c >= 'j' &&  c <= 'x' ||  c == 'z':
goto yystate5
case  c == '\r':
goto yystate6
case  c == 'i':
goto yystate40
case  c == 'y':
goto yystate42
}

yystate40:
 c = y.getc()
switch {
default:
goto yyabort
case  c == '$' ||  c == '.' ||  c >= '0' &&  c <= '9' ||  c >= '@' &&  c <= 'Z' ||  c == '_' ||  c >= 'a' &&  c <= 'f' ||  c >= 'h' &&  c <= 'z':
goto yystate5
case  c == '\r':
goto yystate6
case  c == 'g':
goto yystate41
}

yystate41:
 c = y.getc()
switch {
default:
goto yyrule18
case  c == '$' ||  c == '.' ||  c >= '0' &&  c <= '9' ||  c >= '@' &&  c <= 'Z' ||  c == '_' ||  c >= 'a' &&  c <= 'z':
goto yystate5
case  c == '\r':
goto yystate6
}

yystate42:
 c = y.getc()
switch {
default:
goto yyabort
case  c == '$' ||  c == '.' ||  c >= '0' &&  c <= '9' ||  c >= '@' &&  c <= 'Z' ||  c == '_' ||  c >= 'a' &&  c <= 's' ||  c >= 'u' &&  c <= 'z':
goto yystate5
case  c == '\r':
goto yystate6
case  c == 't':
goto yystate43
}

yystate43:
 c = y.getc()
switch {
default:
goto yyabort
case  c == '$' ||  c == '.' ||  c >= '0' &&  c <= '9' ||  c >= '@' &&  c <= 'Z' ||  c == '_' ||  c >= 'a' &&  c <= 'd' ||  c >= 'f' &&  c <= 'z':
goto yystate5
case  c == '\r':
goto yystate6
case  c == 'e':
goto yystate44
}

yystate44:
 c = y.getc()
switch {
default:
goto yyabort
case  c == '$' ||  c == '.' ||  c >= '0' &&  c <= '9' ||  c >= '@' &&  c <= 'Z' ||  c == '_' ||  c >= 'a' &&  c <= 'n' ||  c >= 'p' &&  c <= 'z':
goto yystate5
case  c == '\r':
goto yystate6
case  c == 'o':
goto yystate45
}

yystate45:
 c = y.getc()
switch {
default:
goto yyabort
case  c == '$' ||  c == '.' ||  c >= '0' &&  c <= '9' ||  c >= '@' &&  c <= 'Z' ||  c == '_' ||  c >= 'a' &&  c <= 'q' ||  c >= 's' &&  c <= 'z':
goto yystate5
case  c == '\r':
goto yystate6
case  c == 'r':
goto yystate46
}

yystate46:
 c = y.getc()
switch {
default:
goto yyabort
case  c == '$' ||  c == '.' ||  c >= '0' &&  c <= '9' ||  c >= '@' &&  c <= 'Z' ||  c == '_' ||  c >= 'a' &&  c <= 'c' ||  c >= 'e' &&  c <= 'z':
goto yystate5
case  c == '\r':
goto yystate6
case  c == 'd':
goto yystate47
}

yystate47:
 c = y.getc()
switch {
default:
goto yyabort
case  c == '$' ||  c == '.' ||  c >= '0' &&  c <= '9' ||  c >= '@' &&  c <= 'Z' ||  c == '_' ||  c >= 'a' &&  c <= 'd' ||  c >= 'f' &&  c <= 'z':
goto yystate5
case  c == '\r':
goto yystate6
case  c == 'e':
goto yystate48
}

yystate48:
 c = y.getc()
switch {
default:
goto yyabort
case  c == '$' ||  c == '.' ||  c >= '0' &&  c <= '9' ||  c >= '@' &&  c <= 'Z' ||  c == '_' ||  c >= 'a' &&  c <= 'q' ||  c >= 's' &&  c <= 'z':
goto yystate5
case  c == '\r':
goto yystate6
case  c == 'r':
goto yystate49
}

yystate49:
 c = y.getc()
switch {
default:
goto yyrule15
case  c == '$' ||  c == '.' ||  c >= '0' &&  c <= '9' ||  c >= '@' &&  c <= 'Z' ||  c == '_' ||  c >= 'a' &&  c <= 'z':
goto yystate5
case  c == '\r':
goto yystate6
}

yystate50:
 c = y.getc()
switch {
default:
goto yyabort
case  c == '$' ||  c == '.' ||  c >= '0' &&  c <= '9' ||  c >= '@' &&  c <= 'Z' ||  c == '_' ||  c >= 'a' &&  c <= 'n' ||  c >= 'p' &&  c <= 't' ||  c >= 'v' &&  c <= 'z':
goto yystate5
case  c == '\r':
goto yystate6
case  c == 'o':
goto yystate51
case  c == 'u':
goto yystate64
}

yystate51:
 c = y.getc()
switch {
default:
goto yyabort
case  c == '$' ||  c == '.' ||  c >= '0' &&  c <= '9' ||  c >= '@' &&  c <= 'Z' ||  c == '_' ||  c >= 'a' &&  c <= 'm' ||  c >= 'o' &&  c <= 'z':
goto yystate5
case  c == '\r':
goto yystate6
case  c == 'n':
goto yystate52
}

yystate52:
 c = y.getc()
switch {
default:
goto yyabort
case  c == '$' ||  c == '.' ||  c >= '0' &&  c <= '9' ||  c >= '@' &&  c <= 'Z' ||  c == '_' ||  c >= 'a' &&  c <= 'r' ||  c >= 'u' &&  c <= 'z':
goto yystate5
case  c == '\r':
goto yystate6
case  c == 's':
goto yystate53
case  c == 't':
goto yystate55
}

yystate53:
 c = y.getc()
switch {
default:
goto yyabort
case  c == '$' ||  c == '.' ||  c >= '0' &&  c <= '9' ||  c >= '@' &&  c <= 'Z' ||  c == '_' ||  c >= 'a' &&  c <= 's' ||  c >= 'u' &&  c <= 'z':
goto yystate5
case  c == '\r':
goto yystate6
case  c == 't':
goto yystate54
}

yystate54:
 c = y.getc()
switch {
default:
goto yyrule10
case  c == '$' ||  c == '.' ||  c >= '0' &&  c <= '9' ||  c >= '@' &&  c <= 'Z' ||  c == '_' ||  c >= 'a' &&  c <= 'z':
goto yystate5
case  c == '\r':
goto yystate6
}

yystate55:
 c = y.getc()
switch {
default:
goto yyabort
case  c == '$' ||  c == '.' ||  c >= '0' &&  c <= '9' ||  c >= '@' &&  c <= 'Z' ||  c == '_' ||  c >= 'a' &&  c <= 'h' ||  c >= 'j' &&  c <= 'z':
goto yystate5
case  c == '\r':
goto yystate6
case  c == 'i':
goto yystate56
}

yystate56:
 c = y.getc()
switch {
default:
goto yyabort
case  c == '$' ||  c == '.' ||  c >= '0' &&  c <= '9' ||  c >= '@' &&  c <= 'Z' ||  c == '_' ||  c >= 'a' &&  c <= 'm' ||  c >= 'o' &&  c <= 'z':
goto yystate5
case  c == '\r':
goto yystate6
case  c == 'n':
goto yystate57
}

yystate57:
 c = y.getc()
switch {
default:
goto yyabort
case  c == '$' ||  c == '.' ||  c >= '0' &&  c <= '9' ||  c >= '@' &&  c <= 'Z' ||  c == '_' ||  c >= 'a' &&  c <= 't' ||  c >= 'v' &&  c <= 'z':
goto yystate5
case  c == '\r':
goto yystate6
case  c == 'u':
goto yystate58
}

yystate58:
 c = y.getc()
switch {
default:
goto yyabort
case  c == '$' ||  c == '.' ||  c >= '0' &&  c <= '9' ||  c >= '@' &&  c <= 'Z' ||  c == '_' ||  c >= 'b' &&  c <= 'z':
goto yystate5
case  c == '\r':
goto yystate6
case  c == 'a':
goto yystate59
}

yystate59:
 c = y.getc()
switch {
default:
goto yyabort
case  c == '$' ||  c == '.' ||  c >= '0' &&  c <= '9' ||  c >= '@' &&  c <= 'Z' ||  c == '_' ||  c >= 'a' &&  c <= 's' ||  c >= 'u' &&  c <= 'z':
goto yystate5
case  c == '\r':
goto yystate6
case  c == 't':
goto yystate60
}

yystate60:
 c = y.getc()
switch {
default:
goto yyabort
case  c == '$' ||  c == '.' ||  c >= '0' &&  c <= '9' ||  c >= '@' &&  c <= 'Z' ||  c == '_' ||  c >= 'a' &&  c <= 'h' ||  c >= 'j' &&  c <= 'z':
goto yystate5
case  c == '\r':
goto yystate6
case  c == 'i':
goto yystate61
}

yystate61:
 c = y.getc()
switch {
default:
goto yyabort
case  c == '$' ||  c == '.' ||  c >= '0' &&  c <= '9' ||  c >= '@' &&  c <= 'Z' ||  c == '_' ||  c >= 'a' &&  c <= 'n' ||  c >= 'p' &&  c <= 'z':
goto yystate5
case  c == '\r':
goto yystate6
case  c == 'o':
goto yystate62
}

yystate62:
 c = y.getc()
switch {
default:
goto yyabort
case  c == '$' ||  c == '.' ||  c >= '0' &&  c <= '9' ||  c >= '@' &&  c <= 'Z' ||  c == '_' ||  c >= 'a' &&  c <= 'm' ||  c >= 'o' &&  c <= 'z':
goto yystate5
case  c == '\r':
goto yystate6
case  c == 'n':
goto yystate63
}

yystate63:
 c = y.getc()
switch {
default:
goto yyrule27
case  c == '$' ||  c == '.' ||  c >= '0' &&  c <= '9' ||  c >= '@' &&  c <= 'Z' ||  c == '_' ||  c >= 'a' &&  c <= 'z':
goto yystate5
case  c == '\r':
goto yystate6
}

yystate64:
 c = y.getc()
switch {
default:
goto yyabort
case  c == '$' ||  c == '.' ||  c >= '0' &&  c <= '9' ||  c >= '@' &&  c <= 'Z' ||  c == '_' ||  c >= 'a' &&  c <= 's' ||  c >= 'u' &&  c <= 'z':
goto yystate5
case  c == '\r':
goto yystate6
case  c == 't':
goto yystate65
}

yystate65:
 c = y.getc()
switch {
default:
goto yyrule29
case  c == '$' ||  c == '.' ||  c >= '0' &&  c <= '9' ||  c >= '@' &&  c <= 'Z' ||  c == '_' ||  c >= 'a' &&  c <= 'z':
goto yystate5
case  c == '\r':
goto yystate6
}

yystate66:
 c = y.getc()
switch {
default:
goto yyabort
case  c == '$' ||  c == '.' ||  c >= '0' &&  c <= '9' ||  c >= '@' &&  c <= 'Z' ||  c == '_' ||  c >= 'a' &&  c <= 'w' ||  c == 'y' ||  c == 'z':
goto yystate5
case  c == '\r':
goto yystate6
case  c == 'x':
goto yystate67
}

yystate67:
 c = y.getc()
switch {
default:
goto yyabort
case  c == '$' ||  c == '.' ||  c >= '0' &&  c <= '9' ||  c >= '@' &&  c <= 'Z' ||  c == '_' ||  c >= 'a' &&  c <= 'o' ||  c >= 'q' &&  c <= 'z':
goto yystate5
case  c == '\r':
goto yystate6
case  c == 'p':
goto yystate68
}

yystate68:
 c = y.getc()
switch {
default:
goto yyabort
case  c == '$' ||  c == '.' ||  c >= '0' &&  c <= '9' ||  c >= '@' &&  c <= 'Z' ||  c == '_' ||  c >= 'a' &&  c <= 'n' ||  c >= 'p' &&  c <= 'z':
goto yystate5
case  c == '\r':
goto yystate6
case  c == 'o':
goto yystate69
}

yystate69:
 c = y.getc()
switch {
default:
goto yyabort
case  c == '$' ||  c == '.' ||  c >= '0' &&  c <= '9' ||  c >= '@' &&  c <= 'Z' ||  c == '_' ||  c >= 'a' &&  c <= 'q' ||  c >= 's' &&  c <= 'z':
goto yystate5
case  c == '\r':
goto yystate6
case  c == 'r':
goto yystate70
}

yystate70:
 c = y.getc()
switch {
default:
goto yyabort
case  c == '$' ||  c == '.' ||  c >= '0' &&  c <= '9' ||  c >= '@' &&  c <= 'Z' ||  c == '_' ||  c >= 'a' &&  c <= 's' ||  c >= 'u' &&  c <= 'z':
goto yystate5
case  c == '\r':
goto yystate6
case  c == 't':
goto yystate71
}

yystate71:
 c = y.getc()
switch {
default:
goto yyrule9
case  c == '$' ||  c == '.' ||  c >= '0' &&  c <= '9' ||  c >= '@' &&  c <= 'Z' ||  c == '_' ||  c >= 'a' &&  c <= 'z':
goto yystate5
case  c == '\r':
goto yystate6
}

yystate72:
 c = y.getc()
switch {
default:
goto yyabort
case  c == '$' ||  c == '.' ||  c >= '0' &&  c <= '9' ||  c >= '@' &&  c <= 'Z' ||  c == '_' ||  c >= 'a' &&  c <= 'n' ||  c >= 'p' &&  c <= 'z':
goto yystate5
case  c == '\r':
goto yystate6
case  c == 'o':
goto yystate73
}

yystate73:
 c = y.getc()
switch {
default:
goto yyabort
case  c == '$' ||  c == '.' ||  c >= '0' &&  c <= '9' ||  c >= '@' &&  c <= 'Z' ||  c == '_' ||  c >= 'a' &&  c <= 'q' ||  c >= 's' &&  c <= 'z':
goto yystate5
case  c == '\r':
goto yystate6
case  c == 'r':
goto yystate74
}

yystate74:
 c = y.getc()
switch {
default:
goto yyabort
case  c == '$' ||  c == '.' ||  c >= '0' &&  c <= '9' ||  c >= '@' &&  c <= 'Z' ||  c == '_' ||  c >= 'a' &&  c <= 'd' ||  c >= 'f' &&  c <= 'z':
goto yystate5
case  c == '\r':
goto yystate6
case  c == 'e':
goto yystate75
}

yystate75:
 c = y.getc()
switch {
default:
goto yyabort
case  c == '$' ||  c == '.' ||  c >= '0' &&  c <= '9' ||  c >= '@' &&  c <= 'Z' ||  c == '_' ||  c >= 'a' &&  c <= 'h' ||  c >= 'j' &&  c <= 'z':
goto yystate5
case  c == '\r':
goto yystate6
case  c == 'i':
goto yystate76
}

yystate76:
 c = y.getc()
switch {
default:
goto yyabort
case  c == '$' ||  c == '.' ||  c >= '0' &&  c <= '9' ||  c >= '@' &&  c <= 'Z' ||  c == '_' ||  c >= 'a' &&  c <= 'f' ||  c >= 'h' &&  c <= 'z':
goto yystate5
case  c == '\r':
goto yystate6
case  c == 'g':
goto yystate77
}

yystate77:
 c = y.getc()
switch {
default:
goto yyabort
case  c == '$' ||  c == '.' ||  c >= '0' &&  c <= '9' ||  c >= '@' &&  c <= 'Z' ||  c == '_' ||  c >= 'a' &&  c <= 'm' ||  c >= 'o' &&  c <= 'z':
goto yystate5
case  c == '\r':
goto yystate6
case  c == 'n':
goto yystate78
}

yystate78:
 c = y.getc()
switch {
default:
goto yyrule36
case  c == '$' ||  c == '.' ||  c >= '0' &&  c <= '9' ||  c >= '@' &&  c <= 'Z' ||  c == '_' ||  c >= 'a' &&  c <= 'z':
goto yystate5
case  c == '\r':
goto yystate6
}

yystate79:
 c = y.getc()
switch {
default:
goto yyabort
case  c == '$' ||  c == '.' ||  c >= '0' &&  c <= '9' ||  c >= '@' &&  c <= 'Z' ||  c == '_' ||  c >= 'a' &&  c <= 'n' ||  c >= 'p' &&  c <= 'z':
goto yystate5
case  c == '\r':
goto yystate6
case  c == 'o':
goto yystate80
}

yystate80:
 c = y.getc()
switch {
default:
goto yyabort
case  c == '$' ||  c == '.' ||  c >= '0' &&  c <= '9' ||  c >= '@' &&  c <= 'Z' ||  c == '_' ||  c >= 'a' &&  c <= 's' ||  c >= 'u' &&  c <= 'z':
goto yystate5
case  c == '\r':
goto yystate6
case  c == 't':
goto yystate81
}

yystate81:
 c = y.getc()
switch {
default:
goto yyabort
case  c == '$' ||  c == '.' ||  c >= '0' &&  c <= '9' ||  c >= '@' &&  c <= 'Z' ||  c == '_' ||  c >= 'a' &&  c <= 'n' ||  c >= 'p' &&  c <= 'z':
goto yystate5
case  c == '\r':
goto yystate6
case  c == 'o':
goto yystate82
}

yystate82:
 c = y.getc()
switch {
default:
goto yyrule28
case  c == '$' ||  c == '.' ||  c >= '0' &&  c <= '9' ||  c >= '@' &&  c <= 'Z' ||  c == '_' ||  c >= 'a' &&  c <= 'z':
goto yystate5
case  c == '\r':
goto yystate6
}

yystate83:
 c = y.getc()
switch {
default:
goto yyabort
case  c == '$' ||  c == '.' ||  c >= '0' &&  c <= '9' ||  c >= '@' &&  c <= 'Z' ||  c == '_' ||  c >= 'a' &&  c <= 'e' ||  c >= 'g' &&  c <= 'l' ||  c >= 'o' &&  c <= 'z':
goto yystate5
case  c == '\r':
goto yystate6
case  c == 'f':
goto yystate84
case  c == 'm':
goto yystate85
case  c == 'n':
goto yystate90
}

yystate84:
 c = y.getc()
switch {
default:
goto yyrule23
case  c == '$' ||  c == '.' ||  c >= '0' &&  c <= '9' ||  c >= '@' &&  c <= 'Z' ||  c == '_' ||  c >= 'a' &&  c <= 'z':
goto yystate5
case  c == '\r':
goto yystate6
}

yystate85:
 c = y.getc()
switch {
default:
goto yyabort
case  c == '$' ||  c == '.' ||  c >= '0' &&  c <= '9' ||  c >= '@' &&  c <= 'Z' ||  c == '_' ||  c >= 'a' &&  c <= 'o' ||  c >= 'q' &&  c <= 'z':
goto yystate5
case  c == '\r':
goto yystate6
case  c == 'p':
goto yystate86
}

yystate86:
 c = y.getc()
switch {
default:
goto yyabort
case  c == '$' ||  c == '.' ||  c >= '0' &&  c <= '9' ||  c >= '@' &&  c <= 'Z' ||  c == '_' ||  c >= 'a' &&  c <= 'n' ||  c >= 'p' &&  c <= 'z':
goto yystate5
case  c == '\r':
goto yystate6
case  c == 'o':
goto yystate87
}

yystate87:
 c = y.getc()
switch {
default:
goto yyabort
case  c == '$' ||  c == '.' ||  c >= '0' &&  c <= '9' ||  c >= '@' &&  c <= 'Z' ||  c == '_' ||  c >= 'a' &&  c <= 'q' ||  c >= 's' &&  c <= 'z':
goto yystate5
case  c == '\r':
goto yystate6
case  c == 'r':
goto yystate88
}

yystate88:
 c = y.getc()
switch {
default:
goto yyabort
case  c == '$' ||  c == '.' ||  c >= '0' &&  c <= '9' ||  c >= '@' &&  c <= 'Z' ||  c == '_' ||  c >= 'a' &&  c <= 's' ||  c >= 'u' &&  c <= 'z':
goto yystate5
case  c == '\r':
goto yystate6
case  c == 't':
goto yystate89
}

yystate89:
 c = y.getc()
switch {
default:
goto yyrule8
case  c == '$' ||  c == '.' ||  c >= '0' &&  c <= '9' ||  c >= '@' &&  c <= 'Z' ||  c == '_' ||  c >= 'a' &&  c <= 'z':
goto yystate5
case  c == '\r':
goto yystate6
}

yystate90:
 c = y.getc()
switch {
default:
goto yyabort
case  c == '$' ||  c == '.' ||  c >= '0' &&  c <= '9' ||  c >= '@' &&  c <= 'Z' ||  c == '_' ||  c >= 'a' &&  c <= 'u' ||  c >= 'w' &&  c <= 'z':
goto yystate5
case  c == '\r':
goto yystate6
case  c == 'v':
goto yystate91
}

yystate91:
 c = y.getc()
switch {
default:
goto yyabort
case  c == '$' ||  c == '.' ||  c >= '0' &&  c <= '9' ||  c >= '@' &&  c <= 'Z' ||  c == '_' ||  c >= 'b' &&  c <= 'z':
goto yystate5
case  c == '\r':
goto yystate6
case  c == 'a':
goto yystate92
}

yystate92:
 c = y.getc()
switch {
default:
goto yyabort
case  c == '$' ||  c == '.' ||  c >= '0' &&  c <= '9' ||  c >= '@' &&  c <= 'Z' ||  c == '_' ||  c >= 'a' &&  c <= 'q' ||  c >= 's' &&  c <= 'z':
goto yystate5
case  c == '\r':
goto yystate6
case  c == 'r':
goto yystate93
}

yystate93:
 c = y.getc()
switch {
default:
goto yyabort
case  c == '$' ||  c == '.' ||  c >= '0' &&  c <= '9' ||  c >= '@' &&  c <= 'Z' ||  c == '_' ||  c >= 'a' &&  c <= 'h' ||  c >= 'j' &&  c <= 'z':
goto yystate5
case  c == '\r':
goto yystate6
case  c == 'i':
goto yystate94
}

yystate94:
 c = y.getc()
switch {
default:
goto yyabort
case  c == '$' ||  c == '.' ||  c >= '0' &&  c <= '9' ||  c >= '@' &&  c <= 'Z' ||  c == '_' ||  c >= 'b' &&  c <= 'z':
goto yystate5
case  c == '\r':
goto yystate6
case  c == 'a':
goto yystate95
}

yystate95:
 c = y.getc()
switch {
default:
goto yyabort
case  c == '$' ||  c == '.' ||  c >= '0' &&  c <= '9' ||  c >= '@' &&  c <= 'Z' ||  c == '_' ||  c >= 'a' &&  c <= 'm' ||  c >= 'o' &&  c <= 'z':
goto yystate5
case  c == '\r':
goto yystate6
case  c == 'n':
goto yystate96
}

yystate96:
 c = y.getc()
switch {
default:
goto yyabort
case  c == '$' ||  c == '.' ||  c >= '0' &&  c <= '9' ||  c >= '@' &&  c <= 'Z' ||  c == '_' ||  c >= 'a' &&  c <= 's' ||  c >= 'u' &&  c <= 'z':
goto yystate5
case  c == '\r':
goto yystate6
case  c == 't':
goto yystate97
}

yystate97:
 c = y.getc()
switch {
default:
goto yyrule12
case  c == '$' ||  c == '.' ||  c >= '0' &&  c <= '9' ||  c >= '@' &&  c <= 'Z' ||  c == '_' ||  c >= 'a' &&  c <= 'z':
goto yystate5
case  c == '\r':
goto yystate6
}

yystate98:
 c = y.getc()
switch {
default:
goto yyabort
case  c == '$' ||  c == '.' ||  c >= '0' &&  c <= '9' ||  c >= '@' &&  c <= 'Z' ||  c == '_' ||  c >= 'a' &&  c <= 't' ||  c >= 'v' &&  c <= 'z':
goto yystate5
case  c == '\r':
goto yystate6
case  c == 'u':
goto yystate99
}

yystate99:
 c = y.getc()
switch {
default:
goto yyabort
case  c == '$' ||  c == '.' ||  c >= '0' &&  c <= '9' ||  c >= '@' &&  c <= 'Z' ||  c == '_' ||  c >= 'a' &&  c <= 'l' ||  c >= 'n' &&  c <= 'z':
goto yystate5
case  c == '\r':
goto yystate6
case  c == 'm':
goto yystate100
}

yystate100:
 c = y.getc()
switch {
default:
goto yyabort
case  c == '$' ||  c == '.' ||  c >= '0' &&  c <= '9' ||  c >= '@' &&  c <= 'Z' ||  c == '_' ||  c >= 'a' &&  c <= 'o' ||  c >= 'q' &&  c <= 'z':
goto yystate5
case  c == '\r':
goto yystate6
case  c == 'p':
goto yystate101
}

yystate101:
 c = y.getc()
switch {
default:
goto yyrule25
case  c == '$' ||  c == '.' ||  c >= '0' &&  c <= '9' ||  c >= '@' &&  c <= 'Z' ||  c == '_' ||  c >= 'a' &&  c <= 'z':
goto yystate5
case  c == '\r':
goto yystate6
}

yystate102:
 c = y.getc()
switch {
default:
goto yyabort
case  c == '$' ||  c == '.' ||  c >= '0' &&  c <= '9' ||  c >= '@' &&  c <= 'Z' ||  c == '_' ||  c >= 'a' &&  c <= 'h' ||  c >= 'j' &&  c <= 'z':
goto yystate5
case  c == '\r':
goto yystate6
case  c == 'i':
goto yystate103
}

yystate103:
 c = y.getc()
switch {
default:
goto yyabort
case  c == '$' ||  c == '.' ||  c >= '0' &&  c <= '9' ||  c >= '@' &&  c <= 'Z' ||  c == '_' ||  c >= 'a' &&  c <= 's' ||  c >= 'u' &&  c <= 'z':
goto yystate5
case  c == '\r':
goto yystate6
case  c == 't':
goto yystate104
}

yystate104:
 c = y.getc()
switch {
default:
goto yyabort
case  c == '$' ||  c == '.' ||  c >= '0' &&  c <= '9' ||  c >= '@' &&  c <= 'Z' ||  c == '_' ||  c >= 'a' &&  c <= 's' ||  c >= 'u' &&  c <= 'z':
goto yystate5
case  c == '\r':
goto yystate6
case  c == 't':
goto yystate105
}

yystate105:
 c = y.getc()
switch {
default:
goto yyabort
case  c == '$' ||  c == '.' ||  c >= '0' &&  c <= '9' ||  c >= '@' &&  c <= 'Z' ||  c == '_' ||  c >= 'a' &&  c <= 'k' ||  c >= 'm' &&  c <= 'z':
goto yystate5
case  c == '\r':
goto yystate6
case  c == 'l':
goto yystate106
}

yystate106:
 c = y.getc()
switch {
default:
goto yyabort
case  c == '$' ||  c == '.' ||  c >= '0' &&  c <= '9' ||  c >= '@' &&  c <= 'Z' ||  c == '_' ||  c >= 'a' &&  c <= 'd' ||  c >= 'f' &&  c <= 'z':
goto yystate5
case  c == '\r':
goto yystate6
case  c == 'e':
goto yystate107
}

yystate107:
 c = y.getc()
switch {
default:
goto yyrule17
case  c == '$' ||  c == '.' ||  c >= '0' &&  c <= '9' ||  c >= '@' &&  c <= 'Z' ||  c == '_' ||  c >= 'a' &&  c <= 'z':
goto yystate5
case  c == '\r':
goto yystate6
}

yystate108:
 c = y.getc()
switch {
default:
goto yyabort
case  c == '$' ||  c == '.' ||  c >= '0' &&  c <= '9' ||  c >= '@' &&  c <= 'Z' ||  c == '_' ||  c >= 'a' &&  c <= 'd' ||  c >= 'f' &&  c <= 'z':
goto yystate5
case  c == '\r':
goto yystate6
case  c == 'e':
goto yystate109
}

yystate109:
 c = y.getc()
switch {
default:
goto yyabort
case  c == '$' ||  c == '.' ||  c >= '0' &&  c <= '9' ||  c >= '@' &&  c <= 'Z' ||  c == '_' ||  c >= 'a' &&  c <= 'l' ||  c >= 'n' &&  c <= 'z':
goto yystate5
case  c == '\r':
goto yystate6
case  c == 'm':
goto yystate110
}

yystate110:
 c = y.getc()
switch {
default:
goto yyabort
case  c == '$' ||  c == '.' ||  c >= '0' &&  c <= '9' ||  c >= '@' &&  c <= 'Z' ||  c == '_' ||  c >= 'a' &&  c <= 'r' ||  c >= 't' &&  c <= 'z':
goto yystate5
case  c == '\r':
goto yystate6
case  c == 's':
goto yystate111
}

yystate111:
 c = y.getc()
switch {
default:
goto yyabort
case  c == '$' ||  c == '.' ||  c >= '0' &&  c <= '9' ||  c >= '@' &&  c <= 'Z' ||  c == '_' ||  c >= 'a' &&  c <= 'h' ||  c >= 'j' &&  c <= 'z':
goto yystate5
case  c == '\r':
goto yystate6
case  c == 'i':
goto yystate112
}

yystate112:
 c = y.getc()
switch {
default:
goto yyabort
case  c == '$' ||  c == '.' ||  c >= '0' &&  c <= '9' ||  c >= '@' &&  c <= 'Z' ||  c == '_' ||  c >= 'a' &&  c <= 'y':
goto yystate5
case  c == '\r':
goto yystate6
case  c == 'z':
goto yystate113
}

yystate113:
 c = y.getc()
switch {
default:
goto yyabort
case  c == '$' ||  c == '.' ||  c >= '0' &&  c <= '9' ||  c >= '@' &&  c <= 'Z' ||  c == '_' ||  c >= 'a' &&  c <= 'd' ||  c >= 'f' &&  c <= 'z':
goto yystate5
case  c == '\r':
goto yystate6
case  c == 'e':
goto yystate114
}

yystate114:
 c = y.getc()
switch {
default:
goto yyrule16
case  c == '$' ||  c == '.' ||  c >= '0' &&  c <= '9' ||  c >= '@' &&  c <= 'Z' ||  c == '_' ||  c >= 'a' &&  c <= 'z':
goto yystate5
case  c == '\r':
goto yystate6
}

yystate115:
 c = y.getc()
switch {
default:
goto yyabort
case  c == '$' ||  c == '.' ||  c >= '0' &&  c <= '9' ||  c >= '@' &&  c <= 'Z' ||  c == '_' ||  c >= 'a' &&  c <= 'd' ||  c >= 'f' &&  c <= 'z':
goto yystate5
case  c == '\r':
goto yystate6
case  c == 'e':
goto yystate116
}

yystate116:
 c = y.getc()
switch {
default:
goto yyabort
case  c == '$' ||  c == '.' ||  c >= '0' &&  c <= '9' ||  c >= '@' &&  c <= 'Z' ||  c == '_' ||  c >= 'a' &&  c <= 'u' ||  c >= 'w' &&  c <= 'z':
goto yystate5
case  c == '\r':
goto yystate6
case  c == 'v':
goto yystate117
}

yystate117:
 c = y.getc()
switch {
default:
goto yyabort
case  c == '$' ||  c == '.' ||  c >= '0' &&  c <= '9' ||  c >= '@' &&  c <= 'Z' ||  c == '_' ||  c >= 'a' &&  c <= 'd' ||  c >= 'f' &&  c <= 'z':
goto yystate5
case  c == '\r':
goto yystate6
case  c == 'e':
goto yystate118
}

yystate118:
 c = y.getc()
switch {
default:
goto yyabort
case  c == '$' ||  c == '.' ||  c >= '0' &&  c <= '9' ||  c >= '@' &&  c <= 'Z' ||  c == '_' ||  c >= 'a' &&  c <= 'q' ||  c >= 's' &&  c <= 'z':
goto yystate5
case  c == '\r':
goto yystate6
case  c == 'r':
goto yystate119
}

yystate119:
 c = y.getc()
switch {
default:
goto yyrule32
case  c == '$' ||  c == '.' ||  c >= '0' &&  c <= '9' ||  c >= '@' &&  c <= 'Z' ||  c == '_' ||  c >= 'a' &&  c <= 'z':
goto yystate5
case  c == '\r':
goto yystate6
}

yystate120:
 c = y.getc()
switch {
default:
goto yyabort
case  c == '$' ||  c == '.' ||  c >= '0' &&  c <= '9' ||  c >= '@' &&  c <= 'Z' ||  c == '_' ||  c >= 'a' &&  c <= 'n' ||  c == 'p' ||  c == 'q' ||  c >= 's' &&  c <= 'z':
goto yystate5
case  c == '\r':
goto yystate6
case  c == 'o':
goto yystate121
case  c == 'r':
goto yystate131
}

yystate121:
 c = y.getc()
switch {
default:
goto yyabort
case  c == '$' ||  c == '.' ||  c >= '0' &&  c <= '9' ||  c >= '@' &&  c <= 'Z' ||  c == '_' ||  c >= 'a' &&  c <= 'h' ||  c >= 'j' &&  c <= 'z':
goto yystate5
case  c == '\r':
goto yystate6
case  c == 'i':
goto yystate122
}

yystate122:
 c = y.getc()
switch {
default:
goto yyabort
case  c == '$' ||  c == '.' ||  c >= '0' &&  c <= '9' ||  c >= '@' &&  c <= 'Z' ||  c == '_' ||  c >= 'a' &&  c <= 'm' ||  c >= 'o' &&  c <= 'z':
goto yystate5
case  c == '\r':
goto yystate6
case  c == 'n':
goto yystate123
}

yystate123:
 c = y.getc()
switch {
default:
goto yyabort
case  c == '$' ||  c == '.' ||  c >= '0' &&  c <= '9' ||  c >= '@' &&  c <= 'Z' ||  c == '_' ||  c >= 'a' &&  c <= 's' ||  c >= 'u' &&  c <= 'z':
goto yystate5
case  c == '\r':
goto yystate6
case  c == 't':
goto yystate124
}

yystate124:
 c = y.getc()
switch {
default:
goto yyabort
case  c == '$' ||  c == '.' ||  c >= '0' &&  c <= '9' ||  c >= '@' &&  c <= 'Z' ||  c == '_' ||  c >= 'a' &&  c <= 'd' ||  c >= 'f' &&  c <= 'z':
goto yystate5
case  c == '\r':
goto yystate6
case  c == 'e':
goto yystate125
}

yystate125:
 c = y.getc()
switch {
default:
goto yyabort
case  c == '$' ||  c == '.' ||  c >= '0' &&  c <= '9' ||  c >= '@' &&  c <= 'Z' ||  c == '_' ||  c >= 'a' &&  c <= 'q' ||  c >= 's' &&  c <= 'z':
goto yystate5
case  c == '\r':
goto yystate6
case  c == 'r':
goto yystate126
}

yystate126:
 c = y.getc()
switch {
default:
goto yyabort
case  c == '$' ||  c == '.' ||  c >= '0' &&  c <= '9' ||  c >= '@' &&  c <= 'Z' ||  c == '_' ||  c >= 'a' &&  c <= 'r' ||  c >= 't' &&  c <= 'z':
goto yystate5
case  c == '\r':
goto yystate6
case  c == 's':
goto yystate127
}

yystate127:
 c = y.getc()
switch {
default:
goto yyabort
case  c == '$' ||  c == '.' ||  c >= '0' &&  c <= '9' ||  c >= '@' &&  c <= 'Z' ||  c == '_' ||  c >= 'a' &&  c <= 'h' ||  c >= 'j' &&  c <= 'z':
goto yystate5
case  c == '\r':
goto yystate6
case  c == 'i':
goto yystate128
}

yystate128:
 c = y.getc()
switch {
default:
goto yyabort
case  c == '$' ||  c == '.' ||  c >= '0' &&  c <= '9' ||  c >= '@' &&  c <= 'Z' ||  c == '_' ||  c >= 'a' &&  c <= 'y':
goto yystate5
case  c == '\r':
goto yystate6
case  c == 'z':
goto yystate129
}

yystate129:
 c = y.getc()
switch {
default:
goto yyabort
case  c == '$' ||  c == '.' ||  c >= '0' &&  c <= '9' ||  c >= '@' &&  c <= 'Z' ||  c == '_' ||  c >= 'a' &&  c <= 'd' ||  c >= 'f' &&  c <= 'z':
goto yystate5
case  c == '\r':
goto yystate6
case  c == 'e':
goto yystate130
}

yystate130:
 c = y.getc()
switch {
default:
goto yyrule19
case  c == '$' ||  c == '.' ||  c >= '0' &&  c <= '9' ||  c >= '@' &&  c <= 'Z' ||  c == '_' ||  c >= 'a' &&  c <= 'z':
goto yystate5
case  c == '\r':
goto yystate6
}

yystate131:
 c = y.getc()
switch {
default:
goto yyabort
case  c == '$' ||  c == '.' ||  c >= '0' &&  c <= '9' ||  c >= '@' &&  c <= 'Z' ||  c == '_' ||  c >= 'b' &&  c <= 'z':
goto yystate5
case  c == '\r':
goto yystate6
case  c == 'a':
goto yystate132
}

yystate132:
 c = y.getc()
switch {
default:
goto yyabort
case  c == '$' ||  c == '.' ||  c >= '0' &&  c <= '9' ||  c >= '@' &&  c <= 'Z' ||  c == '_' ||  c >= 'a' &&  c <= 'f' ||  c >= 'h' &&  c <= 'z':
goto yystate5
case  c == '\r':
goto yystate6
case  c == 'g':
goto yystate133
}

yystate133:
 c = y.getc()
switch {
default:
goto yyabort
case  c == '$' ||  c == '.' ||  c >= '0' &&  c <= '9' ||  c >= '@' &&  c <= 'Z' ||  c == '_' ||  c >= 'a' &&  c <= 'l' ||  c >= 'n' &&  c <= 'z':
goto yystate5
case  c == '\r':
goto yystate6
case  c == 'm':
goto yystate134
}

yystate134:
 c = y.getc()
switch {
default:
goto yyabort
case  c == '$' ||  c == '.' ||  c >= '0' &&  c <= '9' ||  c >= '@' &&  c <= 'Z' ||  c == '_' ||  c >= 'b' &&  c <= 'z':
goto yystate5
case  c == '\r':
goto yystate6
case  c == 'a':
goto yystate135
}

yystate135:
 c = y.getc()
switch {
default:
goto yyrule13
case  c == '$' ||  c == '.' ||  c >= '0' &&  c <= '9' ||  c >= '@' &&  c <= 'Z' ||  c == '_' ||  c >= 'a' &&  c <= 'z':
goto yystate5
case  c == '\r':
goto yystate6
}

yystate136:
 c = y.getc()
switch {
default:
goto yyabort
case  c == '$' ||  c == '.' ||  c >= '0' &&  c <= '9' ||  c >= '@' &&  c <= 'Z' ||  c == '_' ||  c >= 'a' &&  c <= 'd' ||  c >= 'f' &&  c <= 'z':
goto yystate5
case  c == '\r':
goto yystate6
case  c == 'e':
goto yystate137
}

yystate137:
 c = y.getc()
switch {
default:
goto yyabort
case  c == '$' ||  c == '.' ||  c >= '0' &&  c <= '9' ||  c >= '@' &&  c <= 'Z' ||  c == '_' ||  c >= 'b' &&  c <= 's' ||  c >= 'u' &&  c <= 'z':
goto yystate5
case  c == '\r':
goto yystate6
case  c == 'a':
goto yystate138
case  c == 't':
goto yystate141
}

yystate138:
 c = y.getc()
switch {
default:
goto yyabort
case  c == '$' ||  c == '.' ||  c >= '0' &&  c <= '9' ||  c >= '@' &&  c <= 'Z' ||  c == '_' ||  c >= 'a' &&  c <= 'c' ||  c >= 'e' &&  c <= 'z':
goto yystate5
case  c == '\r':
goto yystate6
case  c == 'd':
goto yystate139
}

yystate139:
 c = y.getc()
switch {
default:
goto yyabort
case  c == '$' ||  c == '.' ||  c >= '0' &&  c <= '9' ||  c >= '@' &&  c <= 'Z' ||  c == '_' ||  c >= 'a' &&  c <= 'r' ||  c >= 't' &&  c <= 'z':
goto yystate5
case  c == '\r':
goto yystate6
case  c == 's':
goto yystate140
}

yystate140:
 c = y.getc()
switch {
default:
goto yyrule34
case  c == '$' ||  c == '.' ||  c >= '0' &&  c <= '9' ||  c >= '@' &&  c <= 'Z' ||  c == '_' ||  c >= 'a' &&  c <= 'z':
goto yystate5
case  c == '\r':
goto yystate6
}

yystate141:
 c = y.getc()
switch {
default:
goto yyabort
case  c == '$' ||  c == '.' ||  c >= '0' &&  c <= '9' ||  c >= '@' &&  c <= 'Z' ||  c == '_' ||  c >= 'a' &&  c <= 't' ||  c >= 'v' &&  c <= 'z':
goto yystate5
case  c == '\r':
goto yystate6
case  c == 'u':
goto yystate142
}

yystate142:
 c = y.getc()
switch {
default:
goto yyabort
case  c == '$' ||  c == '.' ||  c >= '0' &&  c <= '9' ||  c >= '@' &&  c <= 'Z' ||  c == '_' ||  c >= 'a' &&  c <= 'q' ||  c >= 's' &&  c <= 'z':
goto yystate5
case  c == '\r':
goto yystate6
case  c == 'r':
goto yystate143
}

yystate143:
 c = y.getc()
switch {
default:
goto yyabort
case  c == '$' ||  c == '.' ||  c >= '0' &&  c <= '9' ||  c >= '@' &&  c <= 'Z' ||  c == '_' ||  c >= 'a' &&  c <= 'm' ||  c >= 'o' &&  c <= 'z':
goto yystate5
case  c == '\r':
goto yystate6
case  c == 'n':
goto yystate144
}

yystate144:
 c = y.getc()
switch {
default:
goto yyrule26
case  c == '$' ||  c == '.' ||  c >= '0' &&  c <= '9' ||  c >= '@' &&  c <= 'Z' ||  c == '_' ||  c >= 'a' &&  c <= 'r' ||  c >= 't' &&  c <= 'z':
goto yystate5
case  c == '\r':
goto yystate6
case  c == 's':
goto yystate145
}

yystate145:
 c = y.getc()
switch {
default:
goto yyrule33
case  c == '$' ||  c == '.' ||  c >= '0' &&  c <= '9' ||  c >= '@' &&  c <= 'Z' ||  c == '_' ||  c >= 'a' &&  c <= 'z':
goto yystate5
case  c == '\r':
goto yystate6
}

yystate146:
 c = y.getc()
switch {
default:
goto yyabort
case  c == '$' ||  c == '.' ||  c >= '0' &&  c <= '9' ||  c >= '@' &&  c <= 'Z' ||  c == '_' ||  c >= 'a' &&  c <= 'd' ||  c >= 'f' &&  c <= 'o' ||  c >= 'q' &&  c <= 's' ||  c == 'u' ||  c == 'v' ||  c >= 'x' &&  c <= 'z':
goto yystate5
case  c == '\r':
goto yystate6
case  c == 'e':
goto yystate147
case  c == 'p':
goto yystate153
case  c == 't':
goto yystate156
case  c == 'w':
goto yystate164
}

yystate147:
 c = y.getc()
switch {
default:
goto yyabort
case  c == '$' ||  c == '.' ||  c >= '0' &&  c <= '9' ||  c >= '@' &&  c <= 'Z' ||  c == '_' ||  c == 'a' ||  c == 'b' ||  c >= 'd' &&  c <= 'z':
goto yystate5
case  c == '\r':
goto yystate6
case  c == 'c':
goto yystate148
}

yystate148:
 c = y.getc()
switch {
default:
goto yyabort
case  c == '$' ||  c == '.' ||  c >= '0' &&  c <= '9' ||  c >= '@' &&  c <= 'Z' ||  c == '_' ||  c >= 'a' &&  c <= 's' ||  c >= 'u' &&  c <= 'z':
goto yystate5
case  c == '\r':
goto yystate6
case  c == 't':
goto yystate149
}

yystate149:
 c = y.getc()
switch {
default:
goto yyabort
case  c == '$' ||  c == '.' ||  c >= '0' &&  c <= '9' ||  c >= '@' &&  c <= 'Z' ||  c == '_' ||  c >= 'a' &&  c <= 'h' ||  c >= 'j' &&  c <= 'z':
goto yystate5
case  c == '\r':
goto yystate6
case  c == 'i':
goto yystate150
}

yystate150:
 c = y.getc()
switch {
default:
goto yyabort
case  c == '$' ||  c == '.' ||  c >= '0' &&  c <= '9' ||  c >= '@' &&  c <= 'Z' ||  c == '_' ||  c >= 'a' &&  c <= 'n' ||  c >= 'p' &&  c <= 'z':
goto yystate5
case  c == '\r':
goto yystate6
case  c == 'o':
goto yystate151
}

yystate151:
 c = y.getc()
switch {
default:
goto yyabort
case  c == '$' ||  c == '.' ||  c >= '0' &&  c <= '9' ||  c >= '@' &&  c <= 'Z' ||  c == '_' ||  c >= 'a' &&  c <= 'm' ||  c >= 'o' &&  c <= 'z':
goto yystate5
case  c == '\r':
goto yystate6
case  c == 'n':
goto yystate152
}

yystate152:
 c = y.getc()
switch {
default:
goto yyrule6
case  c == '$' ||  c == '.' ||  c >= '0' &&  c <= '9' ||  c >= '@' &&  c <= 'Z' ||  c == '_' ||  c >= 'a' &&  c <= 'z':
goto yystate5
case  c == '\r':
goto yystate6
}

yystate153:
 c = y.getc()
switch {
default:
goto yyabort
case  c == '$' ||  c == '.' ||  c >= '0' &&  c <= '9' ||  c >= '@' &&  c <= 'Z' ||  c == '_' ||  c >= 'b' &&  c <= 'z':
goto yystate5
case  c == '\r':
goto yystate6
case  c == 'a':
goto yystate154
}

yystate154:
 c = y.getc()
switch {
default:
goto yyabort
case  c == '$' ||  c == '.' ||  c >= '0' &&  c <= '9' ||  c >= '@' &&  c <= 'Z' ||  c == '_' ||  c >= 'a' &&  c <= 'm' ||  c >= 'o' &&  c <= 'z':
goto yystate5
case  c == '\r':
goto yystate6
case  c == 'n':
goto yystate155
}

yystate155:
 c = y.getc()
switch {
default:
goto yyrule7
case  c == '$' ||  c == '.' ||  c >= '0' &&  c <= '9' ||  c >= '@' &&  c <= 'Z' ||  c == '_' ||  c >= 'a' &&  c <= 'z':
goto yystate5
case  c == '\r':
goto yystate6
}

yystate156:
 c = y.getc()
switch {
default:
goto yyabort
case  c == '$' ||  c == '.' ||  c >= '0' &&  c <= '9' ||  c >= '@' &&  c <= 'Z' ||  c == '_' ||  c >= 'b' &&  c <= 'z':
goto yystate5
case  c == '\r':
goto yystate6
case  c == 'a':
goto yystate157
}

yystate157:
 c = y.getc()
switch {
default:
goto yyabort
case  c == '$' ||  c == '.' ||  c >= '0' &&  c <= '9' ||  c >= '@' &&  c <= 'Z' ||  c == '_' ||  c == 'a' ||  c == 'b' ||  c >= 'd' &&  c <= 'z':
goto yystate5
case  c == '\r':
goto yystate6
case  c == 'c':
goto yystate158
}

yystate158:
 c = y.getc()
switch {
default:
goto yyabort
case  c == '$' ||  c == '.' ||  c >= '0' &&  c <= '9' ||  c >= '@' &&  c <= 'Z' ||  c == '_' ||  c >= 'a' &&  c <= 'j' ||  c >= 'l' &&  c <= 'z':
goto yystate5
case  c == '\r':
goto yystate6
case  c == 'k':
goto yystate159
}

yystate159:
 c = y.getc()
switch {
default:
goto yyabort
case  c == '$' ||  c == '.' ||  c >= '0' &&  c <= '9' ||  c >= '@' &&  c <= 'Z' ||  c == '_' ||  c >= 'a' &&  c <= 'c' ||  c >= 'e' &&  c <= 'z':
goto yystate5
case  c == '\r':
goto yystate6
case  c == 'd':
goto yystate160
}

yystate160:
 c = y.getc()
switch {
default:
goto yyabort
case  c == '$' ||  c == '.' ||  c >= '0' &&  c <= '9' ||  c >= '@' &&  c <= 'Z' ||  c == '_' ||  c >= 'b' &&  c <= 'z':
goto yystate5
case  c == '\r':
goto yystate6
case  c == 'a':
goto yystate161
}

yystate161:
 c = y.getc()
switch {
default:
goto yyabort
case  c == '$' ||  c == '.' ||  c >= '0' &&  c <= '9' ||  c >= '@' &&  c <= 'Z' ||  c == '_' ||  c >= 'a' &&  c <= 's' ||  c >= 'u' &&  c <= 'z':
goto yystate5
case  c == '\r':
goto yystate6
case  c == 't':
goto yystate162
}

yystate162:
 c = y.getc()
switch {
default:
goto yyabort
case  c == '$' ||  c == '.' ||  c >= '0' &&  c <= '9' ||  c >= '@' &&  c <= 'Z' ||  c == '_' ||  c >= 'b' &&  c <= 'z':
goto yystate5
case  c == '\r':
goto yystate6
case  c == 'a':
goto yystate163
}

yystate163:
 c = y.getc()
switch {
default:
goto yyrule22
case  c == '$' ||  c == '.' ||  c >= '0' &&  c <= '9' ||  c >= '@' &&  c <= 'Z' ||  c == '_' ||  c >= 'a' &&  c <= 'z':
goto yystate5
case  c == '\r':
goto yystate6
}

yystate164:
 c = y.getc()
switch {
default:
goto yyabort
case  c == '$' ||  c == '.' ||  c >= '0' &&  c <= '9' ||  c >= '@' &&  c <= 'Z' ||  c == '_' ||  c >= 'a' &&  c <= 'h' ||  c >= 'j' &&  c <= 'z':
goto yystate5
case  c == '\r':
goto yystate6
case  c == 'i':
goto yystate165
}

yystate165:
 c = y.getc()
switch {
default:
goto yyabort
case  c == '$' ||  c == '.' ||  c >= '0' &&  c <= '9' ||  c >= '@' &&  c <= 'Z' ||  c == '_' ||  c >= 'a' &&  c <= 's' ||  c >= 'u' &&  c <= 'z':
goto yystate5
case  c == '\r':
goto yystate6
case  c == 't':
goto yystate166
}

yystate166:
 c = y.getc()
switch {
default:
goto yyabort
case  c == '$' ||  c == '.' ||  c >= '0' &&  c <= '9' ||  c >= '@' &&  c <= 'Z' ||  c == '_' ||  c == 'a' ||  c == 'b' ||  c >= 'd' &&  c <= 'z':
goto yystate5
case  c == '\r':
goto yystate6
case  c == 'c':
goto yystate167
}

yystate167:
 c = y.getc()
switch {
default:
goto yyabort
case  c == '$' ||  c == '.' ||  c >= '0' &&  c <= '9' ||  c >= '@' &&  c <= 'Z' ||  c == '_' ||  c >= 'a' &&  c <= 'g' ||  c >= 'i' &&  c <= 'z':
goto yystate5
case  c == '\r':
goto yystate6
case  c == 'h':
goto yystate168
}

yystate168:
 c = y.getc()
switch {
default:
goto yyrule24
case  c == '$' ||  c == '.' ||  c >= '0' &&  c <= '9' ||  c >= '@' &&  c <= 'Z' ||  c == '_' ||  c >= 'a' &&  c <= 'z':
goto yystate5
case  c == '\r':
goto yystate6
}

yystate169:
 c = y.getc()
switch {
default:
goto yyabort
case  c == '$' ||  c == '.' ||  c >= '0' &&  c <= '9' ||  c >= '@' &&  c <= 'Z' ||  c == '_' ||  c >= 'b' &&  c <= 'x' ||  c == 'z':
goto yystate5
case  c == '\r':
goto yystate6
case  c == 'a':
goto yystate170
case  c == 'y':
goto yystate175
}

yystate170:
 c = y.getc()
switch {
default:
goto yyabort
case  c == '$' ||  c == '.' ||  c >= '0' &&  c <= '9' ||  c >= '@' &&  c <= 'Z' ||  c == '_' ||  c >= 'a' &&  c <= 'q' ||  c >= 's' &&  c <= 'z':
goto yystate5
case  c == '\r':
goto yystate6
case  c == 'r':
goto yystate171
}

yystate171:
 c = y.getc()
switch {
default:
goto yyabort
case  c == '$' ||  c == '.' ||  c >= '0' &&  c <= '9' ||  c >= '@' &&  c <= 'Z' ||  c == '_' ||  c >= 'a' &&  c <= 'f' ||  c >= 'h' &&  c <= 'z':
goto yystate5
case  c == '\r':
goto yystate6
case  c == 'g':
goto yystate172
}

yystate172:
 c = y.getc()
switch {
default:
goto yyabort
case  c == '$' ||  c == '.' ||  c >= '0' &&  c <= '9' ||  c >= '@' &&  c <= 'Z' ||  c == '_' ||  c >= 'a' &&  c <= 'd' ||  c >= 'f' &&  c <= 'z':
goto yystate5
case  c == '\r':
goto yystate6
case  c == 'e':
goto yystate173
}

yystate173:
 c = y.getc()
switch {
default:
goto yyabort
case  c == '$' ||  c == '.' ||  c >= '0' &&  c <= '9' ||  c >= '@' &&  c <= 'Z' ||  c == '_' ||  c >= 'a' &&  c <= 's' ||  c >= 'u' &&  c <= 'z':
goto yystate5
case  c == '\r':
goto yystate6
case  c == 't':
goto yystate174
}

yystate174:
 c = y.getc()
switch {
default:
goto yyrule14
case  c == '$' ||  c == '.' ||  c >= '0' &&  c <= '9' ||  c >= '@' &&  c <= 'Z' ||  c == '_' ||  c >= 'a' &&  c <= 'z':
goto yystate5
case  c == '\r':
goto yystate6
}

yystate175:
 c = y.getc()
switch {
default:
goto yyabort
case  c == '$' ||  c == '.' ||  c >= '0' &&  c <= '9' ||  c >= '@' &&  c <= 'Z' ||  c == '_' ||  c >= 'a' &&  c <= 'o' ||  c >= 'q' &&  c <= 'z':
goto yystate5
case  c == '\r':
goto yystate6
case  c == 'p':
goto yystate176
}

yystate176:
 c = y.getc()
switch {
default:
goto yyabort
case  c == '$' ||  c == '.' ||  c >= '0' &&  c <= '9' ||  c >= '@' &&  c <= 'Z' ||  c == '_' ||  c >= 'a' &&  c <= 'd' ||  c >= 'f' &&  c <= 'z':
goto yystate5
case  c == '\r':
goto yystate6
case  c == 'e':
goto yystate177
}

yystate177:
 c = y.getc()
switch {
default:
goto yyabort
case  c == '$' ||  c == '.' ||  c >= '0' &&  c <= '9' ||  c >= '@' &&  c <= 'Z' ||  c == '_' ||  c >= 'a' &&  c <= 'c' ||  c >= 'e' &&  c <= 'z':
goto yystate5
case  c == '\r':
goto yystate6
case  c == 'd':
goto yystate178
}

yystate178:
 c = y.getc()
switch {
default:
goto yyabort
case  c == '$' ||  c == '.' ||  c >= '0' &&  c <= '9' ||  c >= '@' &&  c <= 'Z' ||  c == '_' ||  c >= 'a' &&  c <= 'd' ||  c >= 'f' &&  c <= 'z':
goto yystate5
case  c == '\r':
goto yystate6
case  c == 'e':
goto yystate179
}

yystate179:
 c = y.getc()
switch {
default:
goto yyabort
case  c == '$' ||  c == '.' ||  c >= '0' &&  c <= '9' ||  c >= '@' &&  c <= 'Z' ||  c == '_' ||  c >= 'a' &&  c <= 'e' ||  c >= 'g' &&  c <= 'z':
goto yystate5
case  c == '\r':
goto yystate6
case  c == 'f':
goto yystate180
}

yystate180:
 c = y.getc()
switch {
default:
goto yyrule11
case  c == '$' ||  c == '.' ||  c >= '0' &&  c <= '9' ||  c >= '@' &&  c <= 'Z' ||  c == '_' ||  c >= 'a' &&  c <= 'z':
goto yystate5
case  c == '\r':
goto yystate6
}

yystate181:
 c = y.getc()
switch {
default:
goto yyabort
case  c == '$' ||  c == '.' ||  c >= '0' &&  c <= '9' ||  c >= '@' &&  c <= 'Z' ||  c == '_' ||  c >= 'a' &&  c <= 'n' ||  c == 'p' ||  c == 'q' ||  c >= 's' &&  c <= 'z':
goto yystate5
case  c == '\r':
goto yystate6
case  c == 'o':
goto yystate182
case  c == 'r':
goto yystate189
}

yystate182:
 c = y.getc()
switch {
default:
goto yyabort
case  c == '$' ||  c == '.' ||  c >= '0' &&  c <= '9' ||  c >= '@' &&  c <= 'Z' ||  c == '_' ||  c >= 'a' &&  c <= 'q' ||  c >= 's' &&  c <= 'z':
goto yystate5
case  c == '\r':
goto yystate6
case  c == 'r':
goto yystate183
}

yystate183:
 c = y.getc()
switch {
default:
goto yyabort
case  c == '$' ||  c == '.' ||  c >= '0' &&  c <= '9' ||  c >= '@' &&  c <= 'Z' ||  c == '_' ||  c >= 'a' &&  c <= 'c' ||  c >= 'e' &&  c <= 'z':
goto yystate5
case  c == '\r':
goto yystate6
case  c == 'd':
goto yystate184
}

yystate184:
 c = y.getc()
switch {
default:
goto yyabort
case  c == '$' ||  c == '.' ||  c >= '0' &&  c <= '9' ||  c >= '@' &&  c <= 'Z' ||  c == '_' ||  c >= 'a' &&  c <= 'r' ||  c >= 't' &&  c <= 'z':
goto yystate5
case  c == '\r':
goto yystate6
case  c == 's':
goto yystate185
}

yystate185:
 c = y.getc()
switch {
default:
goto yyabort
case  c == '$' ||  c == '.' ||  c >= '0' &&  c <= '9' ||  c >= '@' &&  c <= 'Z' ||  c == '_' ||  c >= 'a' &&  c <= 'h' ||  c >= 'j' &&  c <= 'z':
goto yystate5
case  c == '\r':
goto yystate6
case  c == 'i':
goto yystate186
}

yystate186:
 c = y.getc()
switch {
default:
goto yyabort
case  c == '$' ||  c == '.' ||  c >= '0' &&  c <= '9' ||  c >= '@' &&  c <= 'Z' ||  c == '_' ||  c >= 'a' &&  c <= 'y':
goto yystate5
case  c == '\r':
goto yystate6
case  c == 'z':
goto yystate187
}

yystate187:
 c = y.getc()
switch {
default:
goto yyabort
case  c == '$' ||  c == '.' ||  c >= '0' &&  c <= '9' ||  c >= '@' &&  c <= 'Z' ||  c == '_' ||  c >= 'a' &&  c <= 'd' ||  c >= 'f' &&  c <= 'z':
goto yystate5
case  c == '\r':
goto yystate6
case  c == 'e':
goto yystate188
}

yystate188:
 c = y.getc()
switch {
default:
goto yyrule20
case  c == '$' ||  c == '.' ||  c >= '0' &&  c <= '9' ||  c >= '@' &&  c <= 'Z' ||  c == '_' ||  c >= 'a' &&  c <= 'z':
goto yystate5
case  c == '\r':
goto yystate6
}

yystate189:
 c = y.getc()
switch {
default:
goto yyabort
case  c == '$' ||  c == '.' ||  c >= '0' &&  c <= '9' ||  c >= '@' &&  c <= 'Z' ||  c == '_' ||  c >= 'a' &&  c <= 'h' ||  c >= 'j' &&  c <= 'z':
goto yystate5
case  c == '\r':
goto yystate6
case  c == 'i':
goto yystate190
}

yystate190:
 c = y.getc()
switch {
default:
goto yyabort
case  c == '$' ||  c == '.' ||  c >= '0' &&  c <= '9' ||  c >= '@' &&  c <= 'Z' ||  c == '_' ||  c >= 'a' &&  c <= 's' ||  c >= 'u' &&  c <= 'z':
goto yystate5
case  c == '\r':
goto yystate6
case  c == 't':
goto yystate191
}

yystate191:
 c = y.getc()
switch {
default:
goto yyabort
case  c == '$' ||  c == '.' ||  c >= '0' &&  c <= '9' ||  c >= '@' &&  c <= 'Z' ||  c == '_' ||  c >= 'a' &&  c <= 'd' ||  c >= 'f' &&  c <= 'z':
goto yystate5
case  c == '\r':
goto yystate6
case  c == 'e':
goto yystate192
}

yystate192:
 c = y.getc()
switch {
default:
goto yyabort
case  c == '$' ||  c == '.' ||  c >= '0' &&  c <= '9' ||  c >= '@' &&  c <= 'Z' ||  c == '_' ||  c >= 'a' &&  c <= 'r' ||  c >= 't' &&  c <= 'z':
goto yystate5
case  c == '\r':
goto yystate6
case  c == 's':
goto yystate193
}

yystate193:
 c = y.getc()
switch {
default:
goto yyrule35
case  c == '$' ||  c == '.' ||  c >= '0' &&  c <= '9' ||  c >= '@' &&  c <= 'Z' ||  c == '_' ||  c >= 'a' &&  c <= 'z':
goto yystate5
case  c == '\r':
goto yystate6
}

yyrule1: // {intlit}
{

		if lval.intVal, err = strconv.ParseInt(yytext,0 , length(yytext); err != nil {
			log.Fatal(err)
		}
		return INT
}
yyrule2: // {charlit}
{

		lval.charVal = yytext[0]
		return CHAR
}
yyrule3: // {stringlit}
{

		fmt.Sscanf(yytext,"\"%s\"",lval.strVal)
		return STRING
}
yyrule4: // {comment}

goto yystate0
yyrule5: // {name}
{

		lval.symVal = yytext
		return NAME
}
yyrule6: // section
{
	return SECTION
}
yyrule7: // span
{
	return SPAN
}
yyrule8: // import
{
	return IMPORT 
}
yyrule9: // export
{
	return EXPORT 
}
yyrule10: // const
{
	return CONST
}
yyrule11: // typedef
{
	return TYPEDEF
}
yyrule12: // invariant
{
return INVARIANT
}
yyrule13: // pragma
{
	return PRAGMA
}
yyrule14: // target
{
	return TARGET
}
yyrule15: // byteorder
{
return BYTEORDER
}
yyrule16: // memsize
{
	return MEMSIZE
}
yyrule17: // little
{
	return LITTLE
}
yyrule18: // big
{
		return BIG
}
yyrule19: // pointersize
{
return POINTERSIZE
}
yyrule20: // wordsize
{
return WORDSIZE
}
yyrule21: // align
{
	return ALIGN
}
yyrule22: // stackdata
{
  return STACKDATA
}
yyrule23: // if
{
		return IF
}
yyrule24: // switch
{
	return SWITCH
}
yyrule25: // jump
{
	return JUMP 
}
yyrule26: // return
{
	return RETURN
}
yyrule27: // continuation
{
return CONTINUATION
}
yyrule28: // goto
{
	return GOTO
}
yyrule29: // cut
{
		return CUT 
}
yyrule30: // also
{
	return ALSO 
}
yyrule31: // aborts
{
	return ABORTS
}
yyrule32: // never
{
	return NEVER
}
yyrule33: // returns
{
	return RETURNS
}
yyrule34: // reads
{
	return READS
}
yyrule35: // writes
{
	return WRITES	
}
yyrule36: // foreign
{
	return foreign
}
panic("unreachable")

goto yyabort // silence unused label error

yyabort: // no lexem recognized
	y.empty = true
	return int(c)
}
