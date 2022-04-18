package main

import (
	"errors"
	"fmt"
)

func main() {
	factory := NewStructFactory(
		func(c Case) error {
			if c.s == 0 {
				return errors.New("must not be zero")
			}
			return nil
		}, // mainV
		func(l []Case) error {
			if len(l) == 0 {
				return errors.New("case length must not be zero")
			}
			return nil
		}, // listV
		func() map[string]Case {
			return map[string]Case{
				"mpd": {
					m: "hello",
					s: 12,
				},
			}
		}, // mpD
		func(m map[string]Case) error {
			for _, v := range m {
				if len(v.m) == 0 {
					return errors.New("mp value of m length must not be zero")
				}
				if v.s < 0 {
					return errors.New("mp value of s length must not be negative")
				}
			}
			return nil
		}, // mpV
		func() Case {
			return Case{
				m: "hello",
				s: 12,
			}
		}, // contentD
		func(v Case) error {
			if len(v.m) == 0 {
				return errors.New("mp value of m length not be zero")
			}
			if v.s < 0 {
				return errors.New("mp value of s length must not be negative")
			}
			return nil

		}, // contentV
	)
	/*
		NewStruct(
			key string, // required
			memo *string, // optional,nillable
			main Case, // validation
			list *[]Case, // optional,validation
			mp *map[string]Case, // default,validation
			content *Case, // default,optional,validation
		) (*Struct, error)
	*/
	l := []Case{
		{
			m: "a",
			s: 12,
		},
	}
	c := Case{
		m: "hello",
		s: 12,
	}
	s, err := factory.NewStruct(
		"key",
		nil,
		Case{
			m: "",
			s: 2,
		}, // main
		&l,  // list
		nil, // mp
		&c,  // content
	)
	if s != nil {
		fmt.Printf("%#v", s)
	} else {
		fmt.Println(err)
	}

}
