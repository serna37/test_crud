package model

type CateTag struct {
	CategoryId int `db:"category_id"`
	CategoryName string `db:"category_name"`
	TagId int `db:"tag_id"`
	TagName string `db:"tag_name"`
}

type AllData struct {
	CategoryId int `db:"category_id"`
	CategoryName string `db:"category_name"`
	ContentId int `db:"content_id"`
	ContentTitle string `db:"content_title"`
	ContentBody string `db:"content_body"`
	TagId int `db:"tag_id"`
	TagName string `db:"tag_name"`
}
