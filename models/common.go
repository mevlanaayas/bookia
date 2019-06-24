package models

import "fmt"

func Count(tableName string) (count int) {
	err := GetDB().Table(tableName).Count(&count)
	if err != nil {
		fmt.Println(err)
		return 0
	}
	return
}
