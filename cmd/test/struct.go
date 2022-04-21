package main

//go:generate go run ../../ -type Struct
type Struct struct {
	Key     string
	Memo    *string         `sc:"optional,nillable"`
	Main    Case            `sc:"validation"`
	List    []Case          `sc:"optional,validation"`
	Mp      map[string]Case `sc:"default,validation"`
	Content Case            `sc:"default,optional,validation"`
}

type Case struct {
	M string
	S int
}
