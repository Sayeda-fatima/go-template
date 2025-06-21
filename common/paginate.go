package common

import (
	"fmt"
	"strconv"

	"gorm.io/gorm"
)

func Paginate(page int, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page <= 0 {
			page = 1
		}
		switch {
		case pageSize > 100:
			pageSize = 100

		case pageSize <= 0:
			pageSize = 10
		}
		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}
func RawPaginate(page int, pageSize int) string {
	if page <= 0 {
		page = 1
	}
	switch {
	case pageSize > 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}
	offset := (page - 1) * pageSize
	return " LIMIT " + strconv.Itoa(pageSize) + " OFFSET " + strconv.Itoa(offset)
}

func KeySetPaginate(column string, lastValue *uint, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if pageSize > 100 {
			pageSize = 100
		} else if pageSize <= 0 {
			pageSize = 10
		}
		if lastValue != nil {
			db = db.Where(column+" > ?", lastValue)
		}
		return db.Order(column + " ASC").Limit(pageSize)
	}
}
func KeySetPaginateRaw(column string, lastValue *uint, pageSize int) (string, []interface{}) {
	if pageSize > 100 {
		pageSize = 100
	} else if pageSize <= 0 {
		pageSize = 10
	}

	var conditions string
	var args []interface{}

	if lastValue != nil {
		conditions = fmt.Sprintf("AND %s > ?", column)
		args = append(args, *lastValue)
	}

	// Append ORDER BY and LIMIT
	conditions += fmt.Sprintf(" ORDER BY %s ASC LIMIT ?", column)
	args = append(args, pageSize)

	return conditions, args
}
