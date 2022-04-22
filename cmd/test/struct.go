package main

import "time"

//go:generate go run ../../ -type Struct
type Struct struct {
	Key     string
	Memo    *string         `sc:"optional,nillable"`
	Main    Case            `sc:"validation"`
	List    []Case          `sc:"optional,validation"`
	Mp      map[string]Case `sc:"default,validation"`
	Content Case            `sc:"default,optional,validation"`
	D       time.Duration
}

type Case struct {
	M string
	S int
}
