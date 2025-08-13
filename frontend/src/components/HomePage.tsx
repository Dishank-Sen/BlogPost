import { useEffect, useState } from "react";
import Navbar from "./NavBar"

export default function HomePage(){
    const[loggedIn, setloggedIn] = useState<boolean>(false)

    const isLoggedIn =  async() => {
        try {
            const res = await fetch("http://localhost:3000/api/loginStatus", {
                method: "GET",
                credentials: "include"
            });

            if (res.ok) {
                setloggedIn(true)
            } else {
                setloggedIn(false)
            }
        } catch (error) {
            console.error("Error logging in:", error);
            setloggedIn(false)
        }
    }

    // Load saved theme from localStorage
    useEffect(() => {
    const savedTheme = localStorage.getItem("theme");
    if (savedTheme === "dark") {
        document.documentElement.classList.add("dark");
    }
    isLoggedIn();
    }, []);

    useEffect(() => {
        console.log(loggedIn ? "logged in" : "logged out");
    }, [loggedIn]);
    return(
        <>
            <Navbar/>
        </>
    )
}