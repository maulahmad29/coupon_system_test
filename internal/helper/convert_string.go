package helper

import "strings"

func ConvertStringCouponName(s string) string {
	s = strings.TrimSpace(s)
	s = strings.ToUpper(s)
	s = strings.ReplaceAll(s, " ", "_")
	return s
}
