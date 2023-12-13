package utils

func RemoveFirstRune(s string) string {
	for _, r := range s {
		return string(s[len(string(r)):])
	}
	return ""
}
