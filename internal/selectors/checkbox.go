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

	cursorPos := 0
	checkedCount := 0

	byteArr := make([]byte, 3)

	for {
		renderCheckboxItems(decoratedItems, cursorPos, *opts)

		utilities.RecordKeyStroke(byteArr)

		if utilities.IsCtrlC(byteArr) {
			return zeroVal, nil
		}

		if utilities.IsEnter(byteArr) && checkedCount >= int(opts.MinSelection) {
			break
		}

		if utilities.IsSpacebar(byteArr) {
			if !decoratedItems[cursorPos].checked && uint(checkedCount) < opts.MaxSelection {
				decoratedItems[cursorPos].checked = true
				checkedCount++
			} else if decoratedItems[cursorPos].checked {
				decoratedItems[cursorPos].checked = false
				checkedCount--
			}

			continue
		}

		if utilities.IsUpArrow(byteArr) {
			if cursorPos <= 0 {
				cursorPos = len(list) - 1
			} else {
				cursorPos -= 1
			}

			continue
		}

		if utilities.IsDownArrow(byteArr) {
			if cursorPos >= len(list)-1 {
				cursorPos = 0
			} else {
				cursorPos += 1
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
