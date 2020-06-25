package main

import (
	"fmt"
)

type Stringer interface {
	String() string	
}

type user struct {
    name string
}

func (n *user) String() string {
	return n.name
}

func toString(type T Stringer)(item T) string {
	return item.String()
}

func main() {
	s := toString(&user{name: "Bradford"})
	fmt.Println(s)
}