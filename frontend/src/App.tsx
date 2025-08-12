import { Routes, Route } from "react-router-dom"
import Signup from "./components/Signup"
import HomePage from "./components/HomePage"
import Login from "./components/Login"
import { useEffect } from "react"

function App() {
  // Load saved theme from localStorage
  useEffect(() => {
  const savedTheme = localStorage.getItem("theme");
  if (savedTheme === "dark") {
      document.documentElement.classList.add("dark");
  }
  }, []);

  return (
    <Routes>
      <Route path="/" element={<HomePage/>}/>
      <Route path="/signup" element={<Signup/>}/>
      <Route path="/login" element={<Login/>}/>
    </Routes>
  )
}

export default App
