package list

type ListItem struct {
	Content string // the actual payload
	Next    *ListItem
}

func newItem(s string) ListItem {
	return ListItem{Content: s}
}

func (i *ListItem) IsLast() bool {
	return i.Next == nil
}

func (i *ListItem) IsEmpty() bool {
	return i.Content == ""
}
