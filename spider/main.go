package main

import (
	"fmt"
	"go_spider/spider/spiders"
	"go_spider/spider/item"
	"encoding/json"
)


func run(id int) {
	price := &spiders.Autohome{
		Id: id,
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
	// 从数据库读取当前库中最大的mid
	max_mid, _ := item.DB.Table("autohome").Max("specid")

	id, _ := max_mid.(int)

	end := id + 1000

	for id < end {
		run(id)
		id++
	}
	item.DB.Close()
}