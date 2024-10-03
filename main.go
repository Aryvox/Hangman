package main

import (
	"hangman/actions"
	"hangman/graphic"
	"hangman/utils"
	"hangman/values"
	"log"

	"github.com/mattn/go-tty"
)

func main() {
	PressF11()
	utils.ClearTerminal()
	graphic.RefreshMainMenu()
	StartListening()
}

func StartListening() {
	tty, err := tty.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer tty.Close()
	for {
		r, err := tty.ReadRune()
		if err != nil {
			log.Fatal(err)
		}
		if r == 65 || r == 66 { // 65 = fleche haut & 66 = fleche bas
			var maxIndex int
			var funcType func()
			switch values.CurrentPage {
			case "main_menu":
				maxIndex = 3
				funcType = graphic.RefreshMainMenu
			case "credits":
				maxIndex = 1
				funcType = graphic.RefreshCreditsMenu
			case "word_menu": // Changed from "mot" to "word_menu"
				maxIndex = len(values.WordFiles) // Nombre d'options + le bouton retour
				funcType = graphic.RefreshWordMenu
			default:
				continue
			}
			if r == 65 {
				if values.CurrentOption == 0 {
					values.CurrentOption = maxIndex
				} else {
					values.CurrentOption--
				}
			} else if r == 66 {
				if values.CurrentOption == maxIndex {
					values.CurrentOption = 0
				} else {
					values.CurrentOption++
				}
			}
			funcType()
		} else {
			if r == 13 { // 13 = touche entrée...
				switch values.CurrentPage {
				case "main_menu":
					actions.MenuExec()
				case "credits":
					actions.CreditsExec()
				case "word_menu": // Changed from "mot" to "word_menu"
					actions.WordExec()
				default:
					continue
				}
			}
		}
	}
}
