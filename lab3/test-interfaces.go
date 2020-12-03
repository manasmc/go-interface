package main

import (
	"fmt"
)


// Defining the map to store an integer  and person(underlying string)

type pdb  map[int]person
type mdb  map[int]person


// Defining type person (underlying is struct) to pass the values to db

type person struct{
	fname string
	
}


// defining a method to write dada to the mango db map 

func (m mdb) writem(n int, p person) {
      m[n] = p
      
}

// defining the method to write dada to pg db

func (pg pdb) writem(n int, p person) {
	pg[n] = p
	
}

//defining  a method to read data rom mango db

func (m mdb) readm(n int)person{
	return m[n]
}

//defining  a method to read data rom postgress db
func (pg pdb) readm(n int)person{
	return pg[n]
	 
}


//Defining an interface as the abstact class( allow multiple type for value)

type access interface{
	writem(n int, p person)
	readm(n int)person
	
}

// defining function to save dada in dada base

func save(a access, n int, p person){
	a.writem(n,p)
}

// defining function to save dada in dada base

func read(a access, n int)person{
	return a.readm(n)
}


func main (){

m := mdb{}
p := pdb{}

p1 := person{
	fname: "manas",
	
}

save(m, 1, p1)
save(p, 1, p1)


fmt.Println(read(m, 1))


}