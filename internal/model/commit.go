package model

import "time"

// 评论
type Comment struct {
	ID        int       `json:"id" gorm:"primaryKey;autoIncrement"` // 评论唯一标识符
	PostID    int       `json:"post_id" gorm:"not null"`            // 评论对应的文章ID，不能为空
	UserID    int       `json:"user_id" gorm:"not null"`            // 评论者的用户ID，不能为空
	Content   string    `json:"content" gorm:"not null"`            // 评论内容，不能为空
	CreatedAt time.Time `json:"created_at" gorm:"default:now()"`    // 评论创建时间，默认当前时间
}
