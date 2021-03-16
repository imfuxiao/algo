package main

import "fmt"

const (
	length = 8
)

type BitMap struct {
	values []byte
	cap    int
}

func NewBitMap(cap int) *BitMap {
	values := make([]byte, cap/length)
	return &BitMap{
		values: values,
		cap:    cap,
	}
}

func (b *BitMap) Set(value int) *BitMap {
	if value > b.cap || value < 0 {
		return b
	}
	index := value / length
	b.values[index] |= 1 << (value % length)
	return b
}

func (b *BitMap) Exists(value int) bool {
	if value > b.cap || value < 0 {
		return false
	}
	index := value / length
	return b.values[index]&(1<<(value%length)) != 0
}

func main() {
	b := NewBitMap(1000)
	b.Set(10).Set(100).Set(55)
	fmt.Println(b.Exists(55))
	fmt.Println(b.Exists(10))
	fmt.Println(b.Exists(100))
	fmt.Println(b.Exists(90))
}
