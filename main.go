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
	expresion := `
	for c in 1...10{
		if c==9{
			print("Soy 9 xdxd")
		}
	}
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

		result := visitor.Visit(tree, global)
		fmt.Println(result)
	}
	if len(visitor.Errores) > 0 {
		for _, val := range visitor.Errores {
			fmt.Println(val.GetMessage())
		}
	}

}
