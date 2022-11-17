package service

import (
	"errors"

	"gorm.io/gorm"
	"whitepaper233.top/WMBBackend/database"
	"whitepaper233.top/WMBBackend/model"
)


func GetReply(CommentID uint64, Page uint64) []model.Reply {
	db := database.GetDB()

	var replies []model.Reply
	if Page == 0 {
		db.Limit(3).Order("timestamp desc").Where("comment_id = ?", CommentID).Find(&replies)
	} else {
		db.Limit(5).
		Where("comment_id = ? AND reply_id > ? AND reply_id <= ?", CommentID, ((Page - 1) * 5), Page * 5).
		Order("timestamp").
		Find(&replies)
	}

	return replies
}

func CountReply(CommentID uint64) CountData {
	db := database.GetDB()

	var countData CountData
	db.Model(&model.Reply{}).Where("comment_id = ?", CommentID).Count(&countData.Count)

	return countData
}

func NewReply(data *model.Reply) {
	db := database.GetDB()

	var latestReply model.Reply
	err := db.Where("comment_id = ?", data.CommentID).Order("timestamp desc").First(&latestReply).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		data.ReplyID = 1
	} else {
		data.ReplyID = latestReply.ReplyID + 1
	}

	db.Create(data)
}
