package spiders

import (
	"fmt"
	"net/http"
	// "strings"
)


type Autohome struct {
	App_price_url string
}

func (self *Autohome) Get_price() []byte {
	response, err := http.Get(self.App_price_url)
	if err != nil {
		fmt.Println(err)
	}
	var body []byte = make([]byte, 1000, 2000)
	response.Body.Read(body)
	// text := string(body)
	fmt.Println(string(body[:cap(body)]))
	return body
}