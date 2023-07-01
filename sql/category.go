package sql

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
	sql := `
	INSERT INTO mst_category (usr_id, "name", del_flg)
	VALUES($1, $2, false);
	`
	upd, _ := db.Prepare(sql)
	upd.Exec(userid, name)
}

func (category *category) Update(id int, name string) {
	sql := `
	update
		mst_category
	set
		"name" = $1,
	where
		id = $2
	`
	upd, _ := db.Prepare(sql)
	upd.Exec(name, id)
}

func (category *category) Delete(id int) {
	sql := `
	update
		mst_category
	set
		del_flg = true
	where
		id = $1
	`
	upd, _ := db.Prepare(sql)
	upd.Exec(id)
}
