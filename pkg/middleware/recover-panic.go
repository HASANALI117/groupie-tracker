package middleware

import (
    "log"
    "net/http"
)

func RecoverPanic(next func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
    return func(w http.ResponseWriter, r *http.Request) {
        defer func() {
            if err := recover(); err != nil {
                log.Printf("Internal server error: %v", err)
                http.Error(w, "Internal server error", http.StatusInternalServerError)
            }
        }()
        next(w, r)
    }
}