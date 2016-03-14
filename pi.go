package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
	"unicode"
)

const pidigits = "14159265358979323846264338327950288419716939937510"
const maxDigits = len(pidigits)

var numDigits int

func main() {
	flag.IntVar(&numDigits, "n", maxDigits, fmt.Sprintf("The number of digits required to win the game.  0 < n <= %d.", maxDigits))
	flag.Parse()
	if numDigits == 0 || numDigits > maxDigits {
		fmt.Printf("The value of n must be be greater than 0 and less than %d.\n", maxDigits+1)
		os.Exit(1)
	}

	printInstructions()
	var inputChan = make(chan rune)
	var successChan = make(chan bool)
	go readUserInput(inputChan)
	go processUserInput(inputChan, successChan)

	<-successChan
	printSuccess()
}

func processUserInput(inputChan chan rune, successChan chan bool) {
	var currentDigit = 0
	for {
		var r = <-inputChan
		if !unicode.IsDigit(r) {
			continue
		}

		var input = fmt.Sprintf("%c", r)
		var expected = pidigits[currentDigit : currentDigit+1]
		if input != expected {
			fmt.Printf("Wrong Answer, the next digit was %s. Try again!\n", expected)
			os.Exit(0)
		}

		currentDigit++
		if currentDigit == numDigits {
			successChan <- true
		}
	}
}

func readUserInput(inputChan chan rune) {
	var reader = bufio.NewReader(os.Stdin)
	for {
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("An error was encountered.\n", err)
			os.Exit(1)
		}

		trimmed := strings.TrimSpace(input)
		for _, c := range trimmed {
			inputChan <- c
		}
	}
}

func printInstructions() {
	fmt.Printf("------------------\n"+
		"Let's try to memorize the number Pi to %d decimal places.\n"+
		"We'll start you off with the number 3.\n"+
		"Enter what you think the next digit is, and press enter.\n"+
		"Keep going until you win or get one wrong.  Have fun!\n"+
		"------------------\n\n3\n.\n", numDigits)
}

func printSuccess() {
	fmt.Printf("\nWay to go!  You memorized Pi to %d decimal places.\n\n", numDigits)
	fmt.Println("      3.141592653589793238462643383279\n" +
		"    5028841971693993751058209749445923\n" +
		"   07816406286208998628034825342117067\n" +
		"   9821    48086         5132\n" +
		"  823      06647        09384\n" +
		" 46        09550        58223\n" +
		" 17        25359        4081\n" +
		"           2848         1117\n" +
		"           4502         8410\n" +
		"           2701         9385\n" +
		"          21105        55964\n" +
		"          46229        48954\n" +
		"          9303         81964\n" +
		"          4288         10975\n" +
		"         66593         34461\n" +
		"        284756         48233\n" +
		"        78678          31652        71\n" +
		"       2019091         456485       66\n" +
		"      9234603           48610454326648\n" +
		"     2133936            0726024914127\n" +
		"     3724587             00660631558\n" +
		"     817488               152092096\n")
}
