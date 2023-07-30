package initialize

import (
	"base-system-backend/global"
	"base-system-backend/initialize/internal"
	"fmt"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
	"reflect"
	"strings"
)

func InitTrans(locale string) (err error) {
	// 修改gin框架中的Validator引擎属性，实现自定制
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {

		// 注册一个获取json tag的自定义方法
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			label := fld.Tag.Get("label")
			if name == "_" {
				return ""
			}
			if label != "" {
				return name + "|" + fld.Tag.Get("label")

			} else {
				return name

			}
		})

		zhT := zh.New() // 中文翻译器
		enT := en.New() // 英文翻译器

		// 第一个参数是备用（fallback）的语言环境
		// 后面的参数是应该支持的语言环境（支持多个）
		// uni := ut.New(zhT, zhT) 也是可以的
		uni := ut.New(enT, zhT, enT)

		// locale 通常取决于 http 请求头的 'Accept-Language'
		var ok bool
		// 也可以使用 uni.FindTranslator(...) 传入多个locale进行查找
		global.TRANS, ok = uni.GetTranslator(locale)
		if !ok {
			panic(fmt.Errorf("uni.GetTranslator(%s) failed", locale))
		}

		// 注册翻译器
		switch locale {
		case "en":
			err = enTranslations.RegisterDefaultTranslations(v, global.TRANS)
		case "zh":
			err = zhTranslations.RegisterDefaultTranslations(v, global.TRANS)
		default:
			err = enTranslations.RegisterDefaultTranslations(v, global.TRANS)
		}

		//在校验器注册自定义的校验方法
		if err = v.RegisterValidation("enum", internal.ValidateEnum); err != nil {
			return
		}

		//注意！因为这里会使用到trans实例
		//所以这一步注册要放到trans初始化的后面
		if err = v.RegisterTranslation(
			"enum",
			global.TRANS,
			registerTranslator("enum", "{0}不合法"),
			translate,
		); err != nil {
			return
		}
		return
	}
	return
}

// registerTranslator为自定义字段添加翻译功能
func registerTranslator(tag string, msg string) validator.RegisterTranslationsFunc {
	return func(trans ut.Translator) error {
		if err := trans.Add(tag, msg, false); err != nil {
			return err
		}
		return nil
	}
}

// translate 自定义字段的翻译方法
func translate(trans ut.Translator, fe validator.FieldError) string {
	msg, err := trans.T(fe.Tag(), fe.Field())
	if err != nil {
		panic(fmt.Errorf("translator failed: %w", fe.(error)))
	}
	return msg
}
