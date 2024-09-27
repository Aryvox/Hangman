package game

import (
	"hangman/utils"
	"io/ioutil"
)

func main() {
	/*Fonction pour lire le mot.txt*/
	files, err := ioutil.ReadFile("mot.txt")
	if err != nil {
		utils.Writeln(err.Error())
		return
	}
	utils.Writeln(string(files))
}
