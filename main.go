package main

import (
	"fmt"
	"os"
	"strings"
)

func FileRead() (string, error) {
	style := os.Args[2]
	file, err := os.ReadFile(style + ".txt")
	if err != nil {
		return "We couldn't read the data in the file", err
	}
	return string(file), err
}

func FindAndPrint(lines []string, word string) {
	for h := 1; h < 9; h++ { // Satırlar için olan Döngü (Loop for rows)
		for i := 0; i < len(word); i++ { // Kelimenin harfleri için döngü (Loop for letters of the word)
			for lineIndex, line := range lines { // Dosya Satırları için olan Döngü (Loop for File Lines)
				if lineIndex == (int(word[i])-32)*9+h { // Karakterin ASCII değerine göre ilgili satırı bulma işlemi (Find the corresponding line according to the ASCII value of the character)
					fmt.Print(line)
				}
			}
		}
		fmt.Println() // Bir satırın harf döngüsü tamamlandığında alt satıra geçme (Move to the next line when the letter loop of a line is complete)
	}
}

func main() {

	if len(os.Args) == 4 {
		fmt.Println("You have entered too many arguments")
		return

	}
	if len(os.Args) < 3 {
		fmt.Println("You entered an incomplete number of arguments")
		return
	}

	argStr := os.Args[1]

	// Girilen Parametremizde '\n' ifadesini görünce Parçalmaısnı sağladık Bu işlem, kullanıcının girdiği metni yeni satırlara bölmek için yapılmıştır.
	//  "Hello\nGO"  burada  --> "Hello" - "GO" şkelinde iki kelime gibi yazdıracaktır

	statement := strings.Split(argStr, "\\n")

	// When we see '\n' in our entered parameter, we have made it fragmented. This is done to split the text entered by the user into new lines.
	// "Hello\nGO" here --> will print "Hello" - "GO" as two words)

	lines, err := FileRead()
	if err != nil {
		fmt.Println("File read error")
		panic(err)
	}

	for i, word := range statement {
		if word == "" {
			if i != 0 {
				fmt.Println()
			}
			continue
		}

		FindAndPrint(strings.Split(lines, "\n"), word)
	}
}
