package model

type Person struct {
	Name     string
	LastName string
	Age      int
}

func (p Person) ToClient() Client {
	return Client{
		ClientName:     p.Name,
		ClientLastName: p.LastName,
		ClientAge:      p.Age,
	}
}
