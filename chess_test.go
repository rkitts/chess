package chess

import (
	"fmt"
	"testing"
)

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
	state.white |= ksideCastle
	state.white |= qsideCastle
	state.black |= ksideCastle
	state.black |= qsideCastle
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
