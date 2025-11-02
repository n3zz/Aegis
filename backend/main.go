package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
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
		http.Error(w, "Module not found", http.StatusNotFound)
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

func modulesHandler(w http.ResponseWriter, r *http.Request) {
	// Empty basket for all the modules
	var modules []Module

	// Read all entries within the "modules" directory
	entries, err := os.ReadDir("modules")
	if err != nil {
		http.Error(w, "Failed to read modules directory", http.StatusInternalServerError)
		return
	}

	// Loop through each entry within all the modules within the directory
	for _, entry := range entries {

		if entry.IsDir() {
			continue
		}

		if !strings.HasSuffix(entry.Name(), ".json") {
			continue
		}

		// Read the file contents
		filename := "modules/" + entry.Name()
		data, err := os.ReadFile(filename)
		if err != nil {
			fmt.Printf("Failed to read file %s: %v\n", filename, err)
			continue
		}

		var mod Module
		if err := json.Unmarshal(data, &mod); err != nil {
			fmt.Printf("failed to verify file %s: %v\n", filename, err)
			continue
		}

		modules = append(modules, mod)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(modules)

}

func main() {
	http.HandleFunc("/module/", moduleHandler)
	http.HandleFunc("/modules/", modulesHandler)
	fmt.Println("Server running on http://localhost:3000")
	http.ListenAndServe(":3000", nil)
}
