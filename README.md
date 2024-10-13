# STBlog 博客系统

## 技术架构

后端：

- go
- gin
- gorm
- zap

前端：
- vue3
- vite
- vue-router
- element-plus

## 模型设计

```go
package models

import "time"

type User struct {
    ID          int       `json:"id" gorm:"primaryKey"`
    Username    string    `json:"username" gorm:"unique"`
    PasswordHash string    `json:"-"`
    Email       string    `json:"email" gorm:"unique"`
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
}

type Post struct {
    ID        int       `json:"id" gorm:"primaryKey"`
    Title     string    `json:"title"`
    Content   string    `json:"content"`
    AuthorID  int       `json:"author_id"` // 存储作者的用户ID
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
    Published bool      `json:"published"`
}

type Comment struct {
    ID        int       `json:"id" gorm:"primaryKey"`
    PostID    int       `json:"post_id"` // 存储评论对应的文章ID
    UserID    int       `json:"user_id"` // 存储评论者的用户ID
    Content   string    `json:"content"`
    CreatedAt time.Time `json:"created_at"`
}


type Tag struct {
    ID   int    `json:"id" gorm:"primaryKey"`
    Name string `json:"name" gorm:"unique"`
}


type PostTag struct {
    PostID int `json:"post_id"` // 存储文章ID
    TagID  int `json:"tag_id"`  // 存储标签ID
}

```

### 功能模块

1、**用户管理**

- 注册、登录、修改密码、查看用户信息。

2、**文章管理**

创建、编辑、删除、查看文章，支持草稿和发布状态。

3、**评论管理**

- 添加、删除评论，查看评论列表。

4、**标签管理**

- 创建、删除标签，给文章添加标签。

5、**搜索与过滤**

- 根据标签、作者、时间等条件搜索文章。