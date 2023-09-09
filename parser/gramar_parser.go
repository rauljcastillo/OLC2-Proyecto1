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
		"", "'='", "';'", "'func'", "'('", "')'", "'->'", "'{'", "'}'", "'guard'",
		"'else'", "'for'", "'in'", "'while'", "'switch'", "'case'", "'default'",
		"'if'", "'?'", "'+='", "'-='", "'.'", "'at:'", "'!'", "'-'", "'/'",
		"'*'", "'%'", "'+'", "'<'", "'>'", "'<='", "'>='", "'=='", "'!='", "'&&'",
		"'||'", "", "", "", "'print'", "'true'", "'false'", "'Int'", "'Float'",
		"'String'", "'Bool'", "'Character'", "'var'", "'let'", "'return'", "'continue'",
		"'break'", "'append'", "'remove'", "'removeLast'", "'IsEmpty'", "'count'",
		"'['", "']'", "'_'", "'...'", "'inout'", "'&'", "','", "':'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "WS", "COMO", "COMM", "PRINT", "TRUE", "FALSE", "RINT",
		"RFLOAT", "RSTRING", "RBOOL", "RChar", "RVAR", "LET", "RETN", "CONT",
		"BRK", "APPE", "REMOV", "REMLST", "EMPT", "COUNT", "ABR", "CIER", "GUION",
		"POINTS", "INOUT", "PUNTE", "COMMA", "TPOINTS", "CADENA", "ID", "INT",
		"DOUBLE", "CHARAC",
	}
	staticData.RuleNames = []string{
		"s", "block", "production", "preturn", "pasignA", "pfuncion", "pllamada",
		"pparams", "pparamet", "pdeclarArray", "pdefinition", "pguard", "pfor",
		"pifor", "pwhile", "pswitch", "ccase", "pdefault", "stmt", "pif", "pelse",
		"prin", "cexpr", "declara", "asign", "tipo", "pespecial", "argumento",
		"tipoArg", "expr",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 70, 445, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 1, 0, 1, 0, 1, 0, 1, 1,
		5, 1, 65, 8, 1, 10, 1, 12, 1, 68, 9, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1,
		2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 86,
		8, 2, 1, 3, 1, 3, 3, 3, 90, 8, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1,
		4, 3, 4, 99, 8, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 107, 8, 5,
		1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5,
		120, 8, 5, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 126, 8, 5, 1, 6, 1, 6, 1, 6, 1,
		6, 3, 6, 132, 8, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 3, 6, 139, 8, 6, 3, 6,
		141, 8, 6, 1, 7, 1, 7, 1, 7, 5, 7, 146, 8, 7, 10, 7, 12, 7, 149, 9, 7,
		3, 7, 151, 8, 7, 1, 8, 1, 8, 1, 8, 1, 8, 3, 8, 157, 8, 8, 1, 8, 1, 8, 1,
		8, 1, 8, 1, 8, 3, 8, 164, 8, 8, 1, 8, 1, 8, 1, 8, 1, 8, 3, 8, 170, 8, 8,
		1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 3, 9,
		183, 8, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 5, 10, 190, 8, 10, 10, 10,
		12, 10, 193, 9, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1,
		11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12,
		1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 3, 13, 218, 8, 13, 1, 14, 1, 14, 1,
		14, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 15, 4, 15, 230, 8, 15,
		11, 15, 12, 15, 231, 1, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 16, 1,
		16, 3, 16, 242, 8, 16, 1, 17, 1, 17, 1, 17, 1, 17, 3, 17, 248, 8, 17, 1,
		18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 3, 18, 256, 8, 18, 1, 19, 1, 19,
		1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 3, 19, 267, 8, 19, 1,
		20, 1, 20, 1, 20, 1, 20, 3, 20, 273, 8, 20, 1, 21, 1, 21, 1, 21, 1, 21,
		1, 21, 3, 21, 280, 8, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 3, 21, 287,
		8, 21, 3, 21, 289, 8, 21, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 3, 22, 296,
		8, 22, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 3, 23, 305, 8,
		23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23,
		3, 23, 317, 8, 23, 1, 24, 1, 24, 1, 24, 1, 24, 1, 25, 1, 25, 1, 26, 1,
		26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 3, 26, 332, 8, 26, 1, 26, 1, 26,
		1, 26, 1, 26, 1, 26, 1, 26, 3, 26, 340, 8, 26, 1, 26, 1, 26, 1, 26, 1,
		26, 1, 26, 1, 26, 1, 26, 1, 26, 3, 26, 350, 8, 26, 3, 26, 352, 8, 26, 1,
		27, 1, 27, 1, 27, 1, 27, 1, 27, 3, 27, 359, 8, 27, 1, 28, 1, 28, 1, 28,
		1, 28, 1, 28, 1, 28, 3, 28, 367, 8, 28, 1, 29, 1, 29, 1, 29, 1, 29, 1,
		29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29,
		1, 29, 1, 29, 3, 29, 386, 8, 29, 1, 29, 1, 29, 1, 29, 1, 29, 3, 29, 392,
		8, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 3, 29, 399, 8, 29, 1, 29, 1,
		29, 1, 29, 1, 29, 1, 29, 3, 29, 406, 8, 29, 1, 29, 1, 29, 1, 29, 1, 29,
		1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 3, 29, 418, 8, 29, 3, 29, 420,
		8, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1,
		29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 5, 29, 440,
		8, 29, 10, 29, 12, 29, 443, 9, 29, 1, 29, 0, 1, 58, 30, 0, 2, 4, 6, 8,
		10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44,
		46, 48, 50, 52, 54, 56, 58, 0, 11, 2, 0, 60, 60, 67, 67, 1, 0, 48, 49,
		1, 0, 50, 52, 2, 0, 1, 1, 19, 20, 1, 0, 43, 47, 1, 0, 23, 24, 1, 0, 41,
		42, 1, 0, 25, 27, 2, 0, 24, 24, 28, 28, 1, 0, 29, 32, 1, 0, 33, 34, 491,
		0, 60, 1, 0, 0, 0, 2, 66, 1, 0, 0, 0, 4, 85, 1, 0, 0, 0, 6, 87, 1, 0, 0,
		0, 8, 91, 1, 0, 0, 0, 10, 125, 1, 0, 0, 0, 12, 140, 1, 0, 0, 0, 14, 150,
		1, 0, 0, 0, 16, 169, 1, 0, 0, 0, 18, 182, 1, 0, 0, 0, 20, 184, 1, 0, 0,
		0, 22, 196, 1, 0, 0, 0, 24, 204, 1, 0, 0, 0, 26, 217, 1, 0, 0, 0, 28, 219,
		1, 0, 0, 0, 30, 225, 1, 0, 0, 0, 32, 236, 1, 0, 0, 0, 34, 243, 1, 0, 0,
		0, 36, 255, 1, 0, 0, 0, 38, 266, 1, 0, 0, 0, 40, 272, 1, 0, 0, 0, 42, 288,
		1, 0, 0, 0, 44, 295, 1, 0, 0, 0, 46, 316, 1, 0, 0, 0, 48, 318, 1, 0, 0,
		0, 50, 322, 1, 0, 0, 0, 52, 351, 1, 0, 0, 0, 54, 358, 1, 0, 0, 0, 56, 366,
		1, 0, 0, 0, 58, 419, 1, 0, 0, 0, 60, 61, 3, 2, 1, 0, 61, 62, 5, 0, 0, 1,
		62, 1, 1, 0, 0, 0, 63, 65, 3, 4, 2, 0, 64, 63, 1, 0, 0, 0, 65, 68, 1, 0,
		0, 0, 66, 64, 1, 0, 0, 0, 66, 67, 1, 0, 0, 0, 67, 3, 1, 0, 0, 0, 68, 66,
		1, 0, 0, 0, 69, 86, 3, 42, 21, 0, 70, 86, 3, 8, 4, 0, 71, 86, 3, 46, 23,
		0, 72, 86, 3, 48, 24, 0, 73, 86, 3, 38, 19, 0, 74, 86, 3, 30, 15, 0, 75,
		86, 3, 28, 14, 0, 76, 86, 3, 24, 12, 0, 77, 86, 3, 22, 11, 0, 78, 86, 3,
		18, 9, 0, 79, 86, 3, 52, 26, 0, 80, 86, 3, 10, 5, 0, 81, 86, 3, 12, 6,
		0, 82, 86, 3, 6, 3, 0, 83, 86, 5, 51, 0, 0, 84, 86, 5, 52, 0, 0, 85, 69,
		1, 0, 0, 0, 85, 70, 1, 0, 0, 0, 85, 71, 1, 0, 0, 0, 85, 72, 1, 0, 0, 0,
		85, 73, 1, 0, 0, 0, 85, 74, 1, 0, 0, 0, 85, 75, 1, 0, 0, 0, 85, 76, 1,
		0, 0, 0, 85, 77, 1, 0, 0, 0, 85, 78, 1, 0, 0, 0, 85, 79, 1, 0, 0, 0, 85,
		80, 1, 0, 0, 0, 85, 81, 1, 0, 0, 0, 85, 82, 1, 0, 0, 0, 85, 83, 1, 0, 0,
		0, 85, 84, 1, 0, 0, 0, 86, 5, 1, 0, 0, 0, 87, 89, 5, 50, 0, 0, 88, 90,
		3, 58, 29, 0, 89, 88, 1, 0, 0, 0, 89, 90, 1, 0, 0, 0, 90, 7, 1, 0, 0, 0,
		91, 92, 5, 67, 0, 0, 92, 93, 5, 58, 0, 0, 93, 94, 3, 58, 29, 0, 94, 95,
		5, 59, 0, 0, 95, 96, 5, 1, 0, 0, 96, 98, 3, 58, 29, 0, 97, 99, 5, 2, 0,
		0, 98, 97, 1, 0, 0, 0, 98, 99, 1, 0, 0, 0, 99, 9, 1, 0, 0, 0, 100, 101,
		5, 3, 0, 0, 101, 102, 5, 67, 0, 0, 102, 103, 5, 4, 0, 0, 103, 106, 5, 5,
		0, 0, 104, 105, 5, 6, 0, 0, 105, 107, 3, 50, 25, 0, 106, 104, 1, 0, 0,
		0, 106, 107, 1, 0, 0, 0, 107, 108, 1, 0, 0, 0, 108, 109, 5, 7, 0, 0, 109,
		110, 3, 2, 1, 0, 110, 111, 5, 8, 0, 0, 111, 126, 1, 0, 0, 0, 112, 113,
		5, 3, 0, 0, 113, 114, 5, 67, 0, 0, 114, 115, 5, 4, 0, 0, 115, 116, 3, 14,
		7, 0, 116, 119, 5, 5, 0, 0, 117, 118, 5, 6, 0, 0, 118, 120, 3, 50, 25,
		0, 119, 117, 1, 0, 0, 0, 119, 120, 1, 0, 0, 0, 120, 121, 1, 0, 0, 0, 121,
		122, 5, 7, 0, 0, 122, 123, 3, 2, 1, 0, 123, 124, 5, 8, 0, 0, 124, 126,
		1, 0, 0, 0, 125, 100, 1, 0, 0, 0, 125, 112, 1, 0, 0, 0, 126, 11, 1, 0,
		0, 0, 127, 128, 5, 67, 0, 0, 128, 129, 5, 4, 0, 0, 129, 131, 5, 5, 0, 0,
		130, 132, 5, 2, 0, 0, 131, 130, 1, 0, 0, 0, 131, 132, 1, 0, 0, 0, 132,
		141, 1, 0, 0, 0, 133, 134, 5, 67, 0, 0, 134, 135, 5, 4, 0, 0, 135, 136,
		3, 54, 27, 0, 136, 138, 5, 5, 0, 0, 137, 139, 5, 2, 0, 0, 138, 137, 1,
		0, 0, 0, 138, 139, 1, 0, 0, 0, 139, 141, 1, 0, 0, 0, 140, 127, 1, 0, 0,
		0, 140, 133, 1, 0, 0, 0, 141, 13, 1, 0, 0, 0, 142, 147, 3, 16, 8, 0, 143,
		144, 5, 64, 0, 0, 144, 146, 3, 16, 8, 0, 145, 143, 1, 0, 0, 0, 146, 149,
		1, 0, 0, 0, 147, 145, 1, 0, 0, 0, 147, 148, 1, 0, 0, 0, 148, 151, 1, 0,
		0, 0, 149, 147, 1, 0, 0, 0, 150, 142, 1, 0, 0, 0, 150, 151, 1, 0, 0, 0,
		151, 15, 1, 0, 0, 0, 152, 153, 7, 0, 0, 0, 153, 154, 5, 67, 0, 0, 154,
		156, 5, 65, 0, 0, 155, 157, 5, 62, 0, 0, 156, 155, 1, 0, 0, 0, 156, 157,
		1, 0, 0, 0, 157, 158, 1, 0, 0, 0, 158, 170, 3, 50, 25, 0, 159, 160, 7,
		0, 0, 0, 160, 161, 5, 67, 0, 0, 161, 163, 5, 65, 0, 0, 162, 164, 5, 62,
		0, 0, 163, 162, 1, 0, 0, 0, 163, 164, 1, 0, 0, 0, 164, 165, 1, 0, 0, 0,
		165, 166, 5, 58, 0, 0, 166, 167, 3, 50, 25, 0, 167, 168, 5, 59, 0, 0, 168,
		170, 1, 0, 0, 0, 169, 152, 1, 0, 0, 0, 169, 159, 1, 0, 0, 0, 170, 17, 1,
		0, 0, 0, 171, 172, 7, 1, 0, 0, 172, 173, 5, 67, 0, 0, 173, 174, 5, 65,
		0, 0, 174, 175, 5, 58, 0, 0, 175, 176, 3, 50, 25, 0, 176, 177, 5, 59, 0,
		0, 177, 178, 3, 20, 10, 0, 178, 183, 1, 0, 0, 0, 179, 180, 7, 1, 0, 0,
		180, 181, 5, 67, 0, 0, 181, 183, 3, 20, 10, 0, 182, 171, 1, 0, 0, 0, 182,
		179, 1, 0, 0, 0, 183, 19, 1, 0, 0, 0, 184, 185, 5, 1, 0, 0, 185, 186, 5,
		58, 0, 0, 186, 191, 3, 58, 29, 0, 187, 188, 5, 64, 0, 0, 188, 190, 3, 58,
		29, 0, 189, 187, 1, 0, 0, 0, 190, 193, 1, 0, 0, 0, 191, 189, 1, 0, 0, 0,
		191, 192, 1, 0, 0, 0, 192, 194, 1, 0, 0, 0, 193, 191, 1, 0, 0, 0, 194,
		195, 5, 59, 0, 0, 195, 21, 1, 0, 0, 0, 196, 197, 5, 9, 0, 0, 197, 198,
		3, 58, 29, 0, 198, 199, 5, 10, 0, 0, 199, 200, 5, 7, 0, 0, 200, 201, 3,
		2, 1, 0, 201, 202, 7, 2, 0, 0, 202, 203, 5, 8, 0, 0, 203, 23, 1, 0, 0,
		0, 204, 205, 5, 11, 0, 0, 205, 206, 5, 67, 0, 0, 206, 207, 5, 12, 0, 0,
		207, 208, 3, 26, 13, 0, 208, 209, 5, 7, 0, 0, 209, 210, 3, 2, 1, 0, 210,
		211, 5, 8, 0, 0, 211, 25, 1, 0, 0, 0, 212, 213, 3, 58, 29, 0, 213, 214,
		5, 61, 0, 0, 214, 215, 3, 58, 29, 0, 215, 218, 1, 0, 0, 0, 216, 218, 3,
		58, 29, 0, 217, 212, 1, 0, 0, 0, 217, 216, 1, 0, 0, 0, 218, 27, 1, 0, 0,
		0, 219, 220, 5, 13, 0, 0, 220, 221, 3, 58, 29, 0, 221, 222, 5, 7, 0, 0,
		222, 223, 3, 2, 1, 0, 223, 224, 5, 8, 0, 0, 224, 29, 1, 0, 0, 0, 225, 226,
		5, 14, 0, 0, 226, 227, 3, 58, 29, 0, 227, 229, 5, 7, 0, 0, 228, 230, 3,
		32, 16, 0, 229, 228, 1, 0, 0, 0, 230, 231, 1, 0, 0, 0, 231, 229, 1, 0,
		0, 0, 231, 232, 1, 0, 0, 0, 232, 233, 1, 0, 0, 0, 233, 234, 3, 34, 17,
		0, 234, 235, 5, 8, 0, 0, 235, 31, 1, 0, 0, 0, 236, 237, 5, 15, 0, 0, 237,
		238, 3, 58, 29, 0, 238, 239, 5, 65, 0, 0, 239, 241, 3, 2, 1, 0, 240, 242,
		5, 52, 0, 0, 241, 240, 1, 0, 0, 0, 241, 242, 1, 0, 0, 0, 242, 33, 1, 0,
		0, 0, 243, 244, 5, 16, 0, 0, 244, 245, 5, 65, 0, 0, 245, 247, 3, 2, 1,
		0, 246, 248, 5, 52, 0, 0, 247, 246, 1, 0, 0, 0, 247, 248, 1, 0, 0, 0, 248,
		35, 1, 0, 0, 0, 249, 250, 5, 7, 0, 0, 250, 251, 3, 2, 1, 0, 251, 252, 5,
		8, 0, 0, 252, 256, 1, 0, 0, 0, 253, 254, 5, 7, 0, 0, 254, 256, 5, 8, 0,
		0, 255, 249, 1, 0, 0, 0, 255, 253, 1, 0, 0, 0, 256, 37, 1, 0, 0, 0, 257,
		258, 5, 17, 0, 0, 258, 259, 3, 58, 29, 0, 259, 260, 3, 36, 18, 0, 260,
		267, 1, 0, 0, 0, 261, 262, 5, 17, 0, 0, 262, 263, 3, 58, 29, 0, 263, 264,
		3, 36, 18, 0, 264, 265, 3, 40, 20, 0, 265, 267, 1, 0, 0, 0, 266, 257, 1,
		0, 0, 0, 266, 261, 1, 0, 0, 0, 267, 39, 1, 0, 0, 0, 268, 269, 5, 10, 0,
		0, 269, 273, 3, 38, 19, 0, 270, 271, 5, 10, 0, 0, 271, 273, 3, 36, 18,
		0, 272, 268, 1, 0, 0, 0, 272, 270, 1, 0, 0, 0, 273, 41, 1, 0, 0, 0, 274,
		275, 5, 40, 0, 0, 275, 276, 5, 4, 0, 0, 276, 277, 3, 58, 29, 0, 277, 279,
		5, 5, 0, 0, 278, 280, 5, 2, 0, 0, 279, 278, 1, 0, 0, 0, 279, 280, 1, 0,
		0, 0, 280, 289, 1, 0, 0, 0, 281, 282, 5, 40, 0, 0, 282, 283, 5, 4, 0, 0,
		283, 284, 3, 44, 22, 0, 284, 286, 5, 5, 0, 0, 285, 287, 5, 2, 0, 0, 286,
		285, 1, 0, 0, 0, 286, 287, 1, 0, 0, 0, 287, 289, 1, 0, 0, 0, 288, 274,
		1, 0, 0, 0, 288, 281, 1, 0, 0, 0, 289, 43, 1, 0, 0, 0, 290, 291, 3, 58,
		29, 0, 291, 292, 5, 64, 0, 0, 292, 293, 3, 44, 22, 0, 293, 296, 1, 0, 0,
		0, 294, 296, 3, 58, 29, 0, 295, 290, 1, 0, 0, 0, 295, 294, 1, 0, 0, 0,
		296, 45, 1, 0, 0, 0, 297, 298, 7, 1, 0, 0, 298, 299, 5, 67, 0, 0, 299,
		300, 5, 65, 0, 0, 300, 301, 3, 50, 25, 0, 301, 302, 5, 1, 0, 0, 302, 304,
		3, 58, 29, 0, 303, 305, 5, 2, 0, 0, 304, 303, 1, 0, 0, 0, 304, 305, 1,
		0, 0, 0, 305, 317, 1, 0, 0, 0, 306, 307, 7, 1, 0, 0, 307, 308, 5, 67, 0,
		0, 308, 309, 5, 65, 0, 0, 309, 310, 3, 50, 25, 0, 310, 311, 5, 18, 0, 0,
		311, 317, 1, 0, 0, 0, 312, 313, 7, 1, 0, 0, 313, 314, 5, 67, 0, 0, 314,
		315, 5, 1, 0, 0, 315, 317, 3, 58, 29, 0, 316, 297, 1, 0, 0, 0, 316, 306,
		1, 0, 0, 0, 316, 312, 1, 0, 0, 0, 317, 47, 1, 0, 0, 0, 318, 319, 5, 67,
		0, 0, 319, 320, 7, 3, 0, 0, 320, 321, 3, 58, 29, 0, 321, 49, 1, 0, 0, 0,
		322, 323, 7, 4, 0, 0, 323, 51, 1, 0, 0, 0, 324, 325, 5, 67, 0, 0, 325,
		326, 5, 21, 0, 0, 326, 327, 5, 53, 0, 0, 327, 328, 5, 4, 0, 0, 328, 329,
		3, 58, 29, 0, 329, 331, 5, 5, 0, 0, 330, 332, 5, 2, 0, 0, 331, 330, 1,
		0, 0, 0, 331, 332, 1, 0, 0, 0, 332, 352, 1, 0, 0, 0, 333, 334, 5, 67, 0,
		0, 334, 335, 5, 21, 0, 0, 335, 336, 5, 55, 0, 0, 336, 337, 5, 4, 0, 0,
		337, 339, 5, 5, 0, 0, 338, 340, 5, 2, 0, 0, 339, 338, 1, 0, 0, 0, 339,
		340, 1, 0, 0, 0, 340, 352, 1, 0, 0, 0, 341, 342, 5, 67, 0, 0, 342, 343,
		5, 21, 0, 0, 343, 344, 5, 54, 0, 0, 344, 345, 5, 4, 0, 0, 345, 346, 5,
		22, 0, 0, 346, 347, 3, 58, 29, 0, 347, 349, 5, 5, 0, 0, 348, 350, 5, 2,
		0, 0, 349, 348, 1, 0, 0, 0, 349, 350, 1, 0, 0, 0, 350, 352, 1, 0, 0, 0,
		351, 324, 1, 0, 0, 0, 351, 333, 1, 0, 0, 0, 351, 341, 1, 0, 0, 0, 352,
		53, 1, 0, 0, 0, 353, 354, 3, 56, 28, 0, 354, 355, 5, 64, 0, 0, 355, 356,
		3, 54, 27, 0, 356, 359, 1, 0, 0, 0, 357, 359, 3, 56, 28, 0, 358, 353, 1,
		0, 0, 0, 358, 357, 1, 0, 0, 0, 359, 55, 1, 0, 0, 0, 360, 361, 5, 63, 0,
		0, 361, 367, 3, 58, 29, 0, 362, 363, 5, 67, 0, 0, 363, 364, 5, 65, 0, 0,
		364, 367, 3, 58, 29, 0, 365, 367, 3, 58, 29, 0, 366, 360, 1, 0, 0, 0, 366,
		362, 1, 0, 0, 0, 366, 365, 1, 0, 0, 0, 367, 57, 1, 0, 0, 0, 368, 369, 6,
		29, -1, 0, 369, 370, 7, 5, 0, 0, 370, 420, 3, 58, 29, 20, 371, 420, 3,
		12, 6, 0, 372, 373, 5, 4, 0, 0, 373, 374, 3, 58, 29, 0, 374, 375, 5, 5,
		0, 0, 375, 420, 1, 0, 0, 0, 376, 420, 5, 68, 0, 0, 377, 420, 7, 6, 0, 0,
		378, 420, 5, 66, 0, 0, 379, 420, 5, 69, 0, 0, 380, 420, 5, 67, 0, 0, 381,
		382, 5, 67, 0, 0, 382, 383, 5, 21, 0, 0, 383, 385, 5, 56, 0, 0, 384, 386,
		5, 2, 0, 0, 385, 384, 1, 0, 0, 0, 385, 386, 1, 0, 0, 0, 386, 420, 1, 0,
		0, 0, 387, 388, 5, 67, 0, 0, 388, 389, 5, 21, 0, 0, 389, 391, 5, 57, 0,
		0, 390, 392, 5, 2, 0, 0, 391, 390, 1, 0, 0, 0, 391, 392, 1, 0, 0, 0, 392,
		420, 1, 0, 0, 0, 393, 394, 5, 67, 0, 0, 394, 395, 5, 58, 0, 0, 395, 396,
		3, 58, 29, 0, 396, 398, 5, 59, 0, 0, 397, 399, 5, 2, 0, 0, 398, 397, 1,
		0, 0, 0, 398, 399, 1, 0, 0, 0, 399, 420, 1, 0, 0, 0, 400, 401, 5, 43, 0,
		0, 401, 402, 5, 4, 0, 0, 402, 403, 3, 58, 29, 0, 403, 405, 5, 5, 0, 0,
		404, 406, 5, 2, 0, 0, 405, 404, 1, 0, 0, 0, 405, 406, 1, 0, 0, 0, 406,
		420, 1, 0, 0, 0, 407, 408, 5, 44, 0, 0, 408, 409, 5, 4, 0, 0, 409, 410,
		3, 58, 29, 0, 410, 411, 5, 5, 0, 0, 411, 420, 1, 0, 0, 0, 412, 413, 5,
		45, 0, 0, 413, 414, 5, 4, 0, 0, 414, 415, 3, 58, 29, 0, 415, 417, 5, 5,
		0, 0, 416, 418, 5, 2, 0, 0, 417, 416, 1, 0, 0, 0, 417, 418, 1, 0, 0, 0,
		418, 420, 1, 0, 0, 0, 419, 368, 1, 0, 0, 0, 419, 371, 1, 0, 0, 0, 419,
		372, 1, 0, 0, 0, 419, 376, 1, 0, 0, 0, 419, 377, 1, 0, 0, 0, 419, 378,
		1, 0, 0, 0, 419, 379, 1, 0, 0, 0, 419, 380, 1, 0, 0, 0, 419, 381, 1, 0,
		0, 0, 419, 387, 1, 0, 0, 0, 419, 393, 1, 0, 0, 0, 419, 400, 1, 0, 0, 0,
		419, 407, 1, 0, 0, 0, 419, 412, 1, 0, 0, 0, 420, 441, 1, 0, 0, 0, 421,
		422, 10, 19, 0, 0, 422, 423, 7, 7, 0, 0, 423, 440, 3, 58, 29, 20, 424,
		425, 10, 18, 0, 0, 425, 426, 7, 8, 0, 0, 426, 440, 3, 58, 29, 19, 427,
		428, 10, 17, 0, 0, 428, 429, 7, 9, 0, 0, 429, 440, 3, 58, 29, 18, 430,
		431, 10, 16, 0, 0, 431, 432, 7, 10, 0, 0, 432, 440, 3, 58, 29, 17, 433,
		434, 10, 15, 0, 0, 434, 435, 5, 35, 0, 0, 435, 440, 3, 58, 29, 16, 436,
		437, 10, 14, 0, 0, 437, 438, 5, 36, 0, 0, 438, 440, 3, 58, 29, 15, 439,
		421, 1, 0, 0, 0, 439, 424, 1, 0, 0, 0, 439, 427, 1, 0, 0, 0, 439, 430,
		1, 0, 0, 0, 439, 433, 1, 0, 0, 0, 439, 436, 1, 0, 0, 0, 440, 443, 1, 0,
		0, 0, 441, 439, 1, 0, 0, 0, 441, 442, 1, 0, 0, 0, 442, 59, 1, 0, 0, 0,
		443, 441, 1, 0, 0, 0, 44, 66, 85, 89, 98, 106, 119, 125, 131, 138, 140,
		147, 150, 156, 163, 169, 182, 191, 217, 231, 241, 247, 255, 266, 272, 279,
		286, 288, 295, 304, 316, 331, 339, 349, 351, 358, 366, 385, 391, 398, 405,
		417, 419, 439, 441,
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
	GramarParserT__31   = 32
	GramarParserT__32   = 33
	GramarParserT__33   = 34
	GramarParserT__34   = 35
	GramarParserT__35   = 36
	GramarParserWS      = 37
	GramarParserCOMO    = 38
	GramarParserCOMM    = 39
	GramarParserPRINT   = 40
	GramarParserTRUE    = 41
	GramarParserFALSE   = 42
	GramarParserRINT    = 43
	GramarParserRFLOAT  = 44
	GramarParserRSTRING = 45
	GramarParserRBOOL   = 46
	GramarParserRChar   = 47
	GramarParserRVAR    = 48
	GramarParserLET     = 49
	GramarParserRETN    = 50
	GramarParserCONT    = 51
	GramarParserBRK     = 52
	GramarParserAPPE    = 53
	GramarParserREMOV   = 54
	GramarParserREMLST  = 55
	GramarParserEMPT    = 56
	GramarParserCOUNT   = 57
	GramarParserABR     = 58
	GramarParserCIER    = 59
	GramarParserGUION   = 60
	GramarParserPOINTS  = 61
	GramarParserINOUT   = 62
	GramarParserPUNTE   = 63
	GramarParserCOMMA   = 64
	GramarParserTPOINTS = 65
	GramarParserCADENA  = 66
	GramarParserID      = 67
	GramarParserINT     = 68
	GramarParserDOUBLE  = 69
	GramarParserCHARAC  = 70
)

