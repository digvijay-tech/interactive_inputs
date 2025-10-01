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

func Radio[T AcceptedListType](list []T, opts *RadioOptions) (selectedItem T, err error) {
	utilities.HideDefaultTerminalCursor()
	defer utilities.ShowDefaultTerminalCursor()

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
		renderList(list, cursorPos, opts.Title, opts.Description, opts.TextTransform)

		oldState := utilities.EnterRawMode()

		byteArr := make([]byte, 3)
		os.Stdin.Read(byteArr)

		utilities.ExitRawMode(oldState)

		// ctrl+c
		if byteArr[0] == 3 {
			break
		}

		// enter
		if byteArr[0] == 13 {
			break
		}

		// up/down arrow
		if byteArr[0] == 27 && byteArr[1] == 91 {
			switch byteArr[2] {
			case 65:
				if cursorPos <= 0 {
					cursorPos = len(list) - 1
				} else {
					cursorPos -= 1
				}

				continue
			case 66:
				if cursorPos >= len(list)-1 {
					cursorPos = 0
				} else {
					cursorPos += 1
				}

				continue
			}
		}
	}

	return list[cursorPos], nil
}

func renderList[T AcceptedListType](list []T, cursorPos int, title string, desc string, textTransform TextTransform) {
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
		for i, v := range list {
			moveCursor(i, cursorPos, fmt.Sprintf("%v", v))
		}

		return
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
}

func moveCursor(index int, cursorPos int, s string) {
	if index == cursorPos {
		fmt.Printf("> %s\n", s)
	} else {
		fmt.Printf("  %s\n", s)
	}
}
