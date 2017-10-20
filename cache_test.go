package cache_test

import (
	"os"
	"testing"

	"github.com/maximelamure/cache"
)

func TestMain(m *testing.M) {
	cache.Init("127.0.0.1:11211")
	os.Exit(m.Run())
}

func TestPrimitiveTypes(t *testing.T) {
	if err := cache.Set("int", 5, 120); err != nil {
		t.Error("TestPrimitiveType failed for Set int", err)
	}

	var i int
	if !cache.Get("int", &i) {
		t.Error("TestPrimitiveType failed for Get int")
	}

	if i != 5 {
		t.Error("TestPrimitiveType failed for int. The Get returns a different value. Expected 5, got: ", i)
	}

	if err := cache.Set("string", "string", 120); err != nil {
		t.Error("TestPrimitiveType failed for Set string", err)
	}

	var s string
	if !cache.Get("string", &s) {
		t.Error("TestPrimitiveType failed for Get string")
	}

	if s != "string" {
		t.Error("TestPrimitiveType failed for string. The Get returns a different value. Expected `string`, got: ", s)
	}

}

type Person struct {
	Name string
}

func NewPerson(firstName, lastName string) Person {
	return Person{Name: firstName + " " + lastName}
}

func TestStructType(t *testing.T) {
	p := NewPerson("Maxime", "Lamure")
	if err := cache.Set("person", p, 320); err != nil {
		t.Error("TestCache failed for Set Person", err)
	}

	var twin Person
	if !cache.Get("person", &twin) {
		t.Error("TestStructType failed for Get Person")
	}

	if twin != p {
		t.Error("TestStructType failed for Person. The Get returns a different value. Got: ", twin)
	}

}
