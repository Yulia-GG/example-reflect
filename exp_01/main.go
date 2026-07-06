package main

import (
	"fmt"
	"reflect"
)

type InspectStruct struct {
     Age    int
	 name   string
}

func inspectStruct(i any) error {
   val := reflect.ValueOf(i) // получаем значение
   typ := reflect.TypeOf(i)  // получаем тип

// у значения проверяем метод Kind(), передали ли нам указатель
// если передали, то получаем методом Elem() возможность изменять это поле
if val.Kind() == reflect.Ptr {
	val = val.Elem()
	typ = typ.Elem()
}

// проверяем методом Kind(), передали ли рефлекс в структуру,
// если нет, то нам инспектировать ее не надо и делаем ранний выход
if val.Kind() != reflect.Struct {
	fmt.Println("Not a struct")
	return nil
}

// циклом обходим по количеству полей, которые у нас есть
fmt.Printf("Inspecting struct: %s\n", typ.Name())
for idx := 0; idx < val.NumField(); idx ++ {
	field := typ.Field(idx)  // информация о поле
	value := val.Field(idx)  // значение поля

// выводим некую информацию
	fmt.Printf("Field: %s\n", field.Name)
	fmt.Printf("Type: %s\n", field.Type)

// PkgPath это сво-во для неэкспортируемых полей (св-во не пустое)
// для экспортируемых полей сво-во будет пустым
	if field.PkgPath != "" {
		fmt.Printf("PkgPath: %q (неэкспортируемое поле)\n", field.PkgPath)
	} else {
		fmt.Printf("PkgPath: <empty> (экспортируемое поле)\n")
	}

	// проверяется на доступность значения, если может, выводится
	// если неэкспортируемое поле, то вывести не может
	if  value.CanInterface() {
		fmt.Printf("Value: %v\n", value.Interface())
	} else {
		fmt.Printf("Value: <unexported, cannot access>\n")
	}

	// возможность изменения
	fmt.Printf("CanSet: %v\n", value.CanSet())
}
return nil
}


func main() {
	s := InspectStruct{
		Age: 22,          // поле экспортное
		name: "Василий",  // поле неэкспортное
	}

	inspectStruct(s)      // функция инспектирующая структуру
	}
