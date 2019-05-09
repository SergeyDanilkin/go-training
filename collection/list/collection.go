package list

import (
	"fmt"
	"runtime"
)

func New() *Collection {
	return &Collection{}
}

func (c *Collection) Add(elem int) {
	ci := Element{}
	ci.val = elem

	if c.len == 0 {
		c.FirstItem = &ci
	} else {
		ci.PrevItem = c.LastItem
		c.LastItem.NextItem = &ci
	}
	c.LastItem = &ci
	c.len++
}

func (c *Collection) Get(index int) *Element {
	if index > c.len {
		return nil
	}
	i := c.FirstItem
	j := 0
	for j < index-1 {
		j++
		i = i.Next()
	}
	return i
}

func (c *Collection) Remove(index int) {
	if index > c.len {
		return
	}

	i := c.Get(index)

	if i.Next() != nil {
		i.Next().PrevItem = i.Prev()
	} else {
		c.LastItem = i.Prev()
	}

	if i.Prev() != nil {
		i.Prev().NextItem = i.Next()
	} else {
		c.FirstItem = i.Next()
	}

	c.len--
	runtime.GC()
}

func (c *Collection) Print() {
	if c.len == 0 {
		fmt.Println("")
	} else {
		i := c.FirstItem
		for j := 0; j < c.len; j++ {
			fmt.Print(i.Value(), " ")
			i = i.Next()
		}
		fmt.Println("")
	}
}

func (collection *Collection) First() *Element {
	return collection.FirstItem
}

func (collection *Collection) Last() *Element {
	return collection.LastItem
}

func (collection *Collection) Length() int {
	return collection.len
}
