package chess

import (
	"fmt"
	"testing"
)

func TestMovingPawnReturnsCorrectMoves(t *testing.T) {
	chess := New()
	chess.Clear()
	chess.board[squareNameToID["a2"]] = Piece{pcolor: white, ptype: pawn}
	actualMoves := chess.generateMoves(true, "a2")
	if len(actualMoves) != 2 {
		t.Errorf("Expected 2 moves, got %d", len(actualMoves))
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
	var state castlingState
	state.white |= ksideCastleMove
	state.white |= qsideCastleMove
	state.black |= ksideCastleMove
	state.black |= qsideCastleMove
	actual := generateCastlingFEN(state)
	if actual != "KQkq" {
		t.Errorf("Expected 'KQkq' got '%s'", actual)
	}

	state.white = 0
	state.black = 0
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
