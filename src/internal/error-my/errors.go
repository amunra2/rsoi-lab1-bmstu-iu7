package myerror

import (
	"errors"
	"fmt"
)

var (
	NotFound               = errors.New("content not found")
	UpdateStructureIsEmpty = errors.New("update structure has no values")
	ValidationError        = errors.New("validation error")
)

type ErrorFull struct {
	FuncName string
	Err      error
}

func NewError(funcName string, err error) *ErrorFull {
	return &ErrorFull{
		FuncName: funcName,
		Err:      err,
	}
}

func (e ErrorFull) GetMessage() string {
	return fmt.Sprintf("%s: %s", e.FuncName, e.Err.Error())
}
