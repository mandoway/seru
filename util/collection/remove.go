package collection

func RemoveAt[T any](collection []T, index int) []T {
	return append(collection[:index], collection[index+1:]...)
}
