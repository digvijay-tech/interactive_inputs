package selectors

const (
	RIGHTPOINTER_ICON = "\u25B6"
	UPARROW_ICON      = "\u2191"
	DOWNARROW_ICON    = "\u2193"
	RIGHTARROW_ICON   = "\u2192"
	LEFTARROW_ICON    = "\u2190"
	CIRCLE_ICON       = "\u25CB"
	CIRCLEFILLED_ICON = "\u25CF"
)

type AcceptedListType interface {
	string |
		int | int8 | int16 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64 |
		float32 | float64
}

type TextTransform int

const (
	NONE TextTransform = iota
	CAPITALISE
	LOWERCASE
	UPPERCASE
)

func (t TextTransform) String() string {
	switch t {
	case LOWERCASE:
		return "lowercase"
	case UPPERCASE:
		return "uppercase"
	case CAPITALISE:
		return "capitalise"
	default:
		return "none"
	}
}
