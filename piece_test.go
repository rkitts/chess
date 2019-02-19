/* 
 * Tests the Piece structure
 */
package chess

import (
	"testing"
)
func TestNewPieceIsUnspecified(t *testing.T) {
	var aPiece Piece
	
	if aPiece.isUnspecified() == false {
		t.Error("piece ", aPiece, " is not unspecified")
	}
}

func TestPieceWithOneFieldSetIsUnspecified(t *testing.T){
	var aPiece Piece

	aPiece.pcolor = king
	if aPiece.isUnspecified() == false {
		t.Error("piece ", aPiece, " is not unspecified")
	}

	var bPiece Piece
	bPiece.ptype = knight
	if bPiece.isUnspecified() == false {
		t.Error("piece ", bPiece, " is not unspecified")
	}
}
