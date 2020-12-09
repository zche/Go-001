package controllers

import (
	"fmt"
	"week02/dtos"
	"week02/svcs"
)

// GetStudent 获取单个学生数据
func GetStudent(id int) dtos.Student {
	student, err := svcs.GetStudentByID(id)
	if err == nil {
		return *student
	}
	// panic(err)
	// u, ok := err.(interface{ Unwrap() error })
	// if ok {
	// 	err = u.Unwrap()
	// }

	fmt.Printf("main:%+v\n", err)
	return *student
}
