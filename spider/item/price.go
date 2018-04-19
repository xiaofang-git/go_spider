package item

import (
	"fmt"
	"strconv"
	"go_spider/spider/database"
)

var DB = database.Connect() 

type PriceInfo struct {
	Message string
	Result result
}

type result struct {
	Carname string
	Pricemessage pricemessage
}

type pricemessage struct {
	Fullprice string
	Seriesid  int
	Specid  int

}


func (self *PriceInfo) Save() {
	carname := self.Result.Carname
	price, _ := strconv.ParseFloat(self.Result.Pricemessage.Fullprice, 64)
	seriesid := self.Result.Pricemessage.Seriesid
	specid := self.Result.Pricemessage.Specid
	fmt.Println(carname, price, seriesid ,specid)

	spider := DB.Table("autohome")
	id, err := spider.Data(map[string]interface{}{"carname":carname, "price": price, "seriesid": seriesid, "specid": specid}).Insert()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(id, "数据保存成功")
}