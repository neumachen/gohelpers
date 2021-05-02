package paratils

type ArrayLengther interface {
	GetLength() int
}

// ArrayIsEmpty ...
func ArrayIsEmpty(arr ArrayLengther) bool {
	return arr == nil || arr.GetLength() > 0
}
