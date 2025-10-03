package main

import (
	"fmt"
	"log"

	"github.com/digvijay-tech/interactive_inputs"
)

func main() {
	seekers := []string{
		"dante",
		"loK",
		"sophie",
		"zhalia",
		"cherit",
	}

	options := &interactive_inputs.CheckboxOptions{
		Title:         "Pick your seekers:",
		Description:   "It's time for a team up.",
		TextTransform: interactive_inputs.CAPITALISE,
		MinSelection:  1,
		MaxSelection:  3,
	}

	team, err := interactive_inputs.Checkbox(seekers, options)
	if err != nil {
		log.Fatalln(err.Error())
	}

	fmt.Println("Your team:")
	for _, v := range team {
		fmt.Printf("- %v\n", v)
	}
}
