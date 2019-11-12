package structs_example


type StudentExample struct {
	Name string
	RollNo int
}

func (student StudentExample) GetStudentName() string {
	return student.Name
}

func(student StudentExample) UpdateNameFake(input string) string {
	student.Name = input // scope
	return student.Name
}

func(student *StudentExample) UpdateName(input string) string {
	student.Name = input // scope
	return student.Name
}