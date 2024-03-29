package model

import (
	"time"
)

type MstUsr struct {
	Id          int       // ID
	Name        string    // 名前
	UsrLoginId  string    // ログインID
	UsrPassWord string    // パスワード
	AuthToken   string    // 認証トークン
	LastLogin   time.Time // 最終ログイン日時
}

func (MstUsr) TableName() string {
	return "mst_usr"
}

func NewMstUsr(id int, name string, usrLoginId string, usrPassWord string, authToken string, lastLogin time.Time) *MstUsr {
	u := new(MstUsr)
	u.Id = id
	u.Name = name
	u.UsrLoginId = usrLoginId
	u.UsrPassWord = usrPassWord
	u.AuthToken = authToken
	u.LastLogin = lastLogin
	return u
}

type MstCategory struct {
	Id     int    // カテゴリID
	UsrId  int    // ユーザID
	Name   string // カテゴリ名
	DelFlg bool   // 削除フラグ
}

func (MstCategory) TableName() string {
	return "mst_category"
}

func NewMstCategory(id int, usrId int, name string, delFlg bool) *MstCategory {
	u := new(MstCategory)
	u.Id = id
	u.UsrId = usrId
	u.Name = name
	u.DelFlg = delFlg
	return u
}

type MstTag struct {
	Id         int    // タグID
	UsrId      int    // ユーザID
	Name       string // タグ名
	DelFlg     bool   // 削除フラグ
	CategoryId int    // カテゴリID
}

func (MstTag) TableName() string {
	return "mst_tag"
}

func NewMstTag(id int, usrId int, name string, delFlg bool, categoryId int) *MstTag {
	u := new(MstTag)
	u.Id = id
	u.UsrId = usrId
	u.Name = name
	u.DelFlg = delFlg
	u.CategoryId = categoryId
	return u
}

type TrnContents struct {
	Id         int    // コンテンツID
	UsrId      int    // ユーザID
	Contents   string // コンテンツ
	CategoryId int    // カテゴリID
	DelFlg     bool   // 削除フラグ
	Title      string // タイトル
}

func (TrnContents) TableName() string {
	return "trn_contents"
}

func NewTrnContents(id int, usrId int, contents string, categoryId int, delFlg bool, title string) *TrnContents {
	u := new(TrnContents)
	u.Id = id
	u.UsrId = usrId
	u.Contents = contents
	u.CategoryId = categoryId
	u.DelFlg = delFlg
	u.Title = title
	return u
}

type TrnContentsTag struct {
	Id        int // タグ付けID
	ContentId int // コンテンツID
	TagId     int // タグID
}

func (TrnContentsTag) TableName() string {
	return "trn_contents_tag"
}

func NewTrnContentsTag(id int, contentId int, tagId int) *TrnContentsTag {
	u := new(TrnContentsTag)
	u.Id = id
	u.ContentId = contentId
	u.TagId = tagId
	return u
}

type TrnChatRoom struct {
	Id        int       // チャットルームID
	Name      string    // ルーム名
	CreaterId int       // 作成者ID
	DelFlg    bool      // 削除フラグ
	CreatedAt time.Time // 作成日時
}

func (TrnChatRoom) TableName() string {
	return "trn_chat_room"
}

func NewTrnChatRoom(id int, name string, createrId int, delFlg bool, createdAt time.Time) *TrnChatRoom {
	u := new(TrnChatRoom)
	u.Id = id
	u.Name = name
	u.CreaterId = createrId
	u.DelFlg = delFlg
	u.CreatedAt = createdAt
	return u
}

type TrnChatRoomMember struct {
	Id         int       // チャットルーム参加者管理ID
	ChatRoomId int       // チャットルームID
	JoinerId   int       // 参加者ID
	DelFlg     bool      // 削除フラグ
	JoinedAt   time.Time // 参加日時
	Inviter    int       // 招待者
}

func (TrnChatRoomMember) TableName() string {
	return "trn_chat_room_member"
}

func NewTrnChatRoomMember(id int, chatRoomId int, joinerId int, delFlg bool, joinedAt time.Time, inviter int) *TrnChatRoomMember {
	u := new(TrnChatRoomMember)
	u.Id = id
	u.ChatRoomId = chatRoomId
	u.JoinerId = joinerId
	u.DelFlg = delFlg
	u.JoinedAt = joinedAt
	u.Inviter = inviter
	return u
}

type TrnChatMsg struct {
	Id         int       // チャットルーム参加者管理ID
	ChatRoomId int       // チャットルームID
	FromId     int       // 発言者ID
	FromAt     time.Time // 発言日時
	Msg        string    // メッセージ
}

func (TrnChatMsg) TableName() string {
	return "trn_chat_msg"
}

func NewTrnChatMsg(id int, chatRoomId int, fromId int, fromAt time.Time, msg string) *TrnChatMsg {
	u := new(TrnChatMsg)
	u.Id = id
	u.ChatRoomId = chatRoomId
	u.FromId = fromId
	u.FromAt = fromAt
	u.Msg = msg
	return u
}
