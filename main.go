package main

import (
	"fmt"
	"errors"
	"go_basics/classes/first"
	"go_basics/classes/second"
	"go_basics/helpers"
	"go_basics/structs_example"
	"reflect"
)

func printSome() {
	fmt.Println("Someone called!!")
}

func printName(name string) {
	fmt.Println("Hello", name)
}

func formatName(name string) string {
	return fmt.Sprintf("Hello %s Welcome to Golang", name)
}


func multi(input int) (string, error) {
	if input == 2 {
		return "", errors.New("Inout cannot be 2")
	}
	return "he!!", nil
}


type Student struct {
	name string
	rollNo int
}

func getStudent(name string) (Student, error) {
	if name == "Amith" {
		return Student{
			name:   name,
			rollNo: 0,
		}, nil
	}
	return Student{}, errors.New("Name cannot be somehtij3pr")
}

func main() {


	a := 1
	b := &a
	*b = 21

	fmt.Println(a)


	fmt.Println("sasas %T", reflect.TypeOf(b))


	student := structs_example.StudentExample{
		Name: "Amithhththtt",
		RollNo: 1,
	}

	fmt.Println("before ", student.Name)

	student.UpdateName("Second!!!")

	fmt.Println("after ", student.Name)


	fmt.Println(student.GetStudentName())

	myName := "Amith"




	firstClassName := first.FormatName(myName)
	secondClassName := second.FormatName(myName)

	fmt.Println(firstClassName, "sasas", secondClassName)



	fmt.Println(helpers.ToUpperCase("Amith"))

	data, studentErr := getStudent("Amith")
	if studentErr != nil {
		fmt.Println("Student error ", studentErr.Error())
	}
	fmt.Println(data)

	result, err := multi(2)

	if err == nil {
		fmt.Println("sasas", result)
	} else {
		fmt.Println(err)
	}


	// var name string
	name := "Amith!!"
	fmt.Println("HellO!!", name)

	printSome()
	printName(name)

	formmatedName := formatName(name)
	fmt.Println(formmatedName)

}