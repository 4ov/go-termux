package termux

import (
	"fmt"
	"os/exec"
	"reflect"
)

func ReadyArgs(options interface{}) []string {
	result := []string{}
	keys := reflect.TypeOf(options)
	values := reflect.ValueOf(options)
	key_len := keys.NumField()
	for i := 0; i < key_len; i++ {
		key := keys.Field(i)
		value := values.Field(i)
		result = append(result, parseValue(value, key)...)
		// fmt.Print("Type:", key.Type, ",", key.Name, "=", value, "\n")
	}
	return result
}

func CallCommand(command string, args ...string) ([]byte, error) {
	cmd := exec.Command(command, args...)
	return cmd.Output()

}

func parseValue(value reflect.Value, key reflect.StructField) []string {
	result := []string{}
	switch value.Kind() {
	case reflect.String:
		if value.String() != "" {
			result = append(result, key.Tag.Get("arg"), fmt.Sprintf(`%s`, value))
		}
	case reflect.Bool:
		if value.Bool() {
			result = append(result, key.Tag.Get("arg"))
		}
	case reflect.Int:
		if value.Int() != 0 {
			result = append(result, fmt.Sprintf("%s %d", key.Tag.Get("arg"), value.Int()))
		}
	case reflect.Slice:
		result = append(result, key.Tag.Get("arg"))
		s := value.Slice(0, value.Len())
		result = append(result, joinSlice(s, key.Tag.Get("split")))

	//# SORRY, JUST TO PROVE THING
	// case reflect.Interface:
	// 	if value.IsNil() {
	// 		break
	// 	}
	// 	v := value.Interface()
	// 	j, _ := json.Marshal(v)
	// 	result = append(result, key.Tag.Get("arg"), string(j))

	default:
		panic("Unsupported type " + value.Kind().String())
	}
	return result
}

func joinSlice(slice reflect.Value, sep string) string {
	//type of elements
	elem_type := slice.Type().Elem()

	result := ""
	for i := 0; i < slice.Len(); i++ {
		var v string
		switch elem_type.Kind() {
		case reflect.String:
			v = slice.Index(i).String()
		case reflect.Int:
			v = fmt.Sprintf("%d", slice.Index(i).Int())
		case reflect.Bool:
			v = fmt.Sprintf("%t", slice.Index(i).Bool())
		default:
			panic("Unsupported type " + elem_type.Kind().String())
		}

		// v = strings.Trim(v, " ")

		if i == slice.Len()-1 {
			result += fmt.Sprint(v)
		} else {
			result += fmt.Sprintf("%s%s", v, sep)
		}
	}
	return result
}
