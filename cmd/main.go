package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
	"unicode"
)

/*
printing game state
	print word you're guessing
	print hangman state
Derive a word we have to guess
Read user input
	Validate user input
Determine if the letter is a correct guess or not
	if correct, update the guessed letter on the word
	if in-correct, update the hangman state
if word is guessed correctly -> game over, you win
if hangman is complete -> game over , you lose

*/

var inputReader = bufio.NewReader(os.Stdin)

func main() {
	fData, err := os.ReadFile("dict")
	if err != nil {
		fmt.Println("Err is ", err)
		// print any error
	}
	strbuffer := string(fData) // convert read in file to a string

	dictionary := strings.Split(strbuffer, ",")

	targetWord := getRandomWord(dictionary)
	guessedLetters := initializeGuessedWords(targetWord)
	hangmanState := 0

	for !isWordGuessed(targetWord, guessedLetters) && !isHangmanComplete(hangmanState) {

		printGameState(targetWord, guessedLetters, hangmanState)
		input := readInput()
		if len(input) != 1 {
			fmt.Println("Invalid input. Please use letter only ...")
			continue
		}
		if isCorrectGuess(input, targetWord) {
			guessedLetters[rune(input[0])] = true
		} else {

			hangmanState++
		}
	}
	if isWordGuessed(targetWord, guessedLetters) {
		fmt.Println("Congratulations you guessed the word correctly! ...")
	} else {
		fmt.Println("Game Over 2! ...")
	}

}
func initializeGuessedWords(targetWord string) map[rune]bool {

	guessedLetters := map[rune]bool{}
	guessedLetters[unicode.ToLower(rune(targetWord[0]))] = true
	guessedLetters[unicode.ToLower(rune(targetWord[len(targetWord)-1]))] = true

	return guessedLetters
}

func getRandomWord(dictionary []string) string {
	rand.NewSource(time.Now().UnixNano())
	return strings.TrimSpace(dictionary[rand.Intn(len(dictionary))])
}

func printGameState(targetWord string, guessedLetters map[rune]bool, hangmanState int) {
	fmt.Println(getWordGuessingProgress(targetWord, guessedLetters))
	fmt.Println()
	fmt.Println()
	fmt.Println(getHangman(hangmanState))
	fmt.Println()
	fmt.Println()

}
func getWordGuessingProgress(targetWord string, guessedLetters map[rune]bool) string {
	guessedWord := ""
	for _, v := range targetWord {
		if v == ' ' {
			guessedWord += "   "
			continue
		}
		if guessedLetters[unicode.ToLower(v)] {
			guessedWord += fmt.Sprintf("%c", v)
		} else {
			guessedWord += " _ "
		}

	}
	return guessedWord

}

func getHangman(hangmanState int) string {
	file := fmt.Sprintf("%s%d", "states/hangman", hangmanState)

	data, err := os.ReadFile(file)
	if err != nil {
		log.Println("Run into an error printing hang man", hangmanState)
	}
	return string(data)

}

func readInput() string {
	fmt.Print(">")

	input, err := inputReader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	return strings.TrimSpace(input)
}

func isCorrectGuess(input string, targetWord string) bool {
	return strings.Contains(strings.ToLower(targetWord), strings.ToLower(input))
}

func isWordGuessed(targetWord string, guessedLetter map[rune]bool) bool {

	for _, v := range strings.ToLower(targetWord) {
		if !guessedLetter[unicode.ToLower(v)] {
			return false
		}
	}
	return true
}

func isHangmanComplete(hangmanState int) bool {

	return hangmanState >= 9
}
