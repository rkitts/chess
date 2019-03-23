package chess

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

const emptySquare = -1

// PieceType identifies the type of a piece.
type PieceType int

const pawn PieceType = 'p'
const knight PieceType = 'n'
const bishop PieceType = 'b'
const rook PieceType = 'r'
const queen PieceType = 'q'
const king PieceType = 'k'

var promotionTypes = []PieceType{queen, rook, bishop, knight}

// PieceColor is either black or white
type PieceColor int

const black PieceColor = 'b'
const white PieceColor = 'w'

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

var shifts = map[PieceType]uint{
	pawn:   0,
	knight: 1,
	bishop: 2,
	rook:   3,
	queen:  4,
	king:   5}

// Bitwise values for Move.flags
const (
	normalMove      = 1
	captureMove     = 2
	bigPawnMove     = 4
	enpassantMove   = 8
	promotionMove   = 16
	ksideCastleMove = 32
	qsideCastleMove = 64
)

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
	white: {{squareNameToID["a1"], qsideCastleMove},
		{squareNameToID["h1"], ksideCastleMove}},
	black: {{squareNameToID["a8"], qsideCastleMove},
		{squareNameToID["h8"], ksideCastleMove}}}

// The current state or allowability of castling for the board. The members
// contain the bitwise value of ksideCastle and qsideCastle
type castlingState map[PieceColor]int

// Move describes a move on the board
type Move struct {
	turn         PieceColor
	from         int
	to           int
	ptype        PieceType
	flags        int
	promotedType PieceType
	capturedType PieceType
}

var secondRank = map[PieceColor]int{black: rank7, white: rank2}

type kingsLocation map[PieceColor]int

type historyEntry struct {
	move            Move
	kings           kingsLocation
	turn            PieceColor
	castling        castlingState
	enpassantSquare int
	halfMoves       int
	moveNumber      int
}

// Chess defines the current structure of a chess game
type Chess struct {
	board           []Piece
	turn            PieceColor
	enpassantSquare int
	halfMoves       int
	moveNumber      int
	castling        castlingState
	kings           kingsLocation
	history         Stack
	header          map[string]string
}

// New creates a new Chess instance initialized to the starting/default chess position
func New() *Chess {
	retVal := new(Chess)
	retVal.Clear()
	retVal.load(defaultPosition)
	return (retVal)
}

// Reset sets the board to the default/starting position
func (chess *Chess) Reset() {
	chess.load(defaultPosition)
}

func (chess *Chess) get(squareID string) Piece {
	var retVal Piece
	if squareNum, ok := squareNameToID[squareID]; ok {
		retVal = chess.board[squareNum]
	}
	return retVal
}

func (chess *Chess) put(piece Piece, squareName string) error {
	var retVal error

	if squareID, ok := squareNameToID[squareName]; ok == false {
		retVal = fmt.Errorf("%s is not a legal square name", squareName)
	} else {
		retVal = chess.maybeUpdateKings(piece, squareID)
		if retVal == nil {
			chess.board[squareID] = piece
			chess.updateSetup(chess.generateFen())
		}
	}
	return retVal
}

func (chess *Chess) remove(squareID string) Piece {
	var retVal = chess.get(squareID)
	if squareNum, ok := squareNameToID[squareID]; ok {
		var replacementPiece Piece
		chess.board[squareNum] = replacementPiece
	}
	chess.updateSetup(chess.generateFen())
	return retVal
}

func (chess *Chess) maybeUpdateKings(piece Piece, squareID int) error {

	var retVal error

	if piece.ptype == king {
		switch piece.pcolor {
		case white:
			if chess.kings[white] == emptySquare {
				chess.kings[white] = squareID
			} else if chess.kings[white] != squareID {
				retVal = fmt.Errorf("White king already on board")
			}
		case black:
			if chess.kings[black] == emptySquare {
				chess.kings[black] = squareID
			} else if chess.kings[black] != squareID {
				retVal = fmt.Errorf("Black king already on board")
			}
		}
	}
	return retVal
}

