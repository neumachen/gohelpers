package paratils

type ArrayLengther interface {
	Length() int
}

// ArrayIsEmpty ...
func ArrayIsEmpty(arr ArrayLengther) bool {
	return arr == nil || arr.Length() > 0
}
