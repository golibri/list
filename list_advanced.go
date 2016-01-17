package list

// Checks if a specific string exists within the List
func (list *List) Contains(s string) bool {
	return list.IndexOf(s) != -1
}

// Returns the position/index of the (first) given string or -1 if not found
func (list *List) IndexOf(s string) int {
	if list.length == 0 {
		return -1
	}
	tmp := list.first
	for i := 0; i < list.length; i++ {
		if tmp.Content == s {
			return i
		}
		if tmp.Next != nil {
			tmp = tmp.Next
		}
	}
	return -1
}

// remove duplicated entries from List
func (list *List) Uniq() *List {
	if list.length < 2 {
		return list
	}
	m := make(map[string]bool)
	prev := list.first
	m[prev.Content] = true
	tmp := prev.Next
	for i := 0; i < list.length-1; i++ {
		if m[tmp.Content] {
			prev.Next = tmp.Next
			tmp = tmp.Next
		} else {
			m[tmp.Content] = true
			prev = tmp
			tmp = prev.Next
		}
		if tmp == nil {
			break
		}
	}
	return list
}