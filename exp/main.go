package main

import "fmt"

type Cat struct{}

func (c Cat) Speak() {
	fmt.Println("meow")
}

type Dog struct{}

func (d Dog) Speak() {
	fmt.Println("woof")
}

type Husky struct {
	// Dog
	Speaker
}

type SpeakerPrefixer struct {
	Speaker
}

func (sp SpeakerPrefixer) Speak() {
	fmt.Print("Prefix: ")
	sp.Speaker.Speak()
}

type Speaker interface {
	Speak()
}

func main() {
	d := Dog{}
	d.Speak()

	h := Husky{Dog{}}
	h.Speak()

	c := Husky{Cat{}}
	c.Speak()

	c1 := Husky{SpeakerPrefixer{Cat{}}}
	c1.Speak()
}
