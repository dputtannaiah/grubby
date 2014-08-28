//line parser.y:2
package parser

import __yyfmt__ "fmt"

//line parser.y:3
import (
	"github.com/grubby/grubby/ast"
	"strings"
)

var Statements []ast.Node

//line parser.y:16
type RubySymType struct {
	yys          int
	operator     string
	genericValue ast.Node
	genericSlice ast.Nodes
	stringSlice  []string
}

const OPERATOR = 57346
const NODE = 57347
const REF = 57348
const CAPITAL_REF = 57349
const LPAREN = 57350
const RPAREN = 57351
const COMMA = 57352
const DO = 57353
const DEF = 57354
const END = 57355
const IF = 57356
const ELSE = 57357
const ELSIF = 57358
const CLASS = 57359
const MODULE = 57360
const TRUE = 57361
const FALSE = 57362
const LESSTHAN = 57363
const GREATERTHAN = 57364
const EQUALTO = 57365
const BANG = 57366
const COMPLEMENT = 57367
const POSITIVE = 57368
const NEGATIVE = 57369
const STAR = 57370
const WHITESPACE = 57371
const NEWLINE = 57372
const SEMICOLON = 57373
const COLON = 57374
const DOT = 57375
const PIPE = 57376
const SLASH = 57377
const AMPERSAND = 57378
const MODULO = 57379
const CARET = 57380
const LBRACKET = 57381
const RBRACKET = 57382
const LBRACE = 57383
const RBRACE = 57384
const DOLLARSIGN = 57385
const ATSIGN = 57386
const FILE_CONST_REF = 57387
const EOF = 57388

var RubyToknames = []string{
	"OPERATOR",
	"NODE",
	"REF",
	"CAPITAL_REF",
	"LPAREN",
	"RPAREN",
	"COMMA",
	"DO",
	"DEF",
	"END",
	"IF",
	"ELSE",
	"ELSIF",
	"CLASS",
	"MODULE",
	"TRUE",
	"FALSE",
	"LESSTHAN",
	"GREATERTHAN",
	"EQUALTO",
	"BANG",
	"COMPLEMENT",
	"POSITIVE",
	"NEGATIVE",
	"STAR",
	"WHITESPACE",
	"NEWLINE",
	"SEMICOLON",
	"COLON",
	"DOT",
	"PIPE",
	"SLASH",
	"AMPERSAND",
	"MODULO",
	"CARET",
	"LBRACKET",
	"RBRACKET",
	"LBRACE",
	"RBRACE",
	"DOLLARSIGN",
	"ATSIGN",
	"FILE_CONST_REF",
	"EOF",
}
var RubyStatenames = []string{}

const RubyEofCode = 1
const RubyErrCode = 2
const RubyMaxDepth = 200

//line parser.y:554

