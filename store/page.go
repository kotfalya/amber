package store

import (
	"errors"
	"github.com/kotfalya/store/utils"
	"sync"
)

type Page struct {
	leaf  bool
	muRW  sync.RWMutex
	keys  map[string]Key
	leafs []*Page
	req   chan *PageReq
	stop  chan struct{}
}

func NewPage() *Page {
	page := &Page{
		leaf:  true,
		muRW:  sync.RWMutex{},
		keys:  make(map[string]Key, *pageKeysSize),
		leafs: make([]*Page, *pageLeafPoolSize),
		req:   make(chan *PageReq, *pageReqBufferSize),
		stop:  make(chan struct{}),
	}
	go page.start()

	return page
}

func (p *Page) AddReq(req *PageReq) {
	req.start()
	p.req <- req

	return
}

func (p *Page) handler(req *PageReq) {
	switch req.name {
	case "get":
		keyName := req.args[0].(string)
		key, err := p.load(keyName)

		req.AddRes(NewKeyPageRes(key, err))
	case "add":
		key := req.args[0].(Key)
		err := p.add(key)

		req.AddRes(NewEmptyPageRes(err))
	case "remove":

	}
}

func (p *Page) start() {
	sem := utils.NewSemaphore(*pageReqBufferSize)

	for {
		select {
		case req := <-p.req:
			sem.Acquire()
			go func(req *PageReq) {
				defer sem.Release()
				p.handler(req)
			}(req)
		case <-p.stop:
			close(p.req)
			return
		}
	}
}

func (p *Page) Stop() {
	close(p.stop)
}

func (p *Page) load(keyName string) (key Key, err error) {
	p.muRW.RLock()

	if p.leaf {
		defer p.muRW.RUnlock()
		var ok bool

		if key, ok = p.keys[keyName]; !ok {
			err = errors.New(ErrUndefinedKey)
		}

		return
	} else {
		index := 0 // TODO add calculate index function
		child := p.leafs[index]
		p.muRW.RUnlock()

		return child.load(keyName)
	}
}

func (p *Page) add(key Key) (err error) {
	p.muRW.Lock()

	if p.leaf {
		defer p.muRW.Unlock()

		p.keys[key.Name()] = key

		err = nil
	} else {
		index := 0 // TODO add calculate index function
		child := p.leafs[index]
		p.muRW.Unlock()

		err = child.add(key)
	}

	return
}
