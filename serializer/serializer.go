// package serializer api return serializer
package serializer

import (
	"whitepaper233.top/WMBBackend/model"
	"whitepaper233.top/WMBBackend/util"
)

type Resp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type RespWithData struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type Comment struct {
	CommentID    uint64 `json:"comment_id"`
	Avatar       int64  `json:"avatar"`
	Nickname     string `json:"nickname"`
	Time         string `json:"time"`
	EmailAddress string `json:"email_address"`
	Content      string `json:"content"`
}

func (comment *Comment) Construct(commentModel model.Comment) Comment {
	comment.CommentID = commentModel.ID
	comment.Avatar = commentModel.Avatar
	comment.Nickname = commentModel.Nickname
	comment.Time = util.FormatTime(commentModel.Timestamp)
	comment.EmailAddress = commentModel.Email
	comment.Content = commentModel.Content
	return *comment
}

type Reply struct {
	Comment
	ReplyID int64 `json:"reply_id"`
}

func (reply *Reply) Construct(replyModel model.Reply) Reply {
	reply.Comment = reply.Comment.Construct(replyModel.Comment)
	reply.ReplyID = replyModel.ReplyID
	return *reply
}

func Serialize(Code int, Msg string) Resp {
	var resp Resp
	resp.Code = Code
	resp.Message = Msg
	return resp
}

func SerializeWithData(Code int, Msg string, Data interface{}) RespWithData {
	var resp RespWithData
	resp.Code = Code
	resp.Message = Msg
	resp.Data = Data
	return resp
}

func SerializeComment(Code int, Msg string, CommentData []model.Comment) RespWithData {
	var retList = make([]Comment, len(CommentData))
	for index, value := range CommentData {
		retList[index].Construct(value)
	}
	var resp RespWithData
	resp.Code = Code
	resp.Message = Msg
	resp.Data = retList

	return resp
}

func SerializeReply(Code int, Msg string, ReplyData []model.Reply) RespWithData {
	var retList = make([]Reply, len(ReplyData))
	for index, value := range ReplyData {
		retList[index].Construct(value)
	}
	var resp RespWithData
	resp.Code = Code
	resp.Message = Msg
	resp.Data = retList

	return resp
}
