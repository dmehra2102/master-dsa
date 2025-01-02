package hashtable

import "fmt"

/* HashTable structure : will hold an array
	- Insert
	- Search
	- Delete

   Bucket Structure : It's basically an Linked List
	- Insert
	- Search
	- Delete

   hash method

   Init function : for initializing hashtable
*/

const ArraySize int = 7

type HashTable struct {
	array [ArraySize]*bucket
}

type bucket struct {
	head *bucketNode
}

type bucketNode struct {
	key  string
	next *bucketNode
}

func (h *HashTable) Insert(key string) {
	index := hashMethod(key)
	h.array[index].insert(key)
}

func (h *HashTable) Search(key string) bool {
	index := hashMethod(key)
	return h.array[index].search(key)
}

func (h *HashTable) Delete(key string) {
	index := hashMethod(key)
	h.array[index].delete(key)
}

// Methods for bucket
func (b *bucket) insert(key string) {
	newBucket := &bucketNode{key: key}
	newBucket.next = b.head
	b.head = newBucket
}
func (b *bucket) search(key string) bool {
	currentNode := b.head
	for currentNode != nil {
		if currentNode.key == key {
			return true
		}

		currentNode = currentNode.next
	}
	return false
}
func (b *bucket) delete(key string) {

	if b.head == nil {
		return
	}

	if b.head.key == key {
		b.head = b.head.next
		return
	}

	prevNode := b.head

	for prevNode.next != nil {
		if prevNode.next.key == key {
			prevNode.next = prevNode.next.next
			return
		}

		prevNode = prevNode.next
	}
}

func Init() *HashTable {
	result := &HashTable{}

	for i := range result.array {
		result.array[i] = &bucket{}
	}

	return result
}

// method for hashing the key and return hash code
func hashMethod(key string) int {
	sum := 0
	for _, v := range key {
		sum += int(v)
	}

	return sum % ArraySize
}

func HashFunction() {
	hashTable := Init()

	fmt.Printf("Is %s present %t", "Deepanshu", hashTable.Search("Deepanshu"))
	fmt.Println()

	hashTable.Insert("ERIC")
	hashTable.Insert("Deepanshu")
	hashTable.Insert("Mehra")
	fmt.Printf("Is %s present %t", "Deepanshu", hashTable.Search("Deepanshu"))
	fmt.Println()

	hashTable.Delete("ERIC")
	hashCode := hashMethod("Deepanshu")
	fmt.Printf("HashCode : %d", hashCode)
	fmt.Println()

	fmt.Printf("Is %s present %t", "ERIC", hashTable.Search("ERIC"))
	fmt.Println()
}
