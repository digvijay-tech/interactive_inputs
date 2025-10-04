package selectors

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/digvijay-tech/interactive_inputs/internal/utilities"
)

type RadioOptions struct {
	Title         string
	Description   string
	TextTransform TextTransform
}

// TODO: complete scrolling
func Radio[T AcceptedListType](list []T, opts *RadioOptions) (selectedItem T, err error) {
	utilities.HideDefaultTerminalCursor()
	defer utilities.ShowDefaultTerminalCursor()

	_, winHeight := utilities.GetWindowSize()
	enableScroll := false

	if winHeight-len(list) <= 0 || len(list) > 15 {
		enableScroll = true
	}

	// optional param
	if opts == nil {
		opts = &RadioOptions{TextTransform: NONE}
	}

	if len(list) < 1 {
		var zeroVal T
		return zeroVal, errors.New("list is empty")
	}

	cursorPos := 0

	for {
		renderList(list, cursorPos, opts.Title, opts.Description, opts.TextTransform, enableScroll)

		oldState := utilities.EnterRawMode()

		byteArr := make([]byte, 3)
		os.Stdin.Read(byteArr)

		utilities.ExitRawMode(oldState)

		// ctrl+c
		if byteArr[0] == 3 {
			var zeroval T
			return zeroval, nil
		}

		// enter
		if byteArr[0] == 13 {
			break
		}

		// up/down arrow
		if byteArr[0] == 27 && byteArr[1] == 91 {
			switch byteArr[2] {
			case 65:
				if cursorPos <= 0 && !enableScroll {
					cursorPos = len(list) - 1
				} else if cursorPos > 0 {
					cursorPos -= 1
				}

				continue
			case 66:
				if cursorPos >= len(list)-1 && !enableScroll {
					cursorPos = 0
				} else if cursorPos < len(list)-1 {
					cursorPos += 1
				}

				continue
			}
		}
	}

	return list[cursorPos], nil
}

func renderList[T AcceptedListType](list []T, cursorPos int, title string, desc string, textTransform TextTransform, enableScroll bool) {
	utilities.ClearTerminal()

	if strings.TrimSpace(title) != "" {
		fmt.Println(title)
	}

	if strings.TrimSpace(desc) != "" {
		fmt.Printf("%s\n\n", desc)
	}

	listType := utilities.FindType(list, true)
	if listType == "" {
		log.Fatalln("invalid list type")
	}

	if listType != "string" {
		if enableScroll && cursorPos > 0 {
			fmt.Println("\u2191")
		}
		for i, v := range list {
			moveCursor(i, cursorPos, fmt.Sprintf("%v", v))
		}

		if enableScroll && cursorPos < len(list) {
			fmt.Println("\u2193")
		}

		return
	}

	if enableScroll && cursorPos > 0 {
		fmt.Println("\u2191")
	} else {
		fmt.Println("")
	}
	for i, v := range list {
		switch textTransform.String() {
		case "uppercase":
			moveCursor(i, cursorPos, strings.ToUpper(fmt.Sprintf("%v", v)))
		case "lowercase":
			moveCursor(i, cursorPos, strings.ToLower(fmt.Sprintf("%v", v)))
		case "capitalise":
			moveCursor(i, cursorPos, utilities.ToCapitalise(fmt.Sprintf("%v", v)))
		default:
			moveCursor(i, cursorPos, fmt.Sprintf("%v", v))
		}
	}
	if enableScroll && cursorPos < len(list)-1 {
		fmt.Println("\u2193")
	}
}

func moveCursor(index int, cursorPos int, s string) {
	if index == cursorPos {
		fmt.Printf("> %s\n", s)
	} else {
		fmt.Printf("  %s\n", s)
	}
}
