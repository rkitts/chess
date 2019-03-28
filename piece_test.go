/*
 * Tests the Piece structure
 */
package chess

import (
	"testing"
)

func TestNewPieceIsUnspecified(t *testing.T) {
	var aPiece Piece

	if aPiece.IsUnspecified() == false {
		t.Error("piece ", aPiece, " is not unspecified")
	}
}

func TestPieceWithOneFieldSetIsUnspecified(t *testing.T) {
	var aPiece Piece

	aPiece.pcolor = white
	if aPiece.IsUnspecified() == false {
		t.Error("piece ", aPiece, " is not unspecified")
	}

	var bPiece Piece
	bPiece.ptype = knight
	if bPiece.IsUnspecified() == false {
		t.Error("piece ", bPiece, " is not unspecified")
	}
}
