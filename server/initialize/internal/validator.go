package internal

import (
	"base-system-backend/global"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

type Enum interface {
	IsValid() bool
}

func ValidateEnum(fl validator.FieldLevel) bool {
	value := fl.Field().Interface().(Enum)
	return value.IsValid()
}

// RegisterEnum 自定义枚举值校验
func RegisterEnum(v *validator.Validate) {
	_ = v.RegisterValidation("enum", ValidateEnum)

	_ = v.RegisterTranslation("enum", global.TRANS, func(ut ut.Translator) error {
		return ut.Add("enum", "{0}不合法", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("enum", fe.Field())
		return t
	})
}
