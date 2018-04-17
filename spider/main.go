package main

import (
	"go_spider/spider/spiders"
	_ "go_spider/spider/database"
)

func main() {
	price := &spiders.Autohome{
		App_price_url: "http://cars.app.autohome.com.cn/jiage_v7.9.5/cars/ownerpricecinfo-pm1-p30386-c0-sp28533-se3324.json",
	}
	price.Get_price()
}