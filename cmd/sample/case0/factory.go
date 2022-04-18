package main

import (
	"fmt"
)

type structFactory struct {
	mainV    func(Case) error
	listV    func([]Case) error
	mpD      func() map[string]Case
	mpV      func(map[string]Case) error
	contentD func() Case
	contentV func(Case) error
}

func NewStructFactory(
	mainV func(Case) error,
	listV func([]Case) error,
	mpD func() map[string]Case,
	mpV func(map[string]Case) error,
	contentD func() Case,
	contentV func(Case) error,
) *structFactory {
	return &structFactory{
		mainV:    mainV,
		listV:    listV,
		mpD:      mpD,
		mpV:      mpV,
		contentD: contentD,
		contentV: contentV,
	}
}

func (_f *structFactory) NewStruct(
	key string, // required
	memo *string, // optional,nillable
	main Case, // validation
	list *[]Case, // optional,validation
	mp *map[string]Case, // default,validation
	content *Case, // default,optional,validation
) (*Struct, error) {

	// main
	if err := _f.mainV(main); err != nil {
		return nil, fmt.Errorf("`main` validation error: %v", err)
	}

	// list
	var rList []Case
	if list != nil {
		rList = *list
	}

	if err := _f.listV(rList); err != nil {
		return nil, fmt.Errorf("`list` validation error: %v", err)
	}

	// mp
	var rMp map[string]Case
	if mp == nil {
		rMp = _f.mpD()
	} else {
		rMp = *mp
	}

	if err := _f.mpV(rMp); err != nil {
		return nil, fmt.Errorf("`mp` validation error: %v", err)
	}

	// content (optional > default)
	var rContent Case
	if content != nil {
		rContent = *content
	}

	if err := _f.contentV(rContent); err != nil {
		return nil, fmt.Errorf("`content` validation error: %v", err)
	}

	return &Struct{
		key:     key,
		memo:    memo,
		main:    main,
		list:    rList,
		mp:      rMp,
		content: rContent,
	}, nil
}
