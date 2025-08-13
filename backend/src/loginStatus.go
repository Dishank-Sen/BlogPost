package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
)

func handleLoginStatus(w http.ResponseWriter, r *http.Request){
	cookie, err := r.Cookie("auth_token")
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized) // 401
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Not Logged In",
		})
		return
	}

	token, err := jwt.Parse(cookie.Value, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtSecret, nil
	})

	if err != nil || !token.Valid {
		w.WriteHeader(http.StatusUnauthorized)
        json.NewEncoder(w).Encode(map[string]string{
            "message": "Not Logged In",
        })
        return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Logged In",
	})
}