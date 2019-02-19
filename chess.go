package chess

/*
import(
	"strings"
	"strconv"
)
*/

type squareID int

const emptySquare = -1

// PieceType identifies the type of a piece.
type PieceType int

const pawn = 'p'
const knight = 'n'
const bishop = 'b'
const rook = 'r'
const queen = 'q'
const king = 'k'

// PieceColor is either black or white
type PieceColor int

const black = 'b'
const white = 'w'

//Piece represents a piece on the board, having a type such as pawn or bishop and a color, white or black
type Piece struct {
	ptype  PieceType
	pcolor PieceColor
}

const pieceSymbols = "pnbrqkPNBRQK"

// Starting position in FEN
const defaultPosition = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"

// Can't do const arrays, but this is const.
var possibleResults = []string{"1-0", "0-1", "1/2-1/2", "*"}

var pawnOffsets = map[PieceColor][]int{
	black: {16, 32, 17, 15},
	white: {-16, -32, -17, -15}}

var pieceOffsets = map[PieceType][]int{
	knight: {-18, -33, -31, -14, 18, 33, 31, 14},
	bishop: {-17, -15, 17, 15},
	rook:   {-16, 1, 16, -1},
	queen:  {-17, -16, -15, 1, 17, 16, 15, -1},
	king:   {-17, -16, -15, 1, 17, 16, 15, -1}}

var attacks = []int{
	20, 0, 0, 0, 0, 0, 0, 24, 0, 0, 0, 0, 0, 0, 20, 0,
	0, 20, 0, 0, 0, 0, 0, 24, 0, 0, 0, 0, 0, 20, 0, 0,
	0, 0, 20, 0, 0, 0, 0, 24, 0, 0, 0, 0, 20, 0, 0, 0,
	0, 0, 0, 20, 0, 0, 0, 24, 0, 0, 0, 20, 0, 0, 0, 0,
	0, 0, 0, 0, 20, 0, 0, 24, 0, 0, 20, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 20, 2, 24, 2, 20, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 2, 53, 56, 53, 2, 0, 0, 0, 0, 0, 0,
	24, 24, 24, 24, 24, 24, 56, 0, 56, 24, 24, 24, 24, 24, 24, 0,
	0, 0, 0, 0, 0, 2, 53, 56, 53, 2, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 20, 2, 24, 2, 20, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 20, 0, 0, 24, 0, 0, 20, 0, 0, 0, 0, 0,
	0, 0, 0, 20, 0, 0, 0, 24, 0, 0, 0, 20, 0, 0, 0, 0,
	0, 0, 20, 0, 0, 0, 0, 24, 0, 0, 0, 0, 20, 0, 0, 0,
	0, 20, 0, 0, 0, 0, 0, 24, 0, 0, 0, 0, 0, 20, 0, 0,
	20, 0, 0, 0, 0, 0, 0, 24, 0, 0, 0, 0, 0, 0, 20}

var rays = []int{
	17, 0, 0, 0, 0, 0, 0, 16, 0, 0, 0, 0, 0, 0, 15, 0,
	0, 17, 0, 0, 0, 0, 0, 16, 0, 0, 0, 0, 0, 15, 0, 0,
	0, 0, 17, 0, 0, 0, 0, 16, 0, 0, 0, 0, 15, 0, 0, 0,
	0, 0, 0, 17, 0, 0, 0, 16, 0, 0, 0, 15, 0, 0, 0, 0,
	0, 0, 0, 0, 17, 0, 0, 16, 0, 0, 15, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 17, 0, 16, 0, 15, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 17, 16, 15, 0, 0, 0, 0, 0, 0, 0,
	1, 1, 1, 1, 1, 1, 1, 0, -1, -1, -1, -1, -1, -1, -1, 0,
	0, 0, 0, 0, 0, 0, -15, -16, -17, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, -15, 0, -16, 0, -17, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, -15, 0, 0, -16, 0, 0, -17, 0, 0, 0, 0, 0,
	0, 0, 0, -15, 0, 0, 0, -16, 0, 0, 0, -17, 0, 0, 0, 0,
	0, 0, -15, 0, 0, 0, 0, -16, 0, 0, 0, 0, -17, 0, 0, 0,
	0, -15, 0, 0, 0, 0, 0, -16, 0, 0, 0, 0, 0, -17, 0, 0,
	-15, 0, 0, 0, 0, 0, 0, -16, 0, 0, 0, 0, 0, 0, -17}

