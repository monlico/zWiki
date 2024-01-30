package util

import "reflect"

// ConvertToDTO 将模型结构体转换为DTO结构体
func ConvertToDTO(model interface{}, dto interface{}) error {
	modelValue := reflect.ValueOf(model)
	dtoValue := reflect.ValueOf(dto).Elem()

	for i := 0; i < modelValue.NumField(); i++ {
		field := modelValue.Type().Field(i)
		dtoField, found := dtoValue.Type().FieldByName(field.Name)
		if found {
			dtoValue.FieldByName(dtoField.Name).Set(modelValue.Field(i))
		}
	}

	return nil
}

//去重一个切片
func UniqueSlice(sli []interface{}) []interface{} {
	var (
		resMap = make(map[interface{}]struct{}, len(sli))
		res    []interface{}
	)

	for _, v := range sli {
		if _, ok := resMap[v]; !ok {
			res = append(res, v)
			resMap[v] = struct{}{}
		}
	}
	return res
}

//去重一个int切片
func UniqueIntSlice(sli []int) []int {
	var (
		resMap = make(map[int]struct{}, len(sli))
		res    []int
	)

	for _, v := range sli {
		if _, ok := resMap[v]; !ok {
			res = append(res, v)
			resMap[v] = struct{}{}
		}
	}
	return res
}
