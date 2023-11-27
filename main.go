package main

import (
	"fmt"
	"strings"
)

var morseCodeMap = map[rune]string{
	'A': ".-", 'B': "-...", 'C': "-.-.", 'D': "-..", 'E': ".",
	'F': "..-.", 'G': "--.", 'H': "....", 'I': "..", 'J': ".---",
	'K': "-.-", 'L': ".-..", 'M': "--", 'N': "-.", 'O': "---",
	'P': ".--.", 'Q': "--.-", 'R': ".-.", 'S': "...", 'T': "-",
	'U': "..-", 'V': "...-", 'W': ".--", 'X': "-..-", 'Y': "-.--",
	'Z': "--..",
	'1': ".----", '2': "..---", '3': "...--", '4': "....-",
	'5': ".....", '6': "-....", '7': "--...", '8': "---..",
	'9': "----.", '0': "-----",
}

func textToMorse(text string) string {
	text = strings.ToUpper(text)
	var morseCode string

	for _, char := range text {
		if char == ' ' {
			morseCode += " "
		} else if code, found := morseCodeMap[char]; found {
			morseCode += code + " "
		}
	}

	return morseCode
}

func main() {
	sentence := "The quick brown fox jumps over 13 lazy dogs"
	morseCode := textToMorse(sentence)
	fmt.Println(morseCode)
}
