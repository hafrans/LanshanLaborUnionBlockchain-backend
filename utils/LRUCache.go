// Copyright 2020 Hafrans<hafrans@163.com> All rights reserved.

/*
	Package utils provides some tools
 */
package utils


import (
	"errors"
	"fmt"
)

type lnode struct {
	Key  string
	Data interface{}
	Prev *lnode
	Next *lnode
}

// LRUCache
type LRUCache struct {
	nodes    map[string]*lnode
	capacity uint32
	length   uint32
	head     *lnode
	tail     *lnode
}

// The NewLRUCache can create a LRUCache Container with a specified
// capacity
func NewLRUCache(cap uint32) *LRUCache{
	return &LRUCache{
		nodes: make(map[string]*lnode,cap),
		capacity: cap,
		length: 0,
	}
}

// Put k v into the container
func (c *LRUCache) Put(key string, value interface{}) bool {

	if c.capacity == 0 {
		panic("capacity is zero!")
	}

	var node *lnode = nil
	oldnode, ok := c.nodes[key]
	if ok {

		node = oldnode
		node.Data = value
		c.updateNode(node)

	} else {

		node = &lnode{
			Key:  key,
			Data: value,
			Prev: nil,
			Next: nil,
		}

		if c.length == 0 {
			c.head = node
			c.tail = node
			c.length++
		} else {
			if c.length == c.capacity { // full
				c.deleteLastNode()
			}
			node.Next = c.head
			c.head.Prev = node
			c.head = node
			if c.length != c.capacity {
				c.length++
			}
		}

		c.nodes[key] = node

	}

	return true
}

// Get value by key from container
func (c *LRUCache) Get(key string) (result interface{}, err error) {

	node, ok := c.nodes[key]
	if ok {
		// update nodes
		c.updateNode(node)
		return node.Data, nil
	} else {
		return nil, errors.New("key not found")
	}

}


func (c *LRUCache) Remove(key string) bool {

	if c.length == 0 {
		return false
	}

	node, ok := c.nodes[key]

	if ok {
		c.length--
		if c.length == 0 { // head & tail
			c.head = nil
			c.tail = nil
			c.length = 0
		} else if node == c.head {
			c.head = node.Next
		} else if node == c.tail{
			c.deleteLastNode()
		} else{
			if node.Prev != nil {
				node.Prev.Next = node.Next
			}
			if node.Next != nil {
				node.Next.Prev = node.Prev
			}
		}
		delete(c.nodes, key)
		return true
	} else {
		return false
	}
}

func (c *LRUCache) updateNode(node *lnode) {
	if c.head != node { // 如果不是头结点

		if c.tail == node{
			c.tail = node.Prev
		}
		node.Prev.Next = node.Next
		if node.Next != nil {
			node.Next.Prev = node.Prev
		}

		node.Prev = nil
		node.Next = c.head
		c.head.Prev = node
		c.head = node

	}
}

func (c *LRUCache) deleteLastNode() {
	node := c.tail
	if node == c.head{
		node.Prev.Next = nil
	}
	c.tail = node.Prev
	delete(c.nodes, node.Key)
}

func (c LRUCache) GetCapacity() uint32 {
	return c.capacity
}

func (c LRUCache) GetLength() uint32 {
	return c.length
}

func (c LRUCache) DebuggerGetKVList(){
	head := c.head

	for head != nil{
		fmt.Printf("ITEM: %v => %v \n",head.Key,head.Data)
		head = head.Next
	}
}
