package strings

func WildcardMatching(input, pattern string) bool {
	if IsEmpty(input) && IsEmpty(pattern) {
		return true
	}
	if IsNotEmpty(input) && IsEmpty(pattern) {
		return false
	}
	p := []rune(pattern)
	if IsEmpty(input) {
		for _, c := range p {
			if '*' != c {
				return false
			}
		}
		return true
	}
	return wildcardMatchingInternal([]rune(input), p, 0, 0)
}

func wildcardMatchingInternal(input, pattern []rune, iIndex, pIndex int) bool {
	if iIndex >= len(input) && pIndex >= len(pattern) {
		return true
	}
	if iIndex >= len(input) && pattern[pIndex] != '*' {
		return false
	}
	if iIndex >= len(input) && pattern[pIndex] == '*' {
		return wildcardMatchingInternal(input, pattern, iIndex, pIndex+1)
	}
	if iIndex < len(input) && pIndex >= len(pattern) {
		return false
	}
	if input[iIndex] == pattern[pIndex] {
		return wildcardMatchingInternal(input, pattern, iIndex+1, pIndex+1)
	}
	if pattern[pIndex] == '?' {
		return wildcardMatchingInternal(input, pattern, iIndex+1, pIndex+1)
	}
	if pattern[pIndex] == '*' {
		return wildcardMatchingInternal(input, pattern, iIndex+1, pIndex) ||
			wildcardMatchingInternal(input, pattern, iIndex, pIndex+1)
	} else {
		return false
	}
}
