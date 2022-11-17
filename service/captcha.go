package service

import (
	"github.com/mojocn/base64Captcha"
)

func GenerateCaptcha(store base64Captcha.Store) (string, string, error) {
	var driver base64Captcha.Driver
	var driverString base64Captcha.DriverString

	captchaConfig := base64Captcha.DriverString {
		Height: 60,
		Width: 150,
		NoiseCount: 0,
		ShowLineOptions: 2 | 4,
		Length: 4,
		Source: "123456789qwertyuiplkjhgfdsazxcvbnm",
		Fonts: []string{"wqy-microhei.ttc"},
	}

	driverString = captchaConfig
	driver = driverString.ConvertFonts()
	captcha := base64Captcha.NewCaptcha(driver, store)
	id, b64s, err := captcha.Generate()

	return id, b64s, err
}

func CaptchaVerify(store base64Captcha.Store, id string, answer string) bool {
	if store.Verify(id, answer, false) {
		return true
	} else {
		return false
	}
}