package ast

import (
	"go/token"
	"strings"
	"unicode"
	"unicode/utf8"
)

//Interfaces

type Node interface {
	Pos() int
}

type Expr interface {
	Node
	exprNode()
}

type Stmt interface {
	Node
	stmtNode()
}

type Decl interface {
	Node
	declNode()
}

//Statement
//TODO IMPLEMENT SPAN
type (
	BadStmt struct {
		From, To int
	}

	IfStmt struct {
		If   int  //pos of if work
		Cond Expr //condition
		Body *BlockStmt
		Else Stmt
	}

	CaseClause struct {
		Case  int
		List  []Expr
		Colon int // pos of colon in Case: {Stmts}
		Body  []Stmt
	}

	SwitchStmt struct {
		Switch int //pos of switch
		Cond   Expr
		Body   *BlockStmt // this will only hold CaseClauses
	}

	SpanStmt struct {
		Span  int
		Key   Expr
		Value Expr
		Body  []Stmt
	}

	AssignStmt struct {
		Lhs    []Expr
		TokPos int //location
		Rhs    []Expr
	}

	CallStmt struct {
		Pos     int
		Lhs     []Expr
		Conv    String //no idea what this is for
		Func    Expr   //pointer to runtime value
		Params  []Expr
		Targets []Expr
		Body    *BlockStmt
	}

	//continuations not implemented

	//ignoreing continuation thing
	ReturnStmt struct {
		Return  int //pos
		Results []Expr
	}
)

func (s *BadStmt) Pos() int     { return s.From }
func (s *IfStmt) Pos() int      { return s.If }
func (s *CaseClause) Pos() int  { return s.Case }
func (s *SwtichtStmt) Pos() int { return s.Switch }
func (s *SpanStmt) Pos() int    { return s.Span }
func (s *AssignStmt) Pos() int  { return s.TokPos }
func (s *CallStmt) Pos() int    { return s.Pos }
func (s *ReturnStmt) Pos() int  { return s.Return }

func (s *BadStmt) stmtNode()     {}
func (s *IfStmt) stmtNode()      {}
func (s *CaseClause) stmtNode()  {}
func (s *SwtichtStmt) stmtNode() {}
func (s *SpanStmt) stmtNode()    {}
func (s *AssignStmt) stmtNode()  {}
func (s *CallStmt) stmtNode()    {}
func (s *ReturnStmt) stmtNode()  {}

//Dont forget LValue
//----------------------

//Expressions
type (
	BadExpr struct {
		Pos int
	}

	LitExpr struct {
		Pos   int
		Type  Type
		Value String
	}

	NameExpr struct {
		Pos  int
		Name String
	}

	//ignoreing assertions
	MemRefExpr struct {
		Pos  int
		Type Type
		Loc  Expr
	}

	OpExpr struct {
		Pos   int
		Op    String
		Prams []Expr
	}

	NegExpr struct {
		Pos int
		Rhs Expr
	}
	
	LValExpr struct {
		Pos int
		Rhs Expr
	}
)

func (e *BadExpr) Pos() int    { return e.Pos }
func (e *LitExpr) Pos() int    { return e.Pos }
func (e *NameExpr) Pos() int   { return e.Pos }
func (e *MemRefExpr) Pos() int { return e.Pos }
func (e *OpExpr) Pos() int     { return e.Pos }
func (e *NegExpr) Pos() int    { return e.Pos }
func (e *LValExpr) Pos() int    { return e.Pos }
func (e *BadExpr) exprNode()    {}
func (e *LitExpr) exprNode()    {}
func (e *NameExpr) exprNode()   {}
func (e *MemRefExpr) exprNode() {}
func (e *OpExpr) exprNode()     {}
func (e *NegExpr) exprNode()    {}
func (e *LValExpr) exprNode() 	{}

type (
	Type struct {
		Pos      Int
		Width    Int
		PrimType String
	}
)
