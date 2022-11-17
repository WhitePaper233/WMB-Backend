package handler

import (
	"github.com/mojocn/base64Captcha"
	"whitepaper233.top/WMBBackend/service"
)

func HandleCaptchaGenerate(store base64Captcha.Store) (string, string, error) {
	id, bs64, err := service.GenerateCaptcha(store)
	if err != nil {
		return "", "", err
	}

	return id, bs64, nil
}

func HandleCaptchaVerify(store base64Captcha.Store, id string, answer string) bool {
	return service.CaptchaVerify(store, id, answer)
}