// GramarParser rules.
const (
	GramarParserRULE_s            = 0
	GramarParserRULE_block        = 1
	GramarParserRULE_production   = 2
	GramarParserRULE_preturn      = 3
	GramarParserRULE_pasignA      = 4
	GramarParserRULE_pfuncion     = 5
	GramarParserRULE_pllamada     = 6
	GramarParserRULE_pparams      = 7
	GramarParserRULE_pparamet     = 8
	GramarParserRULE_pdeclarArray = 9
	GramarParserRULE_pdefinition  = 10
	GramarParserRULE_pguard       = 11
	GramarParserRULE_pfor         = 12
	GramarParserRULE_pifor        = 13
	GramarParserRULE_pwhile       = 14
	GramarParserRULE_pswitch      = 15
	GramarParserRULE_ccase        = 16
	GramarParserRULE_pdefault     = 17
	GramarParserRULE_stmt         = 18
	GramarParserRULE_pif          = 19
	GramarParserRULE_pelse        = 20
	GramarParserRULE_prin         = 21
	GramarParserRULE_cexpr        = 22
	GramarParserRULE_declara      = 23
	GramarParserRULE_asign        = 24
	GramarParserRULE_tipo         = 25
	GramarParserRULE_pespecial    = 26
	GramarParserRULE_argumento    = 27
	GramarParserRULE_tipoArg      = 28
	GramarParserRULE_expr         = 29
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
		p.SetState(60)
		p.Block()
	}
	{
		p.SetState(61)
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
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(66)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 0, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(63)
				p.Production()
			}

		}
		p.SetState(68)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 0, p.GetParserRuleContext())
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
	PasignA() IPasignAContext
	Declara() IDeclaraContext
	Asign() IAsignContext
	Pif() IPifContext
	Pswitch() IPswitchContext
	Pwhile() IPwhileContext
	Pfor() IPforContext
	Pguard() IPguardContext
	PdeclarArray() IPdeclarArrayContext
	Pespecial() IPespecialContext
	Pfuncion() IPfuncionContext
	Pllamada() IPllamadaContext
	Preturn() IPreturnContext
	CONT() antlr.TerminalNode
	BRK() antlr.TerminalNode

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

