type Stack struct {
    items []rune
}

func (s *Stack) Push(data rune) {
    s.items = append(s.items, data)
}

func (s *Stack) Pop() (rune, bool) {
    if s.IsEmpty() {
        return '0', false
    }
    n := len(s.items) - 1
    item := s.items[n]
    s.items = s.items[:n]
    return item, true
}

func (s *Stack) Peek() (rune, bool) {
    if s.IsEmpty() {
        return '0', false
    }

	return s.items[len(s.items) - 1], true
}

func (s *Stack) IsEmpty() bool {
    return len(s.items) == 0
}

func isValid(s string) bool {
	pairs := map[rune]rune{')': '(', '}': '{', ']': '['}
	stack := &Stack{}

	for _, char := range s {
		opener, isCloser := pairs[char]
		if !isCloser {
			stack.Push(char)
			continue
		}
		top, ok := stack.Pop()
		if !ok || top != opener {
			return false
		}
	}
	return stack.IsEmpty()
}
