package helpers

func ArrayReverse[T any](array *[]T) []T {

	length := len(*array)

	tmpArray := make([]T, length)

	for i := 0; i < length; i++ {
		tmpArray[i] = (*array)[length-i-1]
	}

	return tmpArray
}
