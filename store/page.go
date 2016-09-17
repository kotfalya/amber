package store

import (
	"errors"
	"github.com/kotfalya/store/types"
	"sync"
)

type Page struct {
	leaf     bool
	muRW     sync.RWMutex
	keys     map[string]types.Key
	children []*Page
}

func NewPage() *Page {
	return &Page{
		leaf:     true,
		muRW:     sync.RWMutex{},
		keys:     make(map[string]types.Key, *pageKeysSize),
		children: make([]*Page, *pageChildSize),
	}
}

func (p *Page) Load(keyName string, index int) (types.Key, error) {
	p.muRW.RLock()
	defer p.muRW.RUnlock()

	return p.load(keyName, index)
}

func (p *Page) Add(key types.Key, index int) error {
	p.muRW.Lock()
	defer p.muRW.Unlock()

	return p.add(key, index)
}

func (p *Page) add(key types.Key, index int) (err error) {
	if p.leaf {
		p.keys[key.Name()] = key
	} else {
		err = p.children[index].Add(key, index)
	}

	return err
}

func (p *Page) load(keyName string, index int) (types.Key, error) {
	var (
		key types.Key
		err error
		ok  bool
	)

	if p.leaf {
		key, ok = p.keys[keyName]
		if !ok {
			err = errors.New(ErrUndefinedKey)
		}
	} else {
		key, err = p.children[index].Load(keyName, index)
	}

	return key, err
}