func (s *ProductionContext) PasignA() IPasignAContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPasignAContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPasignAContext)
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

func (s *ProductionContext) Pguard() IPguardContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPguardContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPguardContext)
}

func (s *ProductionContext) PdeclarArray() IPdeclarArrayContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPdeclarArrayContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPdeclarArrayContext)
}

func (s *ProductionContext) Pespecial() IPespecialContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPespecialContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPespecialContext)
}

func (s *ProductionContext) Pfuncion() IPfuncionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPfuncionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPfuncionContext)
}

func (s *ProductionContext) Pllamada() IPllamadaContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPllamadaContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPllamadaContext)
}

func (s *ProductionContext) Preturn() IPreturnContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPreturnContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPreturnContext)
}

func (s *ProductionContext) CONT() antlr.TerminalNode {
	return s.GetToken(GramarParserCONT, 0)
}

func (s *ProductionContext) BRK() antlr.TerminalNode {
	return s.GetToken(GramarParserBRK, 0)
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
	p.SetState(85)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(69)
			p.Prin()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(70)
			p.PasignA()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(71)
			p.Declara()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(72)
			p.Asign()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(73)
			p.Pif()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(74)
			p.Pswitch()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(75)
			p.Pwhile()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(76)
			p.Pfor()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(77)
			p.Pguard()
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(78)
			p.PdeclarArray()
		}

	case 11:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(79)
			p.Pespecial()
		}

	case 12:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(80)
			p.Pfuncion()
		}

	case 13:
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(81)
			p.Pllamada()
		}

	case 14:
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(82)
			p.Preturn()
		}

	case 15:
		p.EnterOuterAlt(localctx, 15)
		{
			p.SetState(83)
			p.Match(GramarParserCONT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 16:
		p.EnterOuterAlt(localctx, 16)
		{
			p.SetState(84)
			p.Match(GramarParserBRK)
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

// IPreturnContext is an interface to support dynamic dispatch.
type IPreturnContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	RETN() antlr.TerminalNode
	Expr() IExprContext

	// IsPreturnContext differentiates from other interfaces.
	IsPreturnContext()
}

type PreturnContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPreturnContext() *PreturnContext {
	var p = new(PreturnContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GramarParserRULE_preturn
	return p
}

func InitEmptyPreturnContext(p *PreturnContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GramarParserRULE_preturn
}

func (*PreturnContext) IsPreturnContext() {}

func NewPreturnContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PreturnContext {
	var p = new(PreturnContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramarParserRULE_preturn

	return p
}

func (s *PreturnContext) GetParser() antlr.Parser { return s.parser }

func (s *PreturnContext) RETN() antlr.TerminalNode {
	return s.GetToken(GramarParserRETN, 0)
}

func (s *PreturnContext) Expr() IExprContext {
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

func (s *PreturnContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PreturnContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PreturnContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GramarVisitor:
		return t.VisitPreturn(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GramarParser) Preturn() (localctx IPreturnContext) {
	localctx = NewPreturnContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, GramarParserRULE_preturn)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(87)
		p.Match(GramarParserRETN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(89)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(88)
			p.expr(0)
		}

	} else if p.HasError() { // JIM
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

// IPasignAContext is an interface to support dynamic dispatch.
type IPasignAContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	ABR() antlr.TerminalNode
	AllExpr() []IExprContext
	Expr(i int) IExprContext
	CIER() antlr.TerminalNode

	// IsPasignAContext differentiates from other interfaces.
	IsPasignAContext()
}

type PasignAContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPasignAContext() *PasignAContext {
	var p = new(PasignAContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GramarParserRULE_pasignA
	return p
}

func InitEmptyPasignAContext(p *PasignAContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GramarParserRULE_pasignA
}

func (*PasignAContext) IsPasignAContext() {}

func NewPasignAContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PasignAContext {
	var p = new(PasignAContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramarParserRULE_pasignA

	return p
}

func (s *PasignAContext) GetParser() antlr.Parser { return s.parser }

func (s *PasignAContext) ID() antlr.TerminalNode {
	return s.GetToken(GramarParserID, 0)
}

func (s *PasignAContext) ABR() antlr.TerminalNode {
	return s.GetToken(GramarParserABR, 0)
}

func (s *PasignAContext) AllExpr() []IExprContext {
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

func (s *PasignAContext) Expr(i int) IExprContext {
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

func (s *PasignAContext) CIER() antlr.TerminalNode {
	return s.GetToken(GramarParserCIER, 0)
}

func (s *PasignAContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PasignAContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PasignAContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GramarVisitor:
		return t.VisitPasignA(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GramarParser) PasignA() (localctx IPasignAContext) {
	localctx = NewPasignAContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, GramarParserRULE_pasignA)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(91)
		p.Match(GramarParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(92)
		p.Match(GramarParserABR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(93)
		p.expr(0)
	}
	{
		p.SetState(94)
		p.Match(GramarParserCIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(95)
		p.Match(GramarParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(96)
		p.expr(0)
	}
	p.SetState(98)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == GramarParserT__1 {
		{
			p.SetState(97)
			p.Match(GramarParserT__1)
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

// IPfuncionContext is an interface to support dynamic dispatch.
type IPfuncionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	Block() IBlockContext
	Tipo() ITipoContext
	Pparams() IPparamsContext

	// IsPfuncionContext differentiates from other interfaces.
	IsPfuncionContext()
}

type PfuncionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPfuncionContext() *PfuncionContext {
	var p = new(PfuncionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GramarParserRULE_pfuncion
	return p
}

func InitEmptyPfuncionContext(p *PfuncionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GramarParserRULE_pfuncion
}

func (*PfuncionContext) IsPfuncionContext() {}

func NewPfuncionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PfuncionContext {
	var p = new(PfuncionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramarParserRULE_pfuncion

	return p
}

func (s *PfuncionContext) GetParser() antlr.Parser { return s.parser }

func (s *PfuncionContext) ID() antlr.TerminalNode {
	return s.GetToken(GramarParserID, 0)
}

func (s *PfuncionContext) Block() IBlockContext {
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

func (s *PfuncionContext) Tipo() ITipoContext {
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

func (s *PfuncionContext) Pparams() IPparamsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPparamsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPparamsContext)
}

func (s *PfuncionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PfuncionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PfuncionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GramarVisitor:
		return t.VisitPfuncion(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GramarParser) Pfuncion() (localctx IPfuncionContext) {
	localctx = NewPfuncionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, GramarParserRULE_pfuncion)
	var _la int

	p.SetState(125)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(100)
			p.Match(GramarParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(101)
			p.Match(GramarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(102)
			p.Match(GramarParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(103)
			p.Match(GramarParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(106)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == GramarParserT__5 {
			{
				p.SetState(104)
				p.Match(GramarParserT__5)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(105)
				p.Tipo()
			}

		}
		{
			p.SetState(108)
			p.Match(GramarParserT__6)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(109)
			p.Block()
		}
		{
			p.SetState(110)
			p.Match(GramarParserT__7)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(112)
			p.Match(GramarParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(113)
			p.Match(GramarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(114)
			p.Match(GramarParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(115)
			p.Pparams()
		}
		{
			p.SetState(116)
			p.Match(GramarParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(119)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == GramarParserT__5 {
			{
				p.SetState(117)
				p.Match(GramarParserT__5)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(118)
				p.Tipo()
			}

		}
		{
			p.SetState(121)
			p.Match(GramarParserT__6)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(122)
			p.Block()
		}
		{
			p.SetState(123)
			p.Match(GramarParserT__7)
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

// IPllamadaContext is an interface to support dynamic dispatch.
type IPllamadaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	Argumento() IArgumentoContext

	// IsPllamadaContext differentiates from other interfaces.
	IsPllamadaContext()
}

type PllamadaContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPllamadaContext() *PllamadaContext {
	var p = new(PllamadaContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GramarParserRULE_pllamada
	return p
}

func InitEmptyPllamadaContext(p *PllamadaContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GramarParserRULE_pllamada
}

func (*PllamadaContext) IsPllamadaContext() {}

func NewPllamadaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PllamadaContext {
	var p = new(PllamadaContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramarParserRULE_pllamada

	return p
}

func (s *PllamadaContext) GetParser() antlr.Parser { return s.parser }

func (s *PllamadaContext) ID() antlr.TerminalNode {
	return s.GetToken(GramarParserID, 0)
}

func (s *PllamadaContext) Argumento() IArgumentoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgumentoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgumentoContext)
}

func (s *PllamadaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PllamadaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PllamadaContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GramarVisitor:
		return t.VisitPllamada(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GramarParser) Pllamada() (localctx IPllamadaContext) {
	localctx = NewPllamadaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, GramarParserRULE_pllamada)
	p.SetState(140)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(127)
			p.Match(GramarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(128)
			p.Match(GramarParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(129)
			p.Match(GramarParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(131)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(130)
				p.Match(GramarParserT__1)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(133)
			p.Match(GramarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(134)
			p.Match(GramarParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(135)
			p.Argumento()
		}
		{
			p.SetState(136)
			p.Match(GramarParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(138)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(137)
				p.Match(GramarParserT__1)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		} else if p.HasError() { // JIM
			goto errorExit
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

// IPparamsContext is an interface to support dynamic dispatch.
type IPparamsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllPparamet() []IPparametContext
	Pparamet(i int) IPparametContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsPparamsContext differentiates from other interfaces.
	IsPparamsContext()
}

type PparamsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPparamsContext() *PparamsContext {
	var p = new(PparamsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GramarParserRULE_pparams
	return p
}

func InitEmptyPparamsContext(p *PparamsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GramarParserRULE_pparams
}

func (*PparamsContext) IsPparamsContext() {}

func NewPparamsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PparamsContext {
	var p = new(PparamsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramarParserRULE_pparams

	return p
}

func (s *PparamsContext) GetParser() antlr.Parser { return s.parser }

func (s *PparamsContext) AllPparamet() []IPparametContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IPparametContext); ok {
			len++
		}
	}

	tst := make([]IPparametContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IPparametContext); ok {
			tst[i] = t.(IPparametContext)
			i++
		}
	}

	return tst
}

func (s *PparamsContext) Pparamet(i int) IPparametContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPparametContext); ok {
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

	return t.(IPparametContext)
}

func (s *PparamsContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(GramarParserCOMMA)
}

func (s *PparamsContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(GramarParserCOMMA, i)
}

func (s *PparamsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PparamsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PparamsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GramarVisitor:
		return t.VisitPparams(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GramarParser) Pparams() (localctx IPparamsContext) {
	localctx = NewPparamsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, GramarParserRULE_pparams)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(150)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == GramarParserGUION || _la == GramarParserID {
		{
			p.SetState(142)
			p.Pparamet()
		}
		p.SetState(147)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == GramarParserCOMMA {
			{
				p.SetState(143)
				p.Match(GramarParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(144)
				p.Pparamet()
			}

			p.SetState(149)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
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

// IPparametContext is an interface to support dynamic dispatch.
type IPparametContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetPid returns the pid token.
	GetPid() antlr.Token

	// SetPid sets the pid token.
	SetPid(antlr.Token)

	// Getter signatures
	AllID() []antlr.TerminalNode
	ID(i int) antlr.TerminalNode
	TPOINTS() antlr.TerminalNode
	Tipo() ITipoContext
	GUION() antlr.TerminalNode
	INOUT() antlr.TerminalNode
	ABR() antlr.TerminalNode
	CIER() antlr.TerminalNode

	// IsPparametContext differentiates from other interfaces.
	IsPparametContext()
}

type PparametContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	pid    antlr.Token
}

func NewEmptyPparametContext() *PparametContext {
	var p = new(PparametContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GramarParserRULE_pparamet
	return p
}

func InitEmptyPparametContext(p *PparametContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GramarParserRULE_pparamet
}

func (*PparametContext) IsPparametContext() {}

func NewPparametContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PparametContext {
	var p = new(PparametContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramarParserRULE_pparamet

	return p
}

func (s *PparametContext) GetParser() antlr.Parser { return s.parser }

func (s *PparametContext) GetPid() antlr.Token { return s.pid }

func (s *PparametContext) SetPid(v antlr.Token) { s.pid = v }

func (s *PparametContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(GramarParserID)
}

func (s *PparametContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(GramarParserID, i)
}

func (s *PparametContext) TPOINTS() antlr.TerminalNode {
	return s.GetToken(GramarParserTPOINTS, 0)
}

func (s *PparametContext) Tipo() ITipoContext {
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

func (s *PparametContext) GUION() antlr.TerminalNode {
	return s.GetToken(GramarParserGUION, 0)
}

func (s *PparametContext) INOUT() antlr.TerminalNode {
	return s.GetToken(GramarParserINOUT, 0)
}

func (s *PparametContext) ABR() antlr.TerminalNode {
	return s.GetToken(GramarParserABR, 0)
}

func (s *PparametContext) CIER() antlr.TerminalNode {
	return s.GetToken(GramarParserCIER, 0)
}

func (s *PparametContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PparametContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PparametContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GramarVisitor:
		return t.VisitPparamet(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GramarParser) Pparamet() (localctx IPparametContext) {
	localctx = NewPparametContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, GramarParserRULE_pparamet)
	var _la int

	p.SetState(169)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 14, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(152)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*PparametContext).pid = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == GramarParserGUION || _la == GramarParserID) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*PparametContext).pid = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(153)
			p.Match(GramarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(154)
			p.Match(GramarParserTPOINTS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(156)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == GramarParserINOUT {
			{
				p.SetState(155)
				p.Match(GramarParserINOUT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(158)
			p.Tipo()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(159)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*PparametContext).pid = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == GramarParserGUION || _la == GramarParserID) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*PparametContext).pid = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(160)
			p.Match(GramarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(161)
			p.Match(GramarParserTPOINTS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(163)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == GramarParserINOUT {
			{
				p.SetState(162)
				p.Match(GramarParserINOUT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(165)
			p.Match(GramarParserABR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(166)
			p.Tipo()
		}
		{
			p.SetState(167)
			p.Match(GramarParserCIER)
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

// IPdeclarArrayContext is an interface to support dynamic dispatch.
type IPdeclarArrayContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	TPOINTS() antlr.TerminalNode
	ABR() antlr.TerminalNode
	Tipo() ITipoContext
	CIER() antlr.TerminalNode
	Pdefinition() IPdefinitionContext
	RVAR() antlr.TerminalNode
	LET() antlr.TerminalNode

	// IsPdeclarArrayContext differentiates from other interfaces.
	IsPdeclarArrayContext()
}

type PdeclarArrayContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPdeclarArrayContext() *PdeclarArrayContext {
	var p = new(PdeclarArrayContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GramarParserRULE_pdeclarArray
	return p
}

func InitEmptyPdeclarArrayContext(p *PdeclarArrayContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GramarParserRULE_pdeclarArray
}

func (*PdeclarArrayContext) IsPdeclarArrayContext() {}

func NewPdeclarArrayContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PdeclarArrayContext {
	var p = new(PdeclarArrayContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramarParserRULE_pdeclarArray

	return p
}

func (s *PdeclarArrayContext) GetParser() antlr.Parser { return s.parser }

func (s *PdeclarArrayContext) ID() antlr.TerminalNode {
	return s.GetToken(GramarParserID, 0)
}

func (s *PdeclarArrayContext) TPOINTS() antlr.TerminalNode {
	return s.GetToken(GramarParserTPOINTS, 0)
}

func (s *PdeclarArrayContext) ABR() antlr.TerminalNode {
	return s.GetToken(GramarParserABR, 0)
}

func (s *PdeclarArrayContext) Tipo() ITipoContext {
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

func (s *PdeclarArrayContext) CIER() antlr.TerminalNode {
	return s.GetToken(GramarParserCIER, 0)
}

func (s *PdeclarArrayContext) Pdefinition() IPdefinitionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPdefinitionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPdefinitionContext)
}

func (s *PdeclarArrayContext) RVAR() antlr.TerminalNode {
	return s.GetToken(GramarParserRVAR, 0)
}

func (s *PdeclarArrayContext) LET() antlr.TerminalNode {
	return s.GetToken(GramarParserLET, 0)
}

func (s *PdeclarArrayContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PdeclarArrayContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PdeclarArrayContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GramarVisitor:
		return t.VisitPdeclarArray(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GramarParser) PdeclarArray() (localctx IPdeclarArrayContext) {
	localctx = NewPdeclarArrayContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, GramarParserRULE_pdeclarArray)
	var _la int

	p.SetState(182)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 15, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(171)
			_la = p.GetTokenStream().LA(1)

			if !(_la == GramarParserRVAR || _la == GramarParserLET) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(172)
			p.Match(GramarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(173)
			p.Match(GramarParserTPOINTS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(174)
			p.Match(GramarParserABR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(175)
			p.Tipo()
		}
		{
			p.SetState(176)
			p.Match(GramarParserCIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(177)
			p.Pdefinition()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(179)
			_la = p.GetTokenStream().LA(1)

			if !(_la == GramarParserRVAR || _la == GramarParserLET) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(180)
			p.Match(GramarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(181)
			p.Pdefinition()
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

// IPdefinitionContext is an interface to support dynamic dispatch.
type IPdefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ABR() antlr.TerminalNode
	CIER() antlr.TerminalNode
	AllExpr() []IExprContext
	Expr(i int) IExprContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsPdefinitionContext differentiates from other interfaces.
	IsPdefinitionContext()
}

type PdefinitionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPdefinitionContext() *PdefinitionContext {
	var p = new(PdefinitionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GramarParserRULE_pdefinition
	return p
}

func InitEmptyPdefinitionContext(p *PdefinitionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GramarParserRULE_pdefinition
}

func (*PdefinitionContext) IsPdefinitionContext() {}

func NewPdefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PdefinitionContext {
	var p = new(PdefinitionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramarParserRULE_pdefinition

	return p
}

func (s *PdefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *PdefinitionContext) ABR() antlr.TerminalNode {
	return s.GetToken(GramarParserABR, 0)
}

func (s *PdefinitionContext) CIER() antlr.TerminalNode {
	return s.GetToken(GramarParserCIER, 0)
}

func (s *PdefinitionContext) AllExpr() []IExprContext {
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

func (s *PdefinitionContext) Expr(i int) IExprContext {
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

func (s *PdefinitionContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(GramarParserCOMMA)
}

func (s *PdefinitionContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(GramarParserCOMMA, i)
}

func (s *PdefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PdefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PdefinitionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GramarVisitor:
		return t.VisitPdefinition(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GramarParser) Pdefinition() (localctx IPdefinitionContext) {
	localctx = NewPdefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, GramarParserRULE_pdefinition)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(184)
		p.Match(GramarParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(185)
		p.Match(GramarParserABR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

	{
		p.SetState(186)
		p.expr(0)
	}
	p.SetState(191)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == GramarParserCOMMA {
		{
			p.SetState(187)
			p.Match(GramarParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(188)
			p.expr(0)
		}

		p.SetState(193)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

	{
		p.SetState(194)
		p.Match(GramarParserCIER)
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

// IPguardContext is an interface to support dynamic dispatch.
type IPguardContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOpera returns the opera token.
	GetOpera() antlr.Token

	// SetOpera sets the opera token.
	SetOpera(antlr.Token)

	// Getter signatures
	Expr() IExprContext
	Block() IBlockContext
	RETN() antlr.TerminalNode
	CONT() antlr.TerminalNode
	BRK() antlr.TerminalNode

	// IsPguardContext differentiates from other interfaces.
	IsPguardContext()
}

type PguardContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	opera  antlr.Token
}

func NewEmptyPguardContext() *PguardContext {
	var p = new(PguardContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GramarParserRULE_pguard
	return p
}

func InitEmptyPguardContext(p *PguardContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GramarParserRULE_pguard
}

func (*PguardContext) IsPguardContext() {}

func NewPguardContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PguardContext {
	var p = new(PguardContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramarParserRULE_pguard

	return p
}

func (s *PguardContext) GetParser() antlr.Parser { return s.parser }

func (s *PguardContext) GetOpera() antlr.Token { return s.opera }

func (s *PguardContext) SetOpera(v antlr.Token) { s.opera = v }

func (s *PguardContext) Expr() IExprContext {
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

func (s *PguardContext) Block() IBlockContext {
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

func (s *PguardContext) RETN() antlr.TerminalNode {
	return s.GetToken(GramarParserRETN, 0)
}

func (s *PguardContext) CONT() antlr.TerminalNode {
	return s.GetToken(GramarParserCONT, 0)
}

func (s *PguardContext) BRK() antlr.TerminalNode {
	return s.GetToken(GramarParserBRK, 0)
}

func (s *PguardContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PguardContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PguardContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GramarVisitor:
		return t.VisitPguard(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GramarParser) Pguard() (localctx IPguardContext) {
	localctx = NewPguardContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, GramarParserRULE_pguard)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(196)
		p.Match(GramarParserT__8)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(197)
		p.expr(0)
	}
	{
		p.SetState(198)
		p.Match(GramarParserT__9)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(199)
		p.Match(GramarParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(200)
		p.Block()
	}
	{
		p.SetState(201)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*PguardContext).opera = _lt

		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&7881299347898368) != 0) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*PguardContext).opera = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(202)
		p.Match(GramarParserT__7)
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
	p.EnterRule(localctx, 24, GramarParserRULE_pfor)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(204)
		p.Match(GramarParserT__10)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(205)
		p.Match(GramarParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(206)
		p.Match(GramarParserT__11)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(207)
		p.Pifor()
	}
	{
		p.SetState(208)
		p.Match(GramarParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(209)
		p.Block()
	}
	{
		p.SetState(210)
		p.Match(GramarParserT__7)
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
	AllExpr() []IExprContext
	Expr(i int) IExprContext
	POINTS() antlr.TerminalNode

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

func (s *PiforContext) AllExpr() []IExprContext {
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

func (s *PiforContext) Expr(i int) IExprContext {
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

func (s *PiforContext) POINTS() antlr.TerminalNode {
	return s.GetToken(GramarParserPOINTS, 0)
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
	p.EnterRule(localctx, 26, GramarParserRULE_pifor)
	p.SetState(217)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 17, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(212)
			p.expr(0)
		}
		{
			p.SetState(213)
			p.Match(GramarParserPOINTS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(214)
			p.expr(0)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(216)
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
	Block() IBlockContext

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

func (s *PwhileContext) Block() IBlockContext {
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
	p.EnterRule(localctx, 28, GramarParserRULE_pwhile)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(219)
		p.Match(GramarParserT__12)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(220)
		p.expr(0)
	}
	{
		p.SetState(221)
		p.Match(GramarParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(222)
		p.Block()
	}
	{
		p.SetState(223)
		p.Match(GramarParserT__7)
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
	p.EnterRule(localctx, 30, GramarParserRULE_pswitch)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(225)
		p.Match(GramarParserT__13)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(226)
		p.expr(0)
	}
	{
		p.SetState(227)
		p.Match(GramarParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(229)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == GramarParserT__14 {
		{
			p.SetState(228)
			p.Ccase()
		}

		p.SetState(231)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(233)
		p.Pdefault()
	}
	{
		p.SetState(234)
		p.Match(GramarParserT__7)
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
	TPOINTS() antlr.TerminalNode
	Block() IBlockContext
	BRK() antlr.TerminalNode

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

func (s *CcaseContext) TPOINTS() antlr.TerminalNode {
	return s.GetToken(GramarParserTPOINTS, 0)
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

func (s *CcaseContext) BRK() antlr.TerminalNode {
	return s.GetToken(GramarParserBRK, 0)
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
	p.EnterRule(localctx, 32, GramarParserRULE_ccase)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(236)
		p.Match(GramarParserT__14)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(237)
		p.expr(0)
	}
	{
		p.SetState(238)
		p.Match(GramarParserTPOINTS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(239)
		p.Block()
	}
	p.SetState(241)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == GramarParserBRK {
		{
			p.SetState(240)
			p.Match(GramarParserBRK)
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
	TPOINTS() antlr.TerminalNode
	Block() IBlockContext
	BRK() antlr.TerminalNode

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

func (s *PdefaultContext) TPOINTS() antlr.TerminalNode {
	return s.GetToken(GramarParserTPOINTS, 0)
}

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

func (s *PdefaultContext) BRK() antlr.TerminalNode {
	return s.GetToken(GramarParserBRK, 0)
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
	p.EnterRule(localctx, 34, GramarParserRULE_pdefault)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(243)
		p.Match(GramarParserT__15)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(244)
		p.Match(GramarParserTPOINTS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(245)
		p.Block()
	}
	p.SetState(247)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == GramarParserBRK {
		{
			p.SetState(246)
			p.Match(GramarParserBRK)
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
	p.EnterRule(localctx, 36, GramarParserRULE_stmt)
	p.SetState(255)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 21, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(249)
			p.Match(GramarParserT__6)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(250)
			p.Block()
		}
		{
			p.SetState(251)
			p.Match(GramarParserT__7)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(253)
			p.Match(GramarParserT__6)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(254)
			p.Match(GramarParserT__7)
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
	p.EnterRule(localctx, 38, GramarParserRULE_pif)
	p.SetState(266)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 22, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(257)
			p.Match(GramarParserT__16)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(258)
			p.expr(0)
		}
		{
			p.SetState(259)
			p.Stmt()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(261)
			p.Match(GramarParserT__16)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(262)
			p.expr(0)
		}
		{
			p.SetState(263)
			p.Stmt()
		}
		{
			p.SetState(264)
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
	p.EnterRule(localctx, 40, GramarParserRULE_pelse)
	p.SetState(272)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 23, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(268)
			p.Match(GramarParserT__9)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(269)
			p.Pif()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(270)
			p.Match(GramarParserT__9)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(271)
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
	Cexpr() ICexprContext

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

func (s *PrinContext) Cexpr() ICexprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICexprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICexprContext)
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
	p.EnterRule(localctx, 42, GramarParserRULE_prin)
	var _la int

	p.SetState(288)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 26, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(274)
			p.Match(GramarParserPRINT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(275)
			p.Match(GramarParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(276)
			p.expr(0)
		}
		{
			p.SetState(277)
			p.Match(GramarParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(279)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == GramarParserT__1 {
			{
				p.SetState(278)
				p.Match(GramarParserT__1)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(281)
			p.Match(GramarParserPRINT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(282)
			p.Match(GramarParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(283)
			p.Cexpr()
		}
		{
			p.SetState(284)
			p.Match(GramarParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(286)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == GramarParserT__1 {
			{
				p.SetState(285)
				p.Match(GramarParserT__1)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
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

// ICexprContext is an interface to support dynamic dispatch.
type ICexprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expr() IExprContext
	COMMA() antlr.TerminalNode
	Cexpr() ICexprContext

	// IsCexprContext differentiates from other interfaces.
	IsCexprContext()
}

type CexprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCexprContext() *CexprContext {
	var p = new(CexprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GramarParserRULE_cexpr
	return p
}

func InitEmptyCexprContext(p *CexprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GramarParserRULE_cexpr
}

func (*CexprContext) IsCexprContext() {}

func NewCexprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CexprContext {
	var p = new(CexprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramarParserRULE_cexpr

	return p
}

func (s *CexprContext) GetParser() antlr.Parser { return s.parser }

func (s *CexprContext) Expr() IExprContext {
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

func (s *CexprContext) COMMA() antlr.TerminalNode {
	return s.GetToken(GramarParserCOMMA, 0)
}

func (s *CexprContext) Cexpr() ICexprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICexprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICexprContext)
}

func (s *CexprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CexprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CexprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GramarVisitor:
		return t.VisitCexpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GramarParser) Cexpr() (localctx ICexprContext) {
	localctx = NewCexprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, GramarParserRULE_cexpr)
	p.SetState(295)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 27, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(290)
			p.expr(0)
		}
		{
			p.SetState(291)
			p.Match(GramarParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(292)
			p.Cexpr()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(294)
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

// IDeclaraContext is an interface to support dynamic dispatch.
type IDeclaraContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	TPOINTS() antlr.TerminalNode
	Tipo() ITipoContext
	Expr() IExprContext
	RVAR() antlr.TerminalNode
	LET() antlr.TerminalNode

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

func (s *DeclaraContext) ID() antlr.TerminalNode {
	return s.GetToken(GramarParserID, 0)
}

func (s *DeclaraContext) TPOINTS() antlr.TerminalNode {
	return s.GetToken(GramarParserTPOINTS, 0)
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

func (s *DeclaraContext) RVAR() antlr.TerminalNode {
	return s.GetToken(GramarParserRVAR, 0)
}

func (s *DeclaraContext) LET() antlr.TerminalNode {
	return s.GetToken(GramarParserLET, 0)
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
	p.EnterRule(localctx, 46, GramarParserRULE_declara)
	var _la int

	p.SetState(316)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 29, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(297)
			_la = p.GetTokenStream().LA(1)

			if !(_la == GramarParserRVAR || _la == GramarParserLET) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(298)
			p.Match(GramarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(299)
			p.Match(GramarParserTPOINTS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(300)
			p.Tipo()
		}
		{
			p.SetState(301)
			p.Match(GramarParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(302)
			p.expr(0)
		}
		p.SetState(304)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == GramarParserT__1 {
			{
				p.SetState(303)
				p.Match(GramarParserT__1)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(306)
			_la = p.GetTokenStream().LA(1)

			if !(_la == GramarParserRVAR || _la == GramarParserLET) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(307)
			p.Match(GramarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(308)
			p.Match(GramarParserTPOINTS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(309)
			p.Tipo()
		}
		{
			p.SetState(310)
			p.Match(GramarParserT__17)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(312)
			_la = p.GetTokenStream().LA(1)

			if !(_la == GramarParserRVAR || _la == GramarParserLET) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(313)
			p.Match(GramarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(314)
			p.Match(GramarParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(315)
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
	p.EnterRule(localctx, 48, GramarParserRULE_asign)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(318)
		p.Match(GramarParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(319)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*AsignContext).op = _lt

		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1572866) != 0) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*AsignContext).op = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(320)
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
	p.EnterRule(localctx, 50, GramarParserRULE_tipo)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(322)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&272678883688448) != 0) {
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

// IPespecialContext is an interface to support dynamic dispatch.
type IPespecialContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	APPE() antlr.TerminalNode
	Expr() IExprContext
	REMLST() antlr.TerminalNode
	REMOV() antlr.TerminalNode

	// IsPespecialContext differentiates from other interfaces.
	IsPespecialContext()
}

type PespecialContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPespecialContext() *PespecialContext {
	var p = new(PespecialContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GramarParserRULE_pespecial
	return p
}

func InitEmptyPespecialContext(p *PespecialContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GramarParserRULE_pespecial
}

func (*PespecialContext) IsPespecialContext() {}

func NewPespecialContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PespecialContext {
	var p = new(PespecialContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramarParserRULE_pespecial

	return p
}

func (s *PespecialContext) GetParser() antlr.Parser { return s.parser }

func (s *PespecialContext) ID() antlr.TerminalNode {
	return s.GetToken(GramarParserID, 0)
}

func (s *PespecialContext) APPE() antlr.TerminalNode {
	return s.GetToken(GramarParserAPPE, 0)
}

func (s *PespecialContext) Expr() IExprContext {
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

func (s *PespecialContext) REMLST() antlr.TerminalNode {
	return s.GetToken(GramarParserREMLST, 0)
}

func (s *PespecialContext) REMOV() antlr.TerminalNode {
	return s.GetToken(GramarParserREMOV, 0)
}

func (s *PespecialContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PespecialContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PespecialContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GramarVisitor:
		return t.VisitPespecial(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GramarParser) Pespecial() (localctx IPespecialContext) {
	localctx = NewPespecialContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, GramarParserRULE_pespecial)
	var _la int

	p.SetState(351)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 33, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(324)
			p.Match(GramarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(325)
			p.Match(GramarParserT__20)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(326)
			p.Match(GramarParserAPPE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(327)
			p.Match(GramarParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(328)
			p.expr(0)
		}
		{
			p.SetState(329)
			p.Match(GramarParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(331)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == GramarParserT__1 {
			{
				p.SetState(330)
				p.Match(GramarParserT__1)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(333)
			p.Match(GramarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(334)
			p.Match(GramarParserT__20)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(335)
			p.Match(GramarParserREMLST)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(336)
			p.Match(GramarParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(337)
			p.Match(GramarParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(339)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == GramarParserT__1 {
			{
				p.SetState(338)
				p.Match(GramarParserT__1)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(341)
			p.Match(GramarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(342)
			p.Match(GramarParserT__20)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(343)
			p.Match(GramarParserREMOV)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(344)
			p.Match(GramarParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(345)
			p.Match(GramarParserT__21)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(346)
			p.expr(0)
		}
		{
			p.SetState(347)
			p.Match(GramarParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(349)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == GramarParserT__1 {
			{
				p.SetState(348)
				p.Match(GramarParserT__1)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
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

// IArgumentoContext is an interface to support dynamic dispatch.
type IArgumentoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TipoArg() ITipoArgContext
	COMMA() antlr.TerminalNode
	Argumento() IArgumentoContext

	// IsArgumentoContext differentiates from other interfaces.
	IsArgumentoContext()
}

type ArgumentoContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgumentoContext() *ArgumentoContext {
	var p = new(ArgumentoContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GramarParserRULE_argumento
	return p
}

func InitEmptyArgumentoContext(p *ArgumentoContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GramarParserRULE_argumento
}

func (*ArgumentoContext) IsArgumentoContext() {}

func NewArgumentoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgumentoContext {
	var p = new(ArgumentoContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramarParserRULE_argumento

	return p
}

func (s *ArgumentoContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgumentoContext) TipoArg() ITipoArgContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITipoArgContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITipoArgContext)
}

func (s *ArgumentoContext) COMMA() antlr.TerminalNode {
	return s.GetToken(GramarParserCOMMA, 0)
}

func (s *ArgumentoContext) Argumento() IArgumentoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgumentoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgumentoContext)
}

func (s *ArgumentoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgumentoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgumentoContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GramarVisitor:
		return t.VisitArgumento(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GramarParser) Argumento() (localctx IArgumentoContext) {
	localctx = NewArgumentoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, GramarParserRULE_argumento)
	p.SetState(358)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 34, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(353)
			p.TipoArg()
		}
		{
			p.SetState(354)
			p.Match(GramarParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(355)
			p.Argumento()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(357)
			p.TipoArg()
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

// ITipoArgContext is an interface to support dynamic dispatch.
type ITipoArgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetExp1 returns the exp1 rule contexts.
	GetExp1() IExprContext

	// GetExp2 returns the exp2 rule contexts.
	GetExp2() IExprContext

	// GetExp3 returns the exp3 rule contexts.
	GetExp3() IExprContext

	// SetExp1 sets the exp1 rule contexts.
	SetExp1(IExprContext)

	// SetExp2 sets the exp2 rule contexts.
	SetExp2(IExprContext)

	// SetExp3 sets the exp3 rule contexts.
	SetExp3(IExprContext)

	// Getter signatures
	PUNTE() antlr.TerminalNode
	Expr() IExprContext
	ID() antlr.TerminalNode
	TPOINTS() antlr.TerminalNode

	// IsTipoArgContext differentiates from other interfaces.
	IsTipoArgContext()
}

type TipoArgContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	exp1   IExprContext
	exp2   IExprContext
	exp3   IExprContext
}

func NewEmptyTipoArgContext() *TipoArgContext {
	var p = new(TipoArgContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GramarParserRULE_tipoArg
	return p
}

func InitEmptyTipoArgContext(p *TipoArgContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GramarParserRULE_tipoArg
}

func (*TipoArgContext) IsTipoArgContext() {}

func NewTipoArgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TipoArgContext {
	var p = new(TipoArgContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramarParserRULE_tipoArg

	return p
}

func (s *TipoArgContext) GetParser() antlr.Parser { return s.parser }

func (s *TipoArgContext) GetExp1() IExprContext { return s.exp1 }

func (s *TipoArgContext) GetExp2() IExprContext { return s.exp2 }

func (s *TipoArgContext) GetExp3() IExprContext { return s.exp3 }

func (s *TipoArgContext) SetExp1(v IExprContext) { s.exp1 = v }

func (s *TipoArgContext) SetExp2(v IExprContext) { s.exp2 = v }

func (s *TipoArgContext) SetExp3(v IExprContext) { s.exp3 = v }

func (s *TipoArgContext) PUNTE() antlr.TerminalNode {
	return s.GetToken(GramarParserPUNTE, 0)
}

func (s *TipoArgContext) Expr() IExprContext {
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

func (s *TipoArgContext) ID() antlr.TerminalNode {
	return s.GetToken(GramarParserID, 0)
}

func (s *TipoArgContext) TPOINTS() antlr.TerminalNode {
	return s.GetToken(GramarParserTPOINTS, 0)
}

func (s *TipoArgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TipoArgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TipoArgContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GramarVisitor:
		return t.VisitTipoArg(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GramarParser) TipoArg() (localctx ITipoArgContext) {
	localctx = NewTipoArgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, GramarParserRULE_tipoArg)
	p.SetState(366)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 35, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(360)
			p.Match(GramarParserPUNTE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(361)

			var _x = p.expr(0)

			localctx.(*TipoArgContext).exp1 = _x
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(362)
			p.Match(GramarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(363)
			p.Match(GramarParserTPOINTS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(364)

			var _x = p.expr(0)

			localctx.(*TipoArgContext).exp2 = _x
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(365)

			var _x = p.expr(0)

			localctx.(*TipoArgContext).exp3 = _x
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

type LlamadaContext struct {
	ExprContext
}

func NewLlamadaContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LlamadaContext {
	var p = new(LlamadaContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *LlamadaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LlamadaContext) Pllamada() IPllamadaContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPllamadaContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPllamadaContext)
}

func (s *LlamadaContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GramarVisitor:
		return t.VisitLlamada(s)

	default:
		return t.VisitChildren(s)
	}
}

type AccesoAContext struct {
	ExprContext
}

func NewAccesoAContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AccesoAContext {
	var p = new(AccesoAContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *AccesoAContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AccesoAContext) ID() antlr.TerminalNode {
	return s.GetToken(GramarParserID, 0)
}

func (s *AccesoAContext) ABR() antlr.TerminalNode {
	return s.GetToken(GramarParserABR, 0)
}

func (s *AccesoAContext) Expr() IExprContext {
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

func (s *AccesoAContext) CIER() antlr.TerminalNode {
	return s.GetToken(GramarParserCIER, 0)
}

func (s *AccesoAContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GramarVisitor:
		return t.VisitAccesoA(s)

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

type EspecialContext struct {
	ExprContext
}

func NewEspecialContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EspecialContext {
	var p = new(EspecialContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *EspecialContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EspecialContext) ID() antlr.TerminalNode {
	return s.GetToken(GramarParserID, 0)
}

func (s *EspecialContext) EMPT() antlr.TerminalNode {
	return s.GetToken(GramarParserEMPT, 0)
}

func (s *EspecialContext) COUNT() antlr.TerminalNode {
	return s.GetToken(GramarParserCOUNT, 0)
}

func (s *EspecialContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GramarVisitor:
		return t.VisitEspecial(s)

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

type EmbebidaContext struct {
	ExprContext
}

func NewEmbebidaContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EmbebidaContext {
	var p = new(EmbebidaContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *EmbebidaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EmbebidaContext) RINT() antlr.TerminalNode {
	return s.GetToken(GramarParserRINT, 0)
}

func (s *EmbebidaContext) Expr() IExprContext {
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

func (s *EmbebidaContext) RFLOAT() antlr.TerminalNode {
	return s.GetToken(GramarParserRFLOAT, 0)
}

func (s *EmbebidaContext) RSTRING() antlr.TerminalNode {
	return s.GetToken(GramarParserRSTRING, 0)
}

func (s *EmbebidaContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GramarVisitor:
		return t.VisitEmbebida(s)

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
	_startState := 58
	p.EnterRecursionRule(localctx, 58, GramarParserRULE_expr, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(419)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 41, p.GetParserRuleContext()) {
	case 1:
		localctx = NewOpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(369)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*OpContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == GramarParserT__22 || _la == GramarParserT__23) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*OpContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(370)

			var _x = p.expr(20)

			localctx.(*OpContext).right = _x
		}

	case 2:
		localctx = NewLlamadaContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(371)
			p.Pllamada()
		}

	case 3:
		localctx = NewParenContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(372)
			p.Match(GramarParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(373)
			p.expr(0)
		}
		{
			p.SetState(374)
			p.Match(GramarParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		localctx = NewLiteralContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(376)
			p.Match(GramarParserINT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		localctx = NewLiteralContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(377)

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

	case 6:
		localctx = NewLiteralContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(378)
			p.Match(GramarParserCADENA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 7:
		localctx = NewLiteralContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(379)
			p.Match(GramarParserDOUBLE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 8:
		localctx = NewAccesoContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(380)
			p.Match(GramarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 9:
		localctx = NewEspecialContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(381)
			p.Match(GramarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(382)
			p.Match(GramarParserT__20)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(383)
			p.Match(GramarParserEMPT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(385)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 36, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(384)
				p.Match(GramarParserT__1)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}

	case 10:
		localctx = NewEspecialContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(387)
			p.Match(GramarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(388)
			p.Match(GramarParserT__20)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(389)
			p.Match(GramarParserCOUNT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(391)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 37, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(390)
				p.Match(GramarParserT__1)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}

	case 11:
		localctx = NewAccesoAContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(393)
			p.Match(GramarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(394)
			p.Match(GramarParserABR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(395)
			p.expr(0)
		}
		{
			p.SetState(396)
			p.Match(GramarParserCIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(398)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 38, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(397)
				p.Match(GramarParserT__1)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}

	case 12:
		localctx = NewEmbebidaContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(400)
			p.Match(GramarParserRINT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(401)
			p.Match(GramarParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(402)
			p.expr(0)
		}
		{
			p.SetState(403)
			p.Match(GramarParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(405)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 39, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(404)
				p.Match(GramarParserT__1)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}

	case 13:
		localctx = NewEmbebidaContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(407)
			p.Match(GramarParserRFLOAT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(408)
			p.Match(GramarParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(409)
			p.expr(0)
		}
		{
			p.SetState(410)
			p.Match(GramarParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 14:
		localctx = NewEmbebidaContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(412)
			p.Match(GramarParserRSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(413)
			p.Match(GramarParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(414)
			p.expr(0)
		}
		{
			p.SetState(415)
			p.Match(GramarParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(417)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 40, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(416)
				p.Match(GramarParserT__1)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(441)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 43, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(439)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 42, p.GetParserRuleContext()) {
			case 1:
				localctx = NewOpContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*OpContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, GramarParserRULE_expr)
				p.SetState(421)

				if !(p.Precpred(p.GetParserRuleContext(), 19)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 19)", ""))
					goto errorExit
				}
				{
					p.SetState(422)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*OpContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&234881024) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*OpContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(423)

					var _x = p.expr(20)

					localctx.(*OpContext).right = _x
				}

			case 2:
				localctx = NewOpContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*OpContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, GramarParserRULE_expr)
				p.SetState(424)

				if !(p.Precpred(p.GetParserRuleContext(), 18)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 18)", ""))
					goto errorExit
				}
				{
					p.SetState(425)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*OpContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == GramarParserT__23 || _la == GramarParserT__27) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*OpContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(426)

					var _x = p.expr(19)

					localctx.(*OpContext).right = _x
				}

			case 3:
				localctx = NewOpContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*OpContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, GramarParserRULE_expr)
				p.SetState(427)

				if !(p.Precpred(p.GetParserRuleContext(), 17)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 17)", ""))
					goto errorExit
				}
				{
					p.SetState(428)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*OpContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&8053063680) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*OpContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(429)

					var _x = p.expr(18)

					localctx.(*OpContext).right = _x
				}

			case 4:
				localctx = NewOpContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*OpContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, GramarParserRULE_expr)
				p.SetState(430)

				if !(p.Precpred(p.GetParserRuleContext(), 16)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 16)", ""))
					goto errorExit
				}
				{
					p.SetState(431)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*OpContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == GramarParserT__32 || _la == GramarParserT__33) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*OpContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(432)

					var _x = p.expr(17)

					localctx.(*OpContext).right = _x
				}

			case 5:
				localctx = NewOpContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*OpContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, GramarParserRULE_expr)
				p.SetState(433)

				if !(p.Precpred(p.GetParserRuleContext(), 15)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 15)", ""))
					goto errorExit
				}
				{
					p.SetState(434)

					var _m = p.Match(GramarParserT__34)

					localctx.(*OpContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(435)

					var _x = p.expr(16)

					localctx.(*OpContext).right = _x
				}

			case 6:
				localctx = NewOpContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*OpContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, GramarParserRULE_expr)
				p.SetState(436)

				if !(p.Precpred(p.GetParserRuleContext(), 14)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 14)", ""))
					goto errorExit
				}
				{
					p.SetState(437)

					var _m = p.Match(GramarParserT__35)

					localctx.(*OpContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(438)

					var _x = p.expr(15)

					localctx.(*OpContext).right = _x
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(443)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 43, p.GetParserRuleContext())
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
	case 29:
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
		return p.Precpred(p.GetParserRuleContext(), 19)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 18)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 17)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 16)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 15)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 14)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
