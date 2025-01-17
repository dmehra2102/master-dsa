package recursion

import "fmt"

func ReverseString(s []byte) {
	result := make([]byte, 0, len(s))

	recursion(s, &result, 0)
	fmt.Println(string(result))
}

func recursion(s []byte, result *[]byte, index int) {
	if index == len(s) {
		return
	}
	recursion(s, result, index+1)
	*result = append(*result, s[index])
}
