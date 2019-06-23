package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	u "github.com/mevlanaayas/bookia/utils"
	"time"
)

type Book struct {
	gorm.Model
	Name        string    `json:"name"`
	CreatedDate time.Time `json:"created_date"`
	CreatedUser string    `json:"created_user"`
	UpdatedDate time.Time `json:"updated_date"`
	UpdatedUser string    `json:"updated_user"`
}

/*
 This struct function validate the required parameters sent through the http request body
returns message and true if the requirement is met
*/
func (book *Book) Validate() (map[string]interface{}, bool) {

	if book.Name == "" {
		return u.Message(false, "Book name should be on the payload"), false
	}

	//All the required parameters are present
	return u.Message(true, "success"), true
}

func (book *Book) Create() map[string]interface{} {

	if resp, ok := book.Validate(); !ok {
		return resp
	}

	GetDB().Create(book)

	resp := u.Message(true, "success")
	resp["book"] = book
	return resp
}

func GetBook(id uint) *Book {

	book := &Book{}
	err := GetDB().Table("books").Where("id = ?", id).First(book).Error
	if err != nil {
		return nil
	}
	return book
}

func GetBooks(username string) []*Book {

	books := make([]*Book, 0)
	err := GetDB().Table("books").Where("created_by = ?", username).Find(&books).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return books
}
