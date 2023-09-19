package main

import "fmt"

func main() {

	set := New()

	set.add(1)
	set.add(2)
	set.add(3)
	set.add(4)

	fmt.Println(set.data)
	set.remove(3)
	fmt.Println(set.data)
	set.remove(2)
	fmt.Println(set.data)
	set.remove(1)
	fmt.Println(set.data)
	set.remove(4)
	fmt.Println(set.data)
	
}

type Set struct {
	data []interface{}
}

func New() *Set {
	return &Set{}
}

func (s *Set) isEmpty() bool {
	return len(s.data) == 0
}

func (s *Set) size() int {
	return len(s.data)
}

func (s *Set) contains(value interface{}) bool {
	index := s.indexOf(value)
	if index != -1 {
		return true
	}
	return false
}

func (s *Set) indexOf(value interface{}) int {
	for i, val := range s.data {
		if val == value {
			return i
		}
	}
	return -1
}

func (s *Set) add(value interface{}) {
	if !s.contains(value) {
		s.data = append(s.data, value)
	}
}

func (s *Set) remove(value interface{}) {
	if !s.isEmpty() {
		index := s.indexOf(value)
		s.data = append(s.data[:index], s.data[index+1:]...)
	}
}
