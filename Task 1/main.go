package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Student struct {
	Name      string
	NoSubject int
	Subjects  map[string]float64
	Average   float64
	Sum       float64
}

func Average(student Student) float64 {
	for _, value := range student.Subjects {
		student.Sum += value
	}
	return student.Sum / float64(student.NoSubject)
}

func main() {

	reader := bufio.NewReader(os.Stdin)

	var student1 Student
	student1.Subjects = make(map[string]float64)

	fmt.Println("Enter your name:")
	student1.Name, _ = reader.ReadString('\n')
	student1.Name = strings.TrimSpace(student1.Name)

	fmt.Println("Enter the number of subjects:")
	noSubjectStr, _ := reader.ReadString('\n')
	student1.NoSubject, _ = strconv.Atoi(strings.TrimSpace(noSubjectStr))

	var (
		subjectname  string
		subjectvalue float64
	)

	for i := 0; i < student1.NoSubject; i++ {
		fmt.Println("Enter the title of the subject:")
		subjectname, _ = reader.ReadString('\n')
		subjectname = strings.TrimSpace(subjectname)

		fmt.Printf("Enter the result of %s:\n", subjectname)
		subjectvalueStr, _ := reader.ReadString('\n')
		subjectvalue, _ = strconv.ParseFloat(strings.TrimSpace(subjectvalueStr), 64)

		for subjectvalue > 100 || subjectvalue < 0 {
			fmt.Println("\nPlease Enter a valid grade !")

			fmt.Printf("Enter the result of %s:\n", subjectname)
			subjectvalueStr, _ := reader.ReadString('\n')
			subjectvalue, _ = strconv.ParseFloat(strings.TrimSpace(subjectvalueStr), 64)
		}

		student1.Subjects[subjectname] = subjectvalue

	}

	student1.Average = Average(student1)

	fmt.Printf("%-20s %-10s\n", student1.Name, "\nStudent Report")
	fmt.Println(strings.Repeat("-", 40))
	fmt.Printf("%-20s %-10s\n", "Subject", "Grade")
	fmt.Println(strings.Repeat("-", 40))
	for subject, grade := range student1.Subjects {
		fmt.Printf("%-20s %-10.2f\n", subject, grade)
	}
	fmt.Println(strings.Repeat("-", 40))
	fmt.Printf("%-20s %-10.2f\n\n", "Average", student1.Average)

}
