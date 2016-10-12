package db

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
		for name, key := range leaf.keys {
			p.keys[name] = key
		}
		leaf.muRW.RUnlock()
	}
}

func (p *Page) moveKeysToLeafs() {
	p.muRW.RLock()
	defer p.muRW.RUnlock()

	for name, key := range p.keys {
		leaf := p.getLeaf(name)
		leaf.add(name, key)
	}
}

func (p *Page) createLeafs() {
	for i := range p.leafs {
		p.leafs[i] = NewPage()
	}
}

func (p *Page) removeLeafs() {
	for i := range p.leafs {
		p.leafs[i] = nil
	}
}
