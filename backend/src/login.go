package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

func handleLogin(w http.ResponseWriter, r *http.Request) {
    var userData struct {
        Email    string `json:"email"`
        Password string `json:"password"`
    }

    // Decode JSON request body
    if err := json.NewDecoder(r.Body).Decode(&userData); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    collection := client.Database("testdb").Collection("users")
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    // Query filter
    filter := bson.M{"email": userData.Email}
    result := collection.FindOne(ctx, filter)
	err := result.Err()

    if err == mongo.ErrNoDocuments {
        // User not found
        w.WriteHeader(http.StatusConflict) // 409
        json.NewEncoder(w).Encode(map[string]string{
            "message": "No user found!",
        })
        return
    }

    var user User
    if err := result.Decode(&user); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
        json.NewEncoder(w).Encode(map[string]string{
            "message": "Internal server error",
        })
        return
    }

	// authenticating
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userData.Password))
    if err != nil {
        w.WriteHeader(http.StatusUnauthorized)
        json.NewEncoder(w).Encode(map[string]string{
            "message": "Invalid credentials",
        })
        return
    }

    payload := jwt.MapClaims{
		"username": user.Username,
        "email": user.Email,
		"exp":      time.Now().Add(time.Hour * 24).Unix(), // expires in 24h
	}

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
    tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
        json.NewEncoder(w).Encode(map[string]string{
            "message": "Internal server error",
        })
        return
	}
    fmt.Println(tokenString)

    http.SetCookie(w, &http.Cookie{
		Name:     "auth_token",
		Value:    tokenString,
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: true,           // prevents JS access
		Secure:   false,          // set to true in production with HTTPS
		Path:     "/",
	})

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Login successfully",
	})
}