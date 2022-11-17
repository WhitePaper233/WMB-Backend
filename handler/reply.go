package handler

import (
	"math/rand"
	"time"

	"whitepaper233.top/WMBBackend/model"
	"whitepaper233.top/WMBBackend/serializer"
	"whitepaper233.top/WMBBackend/service"
)

func HandleQueryUnfoldedReply(CommentID uint64) serializer.RespWithData {
	result := service.GetReply(CommentID, 0)
	ret := serializer.SerializeReply(200, "", result)
	return ret
}

func HandleQueryReplyByPage(CommentID uint64, Page uint64) serializer.RespWithData {
	result := service.GetReply(CommentID, Page)
	ret := serializer.SerializeReply(200, "", result)
	return ret
}

func HandleCountReply(CommentID uint64) serializer.RespWithData {
	result := service.CountReply(CommentID)
	ret := serializer.SerializeWithData(200, "", result)
	return ret
}

func HandleNewReply(CommentID uint64, Nickname string, Email string, Content string) {
	replyData := new(model.Reply)
	rand.Seed(time.Now().Unix())
	replyData.Avatar = rand.Int63n(11)
	replyData.CommentID = CommentID
	replyData.Content = Content
	replyData.Nickname = Nickname
	replyData.Email = Email
	replyData.ID = 0
	replyData.Timestamp = time.Now().Unix()

	service.NewReply(replyData)
}
