package inputs

type acceptedListType interface {
	string | int | int8 | int16 | int32 | int64 | float32 | float64
}
