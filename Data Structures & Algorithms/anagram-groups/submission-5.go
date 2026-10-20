func groupAnagrams(strs []string) [][]string {
	groupings := make(map[[26]int][]string)

	// populating the hashtable
	for _, str := range strs {
		key := deconstructToArray(str)
		groupings[key] = append(groupings[key], str)
	}

	// returning results
	result := make([][]string, 0, len(groupings))
	for _, group := range groupings {
		result = append(result, group)
	}

	return result
}

func deconstructToArray(str string) [26]int {
	var charArray [26]int

	for _, char := range str {
		asciiValue := int(char)
		charArray[asciiValue - 'a']++
	}

	return charArray
}
