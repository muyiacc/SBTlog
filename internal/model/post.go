package model

import "time"

// 文章
type Post struct {
	ID        int       `json:"id" gorm:"primaryKey;autoIncrement"` // 文章唯一标识符
	Title     string    `json:"title" gorm:"not null"`              // 文章标题，不能为空
	Content   string    `json:"content" gorm:"type:text"`           // 文章内容，使用文本类型
	AuthorID  int       `json:"author_id" gorm:"not null"`          // 作者的用户ID，不能为空
	CreatedAt time.Time `json:"created_at" gorm:"default:now()"`    // 文章创建时间，默认当前时间
	UpdatedAt time.Time `json:"updated_at" gorm:"default:now()"`    // 文章更新时间，默认当前时间
	Published bool      `json:"published" gorm:"default:false"`     // 文章是否已发布，默认未发布
}
