package main

import (
	"fmt"
)

type Student struct {
	Name      string
	NoSubject int
	Subjects  map[string]float32
	Average   float32
	Sum       float32
}

func Average(student Student) float32 {
	for _, value := range student.Subjects {
		student.Sum += value
	}
	return student.Sum / float32(student.NoSubject)
}

func main() {

	var student1 Student

	fmt.Println("Enter your name:")
	fmt.Scan(&student1.Name)

	fmt.Println("Enter the number of subjects:")
	fmt.Scan(&student1.NoSubject)

	var subjectname string
	var subjectvalue float32

	for i := 0; i < student1.NoSubject; i++ {
		fmt.Println("Enter the title of the subject:")
		fmt.Scan(&subjectname)

		fmt.Printf("Enter the result of %s:\n", subjectname)
		fmt.Scan(&subjectvalue)

		for subjectvalue > 100 || subjectvalue < 0 {
			fmt.Println("\nPlease Enter a valid grade !")
			fmt.Printf("Enter the result of %s:\n", subjectname)
			fmt.Scan(&subjectvalue)
		}

		student1.Subjects[subjectname] = subjectvalue

	}

	fmt.Println(Average(student1))

}
