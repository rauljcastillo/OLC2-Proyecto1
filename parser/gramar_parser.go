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
		"'if'", "'?'", "'+='", "'-='", "'.'", "'at:'", "'!'", "'-'", "'*'",
		"'/'", "'%'", "'+'", "'<'", "'>'", "'<='", "'>='", "'=='", "'!='", "'&&'",
		"'||'", "", "", "", "'print'", "'true'", "'false'", "'Int'", "'Float'",
		"'String'", "'Bool'", "'Character'", "'var'", "'return'", "'continue'",
		"'break'", "'append'", "'remove'", "'removeLast'", "'IsEmpty'", "'count'",
		"'['", "']'", "'_'", "'...'", "'inout'", "'&'", "','", "':'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "WS", "COMO", "COMM", "PRINT", "TRUE", "FALSE", "RINT",
		"RFLOAT", "RSTRING", "RBOOL", "RChar", "RVAR", "RETN", "CONT", "BRK",
		"APPE", "REMOV", "REMLST", "EMPT", "COUNT", "ABR", "CIER", "GUION",
		"POINTS", "INOUT", "PUNTE", "COMMA", "TPOINTS", "CADENA", "ID", "INT",
		"DOUBLE", "CHARAC",
	}
	staticData.RuleNames = []string{
		"s", "block", "production", "preturn", "pasignA", "pfuncion", "pllamada",
		"parguments", "pargum", "pparams", "pparamet", "pdeclarArray", "pdefinition",
		"pguard", "pfor", "pifor", "pwhile", "pswitch", "ccase", "pdefault",
		"stmt", "pif", "pelse", "prin", "declara", "asign", "tipo", "pespecial",
		"expr",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 69, 421, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 1, 0, 1, 0, 1, 0, 1, 1, 5, 1, 63, 8,
		1, 10, 1, 12, 1, 66, 9, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1,
		2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 84, 8, 2, 1, 3,
		1, 3, 3, 3, 88, 8, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 97,
		8, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 105, 8, 5, 1, 5, 1, 5,
		1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 118, 8, 5,
		1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 124, 8, 5, 1, 6, 1, 6, 1, 6, 1, 6, 3, 6,
		130, 8, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 3, 6, 137, 8, 6, 3, 6, 139, 8,
		6, 1, 7, 1, 7, 1, 7, 5, 7, 144, 8, 7, 10, 7, 12, 7, 147, 9, 7, 1, 8, 3,
		8, 150, 8, 8, 1, 8, 1, 8, 1, 8, 1, 8, 3, 8, 156, 8, 8, 1, 9, 1, 9, 1, 9,
		5, 9, 161, 8, 9, 10, 9, 12, 9, 164, 9, 9, 3, 9, 166, 8, 9, 1, 10, 1, 10,
		1, 10, 1, 10, 3, 10, 172, 8, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 3,
		10, 179, 8, 10, 1, 10, 1, 10, 1, 10, 1, 10, 3, 10, 185, 8, 10, 1, 11, 1,
		11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 3, 11,
		198, 8, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 5, 12, 205, 8, 12, 10, 12,
		12, 12, 208, 9, 12, 3, 12, 210, 8, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12,
		1, 12, 3, 12, 218, 8, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1,
		13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 15,
		1, 15, 1, 15, 1, 15, 1, 15, 3, 15, 241, 8, 15, 1, 16, 1, 16, 1, 16, 1,
		16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 17, 4, 17, 253, 8, 17, 11, 17,
		12, 17, 254, 1, 17, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 3,
		18, 265, 8, 18, 1, 19, 1, 19, 1, 19, 1, 19, 3, 19, 271, 8, 19, 1, 20, 1,
		20, 1, 20, 1, 20, 1, 20, 1, 20, 3, 20, 279, 8, 20, 1, 21, 1, 21, 1, 21,
		1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 3, 21, 290, 8, 21, 1, 22, 1,
		22, 1, 22, 1, 22, 3, 22, 296, 8, 22, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23,
		3, 23, 303, 8, 23, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 3,
		24, 312, 8, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24,
		1, 24, 1, 24, 3, 24, 324, 8, 24, 1, 25, 1, 25, 1, 25, 1, 25, 1, 26, 1,
		26, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 3, 27, 339, 8, 27,
		1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 3, 27, 347, 8, 27, 1, 27, 1,
		27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 3, 27, 357, 8, 27, 3, 27,
		359, 8, 27, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1,
		28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 3, 28, 377, 8, 28,
		1, 28, 1, 28, 1, 28, 1, 28, 3, 28, 383, 8, 28, 1, 28, 1, 28, 1, 28, 1,
		28, 1, 28, 3, 28, 390, 8, 28, 1, 28, 3, 28, 393, 8, 28, 1, 28, 1, 28, 1,
		28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28,
		1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 5, 28, 416, 8,
		28, 10, 28, 12, 28, 419, 9, 28, 1, 28, 0, 1, 56, 29, 0, 2, 4, 6, 8, 10,
		12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46,
		48, 50, 52, 54, 56, 0, 10, 2, 0, 59, 59, 66, 66, 1, 0, 49, 51, 2, 0, 1,
		1, 19, 20, 1, 0, 43, 47, 1, 0, 23, 24, 1, 0, 41, 42, 1, 0, 25, 27, 2, 0,
		24, 24, 28, 28, 1, 0, 29, 32, 1, 0, 33, 34, 463, 0, 58, 1, 0, 0, 0, 2,
		64, 1, 0, 0, 0, 4, 83, 1, 0, 0, 0, 6, 85, 1, 0, 0, 0, 8, 89, 1, 0, 0, 0,
		10, 123, 1, 0, 0, 0, 12, 138, 1, 0, 0, 0, 14, 140, 1, 0, 0, 0, 16, 155,
		1, 0, 0, 0, 18, 165, 1, 0, 0, 0, 20, 184, 1, 0, 0, 0, 22, 197, 1, 0, 0,
		0, 24, 217, 1, 0, 0, 0, 26, 219, 1, 0, 0, 0, 28, 227, 1, 0, 0, 0, 30, 240,
		1, 0, 0, 0, 32, 242, 1, 0, 0, 0, 34, 248, 1, 0, 0, 0, 36, 259, 1, 0, 0,
		0, 38, 266, 1, 0, 0, 0, 40, 278, 1, 0, 0, 0, 42, 289, 1, 0, 0, 0, 44, 295,
		1, 0, 0, 0, 46, 297, 1, 0, 0, 0, 48, 323, 1, 0, 0, 0, 50, 325, 1, 0, 0,
		0, 52, 329, 1, 0, 0, 0, 54, 358, 1, 0, 0, 0, 56, 392, 1, 0, 0, 0, 58, 59,
		3, 2, 1, 0, 59, 60, 5, 0, 0, 1, 60, 1, 1, 0, 0, 0, 61, 63, 3, 4, 2, 0,
		62, 61, 1, 0, 0, 0, 63, 66, 1, 0, 0, 0, 64, 62, 1, 0, 0, 0, 64, 65, 1,
		0, 0, 0, 65, 3, 1, 0, 0, 0, 66, 64, 1, 0, 0, 0, 67, 84, 3, 46, 23, 0, 68,
		84, 3, 8, 4, 0, 69, 84, 3, 48, 24, 0, 70, 84, 3, 50, 25, 0, 71, 84, 3,
		42, 21, 0, 72, 84, 3, 34, 17, 0, 73, 84, 3, 32, 16, 0, 74, 84, 3, 28, 14,
		0, 75, 84, 3, 26, 13, 0, 76, 84, 3, 22, 11, 0, 77, 84, 3, 54, 27, 0, 78,
		84, 3, 10, 5, 0, 79, 84, 3, 12, 6, 0, 80, 84, 3, 6, 3, 0, 81, 84, 5, 50,
		0, 0, 82, 84, 5, 51, 0, 0, 83, 67, 1, 0, 0, 0, 83, 68, 1, 0, 0, 0, 83,
		69, 1, 0, 0, 0, 83, 70, 1, 0, 0, 0, 83, 71, 1, 0, 0, 0, 83, 72, 1, 0, 0,
		0, 83, 73, 1, 0, 0, 0, 83, 74, 1, 0, 0, 0, 83, 75, 1, 0, 0, 0, 83, 76,
		1, 0, 0, 0, 83, 77, 1, 0, 0, 0, 83, 78, 1, 0, 0, 0, 83, 79, 1, 0, 0, 0,
		83, 80, 1, 0, 0, 0, 83, 81, 1, 0, 0, 0, 83, 82, 1, 0, 0, 0, 84, 5, 1, 0,
		0, 0, 85, 87, 5, 49, 0, 0, 86, 88, 3, 56, 28, 0, 87, 86, 1, 0, 0, 0, 87,
		88, 1, 0, 0, 0, 88, 7, 1, 0, 0, 0, 89, 90, 5, 66, 0, 0, 90, 91, 5, 57,
		0, 0, 91, 92, 3, 56, 28, 0, 92, 93, 5, 58, 0, 0, 93, 94, 5, 1, 0, 0, 94,
		96, 3, 56, 28, 0, 95, 97, 5, 2, 0, 0, 96, 95, 1, 0, 0, 0, 96, 97, 1, 0,
		0, 0, 97, 9, 1, 0, 0, 0, 98, 99, 5, 3, 0, 0, 99, 100, 5, 66, 0, 0, 100,
		101, 5, 4, 0, 0, 101, 104, 5, 5, 0, 0, 102, 103, 5, 6, 0, 0, 103, 105,
		3, 52, 26, 0, 104, 102, 1, 0, 0, 0, 104, 105, 1, 0, 0, 0, 105, 106, 1,
		0, 0, 0, 106, 107, 5, 7, 0, 0, 107, 108, 3, 2, 1, 0, 108, 109, 5, 8, 0,
		0, 109, 124, 1, 0, 0, 0, 110, 111, 5, 3, 0, 0, 111, 112, 5, 66, 0, 0, 112,
		113, 5, 4, 0, 0, 113, 114, 3, 18, 9, 0, 114, 117, 5, 5, 0, 0, 115, 116,
		5, 6, 0, 0, 116, 118, 3, 52, 26, 0, 117, 115, 1, 0, 0, 0, 117, 118, 1,
		0, 0, 0, 118, 119, 1, 0, 0, 0, 119, 120, 5, 7, 0, 0, 120, 121, 3, 2, 1,
		0, 121, 122, 5, 8, 0, 0, 122, 124, 1, 0, 0, 0, 123, 98, 1, 0, 0, 0, 123,
		110, 1, 0, 0, 0, 124, 11, 1, 0, 0, 0, 125, 126, 5, 66, 0, 0, 126, 127,
		5, 4, 0, 0, 127, 129, 5, 5, 0, 0, 128, 130, 5, 2, 0, 0, 129, 128, 1, 0,
		0, 0, 129, 130, 1, 0, 0, 0, 130, 139, 1, 0, 0, 0, 131, 132, 5, 66, 0, 0,
		132, 133, 5, 4, 0, 0, 133, 134, 3, 14, 7, 0, 134, 136, 5, 5, 0, 0, 135,
		137, 5, 2, 0, 0, 136, 135, 1, 0, 0, 0, 136, 137, 1, 0, 0, 0, 137, 139,
		1, 0, 0, 0, 138, 125, 1, 0, 0, 0, 138, 131, 1, 0, 0, 0, 139, 13, 1, 0,
		0, 0, 140, 145, 3, 16, 8, 0, 141, 142, 5, 63, 0, 0, 142, 144, 3, 16, 8,
		0, 143, 141, 1, 0, 0, 0, 144, 147, 1, 0, 0, 0, 145, 143, 1, 0, 0, 0, 145,
		146, 1, 0, 0, 0, 146, 15, 1, 0, 0, 0, 147, 145, 1, 0, 0, 0, 148, 150, 5,
		62, 0, 0, 149, 148, 1, 0, 0, 0, 149, 150, 1, 0, 0, 0, 150, 151, 1, 0, 0,
		0, 151, 156, 3, 56, 28, 0, 152, 153, 5, 66, 0, 0, 153, 154, 5, 64, 0, 0,
		154, 156, 3, 56, 28, 0, 155, 149, 1, 0, 0, 0, 155, 152, 1, 0, 0, 0, 156,
		17, 1, 0, 0, 0, 157, 162, 3, 20, 10, 0, 158, 159, 5, 63, 0, 0, 159, 161,
		3, 20, 10, 0, 160, 158, 1, 0, 0, 0, 161, 164, 1, 0, 0, 0, 162, 160, 1,
		0, 0, 0, 162, 163, 1, 0, 0, 0, 163, 166, 1, 0, 0, 0, 164, 162, 1, 0, 0,
		0, 165, 157, 1, 0, 0, 0, 165, 166, 1, 0, 0, 0, 166, 19, 1, 0, 0, 0, 167,
		168, 7, 0, 0, 0, 168, 169, 5, 66, 0, 0, 169, 171, 5, 64, 0, 0, 170, 172,
		5, 61, 0, 0, 171, 170, 1, 0, 0, 0, 171, 172, 1, 0, 0, 0, 172, 173, 1, 0,
		0, 0, 173, 185, 3, 52, 26, 0, 174, 175, 7, 0, 0, 0, 175, 176, 5, 66, 0,
		0, 176, 178, 5, 64, 0, 0, 177, 179, 5, 61, 0, 0, 178, 177, 1, 0, 0, 0,
		178, 179, 1, 0, 0, 0, 179, 180, 1, 0, 0, 0, 180, 181, 5, 57, 0, 0, 181,
		182, 3, 52, 26, 0, 182, 183, 5, 58, 0, 0, 183, 185, 1, 0, 0, 0, 184, 167,
		1, 0, 0, 0, 184, 174, 1, 0, 0, 0, 185, 21, 1, 0, 0, 0, 186, 187, 5, 48,
		0, 0, 187, 188, 5, 66, 0, 0, 188, 189, 5, 64, 0, 0, 189, 190, 5, 57, 0,
		0, 190, 191, 3, 52, 26, 0, 191, 192, 5, 58, 0, 0, 192, 193, 3, 24, 12,
		0, 193, 198, 1, 0, 0, 0, 194, 195, 5, 48, 0, 0, 195, 196, 5, 66, 0, 0,
		196, 198, 3, 24, 12, 0, 197, 186, 1, 0, 0, 0, 197, 194, 1, 0, 0, 0, 198,
		23, 1, 0, 0, 0, 199, 200, 5, 1, 0, 0, 200, 209, 5, 57, 0, 0, 201, 206,
		3, 56, 28, 0, 202, 203, 5, 63, 0, 0, 203, 205, 3, 56, 28, 0, 204, 202,
		1, 0, 0, 0, 205, 208, 1, 0, 0, 0, 206, 204, 1, 0, 0, 0, 206, 207, 1, 0,
		0, 0, 207, 210, 1, 0, 0, 0, 208, 206, 1, 0, 0, 0, 209, 201, 1, 0, 0, 0,
		209, 210, 1, 0, 0, 0, 210, 211, 1, 0, 0, 0, 211, 218, 5, 58, 0, 0, 212,
		213, 5, 1, 0, 0, 213, 214, 5, 57, 0, 0, 214, 215, 3, 56, 28, 0, 215, 216,
		5, 58, 0, 0, 216, 218, 1, 0, 0, 0, 217, 199, 1, 0, 0, 0, 217, 212, 1, 0,
		0, 0, 218, 25, 1, 0, 0, 0, 219, 220, 5, 9, 0, 0, 220, 221, 3, 56, 28, 0,
		221, 222, 5, 10, 0, 0, 222, 223, 5, 7, 0, 0, 223, 224, 3, 2, 1, 0, 224,
		225, 7, 1, 0, 0, 225, 226, 5, 8, 0, 0, 226, 27, 1, 0, 0, 0, 227, 228, 5,
		11, 0, 0, 228, 229, 5, 66, 0, 0, 229, 230, 5, 12, 0, 0, 230, 231, 3, 30,
		15, 0, 231, 232, 5, 7, 0, 0, 232, 233, 3, 2, 1, 0, 233, 234, 5, 8, 0, 0,
		234, 29, 1, 0, 0, 0, 235, 236, 3, 56, 28, 0, 236, 237, 5, 60, 0, 0, 237,
		238, 3, 56, 28, 0, 238, 241, 1, 0, 0, 0, 239, 241, 3, 56, 28, 0, 240, 235,
		1, 0, 0, 0, 240, 239, 1, 0, 0, 0, 241, 31, 1, 0, 0, 0, 242, 243, 5, 13,
		0, 0, 243, 244, 3, 56, 28, 0, 244, 245, 5, 7, 0, 0, 245, 246, 3, 2, 1,
		0, 246, 247, 5, 8, 0, 0, 247, 33, 1, 0, 0, 0, 248, 249, 5, 14, 0, 0, 249,
		250, 3, 56, 28, 0, 250, 252, 5, 7, 0, 0, 251, 253, 3, 36, 18, 0, 252, 251,
		1, 0, 0, 0, 253, 254, 1, 0, 0, 0, 254, 252, 1, 0, 0, 0, 254, 255, 1, 0,
		0, 0, 255, 256, 1, 0, 0, 0, 256, 257, 3, 38, 19, 0, 257, 258, 5, 8, 0,
		0, 258, 35, 1, 0, 0, 0, 259, 260, 5, 15, 0, 0, 260, 261, 3, 56, 28, 0,
		261, 262, 5, 64, 0, 0, 262, 264, 3, 2, 1, 0, 263, 265, 5, 51, 0, 0, 264,
		263, 1, 0, 0, 0, 264, 265, 1, 0, 0, 0, 265, 37, 1, 0, 0, 0, 266, 267, 5,
		16, 0, 0, 267, 268, 5, 64, 0, 0, 268, 270, 3, 2, 1, 0, 269, 271, 5, 51,
		0, 0, 270, 269, 1, 0, 0, 0, 270, 271, 1, 0, 0, 0, 271, 39, 1, 0, 0, 0,
		272, 273, 5, 7, 0, 0, 273, 274, 3, 2, 1, 0, 274, 275, 5, 8, 0, 0, 275,
		279, 1, 0, 0, 0, 276, 277, 5, 7, 0, 0, 277, 279, 5, 8, 0, 0, 278, 272,
		1, 0, 0, 0, 278, 276, 1, 0, 0, 0, 279, 41, 1, 0, 0, 0, 280, 281, 5, 17,
		0, 0, 281, 282, 3, 56, 28, 0, 282, 283, 3, 40, 20, 0, 283, 290, 1, 0, 0,
		0, 284, 285, 5, 17, 0, 0, 285, 286, 3, 56, 28, 0, 286, 287, 3, 40, 20,
		0, 287, 288, 3, 44, 22, 0, 288, 290, 1, 0, 0, 0, 289, 280, 1, 0, 0, 0,
		289, 284, 1, 0, 0, 0, 290, 43, 1, 0, 0, 0, 291, 292, 5, 10, 0, 0, 292,
		296, 3, 42, 21, 0, 293, 294, 5, 10, 0, 0, 294, 296, 3, 40, 20, 0, 295,
		291, 1, 0, 0, 0, 295, 293, 1, 0, 0, 0, 296, 45, 1, 0, 0, 0, 297, 298, 5,
		40, 0, 0, 298, 299, 5, 4, 0, 0, 299, 300, 3, 56, 28, 0, 300, 302, 5, 5,
		0, 0, 301, 303, 5, 2, 0, 0, 302, 301, 1, 0, 0, 0, 302, 303, 1, 0, 0, 0,
		303, 47, 1, 0, 0, 0, 304, 305, 5, 48, 0, 0, 305, 306, 5, 66, 0, 0, 306,
		307, 5, 64, 0, 0, 307, 308, 3, 52, 26, 0, 308, 309, 5, 1, 0, 0, 309, 311,
		3, 56, 28, 0, 310, 312, 5, 2, 0, 0, 311, 310, 1, 0, 0, 0, 311, 312, 1,
		0, 0, 0, 312, 324, 1, 0, 0, 0, 313, 314, 5, 48, 0, 0, 314, 315, 5, 66,
		0, 0, 315, 316, 5, 64, 0, 0, 316, 317, 3, 52, 26, 0, 317, 318, 5, 18, 0,
		0, 318, 324, 1, 0, 0, 0, 319, 320, 5, 48, 0, 0, 320, 321, 5, 66, 0, 0,
		321, 322, 5, 1, 0, 0, 322, 324, 3, 56, 28, 0, 323, 304, 1, 0, 0, 0, 323,
		313, 1, 0, 0, 0, 323, 319, 1, 0, 0, 0, 324, 49, 1, 0, 0, 0, 325, 326, 5,
		66, 0, 0, 326, 327, 7, 2, 0, 0, 327, 328, 3, 56, 28, 0, 328, 51, 1, 0,
		0, 0, 329, 330, 7, 3, 0, 0, 330, 53, 1, 0, 0, 0, 331, 332, 5, 66, 0, 0,
		332, 333, 5, 21, 0, 0, 333, 334, 5, 52, 0, 0, 334, 335, 5, 4, 0, 0, 335,
		336, 3, 56, 28, 0, 336, 338, 5, 5, 0, 0, 337, 339, 5, 2, 0, 0, 338, 337,
		1, 0, 0, 0, 338, 339, 1, 0, 0, 0, 339, 359, 1, 0, 0, 0, 340, 341, 5, 66,
		0, 0, 341, 342, 5, 21, 0, 0, 342, 343, 5, 54, 0, 0, 343, 344, 5, 4, 0,
		0, 344, 346, 5, 5, 0, 0, 345, 347, 5, 2, 0, 0, 346, 345, 1, 0, 0, 0, 346,
		347, 1, 0, 0, 0, 347, 359, 1, 0, 0, 0, 348, 349, 5, 66, 0, 0, 349, 350,
		5, 21, 0, 0, 350, 351, 5, 53, 0, 0, 351, 352, 5, 4, 0, 0, 352, 353, 5,
		22, 0, 0, 353, 354, 3, 56, 28, 0, 354, 356, 5, 5, 0, 0, 355, 357, 5, 2,
		0, 0, 356, 355, 1, 0, 0, 0, 356, 357, 1, 0, 0, 0, 357, 359, 1, 0, 0, 0,
		358, 331, 1, 0, 0, 0, 358, 340, 1, 0, 0, 0, 358, 348, 1, 0, 0, 0, 359,
		55, 1, 0, 0, 0, 360, 361, 6, 28, -1, 0, 361, 362, 7, 4, 0, 0, 362, 393,
		3, 56, 28, 18, 363, 364, 5, 4, 0, 0, 364, 365, 3, 56, 28, 0, 365, 366,
		5, 5, 0, 0, 366, 393, 1, 0, 0, 0, 367, 393, 5, 67, 0, 0, 368, 393, 7, 5,
		0, 0, 369, 393, 5, 65, 0, 0, 370, 393, 5, 68, 0, 0, 371, 393, 5, 66, 0,
		0, 372, 373, 5, 66, 0, 0, 373, 374, 5, 21, 0, 0, 374, 376, 5, 55, 0, 0,
		375, 377, 5, 2, 0, 0, 376, 375, 1, 0, 0, 0, 376, 377, 1, 0, 0, 0, 377,
		393, 1, 0, 0, 0, 378, 379, 5, 66, 0, 0, 379, 380, 5, 21, 0, 0, 380, 382,
		5, 56, 0, 0, 381, 383, 5, 2, 0, 0, 382, 381, 1, 0, 0, 0, 382, 383, 1, 0,
		0, 0, 383, 393, 1, 0, 0, 0, 384, 385, 5, 66, 0, 0, 385, 386, 5, 57, 0,
		0, 386, 387, 3, 56, 28, 0, 387, 389, 5, 58, 0, 0, 388, 390, 5, 2, 0, 0,
		389, 388, 1, 0, 0, 0, 389, 390, 1, 0, 0, 0, 390, 393, 1, 0, 0, 0, 391,
		393, 3, 12, 6, 0, 392, 360, 1, 0, 0, 0, 392, 363, 1, 0, 0, 0, 392, 367,
		1, 0, 0, 0, 392, 368, 1, 0, 0, 0, 392, 369, 1, 0, 0, 0, 392, 370, 1, 0,
		0, 0, 392, 371, 1, 0, 0, 0, 392, 372, 1, 0, 0, 0, 392, 378, 1, 0, 0, 0,
		392, 384, 1, 0, 0, 0, 392, 391, 1, 0, 0, 0, 393, 417, 1, 0, 0, 0, 394,
		395, 10, 17, 0, 0, 395, 396, 7, 6, 0, 0, 396, 416, 3, 56, 28, 18, 397,
		398, 10, 16, 0, 0, 398, 399, 7, 7, 0, 0, 399, 416, 3, 56, 28, 17, 400,
		401, 10, 15, 0, 0, 401, 402, 7, 8, 0, 0, 402, 416, 3, 56, 28, 16, 403,
		404, 10, 14, 0, 0, 404, 405, 7, 9, 0, 0, 405, 416, 3, 56, 28, 15, 406,
		407, 10, 13, 0, 0, 407, 408, 5, 35, 0, 0, 408, 416, 3, 56, 28, 14, 409,
		410, 10, 12, 0, 0, 410, 411, 5, 36, 0, 0, 411, 416, 3, 56, 28, 13, 412,
		413, 10, 11, 0, 0, 413, 414, 5, 63, 0, 0, 414, 416, 3, 56, 28, 12, 415,
		394, 1, 0, 0, 0, 415, 397, 1, 0, 0, 0, 415, 400, 1, 0, 0, 0, 415, 403,
		1, 0, 0, 0, 415, 406, 1, 0, 0, 0, 415, 409, 1, 0, 0, 0, 415, 412, 1, 0,
		0, 0, 416, 419, 1, 0, 0, 0, 417, 415, 1, 0, 0, 0, 417, 418, 1, 0, 0, 0,
		418, 57, 1, 0, 0, 0, 419, 417, 1, 0, 0, 0, 42, 64, 83, 87, 96, 104, 117,
		123, 129, 136, 138, 145, 149, 155, 162, 165, 171, 178, 184, 197, 206, 209,
		217, 240, 254, 264, 270, 278, 289, 295, 302, 311, 323, 338, 346, 356, 358,
		376, 382, 389, 392, 415, 417,
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
	GramarParserRETN    = 49
	GramarParserCONT    = 50
	GramarParserBRK     = 51
	GramarParserAPPE    = 52
	GramarParserREMOV   = 53
	GramarParserREMLST  = 54
	GramarParserEMPT    = 55
	GramarParserCOUNT   = 56
	GramarParserABR     = 57
	GramarParserCIER    = 58
	GramarParserGUION   = 59
	GramarParserPOINTS  = 60
	GramarParserINOUT   = 61
	GramarParserPUNTE   = 62
	GramarParserCOMMA   = 63
	GramarParserTPOINTS = 64
	GramarParserCADENA  = 65
	GramarParserID      = 66
	GramarParserINT     = 67
	GramarParserDOUBLE  = 68
	GramarParserCHARAC  = 69
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
	GramarParserRULE_parguments   = 7
	GramarParserRULE_pargum       = 8
	GramarParserRULE_pparams      = 9
	GramarParserRULE_pparamet     = 10
	GramarParserRULE_pdeclarArray = 11
	GramarParserRULE_pdefinition  = 12
	GramarParserRULE_pguard       = 13
	GramarParserRULE_pfor         = 14
	GramarParserRULE_pifor        = 15
	GramarParserRULE_pwhile       = 16
	GramarParserRULE_pswitch      = 17
	GramarParserRULE_ccase        = 18
	GramarParserRULE_pdefault     = 19
	GramarParserRULE_stmt         = 20
	GramarParserRULE_pif          = 21
	GramarParserRULE_pelse        = 22
	GramarParserRULE_prin         = 23
	GramarParserRULE_declara      = 24
	GramarParserRULE_asign        = 25
	GramarParserRULE_tipo         = 26
	GramarParserRULE_pespecial    = 27
	GramarParserRULE_expr         = 28
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
		p.SetState(58)
		p.Block()
	}
	{
		p.SetState(59)
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
	p.SetState(64)
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
				p.SetState(61)
				p.Production()
			}

		}
		p.SetState(66)
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
	p.SetState(83)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(67)
			p.Prin()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(68)
			p.PasignA()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(69)
			p.Declara()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(70)
			p.Asign()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(71)
			p.Pif()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(72)
			p.Pswitch()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(73)
			p.Pwhile()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(74)
			p.Pfor()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(75)
			p.Pguard()
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(76)
			p.PdeclarArray()
		}

	case 11:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(77)
			p.Pespecial()
		}

	case 12:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(78)
			p.Pfuncion()
		}

	case 13:
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(79)
			p.Pllamada()
		}

	case 14:
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(80)
			p.Preturn()
		}

	case 15:
		p.EnterOuterAlt(localctx, 15)
		{
			p.SetState(81)
			p.Match(GramarParserCONT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 16:
		p.EnterOuterAlt(localctx, 16)
		{
			p.SetState(82)
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
		p.SetState(85)
		p.Match(GramarParserRETN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(87)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(86)
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
		p.SetState(89)
		p.Match(GramarParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(90)
		p.Match(GramarParserABR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(91)
		p.expr(0)
	}
	{
		p.SetState(92)
		p.Match(GramarParserCIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(93)
		p.Match(GramarParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(94)
		p.expr(0)
	}
	p.SetState(96)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == GramarParserT__1 {
		{
			p.SetState(95)
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

	p.SetState(123)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
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
			p.Match(GramarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(100)
			p.Match(GramarParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(101)
			p.Match(GramarParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(104)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == GramarParserT__5 {
			{
				p.SetState(102)
				p.Match(GramarParserT__5)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(103)
				p.Tipo()
			}

		}
		{
			p.SetState(106)
			p.Match(GramarParserT__6)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(107)
			p.Block()
		}
		{
			p.SetState(108)
			p.Match(GramarParserT__7)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(110)
			p.Match(GramarParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(111)
			p.Match(GramarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(112)
			p.Match(GramarParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(113)
			p.Pparams()
		}
		{
			p.SetState(114)
			p.Match(GramarParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(117)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == GramarParserT__5 {
			{
				p.SetState(115)
				p.Match(GramarParserT__5)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(116)
				p.Tipo()
			}

		}
		{
			p.SetState(119)
			p.Match(GramarParserT__6)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(120)
			p.Block()
		}
		{
			p.SetState(121)
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
	Parguments() IPargumentsContext

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

func (s *PllamadaContext) Parguments() IPargumentsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPargumentsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPargumentsContext)
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
	p.SetState(138)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
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
			p.Match(GramarParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(127)
			p.Match(GramarParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(129)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(128)
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
			p.SetState(131)
			p.Match(GramarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(132)
			p.Match(GramarParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(133)
			p.Parguments()
		}
		{
			p.SetState(134)
			p.Match(GramarParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(136)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(135)
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

// IPargumentsContext is an interface to support dynamic dispatch.
type IPargumentsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllPargum() []IPargumContext
	Pargum(i int) IPargumContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsPargumentsContext differentiates from other interfaces.
	IsPargumentsContext()
}

type PargumentsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPargumentsContext() *PargumentsContext {
	var p = new(PargumentsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GramarParserRULE_parguments
	return p
}

func InitEmptyPargumentsContext(p *PargumentsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GramarParserRULE_parguments
}

func (*PargumentsContext) IsPargumentsContext() {}

func NewPargumentsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PargumentsContext {
	var p = new(PargumentsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramarParserRULE_parguments

	return p
}

func (s *PargumentsContext) GetParser() antlr.Parser { return s.parser }

func (s *PargumentsContext) AllPargum() []IPargumContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IPargumContext); ok {
			len++
		}
	}

	tst := make([]IPargumContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IPargumContext); ok {
			tst[i] = t.(IPargumContext)
			i++
		}
	}

	return tst
}

func (s *PargumentsContext) Pargum(i int) IPargumContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPargumContext); ok {
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

	return t.(IPargumContext)
}

func (s *PargumentsContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(GramarParserCOMMA)
}

func (s *PargumentsContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(GramarParserCOMMA, i)
}

func (s *PargumentsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PargumentsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PargumentsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GramarVisitor:
		return t.VisitParguments(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GramarParser) Parguments() (localctx IPargumentsContext) {
	localctx = NewPargumentsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, GramarParserRULE_parguments)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(140)
		p.Pargum()
	}
	p.SetState(145)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == GramarParserCOMMA {
		{
			p.SetState(141)
			p.Match(GramarParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(142)
			p.Pargum()
		}

		p.SetState(147)
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

// IPargumContext is an interface to support dynamic dispatch.
type IPargumContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expr() IExprContext
	PUNTE() antlr.TerminalNode
	ID() antlr.TerminalNode
	TPOINTS() antlr.TerminalNode

	// IsPargumContext differentiates from other interfaces.
	IsPargumContext()
}

type PargumContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPargumContext() *PargumContext {
	var p = new(PargumContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GramarParserRULE_pargum
	return p
}

func InitEmptyPargumContext(p *PargumContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GramarParserRULE_pargum
}

func (*PargumContext) IsPargumContext() {}

func NewPargumContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PargumContext {
	var p = new(PargumContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramarParserRULE_pargum

	return p
}

func (s *PargumContext) GetParser() antlr.Parser { return s.parser }

func (s *PargumContext) Expr() IExprContext {
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

func (s *PargumContext) PUNTE() antlr.TerminalNode {
	return s.GetToken(GramarParserPUNTE, 0)
}

func (s *PargumContext) ID() antlr.TerminalNode {
	return s.GetToken(GramarParserID, 0)
}

func (s *PargumContext) TPOINTS() antlr.TerminalNode {
	return s.GetToken(GramarParserTPOINTS, 0)
}

func (s *PargumContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PargumContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PargumContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GramarVisitor:
		return t.VisitPargum(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GramarParser) Pargum() (localctx IPargumContext) {
	localctx = NewPargumContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, GramarParserRULE_pargum)
	var _la int

	p.SetState(155)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 12, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(149)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == GramarParserPUNTE {
			{
				p.SetState(148)
				p.Match(GramarParserPUNTE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(151)
			p.expr(0)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(152)
			p.Match(GramarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(153)
			p.Match(GramarParserTPOINTS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(154)
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
	p.EnterRule(localctx, 18, GramarParserRULE_pparams)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(165)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == GramarParserGUION || _la == GramarParserID {
		{
			p.SetState(157)
			p.Pparamet()
		}
		p.SetState(162)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == GramarParserCOMMA {
			{
				p.SetState(158)
				p.Match(GramarParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(159)
				p.Pparamet()
			}

			p.SetState(164)
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
	p.EnterRule(localctx, 20, GramarParserRULE_pparamet)
	var _la int

	p.SetState(184)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 17, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(167)

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
			p.SetState(168)
			p.Match(GramarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(169)
			p.Match(GramarParserTPOINTS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(171)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == GramarParserINOUT {
			{
				p.SetState(170)
				p.Match(GramarParserINOUT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(173)
			p.Tipo()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(174)

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
			p.SetState(175)
			p.Match(GramarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(176)
			p.Match(GramarParserTPOINTS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(178)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == GramarParserINOUT {
			{
				p.SetState(177)
				p.Match(GramarParserINOUT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(180)
			p.Match(GramarParserABR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(181)
			p.Tipo()
		}
		{
			p.SetState(182)
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
	RVAR() antlr.TerminalNode
	ID() antlr.TerminalNode
	TPOINTS() antlr.TerminalNode
	ABR() antlr.TerminalNode
	Tipo() ITipoContext
	CIER() antlr.TerminalNode
	Pdefinition() IPdefinitionContext

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

func (s *PdeclarArrayContext) RVAR() antlr.TerminalNode {
	return s.GetToken(GramarParserRVAR, 0)
}

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
	p.EnterRule(localctx, 22, GramarParserRULE_pdeclarArray)
	p.SetState(197)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 18, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(186)
			p.Match(GramarParserRVAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(187)
			p.Match(GramarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(188)
			p.Match(GramarParserTPOINTS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(189)
			p.Match(GramarParserABR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(190)
			p.Tipo()
		}
		{
			p.SetState(191)
			p.Match(GramarParserCIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(192)
			p.Pdefinition()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(194)
			p.Match(GramarParserRVAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(195)
			p.Match(GramarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(196)
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
	p.EnterRule(localctx, 24, GramarParserRULE_pdefinition)
	var _la int

	p.SetState(217)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 21, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(199)
			p.Match(GramarParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(200)
			p.Match(GramarParserABR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(209)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&6597094932496) != 0) || ((int64((_la-65)) & ^0x3f) == 0 && ((int64(1)<<(_la-65))&15) != 0) {
			{
				p.SetState(201)
				p.expr(0)
			}
			p.SetState(206)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			for _la == GramarParserCOMMA {
				{
					p.SetState(202)
					p.Match(GramarParserCOMMA)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(203)
					p.expr(0)
				}

				p.SetState(208)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)
			}

		}
		{
			p.SetState(211)
			p.Match(GramarParserCIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(212)
			p.Match(GramarParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(213)
			p.Match(GramarParserABR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(214)
			p.expr(0)
		}
		{
			p.SetState(215)
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
	p.EnterRule(localctx, 26, GramarParserRULE_pguard)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(219)
		p.Match(GramarParserT__8)
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
		p.Match(GramarParserT__9)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(222)
		p.Match(GramarParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(223)
		p.Block()
	}
	{
		p.SetState(224)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*PguardContext).opera = _lt

		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&3940649673949184) != 0) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*PguardContext).opera = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(225)
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
	p.EnterRule(localctx, 28, GramarParserRULE_pfor)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(227)
		p.Match(GramarParserT__10)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(228)
		p.Match(GramarParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(229)
		p.Match(GramarParserT__11)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(230)
		p.Pifor()
	}
	{
		p.SetState(231)
		p.Match(GramarParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(232)
		p.Block()
	}
	{
		p.SetState(233)
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
	p.EnterRule(localctx, 30, GramarParserRULE_pifor)
	p.SetState(240)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 22, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(235)
			p.expr(0)
		}
		{
			p.SetState(236)
			p.Match(GramarParserPOINTS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(237)
			p.expr(0)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(239)
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
	p.EnterRule(localctx, 32, GramarParserRULE_pwhile)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(242)
		p.Match(GramarParserT__12)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(243)
		p.expr(0)
	}
	{
		p.SetState(244)
		p.Match(GramarParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(245)
		p.Block()
	}
	{
		p.SetState(246)
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
	p.EnterRule(localctx, 34, GramarParserRULE_pswitch)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(248)
		p.Match(GramarParserT__13)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(249)
		p.expr(0)
	}
	{
		p.SetState(250)
		p.Match(GramarParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(252)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == GramarParserT__14 {
		{
			p.SetState(251)
			p.Ccase()
		}

		p.SetState(254)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(256)
		p.Pdefault()
	}
	{
		p.SetState(257)
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
	p.EnterRule(localctx, 36, GramarParserRULE_ccase)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(259)
		p.Match(GramarParserT__14)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(260)
		p.expr(0)
	}
	{
		p.SetState(261)
		p.Match(GramarParserTPOINTS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(262)
		p.Block()
	}
	p.SetState(264)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == GramarParserBRK {
		{
			p.SetState(263)
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
	p.EnterRule(localctx, 38, GramarParserRULE_pdefault)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(266)
		p.Match(GramarParserT__15)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(267)
		p.Match(GramarParserTPOINTS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(268)
		p.Block()
	}
	p.SetState(270)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == GramarParserBRK {
		{
			p.SetState(269)
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
	p.EnterRule(localctx, 40, GramarParserRULE_stmt)
	p.SetState(278)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 26, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(272)
			p.Match(GramarParserT__6)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(273)
			p.Block()
		}
		{
			p.SetState(274)
			p.Match(GramarParserT__7)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(276)
			p.Match(GramarParserT__6)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(277)
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
	p.EnterRule(localctx, 42, GramarParserRULE_pif)
	p.SetState(289)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 27, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(280)
			p.Match(GramarParserT__16)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(281)
			p.expr(0)
		}
		{
			p.SetState(282)
			p.Stmt()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(284)
			p.Match(GramarParserT__16)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(285)
			p.expr(0)
		}
		{
			p.SetState(286)
			p.Stmt()
		}
		{
			p.SetState(287)
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
	p.EnterRule(localctx, 44, GramarParserRULE_pelse)
	p.SetState(295)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 28, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(291)
			p.Match(GramarParserT__9)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(292)
			p.Pif()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(293)
			p.Match(GramarParserT__9)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(294)
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
	p.EnterRule(localctx, 46, GramarParserRULE_prin)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(297)
		p.Match(GramarParserPRINT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(298)
		p.Match(GramarParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(299)
		p.expr(0)
	}
	{
		p.SetState(300)
		p.Match(GramarParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(302)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == GramarParserT__1 {
		{
			p.SetState(301)
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

// IDeclaraContext is an interface to support dynamic dispatch.
type IDeclaraContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	RVAR() antlr.TerminalNode
	ID() antlr.TerminalNode
	TPOINTS() antlr.TerminalNode
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
	p.EnterRule(localctx, 48, GramarParserRULE_declara)
	var _la int

	p.SetState(323)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 31, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(304)
			p.Match(GramarParserRVAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(305)
			p.Match(GramarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(306)
			p.Match(GramarParserTPOINTS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(307)
			p.Tipo()
		}
		{
			p.SetState(308)
			p.Match(GramarParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(309)
			p.expr(0)
		}
		p.SetState(311)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == GramarParserT__1 {
			{
				p.SetState(310)
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
			p.SetState(313)
			p.Match(GramarParserRVAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(314)
			p.Match(GramarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(315)
			p.Match(GramarParserTPOINTS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(316)
			p.Tipo()
		}
		{
			p.SetState(317)
			p.Match(GramarParserT__17)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(319)
			p.Match(GramarParserRVAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(320)
			p.Match(GramarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(321)
			p.Match(GramarParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(322)
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
	p.EnterRule(localctx, 50, GramarParserRULE_asign)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(325)
		p.Match(GramarParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(326)

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
		p.SetState(327)
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
	p.EnterRule(localctx, 52, GramarParserRULE_tipo)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(329)
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
	p.EnterRule(localctx, 54, GramarParserRULE_pespecial)
	var _la int

	p.SetState(358)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 35, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(331)
			p.Match(GramarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(332)
			p.Match(GramarParserT__20)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(333)
			p.Match(GramarParserAPPE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(334)
			p.Match(GramarParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(335)
			p.expr(0)
		}
		{
			p.SetState(336)
			p.Match(GramarParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(338)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == GramarParserT__1 {
			{
				p.SetState(337)
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
			p.SetState(340)
			p.Match(GramarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(341)
			p.Match(GramarParserT__20)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(342)
			p.Match(GramarParserREMLST)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(343)
			p.Match(GramarParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(344)
			p.Match(GramarParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(346)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == GramarParserT__1 {
			{
				p.SetState(345)
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
			p.SetState(348)
			p.Match(GramarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(349)
			p.Match(GramarParserT__20)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(350)
			p.Match(GramarParserREMOV)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(351)
			p.Match(GramarParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(352)
			p.Match(GramarParserT__21)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(353)
			p.expr(0)
		}
		{
			p.SetState(354)
			p.Match(GramarParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(356)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == GramarParserT__1 {
			{
				p.SetState(355)
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

func (s *OpContext) COMMA() antlr.TerminalNode {
	return s.GetToken(GramarParserCOMMA, 0)
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

func (p *GramarParser) Expr() (localctx IExprContext) {
	return p.expr(0)
}

func (p *GramarParser) expr(_p int) (localctx IExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 56
	p.EnterRecursionRule(localctx, 56, GramarParserRULE_expr, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(392)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 39, p.GetParserRuleContext()) {
	case 1:
		localctx = NewOpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(361)

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
			p.SetState(362)

			var _x = p.expr(18)

			localctx.(*OpContext).right = _x
		}

	case 2:
		localctx = NewParenContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(363)
			p.Match(GramarParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(364)
			p.expr(0)
		}
		{
			p.SetState(365)
			p.Match(GramarParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		localctx = NewLiteralContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(367)
			p.Match(GramarParserINT)
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
			p.SetState(368)

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

	case 5:
		localctx = NewLiteralContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(369)
			p.Match(GramarParserCADENA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 6:
		localctx = NewLiteralContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(370)
			p.Match(GramarParserDOUBLE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 7:
		localctx = NewAccesoContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(371)
			p.Match(GramarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 8:
		localctx = NewEspecialContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(372)
			p.Match(GramarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(373)
			p.Match(GramarParserT__20)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(374)
			p.Match(GramarParserEMPT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(376)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 36, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(375)
				p.Match(GramarParserT__1)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}

	case 9:
		localctx = NewEspecialContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(378)
			p.Match(GramarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(379)
			p.Match(GramarParserT__20)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(380)
			p.Match(GramarParserCOUNT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(382)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 37, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(381)
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
		localctx = NewAccesoAContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(384)
			p.Match(GramarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(385)
			p.Match(GramarParserABR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(386)
			p.expr(0)
		}
		{
			p.SetState(387)
			p.Match(GramarParserCIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(389)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 38, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(388)
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
		localctx = NewLlamadaContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(391)
			p.Pllamada()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(417)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 41, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(415)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 40, p.GetParserRuleContext()) {
			case 1:
				localctx = NewOpContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*OpContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, GramarParserRULE_expr)
				p.SetState(394)

				if !(p.Precpred(p.GetParserRuleContext(), 17)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 17)", ""))
					goto errorExit
				}
				{
					p.SetState(395)

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
					p.SetState(396)

					var _x = p.expr(18)

					localctx.(*OpContext).right = _x
				}

			case 2:
				localctx = NewOpContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*OpContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, GramarParserRULE_expr)
				p.SetState(397)

				if !(p.Precpred(p.GetParserRuleContext(), 16)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 16)", ""))
					goto errorExit
				}
				{
					p.SetState(398)

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
					p.SetState(399)

					var _x = p.expr(17)

					localctx.(*OpContext).right = _x
				}

			case 3:
				localctx = NewOpContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*OpContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, GramarParserRULE_expr)
				p.SetState(400)

				if !(p.Precpred(p.GetParserRuleContext(), 15)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 15)", ""))
					goto errorExit
				}
				{
					p.SetState(401)

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
					p.SetState(402)

					var _x = p.expr(16)

					localctx.(*OpContext).right = _x
				}

			case 4:
				localctx = NewOpContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*OpContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, GramarParserRULE_expr)
				p.SetState(403)

				if !(p.Precpred(p.GetParserRuleContext(), 14)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 14)", ""))
					goto errorExit
				}
				{
					p.SetState(404)

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
					p.SetState(405)

					var _x = p.expr(15)

					localctx.(*OpContext).right = _x
				}

			case 5:
				localctx = NewOpContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*OpContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, GramarParserRULE_expr)
				p.SetState(406)

				if !(p.Precpred(p.GetParserRuleContext(), 13)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 13)", ""))
					goto errorExit
				}
				{
					p.SetState(407)

					var _m = p.Match(GramarParserT__34)

					localctx.(*OpContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(408)

					var _x = p.expr(14)

					localctx.(*OpContext).right = _x
				}

			case 6:
				localctx = NewOpContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*OpContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, GramarParserRULE_expr)
				p.SetState(409)

				if !(p.Precpred(p.GetParserRuleContext(), 12)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 12)", ""))
					goto errorExit
				}
				{
					p.SetState(410)

					var _m = p.Match(GramarParserT__35)

					localctx.(*OpContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(411)

					var _x = p.expr(13)

					localctx.(*OpContext).right = _x
				}

			case 7:
				localctx = NewOpContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*OpContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, GramarParserRULE_expr)
				p.SetState(412)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
					goto errorExit
				}
				{
					p.SetState(413)

					var _m = p.Match(GramarParserCOMMA)

					localctx.(*OpContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(414)

					var _x = p.expr(12)

					localctx.(*OpContext).right = _x
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(419)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 41, p.GetParserRuleContext())
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
	case 28:
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
		return p.Precpred(p.GetParserRuleContext(), 17)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 16)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 15)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 14)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 13)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 12)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 11)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
