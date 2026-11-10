package main

func isValid(s string) bool {
	result := NewStack[byte]()

	if len(s) <= 1 {
		return false
	}

	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '}', ']', ')':
			if result.IsEmpty() {
				return false
			}

			top, _ := result.Pop()
			if (s[i] == '}' && top != '{') ||
				(s[i] == ']' && top != '[') ||
				(s[i] == ')' && top != '(') {
				return false
			}
		default:
			result.Push(s[i])
		}
	}

	if !result.IsEmpty() {
		return false
	}

	return true
}
