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

//Initialisation de certaines variables globales
var array []string
var hiddenWord []string
var trial []string
var attempt = 10
var isFind = 0 //Compteur des lettres trouvées
var chosenWord string
var file string

func main() { //Programme principal
	hub()
	if len(os.Args) > 1 {
		file = os.Args[1]
	} else {
		file = "./Ressources/words.txt"
	}
	fileScanner := createScanner(file)
	array = getWords(fileScanner, array)
	rand.Seed(time.Now().UnixNano()) //Initialisation de l'aléatoire
	nombre := rand.Intn(len(array))
	chosenWord = array[nombre]    //Choix du mot à trouver aléatoirement
	hiddenWord = hideToFindWord() //Création du mot caché
	displayedLetters := showToFindLetters()
	isFind = displayedLetters //Ajout des lettres trouvées au compteur
	start()
}

func start() { //Programme de lancement du jeu
	for attempt != 0 { //Boucle d'entrée du jeu'
		verifLettersUsed := 0
		printToFindWord()
		displayHangman(attempt, "./Ressources/Hangman.txt")
		displayProposition(trial)
		letter := trySomething()
		for i := range trial {
			if letter == trial[i] {
				verifLettersUsed++
			}
		}
		if verifLettersUsed != 0 { //Vérification si proposition déjà faites
			println("You already tried that !!!")
			continue
		}
		if verifLettersUsed == 0 { //Ajouts aux propositions passées
			trial = append(trial, letter)
		}
		println("Choosed : ", letter)
		verifGoodProposition := 0
		for i := 0; i < len(chosenWord); i++ { //Vérification si la lettre est présente dans le mot
			if letter == string(chosenWord[i]) && string(hiddenWord[i]) == "_" {
				hiddenWord[i] = string(chosenWord[i])
				isFind++
			} else {
				verifGoodProposition++
			}
		}

		if letter == chosenWord { //Vérification si le mot a été trouvé (via une proposition de mot)
			println("\nCongrats !!! You find the word : ", chosenWord, " with ", attempt, " attempts left")
			break
		}
		if verifGoodProposition == len(chosenWord) { //Modification du compteur d'essai en cas d'échec
			if len(letter) == 1 {
				attempt--
			} else {
				attempt -= 2
				if attempt < 0 {
					attempt = 0
				}
			}
			println("Not present in the word, ", attempt, " attempts remaining")
		}
		if isFind == len(chosenWord) { //Vérification si le mot a été trouvé (via une proposition de lettre)
			println("\nCongrats !!! You find the word : ", chosenWord, " with ", attempt, " attempts left")
			break
		}
		if attempt == 0 { //Vérification s'il reste des tentatives
			loose()
		}
		println()
	}
}

func getWords(fileScanner *bufio.Scanner, array []string) []string { //Programme de récupération des mots du fichier txt
	for fileScanner.Scan() {
		array = append(array, fileScanner.Text())
	}
	return array
}

func showToFindLetters() int { //Choix des lettres affichées dès le début
	lettersToDisplay := (len(hiddenWord) / 2) - 1
	var displayedLetters int
	for i := 0; i < lettersToDisplay; i++ {
		index := rand.Intn(len(hiddenWord))
		if hiddenWord[index] == "_" {
			displayedLetters++
		}
		hiddenWord[index] = string(chosenWord[index])
	}
	return displayedLetters
}

func hideToFindWord() []string { //Programme pour créer le mot caché
	for i := 0; i < len(chosenWord); i++ {
		hiddenWord = append(hiddenWord, "_")
	}
	return hiddenWord
}

func printToFindWord() { //Affichage du mot à découvrir
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

func trySomething() string { //Propositions faites par le joueur
	println("Make a proposition : ")
	lettre := input()
	return lettre
}

func displayProposition(proposition []string) { //Affichage des propositions déjà faites
	println("Here the attempts already tried :")
	for i := range proposition {
		print(proposition[i] + " ")
	}
	println()
}

func createScanner(nomFichier string) *bufio.Scanner { //Programme de création d'un scanner
	file, err := os.Open(nomFichier)
	manageError(err)
	fileScanner := bufio.NewScanner(file)
	return fileScanner
}

func displayHangman(tentative int, nomFichier string) { //Programme pour afficher le hangman selon le nombre de tentatives restantes
	println()
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
	for fileScanner.Scan() {
		if i >= debut && i <= fin {
			println(fileScanner.Text())
		}
		i++
	}
}

func manageError(err error) { //Programme en cas d'erreur (scan)
	if err != nil {
		log.Fatalf("Error when opening file: %s", err)
	}
}

func loose() { //Programme en cas d'échec
	file, err := os.Open("./Ressources/hangman.txt")
	manageError(err)
	fileScanner := bufio.NewScanner(file)
	i := 0
	for fileScanner.Scan() {
		if i >= 72 && i <= 79 {
			println(fileScanner.Text())
		}
		i++
	}
	println("Sorry you loose bro, try again !!!!")
}

func hub() { //Menu de base
	println("}===================================================================================={")
	println("}                                                                                    {")
	println("}            _    _              _   _   _____  __  __              _   _            {")
	println("}-----------| |  | |     /\\     | \\ | | / ____||  \\/  |     /\\     | \\ | |-----------{")
	println("}-----------| |__| |    /  \\    |  \\| || |  __ | \\  / |    /  \\    |  \\| |-----------{")
	println("}-----------|  __  |   / /\\ \\   | . ` || | |_ || |\\/| |   / /\\ \\   | . ` |-----------{")
	println("}-----------| |  | |  / ____ \\  | |\\  || |__| || |  | |  / ____ \\  | |\\  |-----------{")
	println("}-----------|_|  |_| /_/    \\_\\ |_| \\_| \\_____||_|  |_| /_/    \\_\\ |_| \\_|-----------{")
	println("}                                                                                    {")
	println("}                                                                                    {")
	println("}      Welcome to the Hangman game !                                                 {")
	println("}                                                                                    {")
	println("}      The rules are the followings :                                                {")
	println("}      - You're gonna have to find a hidden word                                     {")
	println("}      - You can try to find it letter by letter, or by tapping directly a word      {")
	println("}      - You have a total of 10 attempts                                             {")
	println("}      - If the letter you tried is false, you lose 1 attempt                        {")
	println("}      - If the word you tried is false, you lose 2 attempts                         {")
	println("}      - Some letters are already displayed                                          {")
	println("}      - Displayed letters can be twice in the word, so you have to try them         {")
	println("}                                                                                    {")
	println("}      You win when the word is totally discovered                                   {")
	println("}      If your attempts fall to 0, you will loose and hanged                         {")
	println("}                                                                                    {")
	println("}      Good luck                                                                     {")
	println("}                                                                                    {")
	println("}===================================================================================={")
	println()
	println()
	println()
	println()
}
