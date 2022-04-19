package controllers

import "strconv"

type Student struct {
	Id    int
	Name  string
	Grade int32
}

var students = []*Student{
	{Id: 1, Name: "angga", Grade: 2},
	{Id: 2, Name: "denta", Grade: 2},
	{Id: 3, Name: "fianos", Grade: 3},
}

func GetStudents() []*Student {
	return students
}

func SelectStudent(i string) *Student {
	id, _ := strconv.Atoi(i)
	for _, each := range students {
		if each.Id == id {
			return each
		}
	}

	return nil
}
