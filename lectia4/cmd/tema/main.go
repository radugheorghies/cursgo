package main

import "fmt"

type produs struct {
	Nume      string
	Pret      int
	Cantitate int
}

func (produsCurent *produs) aprovizioneazaStoc(cantitateaprovizionata int) {
	if cantitateaprovizionata < 0 {
		fmt.Println("Nu se poate aproviziona!")
	} else {
		produsCurent.Cantitate += cantitateaprovizionata
		fmt.Println("Am aprovizionat stocul de: ", produsCurent.Nume, "cantitate disponibila:", produsCurent.Cantitate, ", pret: ", produsCurent.Pret)
	}
}

func (produsCurent *produs) cumparaProdus(cantitatecumparata int) {
	pretTotalProdus := 0
	if cantitatecumparata > 0 {

		if produsCurent.Cantitate >= cantitatecumparata {
			pretTotalProdus = cantitatecumparata * produsCurent.Pret
			produsCurent.Cantitate -= cantitatecumparata
			fmt.Println("Ai achizitionat produsul:", produsCurent.Nume, "pretul achizitiei", pretTotalProdus, "Se actualizeaza stocul de produse disponibile", produsCurent.Cantitate, ".")
		} else {
			fmt.Println("Nu avem in stoc cantitatea suficienta de ", produsCurent.Nume, ". Cantitatea disponibila in stoc:", produsCurent.Cantitate)
		}
	} else {
		fmt.Println("Nu se poate achizitiona produsul!")
	}
}

func main() {
	stocProduse := [3]produs{
		{Nume: "Masti chirurgicale",
			Pret: 3,
		},

		{Nume: "Alcool sanitar",
			Pret:      8,
			Cantitate: 190},

		{Nume: "Dezinfectant",
			Pret:      12,
			Cantitate: 200},
	}

	stocProduse[2].aprovizioneazaStoc(-9)
	stocProduse[1].cumparaProdus(180)
	stocProduse[1].cumparaProdus(-11)
	stocProduse[0].cumparaProdus(2)

}
