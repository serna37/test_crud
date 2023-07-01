package sql

import (
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
	GetCateTag(id int) []model.CateTag
	GetAllData(id int) []model.AllData
}

func NewUsr() Usr {
	return &usr{}
}

// ==================
// Imprementation
// ==================
func (usr *usr) Create(name string, loginid string, password string, authtoken string) {
	sql := `
		INSERT INTO mst_usr ("name", usr_login_id, usr_pass_word, auth_token, last_login)
		VALUES($1, $2, $3, $4, CURRENT_TIMESTAMP);
		`
	upd, _ := db.Prepare(sql)
	upd.Exec(name, loginid, password, authtoken)
}

func (usr *usr) Check(loginid string) bool {
	sql := `
	select id from mst_usr where usr_login_id = $1;
	`
	pp, _ := db.Prepare(sql)
	rows, _ := pp.Query(loginid)
	defer rows.Close()
	var res []int
	for rows.Next() {
		var row int
		rows.Scan(&row)
		res = append(res, row)
	}
	return len(res) == 0
}

func (usr *usr) Read(loginid string, password string) model.MstUsr {
	sql := `
		select
			id,
			"name",
			usr_login_id,
			usr_pass_word,
			auth_token,
			last_login
		from
			mst_usr
		where
			usr_login_id = $1
			and usr_pass_word = $2;
		`
	pp, _ := db.Prepare(sql)
	rows, _ := pp.Query(loginid, password)
	defer rows.Close()
	var res []model.MstUsr
	var row model.MstUsr
	for rows.Next() {
		rows.Scan(&row.Id, &row.Name, &row.UsrLoginId, &row.UsrPassWord, &row.AuthToken, &row.LastLogin)
		res = append(res, row)
	}
	if len(res) == 0 {
		return *model.NewMstUsr(-1, "", "", "", "", time.Now())
	}
	return res[0]
}

func (usr *usr) Update(id int, name string, loginid string, password string) {
	var sql string
	if name != "" {
		sql = `
		update
			mst_usr
		set
			"name" = $1,
			usr_login_id = $2,
			usr_pass_word = $3,
		where
			id = $4
		`
	} else {
		sql = `
		update
			mst_usr
		set
			last_login = CURRENT_TIMESTAMP
		where
			id = $1
		`
	}
	upd, _ := db.Prepare(sql)
	if name != "" {
		upd.Exec(name, loginid, password, id)
	} else {
		upd.Exec(id)
	}
}

func (usr *usr) Delete(id int) {
	sql := `
	delete
	from
		mst_usr
	where
		id = $1;
	`
	upd, _ := db.Prepare(sql)
	upd.Exec(id)
}

// have to reduce
func (usr *usr) GetCateTag(id int) []model.CateTag {
	sql := `
select
mc.id category_id,
mc.name category_name,
mt.id tag_id,
mt."name" tag_name
from mst_category mc 
left join mst_tag mt on mc.id = mt.category_id 
where mc.usr_id = $1
and mc.del_flg = false 
and mt.del_flg = false 
;
	`
	pp, _ := db.Prepare(sql)
	rows, _ := pp.Query(id)
	defer rows.Close()
	var res []model.CateTag
	for rows.Next() {
		var row model.CateTag
		rows.Scan(&row.CategoryId, &row.CategoryName, &row.TagId, &row.TagName)
		res = append(res, row)
	}
	return res
}

// have to reduce
func (usr *usr) GetAllData(id int) []model.AllData {
	sql := `
select
mc.id category_id,
mc.name category_name,
tc.id content_id,
tc.title content_title,
tc.contents content_body,
mt.id tag_id,
mt."name" tag_name
from mst_category mc
left join trn_contents tc on tc.category_id = mc.id
left join trn_contents_tag tct on tct.content_id = tc.id
left join mst_tag mt on mt.id = tct.tag_id
where mc.usr_id = $1
and mc.del_flg = false 
and tc.del_flg = false
and mt.del_flg = false
;
	`
	pp, _ := db.Prepare(sql)
	rows, _ := pp.Query(id)
	defer rows.Close()
	var res []model.AllData
	for rows.Next() {
		var row model.AllData
		rows.Scan(&row.CategoryId, &row.CategoryName, &row.ContentId, &row.ContentTitle, &row.ContentBody, &row.TagId, &row.TagName)
		res = append(res, row)
	}
	return res
}
