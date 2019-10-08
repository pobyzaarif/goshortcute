package shortcute

// StringSliceContains : tells whether a contains x
func StringSliceContains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}
