package src

type (
	bucketItem struct {
		key   float64
		value string
	}

	Bucket struct {
		items []*bucketItem
	}
)

// створення бакета
func NewBucket() *Bucket {
	return &Bucket{
		items: make([]*bucketItem, 0, 20),
	}
}

// додавання значення
func (b *Bucket) Put(key float64, value string) {
	indexOfItem := b.getIndexOfItemByKey(key)

	if indexOfItem < 0 {
		newItem := &bucketItem{
			key:   key,
			value: value,
		}
		b.items = append(b.items, newItem)

		return
	}

	b.items[indexOfItem].value = value
}

// отримання значення за ключем
func (b *Bucket) Get(key float64) (string, bool) {
	for _, item := range b.items {
		if item.key == key {
			return item.value, true
		}
	}

	return "", false
}

// видалення значення за ключем
func (b *Bucket) Delete(key float64) {
	indexOfItem := b.getIndexOfItemByKey(key)
	if indexOfItem < 0 {
		return
	}

	b.items[indexOfItem] = nil
}

// returns -1 if there is not such item
func (b *Bucket) getIndexOfItemByKey(key float64) int {
	for idx, item := range b.items {
		if item.key == key {
			return idx
		}
	}

	return -1
}
