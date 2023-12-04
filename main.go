package main

import (
	"fmt"

	"github.com/StevenAndre/collabtasks/settings"
)

func main() {

	s,err := settings.NewSettings()
	if err!=nil{
		panic(err)
	}
	fmt.Println(s)
}
