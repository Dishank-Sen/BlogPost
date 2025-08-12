package main

import (
	"context"
	"encoding/json"
	"net/http"
	"time"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

func handleSignup(w http.ResponseWriter, r *http.Request) {
    var user User
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    collection := client.Database("testdb").Collection("users")
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    // Check if account exists
    filter := bson.M{"email": user.Email}
    result := collection.FindOne(ctx, filter)
    err := result.Err()

    if err == nil {
        // User found
        w.WriteHeader(http.StatusConflict) // 409
        json.NewEncoder(w).Encode(map[string]string{
            "message": "Email already registered!",
        })
        return
    } else if err != mongo.ErrNoDocuments {
        // Some DB error
        w.WriteHeader(http.StatusInternalServerError)
        json.NewEncoder(w).Encode(map[string]string{
            "message": "Internal server error",
        })
        return
    }

	// Hash password
	fmt.Println("hashing...")
    hashPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if err != nil {
		// panic(err)
		w.WriteHeader(http.StatusInternalServerError)
        json.NewEncoder(w).Encode(map[string]string{
            "message": "Failed to save user",
        })
        return
    }
	user.Password = string(hashPassword)
    fmt.Println("Hashed password:", user.Password)

    // Insert new user
    if _, err := collection.InsertOne(ctx, user); err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        json.NewEncoder(w).Encode(map[string]string{
            "message": "Failed to save user",
        })
        return
    }

    w.WriteHeader(http.StatusCreated) // 201
    json.NewEncoder(w).Encode(map[string]string{
        "message": "User registered successfully",
    })
}
