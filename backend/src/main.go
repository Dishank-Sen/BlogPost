package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
    Username string `json:"username"`
    Email    string `json:"email"`
    Password string `json:"password"`
}

var client *mongo.Client
var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

func main(){
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("MONGO_URI not set in .env file")
	}

	client = connectDB(uri)
	port := ":3000"
	http.HandleFunc("/api/signup", withCORS(handleSignup))
	http.HandleFunc("/api/login", withCORS(handleLogin))
	http.HandleFunc("/api/loginStatus", withCORS(handleLoginStatus))
	fmt.Println("Server running on", port)
	http.ListenAndServe(port,nil)
}