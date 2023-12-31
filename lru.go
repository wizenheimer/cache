package cache

import "container/list"

type LRUPolicy struct {
	list    *list.List
	keyNode map[CacheKey]*list.Element
}

func (p *LRUPolicy) Victim() CacheKey {
	element := p.list.Back()
	p.list.Remove(element)
	delete(p.keyNode, element.Value.(CacheKey))
	return element.Value.(CacheKey)
}

func (p *LRUPolicy) Add(key CacheKey) {
	node := p.list.PushFront(key)
	p.keyNode[key] = node
}

func (p *LRUPolicy) Remove(key CacheKey) {
	node, ok := p.keyNode[key]
	if !ok {
		return
	}
	p.list.Remove(node)
	delete(p.keyNode, key)
}

func (p *LRUPolicy) Access(key CacheKey) {
	p.Remove(key)
	p.Add(key)
}
