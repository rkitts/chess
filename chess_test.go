package chess

import (
	"fmt"
	"testing"
)

func TestCastlingHonorsAttackedSquares(t *testing.T) {
	chess := New()
	chess.Clear()

	chess.board[squareNameToID["e1"]] = Piece{pcolor: white, ptype: king}
	chess.castling[white] |= (ksideCastleMove | qsideCastleMove)
	chess.kings[white] = squareNameToID["e1"]

	chess.board[squareNameToID["f2"]] = Piece{pcolor: black, ptype: rook}
	actualMoves := chess.getCastlingMoves(white)
	if len(actualMoves) != 1 {
		t.Errorf("Expected 1 moves, got %d", len(actualMoves))
	}
}
func TestCastlingMoves(t *testing.T) {
	chess := New()
	chess.Clear()
	chess.board[squareNameToID["e1"]] = Piece{pcolor: white, ptype: king}
	chess.castling[white] = (ksideCastleMove | qsideCastleMove)
	chess.kings[white] = squareNameToID["e1"]
	actualMoves := chess.getCastlingMoves(white)
	if len(actualMoves) != 2 {
		t.Errorf("Expected 2 moves, got %d", len(actualMoves))
	}

	chess.castling[white] = ksideCastleMove
	actualMoves = chess.getCastlingMoves(white)
	if len(actualMoves) != 1 {
		t.Errorf("Expected 1 moves, got %d", len(actualMoves))
	}

	chess.castling[white] = qsideCastleMove
	actualMoves = chess.getCastlingMoves(white)
	if len(actualMoves) != 1 {
		t.Errorf("Expected 1 moves, got %d", len(actualMoves))
	}
}

func TestAttacksWithRook(t *testing.T) {
	// This just makes sure the code dealing with sliders works
	chess := New()
	chess.Clear()
	chess.board[squareNameToID["a1"]] = Piece{pcolor: white, ptype: rook}
	chess.board[squareNameToID["a8"]] = Piece{pcolor: black, ptype: rook}
	actual := chess.attacked(white, squareNameToID["a8"])
	if actual != true {
		t.Errorf("Expected true, got %v", actual)
	}
}

func TestAttackedForPawns(t *testing.T) {
	chess := New()
	chess.Clear()
	chess.board[squareNameToID["a2"]] = Piece{pcolor: white, ptype: pawn}
	chess.board[squareNameToID["b3"]] = Piece{pcolor: black, ptype: pawn}
	actual := chess.attacked(white, squareNameToID["b3"])
	if actual != true {
		t.Errorf("Expected true, got %v", actual)
	}
}

func TestKingAndKnightMoveOnlyEightMoves(t *testing.T) {
	chess := New()
	chess.Clear()
	actualMoves := chess.getPieceMoves(squareNameToID["e5"], Piece{pcolor: white, ptype: knight})
	if len(actualMoves) != 8 {
		t.Errorf("Expected 8 moves, got %d", len(actualMoves))
	}
	actualMoves = chess.getPieceMoves(squareNameToID["e5"], Piece{pcolor: white, ptype: king})
	if len(actualMoves) != 8 {
		t.Errorf("Expected 8 moves, got %d", len(actualMoves))
	}

}

func TestKnightMoveOnlyEightMoves(t *testing.T) {
	chess := New()
	chess.Clear()
	actualMoves := chess.getPieceMoves(squareNameToID["e5"], Piece{pcolor: white, ptype: knight})
	if len(actualMoves) != 8 {
		t.Errorf("Expected 8 moves, got %d", len(actualMoves))
	}
}

func TestPieceMoveObservesPieces(t *testing.T) {
	chess := New()
	chess.Clear()
	chess.board[squareNameToID["a2"]] = Piece{pcolor: white, ptype: rook}
	actualMoves := chess.getPieceMoves(squareNameToID["a1"], Piece{pcolor: white, ptype: rook})
	if len(actualMoves) != 7 {
		t.Errorf("Expected 7 moves, got %d", len(actualMoves))
	}
}

func TestPieceMoveEnds(t *testing.T) {
	chess := New()
	chess.Clear()
	actualMoves := chess.getPieceMoves(squareNameToID["a1"], Piece{pcolor: white, ptype: rook})
	if len(actualMoves) != 14 {
		t.Errorf("Expected 14 moves, got %d", len(actualMoves))
	}
}

func TestPawnAttacksEnpassant(t *testing.T) {
	chess := New()
	chess.Clear()
	chess.enpassantSquare = squareNameToID["d6"]
	chess.board[squareNameToID["b2"]] = Piece{pcolor: white, ptype: pawn}
	actualMoves := chess.getPawnAttacks(squareNameToID["c5"], white)
	if len(actualMoves) != 1 {
		t.Errorf("Expected 1 moves, got %d", len(actualMoves))
	} else {
		if actualMoves[0].flags&enpassantMove == 0 {
			t.Errorf("Expected an enpassantMove")
		}
	}
}

