package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"github.com/mattn/go-tty"
)

func runCmd(name string, arg ...string) {
	cmd := exec.Command(name, arg...)
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func ClearTerminal() {
	switch runtime.GOOS {
	case "darwin":
		runCmd("clear")
	case "linux":
		runCmd("clear")
	case "windows":
		runCmd("cmd", "/c", "cls")
	default:
		runCmd("clear")
	}
}

func Write(text string, args ...any) {
	if args != nil && len(args) > 0 {
		fmt.Printf(text, args)
	} else {
		fmt.Print(text)
	}
}

func Writeln(text string, args ...any) {
	if args != nil && len(args) > 0 {
		fmt.Printf(text, args)
		fmt.Println()
	} else {
		fmt.Println(text)
	}
}

func Writeanim(text string, timee time.Duration, args ...any) {
	for _, character := range text {
		fmt.Print(string(rune(character)))
		time.Sleep(timee * time.Millisecond)
	}
}

func ListenInput() string {
	println("listening for input")
	tty, err := tty.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer tty.Close()

	for {
		println("aaa")
		r, err := tty.ReadRune()
		if err != nil {
			log.Fatal(err)
		}
		return string(r)
	}
}

func WaitForInput() string {
	var input string
	fmt.Scanln(&input)
	return input
}

func WaitForNumberInput() (bool, int) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	err := scanner.Err()
	if err != nil {
		log.Fatal(err)
	}
	if !IsNumeric(scanner.Text()) {
		return false, -1
	}
	number, err := strconv.Atoi(scanner.Text())
	if err != nil {
		log.Fatal(err)
		return false, -1
	}
	return true, number
}

func IsNumeric(s string) bool {
	for _, char := range s {
		if !('0' <= char && char <= '9') {
			return false
		}
	}
	return true
}

func CenterText(text string, width int) string {
	return strings.Repeat(" ", width) + text
}

func WriteColorLn(text string, color string) {
	switch color {
	case "red":
		fmt.Printf("\033[31m%s\033[0m\n", text)
	case "green":
		fmt.Printf("\033[32m%s\033[0m\n", text)
	case "yellow":
		fmt.Printf("\033[33m%s\033[0m\n", text)
	case "blue":
		fmt.Printf("\033[34m%s\033[0m\n", text)
	case "magenta":
		fmt.Printf("\033[35m%s\033[0m\n", text)
	case "cyan":
		fmt.Printf("\033[36m%s\033[0m\n", text)
	case "white":
		fmt.Printf("\033[37m%s\033[0m\n", text)
	default:
		fmt.Printf("\033[37m%s\033[0m\n", text)
	}
	fmt.Println()
}

func WriteanimColor(text string, timee time.Duration, color string, args ...any) {
	for _, character := range text {
		switch color {
		case "red":
			fmt.Printf("\033[31m%s\033[0m", string(character))
		case "green":
			fmt.Printf("\033[32m%s\033[0m", string(character))
		case "yellow":
			fmt.Printf("\033[33m%s\033[0m", string(character))
		case "blue":
			fmt.Printf("\033[34m%s\033[0m", string(character))
		case "magenta":
			fmt.Printf("\033[35m%s\033[0m", string(character))
		case "cyan":
			fmt.Printf("\033[36m%s\033[0m", string(character))
		case "white":
			fmt.Printf("\033[37m%s\033[0m", string(character))
		default:
			fmt.Printf("\033[37m%s\033[0m", string(character))
		}
		time.Sleep(timee * time.Millisecond)
	}
}

func WriteanimColorln(text string, timee time.Duration, color string, args ...any) {
	for _, character := range text {
		switch color {
		case "red":
			fmt.Printf("\033[31m%s\033[0m", string(character))
		case "green":
			fmt.Printf("\033[32m%s\033[0m", string(character))
		case "yellow":
			fmt.Printf("\033[33m%s\033[0m", string(character))
		case "blue":
			fmt.Printf("\033[34m%s\033[0m", string(character))
		case "magenta":
			fmt.Printf("\033[35m%s\033[0m", string(character))
		case "cyan":
			fmt.Printf("\033[36m%s\033[0m", string(character))
		case "white":
			fmt.Printf("\033[37m%s\033[0m", string(character))
		default:
			fmt.Printf("\033[37m%s\033[0m", string(character))
		}
		time.Sleep(timee * time.Millisecond)
	}
	fmt.Println()
}

func WriteColor(text string, color string) {
	switch color {
	case "red":
		fmt.Printf("\033[31m%s\033[0m\n", text)
	case "green":
		fmt.Printf("\033[32m%s\033[0m\n", text)
	case "yellow":
		fmt.Printf("\033[33m%s\033[0m\n", text)
	case "blue":
		fmt.Printf("\033[34m%s\033[0m\n", text)
	case "magenta":
		fmt.Printf("\033[35m%s\033[0m\n", text)
	case "cyan":
		fmt.Printf("\033[36m%s\033[0m\n", text)
	case "white":
		fmt.Printf("\033[37m%s\033[0m\n", text)
	default:
		fmt.Printf("\033[37m%s\033[0m\n", text)
	}
}

func PlaySound(path string, rate ...int) {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	if len(rate) < 1 {
		rate = []int{1}
	}

	streamer, format, err := mp3.Decode(f)
	if err != nil {
		log.Fatal(err)
	}
	defer streamer.Close()

	//speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	speaker.Init(format.SampleRate*beep.SampleRate(rate[0]), format.SampleRate.N(time.Second/10))

	done := make(chan bool)
	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		done <- true
	})))

	<-done
}

func ReadWordsFromFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var words []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		word := strings.TrimSpace(scanner.Text())
		if word != "" {
			words = append(words, word)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return words, nil
}
