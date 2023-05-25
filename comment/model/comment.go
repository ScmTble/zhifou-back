package model

type Comment struct {
	Id          string `bson:"_id"` // 评论id
	PostId      string // 动态id
	UserId      string // 评论人id
	Content     string // 评论
	CreatedTime int64
	UpdatedTime int64
	DeletedTime int64
}
