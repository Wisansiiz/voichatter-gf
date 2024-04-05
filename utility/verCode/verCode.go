package verCode

import (
	"github.com/mojocn/base64Captcha"
)

var Store = base64Captcha.DefaultMemStore

func CaptchaVerify(id string, value string) bool {
	return Store.Verify(id, value, true)
}
