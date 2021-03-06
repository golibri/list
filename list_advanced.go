package list

// Checks if the List has zero elements
func (list *List) IsEmpty() bool {
	return list.length == 0
}

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

// remove entries based on a truth-function
func (list *List) Filter(f func(string) bool) {
	if list.length == 0 {
		return
	}
	var start, current *ListItem
	length := 0
	tmp := list.first
	for i := 0; i < list.length; i++ {
		if f(tmp.Content) {
			if length == 0 {
				start = tmp
				current = tmp
			} else {
				current.Next = tmp
				current = tmp
			}
			length++
		}
		if tmp.Next != nil {
			tmp = tmp.Next
		} else {
			break
		}
	}
	list.first = start
	list.last = current
	list.length = length
}

// map function
func (list *List) Map(f func(string) string) {
	tmp := list.first
	for i := 0; i < list.length; i++ {
		tmp.Content = f(tmp.Content)
		tmp = tmp.Next
	}
}

// Iterate over all elements, call f for every element
func (list *List) Each(f func(string)) {
	tmp := list.first
	for i := 0; i < list.length; i++ {
		f(tmp.Content)
		tmp = tmp.Next
	}
}

// return the most commonest item (in case of duplicated items!)
func (list *List) Commonest() string {
	count := make(map[string]int)
	list.Each(func(s string) { count[s]++ })
	var c_string string
	var c_count int
	for k, v := range count {
		if v > c_count {
			c_string = k
			c_count = v
		}
	}
	return c_string
}
