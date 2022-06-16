package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"time"
)

var mistakes int = 0        //how many mistakes the player has made
var typedLetters = []byte{} //letters typed by the player
var word = ""               //word to guess
var hint = ""               //hint to the word to be guessed
var clear map[string]func() //map of commands to clean the screen by operating systems
var words = [][]string{
	{"bookworm", "a person or insect"},
	{"awkward", "hard, uneassy"},
	{"buzzwords", "fashionable word or phrase"},
	{"dwarves", "fantasy humanlike race"},
	{"fluffiness", "soft and light"},
	{"galvanize", "shock or protective layer"},
	{"united states of america", "quite big country"},
}

func main() {
	reset()
	print()
	play()
}

func play() {
	action()
	print()
	if hasLose() {
		fmt.Println("you lose, try again")
		solve()
		fmt.Print("\r\nsolution: ")
		printMatchWord() //print whole word
		// time.Sleep(5 * time.Second)
		restart()
	}
	if hasWon() {
		fmt.Println("you won, congratulations")
		restart()
	}
	play() //play till win or lose
}

func reset() {
	randWord() //draft word
	mistakes = 0
	typedLetters = nil
	print()
}

func solve() {
	typedLetters = nil
	for _, l := range word {
		if byte(l) == ' ' {
			continue
		}
		typedLetters = append(typedLetters, byte(l))
	}
}

func restart() {
	readCharacter("press any key")
	reset()
	play()
}

func randWord() {
	rand.Seed(time.Now().UnixNano())
	min := 0
	max := len(words) - 1
	key := rand.Intn(max-min+1) + min
	word = words[key][0]
	hint = words[key][1]
}

func print() {
	clearScreen()
	printGallow()
	fmt.Print("\r\n")
	printMatchWord()
	fmt.Print("\r\n")
	printAvailableLetters()
	fmt.Print("\r\n")
	printHint()
	fmt.Print("\r\n")
}

func initClear() {
	if len(clear) > 0 {
		return
	}
	clear = make(map[string]func()) //Initialize it
	clear["linux"] = func() {
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func clearScreen() {
	initClear()
	value, ok := clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
	if ok {                          //if we defined a clear func for that platform:
		value() //we execute it
	} else { //unsupported platform
		panic("Your platform is unsupported! I can't clear terminal screen :(")
	}
}

func action() {
	char := getValidChar()
	typedLetters = append(typedLetters, char)
	fmt.Println(hasMatch(char))
	if !hasMatch(char) {
		mistakes++
	}
}

func hasLose() bool {
	return mistakes >= 9
}

func hasWon() bool {
	for _, l := range word {
		if byte(l) == ' ' {
			continue
		}
		if !hasTyped(byte(l)) {
			return false
		}
	}
	return true //player has typed all letters
}

func printAvailableLetters() {
	fmt.Print("\r\n")
	for ch := 97; ch <= 122; ch++ {
		if (ch-97)%9 == 0 && ch != 97 {
			fmt.Print("\r\n")
		}
		pChar := ch
		if hasTyped(byte(ch)) {
			pChar = '_'
		}
		fmt.Printf("%c ", pChar)
	}
}

func printHint() {
	fmt.Print("\r\nhint:" + hint)
}

func hasMatch(char byte) bool {
	for _, l := range word {
		if byte(l) == char {
			return true
		}
	}
	return false
}

func printMatchWord() {
	for _, l := range word {
		p := '_'
		if hasTyped(byte(l)) || byte(l) == ' ' {
			p = l
		}
		fmt.Printf("%c", p)
	}
}

func hasTyped(char byte) bool {
	for _, ch := range typedLetters {
		if char == ch {
			return true
		}
	}
	return false
}

func getValidChar() byte {
	char := readCharacter("enter letter")
	if !isValidLetter(char) {
		fmt.Println("invalid letter, a-z allowed")
		return getValidChar()
	}
	if hasTyped(char) {
		fmt.Printf("\r\nyou have already entered %c", char)
		return getValidChar()
	}
	return char
}

func readCharacter(label string) byte {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("\r\n" + label)
	char, _ := reader.ReadByte()
	return char
}

func isValidLetter(char byte) bool {
	return char >= 97 && char <= 122
}

func printGallow() {
	if mistakes == 0 {
		for i := 0; i < 7; i++ {
			fmt.Print("\r\n")
		}
		return //no bother
	}
	printGallow_7()
	printGallow_6()
	printGallow_5()
	printGallow_4()
	printGallow_3()
	printGallow_2()
	printGallow_1()

	//7  ____
	//6  |/ |
	//5  |  O
	//4  | <|>
	//3  | / \
	//2  |
	//1 ===
}

func printGallow_1() {
	if mistakes > 0 {
		fmt.Println("===")
	} else {
		fmt.Print("\r\n")
	}
}

func printGallow_2() {
	if mistakes > 1 {
		fmt.Println(" |")
	} else {
		fmt.Print("\r\n")
	}
}

func printGallow_3() {
	if mistakes == 1 {
		fmt.Print("\r\n")
		return
	}
	gallow := " |"
	if mistakes == 8 {
		gallow = " | / "
	}
	if mistakes > 8 {
		gallow = " | / \\ "
	}
	fmt.Println(gallow)
}

func printGallow_4() {
	if mistakes == 1 {
		fmt.Print("\r\n")
		return
	}
	gallow := " |"
	if mistakes == 5 {
		gallow = " |  | "
	}
	if mistakes == 6 {
		gallow = " |  |>"
	}
	if mistakes > 6 {
		gallow = " | <|>"
	}
	fmt.Println(gallow)
}

func printGallow_5() {
	if mistakes == 1 {
		fmt.Print("\r\n")
		return
	}
	gallow := " |"
	if mistakes >= 4 {
		gallow = " |  O "
	}
	fmt.Println(gallow)
}

func printGallow_6() {
	if mistakes == 1 {
		fmt.Print("\r\n")
		return
	}
	gallow := " |"
	if mistakes >= 3 {
		gallow = " |/ |"
	}
	fmt.Println(gallow)
}

func printGallow_7() {
	if mistakes == 1 {
		fmt.Print("\r\n")
		return
	}
	if mistakes >= 3 {
		fmt.Println(" ____")
	}
}
