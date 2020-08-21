package tools

import "regexp"

// phone verify
func VerifyPhone(num string) bool {
	reg := `^1([3589]\d|4[5-9]|6[124567]|7[0-8])\d{8}$`
	rgx := regexp.MustCompile(reg)
	return rgx.MatchString(num)
}

// email verify
func VerifyEmail(email string) bool {
	pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*` // 匹配电子邮箱
	// pattern := `^[A-Za-z0-9\u4e00-\u9fa5]+@[a-zA-Z0-9_-]+(\.[a-zA-Z0-9_-]+)+$`
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}
