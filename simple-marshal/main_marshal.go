package main

import (
	"fmt"
	"reflect"
	"strings"
)

// MarshalStruct превращает структуру в JSON-подобную структуру
// на входе принимает любое значение, а на выходе возвращает строку JSON
func MarshalStruct(v any) string {
	val := reflect.ValueOf(v) // получили значение
    typ := reflect.TypeOf(v)  // получили тип

    if val.Kind() != reflect.Struct { return "{}" } // проверяем, если пришла не структура, то просто возвращаем пустой обьект

	var sb strings.Builder
	sb.WriteString("{ ")

	// проходимся по структуре, по ее полям
	for i := 0; i < val.NumField(); i ++ {
	field := typ.Field(i)  // информация о поле
	value := val.Field(i)  // значение

	// читаем тег Json, если есть
	jsonTag := field.Tag.Get("json")

	if jsonTag =="-" { // если поле игнорируется "-", то пропускаем continue
		continue
	}

	if jsonTag =="" {  // если у тега нет имени, то используем имя поля
		jsonTag = field.Name
	}

	// пропускаем неэкспортируемые поля
	if field.PkgPath != "" {
		continue
	}

	// сериализуем простейшие типы
	var strValue string

	// switch проверяет что за поле нам пришло
	switch value.Kind() {
	case reflect.String: strValue = fmt.Sprintf("\"%s\"", value.String())   // если это строка
	case reflect.Int, reflect.Int64: strValue = fmt.Sprintf("%d", value.Int())  // если целочисленное число
	case reflect.Bool: strValue = fmt.Sprintf("%t", value.Bool())  // если bool
	default: strValue = "\"unsupported\""  // другой тип - поле не поддерживается
	}

	sb.WriteString(fmt.Sprintf("\"%s\": %s", jsonTag, strValue ))

	// проверяем, если это не последнее поле, добавляем запятую, так как в json все поля разделены ","
	if i < val.NumField()-1 {
		sb.WriteString(", ")
	}
}

	// закрываем обьект
	sb.WriteString(" }")
return sb.String()
}

type User struct {
	Name string `json:"name" validate:"min=1,max=32"`
	Age int `json:"age"`
	unExported string // не сериализуется
	NoTag string `json:"-"`
	Admin bool `json:"is_admin"`
}

func main() {
	u := User{Name:"Ваня", Age: 25, unExported: "secret", NoTag: "no tag"}
	fmt.Println(MarshalStruct(u))
}

// вызываем go run main_marshal.go