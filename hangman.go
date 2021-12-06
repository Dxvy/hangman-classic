package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

func Input() string { // Fonction pour récupérer le texte écrit dans le cmd et l'utiliser
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("-> ")
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\r\n", "", -1)
	return text
}

func choixLettre() string {
	println("Choisis une lettre (en minuscule) : ")
	lettre := Input()
	return lettre
}

func main() {
	var tableau []string
	var motCache []string
	tentative := 10
	test3 := 0

	// open the file
	file, err := os.Open("./Ressources/words.txt")
	//handle errors while opening
	if err != nil {
		log.Fatalf("Error when opening file: %s", err)
	}
	fileScanner := bufio.NewScanner(file)
	// read line by line
	for fileScanner.Scan() {
		tableau = append(tableau, fileScanner.Text())
	}
	// handle first encountered error while reading
	if err := fileScanner.Err(); err != nil {
		log.Fatalf("Error while reading file: %s", err)
	}
	//Lecture à l'indice de la ligne
	println(len(tableau))
	rand.Seed(time.Now().UnixNano())
	nombre := rand.Intn(len(tableau))
	println(nombre)
	motChoisit := tableau[nombre]
	println(motChoisit)
	var proposition []string
	for i := 0; i < len(motChoisit); i++ {
		motCache = append(motCache, "_")
	}
	for i := range motCache {
		print(motCache[i])
	}
	println()
	lettresAffichees := (len(motCache) / 2) - 1
	for i := 0; i < lettresAffichees; i++ {
		index := rand.Intn(len(motCache))
		motCache[index] = string(motChoisit[index])
	}
	test3 = lettresAffichees
	for i := range motCache {
		print(motCache[i])
	}
	for tentative != 0 {
		println()
		affichageJose(tentative)
		test4 := 0
		println()
		println(test3, " ", len(motChoisit))
		println("Tentative restante :", tentative)
		lettre := choixLettre()
		for i := range proposition {
			if lettre == proposition[i] {
				test4++
			}
		}
		if test4 != 0 {
			println("Cette proposition a déjà été faite !!!")
			continue
		}
		if test4 == 0 {
			proposition = append(proposition, lettre)
		}
		for i := range proposition {
			println(proposition[i])
		}
		println("la lettre choisit est : ", lettre)
		test2 := 0
		for i := 0; i < len(motChoisit); i++ {
			if lettre == string(motChoisit[i]) && string(motCache[i]) == "_" {
				motCache[i] = string(motChoisit[i])
				test3++
			} else {
				test2++
			}
		}
		println("longueur du choix : ", len(lettre))
		if test2 == len(motChoisit) {
			if len(lettre) == 1 {
				tentative--
			} else {
				tentative -= 2
			}
		}
		if test3 == len(motChoisit) {
			println("\ngg Hitler")
			break
		}
		for i := range motCache {
			print(motCache[i])
		}
	}
}

func affichageJose(tentative int) {
	file, err := os.Open("./Ressources/hangman.txt")
	if err != nil {
		log.Fatalf("Error when opening file: %s", err)
	}
	fileScanner := bufio.NewScanner(file)
	i := 0
	var debut int
	var fin int
	if tentative == 9 {
		debut = 0
		fin = 7
	}
	if tentative == 8 {
		debut = 8
		fin = 15
	}
	if tentative == 7 {
		debut = 16
		fin = 23
	}
	if tentative == 6 {
		debut = 24
		fin = 31
	}
	if tentative == 5 {
		debut = 32
		fin = 39
	}
	if tentative == 4 {
		debut = 40
		fin = 47
	}
	if tentative == 3 {
		debut = 48
		fin = 55
	}
	if tentative == 2 {
		debut = 56
		fin = 63
	}
	if tentative == 1 {
		debut = 64
		fin = 71
	}
	if tentative == 0 {
		debut = 72
		fin = 79
	}
	for fileScanner.Scan() {
		if i >= debut && i <= fin {
			println(fileScanner.Text())
		}
		i++
	}
}
