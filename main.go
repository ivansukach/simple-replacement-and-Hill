package main

import (
	"bufio"
	"log"
	"os"
)

var alphabet [33]rune = [33]rune{'а', 'б', 'в', 'г', 'д', 'е', 'ё', 'ж', 'з', 'и', 'й', 'к', 'л', 'м', 'н', 'о', 'п', 'р', 'с', 'т', 'у',
	'ф', 'х', 'ц', 'ч', 'ш', 'щ', 'ъ', 'ы', 'ь', 'э', 'ю', 'я'}
var replacement = make(map[rune]rune, 33)
var keyHillMatrix [2][2]int
var reverseKeyHillMatrix [2][2]int
var det int
var reverseDet int
var amount = 33

func init() {
	replacement['а'] = 'б'
	replacement['б'] = 'я'
	replacement['в'] = 'ж'
	replacement['г'] = 'ь'
	replacement['д'] = 'н'
	replacement['е'] = 'в'
	replacement['ё'] = 'щ'
	replacement['ж'] = 'г'
	replacement['з'] = 'з'
	replacement['и'] = 'п'
	replacement['й'] = 'к'
	replacement['к'] = 'ш'
	replacement['л'] = 'ъ'
	replacement['м'] = 'д'
	replacement['н'] = 'о'
	replacement['о'] = 'м'
	replacement['п'] = 'ё'
	replacement['р'] = 'й'
	replacement['с'] = 'ц'
	replacement['т'] = 'и'
	replacement['у'] = 'ч'
	replacement['ф'] = 'у'
	replacement['х'] = 'ы'
	replacement['ц'] = 'ю'
	replacement['ч'] = 'а'
	replacement['ш'] = 'э'
	replacement['щ'] = 'е'
	replacement['ъ'] = 'т'
	replacement['ы'] = 'л'
	replacement['ь'] = 'с'
	replacement['э'] = 'х'
	replacement['ю'] = 'р'
	replacement['я'] = 'ф'
	keyHillMatrix[0][0] = 11
	keyHillMatrix[0][1] = 7
	keyHillMatrix[1][0] = 24
	keyHillMatrix[1][1] = 26
	det = keyHillMatrix[0][0]*keyHillMatrix[1][1] - keyHillMatrix[0][1]*keyHillMatrix[1][0]
	log.Println("Det: ", det)
	for reverseDet = 0; ; reverseDet++ {
		if (det*reverseDet)%amount == 1 {
			break
		}
	}
	log.Println("reverseDet: ", reverseDet)
	reverseKeyHillMatrix[0][0] = keyHillMatrix[1][1] * reverseDet % amount
	reverseKeyHillMatrix[0][1] = (0 - keyHillMatrix[0][1]) * reverseDet % amount
	reverseKeyHillMatrix[1][0] = (0 - keyHillMatrix[1][0]) * reverseDet % amount
	reverseKeyHillMatrix[1][1] = keyHillMatrix[0][0] * reverseDet % amount
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			for reverseKeyHillMatrix[i][j] < 0 {
				reverseKeyHillMatrix[i][j] += amount
			}
		}
	}
	log.Println("Обратная по модулю матрица:")
	log.Println(reverseKeyHillMatrix[0][0], reverseKeyHillMatrix[0][1])
	log.Println(reverseKeyHillMatrix[1][0], reverseKeyHillMatrix[1][1])

}
func getIndex(value rune) int {
	for index, symbol := range alphabet {
		if symbol == value {
			return index
		}
	}
	return -1
}
func main() {
	encryptedText := ""
	reader := bufio.NewReader(os.Stdin)
	log.Println("Введите текст, который нужно зашифровать")
	t, _, _ := reader.ReadLine()
	text := string(t)
	for index, value := range text {
		log.Printf("%#U starts at byte position %d \n", value, index)
		encryptedText += string(replacement[value])
	}
	log.Println("encryptedText: ", encryptedText)

	decryptedText := ""
	log.Println("Введите текст, который нужно расшифровать")
	t2, _, _ := reader.ReadLine()
	text2 := string(t2)
	indexInAlphabet1 := 0
	indexInAlphabet2 := 0
	indexOfNewChar1 := 0
	indexOfNewChar2 := 0
	stepOdd := false
	for index, value := range text2 {
		log.Printf("%#U starts at byte position %d \n", value, index)
		if stepOdd {
			indexInAlphabet2 = getIndex(value)
			indexOfNewChar1 = (indexInAlphabet1*reverseKeyHillMatrix[0][0] + indexInAlphabet2*reverseKeyHillMatrix[0][1]) % amount
			indexOfNewChar2 = (indexInAlphabet1*reverseKeyHillMatrix[0][1] + indexInAlphabet2*reverseKeyHillMatrix[1][1]) % amount
			decryptedText += string(alphabet[indexOfNewChar1]) + string(alphabet[indexOfNewChar2])
			stepOdd = false
		} else {
			indexInAlphabet1 = getIndex(value)
			stepOdd = true
		}
	}
	log.Println("decryptedText: ", decryptedText)
}
