package cache

import (
	"container/ring"
)

type ClockItem struct {
	key CacheKey
	bit bool
}

type ClockPolicy struct {
	list      *CircularList
	keyNode   map[CacheKey]*ring.Ring
	clockHand *ring.Ring
}

func (p *ClockPolicy) Victim() CacheKey {
	var victimKey CacheKey
	var nodeItem *ClockItem
	for {
		currentNode := (*p.clockHand)
		nodeItem = currentNode.Value.(*ClockItem)
		if nodeItem.bit {
			nodeItem.bit = false
			currentNode.Value = nodeItem
			p.clockHand = currentNode.Next()
		} else {
			victimKey = nodeItem.key
			p.list.Move(p.clockHand.Prev())
			p.clockHand = nil
			p.list.Remove(&currentNode)
			delete(p.keyNode, victimKey)
			return victimKey
		}
	}
}


func (p *ClockPolicy) Add(key CacheKey) {
	node := p.list.Append(&ClockItem{key, true})
	if p.clockHand == nil {
		p.clockHand = node
	}
	p.keyNode[key] = node
}

func (p *ClockPolicy) Remove(key CacheKey) {
	node, ok := p.keyNode[key]
	if !ok {
		return
	}

	if p.clockHand == node {
		p.clockHand = p.clockHand.Prev()
	}
	p.list.Remove(node)
	delete(p.keyNode, key)
}

func (p *ClockPolicy) Access(key CacheKey) {
	node, ok := p.keyNode[key]
	if !ok {
		return
	}
	node.Value = &ClockItem{key, true}
}
