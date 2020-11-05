package validator

import (
	"fmt"
	"github.com/go-playground/locales/zh_Hans_CN"
	unTrans "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTrans "github.com/go-playground/validator/v10/translations/zh"
	"myblog/utils/errmsg"
	"reflect"
)

func Validate(data interface{}) (string, int) {
	validate := validator.New() // 实例化

	uni := unTrans.New(zh_Hans_CN.New()) // 获取翻译

	trans, _ := uni.GetTranslator("zh_Hans_CN") // 获取中文翻译

	err := zhTrans.RegisterDefaultTranslations(validate, trans) // 注册验证 和 翻译
	if err != nil {
		fmt.Println("err: ", err)
	}

	// 将tag label中的字段映射到字段名上
	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		label := field.Tag.Get("label")
		return label
	})

	err = validate.Struct(data)
	if err != nil {
		for _, v := range err.(validator.ValidationErrors) {
			return v.Translate(trans), errmsg.ERROR
		}
	}

	return "", errmsg.SUCCESS
}
