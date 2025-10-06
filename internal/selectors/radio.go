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

func Radio[T AcceptedListType](list []T, opts *RadioOptions) (selectedItem T, err error) {
	utilities.HideDefaultTerminalCursor()
	defer utilities.ShowDefaultTerminalCursor()

	_, winHeight := utilities.GetWindowSize()
	enableScroll := false

	if winHeight-len(list) <= 5 || len(list) > 10 {
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

	byteArr := make([]byte, 3)

	cursorPos, scrollingCursorPos := 0, 0
	var paginatedList []T
	itemsPerPage := 10

	if enableScroll {
		paginatedList = append(paginatedList, list[:itemsPerPage]...)
	} else {
		paginatedList = list
	}

	for {
		renderRadioItems(paginatedList, cursorPos, *opts)

		utilities.RecordKeyStroke(byteArr)

		if utilities.IsCtrlC(byteArr) {
			return zeroVal, nil
		}

		if utilities.IsEnter(byteArr) {
			break
		}

		if utilities.IsUpArrow(byteArr) {
			// scrolling is disabled
			if !enableScroll {
				if cursorPos <= 0 {
					cursorPos = len(paginatedList) - 1
				} else {
					cursorPos -= 1
				}

				continue
			}

			// scrolling is enabled
			if cursorPos > 0 && scrollingCursorPos > 0 {
				cursorPos -= 1
				scrollingCursorPos -= 1
			} else if cursorPos == 0 && scrollingCursorPos > 0 {
				paginatedList = append([]T{}, paginatedList[:itemsPerPage-1]...)
				paginatedList = append([]T{list[scrollingCursorPos-1]}, paginatedList...)
				cursorPos = 0
				scrollingCursorPos -= 1
			} else {
				cursorPos = 0
				scrollingCursorPos = 0
			}

			continue
		}

		if utilities.IsDownArrow(byteArr) {
			// scrolling is disabled
			if !enableScroll {
				if cursorPos >= len(paginatedList)-1 {
					cursorPos = 0
				} else {
					cursorPos += 1
				}

				continue
			}

			// scrolling is enabled
			if cursorPos < len(paginatedList)-1 && scrollingCursorPos < len(list)-1 {
				cursorPos += 1
				scrollingCursorPos += 1
			} else if cursorPos == len(paginatedList)-1 && scrollingCursorPos < len(list)-1 {
				paginatedList = append(paginatedList, list[scrollingCursorPos+1])
				paginatedList = append([]T{}, paginatedList[1:]...)
				cursorPos = len(paginatedList) - 1
				scrollingCursorPos += 1
			} else {
				cursorPos = len(paginatedList) - 1
				scrollingCursorPos = len(list) - 1
			}

			continue
		}
	}

	return paginatedList[cursorPos], nil
}

func renderRadioItems[T AcceptedListType](list []T, cursorPos int, opts RadioOptions) {
	utilities.ClearTerminal()

	displayTitleAndDesc(opts.Title, opts.Description)

	for i, v := range list {
		moveCursor(i, cursorPos, v, opts.TextTransform, opts.GetType(), false) // selected will always be false for radio
	}
}
