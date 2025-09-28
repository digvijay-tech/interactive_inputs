package selectors

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/digvijay-tech/interactive_inputs/internal/utilities"
)

type RadioOptions struct {
	Title       string
	Description string
}

func Radio[T AcceptedListType](list []T, opts *RadioOptions) (selectedItem T, err error) {
	utilities.HideDefaultTerminalCursor()
	defer utilities.ShowDefaultTerminalCursor()

	// optional param
	if opts == nil {
		opts = &RadioOptions{
			Title:       "",
			Description: "",
		}
	}

	if len(list) < 1 {
		var zeroVal T
		return zeroVal, errors.New("list is empty")
	}

	cursorPos := 0

	for {
		renderList(list, cursorPos, opts.Title, opts.Description)

		oldState := utilities.EnterRawMode()

		byteArr := make([]byte, 3)
		os.Stdin.Read(byteArr)

		utilities.ExitRawMode(oldState)

		if byteArr[0] == 3 || byteArr[0] == 13 {
			break
		}

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

func renderList[T AcceptedListType](list []T, cursorPos int, title string, desc string) {
	utilities.ClearTerminal()

	if strings.TrimSpace(title) != "" {
		fmt.Println(title)
	}

	if strings.TrimSpace(desc) != "" {
		fmt.Printf("%s\n\n", desc)
	}

	for i, v := range list {
		if i == cursorPos {
			fmt.Printf("> %v\n", v)
		} else {
			fmt.Printf("  %v\n", v)
		}
	}
}
