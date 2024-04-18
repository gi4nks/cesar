package internal

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

// convertHandler handles HTTP requests to convert integer ranges to Roman numerals.
// It parses the start and end values from the request path parameters,
// converts the range to Roman numerals using the internal package,
// and returns the converted data in JSON format.
func (c *Handler) ConvertHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	start, err := strconv.Atoi(vars["start"])
	if err != nil {
		http.Error(w, "Invalid start value", http.StatusBadRequest)
		return
	}

	endStr, exists := vars["end"]
	var end int
	if !exists {
		end = start // Set end equal to start if end parameter is missing
	} else {
		end, err = strconv.Atoi(endStr)
		if err != nil {
			http.Error(w, "Invalid end value", http.StatusBadRequest)
			return
		}
	}

	// Create a new Converter instance for the given range.
	converter, err := NewConverter(start, end)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Convert the range to Roman numerals and encode the response data as JSON.
	responseData := converter.Convert()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responseData)
}

// startServer starts the HTTP server to handle requests for converting integer ranges to Roman numerals.
// It registers the convertHandler function to handle GET requests at the /convert/{start}/{end} and /convert/{start} endpoints.
func (c *Handler) StartServer() {
	r := mux.NewRouter()
	r.HandleFunc("/convert/{start}/{end}", c.ConvertHandler).Methods("GET")
	r.HandleFunc("/convert/{start}", c.ConvertHandler).Methods("GET")

	fmt.Println("Server is running on port 8080...")
	http.ListenAndServe(":8080", r)
}
