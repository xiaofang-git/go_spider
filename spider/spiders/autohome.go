package spiders

import (
	"fmt"
	"net/http"
	// "strings"
	"strconv"
)


type Autohome struct {
	Id int
}

func (self *Autohome) Get_price() []byte {
	id := strconv.Itoa(self.Id)
	response, err := http.Get("http://cars.app.autohome.com.cn/jiage_v7.9.5/cars/ownerpricecinfo-pm1-p" + id +  "-c0-sp28533-se3324.json")
	fmt.Println("http://cars.app.autohome.com.cn/jiage_v7.9.5/cars/ownerpricecinfo-pm1-p" + id +  "-c0-sp28533-se3324.json")
	if err != nil {
		fmt.Println(err)
	}
	var html []byte = make([]byte, 2048)
	response.Body.Read(html)
	// text := string(body)
	return html
}