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
	test := []rune(lettre)
	if len(lettre) > 1 || test[0] < 'a' || test[0] > 'z' {
		println("On a dit une lettre (en minuscule) ...")
		return choixLettre()
	}
	return lettre
}

func main() {
	var tableau []string
	var motCache []string
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
	for i := range motCache {
		print(motCache[i])
	}
	println()
	lettre := choixLettre()
	println("la lettre choisit est : ", lettre)

	for i := 0; i < len(motChoisit); i++ {
		if lettre == string(motChoisit[i]) {
			println("caca")
		} else {
			println("prout")
		}
	}
	file.Close()
}
