package store

import (
	"errors"
	"sync"

	"github.com/kotfalya/store/key"
	"github.com/kotfalya/store/utils"
)

const (
	ERR_UNDEFINED_KEY = "store:undefined key"
)

type Page struct {
	leaf     bool
	muRW     sync.RWMutex
	keys     map[string]*key.Key
	children []*Page
}

func NewPage() *Page {
	return &Page{
		leaf:     true,
		keys:     make(map[string]*key.Key, *pageKeysSize),
		children: make([]*Page, *pageChildSize),
	}
}

func (p *Page) Exists(keyName string) bool {
	p.muRW.RLock()
	defer p.muRW.RUnlock()

	return p.exists(keyName)
}

func (p *Page) Load(keyName string) (*key.Key, error) {
	p.muRW.RLock()
	defer p.muRW.RUnlock()

	if p.leaf {
		return p.load(keyName)
	} else {
		return p.proxy(keyName)
	}
}

//func (p *Page) Add() {
//
//}

func (p *Page) proxy(keyName string) (*key.Key, error) {
	position := utils.Hash(keyName, pageChildSize) - 1
	p.children[position].Load(keyName)
}

func (p *Page) exists(keyName string) bool {
	_, ok := p.keys[keyName]

	return ok
}

func (p *Page) load(keyName string) (*key.Key, error) {
	if res, ok := p.keys[keyName]; ok {
		return res, nil
	} else {
		return nil, errors.New(ERR_UNDEFINED_KEY)
	}
}

//func (p *Page) Save(keyName string, value interface{}) error {
//    var (
//        res *key.Key
//        ok  bool
//    )
//
//    reflect.
//        res, ok = p.keys[keyName]
//    if !ok {
//        res = New
//    }
//
//    if res, ok = p.keys[keyName]; ok {
//        k
//        return nil
//    } else {
//        return nil, errors.New("store:undefined key")
//    }
//
//}
