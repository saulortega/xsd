package xsd

import (
	"errors"
	"regexp"
	"strings"
	"text/template"
	"time"
)

func normalizedString(s string) string {
	s = strings.Replace(s, "\r\n", " ", -1)
	s = strings.Replace(s, "\n", " ", -1)
	s = strings.Replace(s, "\t", " ", -1)
	return s
}

func token(s string) string {
	s = normalizedString(s)
	s = regexp.MustCompile(`[\s\p{Zs}]{2,}`).ReplaceAllString(s, " ")
	s = strings.TrimSpace(s)
	return s
}

// http://www.datypic.com/sc/xsd/t-xsd_string.html
func String(s string) string {
	return template.HTMLEscapeString(s)
}

// http://www.datypic.com/sc/xsd/t-xsd_normalizedString.html
func NormalizedString(s string) string {
	s = normalizedString(s)
	s = String(s)
	return s
}

// http://www.datypic.com/sc/xsd/t-xsd_token.html
func Token(s string) string {
	s = token(s)
	s = String(s)
	return s
}

// http://www.datypic.com/sc/xsd/t-xsd_language.html
func Language(s string) (string, error) {
	s = token(s)

	enc, err := regexp.MatchString(`[a-zA-Z]{1,8}(-[a-zA-Z0-9]{1,8})*`, s)
	if err != nil || !enc {
		return "", errors.New("wrong format")
	}

	return s, nil
}

// http://www.datypic.com/sc/xsd/t-xsd_Name.html
func Name(s string) (string, error) {
	s = token(s)

	enc, err := regexp.MatchString(`^[a-zA-ZñÑáéíóúÁÉÍÓÚ_:]`, s)
	if err != nil || !enc {
		return "", errors.New("wrong format")
	}

	enc, err = regexp.MatchString(`[^a-zA-ZñÑáéíóúÁÉÍÓÚ0-9_:\-\.]`, s)
	if err != nil || enc {
		return "", errors.New("wrong format")
	}

	return s, nil
}

// http://www.datypic.com/sc/xsd/t-xsd_time.html
func Time(t time.Time) string {
	return t.Format("15:04:05")
}

// http://www.datypic.com/sc/xsd/t-xsd_date.html
func Date(t time.Time) string {
	return t.Format("2006-01-02")
}
