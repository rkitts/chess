package chess

import "testing"

func TestMoveToSan(t *testing.T) {
	chess := New()

	var move Move
	move.from = squareNameToID["e2"]
	move.to = squareNameToID["e4"]
	move.ptype = pawn
	assertSANCorrect(chess, t, "e4", move)

	move.from = squareNameToID["e1"]
	move.to = squareNameToID["e8"]
	move.ptype = queen
	assertSANCorrect(chess, t, "Qe8", move)

	move.from = squareNameToID["h1"]
	move.to = squareNameToID["a8"]
	move.ptype = bishop
	move.capturedType = rook
	move.flags = captureMove
	assertSANCorrect(chess, t, "Bxa8", move)

	move.from = squareNameToID["e5"]
	move.to = squareNameToID["d6"]
	move.ptype = pawn
	move.flags = enpassantMove
	assertSANCorrect(chess, t, "exd6", move)

	chess.Load("4k3/8/4r3/8/3N1N2/8/8/4K3 w - - 0 1")
	move.from = squareNameToID["f4"]
	move.to = squareNameToID["e6"]
	move.flags = captureMove
	move.ptype = knight
	move.capturedType = rook
	assertSANCorrect(chess, t, "Nfxe6", move)
}

func assertSANCorrect(chess *Chess, t *testing.T, expected string, move Move) {
	san := chess.moveToSAN(move)
	if san != expected {
		t.Errorf("Expected %s, got %s", expected, san)
	}
}

func TestIsFile(t *testing.T) {
	if isFile('z') != false {
		t.Errorf("Expected z to be false")
	}

	if isFile('h') == false {
		t.Errorf("Expected z to be true")
	}
}

func TestIsRank(t *testing.T) {
	if isRank('9') != false {
		t.Errorf("Expected 9 to be false")
	}

	if isRank('8') == false {
		t.Errorf("Expected 8 to be true")
	}
}
