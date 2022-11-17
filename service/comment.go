package service

import (
	"whitepaper233.top/WMBBackend/database"
	"whitepaper233.top/WMBBackend/model"
)


func GetComment(Page uint64) []model.Comment {
	db := database.GetDB()
	var comments []model.Comment
	if Page == 0 {
		db.Limit(5).Order("timestamp desc").Find(&comments)
	} else {
		var latestComment model.Comment
		db.Limit(1).Order("timestamp desc").First(&latestComment)

		if latestComment.ID - (Page * 10) < latestComment.ID - ((Page - 1) * 10) {
			db.Limit(10).
			Where("id > ? AND id <= ?", latestComment.ID - (Page * 10), latestComment.ID - ((Page - 1) * 10)).
			Order("timestamp desc").
			Find(&comments)
		} else {
			db.Limit(10).
			Where("id > ? AND id <= ?", 0, latestComment.ID - ((Page - 1) * 10)).
			Order("timestamp desc").
			Find(&comments)
		}
	}
	
	return comments
}

func CountComment() CountData {
	db := database.GetDB()

	var countData CountData
	db.Model(&model.Comment{}).Count(&countData.Count)

	return countData
}

func NewComment(data *model.Comment) {
	db := database.GetDB()

	db.Create(data)
}
