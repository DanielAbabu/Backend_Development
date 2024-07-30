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

func Palindrome(str string) bool {
	var newstr string = Letteronly(strings.ToLower(str))

	var reversed string = ""

	for i := len(newstr) - 1; i >= 0; i-- {
		reversed += string(newstr[i])
	}
	return reversed == newstr
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\nChoose an option:")
		fmt.Println("1. Count word frequency")
		fmt.Println("2. Check if a string is a palindrome")
		fmt.Println("3. Exit\n")

		option, _ := reader.ReadString('\n')
		option = strings.TrimSpace(option)

		switch option {

		case "1":
			fmt.Println("Enter the string you want to count the word frequency:")
			str, _ := reader.ReadString('\n')
			str = strings.TrimSpace(str)

			wordcount := Count(str)

			fmt.Println("\nWord frequency count:")
			fmt.Println(strings.Repeat("-", 40))
			fmt.Printf("%-20s %-10s\n", "Word", "Count")
			fmt.Println(strings.Repeat("-", 40))
			for word, count := range wordcount {
				fmt.Printf("%-20s %-10.2d\n", word, count)
			}

		case "2":
			fmt.Println("Enter the string you want to check if it is a palindrome:")
			str, _ := reader.ReadString('\n')
			str = strings.TrimSpace(str)
			if Palindrome(str) {
				fmt.Printf("\nThe '%s' is a palindrome.\n", str)
			} else {
				fmt.Printf("\nThe '%s' is not a palindrome.\n", str)
			}

		case "3":
			fmt.Println("Exiting the program.")
			return

		default:
			fmt.Println("Invalid option. Please choose again.")
		}
	}

}
