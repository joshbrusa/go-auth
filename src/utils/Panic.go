package utils

type PanicData struct {
	Msg string
	File string
	Line int
}

func NewPanic(panicData PanicData) {
	panic(panicData)
}