var shifts = map[PieceType]int{
	pawn:   0,
	knight: 1,
	bishop: 2,
	rook:   3,
	queen:  4,
	king:   5}

// Bitwise values for Move.flags
const normal = 1
const capture = 2
const bigPawn = 4
const enpassantCapture = 8
const promotion = 16
const ksideCastle = 32
const qsideCastle = 64

const rank1 = 7
const rank2 = 6
const rank3 = 5
const rank4 = 4
const rank5 = 3
const rank6 = 2
const rank7 = 1
const rank8 = 0

var squareNameToID = map[string]int{
	"a8": 0, "b8": 1, "c8": 2, "d8": 3, "e8": 4, "f8": 5, "g8": 6, "h8": 7,
	"a7": 16, "b7": 17, "c7": 18, "d7": 19, "e7": 20, "f7": 21, "g7": 22, "h7": 23,
	"a6": 32, "b6": 33, "c6": 34, "d6": 35, "e6": 36, "f6": 37, "g6": 38, "h6": 39,
	"a5": 48, "b5": 49, "c5": 50, "d5": 51, "e5": 52, "f5": 53, "g5": 54, "h5": 55,
	"a4": 64, "b4": 65, "c4": 66, "d4": 67, "e4": 68, "f4": 69, "g4": 70, "h4": 71,
	"a3": 80, "b3": 81, "c3": 82, "d3": 83, "e3": 84, "f3": 85, "g3": 86, "h3": 87,
	"a2": 96, "b2": 97, "c2": 98, "d2": 99, "e2": 100, "f2": 101, "g2": 102, "h2": 103,
	"a1": 112, "b1": 113, "c1": 114, "d1": 115, "e1": 116, "f1": 117, "g1": 118, "h1": 119}

var rooks = map[PieceColor][][]int{
	white: {{squareNameToID["a1"], qsideCastle},
		{squareNameToID["h1"], ksideCastle}},
	black: {{squareNameToID["a8"], qsideCastle},
		{squareNameToID["h8"], ksideCastle}}}

// The current state or allowability of castling for the board. The members
// contain the bitwise value of ksideCastle and qsideCastlea
type castlingState struct {
	white int8
	black int8
}

// Move describes a move on the board
type Move struct {
	turn  PieceColor
	from  int32
	to    int32
	ptype PieceType
	flags int8
}

type kingsLocation struct {
	white squareID
	black squareID
}

type historyEntry struct {
	move Move
	kingsLocation
	turn            PieceColor
	castling        castlingState
	enpassantSquare squareID
	halfMoves       int32
	moveNumber      int32
}

// Chess defines the current structure of a chess game
type Chess struct {
	board           []Piece
	turn            PieceColor
	enpassantSquare squareID
	halfMoves       int32
	moveNumber      int32
	castling        castlingState
	kings           kingsLocation
	history         []historyEntry
	header          map[string]string
}

// New creates a new Chess instance initialized to the starting/default chess position
func New() *Chess {
	retVal := new(Chess)
	return (retVal)
}

// Clear sets the Chess instance to the starting position
func (chess *Chess) Clear() {
	chess.board = make([]Piece, 128)
	chess.turn = white
	chess.castling = castlingState{0, 0}
	chess.enpassantSquare = emptySquare
	chess.halfMoves = 0
	chess.moveNumber = 1
	chess.history = make([]historyEntry, 0)
	chess.header = make(map[string]string)
	//	chess.updateSetup(chess.generateFen())
}

/*
func (chess *Chess) updateSetup(fen string){
}

func (chess *Chess) generateFen() string{
	emptySquares := 0
	var retVal strings.Builder
	for cntr := squareNameToID["a8"]; cntr <= squareNameToID["h1"]; cntr++ {
		if unspecifiedPiece(chess.board[cntr]){
			emptySquares++
		} else {
			if emptySquares > 0 {
				retVal.WriteString(strconv.Itoa(emptySquares))
				emptySquares = 0
			}
		}
	}
	return retVal.String()
}
*/

func (p Piece) isUnspecified() bool {
	retVal := p.pcolor == 0 || p.ptype == 0
	return (retVal)
}
