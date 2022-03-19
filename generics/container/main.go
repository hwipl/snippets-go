package main

import (
	"log"
)

// Container is a generic container
type Container[T comparable] struct {
	c []T
}

// Add adds element to the container
func (c *Container[T]) Add(element T) {
	for _, e := range c.c {
		if e == element {
			return
		}
	}
	c.c = append(c.c, element)
}

// Remove removes element from the container
func (c *Container[T]) Remove(element T) {
	for i, e := range c.c {
		if e == element {
			c.c = append(c.c[:i], c.c[i+1:]...)
			return
		}
	}
}

// List lists the elements in the container
func (c *Container[T]) List() []T {
	return c.c
}

// NewContainer returns a new Container
func NewContainer[T comparable]() *Container[T] {
	return &Container[T]{}
}

func main() {
	s := NewContainer[string]()
	s.Add("test1")
	s.Add("test2")
	s.Add("test3")
	s.Remove("test2")
	log.Println(s.List())
}
