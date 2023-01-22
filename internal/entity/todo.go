package entity

import "github.com/go-playground/validator"

type Todo struct {
	ID          string `json:"id"`
	Name        string `json:"name" validate:"required"`
	Priority    int    `json:"priority" validate:"required,gt=0"`
	IsCompleted bool   `json:"isCompleted"`
}

func (e *Todo) Validate() error {
	validate := validator.New()

	return validate.Struct(e)
}
