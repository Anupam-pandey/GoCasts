package main

import "fmt"

type bot interface{
	getgreetings() string
}

type englishbot struct{}
type spanishbot struct{}

func main(){

	eb := englishbot{}
	sb := spanishbot{}

	printgreetings(eb)
	printgreetings(sb)


}

func printgreetings(b bot){
	fmt.Println(b.getgreetings())
}

func (eb englishbot) getgreetings() string{
	return "hi"
}


func (eb spanishbot) getgreetings() string{
	return "hola"
}