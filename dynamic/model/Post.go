package model

type Attachment struct {
	Type  string `bson:"type"`
	Label string `bson:"label"`
	Url   string `bson:"url"`
}

type Tag struct {
	Label string `bson:"label"` // 标签名
	Value string `bson:"value"` // id
}

type Post struct {
	Id          string        `bson:"_id"`
	UserId      string        `bson:"userid"`
	Contents    string        `bson:"contents"`
	Attachments []*Attachment `bson:"attachments"`
	Tags        []*Tag        `bson:"tags"`
	CreatedTime int64         `bson:"createdtime"` // 创建时间
	UpdatedTime int64         `bson:"updatedtime"` // 修改时间
	DeletedTime int64         `bson:"deletedtime"` // 删除时间
}
