package sql

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
	sql := `
	INSERT INTO mst_tag (usr_id, "name", del_flg, category_id)
	VALUES($1, $2, false, $3);
	`
	upd, _ := db.Prepare(sql)
	upd.Exec(userid, name, categoryid)
}

func (tag *tag) Update(id int, name string) {
	sql := `
	update
		mst_tag
	set
		"name" = $1
	where
		id = $2;
	`
	upd, _ := db.Prepare(sql)
	upd.Exec(name, id)
}

func (tag *tag) Delete(id int) {
	sql := `
	update
		mst_tag
	set
		del_flg = false
	where
		id = $1;
	`
	upd, _ := db.Prepare(sql)
	upd.Exec(id)
}
