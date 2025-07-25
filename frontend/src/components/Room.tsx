import { useEffect, useRef, useState } from "react";
import { useParams } from "react-router";
import Board from "./Board";
import { IntToGameStatus, type GameStatus } from "../types/game";

type ClientMessage =
  | { action: "reveal"; x: number; y: number }
  | { action: "flag"; x: number; y: number };

function Room() {
  const { roomId } = useParams<{ roomId: string }>();
  const [board, setBoard] = useState(null);
  const [status, setStatus] = useState<GameStatus>("Pending");
  const ws = useRef<WebSocket | null>(null);

  useEffect(() => {
    if (!roomId) return;

    ws.current = new WebSocket(`ws://localhost:8080/rooms/${roomId}/ws`);

    ws.current.onmessage = (event) => {
      const msg = JSON.parse(event.data);
      setBoard(msg.board);
      setStatus(IntToGameStatus[msg.status]);
    };

    return () => ws.current?.close();
  }, [roomId]);

  const sendAction = (action: ClientMessage) => {
    if (ws.current && ws.current.readyState === WebSocket.OPEN) {
      ws.current.send(JSON.stringify(action));
    }
  };

  if (!board) return <div>Loading game...</div>;

  return (
    <div>
      <div>
        Room: {roomId} Status: {status}
      </div>
      <Board
        board={board}
        onReveal={(x, y) => sendAction({ action: "reveal", x, y })}
        onFlag={(x, y) => sendAction({ action: "flag", x, y })}
      />
    </div>
  );
}

export default Room;
