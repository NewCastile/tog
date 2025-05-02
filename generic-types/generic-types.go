package main

import "fmt"

// List represents a singly-linked list that holds
// values of any type.
type List[T any] struct {
	next *List[T]
	val  T
}

type Pet struct {
	name   string
	specie string
	age    int
}

func (l *List[T]) Push(new T) T {
	if l.next == nil {
		l.next = &List[T]{ val: new, next: nil }
	}

	return new
}

func main() {
	ambar := Pet{"Ambar", "dog", 3}
	bruno := Pet{"Bruno", "dog", 2}
	paco := Pet{"Paco", "cat", 1}
	catira := Pet{"Catira", "cat", 2}

	pets := List[Pet]{ val: ambar, next: &List[Pet]{ val: bruno, next: &List[Pet]{ val: paco, next: nil } } }

	fmt.Println(pets.val)

	newPet := pets.next.next.Push(catira)

	fmt.Println(newPet)
	fmt.Println(pets.next.next.next.val)
}
