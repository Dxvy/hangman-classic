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
		println(test3, " ", len(motChoisit))
		println("Tentative restante :", tentative)
		lettre := choixLettre()
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
