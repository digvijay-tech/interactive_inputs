package selectors

import (
	"fmt"
	"log"
	"strings"

	"github.com/digvijay-tech/interactive_inputs/internal/utilities"
)

func moveCursor[T AcceptedListType](index int, cursorPos int, value T, transformation TextTransform, selector string, selected bool) {
	valueType := utilities.FindType(value, false)
	if valueType == "" {
		log.Fatalln("invalid list type")
	}

	// apply text transformation on string type and transform type T to string
	var transformed string
	if valueType == "string" {
		transformed = utilities.TextTransform(transformation.String(), fmt.Sprintf("%v", value))
	} else {
		transformed = fmt.Sprintf("%v", value)
	}

	if selector == "RADIO" {
		if index == cursorPos {
			fmt.Printf("%s %s\n", RIGHTPOINTER_ICON, transformed)
		} else {
			fmt.Printf("  %s\n", transformed)
		}
	}

	if selector == "CHECKBOX" {
		icon := CIRCLE_ICON
		if selected {
			icon = CIRCLEFILLED_ICON
		}

		if index == cursorPos {
			fmt.Printf("%s %s %s\n", RIGHTPOINTER_ICON, icon, transformed)
		} else {
			fmt.Printf("  %s %s\n", icon, transformed)
		}
	}
}

func displayTitleAndDesc(title, desc string) {
	if strings.TrimSpace(title) != "" {
		fmt.Println(title)
	}

	if strings.TrimSpace(desc) != "" {
		fmt.Printf("%s\n\n", desc)
	}
}
