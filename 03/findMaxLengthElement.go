package main

func findMaxLengthElement(strings []string) []string {
	maxLen := 0
	var maxStrings []string
	for _, s := range strings {
		if len(s) > maxLen {
			maxLen = len(s)
			maxStrings = []string{}
			maxStrings = append(maxStrings, s)
		} else if len(s) == maxLen {
			maxStrings = append(maxStrings, s)
		}
	}
	return maxStrings
}
