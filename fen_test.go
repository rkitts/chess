package chess

import (
	"testing"
)

func TestParsingDefaultPositionWorks(t *testing.T) {
	_, err := parseFEN(defaultPosition)
	if err != nil {
		t.Errorf("Error parsing FEN")
	}
}

func TestParsePieceInRankErrsIfUnrecognizedChar(t *testing.T) {
	if err := parsePieceInRank('z'); err == nil {
		t.Errorf("Error not set")
	}
}

func TestParseRankErrsIfLineTooLong(t *testing.T) {
	var ignored = make(map[byte]int)
	if err := parseRank("ppppppppp", ignored); err == nil {
		t.Errorf("Error not set")
	}
}

func TestValidatePieceCounts(t *testing.T) {
	assertValidatePieceCountsReportsError(t, 'r', 3)
	assertValidatePieceCountsReportsError(t, 'n', 3)
	assertValidatePieceCountsReportsError(t, 'b', 3)
	assertValidatePieceCountsReportsError(t, 'q', 2)
	assertValidatePieceCountsReportsError(t, 'k', 2)

	assertValidatePieceCountsReportsError(t, 'R', 3)
	assertValidatePieceCountsReportsError(t, 'N', 3)
	assertValidatePieceCountsReportsError(t, 'B', 3)
	assertValidatePieceCountsReportsError(t, 'Q', 2)
	assertValidatePieceCountsReportsError(t, 'K', 2)
}

func assertValidatePieceCountsReportsError(t *testing.T, piece byte, count int) {
	var pieceCount = make(map[byte]int)
	pieceCount[piece] = count
	if err := validatePieceCounts(pieceCount); err == nil {
		t.Errorf("%c did not report error for %d pieces", piece, count)
	}
}

func TestParseRankReportsErrorIfSquareCountExceeded(t *testing.T) {
	var pieceToCount = make(map[byte]int)
	if err := parseRank("9", pieceToCount); err == nil {
		t.Errorf("Error not set")
	}

	if err := parseRank("333", pieceToCount); err == nil {
		t.Errorf("Error not set")
	}
}

func TestParseCastlingWorks(t *testing.T) {

	result, err := parseCastling("KQkq")
	if err != nil {
		t.Errorf("error set %s", err)
	}
	if (result[white] & ksideCastleMove) == 0 {
		t.Error("white ksideCastle not set")
	}
	if (result[white] & qsideCastleMove) == 0 {
		t.Error("white qsideCastle not set")
	}
	if (result[black] & ksideCastleMove) == 0 {
		t.Error("black ksideCastle not set")
	}
	if (result[black] & qsideCastleMove) == 0 {
		t.Error("black qsideCastle not set")
	}
}

func TestParseCastlingErrorHandling(t *testing.T) {

	if _, err := parseCastling("KQzq"); err == nil {
		t.Error("No error reported")
	}
}