func TestPawnAttacksDiagonals(t *testing.T) {
	chess := New()
	chess.Clear()
	chess.board[squareNameToID["b2"]] = Piece{pcolor: white, ptype: pawn}
	chess.board[squareNameToID["a3"]] = Piece{pcolor: black, ptype: pawn}
	chess.board[squareNameToID["c3"]] = Piece{pcolor: black, ptype: pawn}
	actualMoves := chess.getPawnAttacks(squareNameToID["b2"], white)
	if len(actualMoves) != 2 {
		t.Errorf("Expected 2 moves, got %d", len(actualMoves))
	}
}

func TestMovingBlockedPawnHasNoBigPawnMove(t *testing.T) {
	chess := New()
	chess.Clear()
	chess.board[squareNameToID["a2"]] = Piece{pcolor: white, ptype: pawn}
	chess.board[squareNameToID["a4"]] = Piece{pcolor: white, ptype: pawn}
	actualMoves := chess.getPawnMoves(squareNameToID["a2"], white)
	if len(actualMoves) != 1 {
		t.Errorf("Expected 1 moves, got %d", len(actualMoves))
	}
}

func TestMovingUnblockedPawnReturnsCorrectMoves(t *testing.T) {
	chess := New()
	chess.Clear()
	chess.board[squareNameToID["a2"]] = Piece{pcolor: white, ptype: pawn}
	actualMoves := chess.getPawnMoves(squareNameToID["a2"], white)
	if len(actualMoves) != 2 {
		t.Errorf("Expected 2 moves, got %d", len(actualMoves))
	}

	foundBigPawn := false
	for _, move := range actualMoves {
		if move.flags&bigPawnMove != 0 {
			foundBigPawn = true
			break
		}
	}
	if foundBigPawn == false {
		t.Errorf("Didn't find a bigpawn move")
	}
}

func TestDetermineSquareRangeFailsForInvalidSquare(t *testing.T) {
	chess := New()
	first, last, err := chess.determineSquareRange("b9")
	if err == nil {
		t.Errorf("Error not reported")
	}
	if first != emptySquare || last != emptySquare {
		t.Errorf("Invalid squares returned, %d and %d", first, last)
	}
}

func TestBuildMoveCapturesIfSquareNotEmpty(t *testing.T) {
	chess := New()
	from := squareNameToID["a1"]
	to := squareNameToID["a2"]

	actual := chess.buildMove(from, to, 0, 0)
	if actual.capturedType != pawn {
		t.Errorf("Unexpected captureType '%c'", actual.capturedType)
	}
}

func TestBuildMoveCapturesNothingIfSquareEmpty(t *testing.T) {
	chess := New()
	from := squareNameToID["a2"]
	to := squareNameToID["a3"]

	actual := chess.buildMove(from, to, 0, 0)
	if actual.capturedType != 0 {
		t.Errorf("Unexpected captureType '%c'", actual.capturedType)
	}
}

func TestBuildMoveDoesEnpassant(t *testing.T) {
	chess := New()
	chess.Clear()
	from := squareNameToID["a2"]
	to := squareNameToID["a3"]

	actual := chess.buildMove(from, to, enpassantMove, 0)
	if actual.capturedType != pawn {
		t.Errorf("Expected a pawn, got '%c'", actual.capturedType)
	}
}

func TestRemoveRemoves(t *testing.T) {
	chess := New()
	piece := chess.remove("b2")
	if piece.isUnspecified() {
		t.Errorf("Got an unspecified piece for legal square")
	}
	piece = chess.get("b2")
	if !piece.isUnspecified() {
		t.Errorf("Got an specified piece after removal")
	}
}

func TestGetReturnsSpecifiedPiece(t *testing.T) {
	chess := New()
	piece := chess.get("b2")
	if piece.isUnspecified() {
		t.Errorf("Got an unspecified piece for legal square")
	}
}

func TestGetReturnsUnspecifiedPieceIfBadInput(t *testing.T) {
	chess := New()
	piece := chess.get("b9")
	if piece.isUnspecified() == false {
		t.Errorf("Expected unspecifed piece, got %+v", piece)
	}
}

func TestResetLoadsDefaultPosition(t *testing.T) {
	var retVal Piece
	fmt.Printf("%d", retVal.pcolor)
	chess := New()
	chess.Clear()
	chess.Reset()
	if chess.generateFen() != defaultPosition {
		t.Errorf("Did not get defaultPosition, got '%s'", chess.generateFen())
	}
}

