package model

type Tag struct {
	Id          string `bson:"_id"` // 标签ID
	Tag         string // 标签名
	Avatar      string // tag图标
	QuoteNum    int64  // 引用数
	CreatedTime int64  // 创建时间
	UpdatedTime int64  // 修改时间
	DeletedTime int64  // 删除时间
}
