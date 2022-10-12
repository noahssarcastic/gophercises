package main

import (
	"encoding/json"
	"fmt"
	"html/template"
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
	defaultArcKey := "intro"
	arcKey := defaultArcKey
	if path := r.URL.Path; path != "/" {
		arcKey = r.URL.Path[1:]
	}

	if arc, ok := h.Story[arcKey]; ok {
		displayArc(w, arc)
	} else {
		panic("Arc not found!")
	}

}

func displayArc(w http.ResponseWriter, arc StoryArc) {
	createTemplate(w, arc)
}

func createTemplate(w http.ResponseWriter, arc StoryArc) {
	tmpl, err := template.New("arc.html").ParseFiles("arc.html")
	check(err)
	err = tmpl.Execute(w, arc)
	check(err)
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
