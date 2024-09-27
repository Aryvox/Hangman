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
		game.StartGame() // Ajout de cette ligne pour démarrer le jeu
	case 1:
		utils.ClearTerminal()
		graphic.RefreshCreditsMenu()
	case 2:
		utils.ClearTerminal()
		PressF11()
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
