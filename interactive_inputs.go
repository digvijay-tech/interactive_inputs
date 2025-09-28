// Package interactive_inputs provides utilities for creating interactive terminal user interfaces.
package interactive_inputs

import (
	"log"

	"github.com/digvijay-tech/interactive_inputs/internal/selectors"
	"github.com/digvijay-tech/interactive_inputs/internal/utilities"
)

func init() {
	log.SetPrefix("[interactive_inputs] ")

	if !utilities.IsTerminalCapable() {
		log.Println("WARNING: terminal may not support interactions.")
	}
}

const (
	UPPERCASE  = selectors.UPPERCASE
	LOWERCASE  = selectors.LOWERCASE
	CAPITALISE = selectors.CAPITALISE
	NONE       = selectors.NONE
)

type RadioOptions = selectors.RadioOptions

func Radio[T selectors.AcceptedListType](list []T, opts *RadioOptions) (selectedItem T, err error) {
	return selectors.Radio(list, opts)
}