func (chess *Chess) load(fenToLoad string) error {
	fen, err := parseFEN(fenToLoad)
	if err == nil {
		square := 0
		for cntr := 0; cntr < len(fen.piecePlacement); cntr++ {
			maybePiece := fen.piecePlacement[cntr]
			if maybePiece == '/' {
				square += 8
			} else if unicode.IsDigit(rune(maybePiece)) {
				square += int(maybePiece - '0')
			} else {
				var color PieceColor
				color = black
				if maybePiece < 'a' {
					color = white
				}
				// TODO: This is terrible. Somehow make it so PieceType is a little more
				// robust instead of this upper/lowercase crap.
				pieceType := PieceType(unicode.ToLower(rune(maybePiece)))
				chess.put(Piece{pieceType, color}, algebraic(square))
				square++
			}
		}
		chess.turn = fen.activeColor
		chess.castling = fen.castlingAbility
		if fen.enpassantCapture == "-" {
			chess.enpassantSquare = emptySquare
		} else {
			chess.enpassantSquare = squareNameToID[fen.enpassantCapture]
		}
		chess.halfMoves = fen.halfMoves
		chess.moveNumber = fen.fullMoves
	}
	return err
}

// Clear sets the Chess instance to the starting position
func (chess *Chess) Clear() {
	chess.board = make([]Piece, 128)
	chess.turn = white
	chess.castling = castlingState{white: 0, black: 0}
	chess.enpassantSquare = emptySquare
	chess.halfMoves = 0
	chess.moveNumber = 1
	chess.kings = kingsLocation{white: emptySquare, black: emptySquare}
	chess.kings[black] = emptySquare
	chess.kings[white] = emptySquare
	chess.header = make(map[string]string)
}

// UndoMove takes the most recently pushed history item and undoes it's effects
func (chess *Chess) UndoMove() (Move, bool) {
	var retVal Move
	foundOne := false
	if chess.history.Len() != 0 {
		history := chess.history.Pop().(historyEntry)
		foundOne = true

		chess.applyHistoryEntry(history)
		chess.applyHistoryMove(history.move)
		chess.undoCapture(history.move)
		chess.undoCastling(history.move)
		retVal = history.move
	}
	return retVal, foundOne
}

func (chess *Chess) undoCastling(move Move) {
	if move.flags&(ksideCastleMove|qsideCastleMove) != 0 {
		var castlingTo int
		var castlingFrom int

		if move.flags&ksideCastleMove != 0 {
			castlingTo = move.to + 1
			castlingFrom = move.to - 1
		} else if move.flags&qsideCastleMove != 0 {
			castlingTo = move.to - 2
			castlingFrom = move.to + 1
		}

		chess.board[castlingTo] = chess.board[castlingFrom]
		chess.board[castlingFrom] = Piece{}
	}
}

func (chess *Chess) undoCapture(move Move) {
	ourColor := move.turn
	theirColor := swapColor(ourColor)
	if move.flags&captureMove != 0 {
		chess.board[move.to] = Piece{ptype: move.capturedType, pcolor: theirColor}
	} else if move.flags&enpassantMove != 0 {
		var index int
		if ourColor == black {
			index = move.to - 16
		} else {
			index = move.to + 16
		}
		chess.board[index] = Piece{ptype: pawn, pcolor: theirColor}
	}
}

func (chess *Chess) applyHistoryMove(moveToUndo Move) {
	chess.board[moveToUndo.from] = chess.board[moveToUndo.to]
	// Undo any promotions
	chess.board[moveToUndo.from].ptype = moveToUndo.ptype
	chess.board[moveToUndo.to] = Piece{}
}

func (chess *Chess) applyHistoryEntry(history historyEntry) {
	chess.kings = history.kings
	chess.turn = history.turn
	chess.castling = history.castling
	chess.enpassantSquare = history.enpassantSquare
	chess.halfMoves = history.halfMoves
	chess.moveNumber = history.moveNumber
}

