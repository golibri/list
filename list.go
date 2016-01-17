// Specialized struct to deal with lists of strings
package list

import "strings"

type List struct {
	length int
	first  *ListItem
	last   *ListItem
}

// constructor for a new List
func New() List {
	return List{}
}

// alternative constructor to init a new List from a slice
func FromSlice(slice []string) List {
	list := New()
	for _, s := range slice {
		list.Push(s)
	}
	return list
}

// Return a slice with the contents of the List
func (list *List) Contents() []string {
	slice := make([]string, 0)
	if list.length == 0 {
		return slice
	}
	tmp := list.first
	for i := 0; i < list.length; i++ {
		slice = append(slice, tmp.Content)
		if tmp.IsLast() {
			break
		}
		tmp = tmp.Next
	}
	return slice
}

// Return a string with all strings within the List concatenated
func (list *List) Join(separator string) string {
	return strings.Join(list.Contents(), separator)
}

func (list *List) String() string {
	return list.Join(",")
}

// Clone the whole list (deep)
func (list *List) Clone() List {
	n := New()
	if list.length == 0 {
		return n
	}
	tmp := list.first
	for i := 0; i < list.length; i++ {
		s := tmp.Content
		n.Push(s)
		tmp = tmp.Next
	}
	return n
}
