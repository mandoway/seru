package collection

func MapSlice[T any, R any](slice []T, transform func(it T) (R, error)) []R {
	arr := make([]R, len(slice))
	for i, val := range slice {
		transformed, err := transform(val)
		if err != nil {
			continue
		}
		arr[i] = transformed
	}
	return arr
}