func (chess *Chess) updateSetup(fen string) {
	if chess.history.Len() == 0 {
		if fen != defaultPosition {
			chess.header["SetUp"] = "1"
			chess.header["FEN"] = fen
		} else {
			delete(chess.header, "SetUp")
			delete(chess.header, "FEN")
		}
	}
}

func (chess *Chess) generateFen() string {
	emptySquares := 0
	var retVal strings.Builder
	for rankValue := 0; rankValue <= 112; rankValue += 16 {
		for fileValue := 0; fileValue < 8; fileValue++ {
			squareID := squareNameToID[algebraic(rankValue+fileValue)]
			if chess.board[squareID].isUnspecified() {
				emptySquares++
			} else {
				if emptySquares > 0 {
					retVal.WriteString(strconv.Itoa(emptySquares))
					emptySquares = 0
				}
				pieceCode := rune(chess.board[squareID].ptype)
				if chess.board[squareID].pcolor == white {
					pieceCode = unicode.ToUpper(rune(pieceCode))
				}
				retVal.WriteRune(pieceCode)
			}
			if ((squareID + 1) & 0x88) != 0 {
				if emptySquares > 0 {
					retVal.WriteString(strconv.Itoa(emptySquares))
					emptySquares = 0
				}
				if squareID != squareNameToID["h1"] {
					retVal.WriteString("/")
				}
			}
		}
	}
	retVal.WriteString(" ")
	retVal.WriteRune(rune(chess.turn))
	retVal.WriteString(" ")
	retVal.WriteString(generateCastlingFEN(chess.castling))
	retVal.WriteString(" ")
	retVal.WriteString(generateEnpassantFEN(chess.enpassantSquare))
	retVal.WriteString(" ")
	retVal.WriteString(strconv.Itoa(chess.halfMoves))
	retVal.WriteString(" ")
	retVal.WriteString(strconv.Itoa(chess.moveNumber))
	return retVal.String()
}

func generateEnpassantFEN(squareID int) string {
	retVal := "-"
	if squareID != emptySquare {
		retVal = algebraic(squareID)
	}
	return retVal
}

func generateCastlingFEN(castling castlingState) string {
	var retVal strings.Builder

	if castling[white]&ksideCastleMove != 0 {
		retVal.WriteString("K")
	}
	if castling[white]&qsideCastleMove != 0 {
		retVal.WriteString("Q")
	}

	if castling[black]&ksideCastleMove != 0 {
		retVal.WriteString("k")
	}
	if castling[black]&qsideCastleMove != 0 {
		retVal.WriteString("q")
	}
	if retVal.Len() == 0 {
		retVal.WriteString("-")
	}
	return retVal.String()
}

func (p *Piece) isUnspecified() bool {
	retVal := p.pcolor == 0 || p.ptype == 0
	return (retVal)
}

func (chess *Chess) buildMove(from int, to int, flags int, promotionType PieceType) Move {
	retVal := Move{
		turn:  chess.turn,
		from:  from,
		to:    to,
		flags: flags,
		ptype: chess.board[from].ptype}

	if promotionType != 0 {
		retVal.flags |= promotionMove
		retVal.promotedType = promotionType
	}

	if !chess.board[to].isUnspecified() {
		retVal.capturedType = chess.board[to].ptype
	} else if flags&enpassantMove != 0 {
		retVal.capturedType = pawn
	}
	return retVal
}

