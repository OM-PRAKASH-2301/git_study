package main

import "fmt"

func main() {
	students := make(map[int]string)

	students[101] = "Om Prakash"
	students[102] = "Amit Kumar"
	students[103] = "Rahul Sharma"

	fmt.Println("Student with ID 101:", students[101])

	if name, ok := students[104]; ok {
		fmt.Println("Student with ID 104:", name)
	} else {
		fmt.Println("Student with ID 104 does not exist.")
	}

	students[103] = "Rohit Verma"
	fmt.Println("Updated Student with ID 103:", students[103])

	delete(students, 102)
	fmt.Println("All students after deletion:", students)
}
