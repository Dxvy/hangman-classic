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

var array []string
var hiddenWord []string
var trial []string
var attempt = 10
var isFind = 0
var chosenWord string

func main() {
	fileScanner := createScanner(os.Args[1])
	array = getWords(fileScanner, array)
	rand.Seed(time.Now().UnixNano())
	nombre := rand.Intn(len(array))
	chosenWord = array[nombre]
	hiddenWord = hideToFindWord()
	displayedLetters := showToFindLetters()
	isFind = displayedLetters
	start()
}

func start() {
	println()
	println("Good Luck, you have ", attempt, " attempts")
	for attempt != 0 {
		verifLettersUsed := 0
		println()
		//println("Tentative restante :", attempt)
		println()
		printToFindWord()
		displayHangman(attempt, "./Ressources/Hangman.txt")
		letter := trySomething()
		for i := range trial {
			if letter == trial[i] {
				verifLettersUsed++
			}
		}
		if verifLettersUsed != 0 {
			println("Cette letter a déjà été faite !!!")
			println()
			continue
		}
		if verifLettersUsed == 0 {
			trial = append(trial, letter)
		}
		println()
		println("Choosed : ", letter)
		displayProposition(trial)
		verifGoodProposition := 0
		for i := 0; i < len(chosenWord); i++ {
			if letter == string(chosenWord[i]) && string(hiddenWord[i]) == "_" {
				hiddenWord[i] = string(chosenWord[i])
				isFind++
			} else {
				verifGoodProposition++
			}
		}
		if letter == chosenWord {
			println("\nCongrats !!! You find the word : ", chosenWord, " in ", attempt, " attempts")
			println()
			break
		}
		if verifGoodProposition == len(chosenWord) {
			if len(letter) == 1 {
				attempt--
			} else {
				attempt -= 2
				if attempt < 0 {
					attempt = 0
				}
			}
			println()
			println("Not present in the word, ", attempt, " attempts remaining")
		}
		if isFind == len(chosenWord) {
			println("\nCongrats !!! You find the word : ", chosenWord, " in ", attempt, " attempts")
			println()
			break
		}
		if attempt == 0 {
			loose()
		}
	}
}

func getWords(fileScanner *bufio.Scanner, array []string) []string {
	for fileScanner.Scan() {
		array = append(array, fileScanner.Text())
	}
	return array
}

func showToFindLetters() int {
	lettresAffichees := (len(hiddenWord) / 2) - 1
	for i := 0; i < lettresAffichees; i++ {
		index := rand.Intn(len(hiddenWord))
		hiddenWord[index] = string(chosenWord[index])
	}
	return lettresAffichees
}

func hideToFindWord() []string {
	for i := 0; i < len(chosenWord); i++ {
		hiddenWord = append(hiddenWord, "_")
	}
	return hiddenWord
}

func printToFindWord() {
	for i := range hiddenWord {
		print(hiddenWord[i])
	}
}

func input() string { // Fonction pour récupérer le texte écrit dans le cmd et l'utiliser
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("-> ")
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\r\n", "", -1)
	return text
}

func trySomething() string {
	println("Make a proposition : ")
	lettre := input()
	return lettre
}

func displayProposition(proposition []string) {
	println()
	println("Here the attempts already tried :")
	for i := range proposition {
		print(proposition[i] + " ")
	}
	println()
}

func createScanner(nomFichier string) *bufio.Scanner {
	file, err := os.Open(nomFichier)
	manageError(err)
	fileScanner := bufio.NewScanner(file)
	return fileScanner
}

func displayHangman(tentative int, nomFichier string) {
	file, err := os.Open("./Ressources/hangman.txt")
	manageError(err)
	fileScanner := bufio.NewScanner(file)
	i := 0
	var debut int
	var fin int
	switch tentative {
	case 9:
		debut = 0
		fin = 7
	case 8:
		debut = 8
		fin = 15
	case 7:
		debut = 16
		fin = 23
	case 6:
		debut = 24
		fin = 31
	case 5:
		debut = 32
		fin = 39
	case 4:
		debut = 40
		fin = 47
	case 3:
		debut = 48
		fin = 55
	case 2:
		debut = 56
		fin = 63
	case 1:
		debut = 64
		fin = 71
	}
	println()
	for fileScanner.Scan() {
		if i >= debut && i <= fin {
			println(fileScanner.Text())
		}
		i++
	}
}

func manageError(err error) {
	if err != nil {
		log.Fatalf("Error when opening file: %s", err)
	}
}

func loose() {
	file, err := os.Open("./Ressources/hangman.txt")
	manageError(err)
	fileScanner := bufio.NewScanner(file)
	i := 0
	println()
	for fileScanner.Scan() {
		if i >= 72 && i <= 79 {
			println(fileScanner.Text())
		}
		i++
	}
	println("Sorry you loose bro, try again !!!!")
	println()
}

func win() {

}
