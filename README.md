[![GoDoc](https://godoc.org/github.com/golibri/list?status.svg)](https://godoc.org/github.com/golibri/list)
[![Built with Spacemacs](https://cdn.rawgit.com/syl20bnr/spacemacs/442d025779da2f62fc86c2082703697714db6514/assets/spacemacs-badge.svg)](http://github.com/syl20bnr/spacemacs)

# golibri/list
Convenience library for working with slices of strings (`[]string`). Since Go
doesn't have *generics*, this library is targeted explicitly for dealing with
slices of *strings*.

The underlying data structure is implemented as a single linked list, so most
methods have *linear* runtime, with very little memory overhead.

# installation
`go get github.com/golibri/list`

# usage
````go
import "github.com/golibri/list"

func main() {
    l := list.New() // => []
    l.Push("world") // => ["world"]
    l.Shift("hello") // => ["hello","world"]
    l.Push("again") // => ["hello","world","again"]
    w := l.Pop() // => "again"; ["hello","world"]
    l.PushAll([]string{"hello","world"}) // => ["hello","world","hello","world"]
    l.Uniq() // ["hello","world"]
    l.Contains("hello") // => true
    l.Contents() // []string{"hello","world"}
    // ...
}
````

# license
LGPLv3. (You can use it in commercial projects as you like, but improvements/bugfixes must flow back to this lib.)
