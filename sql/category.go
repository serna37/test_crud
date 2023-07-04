package sql

import (
	"log"
	"test_crud/model"
)

// ==================
// struct def
// ==================
type category struct {
}
type Category interface {
	Create(userid int, name string)
	Update(id int, name string)
	Delete(id int)
}

func NewCategory() Category {
	return &category{}
}

// ==================
// Imprementation
// ==================
func (category *category) Create(userid int, name string) {
	record := model.MstCategory{UsrId: userid, Name: name}
	result := db.Create(&record)
	if result.Error != nil {
		log.Fatal(result.Error.Error())
	}
}

func (category *category) Update(id int, name string) {
	var target model.MstCategory
	db.First(&target, id)
	target.Name = name
	result := db.Save(&target)
	if result.Error != nil {
		log.Fatal(result.Error.Error())
	}
}

func (category *category) Delete(id int) {
	db.Delete(&model.MstCategory{}, id)
}
