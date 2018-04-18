package main

import (
	"fmt"
	"go_spider/spider/spiders"
	// _ "go_spider/spider/database"
	// "go_spider/spider/item"
	"encoding/json"
)


func run() {
	price := &spiders.Autohome{
		App_price_url: "http://cars.app.autohome.com.cn/jiage_v7.9.5/cars/ownerpricecinfo-pm1-p30386-c0-sp28533-se3324.json",
	}
	// json_body := &item.Price{}
	var json_body map[string]interface{}
	html := price.Get_price()
	// fmt.Println(html)
	err := json.Unmarshal(html, &json_body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(json_body)
}


func main() {
	run()
}