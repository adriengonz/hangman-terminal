package main

import (
	"fmt"
	"time"

)

var hidden_word []string
var used_letters []string

func main() {
	PrintHangmanAscii()
	fmt.Println("Bienvenue sur Hangman !")
	fmt.Println("La partie va bientôt commencer..")
	time.Sleep(4 * time.Second)
	Clear()
	tries := 0
	originalWord := WordPicker(RandomNumber())
	Hidden(originalWord)
	RunHangman(originalWord, tries)
}