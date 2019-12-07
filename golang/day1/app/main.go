package main

import (
	"app/model"
	"encoding/json"
	"fmt"
	"strconv"
)

func main() {
	person := model.Person{
		Name:     "fer", //`json:"name"` xml:"name"
		LastName: "sar",
		Age:      34,
	}
	clientModel := model.Client{
		ClientName:     "fer", //`json:"name"` xml:"name"
		ClientLastName: "sar",
		ClientAge:      34,
	}

	fmt.Println(clientModel)

	response, err := json.Marshal(person)
	if err != nil {
		panic("boom!")
	}
	fmt.Println("Person ", person)
	fmt.Println(string(response))

	client := person.ToClient()
	clientJson, err := json.Marshal(client)
	if err != nil {
		panic("boom!")
	}
	fmt.Println(string(clientJson))

	fmt.Println("=======MY MAAAAP========")
	fmt.Println(CreateMap())
	mapJSON, err := json.Marshal(CreateMap())
	if err != nil {
		panic("boom!")
	}
	fmt.Println(string(mapJSON))

	for index, value := range CreateMap() {
		fmt.Println("index: " + index)
		fmt.Println("value:")
		fmt.Println(value)
	}

	fmt.Println("=======Slice========")
	fmt.Println(MySlice())

	fmt.Println("=======Slice copy========")
	MySliceCopy()

}

func CreateMap() map[string]model.Person {
	myMap := map[string]model.Person{}
	myMap["key1"] = model.Person{
		Name:     "Jorge", //`json:"name"` xml:"name"
		LastName: "Castro",
		Age:      30,
	}
	myMap["key2"] = model.Person{
		Name:     "Fer", //`json:"name"` xml:"name"
		LastName: "Fer",
		Age:      34,
	}
	//myMap["key1"] = "mi valor 2 de mapa"
	return myMap
}

func MySlice() []string {
	mySlice := []string{}
	for i := 1; i <= 10; i++ {
		mySlice = append(mySlice, "random string: "+strconv.Itoa(i))
	}

	return mySlice
}

func MySliceCopy() {
	mySlice := []int{1, 2, 3}
	mySlice2 := []int{5, 6, 7, 8}

	var ar [4]int
	newSlice := append(mySlice, mySlice2...)

	fmt.Println("Longitud:")
	fmt.Println(len(mySlice))
	fmt.Println(cap(mySlice))

	fmt.Println("Longitud:")
	fmt.Println(len(mySlice2))
	fmt.Println(cap(mySlice2))

	fmt.Println("Longitud:")
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))
	fmt.Println(len(ar))
	fmt.Println(cap(ar))

	ns := newSlice[2:4]
	fmt.Println(ns)
}
