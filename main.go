package main

import (
	"fmt"
)

type mdb map[int]person

type person struct {
	fname string
}

func (m mdb) save(n int, p person) {
	m[n] = p
}

func (m mdb) retrieve(n int) person {
	return m[n]

}

type accessor interface {
	save(n int, p person)
	retrieve(n int) person
}

func put(a accessor, n int, p person) {
	a.save(n, p)
}

func get(a accessor, n int) person {
	return a.retrieve(n)
}

func main() {
	data := mdb{}

	p1 := person{
		fname: "manas",
	}

	put(data, 1, p1)
	fmt.Println(get(data, 1))

}
