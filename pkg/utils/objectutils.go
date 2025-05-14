package utils

func OrElseEmpty(value interface{}) interface{} {
	if value == nil {
		return ""
	}
	return value
}
