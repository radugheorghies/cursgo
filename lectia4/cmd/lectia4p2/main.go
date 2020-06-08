package main

import "log"

func main() {
	// x := -3
	// if x > -1 {
	// 	// executa,m daca este indeplinita conditia
	// 	log.Println("Da, x este mai mare decat -1")
	// } else {
	// 	// aici se executa cod daca nu este indeplinita consitia
	// 	log.Println("NU, x nu este mai mare decat -1")
	// }

	// age := make(map[string]int)
	// age["radu"] = 41
	// age["marius"] = 27

	// if v, ok := age["vasile"]; ok {
	// 	log.Println("age:", v)
	// 	// ...
	// } else {
	// 	log.Println("Not known.")
	// }

	// for

	// i := 0

	// for {
	// 	log.Println(i)
	// 	i++
	// 	if i >= 20 {
	// 		break
	// 	}
	// }

	// switch

	i := 4
	j := 0

	switch {
	case i == 0:
		log.Println("i este 0")
		break
	case i == 1:
		log.Println("i este 1")
	case i < 5:
		log.Println("i este < 5")
	case j == 0:
		log.Println("j este 0")
	}

}
