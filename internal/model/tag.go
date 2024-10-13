package model

// 标签
type Tag struct {
	ID   int    `json:"id" gorm:"primaryKey;autoIncrement"` // 标签唯一标识符
	Name string `json:"name" gorm:"unique;not null"`        // 标签名称，唯一且不能为空
}
