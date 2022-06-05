package main

import (
	"net/http"
	"fmt"
	"os"
	"io"
)

type logWriter struct{}

func main()  {
	resp, err := http.Get("https://google.com")
	if err==nil{
		// bs := make([]byte,99999)
		// resp.Body.Read(bs)
		// fmt.Println(string(bs))


		// io.Copy(os.Stdout,resp.Body) // since first one is implementing writer and second is implementing reader interface so it'll work
		lw := logWriter{}
		io.Copy(lw,resp.Body)
	}else{
		fmt.Println(err)
		os.Exit(1)
	}
}


func (lw logWriter) Write(bs []byte) (int , error){
	fmt.Println(string(bs))

	return len(bs),nil
}