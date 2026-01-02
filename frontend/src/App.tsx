import { useEffect, useState } from "react";
import "./App.css";

function App() {
  const [message, setMessage] = useState("");

  useEffect(() => {
    fetch("http://localhost:8080")
      .then((res) => res.json())
      .then((data) => {
        console.log(data);
        setMessage(data.message)
      })
      .catch((err) => console.error(err));
  }, []);

  return (
    <div>
      <h1>Frontend: React</h1>
      <p>Backend says: {message}</p>
    </div>
  );
}

export default App;
