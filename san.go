package chess

import (
	"unicode"
)

func (chess *Chess) getDisambigutor(move Move) string {
	retVal := ""

	from := move.from
	to := move.to
	piece := move.ptype

	ambiguities := 0
	sameRank := 0
	sameFile := 0

	moves := chess.Moves(true, "")

	for cntr := range moves {
		ambigFrom := moves[cntr].from
		ambigTo := moves[cntr].to
		ambigPiece := moves[cntr].ptype

		if piece == ambigPiece && from != ambigFrom && to == ambigTo {
			ambiguities++

			if rank(from) == rank(ambigFrom) {
				sameRank++
			}
			if file(from) == file(ambigFrom) {
				sameFile++
			}
		}
	}

	if ambiguities > 0 {
		if sameRank > 0 && sameFile > 0 {
			retVal = algebraic(from)
		} else if sameFile > 0 {
			retVal = algebraic(from)[1:1]
		} else {
			retVal = algebraic(from)[0:1]
		}
	}
	return retVal
}

func (chess *Chess) moveToSAN(move Move) string {
	retVal := ""

	if (move.flags & ksideCastleMove) != 0 {
		retVal = "O-O"
	} else if (move.flags & qsideCastleMove) != 0 {
		retVal = "O-O-O"
	} else {
		if move.ptype != pawn {
			disambig := chess.getDisambigutor(move)
			retVal = string(unicode.ToUpper(rune(move.ptype))) + disambig
		}
		if (move.flags & (captureMove | enpassantMove)) != 0 {
			if move.ptype == pawn {
				retVal += algebraic(move.from)[0:1]
			}
			retVal += "x"
		}
		retVal += algebraic(move.to)
		if (move.flags & promotionMove) != 0 {
			retVal += "=" + string(unicode.ToUpper(rune(move.promotedType)))
		}
	}
	return retVal
}

func isFile(file byte) bool {
	retVal := file >= 'a' && file <= 'h'
	return retVal
}

func isRank(rank byte) bool {
	retVal := rank >= '1' && rank <= '8'
	return retVal
}
