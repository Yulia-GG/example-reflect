package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

type User struct {
    Name string
	Age int
	secret string
}

func main() {
u := User{Name: "Вася",Age: 22, secret: "old secret"} // инициализируем обьект
val := reflect.ValueOf(&u).Elem()  // обязательно указатель передавать &, иначе паника

field := val.FieldByName("Name") // ищем по имени
if field.CanSet() {      // если можем, то изменяем на "Петю"
	field.SetString("Петя")
}

fmt.Println(u.Name)  // будет Петя, т.к.экспортируемое поле

field = val.FieldByName("secret")  // ищем по имени
if field.CanSet() {     // если можем, то изменяем
	field.SetString("new secret")
}

fmt.Println("can set: " + u.secret)  // останется old secret, т.к.неэкспортируемое поле

// обход неэкспортного поля, чтобы его изменить
// используем в тестировании в основном
ptrToSecret := unsafe.Pointer(field.UnsafeAddr()) // получаем у этого поля адрес и преобразуем его в указатель
realPtr := (*string)(ptrToSecret)  // говорим, что по указателю точно лежит строка, дай ее настоящий адрес
*realPtr = "new secret via unsafe"  // уже меняем

fmt.Println(u.secret) //вывели на экран
}
