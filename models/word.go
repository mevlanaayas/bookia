package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	u "github.com/mevlanaayas/bookia/utils"
)

type Word struct {
	gorm.Model
	Word        string `json:"word"`
	BookId      int    `json:"book_id"`
	Sentence    string `json:"sentence"`
	Translate   string `json:"translate"`
	CreatedUser string `json:"created_user"`
	UpdatedUser string `json:"updated_user"`
}

/*
 This struct function validate the required parameters sent through the http request body
returns message and true if the requirement is met
*/
func (word *Word) Validate() (map[string]interface{}, bool) {

	if word.Word == "" {
		return u.Message(false, "Word should be on the payload"), false
	}

	if word.Sentence == "" {
		return u.Message(false, "Sentence should be on the payload"), false
	}

	if word.Translate == "" {
		return u.Message(false, "Translate should be on the payload"), false
	}

	//All the required parameters are present
	return u.Message(true, "success"), true
}

func (word *Word) Create() map[string]interface{} {

	if resp, ok := word.Validate(); !ok {
		return resp
	}

	GetDB().Create(word)

	resp := u.Message(true, "success")
	resp["word"] = word
	return resp
}

func GetWord(id uint) *Word {

	word := &Word{}
	err := GetDB().Table("words").Where("id = ?", id).First(word).Error
	if err != nil {
		return nil
	}
	return word
}

func GetAllWords() []*Word {

	words := make([]*Word, 0)
	err := GetDB().Table("words").Find(&words).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return words
}

func WordCount() int {
	return Count("words")
}

func GetWords(username string) []*Word {

	words := make([]*Word, 0)
	err := GetDB().Table("words").Where("created_user = ?", username).Find(&words).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return words
}

func GetWordsByBook(username string, book_id int) []*Word {

	words := make([]*Word, 0)
	err := GetDB().Table("words").Where("created_user = ?", username).Find(&words).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return words
}
