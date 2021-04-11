package common

import (
	"fmt"
	"github.com/gin-gonic/gin"
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

var trans ut.Translator

func Trans(locale string) (err error) {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterTagNameFunc(func(jsonField reflect.StructField) string {
			jsonName := strings.SplitN(jsonField.Tag.Get("json"), ",", 2)[0]
			return jsonName
		})
		zhTrans := zh.New()
		enTrans := en.New()
		ut := ut.New(enTrans, zhTrans, enTrans)
		var ok bool
		trans, ok = ut.GetTranslator(locale)
		if !ok {
			return fmt.Errorf("ut.GetTranslator(%s) failed", locale)
		}
		switch locale {
		case "en":
			err = enTranslations.RegisterDefaultTranslations(v, trans)
		case "zh":
			err = zhTranslations.RegisterDefaultTranslations(v, trans)
		default:
			err = enTranslations.RegisterDefaultTranslations(v, trans)
		}
		return
	}
	return
}

func resetFields(fields map[string]string) map[string]string {
	m := map[string]string{}
	for field, err := range fields {
		m[field[strings.Index(field, ".")+1:]] = err
	}
	return m
}

func ShouldAndValid(c *gin.Context, obj interface{}) (errs interface{}) {
	if err := Trans("zh"); err != nil {
		return err
	}
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	err := c.ShouldBind(obj)
	if err != nil {
		validErr, ok := err.(validator.ValidationErrors)
		if !ok {
			errs := err.Error()
			return errs
		}
		errs := resetFields(validErr.Translate(trans))
		return errs
	}
	return nil
}
