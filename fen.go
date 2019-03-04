package chess

import (
	"fmt"
	"regexp"
	"strconv"
	"unicode"
)

// Fen deals with FEN encoded strings. See https://en.wikipedia.org/wiki/Forsyth%E2%80%93Edwards_Notation
type Fen struct {
	piecePlacement   string
	activeColor      PieceColor
	castlingAbility  castlingState
	enpassantCapture string
	halfMoves        int
	fullMoves        int
}

func parseFEN(fen string) (*Fen, error) {
	var retVal = new(Fen)
	var err error

	re := regexp.MustCompile("\\s")
	split := re.Split(fen, -1)
	if len(split) != 6 {
		err = fmt.Errorf("Invalid FEN, expected 6 fields, got %d", len(split))
	} else {
		retVal.piecePlacement, err = parsePositions(split[0])
		retVal.activeColor = PieceColor(split[1][0])
		if retVal.activeColor != white && retVal.activeColor != black {
			err = fmt.Errorf("Unrecognized active color '%c'", rune(retVal.activeColor))
		}
		if err == nil {
			retVal.castlingAbility, _ = parseCastling(split[2])
			retVal.enpassantCapture = split[3]
			if retVal.enpassantCapture != "-" {
				if _, ok := squareNameToID[retVal.enpassantCapture]; !ok {
					err = fmt.Errorf("Enpassant target contains invalid square '%s'", retVal.enpassantCapture)
				}
			}
			if err == nil {
				retVal.halfMoves, err = strconv.Atoi(split[4])
				if err == nil {
					retVal.fullMoves, err = strconv.Atoi(split[5])
				}
			}
		}
	}
	return retVal, err
}

func parsePositions(positions string) (string, error) {
	var err error
	re := regexp.MustCompile("/")
	ranks := re.Split(positions, -1)

	if len(ranks) == 8 {
		var pieceToCount = make(map[byte]int)
		for cntr := 0; err == nil && cntr < len(ranks); cntr++ {
			err = parseRank(ranks[cntr], pieceToCount)
		}
		if err == nil {
			err = validatePieceCounts(pieceToCount)
		}
	} else {
		err = fmt.Errorf("Positions contains incorrect number of ranks %d", len(ranks))
	}
	return positions, err
}

func validatePieceCounts(pieceToCount map[byte]int) error {
	var err error
	var pieceName string
	var maxCount int
	for key, value := range pieceToCount {
		switch rune(key) {
		case 'r':
			pieceName = "Black rooks"
			maxCount = 2
		case 'n':
			pieceName = "Black knights"
			maxCount = 2
		case 'b':
			pieceName = "Black bishops"
			maxCount = 2
		case 'q':
			pieceName = "Black queen"
			maxCount = 1
		case 'k':
			pieceName = "Black king"
			maxCount = 1
		case 'p':
			pieceName = "Black pawns"
			maxCount = 8
		case 'R':
			pieceName = "White rooks"
			maxCount = 2
		case 'N':
			pieceName = "White knights"
			maxCount = 2
		case 'B':
			pieceName = "White bishops"
			maxCount = 2
		case 'Q':
			pieceName = "White queen"
			maxCount = 1
		case 'K':
			pieceName = "White king"
			maxCount = 1
		case 'P':
			pieceName = "White pawns"
			maxCount = 8
		}
		if value > maxCount {
			err = fmt.Errorf("%s has %d pieces", pieceName, value)
			break
		}
	}
	return err
}

func parseRank(rank string, pieceToCount map[byte]int) error {
	var squareCount = 0
	var err error

	if len(rank) > 8 {
		err = fmt.Errorf("rank '%s' is too long", rank)
	} else {
		for cntr := 0; err == nil && cntr < len(rank); cntr++ {
			var currChar = rank[cntr]
			if unicode.IsDigit(rune(currChar)) {
				squareCount += int(currChar - '0')
			} else {
				if err = parsePieceInRank(currChar); err == nil {
					pieceToCount[currChar] = pieceToCount[currChar] + 1
					squareCount++
				}
			}
		}
	}

	if err == nil {
		if squareCount > 8 {
			err = fmt.Errorf("rank has more than 8 squares")
		} else if squareCount < 8 {
			err = fmt.Errorf("rank has fewer than 8 squares")
		}
	}

	return err
}

func parsePieceInRank(piece byte) error {
	var retVal error
	switch unicode.ToLower(rune(piece)) {
	case 'r':
	case 'n':
	case 'b':
	case 'q':
	case 'k':
	case 'p':
	default:
		retVal = fmt.Errorf("unrecognized character %c", rune(piece))
	}
	return retVal
}

func parseCastling(fenCastling string) (castlingState, error) {
	var retVal castlingState
	var err error
	for cntr := 0; cntr < len(fenCastling) && err == nil; cntr++ {
		switch fenCastling[cntr] {
		case 'K':
			retVal.white |= ksideCastle
		case 'Q':
			retVal.white |= qsideCastle
		case 'k':
			retVal.black |= ksideCastle
		case 'q':
			retVal.black |= qsideCastle
		case '-':
			break
		default:
			err = fmt.Errorf("unexpected character '%c' in FEN castling encoding", fenCastling[cntr])
		}
	}
	return retVal, err
}
