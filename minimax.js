"use strict"

class MiniMax{
    calculateBestMove(game){
        let newGameMoves = game.ugly_moves();
        let bestEvaluation = -Infinity;
        let bestMove;

        for (let i = 0; i < newGameMoves.length; i++) {
            let newGameMove = newGameMoves[i];
            game.ugly_move(newGameMove);
            let currentEvaluation = this.minimax(game, 1, true);
            game.undo();
            if(currentEvaluation > bestEvaluation){
                bestMove = newGameMove;
                bestEvaluation = currentEvaluation;
            }
        }
        return(bestMove);
    }

    minimax(game, depth, isMaximising){
        let retVal;

        if(depth == 0){
            retVal = -this.evaluateBoard(game.board());
        }
        else if(isMaximising){
            let newMoves = game.ugly_moves();
            let bestEvaluation = -Infinity;
            for(let cntr = 0; cntr < newMoves.length; cntr++){
                game.ugly_move(newMoves[cntr]);
                bestEvaluation = Math.max(bestEvaluation, this.minimax(game, depth - 1, false));
                game.undo();
            }
            retVal = bestEvaluation;
        }
        else{
            let newMoves = game.ugly_moves();
            let bestEvaluation = Infinity;
            for(let cntr = 0; cntr < newMoves.length; cntr++){
                game.ugly_move(newMoves[cntr]);
                bestEvaluation = Math.min(bestEvaluation, this.minimax(game, depth - 1, true));
                game.undo();
            }
            retVal = bestEvaluation;
        }
        return(retVal);
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