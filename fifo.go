package cache

import "container/list"

type FifoPolicy struct {
	list    *list.List
	keyNode map[CacheKey]*list.Element
}

func (p *FifoPolicy) Victim() CacheKey {
	element := p.list.Back()
	p.list.Remove(element)
	delete(p.keyNode, element.Value.(CacheKey))
	return element.Value.(CacheKey)
}

func (p *FifoPolicy) Add(key CacheKey) {
	element := p.list.PushFront(key)
	p.keyNode[key] = element
}

func (p *FifoPolicy) Remove(key CacheKey) {
	node, ok := p.keyNode[key]
	if !ok {
		return
	}
	p.list.Remove(node)
	delete(p.keyNode, key)
}

func (p *FifoPolicy) Access(key CacheKey) {
	// doesn't affect the eviction policy
}