//line yacctab:1
var RubyExca = []int{
	-1, 1,
	1, -1,
	-2, 12,
	-1, 11,
	1, 17,
	4, 17,
	10, 17,
	13, 17,
	15, 17,
	16, 17,
	28, 17,
	30, 17,
	31, 17,
	33, 17,
	34, 17,
	35, 17,
	36, 17,
	42, 17,
	46, 17,
	-2, 12,
	-1, 12,
	23, 12,
	-2, 18,
	-1, 13,
	23, 12,
	-2, 19,
	-1, 14,
	23, 12,
	-2, 20,
	-1, 15,
	23, 12,
	-2, 21,
	-1, 82,
	5, 12,
	6, 12,
	8, 12,
	-2, 44,
	-1, 84,
	10, 12,
	-2, 46,
	-1, 87,
	10, 12,
	-2, 49,
	-1, 88,
	4, 12,
	26, 12,
	27, 12,
	28, 12,
	34, 12,
	35, 12,
	36, 12,
	-2, 50,
	-1, 90,
	1, 17,
	4, 17,
	10, 17,
	13, 17,
	15, 17,
	16, 17,
	28, 17,
	30, 17,
	31, 17,
	33, 17,
	34, 17,
	35, 17,
	36, 17,
	42, 17,
	46, 17,
	-2, 12,
	-1, 92,
	34, 12,
	-2, 8,
	-1, 103,
	4, 12,
	26, 12,
	27, 12,
	28, 12,
	34, 12,
	35, 12,
	36, 12,
	-2, 82,
	-1, 104,
	4, 12,
	26, 12,
	27, 12,
	28, 12,
	34, 12,
	35, 12,
	36, 12,
	-2, 83,
	-1, 105,
	4, 12,
	26, 12,
	27, 12,
	28, 12,
	34, 12,
	35, 12,
	36, 12,
	-2, 84,
	-1, 106,
	4, 12,
	26, 12,
	27, 12,
	28, 12,
	34, 12,
	35, 12,
	36, 12,
	-2, 85,
	-1, 111,
	4, 12,
	28, 12,
	34, 12,
	35, 12,
	36, 12,
	-2, 8,
	-1, 134,
	32, 75,
	-2, 12,
	-1, 141,
	4, 12,
	26, 12,
	27, 12,
	28, 12,
	34, 12,
	35, 12,
	36, 12,
	-2, 47,
	-1, 142,
	4, 12,
	26, 12,
	27, 12,
	28, 12,
	34, 12,
	35, 12,
	36, 12,
	-2, 86,
	-1, 143,
	4, 12,
	26, 12,
	27, 12,
	28, 12,
	34, 12,
	35, 12,
	36, 12,
	-2, 87,
	-1, 144,
	4, 12,
	26, 12,
	27, 12,
	28, 12,
	34, 12,
	35, 12,
	36, 12,
	-2, 88,
	-1, 145,
	4, 12,
	26, 12,
	27, 12,
	28, 12,
	34, 12,
	35, 12,
	36, 12,
	-2, 89,
	-1, 146,
	4, 12,
	26, 12,
	27, 12,
	28, 12,
	34, 12,
	35, 12,
	36, 12,
	-2, 90,
	-1, 147,
	4, 12,
	26, 12,
	27, 12,
	28, 12,
	34, 12,
	35, 12,
	36, 12,
	-2, 91,
	-1, 152,
	4, 12,
	26, 12,
	27, 12,
	28, 12,
	34, 12,
	35, 12,
	36, 12,
	-2, 77,
	-1, 157,
	4, 12,
	28, 12,
	34, 12,
	35, 12,
	36, 12,
	-2, 11,
	-1, 160,
	34, 116,
	-2, 59,
	-1, 161,
	4, 12,
	26, 12,
	27, 12,
	28, 12,
	34, 12,
	35, 12,
	36, 12,
	-2, 78,
	-1, 162,
	4, 12,
	26, 12,
	27, 12,
	28, 12,
	34, 12,
	35, 12,
	36, 12,
	-2, 79,
	-1, 163,
	4, 12,
	26, 12,
	27, 12,
	28, 12,
	34, 12,
	35, 12,
	36, 12,
	-2, 80,
	-1, 164,
	4, 12,
	26, 12,
	27, 12,
	28, 12,
	34, 12,
	35, 12,
	36, 12,
	-2, 81,
	-1, 177,
	4, 17,
	28, 17,
	33, 17,
	34, 17,
	35, 17,
	36, 17,
	-2, 12,
	-1, 187,
	34, 117,
	-2, 60,
	-1, 204,
	4, 12,
	26, 12,
	27, 12,
	28, 12,
	34, 12,
	35, 12,
	36, 12,
	-2, 52,
	-1, 218,
	32, 76,
	-2, 12,
	-1, 229,
	4, 12,
	28, 12,
	34, 12,
	35, 12,
	36, 12,
	-2, 8,
	-1, 241,
	10, 101,
	30, 101,
	42, 101,
	-2, 12,
	-1, 242,
	4, 12,
	28, 12,
	34, 12,
	35, 12,
	36, 12,
	-2, 8,
	-1, 253,
	34, 118,
	-2, 63,
	-1, 256,
	10, 98,
	30, 98,
	42, 98,
	-2, 12,
	-1, 262,
	30, 102,
	42, 102,
	-2, 12,
	-1, 265,
	30, 99,
	42, 99,
	-2, 12,
}

const RubyNprod = 128
const RubyPrivate = 57344

var RubyTokenNames []string
var RubyStates []string

const RubyLast = 742

