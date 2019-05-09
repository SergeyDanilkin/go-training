package list

func (ci Element) Next() *Element {
	return ci.NextItem
}

func (ci Element) Prev() *Element {
	return ci.PrevItem
}

func (ci Element) Value() int {
	return ci.val
}
