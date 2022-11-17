package handler

import (
	"math/rand"
	"time"

	"whitepaper233.top/WMBBackend/model"
	"whitepaper233.top/WMBBackend/serializer"
	"whitepaper233.top/WMBBackend/service"
)

func HandleQueryLatestComment() serializer.RespWithData {
	result := service.GetComment(0)
	ret := serializer.SerializeComment(200, "", result)
	return ret
}

func HandleQueryCommentByPage(Page uint64) serializer.RespWithData {
	result := service.GetComment(Page)
	ret := serializer.SerializeComment(200, "", result)
	return ret
}

func HandleCountComment() serializer.RespWithData {
	result := service.CountComment()
	ret := serializer.SerializeWithData(200, "", result)
	return ret
}

func HandleNewComment(Nickname string, Email string, Content string) {
	commentData := new(model.Comment)
	rand.Seed(time.Now().Unix())
	commentData.Avatar = rand.Int63n(11)
	commentData.Content = Content
	commentData.Nickname = Nickname
	commentData.Email = Email
	commentData.ID = 0
	commentData.Timestamp = time.Now().Unix()

	service.NewComment(commentData)
}
