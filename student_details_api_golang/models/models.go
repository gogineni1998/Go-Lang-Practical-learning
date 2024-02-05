package models

type Student struct {
	Name       string  `json:"name"`
	Age        int     `json:"age"`
	Statistics Marks   `json:"marks,omitempty"`
	Address    Address `json:"address"`
}

type Marks struct {
	Maths     int `json:"maths,omitempty"`
	Physics   int `json:"physics,omitempty"`
	Chemistry int `json:"chemistry,omitempty"`
}

type Address struct {
	Street string `json:"street"`
	City   string `json:"city"`
	County string `json:"county"`
	State  string `json:"state"`
	Pin    int    `json:"pincode"`
}
