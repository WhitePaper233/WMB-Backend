// package main backend-server entering
package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/mojocn/base64Captcha"
	"whitepaper233.top/WMBBackend/router"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())
	captchaStore := base64Captcha.DefaultMemStore

	// comment apis
	app.Get("/api/comment/latest", router.RouteQueryLatestComment)
	app.Get("/api/comment/count", router.RouteCommentCount)
	app.Get("/api/comment/query/page/:page", router.RouteQueryCommentByPage)
	app.Post("/api/comment/new", router.RouteNewComment(captchaStore))

	// reply apis
	app.Get("/api/reply/unfolded/:comment_id", router.RouteQueryUnfoldedReply)
	app.Get("/api/reply/count/:comment_id", router.RouteReplyCount)
	app.Get("/api/reply/query/:comment_id/page/:page", router.RouteQueryReplyByPage)
	app.Post("/api/reply/new", router.RouteNewReply(captchaStore))

	// captcha apis
	app.Get("/api/captcha/new", router.RouteGetCaptcha(captchaStore))
	// app.Post("/api/captcha/verify", router.RouteVerifyCaptcha(captchaStore))

	//  test apis
	// app.Get("/ping/:info", func(c *fiber.Ctx) error {
	// 	info := c.Params("info")
	// 	return c.SendString(info)
	// })

	app.Listen(":3000")
}
