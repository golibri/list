package list

// Return the current length of the List
func (list *List) Length() int {
	return list.length
}

// get first string of the List
func (list *List) First() string {
	if list.length == 0 {
		return ""
	}
	return list.first.Content
}

// get first n strings of the List
func (list *List) FirstN(n int) []string {
	if n < 1 || n >= list.length {
		return emptySlice()
	}
	return list.Slice(0, n-1)
}

// get last string of the List
func (list *List) Last() string {
	if list.length == 0 {
		return ""
	}
	return list.last.Content
}

// get the last n strings of the List
func (list *List) LastN(n int) []string {
	if list.length == 0 {
		return emptySlice()
	}
	return list.Slice(list.length-n, list.length-1)
}

// return a slice [a:b] within List, negative indexes count backwards
func (list *List) Slice(from, to int) []string {
	if list.length == 0 {
		return emptySlice()
	}
	if from >= list.length || to >= list.length || from < 0 || to < 0 || from > to {
		return emptySlice()
	}
	if from == to {
		return oneSlice(list.Get(from))
	}
	diff := to - from + 1
	result := make([]string, diff)
	tmp := list.first
	for i := 0; i < list.length; i++ {
		if i >= from && i <= to {
			result[i-from] = tmp.Content
		}
		if i >= to {
			break
		}
		tmp = tmp.Next
	}
	return result
}

// insert new element at the end of a list
func (list *List) Push(s string) *List {
	item := newItem(s)
	if list.length == 0 {
		list.first = &item
	} else {
		list.last.Next = &item
	}
	list.length++
	list.last = &item
	return list
}

// insert new element for every item in the given slice
func (list *List) PushAll(slice []string) *List {
	for _, s := range slice {
		list.Push(s)
	}
	return list
}

// insert new element at the beginning of a list
func (list *List) Shift(s string) *List {
	if list.length == 0 {
		list.Push(s)
	} else {
		item := newItem(s)
		item.Next = list.first
		list.first = &item
		list.length++
	}
	return list
}

// insert new Element at the beginning of a list for every item in the given slice
func (list *List) ShiftAll(slice []string) *List {
	for _, s := range slice {
		list.Shift(s)
	}
	return list
}

// Get and remove last element of the List
func (list *List) Pop() string {
	if list.length == 0 {
		return ""
	}
	if list.length == 1 {
		s := list.first.Content
		list.first = nil
		list.last = nil
		list.length--
		return s
	}
	s := list.last.Content
	var prev *ListItem
	tmp := list.first
	for i := 0; i < list.length; i++ {
		if i == list.length-2 {
			prev = tmp
			break
		} else {
			tmp = tmp.Next
		}
	}
	prev.Next = nil
	list.length--
	return s
}

// Get and remove the last n elements of the List
func (list *List) PopN(n int) []string {
	if n < 1 || n >= list.length {
		return emptySlice()
	}
	if n == 1 {
		return oneSlice(list.Pop())
	}
	pos := list.length - n
	return list.Slice(pos, -1)
}

// Get and remove first element of the List
func (list *List) Unshift() string {
	if list.length < 2 {
		return list.Pop()
	}
	s := list.first.Content
	list.first = list.first.Next
	list.length--
	return s
}

// Get and remove the first n elements of the List
func (list *List) UnshiftN(n int) []string {
	if n < 1 || n >= list.length {
		return make([]string, 0)
	}
	slice := make([]string, n)
	for i := 0; i < n; i++ {
		slice[i] = list.Unshift()
	}
	return slice
}

// Get string at target position/index within the List
func (list *List) Get(pos int) string {
	if pos < 0 || pos >= list.length {
		return ""
	}
	if pos == 0 {
		return list.first.Content
	}
	tmp := list.first
	for i := 0; i < pos; i++ {
		tmp = tmp.Next
	}
	return tmp.Content
}

// Set the string value at target position/index within the List
func (list *List) Set(pos int, s string) *List {
	if pos < 0 || pos >= list.length {
		return list
	}
	if pos == 0 {
		list.first.Content = s
		return list
	}
	tmp := list.first
	for i := 0; i < pos; i++ {
		tmp = tmp.Next
	}
	tmp.Content = s
	return list
}
