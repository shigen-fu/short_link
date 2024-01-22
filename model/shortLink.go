package model

type ShortLink struct {
	Id         uint64 `gorm:"primaryKey;column:id" json:"id"`                // 主键ID
	OriginURL  string `gorm:"not null;column:origin_url" json:"originUrl"`   // 原始链接
	ShortCode  string `gorm:"not null;column:short_code" json:"shortCode"`   // 短链接代码
	CreateTime int64  `gorm:"not null;column:create_time" json:"createTime"` // 创建时间（时间戳）
	Creator    uint64 `gorm:"not null;column:creator" json:"creator"`        // 创建者
}

func (ShortLink) TableName() string {
	return "short_link"
}
