package main

import (
	"log"

	"github.com/maximelamure/cache"
)

const (
	CacheKey string = "person"
)

type Person struct {
	Name string
}

func NewPerson(firstName, lastName string) Person {
	return Person{Name: firstName + " " + lastName}
}

func main() {
	cache.Init("127.0.0.1:11211")
	p := NewPerson("Maxime", "Lamure")
	if err := cache.Set(CacheKey, p, 320); err != nil {
		log.Fatalln(err)
	}

	var twin Person
	if !cache.Get(CacheKey, &twin) {
		log.Fatalln("MISS")
	}

	if twin == p {
		log.Println("Twins")
	} else {
		log.Fatalln("Not twins")
	}
}
