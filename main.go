package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

func lz77Encode(dictionarySize, bufferSize int, text string) {
	var dictionary, buffer []rune

	for step := 1; len(text) > 0 || len(buffer) > 0; step++ {
		if len(dictionary) > dictionarySize {
			dictionary = dictionary[len(dictionary)-dictionarySize:]
		}

		for len(text) > 0 && len(buffer) < bufferSize {
			r, size := utf8.DecodeRuneInString(text)
			buffer = append(buffer, r)
			text = text[size:]
		}

		maxMatchLength, maxMatchIndex := 0, -1
		for i := 0; i <= len(dictionary); i++ {
			matchLength := 0
			for j := 0; j < len(buffer) && i+j < len(dictionary); j++ {
				if dictionary[i+j] != buffer[j] {
					break
				}
				matchLength++
			}
			if matchLength > maxMatchLength {
				maxMatchLength = matchLength
				maxMatchIndex = i
			}
		}

		fmt.Printf("Шаг %d: ", step)
		fmt.Printf("Словарь: %v ", string(dictionary))
		fmt.Printf("Буфер: %v ", string(buffer))

		if maxMatchLength > 0 && maxMatchLength <= len(buffer) {
			correctIndex := (dictionarySize - len(dictionary) + maxMatchIndex) % dictionarySize
			fmt.Printf("<%d,%d,%s>\n", correctIndex, maxMatchLength, string(buffer[maxMatchLength-1]))
			dictionary = append(dictionary, buffer[:maxMatchLength]...)
			buffer = buffer[maxMatchLength:]
		} else if len(buffer) > 0 {
			fmt.Printf("<0,0,%s>\n", string(buffer[0]))
			dictionary = append(dictionary, buffer[0])
			buffer = buffer[1:]
		}
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите размер словаря: ")
	dictionarySizeInput, _ := reader.ReadString('\n')
	var dictionarySize int
	fmt.Sscanf(dictionarySizeInput, "%d", &dictionarySize)

	fmt.Print("Введите размер буфера: ")
	bufferSizeInput, _ := reader.ReadString('\n')
	var bufferSize int
	fmt.Sscanf(bufferSizeInput, "%d", &bufferSize)

	fmt.Print("Введите сообщение для кодирования: ")
	text, _ := reader.ReadString('\n')
	text = strings.ToUpper(strings.TrimSpace(text))
	text = strings.Replace(text, " ", "_", -1)

	lz77Encode(dictionarySize, bufferSize, text)

	fmt.Println("Надеюсь на добросовестное использование данного продукта ;)")
	fmt.Println("Подпишитесь пожалуйста на мой github или оставьте отзыв - https://github.com/MrInsome/lZ77Encoder")
	fmt.Println("Нажмите Enter для завершения программы")

	_, _ = reader.ReadString('\n')
}
