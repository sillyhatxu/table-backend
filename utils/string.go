package utils

import "strings"

func FormatString(s string) string {
	return strings.ReplaceAll(s, " ", "")
}

func FormatContext(s string) string {
	s = strings.ReplaceAll(s, "：", ":")
	s = strings.ReplaceAll(s, " ", "")
	s = strings.ReplaceAll(s, "	", "")
	s = strings.ReplaceAll(s, "手机号码", "手机号")
	s = strings.ReplaceAll(s, "邮箱号码", "邮箱")
	s = strings.ReplaceAll(s, "邮箱号", "邮箱")
	s = strings.ReplaceAll(s, "身份证号", "身份证")
	return s
}
