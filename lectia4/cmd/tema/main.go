// Tema lectia 4: O structura de date la care sa atasam 2 metode,
// in metodele respective sa folosim instructiunile if, for, switch,
// in functia main sa definiti o variabila de tipul structurii de date,
// si sa apelati cele doua metode.

package main

import "fmt"

type Calculator struct {
	Producator   string
	Model        string
	AnFabricatie int
	TipRAM       string
	CantitateRAM int
	Procesor     string
}

func performanta(cal Calculator) {

	if cal.TipRAM == "DDR4" && (cal.Procesor == "intel i5" || cal.Procesor == "intel i7") {
		fmt.Println("Calculatorul este performant")
	} else {
		fmt.Println("Calculatorul este depasit moral")
	}
}

func (cal Calculator) necesitaUpgrade() {
	switch {
	case cal.CantitateRAM < 8:
		fmt.Println("Adauga mai multa memorie RAM")

	case cal.CantitateRAM > 8:
		fmt.Println("Acest calculator nu necesita upgrade la memoria RAM")

	case cal.Procesor == "intel i3":
		fmt.Println("Verificati ce tip de socket are procesorul, apoi verifica daca gasesiti unul mai performant")

	case cal.Procesor == "intel i5" || cal.Procesor == "intel i7":
		fmt.Println("Puterea de calcul a acestui procesor este mare")

	}
}

func main() {
	calculator1 := Calculator{Producator: "Asus", Model: "Gaming", AnFabricatie: 2019, TipRAM: "DDR4", CantitateRAM: 8, Procesor: "intel i5"}
	fmt.Println("Acest calculator are urmatoarea configuratie: ", calculator1)

	performanta(calculator1)
	calculator1.necesitaUpgrade()

}
