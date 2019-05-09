package list

type Collection struct {
	FirstItem *Element
	LastItem  *Element
	len       int
}

type Element struct {
	val      int
	NextItem *Element
	PrevItem *Element
}
