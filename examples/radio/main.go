package main

import (
	"fmt"
	"log"

	"github.com/digvijay-tech/interactive_inputs"
)

type hero struct {
	name  string
	titan string
	spell string
}

func main() {
	huntik := make(map[string]hero, 0)

	huntik["dante"] = hero{"Dante", "Caliban", "Dragon Feast"}
	huntik["lok"] = hero{"Lok", "Keeperin", "Featherdrop"}
	huntik["sophie"] = hero{"Sophie", "Sebrial", "Hyperstrike"}
	huntik["zhalia"] = hero{"Zhalia", "Garion", "Simple Mind"}
	huntik["cherit"] = hero{"Cherit", "he is a titan ;)", "Soul Burn"}

	// seekers := []string{
	// 	"dante",
	// 	"loK",
	// 	"sophie",
	// 	"zhalia",
	// 	"cherit",
	// }

	seekers := []int{
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13,
	}

	options := &interactive_inputs.RadioOptions{
		Title:         "Pick your seeker:",
		Description:   "Selected seeker comes with a titan and a spell.",
		TextTransform: interactive_inputs.CAPITALISE,
	}

	selectedSeeker, err := interactive_inputs.Radio(seekers, options)
	if err != nil {
		log.Fatalln(err.Error())
	}

	fmt.Println("You Chose:", selectedSeeker)

	// seeker := huntik[strings.ToLower(selectedSeeker)]
	// fmt.Printf("\nName: %s\nFirst Titan: %s\nSpell: %s\n", seeker.name, seeker.titan, seeker.spell)
}
