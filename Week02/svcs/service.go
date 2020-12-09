package svcs

import (
	"database/sql"
	"errors"
	"week02/dao"
	"week02/dtos"
)

// GetStudentByID 根据id获取学生数据
func GetStudentByID(id int) (*dtos.Student, error) {
	data, err := dao.QueryByID(id)
	if errors.Is(err, sql.ErrNoRows) {
		return &dtos.Student{ID: "111", Name: "check", Age: 28}, err
	}
	return data.(*dtos.Student), err
}
