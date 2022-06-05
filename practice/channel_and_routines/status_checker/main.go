package main

import (
	"net/http"
	"fmt"
	"time"
)



func main()  {
	// var c chan string
	c := make(chan string)
	link := []string {
		"https://google.com/",
		"https://facebook.com/",
		"https://gmail.com/",
		"https://amazon.com/",
	}
	verifylink(link, c )

	// for i:=0;i<len(link);i++{
	// 	fmt.Println(<-c)
	// }

	for{
		verifylinkInfinity(c)
	}

}


func verifylink(link []string, c chan string){
	for _,l := range link{
			 go checkstatus(l,c)
	}
}

func verifylinkInfinity( c chan string){
	// for{
	// 	go checkstatus(<-c,c)
	// }

	for l:= range c{
		go func(link string){	
			time.Sleep(5*time.Second)			// lamda function in go (literal function)
			checkstatus(link,c)
			}(l)
	}
}


func checkstatus(l string,c chan string){

	_,err := http.Get(l)
	if err!=nil{
		fmt.Println("something went wrong")
		c<-l

	}else{
		fmt.Println(l," is working!!")
		c<-l
	}
}
