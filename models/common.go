package models

func Count(tableName string) (count int) {
	GetDB().Table(tableName).Count(&count)
	return
}
