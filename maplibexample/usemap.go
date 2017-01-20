package main

import (
	"fmt"
	"github.com/boyzhujian/learngo/maplibexample/lib"
)

func main() {
	lib.Add("dr", "dart")
	fmt.Println(lib.Get("dr"))
	language := lib.GetAll()
	for _, v := range language {
		fmt.Println(v)
	}
}
