package validator

import (
	"fmt"
	"regexp"
	"strings"
	"unicode/utf8"
)

type Validator struct {
	Errors []string
}

func New() *Validator {
	return &Validator{
		Errors: []string{},
	}
}

func (v *Validator) CheckBlank(key, value string) {
	if strings.TrimSpace(value) == "" {
		message := fmt.Sprintf("%s: não pode ser vazio", key)
		v.Errors = append(v.Errors, message)
	}
}

func (v *Validator) CheckMaxLength(key, value string, maxLength int) {
	if utf8.RuneCountInString(value) > maxLength {
		message := fmt.Sprintf("%s: pode conter no máximo %d caracteres", key, maxLength)
		v.Errors = append(v.Errors, message)
	}
}

func (v *Validator) CheckEmail(key, value string) {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	if !emailRegex.MatchString(value) {
		message := fmt.Sprintf("%s: inválido", key)
		v.Errors = append(v.Errors, message)
	}
}

func (v *Validator) HasErrors() bool {
	return len(v.Errors) > 0
}
