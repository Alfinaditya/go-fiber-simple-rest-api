package utils

func UpdateIfNotNil[T any](target *T, source *T) {
	if source != nil {
		*target = *source
	}
}
