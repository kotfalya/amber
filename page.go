package amber

import (
	"errors"
	"math/rand"
	"sync"
	"time"

	"github.com/kotfalya/amber/utils"
)

type Page struct {
	leaf         bool
	scaleStarted bool
	actualSize   uint
	muRW         sync.RWMutex
	keys         map[string]Key
	leafs        []*Page
	req          chan *PageReq
	stop         chan struct{}
	seed         uint32
}

func NewPage() *Page {
	page := &Page{
		leaf:         true,
		scaleStarted: false,
		actualSize:   0,
		muRW:         sync.RWMutex{},
		keys:         make(map[string]Key, *pageKeysSize),
		leafs:        make([]*Page, *pageLeafPoolSize),
		req:          make(chan *PageReq, *pageReqBufferSize),
		stop:         make(chan struct{}),
		seed:         rand.Uint32(),
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
		keyName := req.args[0].(string)
		key := req.args[1].(Key)
		err := p.add(keyName, key)

		req.AddRes(NewEmptyPageRes(err))
	case "remove":

	}
}

func (p *Page) start() {
	sem := utils.NewSemaphore(*pageReqBufferSize)

	scaleTicker := time.NewTicker(time.Second * time.Duration(*checkPageSizeEvery)).C

	for {
		select {
		case req := <-p.req:
			sem.Acquire()
			go func(req *PageReq) {
				defer sem.Release()
				p.handler(req)
			}(req)
		case <-scaleTicker:
			go p.startScaleProcess()
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
		child := p.getLeaf(keyName)
		p.muRW.RUnlock()

		return child.load(keyName)
	}
}

func (p *Page) add(name string, key Key) (err error) {
	p.muRW.Lock()

	if p.leaf {
		defer p.muRW.Unlock()

		p.keys[name] = key
		p.actualSize += 1

		err = nil
	} else {
		child := p.getLeaf(name)
		p.muRW.Unlock()

		err = child.add(name, key)
	}

	return
}
func (p *Page) remove(key Key) (err error) {
	err = nil
	p.actualSize -= 1
	return
}

func (p *Page) getLeaf(keyName string) (leaf *Page) {
	index := utils.GetIndex(keyName, uint32(*pageLeafPoolSize), p.seed)
	leaf = p.leafs[index]
	return
}
