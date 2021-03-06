//go:generate go get github.com/willfaught/goo/cmd/goo
//go:generate goo macro tuples.go ../../tuple_2.go -d {"n":2,"f":["First","Second"]}
//go:generate goo macro tuples.go ../../tuple_3.go -d {"n":3,"f":["First","Second","Third"]}
//go:generate goo macro tuples.go ../../tuple_4.go -d {"n":4,"f":["First","Second","Third","Fourth"]}
//go:generate goo macro tuples.go ../../tuple_5.go -d {"n":5,"f":["First","Second","Third","Fourth","Fifth"]}
//go:generate goo macro tuples.go ../../tuple_6.go -d {"n":6,"f":["First","Second","Third","Fourth","Fifth","Sixth"]}
//go:generate goo macro tuples.go ../../tuple_7.go -d {"n":7,"f":["First","Second","Third","Fourth","Fifth","Sixth","Seventh"]}

package goo

/// {{$n := .n}}

/// {{if false}}
type Pointer interface{}
type Value interface{} /// {{end}}

/// {{if false}}
const __VAR_n__ = 0 /// {{end}}

/// {{if false}}
/// {{$i := 0}}
var __VAR_i__ int /// {{end}}

/// {{if false}}
var _ Equatable = Tuple__VAR_n__{}
var _ Pointer = &Tuple__VAR_n__{}
var _ Value = Tuple__VAR_n__{} /// {{end}}

/// {{if false}}
type Equatable interface{} /// {{end}}

// Tuple__VAR_n__ is a tuple of __VAR_n__ elements.
type Tuple__VAR_n__ [__VAR_n__]interface{}

// Dereference implements Equatable.
func (t *Tuple__VAR_n__) Dereference() Value {
	return *t
}

// Equals implements Equatable.
func (t Tuple__VAR_n__) Equals(other Equatable) bool {
	return t == other.(Tuple__VAR_n__)
}

// NotEquals implements Equatable.
func (t Tuple__VAR_n__) NotEquals(other Equatable) bool {
	return t != other.(Tuple__VAR_n__)
}

// Reference implements Equatable.
func (t Tuple__VAR_n__) Reference() Pointer {
	return &t
}

/// {{range $i, $s := .f}}
// __VAR_s__ returns the element at index __VAR_i__.
func (t Tuple__VAR_n__) __VAR_s__() interface{} {
	return t[__VAR_i__]
}

/// {{end}}
