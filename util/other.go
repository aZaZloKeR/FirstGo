package util

func SendWithoutBlock[T any](message T, to chan T) {
	select {
	case to <- message:
	default:
	}
}
