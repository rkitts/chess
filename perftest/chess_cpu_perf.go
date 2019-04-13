package main

import (
	"fmt"

	"github.com/pkg/profile"
	"github.com/rkitts/chess"
)

func main() {
	defer profile.Start().Stop()
	moves := []string{
		"e4", "Nc6", "d4", "Nf6", "Bc4", "a6", "Qh5", "b6", "Qxf7"}
	chessBoard := chess.New()
	for cntr := 1; cntr < 10; cntr++ {
		fmt.Printf("Loop %d", cntr)
		chessBoard.Reset()
		for moveNum := range moves {
			err := chessBoard.Move(moves[moveNum])
			if err != nil {
				fmt.Printf("Error %v processing move %s", err, moves[moveNum])
			}
		}
		if !chessBoard.InCheckmate() {
			fmt.Printf("Expected to be in checkmate!")
		}
		fmt.Printf("%s", chessBoard.GenerateFen())
	}
}
