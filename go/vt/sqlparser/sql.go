//line sql.y:6
package sqlparser

import __yyfmt__ "fmt"

//line sql.y:6
import "bytes"

func SetParseTree(yylex interface{}, stmt Statement) {
	yylex.(*Tokenizer).ParseTree = stmt
}

func SetAllowComments(yylex interface{}, allow bool) {
	yylex.(*Tokenizer).AllowComments = allow
}

func ForceEOF(yylex interface{}) {
	yylex.(*Tokenizer).ForceEOF = true
}

var (
	SHARE        = []byte("share")
	MODE         = []byte("mode")
	IF_BYTES     = []byte("if")
	VALUES_BYTES = []byte("values")
)

//line sql.y:31
type yySymType struct {
	yys         int
	empty       struct{}
	statement   Statement
	selStmt     SelectStatement
	byt         byte
	bytes       []byte
	bytes2      [][]byte
	str         string
	selectExprs SelectExprs
	selectExpr  SelectExpr
	columns     Columns
	colName     *ColName
	tableExprs  TableExprs
	tableExpr   TableExpr
	smTableExpr SimpleTableExpr
	tableName   *TableName
	indexHints  *IndexHints
	expr        Expr
	boolExpr    BoolExpr
	valExpr     ValExpr
	colTuple    ColTuple
	valExprs    ValExprs
	values      Values
	rowTuple    RowTuple
	subquery    *Subquery
	caseExpr    *CaseExpr
	whens       []*When
	when        *When
	orderBy     OrderBy
	order       *Order
	limit       *Limit
	insRows     InsertRows
	updateExprs UpdateExprs
	updateExpr  *UpdateExpr
}

const LEX_ERROR = 57346
const SELECT = 57347
const INSERT = 57348
const UPDATE = 57349
const DELETE = 57350
const FROM = 57351
const WHERE = 57352
const GROUP = 57353
const HAVING = 57354
const ORDER = 57355
const BY = 57356
const LIMIT = 57357
const FOR = 57358
const ALL = 57359
const DISTINCT = 57360
const AS = 57361
const EXISTS = 57362
const IN = 57363
const IS = 57364
const LIKE = 57365
const BETWEEN = 57366
const NULL = 57367
const TRUE = 57368
const FALSE = 57369
const ASC = 57370
const DESC = 57371
const VALUES = 57372
const INTO = 57373
const DUPLICATE = 57374
const KEY = 57375
const DEFAULT = 57376
const SET = 57377
const LOCK = 57378
const KEYRANGE = 57379
const ID = 57380
const STRING = 57381
const NUMBER = 57382
const VALUE_ARG = 57383
const LIST_ARG = 57384
const COMMENT = 57385
const LE = 57386
const GE = 57387
const NE = 57388
const NULL_SAFE_EQUAL = 57389
const UNION = 57390
const MINUS = 57391
const EXCEPT = 57392
const INTERSECT = 57393
const JOIN = 57394
const STRAIGHT_JOIN = 57395
const LEFT = 57396
const RIGHT = 57397
const INNER = 57398
const OUTER = 57399
const CROSS = 57400
const NATURAL = 57401
const USE = 57402
const FORCE = 57403
const ON = 57404
const OR = 57405
const AND = 57406
const NOT = 57407
const UNARY = 57408
const CASE = 57409
const WHEN = 57410
const THEN = 57411
const ELSE = 57412
const END = 57413
const CREATE = 57414
const ALTER = 57415
const DROP = 57416
const RENAME = 57417
const ANALYZE = 57418
const TABLE = 57419
const INDEX = 57420
const VIEW = 57421
const TO = 57422
const IGNORE = 57423
const IF = 57424
const UNIQUE = 57425
const USING = 57426
const SHOW = 57427
const DESCRIBE = 57428
const EXPLAIN = 57429

var yyToknames = []string{
	"LEX_ERROR",
	"SELECT",
	"INSERT",
	"UPDATE",
	"DELETE",
	"FROM",
	"WHERE",
	"GROUP",
	"HAVING",
	"ORDER",
	"BY",
	"LIMIT",
	"FOR",
	"ALL",
	"DISTINCT",
	"AS",
	"EXISTS",
	"IN",
	"IS",
	"LIKE",
	"BETWEEN",
	"NULL",
	"TRUE",
	"FALSE",
	"ASC",
	"DESC",
	"VALUES",
	"INTO",
	"DUPLICATE",
	"KEY",
	"DEFAULT",
	"SET",
	"LOCK",
	"KEYRANGE",
	"ID",
	"STRING",
	"NUMBER",
	"VALUE_ARG",
	"LIST_ARG",
	"COMMENT",
	"LE",
	"GE",
	"NE",
	"NULL_SAFE_EQUAL",
	" (",
	" =",
	" <",
	" >",
	" ~",
	"UNION",
	"MINUS",
	"EXCEPT",
	"INTERSECT",
	" ,",
	"JOIN",
	"STRAIGHT_JOIN",
	"LEFT",
	"RIGHT",
	"INNER",
	"OUTER",
	"CROSS",
	"NATURAL",
	"USE",
	"FORCE",
	"ON",
	"OR",
	"AND",
	"NOT",
	" &",
	" |",
	" ^",
	" +",
	" -",
	" *",
	" /",
	" %",
	" .",
	"UNARY",
	"CASE",
	"WHEN",
	"THEN",
	"ELSE",
	"END",
	"CREATE",
	"ALTER",
	"DROP",
	"RENAME",
	"ANALYZE",
	"TABLE",
	"INDEX",
	"VIEW",
	"TO",
	"IGNORE",
	"IF",
	"UNIQUE",
	"USING",
	"SHOW",
	"DESCRIBE",
	"EXPLAIN",
}
var yyStatenames = []string{}

const yyEofCode = 1
const yyErrCode = 2
const yyMaxDepth = 200

//line yacctab:1
var yyExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyNprod = 209
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 587

