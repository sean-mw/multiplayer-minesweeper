import React from "react";
import Cell, { type CellType } from "./Cell";

type BoardType = {
  width: number;
  height: number;
  cells: CellType[][];
  mines: number;
  revealed: number;
};

type BoardProps = {
  board: BoardType;
  onReveal: (x: number, y: number) => void;
  onFlag: (x: number, y: number) => void;
};

const Board: React.FC<BoardProps> = ({ board, onReveal, onFlag }) => {
  const gridTemplateColumns = `repeat(${board.width}, 32px)`;

  return (
    <div className={"grid justify-center"} style={{ gridTemplateColumns }}>
      {board.cells.map((row, y) =>
        row.map((cell, x) => (
          <Cell
            cell={cell}
            x={x}
            y={y}
            boardWidth={board.width}
            boardHeight={board.height}
            onLeftClick={() => onReveal(x, y)}
            onRightClick={() => onFlag(x, y)}
          />
        ))
      )}
    </div>
  );
};

export default Board;
