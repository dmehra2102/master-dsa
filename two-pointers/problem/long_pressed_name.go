package twopointers

func IsLongPressedName(name string, typed string) bool {
	i := 0
	j := 0
	N := len(name)
	M := len(typed)

	if N > M {
		return false
	}

	for j < M {
		if i < N && name[i] == typed[j] {
			i++
			j++
		} else if j != 0 && typed[j] == typed[j-1] {
			j++
		} else {
			return false
		}
	}
	return i == N
}
