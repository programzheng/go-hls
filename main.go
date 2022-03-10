package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/programzheng/go-hls/internal/convert"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	go convert.Convert("data/mixkit-stars-in-space-1610.mp4", "data/")

	const dataDir = "data"
	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
	}

	// add a handler for the song files
	http.Handle("/", addHeaders(http.FileServer(http.Dir(dataDir))))
	fmt.Printf("Starting server on %v\n", port)
	log.Printf("Serving %s on HTTP port: %v\n", dataDir, port)

	// serve and log errors
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

// addHeaders will act as middleware to give us CORS support
func addHeaders(h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		h.ServeHTTP(w, r)
	}
}
