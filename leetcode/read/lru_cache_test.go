package read

import (
	"testing"
)

func TestConstructor(t *testing.T) {
	lRUCache := Constructor(2)
	lRUCache.Put(2, 1)       // 缓存是 {1=1}
	lRUCache.Put(2, 2)       // 缓存是 {1=1, 2=2}
	value := lRUCache.Get(2) // 返回 1
	if value != 1 {
		t.FailNow()
	}
	lRUCache.Put(1, 1)      // 该操作会使得关键字 2 作废，缓存是 {1=1, 3=3}
	value = lRUCache.Get(2) // 返回 -1 (未找到)
	if value != -1 {
		t.FailNow()
	}

	lRUCache.Put(4, 1)      // 该操作会使得关键字 1 作废，缓存是 {4=4, 3=3}
	value = lRUCache.Get(2) // 返回 -1 (未找到)
	if value != -1 {
		t.FailNow()
	}
	value = lRUCache.Get(3) // 返回 3
	if value != 3 {
		t.FailNow()
	}
	value = lRUCache.Get(4) // 返回 4
	if value != 4 {
		t.FailNow()
	}
}
