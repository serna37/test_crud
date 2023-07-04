package sql

import (
	"log"
	"test_crud/model"
)

// ==================
// struct def
// ==================
type contents struct {
}
type Contents interface {
	Create(userid int, content string, categoryid int, title string)
	Update(content string, title string, id int)
	Delete(id int)
}

func NewContents() Contents {
	return &contents{}
}

// ==================
// Imprementation
// ==================
func (contents *contents) Create(userid int, content string, categoryid int, title string) {
	record := model.TrnContents{UsrId: userid, Contents: content, CategoryId: categoryid, Title: title}
	result := db.Create(&record)
	if result.Error != nil {
		log.Fatal(result.Error.Error())
	}
}

func (contents *contents) Update(content string, title string, id int) {
	var target model.TrnContents
	db.First(&target, id)
	target.Contents = content
	target.Title = title
	result := db.Save(&target)
	if result.Error != nil {
		log.Fatal(result.Error.Error())
	}
}

func (contents *contents) Delete(id int) {
	db.Delete(&model.TrnContents{}, id)
}
