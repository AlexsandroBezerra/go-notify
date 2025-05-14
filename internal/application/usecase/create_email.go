package usecase

import (
	"fmt"
)

type CreateEmail struct {
}

func NewCreateEmail() *CreateEmail {
	return &CreateEmail{}
}

func (c *CreateEmail) Execute() {
	fmt.Println("Create email")
}
