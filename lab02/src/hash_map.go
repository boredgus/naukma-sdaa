package src

import (
	"fmt"
	"math"
)

// константа золотого розподілу
const GOLDEN_RATIO float64 = 0.6180339887

type HashTable struct {
	size  int
	table []*Bucket
}

// створення хеш-балиці
func NewHashTable(size int) *HashTable {
	return &HashTable{
		size:  size,
		table: make([]*Bucket, size),
	}
}

// додавання елемента
func (h *HashTable) Put(key float64, value string) {
	index := h.hash(key)
	if index >= h.size || index < 0 {
		fmt.Println("> ERROR: hash index is out of range")

		return
	}

	bucket := h.table[index]
	if bucket != nil {
		bucket.Put(key, value)

		return
	}

	newBucket := NewBucket()

	newBucket.Put(key, value)
	h.table[index] = newBucket
}

// видалення елемента
func (h *HashTable) Delete(key float64) {
	index := h.hash(key)
	if index >= h.size || index < 0 {
		fmt.Println("> ERROR: hash index is out of range")

		return
	}
	bucket := h.table[index]
	if bucket != nil {

	}
}

// пошук елемента
func (h *HashTable) Search(key float64) (string, bool) {
	index := h.hash(key)
	if index >= h.size {
		fmt.Println("> ERROR: hash index is out of range")

		return "", false
	}

	bucket := h.table[index]
	if bucket == nil {
		return "", false
	}

	return bucket.Get(key)
}

// виведення непустих бакетів хеш-таблиці у консоль
func (h *HashTable) PrintSelf() {
	fmt.Println("\nHash table state:")
	for idx, bucket := range h.table {
		if bucket != nil {
			fmt.Printf("[%v]\t %v\n", idx, bucket)
		}
	}
}

// функція для генерації хеша
func (h *HashTable) hash(key float64) int {
	hash := key * GOLDEN_RATIO
	fractionalPart := hash - math.Floor(hash)

	return int(math.Floor(float64(h.size) * fractionalPart))
}
