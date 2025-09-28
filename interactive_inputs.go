package interactive_inputs

import (
	"log"

	"github.com/digvijay-tech/interactive_inputs/internal/selectors"
)

func init() {
	log.SetPrefix("[interactive_inputs] ")
}

type RadioOptions = selectors.RadioOptions

func Radio[T selectors.AcceptedListType](list []T, opts *RadioOptions) (selectedItem T, err error) {
	return selectors.Radio(list, opts)
}
