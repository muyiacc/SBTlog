package model

import "time"

// 用户
type User struct {
	ID           int       `json:"id" gorm:"primaryKey;autoIncrement"` // 用户唯一标识符
	Username     string    `json:"username" gorm:"unique;not null"`    // 用户名，唯一且不能为空
	PasswordHash string    `json:"-"`                                  // 哈希后的密码，隐私信息，不返回给客户端
	Email        string    `json:"email" gorm:"unique;not null"`       // 用户邮箱，唯一且不能为空
	CreatedAt    time.Time `json:"created_at" gorm:"default:now()"`    // 用户创建时间，默认当前时间
	UpdatedAt    time.Time `json:"updated_at" gorm:"default:now()"`    // 用户信息更新时间，默认当前时间
}
