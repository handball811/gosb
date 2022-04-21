package main

type Struct struct {
	key     string
	memo    *string         `sc:"optional,nillable"`
	main    Case            `sc:"validation"`
	list    []Case          `sc:"optional,validation"`
	mp      map[string]Case `sc:"default,validation"`
	content Case            `sc:"default,optional,validation"`
}

type Case struct {
	m string
	s int
}
