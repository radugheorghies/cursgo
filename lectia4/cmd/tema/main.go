// Am folosit ca si model de calcul http://www.validare.ro/validare-cnp.html si https://en.wikipedia.org/wiki/National_identification_number#Romania

package main

import ( "log"; "strconv"; "time")
 
type DatePersonale struct {
	Nume1, Nume2, Nume3, Prenume1, Prenume2, Prenume3, Adresa, Judet, Sex, Localitate string
	CNP int
}

func (DatePersonale1 DatePersonale) CheckLenCNP() { 
	y := ""
	if DatePersonale1.CNP > 1000000000000 {y = "CNP-ul are numarul corect de cifre"} else {y = "CNP-ul nu are numarul corect de cifre"}
	log.Println(y)
}

func Sex(C1 int) string{
	if C1 == 9 {
		return "Sex necunoscut (cetatean strain)"
	} else if C1%2 == 1 {
		return "Sex: M"
	} else {
		return "Sex: F"
	}
}

func ValidareData(CNPA [13]int) string{
	//log.Println(CNPA[1:7])
	y :=""
	x := ""
	if CNPA[0] == 1 || CNPA[0] == 2 { x = "19"
	} else if CNPA[0] == 5 || CNPA[0] == 6 { x ="20"
	} else if CNPA[0] == 3 || CNPA[0] == 4 { x ="18"
	} else {log.Println("Imposibilitate validare data ")}

	x = x + strconv.Itoa(CNPA[1]) + strconv.Itoa(CNPA[2]) + "-" + strconv.Itoa(CNPA[3]) + strconv.Itoa(CNPA[4]) + "-" + strconv.Itoa(CNPA[5]) + strconv.Itoa(CNPA[6]) 
	//log.Println(x)
	// validare data luna<=12. numarul corect de zile intr-o luna inclusiv ani bisecti	
	d, err := time.Parse("2006-01-02", x)
	if err != nil {
		y = "Data Invalida"
} else if d.Year() > 0001 {
    y = "Data valida"
}
	
	return y

}

func ValidareJudet(CNPA [13]int) string{
	y := ""
	x := CNPA[7]*10 + CNPA[8]

	// tabelul cu judetele: https://ro.wikipedia.org/wiki/Cod_numeric_personal
	if x>0 && x<53 { 
		y = "Judet Valid"
	} else {y = "!! Judet INVALID !!"}
	return y

}

func ValidareCifraControl(CNPA [13]int) string{
	x := 0
	a := [13]int {2,7,9,1,4,6,3,5,8,2,7,9}
	y := ""
	for i := 0; i <=12; i++ {
		x += CNPA[i]*a[i]
		//log.Println("validare x ", x)
	}
	rest11 := x%11
	if rest11 == 10 && CNPA[12] == 1 { y="Cifra de control corecta"
} else if rest11 == CNPA[12] { y="Cifra de control corecta"
} else { y="!! Cifra de control este INCORECTA !!"
}


	return y
}

func main() {
	var CNPA [13]int
// CNP-ul a fost generat aleatoriu pe http://redoxwap.freehostia.com/cnpgenerator.php
	DatePersonale1 := DatePersonale{Nume1 :"Stefan", Nume2 :"Cel", Nume3 :"Mic", Prenume1 :"Doar", Prenume2 :"In", Prenume3 :"Nume", Adresa :"un pic mai confidential", Judet : "B", Sex : "M", Localitate:"Bucuresti", CNP: 6021025146120}
	CNP2:= DatePersonale1.CNP

	for i :=0; i<=12; i++ {
		x := 12-i
		CNPA[x] = CNP2 % 10
		CNP2 = (int)(CNP2/10)
	}	
	
	// Verificari pe baza CNP-ului
	DatePersonale1.CheckLenCNP()
	log.Println(Sex(CNPA[0]))
	log.Println(ValidareData(CNPA))
	log.Println(ValidareJudet(CNPA))
	log.Println(ValidareCifraControl(CNPA))
	
}



