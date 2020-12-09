package db

import "database/sql"

// QueryDb 查询记录
func QueryDb(id int) (interface{}, error) {
	return nil, sql.ErrNoRows
}
