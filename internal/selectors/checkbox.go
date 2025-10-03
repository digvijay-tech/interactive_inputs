package selectors

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/digvijay-tech/interactive_inputs/internal/utilities"
)

type CheckboxOptions struct {
	Title         string
	Description   string
	MinSelection  uint
	MaxSelection  uint
	TextTransform TextTransform
}

type checkboxItem[T AcceptedListType] struct {
	item    T
	checked bool
}

func Checkbox[T AcceptedListType](list []T, opts *CheckboxOptions) (selectedItems []T, err error) {
	utilities.HideDefaultTerminalCursor()
	defer utilities.ShowDefaultTerminalCursor()

	// default params
	if opts == nil {
		opts = &CheckboxOptions{
			MaxSelection:  uint(len(list)),
			TextTransform: NONE,
		}
	}

	var zeroVal []T

	if len(list) < 1 {
		return zeroVal, errors.New("list is empty")
	}

	if opts.MinSelection > uint(len(list)) {
		return zeroVal, errors.New("MinSelection cannot be greater than the list length")
	}

	var decoratedItems []checkboxItem[T]

	for _, v := range list {
		decoratedItems = append(decoratedItems, checkboxItem[T]{
			item:    v,
			checked: false,
		})
	}

	cursorPos := 0
	checkedCount := 0

	for {
		renderDecoratedList(decoratedItems, cursorPos, opts.Title, opts.Description, opts.TextTransform)

		oldState := utilities.EnterRawMode()

		byteArr := make([]byte, 3)
		os.Stdin.Read(byteArr)

		utilities.ExitRawMode(oldState)

		// ctrl+c
		if byteArr[0] == 3 {
			break
		}

		// enter
		if byteArr[0] == 13 && checkedCount >= int(opts.MinSelection) {
			break
		}

		// spacebar
		if byteArr[0] == 32 {
			if decoratedItems[cursorPos].checked {
				decoratedItems[cursorPos].checked = false
				checkedCount--
			} else {
				decoratedItems[cursorPos].checked = true
				checkedCount++
			}

			continue
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

	// filter checked
	var selected []T
	for _, v := range decoratedItems {
		if v.checked {
			selected = append(selected, v.item)
		}
	}

	return selected, nil
}

func renderDecoratedList[T AcceptedListType](decoratedList []checkboxItem[T], cursorPos int, title string, desc string, textTransform TextTransform) {
	utilities.ClearTerminal()

	if strings.TrimSpace(title) != "" {
		fmt.Println(title)
	}

	if strings.TrimSpace(desc) != "" {
		fmt.Printf("%s\n\n", desc)
	}

	listType := utilities.FindType(decoratedList[0].item, false)
	if listType == "" {
		log.Fatalln("invalid list type")
	}

	if listType != "string" {
		for i, v := range decoratedList {
			moveCursor(i, cursorPos, fmt.Sprintf("%v", v.item))
		}

		return
	}

	// default unchecked icon
	var icon = "○"

	for i, v := range decoratedList {
		if v.checked {
			icon = "●"
		}

		switch textTransform.String() {
		case "uppercase":
			moveCursor(i, cursorPos, icon+" "+strings.ToUpper(fmt.Sprintf("%v", v.item)))
		case "lowercase":
			moveCursor(i, cursorPos, icon+" "+strings.ToLower(fmt.Sprintf("%v", v.item)))
		case "capitalise":
			moveCursor(i, cursorPos, icon+" "+utilities.ToCapitalise(fmt.Sprintf("%v", v.item)))
		default:
			moveCursor(i, cursorPos, fmt.Sprintf("%s %v", icon, v.item))
		}

		// reset for next iteration
		icon = "○"
	}
}
