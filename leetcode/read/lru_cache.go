package read

type Cache struct {
	key, value int
	pre, next  *Cache
}

func newCache(key, value int) *Cache {
	return &Cache{
		key:   key,
		value: value,
	}
}

type LRUCache struct {
	Head        *Cache
	Trail       *Cache
	hset        map[int]*Cache
	cap, length int
}

func Constructor(capacity int) LRUCache {
	h, t := newCache(0, 0), newCache(0, 0)
	h.next, t.pre = t, h
	return LRUCache{
		Head:   h,
		Trail:  t,
		hset:   make(map[int]*Cache, capacity),
		cap:    capacity,
		length: 0,
	}
}

func (this *LRUCache) Get(key int) int {
	if cache, exists := this.hset[key]; exists {
		this.moveToHeader(cache)
		return cache.value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if cache, exists := this.hset[key]; exists {
		if cache.value != value {
			cache.value = value
		}
		this.moveToHeader(cache)
	} else {
		if this.length >= this.cap {
			c := this.deleteFromTrail()
			delete(this.hset, c.key)
			this.length--
		}
		c := newCache(key, value)
		this.addFromHeader(c)
		this.hset[c.key] = c
		this.length++
	}
}

func (this *LRUCache) moveToHeader(cache *Cache) {
	this.deleteLinker(cache)
	this.addFromHeader(cache)
}

func (this *LRUCache) addFromHeader(cache *Cache) {
	next := this.Head.next
	cache.next, cache.pre = next, this.Head
	this.Head.next, next.pre = cache, cache
}

func (this *LRUCache) deleteLinker(cache *Cache) {
	cache.pre.next = cache.next
	cache.next.pre = cache.pre
}

func (this *LRUCache) deleteFromTrail() *Cache {
	t := this.Trail.pre
	t.pre.next, this.Trail.pre = this.Trail, t.pre
	return t
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
