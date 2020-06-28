<<<<<<< HEAD
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

=======
package main

import "log"

type Masina struct {
	Marca        string
	Model        string
	AnFabricatie int
	Combustibil  string
	//PutereCai int
	PutereKW  float64
	Greutate  float64
	Tractiune string
}

func main() {

	//Conversie kw la cai putere cu formula 1 kw = 1,34102 hp

	Masina1 := Masina{Marca: "Subaru", Model: "Impreza", AnFabricatie: 2010, Combustibil: "Benzina", PutereKW: 198, Greutate: 1833, Tractiune: "integrala"}
	Masina2 := Masina{Marca: "BUGATTI", Model: "CHIRON", AnFabricatie: 2019, Combustibil: "Benzina", PutereKW: 1103, Greutate: 1995, Tractiune: "integrala"}

	log.Println("Aceasta masina are caracteristicile: ", Masina1)
	performanta(Masina1)
	Masina1.tuning()

	log.Println("Aceasta masina are caracteristicile: ", Masina2)
	performanta(Masina2)
	Masina2.tuning()

}

func performanta(power Masina) {

	if power.PutereKW > 150 && (power.PutereKW/power.Greutate > 0.5) {
		log.Println("Bolid")
	} else {
		log.Println("Masina de tatae")

	}
	PutereCai := power.PutereKW * 1.34102
	//am printat mai inatai puterea in KW deoarece cu ea a inceput utilizatorul
	log.Println("Raportul putere greutate in kw putere per kg este: ", power.PutereKW/power.Greutate)
	log.Println("Raportul putere greutate in cai putere per kg este: ", PutereCai/power.Greutate)
}

func (fun Masina) tuning() {
	switch {
	case fun.Tractiune == "fata":
		log.Println("no donuts")

	case fun.Tractiune == "spate":
		log.Println("do donuts!!!")

	case fun.Tractiune == "integrala":
		log.Println("no donuts but volvo safe")
	}
>>>>>>> 9de02d14b64f90522fbba7034faf3f9061fa669a
}
