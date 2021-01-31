package single

import "errors"

type ArrayStack struct {
	cap    int
	length int
	values []interface{}
}

func NewArrayStack(cap int) *ArrayStack {
	return &ArrayStack{
		cap:    cap,
		length: 0,
		values: make([]interface{}, cap),
	}
}

func (s *ArrayStack) Pop() (interface{}, error) {
	if s.length == 0 {
		return nil, errors.New("no elements")
	}
	value := s.values[s.length-1]
	s.values[s.length-1] = nil
	s.length--
	return value, nil
}

func (s *ArrayStack) Push(value interface{}) error {
	if s.length == s.cap {
		return errors.New("cap enough")
	}
	s.values[s.length] = value
	s.length++
	return nil
}

func (s *ArrayStack) Print() []interface{} {
	return s.values[:s.length]
}
