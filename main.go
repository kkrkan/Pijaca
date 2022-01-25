package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	var kolicina int
	var trosak int
	var namirnica string
	var novci int
	var biranje string
	var shop_balance int = 10000
	var balance int = 1000
	var cijene = [6]int{10,15,13,19,30,40}
	cls := exec.Command("cmd", "/c", "cls")
	cls.Stdout = os.Stdout 
	var menu string = (`

	  	  █▄▀ █▀█ █▄▀ ▄▀█ █▄░█ █▀█ █░█ ▄▀█   █▀█ █ ░░█ ▄▀█ █▀▀ ▄▀█
		    █░█ █▀▄ █░█ █▀█ █░▀█ █▄█ ▀▄▀ █▀█   █▀▀ █ █▄█ █▀█ █▄▄ █▀█
	_____________________________________________________________________________

	[10$] Brašno           [15$] Kruh            [13$] Voda            [19$] Žvake

	                [30$] Kola                           [40$] Kokain

	______________________________________________________________________________

	`)

	fmt.Print(menu)

	fmt.Print("Unesi broj namirnica koje si kupio ->  ")
	fmt.Scanln(&kolicina)
	for i := 0; i <= kolicina; i++ {
		fmt.Print("Unesi ime namjernice broj ",i,":  ")
		fmt.Scanln(&namirnica)
		if(namirnica == "brašno") {
			trosak += cijene[0]
		} else if(namirnica == "kruh") {
			trosak += cijene[1]
		} else if(namirnica == "voda") {
			trosak += cijene[2]
		} else if(namirnica == "žvake") {
			trosak += cijene[3]
		} else if(namirnica == "kola") {
			trosak += cijene[4]
		} else if(namirnica == "kokain") {
			trosak += cijene[5]
		} else {
			fmt.Println(namirnica, " nema u našoj trgovini!")
		}
	}

	cls.Run()
	fmt.Print(menu)

	if(trosak == 0) {
		fmt.Println("Hvala na posjeti marketa, dođite opet!")
	} else {
		fmt.Println("Sve ukupono ste potrošili [",trosak,"$]")
		fmt.Print("Uplati ovdje ->  ")
		fmt.Scanln(&novci)
		if(novci < trosak) {
			fmt.Println("Žao nam je ali niste dosta uplatili!")
		} else if(novci > trosak) {
			fmt.Print(`
			Jeste li sigurni da želite uplatiti više ? 
			DA/NE 
			Odgovor ->  `)
			fmt.Scanln(&biranje)
			if(biranje == "DA") {
				fmt.Println("Hvala, uspješno ste platili račun!")
				balance -= novci
				shop_balance += novci
				fmt.Println("Tvoje balance: ",balance,"$")
			} else if(biranje == "NE") {
				fmt.Println("Ured, evo vam ostatak nazad!")
				balance -= novci
				shop_balance += novci
				fmt.Println("Tvoje balance: ",balance,"$")
			} 
		} else if(novci == trosak) {
			fmt.Println("Uspješno ste platili račun!")
			balance -= novci
			shop_balance += novci
			fmt.Println("Tvoje balance: ",balance,"$")
		}
	}
}
