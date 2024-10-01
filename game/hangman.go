package game

import (
	"fmt"
	"hangman/graphic"
	"hangman/sounds"
	"hangman/utils"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/mattn/go-tty"
	"golang.org/x/term"
)

func centerString(s string) string {
	width, _, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		return s
	}
	if len(s) >= width {
		return s
	}
	padding := (width - len(s)) / 2
	return strings.Repeat(" ", padding) + s
}

func StartGame() {
	words, err := utils.ReadWordsFromFile("mot.txt")
	if err != nil {
		utils.WriteColorLn(centerString("Erreur lors de la lecture du fichier de mots"), "red")
		time.Sleep(3 * time.Second)
		utils.ClearTerminal()
		graphic.RefreshMainMenu()
		return
	}

	word := chooseRandomWord(words)
	guessedLetters := make(map[rune]bool)
	wrongGuesses := 0
	maxWrongGuesses := 9

	revealRandomLetter(word, guessedLetters)

	for wrongGuesses < maxWrongGuesses {
		utils.ClearTerminal()
		displayHangman(wrongGuesses)
		displayWord(word, guessedLetters)
		displayGuessedLetters(word, guessedLetters)
		utils.WriteColorLn(centerString(fmt.Sprintf("Essais restants : %d", maxWrongGuesses-wrongGuesses)), "cyan")

		utils.WriteColorLn(centerString("Entrez une lettre : "), "cyan")
		guess, err := getInput()
		if err != nil {
			utils.WriteColorLn(centerString("Erreur lors de la saisie"), "red")
			continue
		}

		if len(guess) != 1 {
			utils.WriteColorLn(centerString("Veuillez entrer une seule lettre"), "yellow")
			time.Sleep(1 * time.Second)
			continue
		}

		letter := rune(strings.ToUpper(guess)[0])

		if guessedLetters[letter] {
			utils.WriteColorLn(centerString("Vous avez déjà deviné cette lettre"), "yellow")
			time.Sleep(1 * time.Second)
			continue
		}

		guessedLetters[letter] = true

		if !strings.ContainsRune(word, letter) {
			wrongGuesses++
			utils.WriteColorLn(centerString(fmt.Sprintf("Mauvaise lettre : %c", letter)), "red")
		} else {
			utils.WriteColorLn(centerString(fmt.Sprintf("Bonne lettre : %c", letter)), "green")
		}

		utils.WriteColorLn(centerString("Veuillez patienter avant de rentrer une lettre..."), "yellow")
		time.Sleep(1 * time.Second)

		if isWordGuessed(word, guessedLetters) {
			utils.ClearTerminal()
			displayHangman(wrongGuesses)
			displayWord(word, guessedLetters)
			utils.WriteColorLn(centerString("Félicitations ! Vous avez deviné le mot !"), "green")
			time.Sleep(3 * time.Second)
			utils.ClearTerminal()
			graphic.RefreshMainMenu()
			return
		}
	}

	utils.ClearTerminal()
	utils.PlaySound(sounds.Wasted, 1)
	displayHangman(wrongGuesses)
	displayWord(word, guessedLetters)
	utils.WriteColorLn(centerString("Désolé, vous avez perdu. Le mot était : "+word), "red")
	time.Sleep(3 * time.Second)
	utils.ClearTerminal()
	graphic.RefreshMainMenu()
}

func chooseRandomWord(words []string) string {
	rand.Seed(time.Now().UnixNano())
	return strings.ToUpper(words[rand.Intn(len(words))])
}

func displayHangman(wrongGuesses int) {
	hangmanFile := fmt.Sprintf("hangman_step_%d.txt", wrongGuesses+1)
	content, err := os.ReadFile(filepath.Join("ascii", hangmanFile))
	if err != nil {
		utils.WriteColorLn(centerString("Erreur lors de la lecture du fichier du pendu"), "red")
		return
	}
	centeredHangman := centerString(string(content))
	utils.Writeln(centeredHangman)
}

func displayWord(word string, guessedLetters map[rune]bool) {
	var display string
	for _, letter := range word {
		if guessedLetters[letter] {
			display += string(letter) + " "
		} else {
			display += "_ "
		}
	}
	centeredWord := centerString("Mot : " + display)
	utils.Writeln(centeredWord)
}

func displayGuessedLetters(word string, guessedLetters map[rune]bool) {
	var correctLetters, wrongLetters string
	for letter := range guessedLetters {
		if strings.ContainsRune(word, letter) {
			correctLetters += string(letter) + " "
		} else {
			wrongLetters += string(letter) + " "
		}
	}

	utils.WriteColorLn(centerString("Lettres correctes : "+correctLetters), "green")
	utils.WriteColorLn(centerString("Lettres incorrectes : "+wrongLetters), "red")
}

func isWordGuessed(word string, guessedLetters map[rune]bool) bool {
	for _, letter := range word {
		if !guessedLetters[letter] {
			return false
		}
	}
	return true
}

func getInput() (string, error) {
	tty, err := tty.Open()
	if err != nil {
		return "", err
	}
	defer tty.Close()

	r, err := tty.ReadRune()
	if err != nil {
		return "", err
	}

	utils.WriteColorLn(string(r), "yellow")

	return string(r), nil
}

func revealRandomLetter(word string, guessedLetters map[rune]bool) {
	rand.Seed(time.Now().UnixNano())
	letterIndex := rand.Intn(len(word))
	revealedLetter := rune(word[letterIndex])
	guessedLetters[revealedLetter] = true
	utils.WriteColorLn(centerString(fmt.Sprintf("Une lettre a été révélée : %c", revealedLetter)), "green")
	time.Sleep(2 * time.Second)
}
