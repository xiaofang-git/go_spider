package spiders

import (
	"fmt"
	"net/http"
	// "strings"
)


type Autohome struct {
	App_price_url string
}

func (self *Autohome) Get_price()  {
	response, err := http.Get(self.App_price_url)
	if err != nil {
		fmt.Println(err)
	}else{
		body := make([]byte, 5000)
		response.Body.Read(body)
		text := string(body)
		fmt.Println(text)
	}
}