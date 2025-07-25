import { cn } from "../lib/utils";

export type CellType = {
  isMine: boolean;
  isRevealed: boolean;
  isFlagged: boolean;
  adjacent: number;
};

type CellProps = {
  cell: CellType;
  x: number;
  y: number;
  boardWidth: number;
  boardHeight: number;
  onLeftClick: () => void;
  onRightClick: () => void;
};

const Cell: React.FC<CellProps> = ({
  cell,
  x,
  y,
  boardWidth,
  boardHeight,
  onLeftClick,
  onRightClick,
}) => {
  const backgroundColor = cell.isRevealed ? "bg-gray-300" : "bg-gray-500";
  let borderStyle = "border-t border-l border-black " + backgroundColor;
  if (x === boardWidth - 1) borderStyle += " border-r";
  if (y === boardHeight - 1) borderStyle += " border-b";

  let display = "";
  if (cell.isFlagged) display = "🚩";
  else if (!cell.isRevealed) display = "";
  else if (cell.isMine) display = "💣";
  else if (cell.adjacent > 0) display = cell.adjacent.toString();

  return (
    <button
      className={cn("w-8 h-8", borderStyle, backgroundColor)}
      onClick={onLeftClick}
      onContextMenu={(e) => {
        e.preventDefault();
        onRightClick();
      }}
    >
      {display}
    </button>
  );
};

export default Cell;
