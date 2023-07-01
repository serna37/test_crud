package sql

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
	sql := `
	INSERT INTO trn_contents (usr_id, contents, category_id, del_flg, title)
	VALUES($1, $2, $3, false, $4);
	`
	upd, _ := db.Prepare(sql)
	upd.Exec(userid, content, categoryid, title)
}

func (contents *contents) Update(content string, title string, id int) {
	sql := `
	update
		trn_contents
	set
		contents = $1,
		title = $2
	where
		id = $3;
	`
	upd, _ := db.Prepare(sql)
	upd.Exec(content, title, id)
}

func (contents *contents) Delete(id int) {
	sql := `
	update
		trn_contents
	set
		del_flg = true,
	where
		id = $1;
	`
	upd, _ := db.Prepare(sql)
	upd.Exec(id)
}
