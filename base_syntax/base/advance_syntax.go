package main

import (
	"fmt"
	"reflect"
)

type DeclarationInterface interface {
	NoParamReturnVoid()
	NoParamReturnValue() string
	//ParamReturnVoid(param *string)
	//ParamReturnValue(param *string,a ...[]interface{}) string
}

func reflect1() {
	type DocDetail struct {
		Catalogue bool   `gorm:"type:varchar(35)" json:"catalogue" form:"catalogue"`
		Title     string `gorm:"type:varchar(35)" json:"title" form:"title"`
	}
	d := &DocDetail{}
	valueD := &DocDetail{
		Catalogue: true,
		Title:     "fdsafsd",
	}
	dType := reflect.TypeOf(*d)
	for j := 0; j < 1; j++ {
		_ = dType.Field(j)
	}
	fmt.Printf("name:%v , kind: %v", dType.Name(), dType.Kind())
	dValue := reflect.ValueOf(d).Elem()
	fmt.Printf("kind: %v", dType.Kind())
	fmt.Println(d)
	for i := 0; i < dValue.NumField(); i++ {
		typeField := dType.Field(i)
		valueField := dValue.Field(i)
		fmt.Println("tag gorm:", typeField.Tag.Get("gorm"))
		fmt.Println("tag json:", typeField.Tag.Get("json"))
		fmt.Println(typeField.Name, valueField.CanSet())
	}
	reflect.ValueOf(d).Elem().Set(reflect.ValueOf(*valueD))
	fmt.Println(d)
}

func reflect2() {
	type test struct {
		field1 string
		field2 int
	}
	//prtValue := &test{}
	//t := reflect.ValueOf(prtValue).Elem().Type()
	//fmt.Printf("pointer value: isPtr: %v,",reflect.ValueOf(prtValue).Kind())
	//fmt.Printf("name: %v,fieldCount: %v,kind: %v", t.Name(), t.NumField(),t.Kind())
	//fmt.Println()
	normalValue := &test{}
	fmt.Printf("pointer value: isPtr: %v,", reflect.ValueOf(normalValue).Kind())
	typeValue := reflect.ValueOf(normalValue).Type()
	//当一个interface{}是一个指针类型，可以通过reflect.ValueOf(interface{})获取反射对象后，调用kind()函数获取该值的类型
	//如果kind类型不是struct类型，那么部分反射函数调用时将会panic
	//如果kind类型是slice或者ptr类型，可以再通过Elem()函数获取相关的struct类型的反射类型数据，之后便可以获取struct中的所有field、method等数据了
	fmt.Printf("name: %v,fieldCount: %v,kind: %v", typeValue.Name(), typeValue.NumField(), typeValue.Kind())
}
