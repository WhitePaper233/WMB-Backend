package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mojocn/base64Captcha"
	"whitepaper233.top/WMBBackend/handler"
	"whitepaper233.top/WMBBackend/serializer"
)

type CaptchaData struct {
	CaptchaID   string `json:"id"`
	CaptchaBS64 string `json:"bs64"`
}

type VerifyData struct {
	CaptchaID     string `json:"id"`
	CaptchaAnswer string `json:"answer"`
}

func RouteGetCaptcha(store base64Captcha.Store) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		captcahID, captchaBS64, err := handler.HandleCaptchaGenerate(store)
		if err != nil {
			return c.Status(500).JSON(serializer.Serialize(500, err.Error()))
		}

		captcha := new(CaptchaData)
		captcha.CaptchaID = captcahID
		captcha.CaptchaBS64 = captchaBS64

		return c.Status(200).JSON(serializer.SerializeWithData(200, "", captcha))
	}
}

func RouteVerifyCaptcha(store base64Captcha.Store) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		verifyData := new(VerifyData)
		if err := c.BodyParser(verifyData); err != nil {
			return c.Status(503).JSON(serializer.Serialize(503, err.Error()))
		}

		if handler.HandleCaptchaVerify(store, verifyData.CaptchaID, verifyData.CaptchaAnswer) {
			return c.Status(200).JSON(serializer.Serialize(200, "Verify Succeed"))
		}
		return c.Status(401).JSON(serializer.Serialize(401, "Verify Failed"))
	}
}
