package utils

import (
	"fmt"
	"gin-myboot/global"
	"github.com/go-playground/validator/v10"
	"reflect"
	"strings"
)

func IsNil(i interface{})  bool {
	iv := reflect.ValueOf(i)

	if iv.Kind() == reflect.Ptr {
		return iv.IsNil()
	}
	return false
}

// 自定义错误消息

func GetError(err error, r interface{}) string {
	e := err.Error()
	if strings.HasPrefix(e, "json:")  {
		global.Error("GetError", e)
		return e
	}
	if strings.HasPrefix(e, "parsing")  {
		global.Error("GetError", e)
		return e
	}
	if e == "EOF" {
		global.Error("GetError", e)
		return "字段验证错误"
	}

	errs := err.(validator.ValidationErrors)
	s := reflect.TypeOf(r)
	for _, fieldError := range errs {
		filed, _ := s.FieldByName(fieldError.Field())
		errTag := fieldError.Tag() + "_err"
		// 获取对应binding得错误消息
		errTagText := filed.Tag.Get(errTag)
		// 获取统一错误消息
		errText := filed.Tag.Get("err")
		if errTagText != "" {
			return errTagText
		}
		if errText != "" {
			return errText
		}
		return fieldError.Field() + ":" + fieldError.Tag()
	}
	return ""
}

// 用b的所有字段覆盖a的
// 如果fields不为空, 表示用b的特定字段覆盖a的
// a应该为结构体指针

func CopyFields(a interface{}, b interface{}, fields ...string) (err error) {
	at := reflect.TypeOf(a)
	av := reflect.ValueOf(a)
	bt := reflect.TypeOf(b)
	bv := reflect.ValueOf(b)
	// 简单判断下
	if at.Kind() != reflect.Ptr {
		err = fmt.Errorf("a must be a struct pointer")
		return
	}
	av = reflect.ValueOf(av.Interface())
	// 要复制哪些字段
	_fields := make([]string, 0)
	if len(fields) > 0 {
		_fields = fields
	} else {
		for i := 0; i < bv.NumField(); i++ {
			_fields = append(_fields, bt.Field(i).Name)
		}
	}
	if len(_fields) == 0 {
		fmt.Println("no fields to copy")
		return
	}
	// 复制
	for i := 0; i < len(_fields); i++ {
		name := _fields[i]
		f := av.Elem().FieldByName(name)
		bValue := bv.FieldByName(name)
		// a中有同名的字段并且类型一致才复制
		if f.IsValid() && f.Kind() == bValue.Kind() {
			f.Set(bValue)
		} else {
			fmt.Printf("no such field or different kind, fieldName: %s\n", name)
		}
	}
	return
}


func JudgeIds(id uint64, ids []uint64) (b bool)  {
	flag := false

	for _, _id := range ids {
		if _id == id {
			flag = true
			break
		}
	}

	return flag
}