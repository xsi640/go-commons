package regex

import (
	"go-commons/strings"
	"regexp"
	"unicode"
)

var (
	HtmlColorPattern = regexp.MustCompile("^#?([a-f]|[A-F]|[0-9]){3}(([a-f]|[A-F]|[0-9]){3})?$")
	NumberPattern    = regexp.MustCompile("^[0-9]+$")
	EnglishPattern   = regexp.MustCompile("^[A-Za-z0-9]+$")
	EmailPattern     = regexp.MustCompile("^[\\w-]+(\\.[\\w-]+)*@[\\w-]+(\\.[\\w-]+)+$")
	IPV4Pattern      = regexp.MustCompile("^\\d{1,3}\\.\\d{1,3}\\.\\d{1,3}\\.\\d{1,3}$")
	UrlPattern       = regexp.MustCompile("^(http|https|ftp|rtsp|mms):(\\/\\/|\\\\\\\\)[A-Za-z0-9%\\-_@]+\\.[A-Za-z0-9%\\-_@]+[A-Za-z0-9\\.\\/=\\?%\\-&_~`@:\\+!;]*$")
	ChinesePattern   = regexp.MustCompile("^[\u4e00-\u9fa5]+$")
	DomainPattern    = regexp.MustCompile("([0-9A-Za-z]{2,}\\.[0-9A-Za-z]{2,3}|[0-9A-Za-z]{2,}\\.[0-9A-Za-z]{2,3})$")
)

func MatchHtmlColor(s string) bool {
	return HtmlColorPattern.MatchString(s)
}

func MatchNumber(s string) bool {
	return NumberPattern.MatchString(s)
}

func MatchEnglish(s string) bool {
	return EnglishPattern.MatchString(s)
}

func MatchEmail(s string) bool {
	return EmailPattern.MatchString(s)
}

func MatchIPv4(s string) bool {
	return IPV4Pattern.MatchString(s)
}

func MatchUrl(s string) bool {
	return UrlPattern.MatchString(s)
}

func MatchChinese(s string) bool {
	return ChinesePattern.MatchString(s)
}

func ContainerChinese(s string) bool {
	if strings.IsNotEmpty(s) {
		for _, r := range s {
			if unicode.Is(unicode.Han, r) && MatchChinese(string(r)) {
				return true
			}
		}
	}
	return false
}

func MatchDomain(s string) bool {
	return DomainPattern.MatchString(s)
}
