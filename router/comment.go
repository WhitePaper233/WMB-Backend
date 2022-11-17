// package router routers
package router

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/mojocn/base64Captcha"

	"whitepaper233.top/WMBBackend/handler"
	"whitepaper233.top/WMBBackend/serializer"
)

type CommentQueryData struct {
	Page uint64 `json:"page"`
}

type CommentPostData struct {
	CaptchaVerifyData struct {
		VerifyData
	} `json:"captcha_verify_data"`
	CommentData struct {
		// Avatar    int64  `json:"avatar"`
		Nickname string `json:"nickname"`
		Email    string `json:"email"`
		Content  string `json:"content"`
	} `json:"comment_data"`
}

func RouteQueryLatestComment(c *fiber.Ctx) error {
	return c.Status(200).JSON(
		handler.HandleQueryLatestComment(),
	)
}

func RouteQueryCommentByPage(c *fiber.Ctx) error {
	page, err := strconv.ParseUint(c.Params("page"), 10, 64)
	if err != nil {
		return c.Status(500).JSON(
			serializer.Serialize(500, err.Error()),
		)
	}

	return c.Status(200).JSON(
		handler.HandleQueryCommentByPage(page),
	)
}

func RouteCommentCount(c *fiber.Ctx) error {
	return c.Status(200).JSON(
		handler.HandleCountComment(),
	)
}

func RouteNewComment(store base64Captcha.Store) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		commentPostData := new(CommentPostData)

		if err := c.BodyParser(commentPostData); err != nil {
			return c.Status(500).JSON(
				serializer.Serialize(500, err.Error()),
			)
		}

		if handler.HandleCaptchaVerify(
			store,
			commentPostData.CaptchaVerifyData.CaptchaID,
			commentPostData.CaptchaVerifyData.CaptchaAnswer,
		) {

			if commentPostData.CommentData.Content == "" {
				return c.Status(400).JSON(
					serializer.Serialize(400, "Content can't be empty"),
				)
			}
			if commentPostData.CommentData.Nickname == "" {
				return c.Status(400).JSON(
					serializer.Serialize(400, "Nickname can't be empty"),
				)
			}

			handler.HandleNewComment(
				commentPostData.CommentData.Nickname,
				commentPostData.CommentData.Email,
				commentPostData.CommentData.Content,
			)

			return c.Status(200).JSON(
				serializer.SerializeWithData(200, "Post Succeed", commentPostData),
			)
		}
		return c.Status(401).JSON(
			serializer.Serialize(401, "Captcha Verify Failed"),
		)
	}
}
