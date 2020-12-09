package dao

import (
	"database/sql"
	"week02/db"

	"github.com/pkg/errors"
)

// QueryByID 根据id查询数据
func QueryByID(id int) (interface{}, error) {
	data, err := db.QueryDb(id)
	if err != nil {
		return data, errors.Wrapf(sql.ErrNoRows, "查询db 返回为nil，id为%d", id)
	}
	return data, nil
}
