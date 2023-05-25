package model

type User struct {
	Id          string `bson:"_id"` // 用户ID
	Nickname    string // 昵称
	Username    string // 用户名
	Password    string // MD5密码
	Avatar      string // 用户头像
	IsAdmin     bool   // 是否管理员
	CreatedTime int64  // 创建时间
	UpdatedTime int64  // 修改时间
	DeletedTime int64  // 删除时间
}
