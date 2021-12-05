package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	var tableau []string

	// open the file
	file, err := os.Open("words.txt")

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
	fmt.Println(tableau[5])

	file.Close()
}
