package selectors

import (
	"errors"

	"github.com/digvijay-tech/interactive_inputs/internal/utilities"
)

type CheckboxOptions struct {
	Title         string
	Description   string
	MinSelection  uint
	MaxSelection  uint
	TextTransform TextTransform
}

func (cbo CheckboxOptions) GetType() string { return "CHECKBOX" }

type checkboxItem[T AcceptedListType] struct {
	item    T
	checked bool
}

func Checkbox[T AcceptedListType](list []T, opts *CheckboxOptions) (selectedItems []T, err error) {
	utilities.HideDefaultTerminalCursor()
	defer utilities.ShowDefaultTerminalCursor()

	_, winHeight := utilities.GetWindowSize()
	enableScroll := false

	if winHeight-len(list) <= 5 || len(list) > 10 {
		enableScroll = true
	}

	// override options
	if opts == nil || opts.MaxSelection == 0 || opts.MaxSelection > uint(len(list)) {
		opts = &CheckboxOptions{
			Title:         opts.Title,
			Description:   opts.Description,
			MinSelection:  opts.MinSelection,
			MaxSelection:  uint(len(list)),
			TextTransform: opts.TextTransform,
		}
	}

	var zeroVal []T

	if len(list) < 1 {
		return zeroVal, errors.New("list is empty")
	}

	if opts.MinSelection > uint(len(list)) {
		return zeroVal, errors.New("MinSelection cannot be greater than the list")
	}

	var decoratedItems []checkboxItem[T]

	for _, v := range list {
		decoratedItems = append(decoratedItems, checkboxItem[T]{
			item:    v,
			checked: false,
		})
	}

	byteArr := make([]byte, 3)

	cursorPos, scrollingCursorPos, checkedCount := 0, 0, 0
	var paginatedList []checkboxItem[T]
	itemsPerPage := 10

	if enableScroll {
		paginatedList = append(paginatedList, decoratedItems[:itemsPerPage]...)
	} else {
		paginatedList = decoratedItems
	}

	for {
		renderCheckboxItems(paginatedList, cursorPos, *opts)

		utilities.RecordKeyStroke(byteArr)

		if utilities.IsCtrlC(byteArr) {
			return zeroVal, nil
		}

		if utilities.IsEnter(byteArr) && checkedCount >= int(opts.MinSelection) {
			break
		}

		if utilities.IsSpacebar(byteArr) {
			// scrolling is disabled
			if !enableScroll {
				if !decoratedItems[cursorPos].checked && uint(checkedCount) < opts.MaxSelection {
					decoratedItems[cursorPos].checked = true
					checkedCount++
				} else if decoratedItems[cursorPos].checked {
					decoratedItems[cursorPos].checked = false
					checkedCount--
				}

				continue
			}

			if !decoratedItems[scrollingCursorPos].checked && uint(checkedCount) < opts.MaxSelection {
				decoratedItems[scrollingCursorPos].checked = true
				paginatedList[cursorPos].checked = true
				checkedCount++
			} else if decoratedItems[scrollingCursorPos].checked {
				decoratedItems[scrollingCursorPos].checked = false
				paginatedList[cursorPos].checked = false
				checkedCount--
			}

			continue
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
				paginatedList = append([]checkboxItem[T]{}, paginatedList[:itemsPerPage-1]...)
				paginatedList = append([]checkboxItem[T]{decoratedItems[scrollingCursorPos-1]}, paginatedList...)
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
			if cursorPos < len(paginatedList)-1 && scrollingCursorPos < len(decoratedItems)-1 {
				cursorPos += 1
				scrollingCursorPos += 1
			} else if cursorPos == len(paginatedList)-1 && scrollingCursorPos < len(decoratedItems)-1 {
				paginatedList = append(paginatedList, decoratedItems[scrollingCursorPos+1])
				paginatedList = append([]checkboxItem[T]{}, paginatedList[1:]...)
				cursorPos = len(paginatedList) - 1
				scrollingCursorPos += 1
			} else {
				cursorPos = len(paginatedList) - 1
				scrollingCursorPos = len(decoratedItems) - 1
			}

			continue
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

func renderCheckboxItems[T AcceptedListType](decoratedList []checkboxItem[T], cursorPos int, opts CheckboxOptions) {
	utilities.ClearTerminal()

	displayTitleAndDesc(opts.Title, opts.Description)

	for i, v := range decoratedList {
		moveCursor(i, cursorPos, v.item, opts.TextTransform, opts.GetType(), v.checked)
	}
}
