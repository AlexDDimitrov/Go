package main

func makeGood(s string) string {
	result := NewStack[byte]()
	result.Push(s[0])
	for i := 1; i < len(s); i++ {
		if len(result.Array.data) > 0 {
			if s[i] >= 65 && s[i] <= 90 {
				if (result.Array.Last() - s[i]) == 32 {
					result.Pop()
					continue
				}
			} else if s[i] >= 97 && s[i] <= 122 {
				if s[i]-result.Array.Last() == 32 {
					result.Pop()
					continue
				}
			}
		}

		result.Push(s[i])
	}
	return string(result.Array.data)
}
