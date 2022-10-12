package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type StoryArc struct {
	Title   string
	Story   []string
	Options []struct {
		Text string
		Arc  string
	}
}

type Story map[string]StoryArc

func check(err error) {
	if err != nil {
		panic(err)
	}
}

type AdventureHandler struct {
	Story Story
}

func (h AdventureHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	if path == "/" {
		arc := h.Story["intro"]
		fmt.Fprintln(w, arc)
	} else {
		arc := h.Story[path[1:]]
		fmt.Fprintln(w, arc)
	}
}

func main() {
	jsonBytes, err := os.ReadFile("gopher.json")
	check(err)

	var s Story
	err = json.Unmarshal(jsonBytes, &s)
	check(err)

	handler := AdventureHandler{s}
	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", handler)

}
