package models

type Album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums=[]Album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 80},
	{ID: "2", Title: "Red Train", Artist: "John Coltrane1", Price: 90},
	{ID: "3", Title: "Green Train", Artist: "John Coltrane2", Price: 100},
}
