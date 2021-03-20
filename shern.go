package main

import (
	"flag"
	"fmt"
	"net/http"
	"strings"

	"github.com/rusldv/shern/lib"
)

const (
	// ServerName название сервера
	ServerName = "Shern"
	// ServerVersion текущая версия
	ServerVersion = "0.0.1"
)

var conf = flag.String("config", "./config.json", "Initial configuration file.")

func main() {
	flag.Parse()
	fmt.Println(*conf)
	msg := "ok"
	// Загрузка конфигурационных данных из заданного файла
	cfg, err := lib.ParseConfig(*conf)
	if err != nil {
		msg = "Файл конфигурации" + *conf + "не обнаружен"
		fmt.Println(msg)
		return
	}
	fmt.Println("Конфигурация загружена из файла:", *conf)
	fmt.Println(cfg)
	// Функция в ответ на запрос
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//fmt.Println(r.URL.Query())
		/*
			a := r.URL.Query().Get("a")
			b := r.URL.Query().Get("b")
			fmt.Println(a, b)
		*/
		//fmt.Fprintf(w, "%s %s\n", ServerName, ServerVersion)
		//fmt.Fprintf(w, "Ответ сервера: %s", msg)
		// Получаем каталог для запрашиваемого хоста
		reqHost := cfg.RootDir + r.Host
		// Проверяем каталог с названием хоста
		if !lib.GetHostDir(reqHost) {
			fmt.Println("Директория сайта " + reqHost + "не обнаружена") // TODO
		}
		// Проверяем URI путь запроса
		if lib.IsRootDir(r.URL.Path) {
			// Если запрошен только домен вызываем main/view.html
			tmpl, err := lib.RunComponent(cfg, reqHost, "main", nil, nil)
			if err != nil {
				fmt.Println(err)
			}
			tmpl.Execute(w, nil)
		} else {
			// Иначе вызываем другой компонент и передаем параметры - полный путь запроса
			sl := strings.Split(r.URL.Path, "/")
			tmpl, err := lib.RunComponent(cfg, reqHost, sl[1], sl, nil)
			if err != nil {
				fmt.Println(err)
			}
			// Пробуем запустить контроллер и передать результат его вызова в шаблон
			r := lib.StartAccept(cfg, reqHost, sl[1], "accept.go", "Accept", sl, nil)
			fmt.Println(r)
			
			tmpl.Execute(w, r)
		}
	})
	fmt.Println("Server listening...")
	http.ListenAndServe(":"+cfg.Port, nil)
}
