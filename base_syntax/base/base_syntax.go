package main

import (
	fmtImportAlias "fmt"
)

func main() {
	//fmtImportAlias.Println("Hello world!")
	//declarateVar()
	//forloop()
	////var parentCursor *ImportStruct
	//var subCursor *StructDeclaration
	//subCursor, _ = createStructCursor()
	//fmtImportAlias.Println(subCursor.ToString())
	//fmtImportAlias.Println(subCursor.ImportStruct.ToString())
	//conditionProcess()
	//
	//var inter DeclarationInterface = subCursor
	//fmtImportAlias.Print("inter:NoParamReturnVoid():")
	//inter.NoParamReturnVoid()
	//fmtImportAlias.Println("inter:NoParamReturnValue():", inter.NoParamReturnValue())
	reflect1()
	//reflect2()
}

func declarateVar() {
	i := "test"
	fmtImportAlias.Println(i)
	var testInt int = 321
	fmtImportAlias.Println(testInt)
	var cursorStr *string
	test := "rewrewrew"
	cursorStr = &test
	fmtImportAlias.Println(cursorStr)
	mux, muxStr := 1, "fdsfds"
	fmtImportAlias.Println(mux, muxStr)

}

func conditionProcess() {
	var something *string
	//bool //默认值为false
	//string //默认值为空字符串
	//int int8 int16 int32 int64 //默认值为0
	//uint uint8 uint16 uint32 uint64 uintptr //默认值为0
	//byte // uint8 的别名
	//rune // int32 的别名
	//float32 float64 //默认值为0
	//complex64 complex128 //默认值为0
	//pointers -> nil
	//slices -> nil
	//maps -> nil
	//channels -> nil
	//functions -> nil
	//interfaces -> nil
	fmtImportAlias.Println("enter func conditionProcess")
	if nil != something {
		fmtImportAlias.Println(something)
	}
	valueStr := "tesatesa"
	something = &valueStr
	if nil == something {
		fmtImportAlias.Println(something)
	}
}

func forloop() {
	for i := 0; i < 5; i++ {
		fmtImportAlias.Println("fori:", "", i)
	}
	rangeOfInt := [3]string{"str1", "str2", "str3"}
	for index, element := range rangeOfInt {
		fmtImportAlias.Println("forr: ", index, ":", element)
	}
}

func createStructCursor() (*StructDeclaration, *ImportStruct) {
	var first *StructDeclaration
	var parent *ImportStruct

	first = &StructDeclaration{
		name:  "my name you cant imagine",
		value: "surprise mother fucker",
	}
	parent = &ImportStruct{
		name: "I'm your father",
	}

	first.ImportStruct = parent
	return first, parent
}

func (sd *StructDeclaration) ToString() string {
	return fmtImportAlias.Sprintf("[name: %s,value: %s,import: %s]", sd.name, sd.value, sd.ImportStruct.ToString())
}

func (ps *ImportStruct) ToString() string {
	return fmtImportAlias.Sprintf("[name: %s]", ps.name)
}

func (sd *ImportStruct) NoParamReturnVoid() {
	fmtImportAlias.Println("there is a interface's method implement by sub!!")
}

func (sd *ImportStruct) NoParamReturnValue() string {
	return "there is a interface's method implement by sub!!"
}

/**
结构体的继承只能通过引入对应结构体变量或者匿名声明
*/
type StructDeclaration struct {
	name  string
	value string
	*ImportStruct
}
type ImportStruct struct {
	name string
}
