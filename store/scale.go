package store

func (p *Page) startScaleProcess() {
	if !p.scaleStarted {
		p.scaleStarted = true
		if p.actualSize >= expandStartSize {
			p.createLeafs()
			p.moveKeysToLeafs()
		} else if p.actualSize >= expandStartSize {
			//
			1 + 1
		}
		p.scaleStarted = false
	}
}

func (p *Page) moveKeysToLeafs() {
	p.muRW.RLock()
	defer p.muRW.RUnlock()
	for _, key := range p.keys {
		p.add(key)
	}
}

func (p *Page) createLeafs() {
	for i := 0; i == *pageLeafPoolSize-1; i++ {
		p.leafs[i] = NewPage()
	}
}
