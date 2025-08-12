package main

import (
	"fmt"
	"net/http"
)

func withCORS(h http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.Method)
        // Add CORS headers before doing anything else
        w.Header().Set("Access-Control-Allow-Origin", "*")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
        w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")

        // If it's a preflight request, stop here
        if r.Method == http.MethodOptions {
            w.WriteHeader(http.StatusOK)
            return
        }

        // Call the original handler
        h(w, r)
    }
}