func TestUpdateSetupWithDefaultPosition(t *testing.T) {
	chess := New()
	chess.header["SetUp"] = "Expected"
	chess.header["FEN"] = "Expected"
	chess.updateSetup(defaultPosition)
	if _, ok := chess.header["SetUp"]; ok == true {
		t.Errorf("Header was not cleared, got '%s'", chess.header["SetUp"])
	}
	if _, ok := chess.header["FEN"]; ok == true {
		t.Errorf("Header was not cleared, got '%s'", chess.header["FEN"])
	}
}

func TestUpdateSetupDoesNothingIfHistoryPresent(t *testing.T) {
	chess := New()
	chess.history = append(chess.history, historyEntry{})
	chess.header["SetUp"] = "Expected"
	chess.header["FEN"] = "Expected"
	chess.updateSetup(defaultPosition)
	if chess.header["SetUp"] != "Expected" || chess.header["FEN"] != "Expected" {
		t.Errorf("Header was unexpectedly modified")
	}
}
func TestGenerateEnpassantFEN(t *testing.T) {
	squareID := squareNameToID["a8"]
	actual := generateEnpassantFEN(squareID)
	if actual != "a8" {
		t.Errorf("Expected 'a8' but got '%s'", actual)
	}

	actual = generateEnpassantFEN(emptySquare)
	if actual != "-" {
		t.Errorf("Expected '-' but got '%s'", actual)
	}
}

func TestGenerateCastlingFEN(t *testing.T) {
	var state = make(castlingState)
	state[white] |= ksideCastleMove
	state[white] |= qsideCastleMove
	state[black] |= ksideCastleMove
	state[black] |= qsideCastleMove
	actual := generateCastlingFEN(state)
	if actual != "KQkq" {
		t.Errorf("Expected 'KQkq' got '%s'", actual)
	}

	state[white] = 0
	state[black] = 0
	actual = generateCastlingFEN(state)
	if actual != "-" {
		t.Errorf("Expected '-' got '%s'", actual)
	}
}
func TestClearedBoardGeneratesEmptyFEN(t *testing.T) {
	chess := New()
	chess.Clear()
	expected := "8/8/8/8/8/8/8/8 w - - 0 1"
	actual := chess.generateFen()
	if actual != expected {
		t.Errorf("Expected '%s' got '%s'", expected, actual)
	}
}

func TestGenerateFenWorks(t *testing.T) {
	chess := New()
	actual := chess.generateFen()
	if actual != defaultPosition {
		t.Errorf("Got unexpected '%s', expected '%s'", actual, defaultPosition)
	}
}
func TestPuttingToInvalidSquareIsError(t *testing.T) {
	chess := New()
	piece := Piece{pcolor: white, ptype: pawn}
	if err := chess.put(piece, "z2"); err == nil {
		t.Errorf("Error not returned")
	}
}

func TestPlaceKingsReportsNoErrorIfPlacedOnSameSquare(t *testing.T) {

	chess := New()
	chess.Clear()
	piece := Piece{ptype: king, pcolor: white}

	if err := chess.maybeUpdateKings(piece, 0); err != nil {
		t.Errorf("Unexpected error putting king on board")
	}
	if err := chess.maybeUpdateKings(piece, 0); err != nil {
		t.Errorf("Error not returned")
	}

	piece.pcolor = black

	if err := chess.maybeUpdateKings(piece, 0); err != nil {
		t.Errorf("Unexpected error putting king on board")
	}
	if err := chess.maybeUpdateKings(piece, 0); err != nil {
		t.Errorf("Error not returned")
	}

}

func TestPlaceKingsReportsErrorIfAlreadyPlaced(t *testing.T) {
	chess := New()
	chess.Clear()
	piece := Piece{ptype: king, pcolor: white}

	if err := chess.maybeUpdateKings(piece, 0); err != nil {
		t.Errorf("Unexpected error putting king on board")
	}
	if err := chess.maybeUpdateKings(piece, 1); err == nil {
		t.Errorf("Error when adding king to same square")
	}

	piece.pcolor = black

	if err := chess.maybeUpdateKings(piece, 0); err != nil {
		t.Errorf("Unexpected error putting king on board")
	}
	if err := chess.maybeUpdateKings(piece, 1); err == nil {
		t.Errorf("Error when adding king to same square")
	}

}

func TestAlgebraicWorks(t *testing.T) {
	for rankValue := 0; rankValue <= 112; rankValue += 16 {
		for fileValue := 0; fileValue < 8; fileValue++ {
			squareValue := rankValue + fileValue
			result := algebraic(squareValue)
			if squareNameToID[result] != squareValue {
				t.Errorf("Expected %d, got %d for square %s", squareValue, squareNameToID[result], result)
			}
		}
	}
}
