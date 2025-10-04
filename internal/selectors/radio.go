package selectors

import (
	"errors"

	"github.com/digvijay-tech/interactive_inputs/internal/utilities"
)

type RadioOptions struct {
	Title         string
	Description   string
	TextTransform TextTransform
}

func (ro RadioOptions) GetType() string { return "RADIO" }

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

	var zeroVal T

	if len(list) < 1 {
		return zeroVal, errors.New("list is empty")
	}

	cursorPos := 0

	byteArr := make([]byte, 3)

	for {
		renderRadioItems(list, cursorPos, *opts)

		utilities.RecordKeyStroke(byteArr)

		if utilities.IsCtrlC(byteArr) {
			return zeroVal, nil
		}

		if utilities.IsEnter(byteArr) {
			break
		}

		if utilities.IsUpArrow(byteArr) {
			if cursorPos <= 0 && !enableScroll {
				cursorPos = len(list) - 1
			} else if cursorPos > 0 {
				cursorPos -= 1
			}

			continue
		}

		if utilities.IsDownArrow(byteArr) {
			if cursorPos >= len(list)-1 && !enableScroll {
				cursorPos = 0
			} else if cursorPos < len(list)-1 {
				cursorPos += 1
			}

			continue
		}
	}

	return list[cursorPos], nil
}

func renderRadioItems[T AcceptedListType](list []T, cursorPos int, opts RadioOptions) {
	utilities.ClearTerminal()

	displayTitleAndDesc(opts.Title, opts.Description)

	for i, v := range list {
		moveCursor(i, cursorPos, v, opts.TextTransform, opts.GetType(), false) // selected will always be false for radio
	}
}
