/*
The Super Naive Generics Approach!

TODO: 	Find a simple syntax for initializing the translator
	Doesn't matter, all that needs to happen is that we have
	to know the type that uses a generic type, and it's target
	concrete type.
	
	Example:
	
		type TYPE1 int
		s := map[string]TYPE1
		
	This is probably the easiest one, because it will get
	converted into the form of:
	
		// type TYPE1 int
		s := map[string]int
	
	And the declaration will get commented out, because
	it's no longer needed.


TODO: Allow Multiple usages, clone code and inline it.
	x := Set_Type1
	y := Set_Type2
	z := Set_Type3

*/

package main

import (
	"fmt"
)


type Set struct {
	itemMap map[type1]bool
	items[] type1
}


func (s *Set) TypeName() string {
	return "type1"
}


func (s *Set) Add(item type1) {
	for _, v := range s.items {
		if v == item {
			return
		}
	}
	s.items = append(s.items, item)
}


func (s *Set) Remove(item type1) {
	for i, v := range s.items {
		if v == item {
			if len(s.items) <= 1 {
				s.Clear()
				return
			}
			last := len(s.items) - 1
			s.swap(i, last)
			s.items = s.items[:last]
			return
		}
	}
}

func (s *Set) swap(i, j int) {
	s.items[i], s.items[j] = s.items[j], s.items[i]
}

func (s *Set) Clear() {
	s.items = []type1{}
}

func (s *Set) Contains(item type1) bool {
	for _, v := range s.items {
		if v == item {
			return true
		}
	}
	return false
}


func (s Set) String() string {
	return fmt.Sprint(s.items)
}


func NewSet() *Set {
	return &Set{
		itemMap: map[type1]bool{},
	}
}

func main() {
	fmt.Println("Generic experiment!")
	fmt.Println("Using Type: type1")
	x := NewSet() // TODO: declare typename = int
	x.Add(1)
	x.Add(2)
	fmt.Println("after add:", x)
	x.Add(1)
	x.Add(7)
	fmt.Println("after more add", x)
	x.Remove(2)
	fmt.Println("after remove", x)
	fmt.Println("Set contains the type: ", x.TypeName())
}



