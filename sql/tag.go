package sql

import (
	"log"
	"test_crud/model"
)

// ==================
// struct def
// ==================
type tag struct {
}
type Tag interface {
	Create(userid int, name string, categoryid int)
	Update(id int, name string)
	Delete(id int)
}

func NewTag() Tag {
	return &tag{}
}

// ==================
// Imprementation
// ==================
func (tag *tag) Create(userid int, name string, categoryid int) {
	record := model.MstTag{UsrId: userid, Name: name, CategoryId: categoryid}
	result := db.Create(&record)
	if result.Error != nil {
		log.Fatal(result.Error.Error())
	}
}

func (tag *tag) Update(id int, name string) {
	var target model.MstTag
	db.First(&target, id)
	target.Name = name
	result := db.Save(&target)
	if result.Error != nil {
		log.Fatal(result.Error.Error())
	}
}

func (tag *tag) Delete(id int) {
	db.Delete(&model.MstTag{}, id)
}
