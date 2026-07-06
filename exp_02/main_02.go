package exp02

import (
	"fmt"
	"reflect"
)

// структура пользователя
type User struct {
    Age int `json:"age" bson:"age" gorm:"pk"`  // у возраста указано 3 тега
	Name string `json:"display_name"`
}

// у структуры есть метод Greeting
func (u User) Greeting(name string) string {
	return fmt.Sprintf("Привет, %s! Меня зовут %s, мне %d лет", name, u.Name, u.Age)
	}

func main() {
	s := User{Age: 22, Name: "Василий"}

	val := reflect.ValueOf(s) // получили значение
    typ := reflect.TypeOf(s)  // получили тип

	// начинаем итерироваться по полям
	for i := 0; i < val.NumField(); i ++ {
	field := typ.Field(i)  // информация о поле
	value := val.Field(i)
	fmt.Println("___")
	fmt.Println("Name:", field.Name)
	fmt.Println("Value:", value.Interface())
    fmt.Println("JSON tag:", field.Tag.Get("json"))
	}

	m := val.MethodByName("Greeting")  // получаем значение по имени
	if m.IsValid() {  // проверяем, что имя валидное
		args := []reflect.Value{reflect.ValueOf("Маша")} // формируем параметры для этой функции
		fmt.Printf("call method: %v\n", m.Call(args)) // вызываем метод m.Call(), передав аргументы
	}
}