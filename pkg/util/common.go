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
