package selectors

type AcceptedListType interface {
	string | int | int8 | int16 | int32 | int64 | float32 | float64
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
