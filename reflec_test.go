package main

import (
	"reflect"
	"testing"
)

/*
Golang反射三大法则
https://blog.golang.org/laws-of-reflection
* 从接口值(Interface Value)到反射对象(reflect.Value/reflect.Type)
* 从反射对象(reflect.Value/reflect.Type)到接口值(Interface Value)
* 要修改一个反射对象, 值必须是可设置的(可寻址)


*/

type Equip struct {
	Name   string
	Profit int
}

type Bag struct {
	ID    int64
	Award string
}

type Profile1 struct {
	HP int

	HPAddr *int

	Equip

	*Bag
}

func makeProfileInstance() *Profile1 {
	return &Profile1{HP: 100, HPAddr: nil, Bag: nil}
}

//================取类型信息===================
func TestReflectType(t *testing.T) {

	p := makeProfileInstance()

	typeOfP := reflect.TypeOf(p)

	// 指针类型, 必须取Elem()才能进行NumField操作

	// 类型本身
	t.Log("typeOfP", typeOfP)

	// 指针种类
	t.Log("typeOfP.Kind()", typeOfP.Kind())

	// 指针没有名字
	t.Log("typeOfP.Name()", typeOfP.Name())

	// 只有取Elem()才有名字
	t.Log("typeOfP.Elem().Name()", typeOfP.Elem().Name())

	// 指针只有取Elem()才能调用NumField(), 否则报错 NumField of non-struct type
	t.Log("typeOfP.Elem().NumField()", typeOfP.Elem().NumField())

}

//================取值信息===================
func TestReflectValue(t *testing.T) {

	p := makeProfileInstance()

	valueOfP := reflect.ValueOf(p)
	// 值本身
	t.Log("valueOfP", valueOfP)

	// Value的Type() 等效于原值的TypeOf()
	t.Log("valueOfP.Type()", valueOfP.Type())

	// 值的数据(原值的指针类型)
	t.Log("valueOfP.Interface()", valueOfP.Interface())

	// 值的元素类型
	t.Log("valueOfP.Elem().Type()", valueOfP.Elem().Type())

	// 值的元素数据(原值)
	t.Log("valueOfP.Elem().Interface()", valueOfP.Elem().Interface())

	// 指针只有取Elem才能调用NumField(), 否则报错 call of reflect.Value.NumField on ptr Value
	t.Log("valueOfP.Elem().NumField()", valueOfP.Elem().NumField())

}

//================普通类型的赋值===================
func TestReflectAssignFieldOfValue(t *testing.T) {

	p := makeProfileInstance()

	valueOfP := reflect.ValueOf(p)

	HPField := valueOfP.Elem().FieldByName("HP")

	t.Log("HPField", HPField)

	// 值类型
	t.Log("HPField.Type()", HPField.Type())

	// 注意:只有值所在的结构体是指针类型才可以被Set
	HPField.SetInt(1234)

	/* 以下代码会报错reflect.Value.SetInt using unaddressable value
	p2 := Profile1{100, nil}
	valueOfP2 := reflect.ValueOf(p2)
	valueOfP2.FieldByName("HP").SetInt(1234)
	*/

	t.Log("reflect set HPAddr", p.HP, HPField.Interface())

}

//================指针普通类型的赋值===================
func TestReflectAssignFieldOfPtrValue(t *testing.T) {

	p := makeProfileInstance()

	valueOfP := reflect.ValueOf(p)

	HPAddrField := valueOfP.Elem().FieldByName("HPAddr")

	// 值此时是nil
	t.Log("HPAddrField", HPAddrField)

	// 值类型
	t.Log("HPAddrField.Type()", HPAddrField.Type())

	// 取实际类型
	t.Log("HPAddrField.Type().Elem()", HPAddrField.Type().Elem())

	// reflect.New()的参数类型, 必须是非指针类型
	// 使用HPAddrField.Elem().Type()会报call of reflect.Value.Type on zero Value
	HPAddrValue := reflect.New(HPAddrField.Type().Elem())

	// reflect.New()的值是一个指针
	t.Log("HPAddrValue", HPAddrValue)

	// reflect.New()的是指针类型
	t.Log("HPAddrValue.Type()", HPAddrValue.Type())

	// reflect.New()的本体值
	t.Log("HPAddrValue.Interface()", HPAddrValue.Interface())

	// reflect.New()的返回值是一个指针类型, 因此取Elem()
	HPAddrValue.Elem().SetInt(1234)

	// 将指针值设置到原字段
	HPAddrField.Set(HPAddrValue)
	t.Log("reflect set HPAddr", *p.HPAddr, HPAddrField.Elem().Interface())

}

//================匿名结构体实例成员复制===================
func TestReflectAssignAnonymousStructField(t *testing.T) {

	p := makeProfileInstance()

	valueOfP := reflect.ValueOf(p)

	EquipFiled := valueOfP.Elem().FieldByName("Equip")

	// 值此时是nil
	t.Log("EquipFiled", EquipFiled)

	// 能否设置
	t.Log("EquipFiled.CanSet()", EquipFiled.CanSet())

	// 值类型
	t.Log("EquipFiled.Type()", EquipFiled.Type())

	// 取地址
	t.Log("EquipFiled.Addr()", EquipFiled.Addr())

	// 指针类型
	t.Log("EquipFiled.Addr().Type()", EquipFiled.Addr().Type())

	// 取字段地址
	EquipFiledOrigin := EquipFiled.Addr().Interface()

	t.Log("EquipFiled.Addr().Interface()", EquipFiledOrigin)

	valueOfEquipFiledOrigin := reflect.ValueOf(EquipFiledOrigin)

	// 指针本身不能被设置
	t.Log("valueOfEquipFiledOrigin.CanSet()", valueOfEquipFiledOrigin.CanSet())

	// 指针指向的元素可以被设置
	t.Log("valueOfEquipFiledOrigin.Elem().CanSet()", valueOfEquipFiledOrigin.Elem().CanSet())

	NameField := valueOfEquipFiledOrigin.Elem().FieldByName("Name")

	NameField.SetString("hello")

	/*
		// 以下代码因为valueOfEquipFiledOrigin2结构体为不可寻址(值拷贝的), 所以其字段也是无法设置的
		// 报错reflect.Value.SetString using unaddressable value

		valueOfEquipFiledOrigin2 := reflect.ValueOf(EquipFiled.Interface())

		valueOfEquipFiledOrigin2.FieldByName("Name").SetString("hello")

	*/

	t.Log("reflect set NameField", p.Equip.Name, NameField.Interface())

}
