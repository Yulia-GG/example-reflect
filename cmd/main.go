package main

import (
     "reflect"
     "regexp"
     "strconv"
     "strings"
	 "fmt"
)

type User struct {
     Name   string `validate:"min=3"`
     Age    int     `validate:"min=18;max=65"`
     Email string `validate:"regexp=^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\\.[a-zA-Z]{2,}$"`
}

// atoi для преобразования строки в число, т.к. в тегах структуры значения обычно хранятся как строки
func atoi(s string) int {
    n, _ := strconv.Atoi(s)
    return n
}

func Validate(v any) error {
    val := reflect.ValueOf(v)
    typ := reflect.TypeOf(v)

    // итерируемся по полям
    for i := 0; i < val.NumField(); i++ {
       field := val.Field(i)
       tag := typ.Field(i).Tag.Get("validate") // получаем тег
       if tag == "" {
          continue
       }

      // разбиваем теги (строку пр.:“min=3;max=10;required=true”) по правилам валидации
      rules := strings.Split(tag, ";")  // разбиваем по ; и получаем слайс строк (пр.:[“min=3”, “max=10”, “required=true”])
       for _, rule := range rules {
          parts := strings.SplitN(rule, "=", 2)  // “min=3” превратится в: [“min”, “3”]
          if len(parts) < 2 {
             continue
          }

          // само правило валидации
          key, value := parts[0], parts[1]
          switch key {
          case "min":
             switch field.Kind() {
             case reflect.String:  // если поле строка
                if len([]rune(field.String())) < atoi(value) {
                   return fmt.Errorf("field %s: min length %s", typ.Field(i).Name, value)
                }
             case reflect.Int:
                if field.Int() < int64(atoi(value)) {
                   return fmt.Errorf("field %s: min value %s", typ.Field(i).Name, value)
                }
             }
          case "max":
             switch field.Kind() {
             case reflect.String:
                if len(field.String()) > atoi(value) {
                   return fmt.Errorf("field %s: max length %s", typ.Field(i).Name, value)
                }
             case reflect.Int:
                if field.Int() > int64(atoi(value)) {
                   return fmt.Errorf("field %s: max value %s", typ.Field(i).Name, value)
                }
             }
          case "regexp":
             if field.Kind() == reflect.String && !regexp.MustCompile(value).MatchString(field.String()) {
                return fmt.Errorf("field %s: invalid format", typ.Field(i).Name)
             }
          }
       }
    }
    return nil
}

func main() {
if err := Validate(User{Name: "Ив", Age: 18, Email:
"test@example.com"}); err != nil {
fmt.Println("Validation error:", err)
}

if err := Validate(User{Name: "Иван", Age: 70, Email:
"test@example.com"}); err != nil {
fmt.Println("Validation error:", err)
}

if err := Validate(User{Name: "Иван", Age: 35, Email:
"invalid email"}); err != nil {
fmt.Println("Validation error:", err)
}

if err := Validate(User{Name: "Иван", Age: 35, Email:
"test@example.com"}); err != nil {
fmt.Println("Validation error:", err)
}
}