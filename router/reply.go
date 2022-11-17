package router

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/mojocn/base64Captcha"

	"whitepaper233.top/WMBBackend/handler"
	"whitepaper233.top/WMBBackend/serializer"
)


type ReplyPostData struct {
	CaptchaVerifyData struct {
		VerifyData
	} `json:"captcha_verify_data"`
	ReplyData struct {
		// Avatar    int64  `json:"avatar"`
		CommentID uint64 `json:"comment_id"`
		Nickname  string `json:"nickname"`
		Email     string `json:"email"`
		Content   string `json:"content"`
	} `json:"reply_data"`
}

func RouteQueryUnfoldedReply(c *fiber.Ctx) error {
	comment_id, err := strconv.ParseUint(c.Params("comment_id"), 10, 64)
	if err != nil {
		return c.Status(500).JSON(
			serializer.Serialize(500, err.Error()),
		)
	}

	return c.Status(200).JSON(
		handler.HandleQueryUnfoldedReply(comment_id),
	)
}

func RouteQueryReplyByPage(c *fiber.Ctx) error {
	comment_id, err := strconv.ParseUint(c.Params("comment_id"), 10, 64)
	if err != nil {
		return c.Status(500).JSON(
			serializer.Serialize(500, err.Error()),
		)
	}
	page, err := strconv.ParseUint(c.Params("page"), 10, 64)
	if err != nil {
		return c.Status(500).JSON(
			serializer.Serialize(500, err.Error()),
		)
	}

	return c.Status(200).JSON(
		handler.HandleQueryReplyByPage(comment_id, page),
	)
}

func RouteReplyCount(c *fiber.Ctx) error {
	comment_id, err := strconv.ParseUint(c.Params("comment_id"), 10, 64)
	if err != nil {
		return c.Status(500).JSON(
			serializer.Serialize(500, err.Error()),
		)
	}

	return c.Status(200).JSON(
		handler.HandleCountReply(comment_id),
	)
}

func RouteNewReply(store base64Captcha.Store) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		replyPostData := new(ReplyPostData)

		if err := c.BodyParser(replyPostData); err != nil {
			return c.Status(500).JSON(
				serializer.Serialize(500, err.Error()),
			)
		}

		if handler.HandleCaptchaVerify(
			store,
			replyPostData.CaptchaVerifyData.CaptchaID,
			replyPostData.CaptchaVerifyData.CaptchaAnswer,
		) {

			if replyPostData.ReplyData.Content == "" {
				return c.Status(400).JSON(
					serializer.Serialize(400, "Content can't be empty"),
				)
			}
			if replyPostData.ReplyData.Nickname == "" {
				return c.Status(400).JSON(
					serializer.Serialize(400, "Nickname can't be empty"),
				)
			}

			handler.HandleNewReply(
				replyPostData.ReplyData.CommentID,
				replyPostData.ReplyData.Nickname,
				replyPostData.ReplyData.Email,
				replyPostData.ReplyData.Content,
			)

			return c.Status(200).JSON(
				serializer.Serialize(200, "Post Succeed"),
			)
		}
		return c.Status(401).JSON(
			serializer.Serialize(401, "Captcha Verify Failed"),
		)
	}
}
