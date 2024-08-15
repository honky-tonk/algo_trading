package main

import (
	"algo_trading/pkg/assets"
	"fmt"
)

func main() {
	//net.Get("https://www.google.com/")
	//net.GetFromUrl("https://www.baidu.com/")
	s := assets.ParseConfig()
	err := s.GetPrice()
	if err != nil {
		fmt.Println(err.Error())
	}

}
