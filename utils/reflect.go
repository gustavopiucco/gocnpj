package utils

import (
	"fmt"
	"reflect"
	"strconv"
)

func ClearSlice(slice any) {
	sliceVal := reflect.ValueOf(slice).Elem()
	sliceVal.Set(reflect.MakeSlice(sliceVal.Type(), 0, 0))
}

func AppendLinesToSlice(lines [][]string, slice any) error {
	// Obter o valor e tipo refletido do slice
	sliceVal := reflect.ValueOf(slice).Elem()
	sliceType := sliceVal.Type()

	if sliceType.Kind() != reflect.Slice {
		return fmt.Errorf("provided value is not a slice")
	}

	sliceVal.Set(sliceVal.Slice(0, 0))

	// Verificar o tipo dos elementos do slice
	elemType := sliceType.Elem()

	// Iterar sobre cada linha e adicionar ao slice
	for _, line := range lines {
		// Criar uma nova instância do tipo dos elementos do slice
		elemVal := reflect.New(elemType).Elem()

		// Verificar se o número de campos corresponde ao número de colunas na linha
		if len(line) != elemType.NumField() {
			return fmt.Errorf("line length does not match number of fields in struct")
		}

		for i := 0; i < elemType.NumField(); i++ {
			field := elemVal.Field(i)
			if !field.CanSet() {
				return fmt.Errorf("field %s cannot be set", elemType.Field(i).Name)
			}

			value := line[i]
			switch field.Kind() {
			case reflect.String:
				field.SetString(value)
			case reflect.Int:
				intValue, err := strconv.Atoi(value)
				if err != nil {
					return fmt.Errorf("failed to convert %s to int: %v", value, err)
				}
				field.SetInt(int64(intValue))
			default:
				return fmt.Errorf("unsupported field type: %s", field.Type())
			}
		}

		// Adicionar o novo elemento ao slice
		sliceVal.Set(reflect.Append(sliceVal, elemVal))
	}

	return nil
}
