export type GameStatus = "Pending" | "Won" | "Lost";

export const IntToGameStatus: Record<number, GameStatus> = {
  0: "Pending",
  1: "Won",
  2: "Lost",
};