var yyAct = []int{

	94, 300, 161, 370, 91, 85, 337, 252, 62, 164,
	92, 200, 292, 243, 80, 211, 63, 180, 90, 163,
	3, 138, 137, 81, 350, 188, 50, 105, 76, 68,
	379, 379, 259, 263, 264, 265, 266, 267, 65, 268,
	269, 70, 64, 38, 73, 40, 348, 53, 77, 41,
	126, 347, 51, 52, 97, 232, 379, 132, 86, 102,
	103, 104, 234, 43, 110, 44, 298, 132, 317, 319,
	122, 98, 84, 99, 100, 101, 381, 380, 321, 130,
	132, 234, 89, 346, 134, 69, 108, 72, 28, 29,
	30, 31, 165, 49, 160, 162, 166, 123, 318, 45,
	125, 328, 378, 326, 244, 88, 290, 274, 323, 106,
	107, 82, 297, 286, 174, 65, 111, 136, 65, 64,
	184, 183, 64, 178, 244, 170, 284, 235, 46, 47,
	48, 109, 119, 115, 182, 86, 206, 184, 233, 150,
	151, 152, 210, 208, 209, 218, 219, 185, 224, 225,
	226, 227, 228, 229, 230, 231, 205, 198, 121, 204,
	357, 358, 137, 71, 138, 137, 138, 137, 213, 117,
	236, 86, 86, 343, 293, 255, 194, 65, 65, 330,
	293, 64, 250, 238, 240, 248, 129, 254, 345, 256,
	117, 241, 223, 220, 222, 247, 251, 192, 311, 344,
	315, 195, 207, 312, 145, 146, 147, 148, 149, 150,
	151, 152, 314, 309, 273, 236, 257, 260, 310, 277,
	278, 181, 275, 148, 149, 150, 151, 152, 313, 181,
	204, 276, 59, 28, 29, 30, 31, 283, 221, 131,
	234, 355, 86, 213, 332, 14, 287, 118, 203, 214,
	291, 191, 193, 190, 285, 212, 289, 295, 202, 299,
	97, 296, 176, 75, 14, 102, 103, 104, 261, 365,
	110, 364, 363, 307, 308, 177, 117, 98, 66, 99,
	100, 101, 325, 167, 172, 171, 169, 132, 89, 204,
	204, 113, 108, 329, 116, 168, 112, 203, 272, 65,
	135, 334, 71, 333, 335, 338, 327, 202, 66, 322,
	320, 88, 78, 304, 303, 106, 107, 271, 197, 71,
	196, 179, 111, 14, 15, 16, 17, 349, 127, 124,
	339, 120, 60, 351, 79, 74, 114, 109, 352, 376,
	102, 103, 104, 353, 331, 236, 383, 360, 359, 362,
	58, 14, 361, 18, 99, 100, 101, 367, 338, 377,
	186, 369, 368, 128, 371, 371, 371, 65, 372, 373,
	239, 64, 97, 56, 54, 374, 246, 102, 103, 104,
	384, 301, 110, 342, 385, 215, 386, 216, 217, 98,
	84, 99, 100, 101, 302, 354, 282, 280, 281, 253,
	89, 341, 306, 181, 108, 19, 20, 22, 21, 23,
	145, 146, 147, 148, 149, 150, 151, 152, 24, 25,
	26, 61, 382, 88, 366, 97, 14, 106, 107, 82,
	102, 103, 104, 33, 111, 110, 187, 39, 258, 14,
	189, 42, 98, 66, 99, 100, 101, 67, 249, 109,
	32, 175, 375, 89, 356, 237, 336, 108, 340, 102,
	103, 104, 305, 288, 110, 173, 34, 35, 36, 37,
	242, 96, 66, 99, 100, 101, 88, 93, 95, 294,
	106, 107, 167, 245, 139, 87, 108, 111, 316, 201,
	102, 103, 104, 262, 199, 110, 83, 270, 140, 144,
	142, 143, 109, 66, 99, 100, 101, 133, 55, 106,
	107, 27, 57, 167, 13, 12, 111, 108, 11, 10,
	9, 156, 157, 158, 159, 8, 153, 154, 155, 7,
	324, 109, 145, 146, 147, 148, 149, 150, 151, 152,
	106, 107, 6, 5, 4, 2, 1, 111, 141, 145,
	146, 147, 148, 149, 150, 151, 152, 0, 0, 0,
	0, 279, 109, 145, 146, 147, 148, 149, 150, 151,
	152, 145, 146, 147, 148, 149, 150, 151, 152, 263,
	264, 265, 266, 267, 0, 268, 269,
}
var yyPact = []int{

	318, -1000, -1000, 180, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -49,
	-31, 7, 36, 1, -1000, -1000, -1000, 421, 357, -1000,
	-1000, -1000, 355, -1000, 319, 294, 412, 270, -68, -8,
	264, -1000, -5, 264, -1000, 297, -69, 264, -69, 296,
	-1000, -1000, -1000, -1000, -1000, 34, -1000, 253, 294, 301,
	53, 294, 133, -1000, 198, -1000, 52, 293, 87, 264,
	-1000, -1000, 291, -1000, -45, 290, 343, 118, 264, -1000,
	230, -1000, -1000, 281, 37, 97, 477, -1000, 405, 240,
	-1000, -1000, -1000, 465, 247, 238, -1000, 237, 236, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, 465, -1000, 227, 270, 283, 393, 270, 465, 264,
	-1000, 340, -74, -1000, 163, -1000, 282, -1000, -1000, 280,
	-1000, 210, 34, -1000, -1000, 264, 125, 405, 405, 465,
	207, 364, 465, 465, 167, 465, 465, 465, 465, 465,
	465, 465, 465, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, 477, -48, 35, 24, 477, -1000, 434, 352, 34,
	-1000, 421, 315, 41, 499, 346, 270, 270, 219, -1000,
	386, 405, -1000, 499, -1000, -1000, -1000, 107, 264, -1000,
	-63, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 211,
	521, 279, 259, 27, -1000, -1000, -1000, -1000, -1000, 92,
	499, -1000, 434, -1000, -1000, 207, 465, 465, 499, 491,
	-1000, 371, -1000, -1000, 148, 148, 148, 62, 62, -1000,
	-1000, -1000, -1000, -1000, 465, -1000, 499, -1000, 23, 34,
	10, 189, 21, -1000, 405, 106, 235, 180, 112, 9,
	-1000, 386, 366, 380, 97, 276, -1000, -1000, 275, -1000,
	391, 210, 210, -1000, -1000, 155, 140, 170, 154, 142,
	2, -1000, 272, -25, 271, 5, -1000, 499, 460, 465,
	-1000, -1000, -1000, 499, -1000, 0, -1000, 315, 15, -1000,
	465, 95, -1000, 312, 187, -1000, -1000, -1000, 270, 366,
	-1000, 465, 465, -1000, -1000, 389, 369, 521, 105, -1000,
	141, -1000, 130, -1000, -1000, -1000, -1000, -10, -42, -47,
	-1000, -1000, -1000, -1000, 465, 499, -1000, -79, -1000, 499,
	465, 305, 235, -1000, -1000, 338, 184, -1000, 132, -1000,
	386, 405, 465, 405, -1000, -1000, 224, 223, 221, 499,
	-1000, 499, 417, -1000, 465, 465, -1000, -1000, -1000, 366,
	97, 183, 97, 264, 264, 264, 270, 499, -1000, 323,
	-1, -1000, -26, -27, 133, -1000, 415, 325, -1000, 264,
	-1000, -1000, -1000, 264, -1000, 264, -1000,
}
var yyPgo = []int{

	0, 546, 545, 19, 544, 543, 542, 529, 525, 520,
	519, 518, 515, 514, 450, 512, 511, 508, 14, 23,
	507, 497, 496, 494, 11, 493, 489, 232, 488, 3,
	17, 5, 485, 484, 483, 18, 2, 15, 9, 479,
	10, 478, 27, 477, 4, 471, 470, 13, 465, 463,
	462, 458, 7, 456, 6, 454, 1, 452, 451, 448,
	12, 8, 16, 263, 447, 441, 440, 438, 437, 436,
	0, 26, 433,
}
var yyR1 = []int{

	0, 1, 2, 2, 2, 2, 2, 2, 2, 2,
	2, 2, 2, 3, 3, 4, 4, 5, 6, 7,
	8, 8, 8, 9, 9, 9, 10, 11, 11, 11,
	12, 13, 13, 13, 72, 14, 15, 15, 16, 16,
	16, 16, 16, 17, 17, 18, 18, 19, 19, 19,
	22, 22, 20, 20, 20, 23, 23, 24, 24, 24,
	24, 21, 21, 21, 25, 25, 25, 25, 25, 25,
	25, 25, 25, 26, 26, 26, 27, 27, 28, 28,
	28, 28, 29, 29, 30, 30, 31, 31, 31, 31,
	31, 32, 32, 32, 32, 32, 32, 32, 32, 32,
	32, 32, 32, 32, 32, 32, 33, 33, 33, 33,
	33, 33, 33, 37, 37, 37, 42, 38, 38, 36,
	36, 36, 36, 36, 36, 36, 36, 36, 36, 36,
	36, 36, 36, 36, 36, 36, 41, 41, 43, 43,
	43, 45, 48, 48, 46, 46, 47, 49, 49, 44,
	44, 35, 35, 35, 35, 35, 35, 50, 50, 51,
	51, 52, 52, 53, 53, 54, 55, 55, 55, 56,
	56, 56, 57, 57, 57, 58, 58, 59, 59, 60,
	60, 34, 34, 39, 39, 40, 40, 61, 61, 62,
	63, 63, 64, 64, 65, 65, 66, 66, 66, 66,
	66, 67, 67, 68, 68, 69, 69, 70, 71,
}
var yyR2 = []int{

	0, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 12, 3, 7, 7, 8, 7, 3,
	5, 8, 4, 6, 7, 4, 5, 4, 5, 5,
	3, 2, 2, 2, 0, 2, 0, 2, 1, 2,
	1, 1, 1, 0, 1, 1, 3, 1, 2, 3,
	1, 1, 0, 1, 2, 1, 3, 3, 3, 3,
	5, 0, 1, 2, 1, 1, 2, 3, 2, 3,
	2, 2, 2, 1, 3, 1, 1, 3, 0, 5,
	5, 5, 1, 3, 0, 2, 1, 3, 3, 2,
	3, 3, 3, 4, 3, 4, 5, 6, 3, 4,
	3, 4, 3, 4, 2, 6, 1, 1, 1, 1,
	1, 1, 1, 3, 1, 1, 3, 1, 3, 1,
	1, 1, 3, 3, 3, 3, 3, 3, 3, 3,
	2, 3, 4, 5, 4, 1, 1, 1, 1, 1,
	1, 5, 0, 1, 1, 2, 4, 0, 2, 1,
	3, 1, 1, 1, 1, 1, 1, 0, 3, 0,
	2, 0, 3, 1, 3, 2, 0, 1, 1, 0,
	2, 4, 0, 2, 4, 0, 3, 1, 3, 0,
	5, 2, 1, 1, 3, 3, 1, 1, 3, 3,
	0, 2, 0, 3, 0, 1, 1, 1, 1, 1,
	1, 0, 1, 0, 1, 0, 2, 1, 0,
}
var yyChk = []int{

	-1000, -1, -2, -3, -4, -5, -6, -7, -8, -9,
	-10, -11, -12, -13, 5, 6, 7, 8, 35, 87,
	88, 90, 89, 91, 100, 101, 102, -16, 53, 54,
	55, 56, -14, -72, -14, -14, -14, -14, 92, -68,
	94, 98, -65, 94, 96, 92, 92, 93, 94, 92,
	-71, -71, -71, -3, 17, -17, 18, -15, 31, -27,
	38, 9, -61, -62, -44, -70, 38, -64, 97, 93,
	-70, 38, 92, -70, 38, -63, 97, -70, -63, 38,
	-18, -19, 77, -22, 38, -31, -36, -32, 71, 48,
	-35, -44, -40, -43, -70, -41, -45, 20, 37, 39,
	40, 41, 25, 26, 27, -42, 75, 76, 52, 97,
	30, 82, 43, -27, 35, 80, -27, 57, 49, 80,
	38, 71, -70, -71, 38, -71, 95, 38, 20, 68,
	-70, 9, 57, -20, -70, 19, 80, 70, 69, -33,
	21, 71, 23, 24, 22, 72, 73, 74, 75, 76,
	77, 78, 79, 49, 50, 51, 44, 45, 46, 47,
	-31, -36, -31, -3, -38, -36, -36, 48, 48, 48,
	-42, 48, 48, -48, -36, -58, 35, 48, -61, 38,
	-30, 10, -62, -36, -70, -71, 20, -69, 99, -66,
	90, 88, 34, 89, 13, 38, 38, 38, -71, -23,
	-24, -26, 48, 38, -42, -19, -70, 77, -31, -31,
	-36, -37, 48, -42, 42, 21, 23, 24, -36, -36,
	26, 71, 27, 25, -36, -36, -36, -36, -36, -36,
	-36, -36, 103, 103, 57, 103, -36, 103, -18, 18,
	-18, -35, -46, -47, 83, -34, 30, -3, -61, -59,
	-44, -30, -52, 13, -31, 68, -70, -71, -67, 95,
	-30, 57, -25, 58, 59, 60, 61, 62, 64, 65,
	-21, 38, 19, -24, 80, -38, -37, -36, -36, 70,
	26, 27, 25, -36, 103, -18, 103, 57, -49, -47,
	85, -31, -60, 68, -39, -40, -60, 103, 57, -52,
	-56, 15, 14, 38, 38, -50, 11, -24, -24, 58,
	63, 58, 63, 58, 58, 58, -28, 66, 96, 67,
	38, 103, 38, 103, 70, -36, 103, -35, 86, -36,
	84, 32, 57, -44, -56, -36, -53, -54, -36, -71,
	-51, 12, 14, 68, 58, 58, 93, 93, 93, -36,
	103, -36, 33, -40, 57, 57, -55, 28, 29, -52,
	-31, -38, -31, 48, 48, 48, 7, -36, -54, -56,
	-29, -70, -29, -29, -61, -57, 16, 36, 103, 57,
	103, 103, 7, 21, -70, -70, -70,
}
var yyDef = []int{

	0, -2, 1, 2, 3, 4, 5, 6, 7, 8,
	9, 10, 11, 12, 34, 34, 34, 34, 34, 203,
	194, 0, 0, 0, 208, 208, 208, 0, 38, 40,
	41, 42, 43, 36, 0, 0, 0, 0, 192, 0,
	0, 204, 0, 0, 195, 0, 190, 0, 190, 0,
	31, 32, 33, 14, 39, 0, 44, 35, 0, 0,
	76, 0, 19, 187, 0, 149, 207, 0, 0, 0,
	208, 207, 0, 208, 0, 0, 0, 0, 0, 30,
	0, 45, 47, 52, 207, 50, 51, 86, 0, 0,
	119, 120, 121, 0, 149, 0, 135, 0, 0, 151,
	152, 153, 154, 155, 156, 186, 138, 139, 140, 136,
	137, 142, 37, 175, 0, 0, 84, 0, 0, 0,
	208, 0, 205, 22, 0, 25, 0, 27, 191, 0,
	208, 0, 0, 48, 53, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 106, 107, 108, 109, 110, 111, 112,
	89, 0, 0, 0, 0, 117, 130, 0, 0, 0,
	104, 0, 0, 0, 143, 0, 0, 0, 84, 77,
	161, 0, 188, 189, 150, 20, 193, 0, 0, 208,
	201, 196, 197, 198, 199, 200, 26, 28, 29, 84,
	55, 61, 0, 73, 75, 46, 54, 49, 87, 88,
	91, 92, 0, 114, 115, 0, 0, 0, 94, 0,
	98, 0, 100, 102, 122, 123, 124, 125, 126, 127,
	128, 129, 90, 116, 0, 185, 117, 131, 0, 0,
	0, 0, 147, 144, 0, 179, 0, 182, 179, 0,
	177, 161, 169, 0, 85, 0, 206, 23, 0, 202,
	157, 0, 0, 64, 65, 0, 0, 0, 0, 0,
	78, 62, 0, 0, 0, 0, 93, 95, 0, 0,
	99, 101, 103, 118, 132, 0, 134, 0, 0, 145,
	0, 0, 15, 0, 181, 183, 16, 176, 0, 169,
	18, 0, 0, 208, 24, 159, 0, 56, 59, 66,
	0, 68, 0, 70, 71, 72, 57, 0, 0, 0,
	63, 58, 74, 113, 0, 96, 133, 0, 141, 148,
	0, 0, 0, 178, 17, 170, 162, 163, 166, 21,
	161, 0, 0, 0, 67, 69, 0, 0, 0, 97,
	105, 146, 0, 184, 0, 0, 165, 167, 168, 169,
	160, 158, 60, 0, 0, 0, 0, 171, 164, 172,
	0, 82, 0, 0, 180, 13, 0, 0, 79, 0,
	80, 81, 173, 0, 83, 0, 174,
}
var yyTok1 = []int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 79, 72, 3,
	48, 103, 77, 75, 57, 76, 80, 78, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	50, 49, 51, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 74, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 73, 3, 52,
}
var yyTok2 = []int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46, 47, 53, 54, 55, 56,
	58, 59, 60, 61, 62, 63, 64, 65, 66, 67,
	68, 69, 70, 71, 81, 82, 83, 84, 85, 86,
	87, 88, 89, 90, 91, 92, 93, 94, 95, 96,
	97, 98, 99, 100, 101, 102,
}
var yyTok3 = []int{
	0,
}

