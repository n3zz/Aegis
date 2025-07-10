package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Module struct {
	Title          string `json:"title"`
	Slug           string `json:"slug"`
	Summary        string `json:"summary"`
	Difficulty     string `json:"difficulty"`
	Task           string `json:"task"`
	ExpectedAnswer string `json:"expectedAnswer"`
}

func moduleHandler(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/module/")
	filename := fmt.Sprintf("modules/%s.json", path)

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		http.Error(w, "Module no found", http.StatusNotFound)
		return
	}

	var mod Module
	if err := json.Unmarshal(data, &mod); err != nil {
		http.Error(w, "Invalid module format", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(mod)
}

func main() {
	http.HandleFunc("/module/", moduleHandler)
	fmt.Println("Server running on http://localhost:3000")
	http.ListenAndServe(":3000", nil)
}
