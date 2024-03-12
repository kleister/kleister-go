package kleister

// PtrString simple converts a pointer string to a regular string.
func PtrString(val *string) string {
	if val == nil {
		return ""
	}

	return *val
}
