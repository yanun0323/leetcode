package main

import "container/list"

type LRUCache struct {
	Cap  int
	Set  map[int]*list.Element
	List *list.List
}

type Elem struct {
	Key, Value int
}

func NewElem(key, value int) Elem {
	return Elem{Key: key, Value: value}
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		Cap:  capacity,
		Set:  make(map[int]*list.Element, 0),
		List: list.New(),
	}

}

func (cache *LRUCache) Get(key int) int {
	if _, exist := cache.Set[key]; exist {
		cache.List.MoveToBack(cache.Set[key])
		return cache.Set[key].Value.(Elem).Value
	}
	return -1
}

func (cache *LRUCache) Put(key int, value int) {
	if _, exist := cache.Set[key]; exist {
		elem := cache.Set[key].Value.(Elem)
		elem.Value = value
		cache.Set[key].Value = elem
		cache.List.MoveToBack(cache.Set[key])
		return
	}
	cache.Set[key] = cache.List.PushBack(Elem{Key: key, Value: value})

	if len(cache.Set) > cache.Cap {
		elem := cache.List.Front()
		cache.List.Remove(elem)
		delete(cache.Set, elem.Value.(Elem).Key)
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
