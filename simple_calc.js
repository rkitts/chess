"use strict"

class SimpleCalculator extends MoveCalculator{

    calculateBestMove(game) {
		let newGameMoves = game.ugly_moves();
		let bestMove = null;
		//use any negative large number
		let bestValue = -9999;

		for (let i = 0; i < newGameMoves.length; i++) {
			let newGameMove = newGameMoves[i];
			game.ugly_move(newGameMove);

			//take the negative as AI plays as black

			let boardValue = -this.evaluateBoard(game.board())
			game.undo();
			if (boardValue > bestValue) {
				bestValue = boardValue;
				bestMove = newGameMove
			}
		}

		return bestMove;
	}
}
