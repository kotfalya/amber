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

	if p.leaf {
		defer p.muRW.RUnlock()
		return p.load(keyName, index)
	} else {
		child := p.children[index]
		p.muRW.RUnlock()

		return child.Load(keyName, index)
	}
}

func (p *Page) Add(key types.Key, index int) error {
	p.muRW.Lock()

	if p.leaf {
		defer p.muRW.Unlock()
		return p.add(key, index)
	} else {
		child := p.children[index]
		p.muRW.RUnlock()

		return child.Add(key, index)
	}
}

func (p *Page) add(key types.Key, index int) (err error) {
	err = nil
	p.keys[key.Name()] = key

	return
}

func (p *Page) load(keyName string, index int) (types.Key, error) {
	var (
		key types.Key
		err error
		ok  bool
	)

	if key, ok = p.keys[keyName]; !ok {
		err = errors.New(ErrUndefinedKey)
	}

	return key, err
}
