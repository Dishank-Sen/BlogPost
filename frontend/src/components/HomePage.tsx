import { useEffect } from "react";
import Navbar from "./NavBar"

export default function HomePage(){
    // Load saved theme from localStorage
    useEffect(() => {
    const savedTheme = localStorage.getItem("theme");
    if (savedTheme === "dark") {
        document.documentElement.classList.add("dark");
    }
    }, []);
    return(
        <>
            <Navbar/>
        </>
    )
}