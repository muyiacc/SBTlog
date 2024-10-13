package model

// 文章和标签
type PostTag struct {
	PostID int `json:"post_id" gorm:"not null"` // 文章ID，不能为空
	TagID  int `json:"tag_id" gorm:"not null"`  // 标签ID，不能为空
}
