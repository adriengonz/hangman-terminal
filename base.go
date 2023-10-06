package main

import (
	"fmt"
)

func RunHangman(word string, try int) {
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
		presentInWord := false
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
			fmt.Printf("La lettre %v n'est pas dans le mot\n", letter)
			if try < 9 {
				try++
				fmt.Println("Il ne vous reste plus que", 10-try, "essais")
			}
		}
	}
}