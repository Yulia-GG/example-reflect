package httprouter

import (
	"net/http"
	"reflect"
	"strings"
	"fmt"
)

type Controller struct{}

func (c Controller) Hello(w http.ResponseWriter, _*http.Request) {
	_, _ = fmt.Fprintf(w, "Привет! Это метод Hello() через reflection\n")
}

func (c Controller) Bye(w http.ResponseWriter, _*http.Request) {
	_, _ = fmt.Fprintf(w, "Пока! Это метод Bye() через reflection\n")
}

// есть Router куда встраиваем контроллер
type Router struct {
	controller any
}

// конструктор Router, куда в качестве зависимости принимается контроллер и отдается роутер
func NewRouter(controller any) *Router { return &Router{controller: controller}}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	methodName := strings.Title(req.URL.Path[1:]) // получает имя метода из пути, который нам прислали
	val := reflect.ValueOf(r.controller)  // получает значение контроллера

	method := val.MethodByName(methodName)  // поискали метод
	if !method.IsValid() {
		http.Error(w, "404 not found", http.StatusNotFound)  // если не нашли
	return
	}

	args := []reflect.Value {  // если нашли
		reflect.ValueOf(w),
		reflect.ValueOf(req),
	}
	method.Call(args)
}

 func main () {
	ctrl := new(Controller)  // создаем контроллер
	router := NewRouter(ctrl)  // создаем роутер

	fmt.Println("Сервер запущен на http://localhost:8080/")

	if err := http.ListenAndServe(":8080", router); err != nil { return }
 }