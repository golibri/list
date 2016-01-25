package list

import (
	"strings"
	"testing"
)

func TestAccessors(t *testing.T) {
	str := "a b c d e d e d"
	slice := strings.Fields(str)
	list := FromSlice(slice)

	if strings.Join(slice, ",") != list.Join(",") {
		t.Errorf("construction/joining error: %v", list.Join(","))
	}
	if list.First() != "a" {
		t.Errorf("first element should be 'a', got: %v", list.First())
	}
	if list.Last() != "d" {
		t.Errorf("last element should be 'd', got: %v", list.Last())
	}
	if sliceMismatch(list.FirstN(3), []string{"a", "b", "c"}) {
		t.Errorf("first n elements mismatch: %v", list.FirstN(3))
	}
	if sliceMismatch(list.LastN(2), []string{"e", "d"}) {
		t.Errorf("last n elements mismatch: %v", list.LastN(2))
	}
	if sliceMismatch(list.Slice(2, 4), []string{"c", "d", "e"}) {
		t.Errorf("slice mismatch: %v", list.Slice(2, 4))
	}

	if list.Commonest() != "d" {
		t.Errorf("commonest element error, got: %v", list.Commonest())
	}

	list.Push("x").Shift("y")

	if list.First() != "y" {
		t.Errorf("shift error, first element should be 'y', got: %v: ", list.First())
	}
	if list.Last() != "x" {
		t.Errorf("push error, last element should be 'x', got: %v", list.Last())
	}

	a := list.Pop()
	if a != "x" {
		t.Errorf("pop error, expected: 'x', got: %v", a)
	}
	b := list.Unshift()
	if b != "y" {
		t.Errorf("unshift error, expected: 'y', got: %v", b)
	}

	list.PushAll([]string{"1", "2", "3"})
	list.PopN(3)
	list.ShiftAll([]string{"1", "2", "3"})
	list.UnshiftN(3)

	if sliceMismatch(list.Contents(), slice) {
		t.Errorf("error pushing n items around, %v", list.Contents())
	}

	if list.IndexOf("a") != 0 {
		t.Errorf("error indexOf: expected: 0, got: %v", list.IndexOf("a"))
	}

	if !list.Contains("a") {
		t.Errorf("error list contain")
	}

	if list.Contains("9") {
		t.Errorf("error list contain")
	}

	list.Uniq()
	if sliceMismatch(list.Contents(), []string{"a", "b", "c", "d", "e"}) {
		t.Errorf("uniq error, got: %v", list.Contents())
	}

	list.Filter(func(s string) bool { return s != "c" })
	if sliceMismatch(list.Contents(), []string{"a", "b", "d", "e"}) {
		t.Errorf("filter error, got: %v", list.Contents())
	}

	list.Map(func(s string) string { return s + s })
	if sliceMismatch(list.Contents(), []string{"aa", "bb", "dd", "ee"}) {
		t.Errorf("map error, got: %v", list.Contents())
	}
}

func sliceMismatch(a, b []string) bool {
	if len(a) != len(b) {
		return true
	}
	for i, _ := range a {
		if a[i] != b[i] {
			return true
		}
	}
	return false
}
