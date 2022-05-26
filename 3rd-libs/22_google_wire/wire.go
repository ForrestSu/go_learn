//go:build wireinject
// +build wireinject

package main

import "github.com/google/wire"

//go:generate wire
func InitializeEvent(msg string) Event {
	panic(wire.Build(NewEvent, NewGreeter, NewMessage))
}
