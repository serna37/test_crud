package model

// get category > tags
type CateTags struct {
	Id      int
	UsrId   int
	Name    string
	DelFlg  bool
	MstTags []MstTag `gorm:"foreignKey:CategoryId"`
}

func (CateTags) TableName() string { return "mst_category" }

// get category > contents > tags
type AllData struct {
	Id         int `gorm:"references:id"`
	UsrId      int
	Name       string
	DelFlg     bool
	AllContent []ContentTags `gorm:"foreignKey:CategoryId"`
}

func (AllData) TableName() string { return "mst_category" }

type ContentTags struct {
	Id         int    `gorm:"references:id"`
	UsrId      int    // ユーザID
	Contents   string // コンテンツ
	CategoryId int
	DelFlg     bool      // 削除フラグ
	Title      string    // タイトル
	Tags       []RefTags `gorm:"foreignKey:ContentId"`
}

func (ContentTags) TableName() string { return "trn_contents" }

type RefTags struct {
	Id        int    // タグ付けID
	ContentId int    // コンテンツID
	TagId     int    `gorm:"references:tag_id"`
	MstTags   MstTag `gorm:"foreignKey:Id"`
}

func (RefTags) TableName() string { return "trn_contents_tag" }
