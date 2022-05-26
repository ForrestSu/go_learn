package dao

import (
	"context"
	"fmt"
)

type Student interface {
	Invoke(ctx context.Context, reqbody interface{}, rspbody interface{}, opts ...int) error
}

var DefaultStudent = New()

func New() Student {
	return &student{}
}

// impl
type student struct{}

func (s *student) Invoke(ctx context.Context, reqbody interface{}, rspbody interface{}, opts ...int) error {
	fmt.Println("student.invoke()...")
	return nil
}
