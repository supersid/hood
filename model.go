package hood

import (
	"errors"
	"reflect"
)

func modelFieldOrTableName(i interface{}) string {
	return snakeCase(reflect.TypeOf(i).Elem().Name())
}

func modelMap(model interface{}) (*Model, error) {
	v := reflect.Indirect(reflect.ValueOf(model))
	if v.Kind() != reflect.Struct {
		return nil, errors.New("model is not a struct")
	}
	t := v.Type()
	m := &Model{
		Pk:     nil,
		Table:  modelFieldOrTableName(model),
		Fields: make(map[string]interface{}),
	}
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Tag == "PK" {
			m.Pk = &Pk{
				Name: snakeCase(string(field.Name)),
				Type: field.Type,
			}
		}
		key := snakeCase(field.Name)
		m.Fields[key] = v.FieldByName(field.Name).Interface()
	}
	return m, nil
}
