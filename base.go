package main

import (
	"fmt"
	"time"
)

func RunHangman(word string, try int) {
	presentInWord := false
	wordNotRevealed := false
	draw_hangman(try)
	letter := ""
	fmt.Println(letter)

	fmt.Println("Déjà utilisé(s) : ", used_letters)
	fmt.Println("\nEntrez votre choix :")
	fmt.Scanln(&letter)
	for _, char := range used_letters {
		if letter == string(char) {
			fmt.Println("Essayez autre chose, vous avez déjà essayé ceci !")
			RunHangman(word, try)
		}
	}
	used_letters = append(used_letters, letter)

	if len(letter) < 1 || len(letter) > 1 {
		fmt.Println("Impossible, entrez une seule lettre.")
		RunHangman(word, try)
	} else {
		for i := 0; i < len(word); i++ {
			if letter == string(word[i]) {
				presentInWord = true
			}
		}
		if presentInWord {
			fmt.Println("La lettre", letter, "est dans le mot")
			for index, char := range word {
				if letter == string(char) {
					hidden_word[index] = letter
				}
			}
		} else {
			fmt.Println("La lettre", letter, "n'est pas dans le mot")
			if try < 9 {
				try++
				fmt.Println("Il ne vous reste plus que", 10-try, "essais")
			} else {
				fmt.Println("Perdu ! Vous n'avez plus d'essais disponibles")
				fmt.Println("Le mot était", word)
				time.Sleep(3 * time.Second)
				Exit()
			}
		}
		Clear()
		fmt.Println(hidden_word)
	}
	for index, _ := range word {
		if hidden_word[index] == "_" {
			wordNotRevealed = true
			RunHangman(word, try)
			break
		}
	}
	if wordNotRevealed == false {
		fmt.Println("Felicitations ! Vous avez déviné !")
		time.Sleep(3 * time.Second)
		Exit()
	}
}