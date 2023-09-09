package main

import (
	"fmt"
	"main/parser"
	"main/prueba"

	"github.com/antlr4-go/antlr/v4"
)

type ErrorListener struct {
	*antlr.DefaultErrorListener
	foundError bool
}

func NewError() *ErrorListener {
	return &ErrorListener{}
}

func (s *ErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {

	fmt.Printf("Error sintáctico en la línea %d, columna %d: %s\n", line, column, msg)
	s.foundError = true
}

func main() {
	/*
		exp := `
		func Nuevo(_ a: Int, num2 b: Int) {
			print(a)
			print(b)
		}

		Nuevo(2,num2: 3);
		`
	*/
	expresion := `
	func ackermanPuntosMenos(_ m: Int , _ n:Int)->Int {
		if (m == 0) {
			return n + 1
		} else if (m > 0 && n == 0) {
			return ackermanPuntosMenos(m - 1, 1)
		} else {
			return ackermanPuntosMenos(m - 1, ackermanPuntosMenos(m, n - 1))
		}
	}

	ackermanPuntosMenos(3,1,ackermanPuntosMenos(3,1))
	`

	is := antlr.NewInputStream(expresion)
	lexer := parser.NewGramarLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewGramarParser(stream)
	p.BuildParseTrees = true
	tt := NewError()
	p.RemoveErrorListeners()
	p.AddErrorListener(tt)
	tree := p.S()
	visitor := prueba.NewVisitor()
	if !tt.foundError {
		global := prueba.NewAmbiente(nil)

		visitor.Visit(tree, global)
		fmt.Println(visitor.Cadena)
	}
	if len(visitor.Errores) > 0 {
		for _, val := range visitor.Errores {
			fmt.Println(val.GetMessage())
		}
	}

}
