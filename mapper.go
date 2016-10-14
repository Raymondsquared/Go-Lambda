package main

import (
	"encoding/base64"
	"encoding/json"
	"log"
)

type SimpleData struct {
	users       []User
	images      []string
	coordinates Coordinates
	price       string
}

type User struct {
	id      int
	name    string
	work    string
	email   string
	dob     string
	address string
	city    string
	optedin bool
}

type Coordinates struct {
	x float64
	y float64
}

type StandardMessage struct {
	Id   string
	Type string
	data string
}

func simpleMapping() {

	coordinates := Coordinates{
		x: 1.11,
		y: -2.22,
	}

	user1 := User{
		id:      0,
		name:    "name1",
		work:    "work1",
		email:   "email1@one.com",
		dob:     "1/1/2011",
		address: "address1",
		city:    "city1",
		optedin: false,
	}

	user2 := User{
		id:      0,
		name:    "name2",
		work:    "work2",
		email:   "email2@two.com",
		dob:     "2/2/2022",
		address: "address2",
		city:    "city2",
		optedin: true,
	}

	var data = SimpleData{
		users:       []User{user1, user2},
		images:      []string{"image1.png", "image2.jpg"},
		coordinates: coordinates,
		price:       "1.1",
	}

	jsonData, errJsonMarshal := json.Marshal(data)

	if errJsonMarshal != nil {
		log.Fatal(errJsonMarshal)
	}

	var msg = StandardMessage{
		Id:   "12321wadkwajdkj1kj21ojeko21nk1",
		Type: "v1/12312/adawd/awd12213/21321/adwa/dwa/d21321",
		data: base64.StdEncoding.EncodeToString([]byte(jsonData)),
	}

	log.Printf("%+v\n", msg)
}
