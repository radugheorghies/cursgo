package main

import "log"

type Person struct {
	Name string
	Age  int
}

func myAgeNextYear(p *Person) {
	p.Age++
}

func (p *Person) myAgeAfter2Years() {
	p.Age += 2
	log.Println("Me after 2 years", p)
}

func main() {
	// recap pointers
	me := &Person{Name: "Radu Gheorghies", Age: 41}
	log.Println(me)
	me.myAgeAfter2Years()
	// myAgeNextYear(me)
	// log.Println("Me after one year", me)
}
