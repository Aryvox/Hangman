package graphic

import (
	"hangman/ascii"
	"hangman/utils"
	"hangman/values"
)

func RefreshCreditsMenu() {
	utils.ClearTerminal()
	utils.Writeln(ascii.Me)

	switch values.CurrentOption {
	case 0:
		utils.WriteColor(utils.CenterText("> Quitter                 |", 65), "red")
		utils.Writeln(utils.CenterText("===========================", 65))
		break
	}
}
