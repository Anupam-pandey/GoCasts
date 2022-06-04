package main

import (
	"net/http"
	"fmt"
)
func main()  {
	resp, err := http.Get("https://google.com")
	if err==nil{
		fmt.Println(resp.Body)
	}else{
		fmt.Println(err)
	}


}