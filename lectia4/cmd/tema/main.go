package main

import "log"

type Square struct {
	width int
}

func (sQr *Square) area() int {
	return sQr.width * sQr.width
}

func (sQr *Square) perimeter() int {
	return sQr.width * 4
}

func sQrAsessment(size *Square) {
	if size.width > 10 {
		log.Println("Big Square")
	} else {
		log.Println("Small Square")
	}
}

func (squ *Square) sQrSwitch() {

	switch {
	case squ.width < 5:
		log.Println("Switch less than 5")
	case squ.width == 5:
		log.Println("Switch equals 5")
	case squ.width > 5:
		log.Println("Switch bigger than 5")
	}
}

func main() {
	sQr := &Square{width: 12}
	log.Println("area :", sQr.area())
	log.Println("perimeter :", sQr.perimeter())
	sQrAsessment(sQr)
	sQr.sQrSwitch()

}
