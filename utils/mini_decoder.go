package utils

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"time"
)

//用map填充结构
func MiniDecodeStruct(data map[string]interface{}, obj interface{}) error {
	structVal := reflect.ValueOf(obj).Elem()
	structType := structVal.Type()
	for i := 0; i < structType.NumField(); i++ {
		fieldType := structType.Field(i)
		field := structVal.Field(i)

		if !field.IsValid() {
			panic("field is not valid")
		}

		if !field.CanSet() {
			continue
		}

		fieldName := fieldType.Name
		tagValue := fieldType.Tag.Get("json")

		fmt.Println("fieldname:", fieldName, " --> tag: ", tagValue)
	}
	return nil
}

//用map的值替换结构的值
func decodeValue(obj interface{}, name string, value interface{}) error {
	//	structType := reflect.TypeOf(obj).Elem()
	//	structTagValue, ok := structType.FieldByName(name)
	//	if !ok {
	//		return nil
	//	}
	//	tagname := structTagValue.Tag.Get(name)

	structValue := reflect.ValueOf(obj).Elem()        //结构体属性值
	structFieldValue := structValue.FieldByName(name) //结构体单个属性值

	if !structFieldValue.IsValid() {
		return fmt.Errorf("No such field: %s in obj", name)
	}

	if !structFieldValue.CanSet() {
		return fmt.Errorf("Cannot set %s field value", name)
	}

	structFieldType := structFieldValue.Type() //结构体的类型
	val := reflect.ValueOf(value)              //map值的反射值

	var err error
	if structFieldType != val.Type() {
		val, err = typeconver(fmt.Sprintf("%v", value), structFieldValue.Type().Name()) //类型转换
		if err != nil {
			return err
		}
	}

	structFieldValue.Set(val)
	return nil
}

//类型转换
func typeconver(value string, ntype string) (reflect.Value, error) {
	if ntype == "string" {
		return reflect.ValueOf(value), nil
	} else if ntype == "time.Time" {
		t, err := time.ParseInLocation("2006-01-02 15:04:05", value, time.Local)
		return reflect.ValueOf(t), err
	} else if ntype == "Time" {
		t, err := time.ParseInLocation("2006-01-02 15:04:05", value, time.Local)
		return reflect.ValueOf(t), err
	} else if ntype == "int" {
		i, err := strconv.Atoi(value)
		return reflect.ValueOf(i), err
	} else if ntype == "int8" {
		i, err := strconv.ParseInt(value, 10, 64)
		return reflect.ValueOf(int8(i)), err
	} else if ntype == "int32" {
		i, err := strconv.ParseInt(value, 10, 64)
		return reflect.ValueOf(int64(i)), err
	} else if ntype == "int64" {
		i, err := strconv.ParseInt(value, 10, 64)
		return reflect.ValueOf(i), err
	} else if ntype == "float32" {
		i, err := strconv.ParseFloat(value, 64)
		return reflect.ValueOf(float32(i)), err
	} else if ntype == "float64" {
		i, err := strconv.ParseFloat(value, 64)
		return reflect.ValueOf(i), err
	}

	//else if .......增加其他一些类型的转换

	return reflect.ValueOf(value), errors.New("未知的类型：" + ntype)
}
