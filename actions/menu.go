package actions

import (
	"hangman/game"
	"hangman/graphic"
	"hangman/utils"
	"hangman/values"
	"os"
	"syscall"
	"time"
)

var (
	user32               = syscall.NewLazyDLL("user32.dll")
	procKeybdEvent       = user32.NewProc("keybd_event")
	VK_F11          byte = 0x7A
	KEYEVENTF_KEYUP      = 0x0002
)

func PressF11() {
	procKeybdEvent.Call(uintptr(VK_F11), 0, 0, 0)
	time.Sleep(100 * time.Millisecond)
	procKeybdEvent.Call(uintptr(VK_F11), 0, uintptr(KEYEVENTF_KEYUP), 0)
}

func MenuExec() {
	switch values.CurrentOption {
	case 0:
		utils.ClearTerminal()
		game.StartGame()
	case 1:
		values.CurrentPage = "credits"
		utils.ClearTerminal()
		graphic.RefreshCreditsMenu()
	case 2:
		values.CurrentPage = "word_menu" // Changed from "mot" to "word_menu"
		values.CurrentOption = 0         // Reset l'option sélectionnée
		utils.ClearTerminal()
		graphic.RefreshWordMenu()
	case 3:
		PressF11()
		utils.ClearTerminal()
		os.Exit(0)
	}
}

func CreditsExec() {
	switch values.CurrentOption {
	case 0:
		values.CurrentOption = 0
		values.CurrentPage = "credits"
		values.CurrentOptionMax = 0
		utils.ClearTerminal()
		graphic.RefreshMainMenu()
	}
}

func WordExec() {
	if values.CurrentOption < len(values.WordFiles) {
		// Si on n'a pas sélectionné "Retour"
		values.CurrentWordFile = values.WordFiles[values.CurrentOption].Path
		utils.WriteColorLn(utils.CenterText("Liste de mots changée !", 65), "green")
		utils.WriteColorLn(utils.CenterText("Nouvelle liste : "+values.WordFiles[values.CurrentOption].Name, 65), "cyan")
		time.Sleep(1 * time.Second)
		values.CurrentOption = 0
		values.CurrentPage = "main_menu"
		graphic.RefreshMainMenu()
	} else {
		// Option Retour
		values.CurrentOption = 0
		values.CurrentPage = "main_menu"
		graphic.RefreshMainMenu()
	}
}
