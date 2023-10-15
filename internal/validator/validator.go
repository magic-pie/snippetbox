package validator

import (
	"regexp"
	"strings"
	"unicode/utf8"
)

var EmailRX = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

type Validator struct {
	NonFieldErrors []string
	FieldErrors    map[string]string
}

func (v *Validator) Valid() bool {
	return len(v.FieldErrors) == 0 && len(v.NonFieldErrors) == 0
}

func (v *Validator) AddFieldError(key, message string) {
	if v.FieldErrors == nil {
		v.FieldErrors = make(map[string]string)
	}

	if _, exists := v.FieldErrors[key]; !exists {
		v.FieldErrors[key] = message
	}
}

func (v *Validator) AddNonFieldError(message string) {
	v.NonFieldErrors = append(v.NonFieldErrors, message)
}

func (v *Validator) CheckField(ok bool, key, message string) {
	if !ok {
		v.AddFieldError(key, message)
	}
}

func NotBlank(s string) bool {
	return strings.TrimSpace(s) != ""
}

func MinChars(s string, n int) bool {
	return utf8.RuneCountInString(s) >= n
}

func MaxChars(s string, n int) bool {
	return utf8.RuneCountInString(s) <= n
}

func Matches(s string, re *regexp.Regexp) bool {
	return re.MatchString(s)
}

func PermittedInt(v int, permittedValues ...int) bool {
	for i := range permittedValues {
		if v == permittedValues[i] {
			return true
		}
	}

	return false
}
