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
	var html []byte = make([]byte, 2048)
	response.Body.Read(html)
	// text := string(body)
	return html
}