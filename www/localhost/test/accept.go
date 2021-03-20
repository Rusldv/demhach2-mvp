package test

import "fmt"

func Accept() map[string]string {
	println("А вывести можно тут")
	fmt.Println("А так")
	return map[string]string {
		"red": "красный",
		"green": "зеленый",
		"blue": "синий",
	}
}