func (chess *Chess) makeMove(moveToMake Move) {
	chess.pushHistory(moveToMake)

	ourColor := chess.turn
	theirColor := swapColor(ourColor)

	chess.board[moveToMake.to] = chess.board[moveToMake.from]

	if moveToMake.flags&enpassantMove != 0 {
		if ourColor == black {
			chess.board[moveToMake.to-16] = Piece{}
		} else {
			chess.board[moveToMake.to+16] = Piece{}
		}
	}

	if moveToMake.flags&promotionMove != 0 {
		chess.board[moveToMake.to] = Piece{pcolor: ourColor, ptype: moveToMake.promotedType}
	}

	if chess.board[moveToMake.to].ptype == king {
		chess.kings[ourColor] = moveToMake.to

		if moveToMake.flags&ksideCastleMove != 0 {
			// Move the rook next to the king
			castlingTo := moveToMake.to - 1
			castlingFrom := moveToMake.to + 1
			chess.board[castlingTo] = chess.board[castlingFrom]
			chess.board[castlingFrom] = Piece{}
		} else if moveToMake.flags&qsideCastleMove != 0 {
			castlingTo := moveToMake.to + 1
			castlingFrom := moveToMake.to - 2
			chess.board[castlingTo] = chess.board[castlingFrom]
			chess.board[castlingFrom] = Piece{}
		}
		chess.castling[ourColor] = 0
	}

	// Turn off castling if we move a rook
	if chess.castling[ourColor] != 0 {
		rLength := len(rooks[ourColor])
		for cntr := 0; cntr < rLength; cntr++ {
			if moveToMake.from == rooks[ourColor][cntr][0] &&
				chess.castling[ourColor]&rooks[ourColor][cntr][1] != 0 {
				chess.castling[ourColor] ^= rooks[ourColor][cntr][1]
				break
			}
		}
	}

	// Turn off castling if we capture a rook
	if chess.castling[theirColor] != 0 {
		rLength := len(rooks[theirColor])
		for cntr := 0; cntr < rLength; cntr++ {
			if moveToMake.to == rooks[theirColor][cntr][0] &&
				chess.castling[theirColor]&rooks[theirColor][cntr][1] != 0 {
				chess.castling[theirColor] ^= rooks[theirColor][cntr][1]
				break
			}
		}
	}

	chess.updateEnpassantSquare(moveToMake)
	chess.updateMoveCounters(moveToMake)
	chess.turn = swapColor(ourColor)
}

func (chess *Chess) updateMoveCounters(move Move) {
	/* reset the 50 move counter if a pawn is moved or a piece is captured */
	if move.ptype == pawn {
		chess.halfMoves = 0
	} else if (move.flags & (captureMove | enpassantMove)) != 0 {
		chess.halfMoves = 0
	} else {
		chess.halfMoves++
	}

	if move.turn == black {
		chess.moveNumber++
	}
}

func (chess *Chess) updateEnpassantSquare(move Move) {
	// If big pawn move, update the enpassant square
	if move.flags&bigPawnMove != 0 {
		if move.turn == black {
			chess.enpassantSquare = move.to - 16
		} else {
			chess.enpassantSquare = move.to + 16
		}
	} else {
		chess.enpassantSquare = emptySquare
	}
}

func (chess *Chess) pushHistory(move Move) {
	var entry historyEntry

	entry.halfMoves = chess.halfMoves
	entry.moveNumber = chess.moveNumber
	entry.enpassantSquare = chess.enpassantSquare
	entry.move = move
	entry.turn = chess.turn
	entry.kings = make(kingsLocation)
	entry.kings[white] = chess.kings[white]
	entry.kings[black] = chess.kings[black]
	entry.castling = make(castlingState)
	entry.castling[white] = chess.castling[white]
	entry.castling[black] = chess.castling[black]
	chess.history.Push(entry)
}

