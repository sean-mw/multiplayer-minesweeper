import { useState } from "react";
import { useNavigate } from "react-router";

function Home() {
  const [loading, setLoading] = useState(false);
  const navigate = useNavigate();

  const createRoom = async () => {
    setLoading(true);
    try {
      const res = await fetch("/rooms", {
        method: "POST",
      });
      if (!res.ok) throw new Error("Failed to create room");
      const data = await res.json();
      navigate(`/room/${data.roomId}`);
    } catch (err) {
      console.error("Error creating room:", err);
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="flex flex-col gap-8">
      <div className="font-semibold text-3xl">Multiplayer Minesweeper</div>
      <button
        className="p-3 font-semibold border-3 border-black rounded
        hover:bg-gray-200 hover:cursor-pointer
          disabled:opacity-50 disabled:cursor-not-allowed"
        onClick={createRoom}
        disabled={loading}
      >
        {loading ? "Creating..." : "Create New Room"}
      </button>
    </div>
  );
}

export default Home;
