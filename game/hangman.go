package game

import (
	"fmt"
	"hangman/ascii"
	"hangman/graphic"
	"hangman/utils"
	"math/rand"
	"strings"
	"time"

	"os"

	"github.com/mattn/go-tty"
	"golang.org/x/term"
)

// Helper function to center a string based on terminal width
func centerString(s string) string {
	width, _, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		// Si on ne peut pas obtenir la taille du terminal, retourner la chaîne telle quelle
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
		return
	}

	word := chooseRandomWord(words)
	guessedLetters := make(map[rune]bool)
	wrongGuesses := 0
	maxWrongGuesses := 9

	// Révéler une lettre aléatoire au début du jeu
	revealRandomLetter(word, guessedLetters)

	for wrongGuesses < maxWrongGuesses {
		utils.ClearTerminal()
		displayHangman(wrongGuesses)
		displayWord(word, guessedLetters)
		displayGuessedLetters(word, guessedLetters)
		utils.WriteColorLn(centerString(fmt.Sprintf("Essais restants : %d", maxWrongGuesses-wrongGuesses)), "cyan")

		utils.WriteColor(centerString("Entrez une lettre : "), "cyan")
		guess, err := getInput()
		if err != nil {
			utils.WriteColorLn(centerString("Erreur lors de la saisie"), "red")
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
			time.Sleep(1 * time.Second)
		} else {
			utils.WriteColorLn(centerString(fmt.Sprintf("Bonne lettre : %c", letter)), "green")
			time.Sleep(1 * time.Second)
		}

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
	displayHangman(wrongGuesses)
	utils.WriteColorLn(centerString("Désolé, vous avez perdu. Le mot était : "+word), "red")
	time.Sleep(3 * time.Second)
	utils.ClearTerminal()
	graphic.RefreshMainMenu()
}

// Choisir un mot aléatoire parmi la liste
func chooseRandomWord(words []string) string {
	rand.Seed(time.Now().UnixNano())
	return strings.ToUpper(words[rand.Intn(len(words))])
}

// Afficher l'état du pendu
func displayHangman(wrongGuesses int) {
	hangmanASCII := []string{
		ascii.PenduStep1,
		ascii.PenduStep2,
		ascii.PenduStep3,
		ascii.PenduStep4,
		ascii.PenduStep5,
		ascii.PenduStep6,
		ascii.PenduStep7,
		ascii.PenduStep8,
		ascii.PenduStep9,
		ascii.PenduStep10,
	}

	centeredHangman := centerString(hangmanASCII[wrongGuesses])
	utils.Writeln(centeredHangman)
}

// Afficher le mot avec les lettres devinées
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

// Afficher les lettres devinées, séparées en correctes et incorrectes
func displayGuessedLetters(word string, guessedLetters map[rune]bool) {
	var correctLetters, wrongLetters string
	for letter := range guessedLetters {
		if strings.ContainsRune(word, letter) {
			correctLetters += string(letter) + " "
		} else {
			wrongLetters += string(letter) + " "
		}
	}
	guessed := "Lettres correctes : " + correctLetters + "\nLettres incorrectes : " + wrongLetters
	centeredGuessed := centerString(guessed)
	utils.Writeln(centeredGuessed)
}

// Vérifier si le mot a été entièrement deviné
func isWordGuessed(word string, guessedLetters map[rune]bool) bool {
	for _, letter := range word {
		if !guessedLetters[letter] {
			return false
		}
	}
	return true
}

// Obtenir l'entrée de l'utilisateur et afficher la lettre immédiatement
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

	// Afficher la lettre saisie immédiatement
	utils.WriteColor(string(r), "yellow")

	// Supprimer l'attente d'appui sur Entrée
	// utils.WriteColor(" (Appuyez sur Entrée pour confirmer) ", "yellow")
	// _, err = tty.ReadRune() // Attendre l'appui sur Entrée
	// if err != nil {
	// 	return "", err
	// }

	return string(r), nil
}

// Révéler une lettre aléatoire au début du jeu
func revealRandomLetter(word string, guessedLetters map[rune]bool) {
	rand.Seed(time.Now().UnixNano())
	letterIndex := rand.Intn(len(word))
	guessedLetters[rune(word[letterIndex])] = true
	utils.WriteColorLn(centerString(fmt.Sprintf("Une lettre a été révélée : %c", word[letterIndex])), "green")
	time.Sleep(2 * time.Second)
}
