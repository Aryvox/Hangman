package graphic

import (
	"hangman/ascii"
	"hangman/utils"
	"hangman/values"
)

func RefreshMainMenu() {
	utils.ClearTerminal()

	utils.Writeln("")
	utils.Writeln("")
	utils.Writeln(ascii.MenuTitle)

	utils.Writeln(utils.CenterText("====== Jeu du Pendu !======", 65))
	switch values.CurrentOption {
	case 0:
		utils.WriteColor(utils.CenterText("> Jouer au Pendu          |", 65), "green")
		utils.Writeln(utils.CenterText("| Crédits                 |", 65))
		utils.Writeln(utils.CenterText("| Changer les mots        |", 65))
		utils.Writeln(utils.CenterText("| Quitter                 |", 65))
		utils.Writeln(utils.CenterText("===========================", 65))
		utils.Writeln("")
		utils.Writeln("")
		utils.Writeln(ascii.Pendu)
		break
	case 1:
		utils.Writeln(utils.CenterText("| Jouer au Pendu          |", 65))
		utils.WriteColor(utils.CenterText("> Crédits                 |", 65), "yellow")
		utils.Writeln(utils.CenterText("| Changer les mots        |", 65))
		utils.Writeln(utils.CenterText("| Quitter                 |", 65))
		utils.Writeln(utils.CenterText("===========================", 65))
		utils.Writeln("")
		utils.Writeln("")
		utils.Writeln(ascii.Credits)
		break
	case 2:
		utils.Writeln(utils.CenterText("| Jouer au Pendu          |", 65))
		utils.Writeln(utils.CenterText("| Crédits                 |", 65))
		utils.WriteColor(utils.CenterText("> Changer les mots        |", 65), "yellow")
		utils.Writeln(utils.CenterText("| Quitter                 |", 65))
		utils.Writeln(utils.CenterText("===========================", 65))
		break
	case 3:
		utils.Writeln(utils.CenterText("| Jouer au Pendu          |", 65))
		utils.Writeln(utils.CenterText("| Crédits                 |", 65))
		utils.Writeln(utils.CenterText("| Changer les mots        |", 65))
		utils.WriteColor(utils.CenterText("> Quitter                 |", 65), "red")
		utils.Writeln(utils.CenterText("===========================", 65))
		utils.Writeln("")
		utils.Writeln("")
		utils.Writeln(ascii.ExitDoor)
		break
	}
}
