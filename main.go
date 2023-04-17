package main

// Import format and time packages
import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

// Do the things
func main() {
	http.HandleFunc("/hello", helloHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

type Greeting struct {
	Message   string    `json:"message"`
	Timestamp time.Time `json:"timestamp"`
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	message := map[string]string{"message": "Automate all of the things!", "timestamp": time.Now().Format(time.RFC3339)}
	response, err := json.MarshalIndent(message, "", "  ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
