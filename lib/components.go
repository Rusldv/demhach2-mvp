package lib

import (
	"fmt"
	"html/template"
	"os"
	"io/ioutil"

	"github.com/traefik/yaegi/interp"
	"github.com/traefik/yaegi/stdlib"
)

// GetView проверяет наличие элемента каталога компонента
func GetView(cfg *Config, name string) (*template.Template, error) {
	_, err := os.Stat(name)
	if err != nil {
		if os.IsNotExist(err) {
			tmpl, err := template.ParseFiles(cfg.File404)
			if err != nil {
				return nil, err
			}
			return tmpl, nil
		}
		//fmt.Println(err)
		tmpl, err := template.ParseFiles(cfg.FileError)
		if err != nil {
			return nil, err
		}
		return tmpl, nil
	}
	tmpl, err := template.ParseFiles(name)
	if err != nil {
		return nil, err
	}
	return tmpl, nil
}

// StartAccept вызывает контроллер компонента
func StartAccept(cfg *Config, reqHost string, name string, fname string, actFunc string, items []string, args map[string]string) map[string]string {
	fmt.Println("StartAccept", name, fname, actFunc, items, args)
	path := reqHost+"/"+name+"/"+fname
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("No file accept.go in", path)
			return nil
		}
		fmt.Println(err)
		return nil
	}
	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	src := string(data)
	println(src)
	i := interp.New(interp.Options{})
	i.Use(stdlib.Symbols)
	_, err = i.Eval(src)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	sing := name + "." + actFunc
	fmt.Println(sing)
	v, err := i.Eval(sing)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	fx := v.Interface().(func() map[string]string)
	res := fx()
	
	return res
}

// RunComponent выполняет указанный компонент
func RunComponent(cfg *Config, reqHost string, name string, items []string, args map[string]string) (*template.Template, error) {
	//fmt.Println("hostDir:", reqHost)
	//fmt.Println("name:", name)
	//fmt.Println("args:", args)
	// проверяем, чтобы присутствовал и был каталогом
	t, err := GetView(cfg, reqHost+"/"+name+"/view.html")
	if err != nil {
		return nil, err
	}
	return t, nil
}
