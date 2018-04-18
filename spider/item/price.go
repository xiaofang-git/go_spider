package item

type Price struct {
	Carname string	`json:"carname"`
	Price float64	`json:"fullprice"`	
	Seriesid int   `json:"seriesid"`
	Specid int		`json:"specid"`
}