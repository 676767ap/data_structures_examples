package hash_table

import (
	"hash/fnv"
)

type HashTable struct {
	data [][]string
	size int
}

func NewHashTable(size int) *HashTable {
	data := make([][]string, size)
	return &HashTable{data: data, size: size}
}

func (ht *HashTable) hash(key string) int {
	hash := fnv.New32a()
	hash.Write([]byte(key))
	return int(hash.Sum32()) % ht.size
}

func (ht *HashTable) Insert(key, value string) {
	index := ht.hash(key)
	ht.data[index] = append(ht.data[index], key+"="+value)
}

func (ht *HashTable) Get(key string) string {
	index := ht.hash(key)
	for _, pair := range ht.data[index] {
		if pair[:len(key)] == key {
			return pair[len(key)+1:]
		}
	}
	return ""
}

// func main() {
// 	ht := NewHashTable(10)
// 	ht.Insert("name", "Alice")
// 	ht.Insert("age", "30")
// 	fmt.Println(ht.Get("name")) // Output: Alice
// 	fmt.Println(ht.Get("age"))  // Output: 30
// }
