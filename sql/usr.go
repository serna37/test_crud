package sql

import (
	"log"
	"test_crud/model"
	"time"
)

// ==================
// struct def
// ==================
type usr struct {
}
type Usr interface {
	Create(name string, loginid string, password string, authtoken string)
	Check(loginid string) bool
	Read(loginid string, password string) model.MstUsr
	Update(id int, name string, loginid string, password string)
	Delete(id int)
	GetCateTag(id int) []model.CateTags
	GetAllData(id int) []model.AllData
}

func NewUsr() Usr {
	return &usr{}
}

// ==================
// Imprementation
// ==================
func (usr *usr) Create(name string, loginid string, password string, authtoken string) {
	record := model.MstUsr{Name: name, UsrLoginId: loginid, UsrPassWord: password, AuthToken: authtoken, LastLogin: time.Now()}
	result := db.Create(&record)
	if result.Error != nil {
		log.Fatal(result.Error.Error())
	}
}

func (usr *usr) Check(loginid string) bool {
	var row model.MstUsr
	db.Where("usr_login_id = ?", loginid).First(&row)
	if row.Name != "" {
		return false
	}
	return true
}

func (usr *usr) Read(loginid string, password string) model.MstUsr {
	var row model.MstUsr
	db.Where("usr_login_id = ? AND usr_pass_word = ?", loginid, password).First(&row)
	return row
}

func (usr *usr) Update(id int, name string, loginid string, password string) {
	if name != "" {
		// update user info
		var target model.MstUsr
		db.First(&target, id)
		target.Name = name
		target.UsrLoginId = loginid
		target.UsrPassWord = password
		result := db.Save(&target)
		if result.Error != nil {
			log.Fatal(result.Error.Error())
		}
	} else {
		// login
		var target model.MstUsr
		db.First(&target, id)
		target.LastLogin = time.Now()
		result := db.Save(&target)
		if result.Error != nil {
			log.Fatal(result.Error.Error())
		}
	}
}

func (usr *usr) Delete(id int) {
	db.Delete(&model.MstUsr{}, id)
}

func (usr *usr) GetCateTag(id int) []model.CateTags {
	var cate []model.CateTags
	db.Preload("MstTags", "del_flg = ?", false).Find(&cate, "mst_category.usr_id = ? AND mst_category.del_flg = false", id)
	return cate
}

func (usr *usr) GetAllData(id int) []model.AllData {
	var result []model.AllData
	db.Preload("AllContent", "trn_contents.del_flg = ?", false).Preload("AllContent.Tags.MstTags", "mst_tag.del_flg = ?", false).Find(&result, "mst_category.usr_id = ? AND mst_category.del_flg = false", id)
	return result
}
