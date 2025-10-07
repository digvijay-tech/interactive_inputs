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

// Display scroll indicating arrow for either up or down, 'orientation' can be "1" for up arrow and "-1"
// for down arrow. 'position' is the cursor position of main array/slice and 'arraySize' is the length of
// main array/slice.
func displayScrollIndicator(orientation int8, position int, arraySize int) {
	switch orientation {
	case 1:
		if position > 0 {
			fmt.Println(UPARROW_ICON)
			return
		}

		fmt.Printf("\n")
	case -1:
		if position < arraySize-1 {
			fmt.Println(DOWNARROW_ICON)
			return
		}

		fmt.Printf("\n")
	}
}