//line yaccpar:1

/*	parser for yacc output	*/

var yyDebug = 0

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

const yyFlag = -1000

func yyTokname(c int) string {
	// 4 is TOKSTART above
	if c >= 4 && c-4 < len(yyToknames) {
		if yyToknames[c-4] != "" {
			return yyToknames[c-4]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yylex1(lex yyLexer, lval *yySymType) int {
	c := 0
	char := lex.Lex(lval)
	if char <= 0 {
		c = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		c = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			c = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		c = yyTok3[i+0]
		if c == char {
			c = yyTok3[i+1]
			goto out
		}
	}

out:
	if c == 0 {
		c = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(c), uint(char))
	}
	return c
}

func yyParse(yylex yyLexer) int {
	var yyn int
	var yylval yySymType
	var yyVAL yySymType
	yyS := make([]yySymType, yyMaxDepth)

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yychar := -1
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yychar), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = yyPact[yystate]
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yychar < 0 {
		yychar = yylex1(yylex, &yylval)
	}
	yyn += yychar
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yychar { /* valid shift */
		yychar = -1
		yyVAL = yylval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = yyDef[yystate]
	if yyn == -2 {
		if yychar < 0 {
			yychar = yylex1(yylex, &yylval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && yyExca[xi+1] == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = yyExca[xi+0]
			if yyn < 0 || yyn == yychar {
				break
			}
		}
		yyn = yyExca[xi+1]
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error("syntax error")
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yychar))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = yyPact[yyS[yyp].yys] + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = yyAct[yyn] /* simulate a shift of "error" */
					if yyChk[yystate] == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yychar))
			}
			if yychar == yyEofCode {
				goto ret1
			}
			yychar = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= yyR2[yyn]
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = yyR1[yyn]
	yyg := yyPgo[yyn]
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = yyAct[yyg]
	} else {
		yystate = yyAct[yyj]
		if yyChk[yystate] != -yyn {
			yystate = yyAct[yyg]
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 1:
		//line sql.y:151
		{
			SetParseTree(yylex, yyS[yypt-0].statement)
		}
	case 2:
		//line sql.y:157
		{
			yyVAL.statement = yyS[yypt-0].selStmt
		}
	case 3:
		yyVAL.statement = yyS[yypt-0].statement
	case 4:
		yyVAL.statement = yyS[yypt-0].statement
	case 5:
		yyVAL.statement = yyS[yypt-0].statement
	case 6:
		yyVAL.statement = yyS[yypt-0].statement
	case 7:
		yyVAL.statement = yyS[yypt-0].statement
	case 8:
		yyVAL.statement = yyS[yypt-0].statement
	case 9:
		yyVAL.statement = yyS[yypt-0].statement
	case 10:
		yyVAL.statement = yyS[yypt-0].statement
	case 11:
		yyVAL.statement = yyS[yypt-0].statement
	case 12:
		yyVAL.statement = yyS[yypt-0].statement
	case 13:
		//line sql.y:173
		{
			yyVAL.selStmt = &Select{Comments: Comments(yyS[yypt-10].bytes2), Distinct: yyS[yypt-9].str, SelectExprs: yyS[yypt-8].selectExprs, From: yyS[yypt-6].tableExprs, Where: NewWhere(AST_WHERE, yyS[yypt-5].boolExpr), GroupBy: GroupBy(yyS[yypt-4].valExprs), Having: NewWhere(AST_HAVING, yyS[yypt-3].boolExpr), OrderBy: yyS[yypt-2].orderBy, Limit: yyS[yypt-1].limit, Lock: yyS[yypt-0].str}
		}
	case 14:
		//line sql.y:177
		{
			yyVAL.selStmt = &Union{Type: yyS[yypt-1].str, Left: yyS[yypt-2].selStmt, Right: yyS[yypt-0].selStmt}
		}
	case 15:
		//line sql.y:183
		{
			yyVAL.statement = &Insert{Comments: Comments(yyS[yypt-5].bytes2), Table: yyS[yypt-3].tableName, Columns: yyS[yypt-2].columns, Rows: yyS[yypt-1].insRows, OnDup: OnDup(yyS[yypt-0].updateExprs)}
		}
	case 16:
		//line sql.y:187
		{
			cols := make(Columns, 0, len(yyS[yypt-1].updateExprs))
			vals := make(ValTuple, 0, len(yyS[yypt-1].updateExprs))
			for _, col := range yyS[yypt-1].updateExprs {
				cols = append(cols, &NonStarExpr{Expr: col.Name})
				vals = append(vals, col.Expr)
			}
			yyVAL.statement = &Insert{Comments: Comments(yyS[yypt-5].bytes2), Table: yyS[yypt-3].tableName, Columns: cols, Rows: Values{vals}, OnDup: OnDup(yyS[yypt-0].updateExprs)}
		}
	case 17:
		//line sql.y:199
		{
			yyVAL.statement = &Update{Comments: Comments(yyS[yypt-6].bytes2), Table: yyS[yypt-5].tableName, Exprs: yyS[yypt-3].updateExprs, Where: NewWhere(AST_WHERE, yyS[yypt-2].boolExpr), OrderBy: yyS[yypt-1].orderBy, Limit: yyS[yypt-0].limit}
		}
	case 18:
		//line sql.y:205
		{
			yyVAL.statement = &Delete{Comments: Comments(yyS[yypt-5].bytes2), Table: yyS[yypt-3].tableName, Where: NewWhere(AST_WHERE, yyS[yypt-2].boolExpr), OrderBy: yyS[yypt-1].orderBy, Limit: yyS[yypt-0].limit}
		}
	case 19:
		//line sql.y:211
		{
			yyVAL.statement = &Set{Comments: Comments(yyS[yypt-1].bytes2), Exprs: yyS[yypt-0].updateExprs}
		}
	case 20:
		//line sql.y:217
		{
			yyVAL.statement = &DDL{Action: AST_CREATE, NewName: yyS[yypt-1].bytes}
		}
	case 21:
		//line sql.y:221
		{
			// Change this to an alter statement
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyS[yypt-1].bytes, NewName: yyS[yypt-1].bytes}
		}
	case 22:
		//line sql.y:226
		{
			yyVAL.statement = &DDL{Action: AST_CREATE, NewName: yyS[yypt-1].bytes}
		}
	case 23:
		//line sql.y:232
		{
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyS[yypt-2].bytes, NewName: yyS[yypt-2].bytes}
		}
	case 24:
		//line sql.y:236
		{
			// Change this to a rename statement
			yyVAL.statement = &DDL{Action: AST_RENAME, Table: yyS[yypt-3].bytes, NewName: yyS[yypt-0].bytes}
		}
	case 25:
		//line sql.y:241
		{
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyS[yypt-1].bytes, NewName: yyS[yypt-1].bytes}
		}
	case 26:
		//line sql.y:247
		{
			yyVAL.statement = &DDL{Action: AST_RENAME, Table: yyS[yypt-2].bytes, NewName: yyS[yypt-0].bytes}
		}
	case 27:
		//line sql.y:253
		{
			yyVAL.statement = &DDL{Action: AST_DROP, Table: yyS[yypt-0].bytes}
		}
	case 28:
		//line sql.y:257
		{
			// Change this to an alter statement
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyS[yypt-0].bytes, NewName: yyS[yypt-0].bytes}
		}
	case 29:
		//line sql.y:262
		{
			yyVAL.statement = &DDL{Action: AST_DROP, Table: yyS[yypt-1].bytes}
		}
	case 30:
		//line sql.y:268
		{
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyS[yypt-0].bytes, NewName: yyS[yypt-0].bytes}
		}
	case 31:
		//line sql.y:274
		{
			yyVAL.statement = &Other{}
		}
	case 32:
		//line sql.y:278
		{
			yyVAL.statement = &Other{}
		}
	case 33:
		//line sql.y:282
		{
			yyVAL.statement = &Other{}
		}
	case 34:
		//line sql.y:287
		{
			SetAllowComments(yylex, true)
		}
	case 35:
		//line sql.y:291
		{
			yyVAL.bytes2 = yyS[yypt-0].bytes2
			SetAllowComments(yylex, false)
		}
	case 36:
		//line sql.y:297
		{
			yyVAL.bytes2 = nil
		}
	case 37:
		//line sql.y:301
		{
			yyVAL.bytes2 = append(yyS[yypt-1].bytes2, yyS[yypt-0].bytes)
		}
	case 38:
		//line sql.y:307
		{
			yyVAL.str = AST_UNION
		}
	case 39:
		//line sql.y:311
		{
			yyVAL.str = AST_UNION_ALL
		}
	case 40:
		//line sql.y:315
		{
			yyVAL.str = AST_SET_MINUS
		}
	case 41:
		//line sql.y:319
		{
			yyVAL.str = AST_EXCEPT
		}
	case 42:
		//line sql.y:323
		{
			yyVAL.str = AST_INTERSECT
		}
	case 43:
		//line sql.y:328
		{
			yyVAL.str = ""
		}
	case 44:
		//line sql.y:332
		{
			yyVAL.str = AST_DISTINCT
		}
	case 45:
		//line sql.y:338
		{
			yyVAL.selectExprs = SelectExprs{yyS[yypt-0].selectExpr}
		}
	case 46:
		//line sql.y:342
		{
			yyVAL.selectExprs = append(yyVAL.selectExprs, yyS[yypt-0].selectExpr)
		}
	case 47:
		//line sql.y:348
		{
			yyVAL.selectExpr = &StarExpr{}
		}
	case 48:
		//line sql.y:352
		{
			yyVAL.selectExpr = &NonStarExpr{Expr: yyS[yypt-1].expr, As: yyS[yypt-0].bytes}
		}
	case 49:
		//line sql.y:356
		{
			yyVAL.selectExpr = &StarExpr{TableName: yyS[yypt-2].bytes}
		}
	case 50:
		//line sql.y:362
		{
			yyVAL.expr = yyS[yypt-0].boolExpr
		}
	case 51:
		//line sql.y:366
		{
			yyVAL.expr = yyS[yypt-0].valExpr
		}
	case 52:
		//line sql.y:371
		{
			yyVAL.bytes = nil
		}
	case 53:
		//line sql.y:375
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 54:
		//line sql.y:379
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 55:
		//line sql.y:385
		{
			yyVAL.tableExprs = TableExprs{yyS[yypt-0].tableExpr}
		}
	case 56:
		//line sql.y:389
		{
			yyVAL.tableExprs = append(yyVAL.tableExprs, yyS[yypt-0].tableExpr)
		}
	case 57:
		//line sql.y:395
		{
			yyVAL.tableExpr = &AliasedTableExpr{Expr: yyS[yypt-2].smTableExpr, As: yyS[yypt-1].bytes, Hints: yyS[yypt-0].indexHints}
		}
	case 58:
		//line sql.y:399
		{
			yyVAL.tableExpr = &ParenTableExpr{Expr: yyS[yypt-1].tableExpr}
		}
	case 59:
		//line sql.y:403
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyS[yypt-2].tableExpr, Join: yyS[yypt-1].str, RightExpr: yyS[yypt-0].tableExpr}
		}
	case 60:
		//line sql.y:407
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyS[yypt-4].tableExpr, Join: yyS[yypt-3].str, RightExpr: yyS[yypt-2].tableExpr, On: yyS[yypt-0].boolExpr}
		}
	case 61:
		//line sql.y:412
		{
			yyVAL.bytes = nil
		}
	case 62:
		//line sql.y:416
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 63:
		//line sql.y:420
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 64:
		//line sql.y:426
		{
			yyVAL.str = AST_JOIN
		}
	case 65:
		//line sql.y:430
		{
			yyVAL.str = AST_STRAIGHT_JOIN
		}
	case 66:
		//line sql.y:434
		{
			yyVAL.str = AST_LEFT_JOIN
		}
	case 67:
		//line sql.y:438
		{
			yyVAL.str = AST_LEFT_JOIN
		}
	case 68:
		//line sql.y:442
		{
			yyVAL.str = AST_RIGHT_JOIN
		}
	case 69:
		//line sql.y:446
		{
			yyVAL.str = AST_RIGHT_JOIN
		}
	case 70:
		//line sql.y:450
		{
			yyVAL.str = AST_JOIN
		}
	case 71:
		//line sql.y:454
		{
			yyVAL.str = AST_CROSS_JOIN
		}
	case 72:
		//line sql.y:458
		{
			yyVAL.str = AST_NATURAL_JOIN
		}
	case 73:
		//line sql.y:464
		{
			yyVAL.smTableExpr = &TableName{Name: yyS[yypt-0].bytes}
		}
	case 74:
		//line sql.y:468
		{
			yyVAL.smTableExpr = &TableName{Qualifier: yyS[yypt-2].bytes, Name: yyS[yypt-0].bytes}
		}
	case 75:
		//line sql.y:472
		{
			yyVAL.smTableExpr = yyS[yypt-0].subquery
		}
	case 76:
		//line sql.y:478
		{
			yyVAL.tableName = &TableName{Name: yyS[yypt-0].bytes}
		}
	case 77:
		//line sql.y:482
		{
			yyVAL.tableName = &TableName{Qualifier: yyS[yypt-2].bytes, Name: yyS[yypt-0].bytes}
		}
	case 78:
		//line sql.y:487
		{
			yyVAL.indexHints = nil
		}
	case 79:
		//line sql.y:491
		{
			yyVAL.indexHints = &IndexHints{Type: AST_USE, Indexes: yyS[yypt-1].bytes2}
		}
	case 80:
		//line sql.y:495
		{
			yyVAL.indexHints = &IndexHints{Type: AST_IGNORE, Indexes: yyS[yypt-1].bytes2}
		}
	case 81:
		//line sql.y:499
		{
			yyVAL.indexHints = &IndexHints{Type: AST_FORCE, Indexes: yyS[yypt-1].bytes2}
		}
	case 82:
		//line sql.y:505
		{
			yyVAL.bytes2 = [][]byte{yyS[yypt-0].bytes}
		}
	case 83:
		//line sql.y:509
		{
			yyVAL.bytes2 = append(yyS[yypt-2].bytes2, yyS[yypt-0].bytes)
		}
	case 84:
		//line sql.y:514
		{
			yyVAL.boolExpr = nil
		}
	case 85:
		//line sql.y:518
		{
			yyVAL.boolExpr = yyS[yypt-0].boolExpr
		}
	case 86:
		yyVAL.boolExpr = yyS[yypt-0].boolExpr
	case 87:
		//line sql.y:525
		{
			yyVAL.boolExpr = &AndExpr{Left: yyS[yypt-2].boolExpr, Right: yyS[yypt-0].boolExpr}
		}
	case 88:
		//line sql.y:529
		{
			yyVAL.boolExpr = &OrExpr{Left: yyS[yypt-2].boolExpr, Right: yyS[yypt-0].boolExpr}
		}
	case 89:
		//line sql.y:533
		{
			yyVAL.boolExpr = &NotExpr{Expr: yyS[yypt-0].boolExpr}
		}
	case 90:
		//line sql.y:537
		{
			yyVAL.boolExpr = &ParenBoolExpr{Expr: yyS[yypt-1].boolExpr}
		}
	case 91:
		//line sql.y:543
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyS[yypt-2].valExpr, Operator: yyS[yypt-1].str, Right: yyS[yypt-0].valExpr}
		}
	case 92:
		//line sql.y:547
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyS[yypt-2].valExpr, Operator: AST_IN, Right: yyS[yypt-0].colTuple}
		}
	case 93:
		//line sql.y:551
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyS[yypt-3].valExpr, Operator: AST_NOT_IN, Right: yyS[yypt-0].colTuple}
		}
	case 94:
		//line sql.y:555
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyS[yypt-2].valExpr, Operator: AST_LIKE, Right: yyS[yypt-0].valExpr}
		}
	case 95:
		//line sql.y:559
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyS[yypt-3].valExpr, Operator: AST_NOT_LIKE, Right: yyS[yypt-0].valExpr}
		}
	case 96:
		//line sql.y:563
		{
			yyVAL.boolExpr = &RangeCond{Left: yyS[yypt-4].valExpr, Operator: AST_BETWEEN, From: yyS[yypt-2].valExpr, To: yyS[yypt-0].valExpr}
		}
	case 97:
		//line sql.y:567
		{
			yyVAL.boolExpr = &RangeCond{Left: yyS[yypt-5].valExpr, Operator: AST_NOT_BETWEEN, From: yyS[yypt-2].valExpr, To: yyS[yypt-0].valExpr}
		}
	case 98:
		//line sql.y:571
		{
			yyVAL.boolExpr = &TrueCheck{Operator: AST_IS_TRUE, Expr: yyS[yypt-2].valExpr}
		}
	case 99:
		//line sql.y:575
		{
			yyVAL.boolExpr = &TrueCheck{Operator: AST_IS_NOT_TRUE, Expr: yyS[yypt-3].valExpr}
		}
	case 100:
		//line sql.y:579
		{
			yyVAL.boolExpr = &FalseCheck{Operator: AST_IS_FALSE, Expr: yyS[yypt-2].valExpr}
		}
	case 101:
		//line sql.y:583
		{
			yyVAL.boolExpr = &FalseCheck{Operator: AST_IS_NOT_FALSE, Expr: yyS[yypt-3].valExpr}
		}
	case 102:
		//line sql.y:587
		{
			yyVAL.boolExpr = &NullCheck{Operator: AST_IS_NULL, Expr: yyS[yypt-2].valExpr}
		}
	case 103:
		//line sql.y:591
		{
			yyVAL.boolExpr = &NullCheck{Operator: AST_IS_NOT_NULL, Expr: yyS[yypt-3].valExpr}
		}
	case 104:
		//line sql.y:595
		{
			yyVAL.boolExpr = &ExistsExpr{Subquery: yyS[yypt-0].subquery}
		}
	case 105:
		//line sql.y:599
		{
			yyVAL.boolExpr = &KeyrangeExpr{Start: yyS[yypt-3].valExpr, End: yyS[yypt-1].valExpr}
		}
	case 106:
		//line sql.y:605
		{
			yyVAL.str = AST_EQ
		}
	case 107:
		//line sql.y:609
		{
			yyVAL.str = AST_LT
		}
	case 108:
		//line sql.y:613
		{
			yyVAL.str = AST_GT
		}
	case 109:
		//line sql.y:617
		{
			yyVAL.str = AST_LE
		}
	case 110:
		//line sql.y:621
		{
			yyVAL.str = AST_GE
		}
	case 111:
		//line sql.y:625
		{
			yyVAL.str = AST_NE
		}
	case 112:
		//line sql.y:629
		{
			yyVAL.str = AST_NSE
		}
	case 113:
		//line sql.y:635
		{
			yyVAL.colTuple = ValTuple(yyS[yypt-1].valExprs)
		}
	case 114:
		//line sql.y:639
		{
			yyVAL.colTuple = yyS[yypt-0].subquery
		}
	case 115:
		//line sql.y:643
		{
			yyVAL.colTuple = ListArg(yyS[yypt-0].bytes)
		}
	case 116:
		//line sql.y:649
		{
			yyVAL.subquery = &Subquery{yyS[yypt-1].selStmt}
		}
	case 117:
		//line sql.y:655
		{
			yyVAL.valExprs = ValExprs{yyS[yypt-0].valExpr}
		}
	case 118:
		//line sql.y:659
		{
			yyVAL.valExprs = append(yyS[yypt-2].valExprs, yyS[yypt-0].valExpr)
		}
	case 119:
		//line sql.y:665
		{
			yyVAL.valExpr = yyS[yypt-0].valExpr
		}
	case 120:
		//line sql.y:669
		{
			yyVAL.valExpr = yyS[yypt-0].colName
		}
	case 121:
		//line sql.y:673
		{
			yyVAL.valExpr = yyS[yypt-0].rowTuple
		}
	case 122:
		//line sql.y:677
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_BITAND, Right: yyS[yypt-0].valExpr}
		}
	case 123:
		//line sql.y:681
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_BITOR, Right: yyS[yypt-0].valExpr}
		}
	case 124:
		//line sql.y:685
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_BITXOR, Right: yyS[yypt-0].valExpr}
		}
	case 125:
		//line sql.y:689
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_PLUS, Right: yyS[yypt-0].valExpr}
		}
	case 126:
		//line sql.y:693
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_MINUS, Right: yyS[yypt-0].valExpr}
		}
	case 127:
		//line sql.y:697
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_MULT, Right: yyS[yypt-0].valExpr}
		}
	case 128:
		//line sql.y:701
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_DIV, Right: yyS[yypt-0].valExpr}
		}
	case 129:
		//line sql.y:705
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_MOD, Right: yyS[yypt-0].valExpr}
		}
	case 130:
		//line sql.y:709
		{
			if num, ok := yyS[yypt-0].valExpr.(NumVal); ok {
				switch yyS[yypt-1].byt {
				case '-':
					yyVAL.valExpr = append(NumVal("-"), num...)
				case '+':
					yyVAL.valExpr = num
				default:
					yyVAL.valExpr = &UnaryExpr{Operator: yyS[yypt-1].byt, Expr: yyS[yypt-0].valExpr}
				}
			} else {
				yyVAL.valExpr = &UnaryExpr{Operator: yyS[yypt-1].byt, Expr: yyS[yypt-0].valExpr}
			}
		}
	case 131:
		//line sql.y:724
		{
			yyVAL.valExpr = &FuncExpr{Name: yyS[yypt-2].bytes}
		}
	case 132:
		//line sql.y:728
		{
			yyVAL.valExpr = &FuncExpr{Name: yyS[yypt-3].bytes, Exprs: yyS[yypt-1].selectExprs}
		}
	case 133:
		//line sql.y:732
		{
			yyVAL.valExpr = &FuncExpr{Name: yyS[yypt-4].bytes, Distinct: true, Exprs: yyS[yypt-1].selectExprs}
		}
	case 134:
		//line sql.y:736
		{
			yyVAL.valExpr = &FuncExpr{Name: yyS[yypt-3].bytes, Exprs: yyS[yypt-1].selectExprs}
		}
	case 135:
		//line sql.y:740
		{
			yyVAL.valExpr = yyS[yypt-0].caseExpr
		}
	case 136:
		//line sql.y:746
		{
			yyVAL.bytes = IF_BYTES
		}
	case 137:
		//line sql.y:750
		{
			yyVAL.bytes = VALUES_BYTES
		}
	case 138:
		//line sql.y:756
		{
			yyVAL.byt = AST_UPLUS
		}
	case 139:
		//line sql.y:760
		{
			yyVAL.byt = AST_UMINUS
		}
	case 140:
		//line sql.y:764
		{
			yyVAL.byt = AST_TILDA
		}
	case 141:
		//line sql.y:770
		{
			yyVAL.caseExpr = &CaseExpr{Expr: yyS[yypt-3].valExpr, Whens: yyS[yypt-2].whens, Else: yyS[yypt-1].valExpr}
		}
	case 142:
		//line sql.y:775
		{
			yyVAL.valExpr = nil
		}
	case 143:
		//line sql.y:779
		{
			yyVAL.valExpr = yyS[yypt-0].valExpr
		}
	case 144:
		//line sql.y:785
		{
			yyVAL.whens = []*When{yyS[yypt-0].when}
		}
	case 145:
		//line sql.y:789
		{
			yyVAL.whens = append(yyS[yypt-1].whens, yyS[yypt-0].when)
		}
	case 146:
		//line sql.y:795
		{
			yyVAL.when = &When{Cond: yyS[yypt-2].boolExpr, Val: yyS[yypt-0].valExpr}
		}
	case 147:
		//line sql.y:800
		{
			yyVAL.valExpr = nil
		}
	case 148:
		//line sql.y:804
		{
			yyVAL.valExpr = yyS[yypt-0].valExpr
		}
	case 149:
		//line sql.y:810
		{
			yyVAL.colName = &ColName{Name: yyS[yypt-0].bytes}
		}
	case 150:
		//line sql.y:814
		{
			yyVAL.colName = &ColName{Qualifier: yyS[yypt-2].bytes, Name: yyS[yypt-0].bytes}
		}
	case 151:
		//line sql.y:820
		{
			yyVAL.valExpr = StrVal(yyS[yypt-0].bytes)
		}
	case 152:
		//line sql.y:824
		{
			yyVAL.valExpr = NumVal(yyS[yypt-0].bytes)
		}
	case 153:
		//line sql.y:828
		{
			yyVAL.valExpr = ValArg(yyS[yypt-0].bytes)
		}
	case 154:
		//line sql.y:832
		{
			yyVAL.valExpr = &NullVal{}
		}
	case 155:
		//line sql.y:836
		{
			yyVAL.valExpr = &TrueVal{}
		}
	case 156:
		//line sql.y:840
		{
			yyVAL.valExpr = &FalseVal{}
		}
	case 157:
		//line sql.y:845
		{
			yyVAL.valExprs = nil
		}
	case 158:
		//line sql.y:849
		{
			yyVAL.valExprs = yyS[yypt-0].valExprs
		}
	case 159:
		//line sql.y:854
		{
			yyVAL.boolExpr = nil
		}
	case 160:
		//line sql.y:858
		{
			yyVAL.boolExpr = yyS[yypt-0].boolExpr
		}
	case 161:
		//line sql.y:863
		{
			yyVAL.orderBy = nil
		}
	case 162:
		//line sql.y:867
		{
			yyVAL.orderBy = yyS[yypt-0].orderBy
		}
	case 163:
		//line sql.y:873
		{
			yyVAL.orderBy = OrderBy{yyS[yypt-0].order}
		}
	case 164:
		//line sql.y:877
		{
			yyVAL.orderBy = append(yyS[yypt-2].orderBy, yyS[yypt-0].order)
		}
	case 165:
		//line sql.y:883
		{
			yyVAL.order = &Order{Expr: yyS[yypt-1].valExpr, Direction: yyS[yypt-0].str}
		}
	case 166:
		//line sql.y:888
		{
			yyVAL.str = AST_ASC
		}
	case 167:
		//line sql.y:892
		{
			yyVAL.str = AST_ASC
		}
	case 168:
		//line sql.y:896
		{
			yyVAL.str = AST_DESC
		}
	case 169:
		//line sql.y:901
		{
			yyVAL.limit = nil
		}
	case 170:
		//line sql.y:905
		{
			yyVAL.limit = &Limit{Rowcount: yyS[yypt-0].valExpr}
		}
	case 171:
		//line sql.y:909
		{
			yyVAL.limit = &Limit{Offset: yyS[yypt-2].valExpr, Rowcount: yyS[yypt-0].valExpr}
		}
	case 172:
		//line sql.y:914
		{
			yyVAL.str = ""
		}
	case 173:
		//line sql.y:918
		{
			yyVAL.str = AST_FOR_UPDATE
		}
	case 174:
		//line sql.y:922
		{
			if !bytes.Equal(yyS[yypt-1].bytes, SHARE) {
				yylex.Error("expecting share")
				return 1
			}
			if !bytes.Equal(yyS[yypt-0].bytes, MODE) {
				yylex.Error("expecting mode")
				return 1
			}
			yyVAL.str = AST_SHARE_MODE
		}
	case 175:
		//line sql.y:935
		{
			yyVAL.columns = nil
		}
	case 176:
		//line sql.y:939
		{
			yyVAL.columns = yyS[yypt-1].columns
		}
	case 177:
		//line sql.y:945
		{
			yyVAL.columns = Columns{&NonStarExpr{Expr: yyS[yypt-0].colName}}
		}
	case 178:
		//line sql.y:949
		{
			yyVAL.columns = append(yyVAL.columns, &NonStarExpr{Expr: yyS[yypt-0].colName})
		}
	case 179:
		//line sql.y:954
		{
			yyVAL.updateExprs = nil
		}
	case 180:
		//line sql.y:958
		{
			yyVAL.updateExprs = yyS[yypt-0].updateExprs
		}
	case 181:
		//line sql.y:964
		{
			yyVAL.insRows = yyS[yypt-0].values
		}
	case 182:
		//line sql.y:968
		{
			yyVAL.insRows = yyS[yypt-0].selStmt
		}
	case 183:
		//line sql.y:974
		{
			yyVAL.values = Values{yyS[yypt-0].rowTuple}
		}
	case 184:
		//line sql.y:978
		{
			yyVAL.values = append(yyS[yypt-2].values, yyS[yypt-0].rowTuple)
		}
	case 185:
		//line sql.y:984
		{
			yyVAL.rowTuple = ValTuple(yyS[yypt-1].valExprs)
		}
	case 186:
		//line sql.y:988
		{
			yyVAL.rowTuple = yyS[yypt-0].subquery
		}
	case 187:
		//line sql.y:994
		{
			yyVAL.updateExprs = UpdateExprs{yyS[yypt-0].updateExpr}
		}
	case 188:
		//line sql.y:998
		{
			yyVAL.updateExprs = append(yyS[yypt-2].updateExprs, yyS[yypt-0].updateExpr)
		}
	case 189:
		//line sql.y:1004
		{
			yyVAL.updateExpr = &UpdateExpr{Name: yyS[yypt-2].colName, Expr: yyS[yypt-0].valExpr}
		}
	case 190:
		//line sql.y:1009
		{
			yyVAL.empty = struct{}{}
		}
	case 191:
		//line sql.y:1011
		{
			yyVAL.empty = struct{}{}
		}
	case 192:
		//line sql.y:1014
		{
			yyVAL.empty = struct{}{}
		}
	case 193:
		//line sql.y:1016
		{
			yyVAL.empty = struct{}{}
		}
	case 194:
		//line sql.y:1019
		{
			yyVAL.empty = struct{}{}
		}
	case 195:
		//line sql.y:1021
		{
			yyVAL.empty = struct{}{}
		}
	case 196:
		//line sql.y:1025
		{
			yyVAL.empty = struct{}{}
		}
	case 197:
		//line sql.y:1027
		{
			yyVAL.empty = struct{}{}
		}
	case 198:
		//line sql.y:1029
		{
			yyVAL.empty = struct{}{}
		}
	case 199:
		//line sql.y:1031
		{
			yyVAL.empty = struct{}{}
		}
	case 200:
		//line sql.y:1033
		{
			yyVAL.empty = struct{}{}
		}
	case 201:
		//line sql.y:1036
		{
			yyVAL.empty = struct{}{}
		}
	case 202:
		//line sql.y:1038
		{
			yyVAL.empty = struct{}{}
		}
	case 203:
		//line sql.y:1041
		{
			yyVAL.empty = struct{}{}
		}
	case 204:
		//line sql.y:1043
		{
			yyVAL.empty = struct{}{}
		}
	case 205:
		//line sql.y:1046
		{
			yyVAL.empty = struct{}{}
		}
	case 206:
		//line sql.y:1048
		{
			yyVAL.empty = struct{}{}
		}
	case 207:
		//line sql.y:1052
		{
			yyVAL.bytes = bytes.ToLower(yyS[yypt-0].bytes)
		}
	case 208:
		//line sql.y:1057
		{
			ForceEOF(yylex)
		}
	}
	goto yystack /* stack new state and value */
}
