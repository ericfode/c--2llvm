go build -a github.com/cznic/golex
golex lex.l
go fmt lex.yy.go
go fmt ast.go