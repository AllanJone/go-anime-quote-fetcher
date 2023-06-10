package handler

import (
	"encoding/json"
	"net/http"
	"os"
)

type QuoteResponse struct {
	ID     int    `json:"_id"`
	Quote  string `json:"quote"`
	Anime  string `json:"anime"`
	Author string `json:"author"`
}

var (
	authToken = os.Getenv("AUTH_TOKEN")
	url       = "https://waifu.it/api/quote"
)

func fetchAndParseQuote() (*QuoteResponse, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", authToken)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var quote QuoteResponse
	if err := json.NewDecoder(resp.Body).Decode(&quote); err != nil {
		return nil, err
	}
	return &quote, nil
}

func Handler(w http.ResponseWriter, r *http.Request) {
	quote, err := fetchAndParseQuote()
	if err != nil {
		http.Error(w, "Failed to fetch and parse quote", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(quote)
}
