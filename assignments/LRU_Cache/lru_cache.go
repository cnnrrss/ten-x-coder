package main

// LRUCache using map and doubly linked list
type LRUCache struct {
	head, tail *node
	store      map[int]*node
	cap        int
}

// NOTE: could use container/list from std lib
// for dll but decreased performance

// node of dll
type node struct {
	prev, next *node
	key, val   int
}

// newNode is a constructor to create a new node
func newNode(key, val int) *node {
	return &node{key: key, val: val}
}

// Constructor rolls out a new LRUCache
func Constructor(capacity int) LRUCache {
	return LRUCache{cap: capacity, store: make(map[int]*node)}
}

func (cache *LRUCache) removeFromList(n *node) {
	if n == cache.head {
		cache.head = n.next
	}
	if n == cache.tail {
		cache.tail = n.prev
	}
	if n.next != nil {
		n.next.prev = n.prev
	}
	if n.prev != nil {
		n.prev.next = n.next
	}
}

func (cache *LRUCache) addToList(n *node) {
	n.prev = nil
	n.next = cache.head
	if n.next != nil {
		n.next.prev = n
	}
	cache.head = n
	if cache.tail == nil {
		cache.tail = n
	}
}

// Get retrieves value from the cache using key
func (cache *LRUCache) Get(key int) int {
	n, ok := cache.store[key]
	if !ok {
		return -1
	}
	cache.removeFromList(n)
	cache.addToList(n)
	return n.val
}

// Put sets a key, val pair to the cache
func (cache *LRUCache) Put(key, val int) {
	n, ok := cache.store[key]
	if !ok {
		n = newNode(key, val)
		cache.store[key] = n
	} else {
		n.val = val
		cache.removeFromList(n)
	}
	cache.addToList(n)
	if len(cache.store) > cache.cap {
		n = cache.tail
		if n != nil {
			cache.removeFromList(n)
			delete(cache.store, n.key)
		}
	}
}

func main() {
	cach := Constructor(1)
	cach.Put(2, 1)
}
