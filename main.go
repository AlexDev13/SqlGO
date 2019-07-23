package main

import (
	"fmt"
	"os"
	"io"

)

func main(){
	fmt.Println("Hello!")
	file, err := os.Open("products_backupes.sql")
		if err != nil{
			fmt.Println(err)
			os.Exit(1)
		}
	defer file.Close()


	data := make([]byte,64)
	for{
		n,err :=file.Read(data)
			if err == io.EOF{
				break
			}
			sqlData := data[:n]
		fmt.Println(sqlData)
	}


}

