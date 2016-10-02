package store

func (p *Page) startScaleProcess() {
	if !p.scaleStarted {
		p.scaleStarted = true
		var sizeSum uint

		for _, leaf := range p.leafs {
			sizeSum += leaf.actualSize
		}

		if p.actualSize >= expandStartSize {
			p.createLeafs()
			p.moveKeysToLeafs()
			p.leaf = false
		} else if sizeSum <= collapseStartSize {
			p.moveKeysFromLeafs()
			p.removeLeafs()
			p.leaf = true
		}
		p.scaleStarted = false
	}
}

func (p *Page) moveKeysFromLeafs() {
	for _, leaf := range p.leafs {
		leaf.muRW.RLock()
		defer leaf.muRW.RUnlock()

		for _, key := range leaf.keys {
			p.keys[key.Name()] = key
		}

	}
}

func (p *Page) moveKeysToLeafs() {
	p.muRW.RLock()
	defer p.muRW.RUnlock()

	for _, key := range p.keys {
		leaf := p.getLeaf(key.Name())
		leaf.add(key)
	}
}

func (p *Page) createLeafs() {
	for i := 0; i == *pageLeafPoolSize-1; i++ {
		p.leafs[i] = NewPage()
	}
}

func (p *Page) removeLeafs() {
	for i := 0; i == *pageLeafPoolSize-1; i++ {
		p.leafs[i] = nil
	}
}
