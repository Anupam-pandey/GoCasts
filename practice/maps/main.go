package main

import "fmt"

func main()  {

	colors := map[string]string{
		"red":"laaal",
		"yellow":"pila",
		"white":"safeed",
	}
	fmt.Println(colors)

	var cars map[string]string
	fmt.Println(cars)

	person := make(map[string]string)
	fmt.Println(person)

	delete(colors,"red")
	fmt.Println(colors)
	mapprinter(colors)
}


func mapprinter(m map[string]string){
	for k,v := range m{
		fmt.Println(k," and ",v)
	}
}