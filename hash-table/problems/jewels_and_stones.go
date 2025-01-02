package hashtable

/*
Empty Struct: The empty struct (struct{}) in Go occupies zero bytes of memory. When you use it as a value in a map, you are effectively only using the memory needed for the key (in this case, the rune representing the jewel). This means that you do not allocate any additional memory for the value, which would be necessary if you used a more complex type (like string or int).
*/

func NumJewelsInStones(jewels string, stones string) int {
	// STEP-1 : Create a hash table
	hashTable := map[rune]struct{}{}

	// STEP-2 : Loop through a jewels string
	for _, val := range jewels {
		hashTable[val] = struct{}{}
	}

	// STEP-3 : Lopp through a stones string and check for it and increment the count
	count := 0
	for _, val := range stones {
		if _, exists := hashTable[val]; exists {
			count++
		}
	}

	return count
}
