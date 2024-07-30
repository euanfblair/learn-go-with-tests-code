package main

import "fmt"



var languageGreetings = map[string]string{
	"English":"Hello, ",
	"Spanish":"Hola, ",
	"French": "Bonjour, ",
}



func Hello(name, language string) string {
	greeting, ok := languageGreetings[language]
	if name == "" {
		name = "World"
	}
	if ok {
		return greeting + name
	}
	return languageGreetings["English"] + name
}

func main() {
	fmt.Println(Hello("World", "English"))
}
