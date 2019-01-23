"use strict"

class MiniMax extends MoveCalculator{
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
}