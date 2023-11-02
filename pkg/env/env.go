package env

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
)

func Unmarshal(val interface{}) error {
	v := reflect.ValueOf(val).Elem()
	t := v.Type()

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := os.Getenv(field.Tag.Get("env"))
		f := v.FieldByName(field.Name)
		if f.Kind() == reflect.Struct {
			err := Unmarshal(f.Addr().Interface())
			if err != nil {
				return err
			}
			continue
		}

		if value != "" {
			switch f.Kind() {
			case reflect.String:
				f.SetString(value)
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				iv, err := strconv.ParseInt(value, 10, 64)
				if err != nil {
					return fmt.Errorf("env: failed to parse %s as int: %v", field.Name, err)
				}
				f.SetInt(iv)
			case reflect.Bool:
				bv, err := strconv.ParseBool(value)
				if err != nil {
					return fmt.Errorf("env: failed to parse %s as bool: %v", field.Name, err)
				}
				f.SetBool(bv)
			case reflect.Float32, reflect.Float64:
				fv, err := strconv.ParseFloat(value, 64)
				if err != nil {
					return fmt.Errorf("env: failed to parse %s as float: %v", field.Name, err)
				}
				f.SetFloat(fv)
			case reflect.Ptr:
				if f.Elem().Kind() == reflect.Struct {
					if err := Unmarshal(f.Interface()); err != nil {
						return err
					}
				}
			default:
				return fmt.Errorf("env: invalid type '%s'", f.Kind())
			}
		}
	}
	return nil
}
