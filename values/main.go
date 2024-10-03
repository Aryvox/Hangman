package values

var CurrentPage string = "main_menu"
var CurrentOption int = 0
var CurrentOptionMax int = 1

type WordFile struct {
	Name        string
	Path        string
	Description string
}

var (
	CurrentWordFile = "mot.txt"
	WordFiles       = []WordFile{
		{
			Name:        "Français",
			Path:        "mot.txt",
			Description: "Liste de mots en français",
		},
		{
			Name:        "Anglais",
			Path:        "mot_anglais.txt",
			Description: "Liste de mots en anglais",
		},
	}
)
