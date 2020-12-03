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

// defining a struct to the accessor interface
type service struct {
	x accessor
}

func (ps service) get(n int) person {
	return ps.x.retrieve(n)
}

func (ps service) put(n int, p person) {
	ps.x.save(n, p)
}

func main() {

	data := mdb{}

	ps := service{
		x: mdb{},
	}
	// write data to the map and type person and underlying type is struct
	p1 := person{
		fname: "manas",
	}

	p2 := person{
		fname: "dil",
	}
	// Method 1
	put(data, 1, p1)
	// Method 2
	put(ps.x, 2, p2)
	put(ps.x, 1, p1)

	fmt.Println(get(data, 1))
	fmt.Println("-------------------------------")
	fmt.Println(ps.get(1))
	fmt.Println(ps.get(2))

}