var RubyAct = []int{

	124, 189, 6, 107, 125, 89, 100, 110, 52, 8,
	53, 83, 55, 56, 57, 58, 59, 91, 90, 12,
	86, 60, 61, 92, 39, 2, 51, 249, 247, 40,
	41, 42, 43, 172, 212, 85, 44, 45, 46, 47,
	65, 66, 67, 266, 160, 68, 69, 70, 71, 72,
	73, 48, 74, 49, 7, 38, 37, 50, 54, 62,
	258, 7, 7, 173, 198, 231, 54, 101, 101, 7,
	4, 5, 211, 192, 102, 170, 112, 113, 114, 115,
	116, 117, 118, 119, 264, 120, 121, 122, 123, 246,
	75, 55, 191, 126, 127, 128, 129, 130, 7, 138,
	131, 214, 215, 7, 7, 214, 215, 54, 137, 259,
	133, 139, 76, 77, 78, 136, 140, 96, 95, 132,
	81, 79, 80, 94, 93, 261, 153, 75, 240, 232,
	199, 148, 201, 200, 101, 169, 224, 167, 207, 208,
	168, 171, 222, 166, 92, 154, 225, 206, 172, 76,
	77, 78, 182, 151, 183, 184, 157, 81, 79, 80,
	185, 75, 218, 9, 188, 150, 149, 255, 86, 97,
	98, 63, 64, 194, 134, 195, 196, 197, 55, 219,
	253, 202, 99, 76, 77, 78, 203, 82, 205, 213,
	209, 81, 79, 80, 216, 210, 217, 219, 220, 226,
	135, 227, 84, 221, 223, 87, 228, 109, 187, 109,
	108, 175, 88, 186, 174, 230, 1, 165, 233, 235,
	179, 159, 236, 32, 238, 103, 104, 105, 106, 31,
	237, 111, 239, 245, 243, 30, 29, 28, 248, 27,
	250, 251, 26, 25, 24, 23, 35, 252, 19, 13,
	18, 17, 14, 20, 36, 257, 16, 15, 33, 260,
	22, 34, 263, 21, 3, 0, 0, 0, 0, 141,
	142, 143, 144, 145, 146, 147, 0, 0, 152, 0,
	0, 0, 0, 0, 161, 162, 163, 164, 0, 0,
	10, 11, 12, 0, 0, 0, 176, 39, 178, 51,
	181, 180, 40, 41, 42, 43, 0, 0, 0, 44,
	45, 46, 47, 0, 156, 155, 0, 0, 0, 0,
	0, 0, 0, 0, 48, 0, 49, 0, 38, 37,
	50, 0, 0, 10, 11, 12, 0, 0, 0, 204,
	39, 244, 51, 0, 0, 40, 41, 42, 43, 0,
	0, 0, 44, 45, 46, 47, 0, 156, 155, 229,
	0, 0, 0, 0, 0, 0, 0, 48, 0, 49,
	0, 38, 37, 50, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 241, 242, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 254, 0, 0, 256, 0,
	0, 0, 0, 0, 10, 11, 12, 262, 0, 0,
	265, 39, 234, 51, 0, 0, 40, 41, 42, 43,
	0, 0, 0, 44, 45, 46, 47, 0, 156, 155,
	0, 0, 0, 0, 0, 0, 0, 0, 48, 0,
	49, 0, 38, 37, 50, 10, 11, 12, 0, 0,
	0, 0, 39, 193, 51, 0, 0, 40, 41, 42,
	43, 0, 0, 0, 44, 45, 46, 47, 0, 156,
	155, 0, 0, 0, 0, 0, 0, 0, 0, 48,
	0, 49, 0, 38, 37, 50, 10, 11, 12, 0,
	0, 0, 0, 39, 190, 51, 0, 0, 40, 41,
	42, 43, 0, 0, 0, 44, 45, 46, 47, 0,
	156, 155, 0, 0, 0, 0, 0, 0, 0, 0,
	48, 0, 49, 0, 38, 37, 50, 10, 11, 12,
	0, 0, 0, 0, 39, 158, 51, 0, 0, 40,
	41, 42, 43, 0, 0, 0, 44, 45, 46, 47,
	0, 156, 155, 0, 0, 0, 0, 0, 0, 0,
	0, 48, 0, 49, 0, 38, 37, 50, 10, 11,
	12, 0, 0, 0, 0, 39, 0, 51, 0, 0,
	40, 41, 42, 43, 0, 0, 0, 44, 45, 46,
	47, 0, 156, 155, 0, 0, 0, 0, 0, 0,
	0, 0, 48, 0, 49, 0, 38, 37, 50, 10,
	11, 12, 0, 0, 0, 92, 39, 0, 51, 0,
	0, 40, 41, 42, 43, 0, 0, 0, 44, 45,
	46, 47, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 48, 0, 49, 0, 38, 37, 50,
	10, 11, 12, 0, 0, 0, 0, 39, 0, 51,
	0, 0, 40, 41, 42, 43, 0, 0, 0, 44,
	45, 46, 47, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 48, 0, 49, 0, 38, 37,
	50, 10, 177, 12, 0, 0, 0, 0, 39, 0,
	51, 0, 0, 40, 41, 42, 43, 0, 0, 0,
	44, 45, 46, 47, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 48, 0, 49, 0, 38,
	37, 50,
}
var RubyPact = []int{

	-21, 40, -1000, -37, -1000, -1000, 655, 75, -1000, 25,
	-1000, 75, 75, 75, 75, 75, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, 15, 165, 75,
	75, 75, -1000, -1000, 75, 75, 75, 75, 75, 75,
	-1000, 75, -1000, 157, 181, 12, 101, 100, 95, 94,
	-1000, -1000, 163, -1000, -1000, 176, 75, 75, 655, 655,
	655, 655, 204, -1000, 655, 75, 75, 75, 75, 75,
	75, 75, 75, -1000, 75, 75, 75, 75, 25, -1000,
	75, -1000, 75, 75, 75, 75, 75, -1000, -1000, 75,
	89, 167, 85, 25, 25, 25, 25, 75, -1000, -1000,
	69, 25, 655, 655, 655, 655, 655, 655, 655, 160,
	143, 655, 204, 135, 157, 532, 10, 655, 655, 655,
	655, 160, -1000, 75, 75, 43, -1000, 23, -1000, 696,
	285, 25, 25, 25, 25, 25, 25, 25, -1000, -1000,
	-1000, 75, 25, 75, 75, -1000, -1000, 25, -1000, 75,
	202, 25, 25, 25, 25, -1000, -1000, 491, 62, -1000,
	41, 450, 75, -1000, 75, 75, 25, 32, -1000, 117,
	75, -1000, 614, 138, 133, 42, 0, -1000, 75, 72,
	-1000, -1000, 155, -1000, 192, 132, 126, 123, 75, -1000,
	75, -1000, 655, 573, 25, -1000, -1000, -1000, -1000, -1000,
	35, -1000, -1000, 119, -1000, -1000, 573, 409, 75, -1000,
	-1000, 69, -1000, 69, -1000, 106, 655, 655, 573, 25,
	328, -1000, 75, 76, -1000, -1000, -14, 69, -15, 69,
	75, 25, 25, 573, -1000, 174, -1000, -1000, 655, -1000,
	161, 655, 573, -1000, 25, 28, 25, 86, 75, 103,
	655, 75, 74, 655, -1000, 33, -1000,
}
var RubyPgo = []int{

	0, 265, 156, 264, 263, 261, 5, 260, 258, 257,
	256, 254, 253, 252, 251, 250, 249, 248, 6, 246,
	245, 244, 243, 242, 239, 237, 236, 235, 229, 223,
	4, 11, 221, 220, 217, 216, 214, 3, 213, 211,
	205, 202, 200, 0, 7, 1, 195,
}
var RubyR1 = []int{

	0, 35, 35, 35, 35, 3, 3, 3, 30, 30,
	30, 30, 43, 43, 44, 44, 2, 2, 2, 2,
	2, 2, 2, 2, 2, 2, 2, 2, 2, 2,
	2, 2, 2, 2, 2, 2, 2, 2, 2, 2,
	2, 2, 2, 10, 10, 10, 10, 10, 31, 31,
	41, 41, 41, 41, 40, 40, 40, 40, 40, 37,
	37, 37, 37, 37, 45, 45, 45, 14, 34, 34,
	15, 15, 17, 18, 18, 42, 42, 12, 12, 12,
	12, 12, 20, 21, 22, 23, 24, 25, 26, 27,
	28, 29, 4, 7, 8, 5, 5, 36, 36, 36,
	36, 39, 39, 39, 1, 1, 9, 9, 16, 16,
	13, 13, 19, 6, 6, 32, 38, 38, 38, 46,
	46, 11, 11, 33, 33, 33, 33, 33,
}
var RubyR2 = []int{

	0, 0, 1, 2, 3, 1, 1, 3, 0, 2,
	2, 2, 0, 2, 0, 2, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 3, 3, 5, 3, 5, 5, 1,
	1, 1, 5, 5, 1, 1, 5, 5, 5, 0,
	1, 1, 5, 5, 0, 2, 2, 9, 0, 1,
	6, 8, 6, 3, 6, 1, 4, 5, 5, 5,
	5, 5, 3, 3, 3, 3, 5, 5, 5, 5,
	5, 5, 1, 1, 5, 9, 9, 0, 6, 11,
	12, 4, 9, 10, 0, 1, 2, 2, 2, 2,
	3, 3, 1, 3, 7, 3, 0, 1, 5, 1,
	2, 5, 6, 0, 5, 3, 4, 2,
}
var RubyChk = []int{

	-1000, -35, 46, -3, 30, 31, -43, 29, 46, -2,
	5, 6, 7, -16, -13, -9, -10, -14, -15, -17,
	-12, -4, -7, -20, -21, -22, -23, -24, -25, -26,
	-27, -28, -29, -8, -5, -19, -11, 44, 43, 12,
	17, 18, 19, 20, 24, 25, 26, 27, 39, 41,
	45, 14, -43, -43, 33, -43, -43, -43, -43, -43,
	6, 7, 44, 6, 7, -43, -43, -43, -43, -43,
	-43, -43, -43, -43, -43, 4, 26, 27, 28, 35,
	36, 34, 6, -31, -41, 23, 8, -40, -2, -6,
	6, 5, 11, 23, 23, 23, 23, 6, 7, 6,
	-18, -43, -18, -2, -2, -2, -2, -37, 6, 5,
	-44, -2, -43, -43, -43, -43, -43, -43, -43, -43,
	-43, -43, -43, -43, -43, -30, -43, -43, -43, -43,
	-43, -43, 30, 21, 7, -42, 30, -43, 30, -43,
	-30, -2, -2, -2, -2, -2, -2, -2, -31, 6,
	5, 10, -2, -37, 10, 30, 29, -2, 13, -32,
	34, -2, -2, -2, -2, -34, -31, -30, -18, -43,
	32, -30, 10, 40, -36, -39, -2, 6, 13, -33,
	16, 15, -43, -43, -43, -43, -38, 6, -37, -45,
	13, 30, 32, 13, -43, -43, -43, -43, 32, 13,
	16, 15, -43, -30, -2, -6, 9, 5, 6, -6,
	-46, 30, 34, -43, 29, 30, -30, -30, 7, 5,
	6, -44, 10, -44, 10, 23, -43, -43, -30, -2,
	-30, 30, 10, -45, 13, -43, -43, -44, -43, -44,
	22, -2, -2, -30, 13, -43, 13, 42, -43, 42,
	-43, -43, -30, 6, -2, 6, -2, -43, 32, 23,
	-43, 22, -2, -43, 10, -2, 10,
}
var RubyDef = []int{

	1, -2, 2, 3, 5, 6, 0, 12, 4, 12,
	16, -2, -2, -2, -2, -2, 22, 23, 24, 25,
	26, 27, 28, 29, 30, 31, 32, 33, 34, 35,
	36, 37, 38, 39, 40, 41, 42, 0, 0, 12,
	12, 12, 92, 93, 12, 12, 12, 12, 12, 12,
	112, 12, 13, 7, 0, 0, 0, 0, 0, 0,
	108, 109, 0, 106, 107, 0, 12, 12, 0, 0,
	0, 0, 59, 14, 0, 12, 12, 12, 12, 12,
	12, 12, -2, 43, -2, 12, 12, -2, -2, 51,
	-2, 16, -2, 12, 12, 12, 12, 110, 111, 12,
	0, 0, 0, -2, -2, -2, -2, 12, 60, 61,
	12, -2, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 59, 0, 0, 0, 0, 0, 0, 0,
	0, 68, 8, 12, -2, 0, 8, 0, 15, 97,
	0, -2, -2, -2, -2, -2, -2, -2, 45, 54,
	55, 12, -2, 12, 12, 9, 10, -2, 113, 12,
	-2, -2, -2, -2, -2, 64, 69, 0, 0, 73,
	0, 0, 12, 94, 12, 12, 12, -2, 121, 0,
	12, 8, 0, 0, 0, 0, 0, -2, 12, 8,
	70, 8, 0, 72, 0, 14, 14, 0, 12, 122,
	12, 8, 0, 127, -2, 53, 48, 56, 57, 58,
	8, 119, 115, 0, 65, 66, 64, 0, -2, 62,
	63, 12, 14, 12, 14, 0, 0, 0, 125, -2,
	0, 120, 12, 0, 71, 74, 0, 12, 0, 12,
	12, -2, -2, 126, 114, 0, 67, 95, 0, 96,
	0, 0, 124, -2, 12, 0, -2, 0, 12, 0,
	0, 12, -2, 0, 103, -2, 100,
}
var RubyTok1 = []int{

	1,
}
var RubyTok2 = []int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46,
}
var RubyTok3 = []int{
	0,
}

