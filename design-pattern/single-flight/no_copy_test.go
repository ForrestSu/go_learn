package main

import "testing"

// check by go vet
type noCopy struct{}

func (*noCopy) Lock()   {}
func (*noCopy) Unlock() {}

// Singleton 单例
type Singleton struct {
	noCopy noCopy
}

//go:generate go vet .
func TestSingleton(t *testing.T) {
	s := Singleton{}
	copied := s
	money := 1.02
	t.Logf("%v, %v, %v\n", money, &s, &copied)
}
