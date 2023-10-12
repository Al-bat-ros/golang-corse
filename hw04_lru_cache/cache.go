package hw04lrucache

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
}

type cacheItem struct {
	Key   Key
	Value interface{}
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}

func (c *lruCache) Set(key Key, value any) bool {
	if elem, exists := c.items[key]; exists {
		elem.Value.(*cacheItem).Value = value
		c.queue.MoveToFront(elem)

		return true
	}

	if c.queue.Len() >= c.capacity {
		lastElem := c.queue.Back()

		if lastElem != nil {
			delete(c.items, lastElem.Value.(*cacheItem).Key)
			c.queue.Remove(lastElem)
		}
	}

	newItem := &cacheItem{key, value}

	elem := c.queue.PushFront(newItem)
	c.items[key] = elem

	return false
}

func (c *lruCache) Get(key Key) (any, bool) {
	if elem, exists := c.items[key]; exists {
		c.queue.MoveToFront(elem)

		return elem.Value.(*cacheItem).Value, true
	}

	return nil, false
}

func (c *lruCache) Clear() {
	c.items = make(map[Key]*ListItem)
}
