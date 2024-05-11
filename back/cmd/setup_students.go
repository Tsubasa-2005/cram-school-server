package main

import (
	"cram-school-reserve-server/back/infra/rdb"
	"fmt"
)

func main() {
	// 学生のインスタンスを作成
	students := []rdb.Student{
		{ID: "a002", Name: "Student One", Password: "0"},
		{ID: "b002", Name: "Student Two", Password: "0"},
		{ID: "c002", Name: "Student Three", Password: "0"},
		{ID: "d002", Name: "Student Four", Password: "0"},
		{ID: "e002", Name: "Student Five", Password: "0"},
	}

	// データベースに各学生を追加
	for _, student := range students {
		err := rdb.CreateStudent(student)
		if err != nil {
			fmt.Println("Error adding student:", err)
		} else {
			fmt.Println("Student added successfully:", student.Name)
		}
	}
}
