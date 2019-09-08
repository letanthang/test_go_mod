package main

import (
	"fmt"

	"github.com/letanthang/validator"
)

type Student struct {
	ID        int
	FirstName string `json:"first_name" bson:"last_name" validate:"required"`
	LastName  string `json:"last_name" bson:"last_name" validate:"min=3"`
}

func main() {
	aStudent := Student{ID: 5, FirstName: "ABC"}
	if err := validator.Validate(aStudent); err != nil {
		fmt.Println(err)
	}
	fmt.Println("hehe")
}
