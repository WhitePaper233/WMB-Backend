// package models database models
package model

type Comment struct {
	ID        uint64 `gorm:"primarykey"`
	Avatar    int64  `gorm:"avatar"`
	Nickname  string `gorm:"nickname"`
	Timestamp int64  `gorm:"timestamp"`
	Email     string `gorm:"email"`
	Content   string `gorm:"content"`
}

type Reply struct {
	Comment
	CommentID uint64 `gorm:"comment_id"`
	ReplyID   int64  `gorm:"reply_id"`
}
