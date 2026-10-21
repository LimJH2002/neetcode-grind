type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	var sb strings.Builder
	for _, str := range strs {
		sb.WriteString(strconv.Itoa(len(str)))
		sb.WriteString("#")
		sb.WriteString(str)
	}

	return sb.String()
}

func (s *Solution) Decode(encoded string) []string {
	result := make([]string, 0)

	for i, j := 0, 0; j < len(encoded); j++ {
		if encoded[j] == '#' {
			length, _ := strconv.Atoi(encoded[i : j])
			result = append(result, encoded[j+1 : j+1+length])
			j = j + length + 1
			i = j
		}
	}

	return result
}
