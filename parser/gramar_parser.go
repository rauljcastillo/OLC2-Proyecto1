// Code generated from Gramar.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // Gramar

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type GramarParser struct {
	*antlr.BaseParser
}

var GramarParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func gramarParserInit() {
	staticData := &GramarParserStaticData
	staticData.LiteralNames = []string{
		"", "'for'", "'in'", "'{'", "'}'", "'...'", "'while'", "'switch'", "'case'",
		"':'", "'break'", "'default'", "'if'", "'else'", "'('", "')'", "'='",
		"'?'", "'+='", "'-='", "'!'", "'-'", "'*'", "'/'", "'%'", "'+'", "'<'",
		"'>'", "'<='", "'>='", "'=='", "'!='", "", "", "", "'print'", "'true'",
		"'false'", "'Int'", "'Float'", "'String'", "'Bool'", "'Character'",
		"'var'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "WS", "COMO",
		"COMM", "PRINT", "TRUE", "FALSE", "RINT", "RFLOAT", "RSTRING", "RBOOL",
		"RChar", "RVAR", "CADENA", "ID", "INT", "DOUBLE",
	}
	staticData.RuleNames = []string{
		"s", "block", "production", "pfor", "pifor", "pwhile", "pswitch", "ccase",
		"pdefault", "stmt", "pif", "pelse", "prin", "declara", "asign", "tipo",
		"expr",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 47, 181, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 1, 0, 1, 0, 1, 0, 1, 1, 5, 1, 39, 8, 1, 10, 1, 12, 1, 42,
		9, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 51, 8, 2, 1, 3, 1,
		3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 65,
		8, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 4, 6, 75, 8, 6, 11,
		6, 12, 6, 76, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 3, 7, 87,
		8, 7, 1, 8, 1, 8, 1, 8, 1, 8, 3, 8, 93, 8, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1,
		9, 1, 9, 3, 9, 101, 8, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1,
		10, 1, 10, 1, 10, 3, 10, 112, 8, 10, 1, 11, 1, 11, 1, 11, 1, 11, 3, 11,
		118, 8, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1,
		13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13,
		1, 13, 1, 13, 1, 13, 3, 13, 142, 8, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1,
		15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16,
		1, 16, 1, 16, 1, 16, 3, 16, 162, 8, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1,
		16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 5, 16, 176, 8, 16,
		10, 16, 12, 16, 179, 9, 16, 1, 16, 0, 1, 32, 17, 0, 2, 4, 6, 8, 10, 12,
		14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 0, 8, 2, 0, 16, 16, 18, 19, 1,
		0, 38, 42, 1, 0, 20, 21, 1, 0, 36, 37, 1, 0, 22, 24, 2, 0, 21, 21, 25,
		25, 1, 0, 26, 29, 1, 0, 30, 31, 189, 0, 34, 1, 0, 0, 0, 2, 40, 1, 0, 0,
		0, 4, 50, 1, 0, 0, 0, 6, 52, 1, 0, 0, 0, 8, 64, 1, 0, 0, 0, 10, 66, 1,
		0, 0, 0, 12, 70, 1, 0, 0, 0, 14, 81, 1, 0, 0, 0, 16, 88, 1, 0, 0, 0, 18,
		100, 1, 0, 0, 0, 20, 111, 1, 0, 0, 0, 22, 117, 1, 0, 0, 0, 24, 119, 1,
		0, 0, 0, 26, 141, 1, 0, 0, 0, 28, 143, 1, 0, 0, 0, 30, 147, 1, 0, 0, 0,
		32, 161, 1, 0, 0, 0, 34, 35, 3, 2, 1, 0, 35, 36, 5, 0, 0, 1, 36, 1, 1,
		0, 0, 0, 37, 39, 3, 4, 2, 0, 38, 37, 1, 0, 0, 0, 39, 42, 1, 0, 0, 0, 40,
		38, 1, 0, 0, 0, 40, 41, 1, 0, 0, 0, 41, 3, 1, 0, 0, 0, 42, 40, 1, 0, 0,
		0, 43, 51, 3, 24, 12, 0, 44, 51, 3, 26, 13, 0, 45, 51, 3, 28, 14, 0, 46,
		51, 3, 20, 10, 0, 47, 51, 3, 12, 6, 0, 48, 51, 3, 10, 5, 0, 49, 51, 3,
		6, 3, 0, 50, 43, 1, 0, 0, 0, 50, 44, 1, 0, 0, 0, 50, 45, 1, 0, 0, 0, 50,
		46, 1, 0, 0, 0, 50, 47, 1, 0, 0, 0, 50, 48, 1, 0, 0, 0, 50, 49, 1, 0, 0,
		0, 51, 5, 1, 0, 0, 0, 52, 53, 5, 1, 0, 0, 53, 54, 5, 45, 0, 0, 54, 55,
		5, 2, 0, 0, 55, 56, 3, 8, 4, 0, 56, 57, 5, 3, 0, 0, 57, 58, 3, 2, 1, 0,
		58, 59, 5, 4, 0, 0, 59, 7, 1, 0, 0, 0, 60, 61, 5, 46, 0, 0, 61, 62, 5,
		5, 0, 0, 62, 65, 5, 46, 0, 0, 63, 65, 3, 32, 16, 0, 64, 60, 1, 0, 0, 0,
		64, 63, 1, 0, 0, 0, 65, 9, 1, 0, 0, 0, 66, 67, 5, 6, 0, 0, 67, 68, 3, 32,
		16, 0, 68, 69, 3, 18, 9, 0, 69, 11, 1, 0, 0, 0, 70, 71, 5, 7, 0, 0, 71,
		72, 3, 32, 16, 0, 72, 74, 5, 3, 0, 0, 73, 75, 3, 14, 7, 0, 74, 73, 1, 0,
		0, 0, 75, 76, 1, 0, 0, 0, 76, 74, 1, 0, 0, 0, 76, 77, 1, 0, 0, 0, 77, 78,
		1, 0, 0, 0, 78, 79, 3, 16, 8, 0, 79, 80, 5, 4, 0, 0, 80, 13, 1, 0, 0, 0,
		81, 82, 5, 8, 0, 0, 82, 83, 3, 32, 16, 0, 83, 84, 5, 9, 0, 0, 84, 86, 3,
		2, 1, 0, 85, 87, 5, 10, 0, 0, 86, 85, 1, 0, 0, 0, 86, 87, 1, 0, 0, 0, 87,
		15, 1, 0, 0, 0, 88, 89, 5, 11, 0, 0, 89, 90, 5, 9, 0, 0, 90, 92, 3, 2,
		1, 0, 91, 93, 5, 10, 0, 0, 92, 91, 1, 0, 0, 0, 92, 93, 1, 0, 0, 0, 93,
		17, 1, 0, 0, 0, 94, 95, 5, 3, 0, 0, 95, 96, 3, 2, 1, 0, 96, 97, 5, 4, 0,
		0, 97, 101, 1, 0, 0, 0, 98, 99, 5, 3, 0, 0, 99, 101, 5, 4, 0, 0, 100, 94,
		1, 0, 0, 0, 100, 98, 1, 0, 0, 0, 101, 19, 1, 0, 0, 0, 102, 103, 5, 12,
		0, 0, 103, 104, 3, 32, 16, 0, 104, 105, 3, 18, 9, 0, 105, 112, 1, 0, 0,
		0, 106, 107, 5, 12, 0, 0, 107, 108, 3, 32, 16, 0, 108, 109, 3, 18, 9, 0,
		109, 110, 3, 22, 11, 0, 110, 112, 1, 0, 0, 0, 111, 102, 1, 0, 0, 0, 111,
		106, 1, 0, 0, 0, 112, 21, 1, 0, 0, 0, 113, 114, 5, 13, 0, 0, 114, 118,
		3, 20, 10, 0, 115, 116, 5, 13, 0, 0, 116, 118, 3, 18, 9, 0, 117, 113, 1,
		0, 0, 0, 117, 115, 1, 0, 0, 0, 118, 23, 1, 0, 0, 0, 119, 120, 5, 35, 0,
		0, 120, 121, 5, 14, 0, 0, 121, 122, 3, 32, 16, 0, 122, 123, 5, 15, 0, 0,
		123, 25, 1, 0, 0, 0, 124, 125, 5, 43, 0, 0, 125, 126, 5, 45, 0, 0, 126,
		127, 5, 9, 0, 0, 127, 128, 3, 30, 15, 0, 128, 129, 5, 16, 0, 0, 129, 130,
		3, 32, 16, 0, 130, 142, 1, 0, 0, 0, 131, 132, 5, 43, 0, 0, 132, 133, 5,
		45, 0, 0, 133, 134, 5, 9, 0, 0, 134, 135, 3, 30, 15, 0, 135, 136, 5, 17,
		0, 0, 136, 142, 1, 0, 0, 0, 137, 138, 5, 43, 0, 0, 138, 139, 5, 45, 0,
		0, 139, 140, 5, 16, 0, 0, 140, 142, 3, 32, 16, 0, 141, 124, 1, 0, 0, 0,
		141, 131, 1, 0, 0, 0, 141, 137, 1, 0, 0, 0, 142, 27, 1, 0, 0, 0, 143, 144,
		5, 45, 0, 0, 144, 145, 7, 0, 0, 0, 145, 146, 3, 32, 16, 0, 146, 29, 1,
		0, 0, 0, 147, 148, 7, 1, 0, 0, 148, 31, 1, 0, 0, 0, 149, 150, 6, 16, -1,
		0, 150, 151, 7, 2, 0, 0, 151, 162, 3, 32, 16, 11, 152, 153, 5, 14, 0, 0,
		153, 154, 3, 32, 16, 0, 154, 155, 5, 15, 0, 0, 155, 162, 1, 0, 0, 0, 156,
		162, 5, 46, 0, 0, 157, 162, 7, 3, 0, 0, 158, 162, 5, 45, 0, 0, 159, 162,
		5, 44, 0, 0, 160, 162, 5, 47, 0, 0, 161, 149, 1, 0, 0, 0, 161, 152, 1,
		0, 0, 0, 161, 156, 1, 0, 0, 0, 161, 157, 1, 0, 0, 0, 161, 158, 1, 0, 0,
		0, 161, 159, 1, 0, 0, 0, 161, 160, 1, 0, 0, 0, 162, 177, 1, 0, 0, 0, 163,
		164, 10, 10, 0, 0, 164, 165, 7, 4, 0, 0, 165, 176, 3, 32, 16, 11, 166,
		167, 10, 9, 0, 0, 167, 168, 7, 5, 0, 0, 168, 176, 3, 32, 16, 10, 169, 170,
		10, 8, 0, 0, 170, 171, 7, 6, 0, 0, 171, 176, 3, 32, 16, 9, 172, 173, 10,
		7, 0, 0, 173, 174, 7, 7, 0, 0, 174, 176, 3, 32, 16, 8, 175, 163, 1, 0,
		0, 0, 175, 166, 1, 0, 0, 0, 175, 169, 1, 0, 0, 0, 175, 172, 1, 0, 0, 0,
		176, 179, 1, 0, 0, 0, 177, 175, 1, 0, 0, 0, 177, 178, 1, 0, 0, 0, 178,
		33, 1, 0, 0, 0, 179, 177, 1, 0, 0, 0, 13, 40, 50, 64, 76, 86, 92, 100,
		111, 117, 141, 161, 175, 177,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// GramarParserInit initializes any static state used to implement GramarParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewGramarParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func GramarParserInit() {
	staticData := &GramarParserStaticData
	staticData.once.Do(gramarParserInit)
}

// NewGramarParser produces a new parser instance for the optional input antlr.TokenStream.
func NewGramarParser(input antlr.TokenStream) *GramarParser {
	GramarParserInit()
	this := new(GramarParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &GramarParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "Gramar.g4"

	return this
}

// GramarParser tokens.
const (
	GramarParserEOF     = antlr.TokenEOF
	GramarParserT__0    = 1
	GramarParserT__1    = 2
	GramarParserT__2    = 3
	GramarParserT__3    = 4
	GramarParserT__4    = 5
	GramarParserT__5    = 6
	GramarParserT__6    = 7
	GramarParserT__7    = 8
	GramarParserT__8    = 9
	GramarParserT__9    = 10
	GramarParserT__10   = 11
	GramarParserT__11   = 12
	GramarParserT__12   = 13
	GramarParserT__13   = 14
	GramarParserT__14   = 15
	GramarParserT__15   = 16
	GramarParserT__16   = 17
	GramarParserT__17   = 18
	GramarParserT__18   = 19
	GramarParserT__19   = 20
	GramarParserT__20   = 21
	GramarParserT__21   = 22
	GramarParserT__22   = 23
	GramarParserT__23   = 24
	GramarParserT__24   = 25
	GramarParserT__25   = 26
	GramarParserT__26   = 27
	GramarParserT__27   = 28
	GramarParserT__28   = 29
	GramarParserT__29   = 30
	GramarParserT__30   = 31
	GramarParserWS      = 32
	GramarParserCOMO    = 33
	GramarParserCOMM    = 34
	GramarParserPRINT   = 35
	GramarParserTRUE    = 36
	GramarParserFALSE   = 37
	GramarParserRINT    = 38
	GramarParserRFLOAT  = 39
	GramarParserRSTRING = 40
	GramarParserRBOOL   = 41
	GramarParserRChar   = 42
	GramarParserRVAR    = 43
	GramarParserCADENA  = 44
	GramarParserID      = 45
	GramarParserINT     = 46
	GramarParserDOUBLE  = 47
)

// GramarParser rules.
const (
	GramarParserRULE_s          = 0
	GramarParserRULE_block      = 1
	GramarParserRULE_production = 2
	GramarParserRULE_pfor       = 3
	GramarParserRULE_pifor      = 4
	GramarParserRULE_pwhile     = 5
	GramarParserRULE_pswitch    = 6
	GramarParserRULE_ccase      = 7
	GramarParserRULE_pdefault   = 8
	GramarParserRULE_stmt       = 9
	GramarParserRULE_pif        = 10
	GramarParserRULE_pelse      = 11
	GramarParserRULE_prin       = 12
	GramarParserRULE_declara    = 13
	GramarParserRULE_asign      = 14
	GramarParserRULE_tipo       = 15
	GramarParserRULE_expr       = 16
)

// ISContext is an interface to support dynamic dispatch.
type ISContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Block() IBlockContext
	EOF() antlr.TerminalNode

	// IsSContext differentiates from other interfaces.
	IsSContext()
}

type SContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySContext() *SContext {
	var p = new(SContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GramarParserRULE_s
	return p
}

func InitEmptySContext(p *SContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GramarParserRULE_s
}

func (*SContext) IsSContext() {}

func NewSContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SContext {
	var p = new(SContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramarParserRULE_s

	return p
}

func (s *SContext) GetParser() antlr.Parser { return s.parser }

func (s *SContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *SContext) EOF() antlr.TerminalNode {
	return s.GetToken(GramarParserEOF, 0)
}

func (s *SContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GramarVisitor:
		return t.VisitS(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GramarParser) S() (localctx ISContext) {
	localctx = NewSContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, GramarParserRULE_s)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(34)
		p.Block()
	}
	{
		p.SetState(35)
		p.Match(GramarParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBlockContext is an interface to support dynamic dispatch.
type IBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllProduction() []IProductionContext
	Production(i int) IProductionContext

	// IsBlockContext differentiates from other interfaces.
	IsBlockContext()
}

type BlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockContext() *BlockContext {
	var p = new(BlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GramarParserRULE_block
	return p
}

func InitEmptyBlockContext(p *BlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GramarParserRULE_block
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramarParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) AllProduction() []IProductionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IProductionContext); ok {
			len++
		}
	}

	tst := make([]IProductionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IProductionContext); ok {
			tst[i] = t.(IProductionContext)
			i++
		}
	}

	return tst
}

func (s *BlockContext) Production(i int) IProductionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IProductionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IProductionContext)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GramarVisitor:
		return t.VisitBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GramarParser) Block() (localctx IBlockContext) {
	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, GramarParserRULE_block)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(40)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&44014824853698) != 0 {
		{
			p.SetState(37)
			p.Production()
		}

		p.SetState(42)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IProductionContext is an interface to support dynamic dispatch.
type IProductionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Prin() IPrinContext
	Declara() IDeclaraContext
	Asign() IAsignContext
	Pif() IPifContext
	Pswitch() IPswitchContext
	Pwhile() IPwhileContext
	Pfor() IPforContext

	// IsProductionContext differentiates from other interfaces.
	IsProductionContext()
}

type ProductionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProductionContext() *ProductionContext {
	var p = new(ProductionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GramarParserRULE_production
	return p
}

func InitEmptyProductionContext(p *ProductionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GramarParserRULE_production
}

func (*ProductionContext) IsProductionContext() {}

func NewProductionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProductionContext {
	var p = new(ProductionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramarParserRULE_production

	return p
}

func (s *ProductionContext) GetParser() antlr.Parser { return s.parser }

func (s *ProductionContext) Prin() IPrinContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrinContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrinContext)
}

func (s *ProductionContext) Declara() IDeclaraContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclaraContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclaraContext)
}

func (s *ProductionContext) Asign() IAsignContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAsignContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAsignContext)
}

func (s *ProductionContext) Pif() IPifContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPifContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPifContext)
}

func (s *ProductionContext) Pswitch() IPswitchContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPswitchContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPswitchContext)
}

func (s *ProductionContext) Pwhile() IPwhileContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPwhileContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPwhileContext)
}

func (s *ProductionContext) Pfor() IPforContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPforContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPforContext)
}

func (s *ProductionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProductionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProductionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GramarVisitor:
		return t.VisitProduction(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GramarParser) Production() (localctx IProductionContext) {
	localctx = NewProductionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, GramarParserRULE_production)
	p.SetState(50)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case GramarParserPRINT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(43)
			p.Prin()
		}

	case GramarParserRVAR:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(44)
			p.Declara()
		}

	case GramarParserID:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(45)
			p.Asign()
		}

	case GramarParserT__11:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(46)
			p.Pif()
		}

	case GramarParserT__6:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(47)
			p.Pswitch()
		}

	case GramarParserT__5:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(48)
			p.Pwhile()
		}

	case GramarParserT__0:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(49)
			p.Pfor()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPforContext is an interface to support dynamic dispatch.
type IPforContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	Pifor() IPiforContext
	Block() IBlockContext

	// IsPforContext differentiates from other interfaces.
	IsPforContext()
}

type PforContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPforContext() *PforContext {
	var p = new(PforContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GramarParserRULE_pfor
	return p
}

func InitEmptyPforContext(p *PforContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GramarParserRULE_pfor
}

func (*PforContext) IsPforContext() {}

func NewPforContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PforContext {
	var p = new(PforContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramarParserRULE_pfor

	return p
}

func (s *PforContext) GetParser() antlr.Parser { return s.parser }

func (s *PforContext) ID() antlr.TerminalNode {
	return s.GetToken(GramarParserID, 0)
}

func (s *PforContext) Pifor() IPiforContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPiforContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPiforContext)
}

func (s *PforContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *PforContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PforContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PforContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GramarVisitor:
		return t.VisitPfor(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GramarParser) Pfor() (localctx IPforContext) {
	localctx = NewPforContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, GramarParserRULE_pfor)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(52)
		p.Match(GramarParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(53)
		p.Match(GramarParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(54)
		p.Match(GramarParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(55)
		p.Pifor()
	}
	{
		p.SetState(56)
		p.Match(GramarParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(57)
		p.Block()
	}
	{
		p.SetState(58)
		p.Match(GramarParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPiforContext is an interface to support dynamic dispatch.
type IPiforContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllINT() []antlr.TerminalNode
	INT(i int) antlr.TerminalNode
	Expr() IExprContext

	// IsPiforContext differentiates from other interfaces.
	IsPiforContext()
}

type PiforContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPiforContext() *PiforContext {
	var p = new(PiforContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GramarParserRULE_pifor
	return p
}

func InitEmptyPiforContext(p *PiforContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GramarParserRULE_pifor
}

func (*PiforContext) IsPiforContext() {}

func NewPiforContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PiforContext {
	var p = new(PiforContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramarParserRULE_pifor

	return p
}

func (s *PiforContext) GetParser() antlr.Parser { return s.parser }

func (s *PiforContext) AllINT() []antlr.TerminalNode {
	return s.GetTokens(GramarParserINT)
}

func (s *PiforContext) INT(i int) antlr.TerminalNode {
	return s.GetToken(GramarParserINT, i)
}

func (s *PiforContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *PiforContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PiforContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PiforContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GramarVisitor:
		return t.VisitPifor(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GramarParser) Pifor() (localctx IPiforContext) {
	localctx = NewPiforContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, GramarParserRULE_pifor)
	p.SetState(64)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(60)
			p.Match(GramarParserINT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(61)
			p.Match(GramarParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(62)
			p.Match(GramarParserINT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(63)
			p.expr(0)
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPwhileContext is an interface to support dynamic dispatch.
type IPwhileContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expr() IExprContext
	Stmt() IStmtContext

	// IsPwhileContext differentiates from other interfaces.
	IsPwhileContext()
}

type PwhileContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPwhileContext() *PwhileContext {
	var p = new(PwhileContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GramarParserRULE_pwhile
	return p
}

func InitEmptyPwhileContext(p *PwhileContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GramarParserRULE_pwhile
}

func (*PwhileContext) IsPwhileContext() {}

func NewPwhileContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PwhileContext {
	var p = new(PwhileContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramarParserRULE_pwhile

	return p
}

func (s *PwhileContext) GetParser() antlr.Parser { return s.parser }

func (s *PwhileContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *PwhileContext) Stmt() IStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStmtContext)
}

func (s *PwhileContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PwhileContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PwhileContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GramarVisitor:
		return t.VisitPwhile(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GramarParser) Pwhile() (localctx IPwhileContext) {
	localctx = NewPwhileContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, GramarParserRULE_pwhile)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(66)
		p.Match(GramarParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(67)
		p.expr(0)
	}
	{
		p.SetState(68)
		p.Stmt()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPswitchContext is an interface to support dynamic dispatch.
type IPswitchContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expr() IExprContext
	Pdefault() IPdefaultContext
	AllCcase() []ICcaseContext
	Ccase(i int) ICcaseContext

	// IsPswitchContext differentiates from other interfaces.
	IsPswitchContext()
}

type PswitchContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPswitchContext() *PswitchContext {
	var p = new(PswitchContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GramarParserRULE_pswitch
	return p
}

func InitEmptyPswitchContext(p *PswitchContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GramarParserRULE_pswitch
}

func (*PswitchContext) IsPswitchContext() {}

func NewPswitchContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PswitchContext {
	var p = new(PswitchContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramarParserRULE_pswitch

	return p
}

func (s *PswitchContext) GetParser() antlr.Parser { return s.parser }

func (s *PswitchContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *PswitchContext) Pdefault() IPdefaultContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPdefaultContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPdefaultContext)
}

func (s *PswitchContext) AllCcase() []ICcaseContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ICcaseContext); ok {
			len++
		}
	}

	tst := make([]ICcaseContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ICcaseContext); ok {
			tst[i] = t.(ICcaseContext)
			i++
		}
	}

	return tst
}

func (s *PswitchContext) Ccase(i int) ICcaseContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICcaseContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICcaseContext)
}

func (s *PswitchContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PswitchContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PswitchContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GramarVisitor:
		return t.VisitPswitch(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GramarParser) Pswitch() (localctx IPswitchContext) {
	localctx = NewPswitchContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, GramarParserRULE_pswitch)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(70)
		p.Match(GramarParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(71)
		p.expr(0)
	}
	{
		p.SetState(72)
		p.Match(GramarParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(74)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == GramarParserT__7 {
		{
			p.SetState(73)
			p.Ccase()
		}

		p.SetState(76)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(78)
		p.Pdefault()
	}
	{
		p.SetState(79)
		p.Match(GramarParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICcaseContext is an interface to support dynamic dispatch.
type ICcaseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expr() IExprContext
	Block() IBlockContext

	// IsCcaseContext differentiates from other interfaces.
	IsCcaseContext()
}

type CcaseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCcaseContext() *CcaseContext {
	var p = new(CcaseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GramarParserRULE_ccase
	return p
}

func InitEmptyCcaseContext(p *CcaseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GramarParserRULE_ccase
}

func (*CcaseContext) IsCcaseContext() {}

func NewCcaseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CcaseContext {
	var p = new(CcaseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramarParserRULE_ccase

	return p
}

func (s *CcaseContext) GetParser() antlr.Parser { return s.parser }

func (s *CcaseContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *CcaseContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *CcaseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CcaseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CcaseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GramarVisitor:
		return t.VisitCcase(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GramarParser) Ccase() (localctx ICcaseContext) {
	localctx = NewCcaseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, GramarParserRULE_ccase)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(81)
		p.Match(GramarParserT__7)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(82)
		p.expr(0)
	}
	{
		p.SetState(83)
		p.Match(GramarParserT__8)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(84)
		p.Block()
	}
	p.SetState(86)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == GramarParserT__9 {
		{
			p.SetState(85)
			p.Match(GramarParserT__9)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPdefaultContext is an interface to support dynamic dispatch.
type IPdefaultContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Block() IBlockContext

	// IsPdefaultContext differentiates from other interfaces.
	IsPdefaultContext()
}

type PdefaultContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPdefaultContext() *PdefaultContext {
	var p = new(PdefaultContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GramarParserRULE_pdefault
	return p
}

func InitEmptyPdefaultContext(p *PdefaultContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GramarParserRULE_pdefault
}

func (*PdefaultContext) IsPdefaultContext() {}

func NewPdefaultContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PdefaultContext {
	var p = new(PdefaultContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramarParserRULE_pdefault

	return p
}

func (s *PdefaultContext) GetParser() antlr.Parser { return s.parser }

func (s *PdefaultContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *PdefaultContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PdefaultContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PdefaultContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GramarVisitor:
		return t.VisitPdefault(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GramarParser) Pdefault() (localctx IPdefaultContext) {
	localctx = NewPdefaultContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, GramarParserRULE_pdefault)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(88)
		p.Match(GramarParserT__10)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(89)
		p.Match(GramarParserT__8)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(90)
		p.Block()
	}
	p.SetState(92)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == GramarParserT__9 {
		{
			p.SetState(91)
			p.Match(GramarParserT__9)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStmtContext is an interface to support dynamic dispatch.
type IStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Block() IBlockContext

	// IsStmtContext differentiates from other interfaces.
	IsStmtContext()
}

type StmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStmtContext() *StmtContext {
	var p = new(StmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GramarParserRULE_stmt
	return p
}

func InitEmptyStmtContext(p *StmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GramarParserRULE_stmt
}

func (*StmtContext) IsStmtContext() {}

func NewStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StmtContext {
	var p = new(StmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramarParserRULE_stmt

	return p
}

func (s *StmtContext) GetParser() antlr.Parser { return s.parser }

func (s *StmtContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *StmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GramarVisitor:
		return t.VisitStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GramarParser) Stmt() (localctx IStmtContext) {
	localctx = NewStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, GramarParserRULE_stmt)
	p.SetState(100)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(94)
			p.Match(GramarParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(95)
			p.Block()
		}
		{
			p.SetState(96)
			p.Match(GramarParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(98)
			p.Match(GramarParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(99)
			p.Match(GramarParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPifContext is an interface to support dynamic dispatch.
type IPifContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expr() IExprContext
	Stmt() IStmtContext
	Pelse() IPelseContext

	// IsPifContext differentiates from other interfaces.
	IsPifContext()
}

type PifContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPifContext() *PifContext {
	var p = new(PifContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GramarParserRULE_pif
	return p
}

func InitEmptyPifContext(p *PifContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GramarParserRULE_pif
}

func (*PifContext) IsPifContext() {}

func NewPifContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PifContext {
	var p = new(PifContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramarParserRULE_pif

	return p
}

func (s *PifContext) GetParser() antlr.Parser { return s.parser }

func (s *PifContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *PifContext) Stmt() IStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStmtContext)
}

func (s *PifContext) Pelse() IPelseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPelseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPelseContext)
}

func (s *PifContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PifContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PifContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GramarVisitor:
		return t.VisitPif(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GramarParser) Pif() (localctx IPifContext) {
	localctx = NewPifContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, GramarParserRULE_pif)
	p.SetState(111)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(102)
			p.Match(GramarParserT__11)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(103)
			p.expr(0)
		}
		{
			p.SetState(104)
			p.Stmt()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(106)
			p.Match(GramarParserT__11)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(107)
			p.expr(0)
		}
		{
			p.SetState(108)
			p.Stmt()
		}
		{
			p.SetState(109)
			p.Pelse()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPelseContext is an interface to support dynamic dispatch.
type IPelseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Pif() IPifContext
	Stmt() IStmtContext

	// IsPelseContext differentiates from other interfaces.
	IsPelseContext()
}

type PelseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPelseContext() *PelseContext {
	var p = new(PelseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GramarParserRULE_pelse
	return p
}

func InitEmptyPelseContext(p *PelseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GramarParserRULE_pelse
}

func (*PelseContext) IsPelseContext() {}

func NewPelseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PelseContext {
	var p = new(PelseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramarParserRULE_pelse

	return p
}

func (s *PelseContext) GetParser() antlr.Parser { return s.parser }

func (s *PelseContext) Pif() IPifContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPifContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPifContext)
}

func (s *PelseContext) Stmt() IStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStmtContext)
}

func (s *PelseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PelseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PelseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GramarVisitor:
		return t.VisitPelse(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GramarParser) Pelse() (localctx IPelseContext) {
	localctx = NewPelseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, GramarParserRULE_pelse)
	p.SetState(117)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(113)
			p.Match(GramarParserT__12)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(114)
			p.Pif()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(115)
			p.Match(GramarParserT__12)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(116)
			p.Stmt()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPrinContext is an interface to support dynamic dispatch.
type IPrinContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PRINT() antlr.TerminalNode
	Expr() IExprContext

	// IsPrinContext differentiates from other interfaces.
	IsPrinContext()
}

type PrinContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrinContext() *PrinContext {
	var p = new(PrinContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GramarParserRULE_prin
	return p
}

func InitEmptyPrinContext(p *PrinContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GramarParserRULE_prin
}

func (*PrinContext) IsPrinContext() {}

func NewPrinContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrinContext {
	var p = new(PrinContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramarParserRULE_prin

	return p
}

func (s *PrinContext) GetParser() antlr.Parser { return s.parser }

func (s *PrinContext) PRINT() antlr.TerminalNode {
	return s.GetToken(GramarParserPRINT, 0)
}

func (s *PrinContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *PrinContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrinContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrinContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GramarVisitor:
		return t.VisitPrin(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GramarParser) Prin() (localctx IPrinContext) {
	localctx = NewPrinContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, GramarParserRULE_prin)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(119)
		p.Match(GramarParserPRINT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(120)
		p.Match(GramarParserT__13)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(121)
		p.expr(0)
	}
	{
		p.SetState(122)
		p.Match(GramarParserT__14)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDeclaraContext is an interface to support dynamic dispatch.
type IDeclaraContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	RVAR() antlr.TerminalNode
	ID() antlr.TerminalNode
	Tipo() ITipoContext
	Expr() IExprContext

	// IsDeclaraContext differentiates from other interfaces.
	IsDeclaraContext()
}

type DeclaraContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclaraContext() *DeclaraContext {
	var p = new(DeclaraContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GramarParserRULE_declara
	return p
}

func InitEmptyDeclaraContext(p *DeclaraContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GramarParserRULE_declara
}

func (*DeclaraContext) IsDeclaraContext() {}

func NewDeclaraContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclaraContext {
	var p = new(DeclaraContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramarParserRULE_declara

	return p
}

func (s *DeclaraContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclaraContext) RVAR() antlr.TerminalNode {
	return s.GetToken(GramarParserRVAR, 0)
}

func (s *DeclaraContext) ID() antlr.TerminalNode {
	return s.GetToken(GramarParserID, 0)
}

func (s *DeclaraContext) Tipo() ITipoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITipoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITipoContext)
}

func (s *DeclaraContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *DeclaraContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclaraContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclaraContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GramarVisitor:
		return t.VisitDeclara(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GramarParser) Declara() (localctx IDeclaraContext) {
	localctx = NewDeclaraContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, GramarParserRULE_declara)
	p.SetState(141)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(124)
			p.Match(GramarParserRVAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(125)
			p.Match(GramarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(126)
			p.Match(GramarParserT__8)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(127)
			p.Tipo()
		}
		{
			p.SetState(128)
			p.Match(GramarParserT__15)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(129)
			p.expr(0)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(131)
			p.Match(GramarParserRVAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(132)
			p.Match(GramarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(133)
			p.Match(GramarParserT__8)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(134)
			p.Tipo()
		}
		{
			p.SetState(135)
			p.Match(GramarParserT__16)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(137)
			p.Match(GramarParserRVAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(138)
			p.Match(GramarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(139)
			p.Match(GramarParserT__15)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(140)
			p.expr(0)
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAsignContext is an interface to support dynamic dispatch.
type IAsignContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// Getter signatures
	ID() antlr.TerminalNode
	Expr() IExprContext

	// IsAsignContext differentiates from other interfaces.
	IsAsignContext()
}

type AsignContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	op     antlr.Token
}

func NewEmptyAsignContext() *AsignContext {
	var p = new(AsignContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GramarParserRULE_asign
	return p
}

func InitEmptyAsignContext(p *AsignContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GramarParserRULE_asign
}

func (*AsignContext) IsAsignContext() {}

func NewAsignContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AsignContext {
	var p = new(AsignContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramarParserRULE_asign

	return p
}

func (s *AsignContext) GetParser() antlr.Parser { return s.parser }

func (s *AsignContext) GetOp() antlr.Token { return s.op }

func (s *AsignContext) SetOp(v antlr.Token) { s.op = v }

func (s *AsignContext) ID() antlr.TerminalNode {
	return s.GetToken(GramarParserID, 0)
}

func (s *AsignContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *AsignContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AsignContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AsignContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GramarVisitor:
		return t.VisitAsign(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GramarParser) Asign() (localctx IAsignContext) {
	localctx = NewAsignContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, GramarParserRULE_asign)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(143)
		p.Match(GramarParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(144)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*AsignContext).op = _lt

		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&851968) != 0) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*AsignContext).op = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(145)
		p.expr(0)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITipoContext is an interface to support dynamic dispatch.
type ITipoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	RINT() antlr.TerminalNode
	RFLOAT() antlr.TerminalNode
	RSTRING() antlr.TerminalNode
	RBOOL() antlr.TerminalNode
	RChar() antlr.TerminalNode

	// IsTipoContext differentiates from other interfaces.
	IsTipoContext()
}

type TipoContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTipoContext() *TipoContext {
	var p = new(TipoContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GramarParserRULE_tipo
	return p
}

func InitEmptyTipoContext(p *TipoContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GramarParserRULE_tipo
}

func (*TipoContext) IsTipoContext() {}

func NewTipoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TipoContext {
	var p = new(TipoContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramarParserRULE_tipo

	return p
}

func (s *TipoContext) GetParser() antlr.Parser { return s.parser }

func (s *TipoContext) RINT() antlr.TerminalNode {
	return s.GetToken(GramarParserRINT, 0)
}

func (s *TipoContext) RFLOAT() antlr.TerminalNode {
	return s.GetToken(GramarParserRFLOAT, 0)
}

func (s *TipoContext) RSTRING() antlr.TerminalNode {
	return s.GetToken(GramarParserRSTRING, 0)
}

func (s *TipoContext) RBOOL() antlr.TerminalNode {
	return s.GetToken(GramarParserRBOOL, 0)
}

func (s *TipoContext) RChar() antlr.TerminalNode {
	return s.GetToken(GramarParserRChar, 0)
}

func (s *TipoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TipoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TipoContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GramarVisitor:
		return t.VisitTipo(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GramarParser) Tipo() (localctx ITipoContext) {
	localctx = NewTipoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, GramarParserRULE_tipo)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(147)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&8521215115264) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GramarParserRULE_expr
	return p
}

func InitEmptyExprContext(p *ExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GramarParserRULE_expr
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramarParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) CopyAll(ctx *ExprContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type OpContext struct {
	ExprContext
	left  IExprContext
	op    antlr.Token
	right IExprContext
}

func NewOpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OpContext {
	var p = new(OpContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *OpContext) GetOp() antlr.Token { return s.op }

func (s *OpContext) SetOp(v antlr.Token) { s.op = v }

func (s *OpContext) GetLeft() IExprContext { return s.left }

func (s *OpContext) GetRight() IExprContext { return s.right }

func (s *OpContext) SetLeft(v IExprContext) { s.left = v }

func (s *OpContext) SetRight(v IExprContext) { s.right = v }

func (s *OpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OpContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *OpContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *OpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GramarVisitor:
		return t.VisitOp(s)

	default:
		return t.VisitChildren(s)
	}
}

type LiteralContext struct {
	ExprContext
	logico antlr.Token
}

func NewLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LiteralContext {
	var p = new(LiteralContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *LiteralContext) GetLogico() antlr.Token { return s.logico }

func (s *LiteralContext) SetLogico(v antlr.Token) { s.logico = v }

func (s *LiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralContext) INT() antlr.TerminalNode {
	return s.GetToken(GramarParserINT, 0)
}

func (s *LiteralContext) TRUE() antlr.TerminalNode {
	return s.GetToken(GramarParserTRUE, 0)
}

func (s *LiteralContext) FALSE() antlr.TerminalNode {
	return s.GetToken(GramarParserFALSE, 0)
}

func (s *LiteralContext) CADENA() antlr.TerminalNode {
	return s.GetToken(GramarParserCADENA, 0)
}

func (s *LiteralContext) DOUBLE() antlr.TerminalNode {
	return s.GetToken(GramarParserDOUBLE, 0)
}

func (s *LiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GramarVisitor:
		return t.VisitLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

type AccesoContext struct {
	ExprContext
}

func NewAccesoContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AccesoContext {
	var p = new(AccesoContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *AccesoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AccesoContext) ID() antlr.TerminalNode {
	return s.GetToken(GramarParserID, 0)
}

func (s *AccesoContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GramarVisitor:
		return t.VisitAcceso(s)

	default:
		return t.VisitChildren(s)
	}
}

type ParenContext struct {
	ExprContext
}

func NewParenContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParenContext {
	var p = new(ParenContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *ParenContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParenContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ParenContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GramarVisitor:
		return t.VisitParen(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GramarParser) Expr() (localctx IExprContext) {
	return p.expr(0)
}

func (p *GramarParser) expr(_p int) (localctx IExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 32
	p.EnterRecursionRule(localctx, 32, GramarParserRULE_expr, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(161)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case GramarParserT__19, GramarParserT__20:
		localctx = NewOpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(150)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*OpContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == GramarParserT__19 || _la == GramarParserT__20) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*OpContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(151)

			var _x = p.expr(11)

			localctx.(*OpContext).right = _x
		}

	case GramarParserT__13:
		localctx = NewParenContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(152)
			p.Match(GramarParserT__13)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(153)
			p.expr(0)
		}
		{
			p.SetState(154)
			p.Match(GramarParserT__14)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case GramarParserINT:
		localctx = NewLiteralContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(156)
			p.Match(GramarParserINT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case GramarParserTRUE, GramarParserFALSE:
		localctx = NewLiteralContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(157)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*LiteralContext).logico = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == GramarParserTRUE || _la == GramarParserFALSE) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*LiteralContext).logico = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	case GramarParserID:
		localctx = NewAccesoContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(158)
			p.Match(GramarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case GramarParserCADENA:
		localctx = NewLiteralContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(159)
			p.Match(GramarParserCADENA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case GramarParserDOUBLE:
		localctx = NewLiteralContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(160)
			p.Match(GramarParserDOUBLE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(177)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 12, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(175)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext()) {
			case 1:
				localctx = NewOpContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*OpContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, GramarParserRULE_expr)
				p.SetState(163)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
					goto errorExit
				}
				{
					p.SetState(164)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*OpContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&29360128) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*OpContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(165)

					var _x = p.expr(11)

					localctx.(*OpContext).right = _x
				}

			case 2:
				localctx = NewOpContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*OpContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, GramarParserRULE_expr)
				p.SetState(166)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
					goto errorExit
				}
				{
					p.SetState(167)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*OpContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == GramarParserT__20 || _la == GramarParserT__24) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*OpContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(168)

					var _x = p.expr(10)

					localctx.(*OpContext).right = _x
				}

			case 3:
				localctx = NewOpContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*OpContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, GramarParserRULE_expr)
				p.SetState(169)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
					goto errorExit
				}
				{
					p.SetState(170)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*OpContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1006632960) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*OpContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(171)

					var _x = p.expr(9)

					localctx.(*OpContext).right = _x
				}

			case 4:
				localctx = NewOpContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*OpContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, GramarParserRULE_expr)
				p.SetState(172)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
					goto errorExit
				}
				{
					p.SetState(173)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*OpContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == GramarParserT__29 || _la == GramarParserT__30) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*OpContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(174)

					var _x = p.expr(8)

					localctx.(*OpContext).right = _x
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(179)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 12, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

func (p *GramarParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 16:
		var t *ExprContext = nil
		if localctx != nil {
			t = localctx.(*ExprContext)
		}
		return p.Expr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *GramarParser) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 10)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 9)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 8)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 7)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
