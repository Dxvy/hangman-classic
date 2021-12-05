package main

import (
	"bufio"
	"log"
	"math/rand"
	"os"
	"time"
)

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
	file.Close()
}
