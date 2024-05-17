package main

import (
	"fmt"
	"github.com/lclpedro/go-ratelimiter/internal"
	"log"
	"net/http"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	store, err := internal.NewRedisStore(os.Getenv("REDIS_ADDR"))
	if err != nil {
		log.Fatal("Erro ao conectar ao Redis:", err)
	}

	limiter := internal.NewRateLimiter(store, getExpirationTime())

	r := mux.NewRouter()
	r.Use(limitMiddleware(limiter))

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Requisição permitida!")
	})

	addr := ":8080"
	fmt.Println("Server running in", addr)
	log.Fatal(http.ListenAndServe(addr, r))
}

func limitMiddleware(limiter *internal.RateLimiter) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			identifier := getIdentifier(r)
			limiter.SetRate(getRateLimit(identifier))
			fmt.Println("Limiter:", limiter)
			if !limiter.Allow(identifier) {
				http.Error(w, "Limite de requisições excedido", http.StatusTooManyRequests)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

func getIdentifier(r *http.Request) string {
	token := r.Header.Get("X-API-KEY")
	if token != "" {
		return token
	}
	remoteAddr := strings.Split(r.RemoteAddr, ":")[0]
	return remoteAddr
}

func getRateLimit(identifier string) int64 {

	tokensPermitted := os.Getenv("TOKENS_PERMITED")
	tokens := strings.Split(tokensPermitted, "|")
	if ok := slices.Contains(tokens, identifier); ok {
		fmt.Println("Token Definition!")
		rate, _ := strconv.Atoi(os.Getenv("TOKEN_RATE_LIMIT"))
		return int64(rate)
	}
	fmt.Println("IP Definition!")
	rate, _ := strconv.Atoi(os.Getenv("IP_RATE_LIMIT"))
	return int64(rate)
}

func getExpirationTime() time.Duration {
	expirationTime, _ := strconv.Atoi(os.Getenv("EXPIRATION_TIME"))
	return time.Duration(expirationTime) * time.Second
}
