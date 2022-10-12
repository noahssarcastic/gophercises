package main

import (
	"encoding/json"
	"fmt"
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

type Adventure map[string]StoryArc

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	jsonBytes, err := os.ReadFile("gopher.json")
	check(err)
	var adv Adventure
	err = json.Unmarshal(jsonBytes, &adv)
	check(err)

	for k, v := range adv {
		fmt.Println(k, v)
	}
}
