package main

import "log"

type Vehicol struct {
	nrRoti      int
	consum      int
	combustibil int
}

func (v *Vehicol) parcurgeDistanta(d int) bool {

	if d*v.consum < v.combustibil {
		v.combustibil -= d * v.consum
		return true
	}

	return false
}

func (v *Vehicol) alimenteaza(cantitate int) {
	v.combustibil += cantitate
}

func (v Vehicol) tipVehicol() string {
	switch v.nrRoti {
	case 0:
		return "Not a vehicol"
	case 1:
		return "Monociclu"
	case 2:
		return "Motocicleta"
	case 3:
		return "Tricicleta"
	case 4:
		return "Masina"
	default:
		return "Ceva mai mare:))"
	}
}

func main() {

	v0 := Vehicol{0, 10, 100}
	log.Println(v0.tipVehicol())
	log.Println(v0.parcurgeDistanta(9))
	log.Println(v0.parcurgeDistanta(3))
	v0.alimenteaza(100)
	log.Println(v0.parcurgeDistanta(3))

	v1 := Vehicol{4, 7, 150}
	log.Println(v1.tipVehicol())
	log.Println(v1.parcurgeDistanta(9))

}