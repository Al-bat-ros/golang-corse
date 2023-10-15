package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	len   int
	front *ListItem
	back  *ListItem
}

func (l *list) Len() int {
	return l.len
}

func (l *list) Front() *ListItem {
	return l.front
}

func (l *list) Back() *ListItem {
	return l.back
}

func (l *list) PushFront(v any) *ListItem {
	item := &ListItem{Value: v, Next: l.front}

	if l.front != nil {
		l.front.Prev = item
	} else {
		l.back = item
	}
	l.front = item
	l.len++
	return item
}

func (l *list) PushBack(v any) *ListItem {
	item := &ListItem{Value: v, Prev: l.back}

	if l.back != nil {
		l.back.Next = item
	} else {
		l.front = item
	}

	l.back = item
	l.len++
	return item
}

func (l *list) Remove(i *ListItem) {
	if i.Prev != nil {
		i.Prev.Next = i.Next
	} else {
		l.front = i.Next
	}

	if i.Next != nil {
		i.Next.Prev = i.Prev
	} else {
		l.back = i.Prev
	}
	i.Prev, i.Next = nil, nil
	l.len--
}

func (l *list) MoveToFront(i *ListItem) {
	if l.len <= 1 || l.front == i {
		return
	}
	l.Remove(i)
	l.PushFront(i.Value)
}

func NewList() List {
	return new(list)
}