func (chess *Chess) generateMoves(legalMoves bool, singleSquareName string) []Move {

	var retVal []Move
	ourColor := chess.turn
	firstSquare, lastSquare, err := chess.determineSquareRange(singleSquareName)
	if err == nil {
		var allMoves []Move
		for cntr := firstSquare; cntr <= lastSquare; cntr++ {
			if cntr&0x88 != 0 {
				cntr += 7
				continue
			}
			currPiece := chess.board[cntr]
			if currPiece.isUnspecified() || currPiece.pcolor != ourColor {
				continue
			}
			if currPiece.ptype == pawn {
				// Pawn moves...
				allMoves = append(allMoves, chess.getPawnMoves(cntr, ourColor)...)
				allMoves = append(allMoves, chess.getPawnAttacks(cntr, ourColor)...)
			} else {
				allMoves = append(allMoves, chess.getPieceMoves(cntr, currPiece)...)
			}
		}

		singleSquare := firstSquare == lastSquare
		if !singleSquare || lastSquare == chess.kings[ourColor] {
			allMoves = append(allMoves, chess.getCastlingMoves(ourColor)...)
		}

		if legalMoves {
			for _, move := range allMoves {
				chess.makeMove(move)
				if !chess.kingAttacked(chess.turn) {
					retVal = append(retVal, move)
				}
			}
		} else {
			retVal = allMoves
		}
	}
	return retVal
}

func (chess *Chess) getCastlingMoves(ourColor PieceColor) []Move {
	var retVal []Move

	theirColor := swapColor(ourColor)
	if chess.castling[ourColor]&ksideCastleMove != 0 {
		castlingFrom := chess.kings[ourColor]
		castlingTo := castlingFrom + 2
		if chess.board[castlingFrom+1].isUnspecified() &&
			chess.board[castlingTo].isUnspecified() &&
			!chess.attacked(theirColor, castlingFrom) &&
			!chess.attacked(theirColor, castlingFrom+1) &&
			!chess.attacked(theirColor, castlingTo) {
			retVal = append(retVal, chess.addMove(castlingFrom, castlingTo, ksideCastleMove)...)
		}
	}
	if chess.castling[ourColor]&qsideCastleMove != 0 {
		castlingFrom := chess.kings[ourColor]
		castlingTo := castlingFrom - 2
		if chess.board[castlingFrom-1].isUnspecified() &&
			chess.board[castlingTo].isUnspecified() &&
			chess.board[castlingFrom-3].isUnspecified() &&
			!chess.attacked(theirColor, castlingFrom) &&
			!chess.attacked(theirColor, castlingFrom-1) &&
			!chess.attacked(theirColor, castlingTo) {
			retVal = append(retVal, chess.addMove(castlingFrom, castlingTo, qsideCastleMove)...)
		}
	}
	return retVal
}

func (chess *Chess) getPieceMoves(fromSquare int, currPiece Piece) []Move {
	var retVal []Move

	for _, offset := range pieceOffsets[currPiece.ptype] {
		squareNum := fromSquare
		for {
			squareNum += offset
			if squareNum&0x88 != 0 {
				break
			}
			if chess.board[squareNum].isUnspecified() {
				retVal = append(retVal, chess.addMove(fromSquare, squareNum, normalMove)...)
			} else {
				if chess.board[squareNum].pcolor != currPiece.pcolor {
					retVal = append(retVal, chess.addMove(fromSquare, squareNum, captureMove)...)
				}
				break
			}
			// King and knights only go to a single square
			if currPiece.ptype == knight || currPiece.ptype == king {
				break
			}
		}
	}
	return retVal
}

func (chess *Chess) getPawnAttacks(fromSquare int, ourColor PieceColor) []Move {
	var retVal []Move

	squareNum := fromSquare + pawnOffsets[ourColor][0]
	for offsetIndex := 2; offsetIndex < 4; offsetIndex++ {
		squareNum = fromSquare + pawnOffsets[ourColor][offsetIndex]
		if squareNum&0x88 != 0 {
			continue
		}
		if chess.board[squareNum].pcolor == swapColor(ourColor) {
			retVal = append(retVal, chess.addMove(fromSquare, squareNum, captureMove)...)
		} else if squareNum == chess.enpassantSquare {
			retVal = append(retVal, chess.addMove(fromSquare, squareNum, enpassantMove)...)
		}
	}
	return retVal
}

