package main

import (
	"fmt"
	"main/parser"
	"main/prueba"
	"time"

	"github.com/antlr4-go/antlr/v4"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type ErrorListener struct {
	*antlr.DefaultErrorListener
	foundError bool
	errors     string
}

type Mensaje struct {
	Texto string `json:"texto"`
}

func NewError() *ErrorListener {
	return &ErrorListener{}
}

func (s *ErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	s.errors += fmt.Sprintf("Error sintáctico en la línea %d, columna %d", line, column) + "\n"
	s.foundError = true
}

func main() {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"PUT", "PATCH", "POST", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	router.POST("/texto", func(ctx *gin.Context) {
		var temp Mensaje
		var errors string = ""
		ctx.BindJSON(&temp)
		is := antlr.NewInputStream(temp.Texto)
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
				errors += val.GetMessage() + "\n"
			}
		}
		ctx.JSON(200, gin.H{
			"texto":   visitor.Cadena,
			"errores": errors,
			"sintax":  tt.errors,
		})
	})

	router.Run("localhost:4000")
}
