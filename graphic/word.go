package graphic

import (
	"hangman/ascii"
	"hangman/utils"
	"hangman/values"
)

func RefreshWordMenu() {
	utils.ClearTerminal()
	utils.Writeln("")
	utils.Writeln("")
	utils.Writeln(ascii.WordSelec)
	utils.Writeln(utils.CenterText("=== Sélection de la liste de mots ===", 65))

	for i, wordFile := range values.WordFiles {
		text := ""
		if i == values.CurrentOption {
			if wordFile.Path == values.CurrentWordFile {
				text = "> " + wordFile.Name + " (actuel)"
				utils.WriteColor(utils.CenterText(text, 65), "green")
			} else {
				text = "> " + wordFile.Name + ""
				utils.WriteColor(utils.CenterText(text, 65), "yellow")
			}
		} else {
			if wordFile.Path == values.CurrentWordFile {
				text = "| " + wordFile.Name + " (actuel)"
			} else {
				text = "| " + wordFile.Name + ""
			}
			utils.Writeln(utils.CenterText(text, 65))
		}
	}

	// Option Retour
	if len(values.WordFiles) == values.CurrentOption {
		utils.WriteColor(utils.CenterText("> Retour", 65), "red")
	} else {
		utils.Writeln(utils.CenterText("| Retour", 65))
	}

	utils.Writeln(utils.CenterText("=====================================", 65))
	utils.Writeln("")

	// Afficher la description du fichier sélectionné
	if values.CurrentOption < len(values.WordFiles) {
		utils.WriteColor(utils.CenterText(values.WordFiles[values.CurrentOption].Description, 65), "cyan")
	}
}