func (chess *Chess) getPawnMoves(fromSquare int, ourColor PieceColor) []Move {
	var retVal []Move

	squareNum := fromSquare + pawnOffsets[ourColor][0]
	if chess.board[squareNum].isUnspecified() {
		retVal = append(retVal, chess.addMove(fromSquare, squareNum, normalMove)...)
		squareNum = fromSquare + pawnOffsets[ourColor][1]
		if secondRank[ourColor] == rank(fromSquare) && chess.board[squareNum].isUnspecified() {
			retVal = append(retVal, chess.addMove(squareNum, squareNum, bigPawnMove)...)
		}
	}
	return retVal
}

func (chess *Chess) addMove(from int, to int, flags int) []Move {
	var retVal []Move
	// Are we promoting a pawn?
	if chess.board[from].ptype == pawn && (rank(to) == rank8 || rank(to) == rank1) {
		for _, promotionType := range promotionTypes {
			retVal = append(retVal, chess.buildMove(from, to, flags, promotionType))
		}
	} else {
		retVal = append(retVal, chess.buildMove(from, to, flags, 0))
	}
	return retVal
}

func (chess *Chess) attacked(colorAttacking PieceColor, squareNumAttacked int) bool {
	retVal := false

	for cntr := squareNameToID["a8"]; retVal == false && cntr < squareNameToID["h1"]; cntr++ {
		if cntr&0x88 != 0 {
			cntr += 7
			continue
		}
		if chess.board[cntr].isUnspecified() || chess.board[cntr].pcolor != colorAttacking {
			continue
		}
		piece := chess.board[cntr]
		difference := cntr - squareNumAttacked
		index := difference + 119

		if (attacks[index] & (1 << shifts[piece.ptype])) != 0 {
			if piece.ptype == pawn {
				if difference > 0 {
					if piece.pcolor == white {
						retVal = true
						break
					}
				} else {
					if piece.pcolor == black {
						retVal = true
						break
					}
				}
				continue
			}
			if piece.ptype == knight || piece.ptype == king {
				retVal = true
				break
			}

			offset := rays[index]
			j := cntr + offset
			blocked := false
			for j != squareNumAttacked {
				if !chess.board[j].isUnspecified() {
					blocked = true
					break
				}
				j += offset
			}
			if !blocked {
				retVal = true
				break
			}
		}
	}

	return retVal
}

func (chess *Chess) kingAttacked(color PieceColor) bool {
	retVal := chess.attacked(swapColor(color), chess.kings[color])
	return retVal
}

func (chess *Chess) inCheck() bool {
	retVal := chess.kingAttacked(chess.turn)
	return retVal
}

func (chess *Chess) inCheckmate() bool {
	retVal := chess.inCheck() && len(chess.generateMoves(true, "")) == 0
	return retVal
}

func (chess *Chess) inStalemate() bool {
	retVal := !chess.inCheck() && len(chess.generateMoves(true, "")) == 0
	return retVal
}

func (chess *Chess) determineSquareRange(singleSquareName string) (int, int, error) {
	var err error
	firstSquare := emptySquare
	lastSquare := emptySquare

	// Are we generating moves for a single square?
	if singleSquareName != "" {
		if _, ok := squareNameToID[singleSquareName]; ok {
			firstSquare = squareNameToID[singleSquareName]
			lastSquare = firstSquare
		} else {
			err = fmt.Errorf("Invalid square name '%s'", singleSquareName)
		}
	} else {
		firstSquare = squareNameToID["a8"]
		lastSquare = squareNameToID["h1"]
	}
	return firstSquare, lastSquare, err
}

func swapColor(color PieceColor) PieceColor {
	if color == white {
		return black
	}
	return white
}

func rank(square int) int {
	return square >> 4
}

func file(square int) int {
	return square & 15
}

func algebraic(square int) string {
	f := file(square)
	r := rank(square)

	var retVal bytes.Buffer
	retVal.WriteString("abcdefgh"[f : f+1])
	retVal.WriteString("87654321"[r : r+1])
	return retVal.String()
}
