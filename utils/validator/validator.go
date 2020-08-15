package validator

import (
	"fmt"
	"ginblog/utils/errmsg"
	"github.com/go-playground/locales/zh_Hans_CN"
	unTrans "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTrans "github.com/go-playground/validator/v10/translations/zh"
	"reflect"
)

func Validate(data interface{}) (string, int) {
	validate := validator.New()
	uni := unTrans.New(zh_Hans_CN.New())
	trans, _ := uni.GetTranslator("zh_Hans_CN")

	// 中文提示
	err := zhTrans.RegisterDefaultTranslations(validate, trans)
	if err != nil {
		fmt.Println("err:", err)
	}
	// 映射db中的标签
	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		label := field.Tag.Get("label")
		return label
	})
	// 验证是结构体
	err = validate.Struct(data)
	if err != nil {
		for _, v := range err.(validator.ValidationErrors) {
			return v.Translate(trans), errmsg.ERROR
		}
	}
	return "", errmsg.SUCCSE
}
