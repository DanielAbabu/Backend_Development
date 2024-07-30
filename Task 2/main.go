package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func Letteronly(word string) string {
	newword := ""
	for i := 0; i < len(word); i++ {
		if unicode.IsLetter(rune(word[i])) {
			newword += string(word[i])
		}
	}
	return newword

}

func Count(s string) map[string]int {
	dict := make(map[string]int)

	arr := strings.Fields(s)
	for _, word := range arr {
		word = strings.ToLower(word)
		word = Letteronly(word)

		if _, ok := dict[word]; !ok {
			dict[word] = 1
		} else {
			dict[word] += 1
		}
	}

	return dict
}

func Plaindrome(str string) bool {
	var newstr string = Letteronly(strings.ToLower(str))

	var reversed string = ""

	for i := len(newstr) - 1; i >= 0; i-- {
		reversed += string(newstr[i])
	}
	fmt.Println(newstr)
	return reversed == newstr
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter the String you want to count the word Frequency:")
	str, _ := reader.ReadString('\n')
	fmt.Println(Count(str))

	fmt.Println("Enter the String you want to check if it is Palindrom:")
	strr, _ := reader.ReadString('\n')
	fmt.Println(Plaindrome(strr))

}
