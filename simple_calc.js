"use strict"

class SimpleCalculator{

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

	evaluateBoard(board) {
		let totalEvaluation = 0;
		for (let i = 0; i < 8; i++) {
			for (let j = 0; j < 8; j++) {
				totalEvaluation = totalEvaluation + this.getPieceValue(board[i][j]);
			}
		}
		return totalEvaluation;
	}

	getPieceValue(piece) {
		if (piece === null) {
			return 0;
		}
		let getAbsoluteValue = function (piece) {
			if (piece.type === 'p') {
				return 10;
			} else if (piece.type === 'r') {
				return 50;
			} else if (piece.type === 'n') {
				return 30;
			} else if (piece.type === 'b') {
				return 30 ;
			} else if (piece.type === 'q') {
				return 90;
			} else if (piece.type === 'k') {
				return 900;
			}
			throw "Unknown piece type: " + piece.type;
		};
	
		let absoluteValue = getAbsoluteValue(piece, piece.color === 'w');
		return piece.color === 'w' ? absoluteValue : -absoluteValue;
	}
}
