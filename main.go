package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"

	"github.com/redraskal/r6-dissect/dissect"
)

func handleParse(w http.ResponseWriter, r *http.Request) {
	buf, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	dr, err := dissect.NewReader(bytes.NewReader(buf))
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	if err := dr.Read(); !dissect.Ok(err) {
		http.Error(w, err.Error(), 500)
		return
	}

	out := map[string]interface{}{
		"header":        dr.Header,
		"matchFeedback": dr.MatchFeedback,
		"scoreboard":    dr.Scoreboard,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(out)
}

func main() {
	http.HandleFunc("/parse", handleParse)
	port := os.Getenv("PORT")
	if port == "" {
		port = "10000"
	}
	http.ListenAndServe(":"+port, nil)
}
