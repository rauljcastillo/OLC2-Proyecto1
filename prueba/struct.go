package prueba

import (
	"main/parser"

	"github.com/antlr4-go/antlr/v4"
)

type Visitor struct {
	antlr.ParseTreeVisitor
	Errores []Error
}

func NewVisitor() *Visitor {
	return &Visitor{ParseTreeVisitor: &parser.BaseGramarVisitor{}, Errores: make([]Error, 0)}
}

func (l *Visitor) Visit(tree antlr.ParseTree, entorno *Ambiente) interface{} {
	switch val := tree.(type) {
	case *parser.SContext:
		return l.VisitS(val, entorno)
	case *parser.BlockContext:
		return l.VisitBlock(val, entorno)
	case *parser.ProductionContext:
		return l.VisitProduction(val, entorno)
	case *parser.PrinContext:
		return l.VisitPrint(val, entorno)
	case *parser.LiteralContext:
		return l.VisitLiteral(val)
	case *parser.OpContext:
		return l.VisitOp(val, entorno)
	case *parser.ParenContext:
		return l.VisitParen(val, entorno)
	case *parser.DeclaraContext:
		return l.VisitDeclara(val, entorno)
	case *parser.TipoContext:
		return l.VisitTipo(val)
	case *parser.AccesoContext:
		return l.VisitAcceso(val, entorno)
	case *parser.PifContext:
		return l.VisitPif(val, entorno)
	case *parser.StmtContext:
		return l.VisitStmt(val, entorno)
	case *parser.AsignContext:
		return l.VisitAsign(val, entorno)
	case *parser.PelseContext:
		return l.VisitPelse(val, entorno)
	case *parser.PswitchContext:
		return l.VisitPswitch(val, entorno)
	case *parser.CcaseContext:
		return l.VisitCcase(val, entorno)
	case *parser.PdefaultContext:
		return l.VisitPdefault(val, entorno)
	case *parser.PwhileContext:
		return l.VisitPwhile(val, entorno)
	case *parser.PforContext:
		return l.VisitPfor(val, entorno)
	case *parser.PiforContext:
		return l.VisitPifor(val, entorno)
	}
	return Valor{val: int64(0), tipo: INT}
}

func (l *Visitor) VisitS(ctx *parser.SContext, entorno *Ambiente) interface{} {
	return l.Visit(ctx.Block(), entorno)
}

func (l *Visitor) VisitBlock(ctx *parser.BlockContext, entorno *Ambiente) interface{} {
	valor := ""
	for i := 0; i < ctx.GetChildCount(); i++ {
		if val, ok := l.Visit(ctx.Production(i), entorno).(string); ok {
			valor += val
		}
	}
	return valor
}

func (l *Visitor) VisitProduction(ctx *parser.ProductionContext, entorno *Ambiente) interface{} {
	if ctx.Prin() != nil {
		return l.Visit(ctx.Prin(), entorno)
	} else if ctx.Declara() != nil {
		return l.Visit(ctx.Declara(), entorno)
	} else if ctx.Pif() != nil {
		return l.Visit(ctx.Pif(), entorno)
	} else if ctx.Asign() != nil {
		return l.Visit(ctx.Asign(), entorno)
	} else if ctx.Pswitch() != nil {
		return l.Visit(ctx.Pswitch(), entorno)
	} else if ctx.Pwhile() != nil {
		return l.Visit(ctx.Pwhile(), entorno)
	} else if ctx.Pfor() != nil {
		return l.Visit(ctx.Pfor(), entorno)
	}
	return 0
}
