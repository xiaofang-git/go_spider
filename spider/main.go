package main

import (
	"fmt"
	"go_spider/spider/spiders"
	"go_spider/spider/item"
	"encoding/json"
)


func run() {
	price := &spiders.Autohome{
		App_price_url: "http://cars.app.autohome.com.cn/jiage_v7.9.5/cars/ownerpricecinfo-pm1-p30396-c0-sp28533-se3324.json",
	}
	var info item.PriceInfo
	html := price.Get_price()

	for index := 0; index < len(html); index++ {
		if html[index] == 0 {
			tmp := html[:index]
			err := json.Unmarshal(tmp, &info)
			if err != nil {
				fmt.Println(err)
			}
			
			info.Save()

			break
		}
	}
	
}


func main() {
	run()
}