//line yaccpar:1

/*	parser for yacc output	*/

var RubyDebug = 0

type RubyLexer interface {
	Lex(lval *RubySymType) int
	Error(s string)
}

const RubyFlag = -1000

func RubyTokname(c int) string {
	// 4 is TOKSTART above
	if c >= 4 && c-4 < len(RubyToknames) {
		if RubyToknames[c-4] != "" {
			return RubyToknames[c-4]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func RubyStatname(s int) string {
	if s >= 0 && s < len(RubyStatenames) {
		if RubyStatenames[s] != "" {
			return RubyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func Rubylex1(lex RubyLexer, lval *RubySymType) int {
	c := 0
	char := lex.Lex(lval)
	if char <= 0 {
		c = RubyTok1[0]
		goto out
	}
	if char < len(RubyTok1) {
		c = RubyTok1[char]
		goto out
	}
	if char >= RubyPrivate {
		if char < RubyPrivate+len(RubyTok2) {
			c = RubyTok2[char-RubyPrivate]
			goto out
		}
	}
	for i := 0; i < len(RubyTok3); i += 2 {
		c = RubyTok3[i+0]
		if c == char {
			c = RubyTok3[i+1]
			goto out
		}
	}

out:
	if c == 0 {
		c = RubyTok2[1] /* unknown char */
	}
	if RubyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", RubyTokname(c), uint(char))
	}
	return c
}

func RubyParse(Rubylex RubyLexer) int {
	var Rubyn int
	var Rubylval RubySymType
	var RubyVAL RubySymType
	RubyS := make([]RubySymType, RubyMaxDepth)

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	Rubystate := 0
	Rubychar := -1
	Rubyp := -1
	goto Rubystack

ret0:
	return 0

ret1:
	return 1

Rubystack:
	/* put a state and value onto the stack */
	if RubyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", RubyTokname(Rubychar), RubyStatname(Rubystate))
	}

	Rubyp++
	if Rubyp >= len(RubyS) {
		nyys := make([]RubySymType, len(RubyS)*2)
		copy(nyys, RubyS)
		RubyS = nyys
	}
	RubyS[Rubyp] = RubyVAL
	RubyS[Rubyp].yys = Rubystate

Rubynewstate:
	Rubyn = RubyPact[Rubystate]
	if Rubyn <= RubyFlag {
		goto Rubydefault /* simple state */
	}
	if Rubychar < 0 {
		Rubychar = Rubylex1(Rubylex, &Rubylval)
	}
	Rubyn += Rubychar
	if Rubyn < 0 || Rubyn >= RubyLast {
		goto Rubydefault
	}
	Rubyn = RubyAct[Rubyn]
	if RubyChk[Rubyn] == Rubychar { /* valid shift */
		Rubychar = -1
		RubyVAL = Rubylval
		Rubystate = Rubyn
		if Errflag > 0 {
			Errflag--
		}
		goto Rubystack
	}

Rubydefault:
	/* default state action */
	Rubyn = RubyDef[Rubystate]
	if Rubyn == -2 {
		if Rubychar < 0 {
			Rubychar = Rubylex1(Rubylex, &Rubylval)
		}

		/* look through exception table */
		xi := 0
		for {
			if RubyExca[xi+0] == -1 && RubyExca[xi+1] == Rubystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			Rubyn = RubyExca[xi+0]
			if Rubyn < 0 || Rubyn == Rubychar {
				break
			}
		}
		Rubyn = RubyExca[xi+1]
		if Rubyn < 0 {
			goto ret0
		}
	}
	if Rubyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			Rubylex.Error("syntax error")
			Nerrs++
			if RubyDebug >= 1 {
				__yyfmt__.Printf("%s", RubyStatname(Rubystate))
				__yyfmt__.Printf(" saw %s\n", RubyTokname(Rubychar))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for Rubyp >= 0 {
				Rubyn = RubyPact[RubyS[Rubyp].yys] + RubyErrCode
				if Rubyn >= 0 && Rubyn < RubyLast {
					Rubystate = RubyAct[Rubyn] /* simulate a shift of "error" */
					if RubyChk[Rubystate] == RubyErrCode {
						goto Rubystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if RubyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", RubyS[Rubyp].yys)
				}
				Rubyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if RubyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", RubyTokname(Rubychar))
			}
			if Rubychar == RubyEofCode {
				goto ret1
			}
			Rubychar = -1
			goto Rubynewstate /* try again in the same state */
		}
	}

	/* reduction by production Rubyn */
	if RubyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", Rubyn, RubyStatname(Rubystate))
	}

	Rubynt := Rubyn
	Rubypt := Rubyp
	_ = Rubypt // guard against "declared and not used"

	Rubyp -= RubyR2[Rubyn]
	RubyVAL = RubyS[Rubyp+1]

	/* consult goto table to find next state */
	Rubyn = RubyR1[Rubyn]
	Rubyg := RubyPgo[Rubyn]
	Rubyj := Rubyg + RubyS[Rubyp].yys + 1

	if Rubyj >= RubyLast {
		Rubystate = RubyAct[Rubyg]
	} else {
		Rubystate = RubyAct[Rubyj]
		if RubyChk[Rubystate] != -Rubyn {
			Rubystate = RubyAct[Rubyg]
		}
	}
	// dummy call; replaced with literal code
	switch Rubynt {

	case 1:
		//line parser.y:138
		{
			Statements = []ast.Node{}
		}
	case 2:
		//line parser.y:140
		{
			Statements = []ast.Node{}
		}
	case 3:
		//line parser.y:142
		{
			if RubyS[Rubypt-0].genericValue != nil {
				Statements = append(Statements, RubyS[Rubypt-0].genericValue)
			}
		}
	case 4:
		//line parser.y:148
		{
			if RubyS[Rubypt-1].genericValue != nil {
				Statements = append(Statements, RubyS[Rubypt-1].genericValue)
			}
		}
	case 5:
		//line parser.y:154
		{
			RubyVAL.genericValue = nil
		}
	case 6:
		//line parser.y:155
		{
			RubyVAL.genericValue = nil
		}
	case 7:
		//line parser.y:156
		{
			RubyVAL.genericValue = RubyS[Rubypt-1].genericValue
		}
	case 8:
		//line parser.y:159
		{
			RubyVAL.genericSlice = []ast.Node{}
		}
	case 9:
		//line parser.y:161
		{ /* do nothing */
		}
	case 10:
		//line parser.y:163
		{ /* do nothing */
		}
	case 11:
		//line parser.y:165
		{
			if RubyS[Rubypt-0].genericValue != nil {
				RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
			}
		}
	case 16:
		RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
	case 17:
		RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
	case 18:
		RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
	case 19:
		RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
	case 20:
		RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
	case 21:
		RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
	case 22:
		RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
	case 23:
		RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
	case 24:
		RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
	case 25:
		RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
	case 26:
		RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
	case 27:
		RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
	case 28:
		RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
	case 29:
		RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
	case 30:
		RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
	case 31:
		RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
	case 32:
		RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
	case 33:
		RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
	case 34:
		RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
	case 35:
		RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
	case 36:
		RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
	case 37:
		RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
	case 38:
		RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
	case 39:
		RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
	case 40:
		RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
	case 41:
		RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
	case 42:
		RubyVAL.genericValue = RubyS[Rubypt-0].genericValue
	case 43:
		//line parser.y:177
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-2].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-0].genericSlice,
			}
		}
	case 44:
		//line parser.y:184
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-2].genericValue,
				Func:   RubyS[Rubypt-0].genericValue.(ast.BareReference),
			}
		}
	case 45:
		//line parser.y:191
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-4].genericValue,
				Func:   RubyS[Rubypt-2].genericValue.(ast.BareReference),
				Args:   RubyS[Rubypt-0].genericSlice,
			}
		}
	case 46:
		//line parser.y:199
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func: RubyS[Rubypt-2].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-0].genericSlice,
			}
		}
	case 47:
		//line parser.y:206
		{
			RubyVAL.genericValue = ast.CallExpression{
				Func:   ast.BareReference{Name: RubyS[Rubypt-2].operator},
				Target: RubyS[Rubypt-4].genericValue,
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 48:
		//line parser.y:215
		{
			RubyVAL.genericSlice = RubyS[Rubypt-2].genericSlice
		}
	case 49:
		//line parser.y:217
		{
			RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
		}
	case 50:
		//line parser.y:220
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 51:
		//line parser.y:222
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 52:
		//line parser.y:224
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 53:
		//line parser.y:226
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 54:
		//line parser.y:229
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 55:
		//line parser.y:231
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 56:
		//line parser.y:233
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 57:
		//line parser.y:235
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 58:
		//line parser.y:237
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 59:
		//line parser.y:239
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 60:
		//line parser.y:241
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 61:
		//line parser.y:243
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 62:
		//line parser.y:245
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 63:
		//line parser.y:247
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 67:
		//line parser.y:257
		{
			RubyVAL.genericValue = ast.FuncDecl{
				Name: RubyS[Rubypt-6].genericValue.(ast.BareReference),
				Args: RubyS[Rubypt-4].genericSlice,
				Body: RubyS[Rubypt-2].genericSlice,
			}
		}
	case 68:
		//line parser.y:265
		{
			RubyVAL.genericSlice = []ast.Node{}
		}
	case 69:
		//line parser.y:266
		{
			RubyVAL.genericSlice = RubyS[Rubypt-0].genericSlice
		}
	case 70:
		//line parser.y:269
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name: RubyS[Rubypt-3].genericValue.(ast.Class).Name,
				Body: RubyS[Rubypt-1].genericSlice,
			}
		}
	case 71:
		//line parser.y:276
		{
			RubyVAL.genericValue = ast.ClassDecl{
				Name:       RubyS[Rubypt-5].genericValue.(ast.Class).Name,
				SuperClass: RubyS[Rubypt-3].genericValue.(ast.Class),
				Namespace:  RubyS[Rubypt-5].genericValue.(ast.Class).Namespace,
				Body:       RubyS[Rubypt-1].genericSlice,
			}
		}
	case 72:
		//line parser.y:286
		{
			RubyVAL.genericValue = ast.ModuleDecl{
				Name:      RubyS[Rubypt-3].genericValue.(ast.Class).Name,
				Namespace: RubyS[Rubypt-3].genericValue.(ast.Class).Namespace,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 73:
		//line parser.y:295
		{
			RubyVAL.genericValue = ast.Class{
				Name: RubyS[Rubypt-1].genericValue.(ast.BareReference).Name,
			}
		}
	case 74:
		//line parser.y:301
		{
			RubyVAL.genericValue = ast.Class{
				Name:      RubyS[Rubypt-1].genericValue.(ast.BareReference).Name,
				Namespace: strings.Join(RubyS[Rubypt-4].stringSlice, "::"),
			}
		}
	case 75:
		//line parser.y:309
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 76:
		//line parser.y:313
		{
			RubyVAL.stringSlice = append(RubyVAL.stringSlice, RubyS[Rubypt-0].genericValue.(ast.BareReference).Name)
		}
	case 77:
		//line parser.y:318
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-4].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 78:
		//line parser.y:325
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-4].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 79:
		//line parser.y:332
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-4].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 80:
		//line parser.y:339
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-4].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 81:
		//line parser.y:346
		{
			RubyVAL.genericValue = ast.Assignment{
				LHS: RubyS[Rubypt-4].genericValue,
				RHS: RubyS[Rubypt-0].genericValue,
			}
		}
	case 82:
		//line parser.y:353
		{
			RubyVAL.genericValue = ast.Negation{Target: RubyS[Rubypt-0].genericValue}
		}
	case 83:
		//line parser.y:354
		{
			RubyVAL.genericValue = ast.Complement{Target: RubyS[Rubypt-0].genericValue}
		}
	case 84:
		//line parser.y:355
		{
			RubyVAL.genericValue = ast.Positive{Target: RubyS[Rubypt-0].genericValue}
		}
	case 85:
		//line parser.y:356
		{
			RubyVAL.genericValue = ast.Negative{Target: RubyS[Rubypt-0].genericValue}
		}
	case 86:
		//line parser.y:359
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-4].genericValue,
				Func:   ast.BareReference{Name: "+"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 87:
		//line parser.y:368
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-4].genericValue,
				Func:   ast.BareReference{Name: "-"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 88:
		//line parser.y:377
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-4].genericValue,
				Func:   ast.BareReference{Name: "*"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 89:
		//line parser.y:386
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-4].genericValue,
				Func:   ast.BareReference{Name: "/"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 90:
		//line parser.y:395
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-4].genericValue,
				Func:   ast.BareReference{Name: "&"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 91:
		//line parser.y:404
		{
			RubyVAL.genericValue = ast.CallExpression{
				Target: RubyS[Rubypt-4].genericValue,
				Func:   ast.BareReference{Name: "|"},
				Args:   []ast.Node{RubyS[Rubypt-0].genericValue},
			}
		}
	case 92:
		//line parser.y:412
		{
			RubyVAL.genericValue = ast.Boolean{Value: true}
		}
	case 93:
		//line parser.y:413
		{
			RubyVAL.genericValue = ast.Boolean{Value: false}
		}
	case 94:
		//line parser.y:415
		{
			RubyVAL.genericValue = ast.Array{Nodes: RubyS[Rubypt-2].genericSlice}
		}
	case 95:
		//line parser.y:419
		{
			pairs := []ast.HashKeyValuePair{}
			for _, node := range RubyS[Rubypt-4].genericSlice {
				pairs = append(pairs, node.(ast.HashKeyValuePair))
			}
			RubyVAL.genericValue = ast.Hash{Pairs: pairs}
		}
	case 96:
		//line parser.y:427
		{
			pairs := []ast.HashKeyValuePair{}
			for _, node := range RubyS[Rubypt-4].genericSlice {
				pairs = append(pairs, node.(ast.HashKeyValuePair))
			}
			RubyVAL.genericValue = ast.Hash{Pairs: pairs}
		}
	case 97:
		//line parser.y:435
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 98:
		//line parser.y:437
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-5].genericValue, Value: RubyS[Rubypt-0].genericValue})
		}
	case 99:
		//line parser.y:439
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-5].genericValue, Value: RubyS[Rubypt-0].genericValue})
		}
	case 100:
		//line parser.y:441
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{Key: RubyS[Rubypt-6].genericValue, Value: RubyS[Rubypt-1].genericValue})
		}
	case 101:
		//line parser.y:444
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-3].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-0].genericValue,
			})
		}
	case 102:
		//line parser.y:451
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-3].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-0].genericValue,
			})
		}
	case 103:
		//line parser.y:458
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.HashKeyValuePair{
				Key:   ast.Symbol{Name: RubyS[Rubypt-4].genericValue.(ast.BareReference).Name},
				Value: RubyS[Rubypt-1].genericValue,
			})
		}
	case 104:
		//line parser.y:466
		{
			RubyVAL.genericValue = nil
		}
	case 105:
		//line parser.y:467
		{
			RubyVAL.genericValue = nil
		}
	case 106:
		//line parser.y:470
		{
			RubyVAL.genericValue = ast.GlobalVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 107:
		//line parser.y:472
		{
			RubyVAL.genericValue = ast.GlobalVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 108:
		//line parser.y:475
		{
			RubyVAL.genericValue = ast.InstanceVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 109:
		//line parser.y:477
		{
			RubyVAL.genericValue = ast.InstanceVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 110:
		//line parser.y:480
		{
			RubyVAL.genericValue = ast.ClassVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 111:
		//line parser.y:482
		{
			RubyVAL.genericValue = ast.ClassVariable{Name: RubyS[Rubypt-0].genericValue.(ast.BareReference).Name}
		}
	case 112:
		//line parser.y:485
		{
			RubyVAL.genericValue = ast.FileNameConstReference{}
		}
	case 113:
		//line parser.y:488
		{
			RubyVAL.genericValue = ast.Block{Body: RubyS[Rubypt-1].genericSlice}
		}
	case 114:
		//line parser.y:490
		{
			RubyVAL.genericValue = ast.Block{
				Body: RubyS[Rubypt-1].genericSlice,
				Args: RubyS[Rubypt-4].genericSlice,
			}
		}
	case 115:
		//line parser.y:498
		{
			RubyVAL.genericSlice = RubyS[Rubypt-1].genericSlice
		}
	case 116:
		//line parser.y:500
		{
			RubyVAL.genericSlice = ast.Nodes{}
		}
	case 117:
		//line parser.y:502
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 118:
		//line parser.y:504
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, RubyS[Rubypt-0].genericValue)
		}
	case 121:
		//line parser.y:509
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-2].genericValue,
				Body:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 122:
		//line parser.y:516
		{
			RubyVAL.genericValue = ast.IfBlock{
				Condition: RubyS[Rubypt-3].genericValue,
				Body:      RubyS[Rubypt-2].genericSlice,
				Else:      RubyS[Rubypt-1].genericSlice,
			}
		}
	case 123:
		//line parser.y:524
		{
			RubyVAL.genericSlice = []ast.Node{}
		}
	case 124:
		//line parser.y:526
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 125:
		//line parser.y:533
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 126:
		//line parser.y:540
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: RubyS[Rubypt-1].genericValue,
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	case 127:
		//line parser.y:547
		{
			RubyVAL.genericSlice = append(RubyVAL.genericSlice, ast.IfBlock{
				Condition: ast.Boolean{Value: true},
				Body:      RubyS[Rubypt-0].genericSlice,
			})
		}
	}
	goto Rubystack /* stack new state and value */
}
