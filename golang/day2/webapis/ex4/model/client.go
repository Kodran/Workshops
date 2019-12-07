package model

type Client struct {
	ID       int
	Name     string `json:"name" xml:"Name"`
	LastName string `json:"lastName" xml:"LastName"`
	Age      int    `json:"age" xml:"Age"`
}
