package prueba

import (
	"fmt"
	"main/parser"
)

func (l *Visitor) VisitPrint(ctx *parser.PrinContext, entorno *Ambiente) interface{} {
	a := l.Visit(ctx.Expr(), entorno).(Valor)
	if a.val != nil {
		return fmt.Sprint(a.val) + "\n"
	}
	return 0
}
