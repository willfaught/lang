package goo

var _ Map = MapStringComplex64(nil)

var _ Pointer = &MapStringComplex64{}

// MapStringComplex64 is a map from string to complex64.
type MapStringComplex64 map[string]complex64

// Delete implements Map.
func (m MapStringComplex64) Delete(k interface{}) {
	delete(m, k.(string))
}

// Dereference implements Map.
func (m *MapStringComplex64) Dereference() Value {
	return *m
}

// Equals implements Map.
func (m MapStringComplex64) Equals(other Equatable) bool {
	var n = other.(MapStringComplex64)

	if len(n) != len(m) {
		return false
	}

	for k, v := range m {
		if nv, ok := n[k]; !ok {
			return false
		} else if nv != v {
			return false
		}
	}

	return true
}

// Get implements Map.
func (m MapStringComplex64) Get(k interface{}) interface{} {
	return m[k.(string)]
}

// GetCheck implements Map.
func (m MapStringComplex64) GetCheck(k interface{}) (interface{}, bool) {
	var v, ok = m[k.(string)]

	return v, ok
}

// KeyValues implements Map.
func (m MapStringComplex64) KeyValues() [][2]interface{} {
	var kvs [][2]interface{}

	for k, v := range m {
		kvs = append(kvs, [2]interface{}{k, v})
	}

	return kvs
}

// Keys implements Map.
func (m MapStringComplex64) Keys() []interface{} {
	var ks []interface{}

	for k := range m {
		ks = append(ks, k)
	}

	return ks
}

// Len implements Map.
func (m MapStringComplex64) Len() int {
	return len(m)
}

// Make implements Map.
func (m MapStringComplex64) Make(c int) Map {
	return make(MapStringComplex64, c)
}

// NotEquals implements Map.
func (m MapStringComplex64) NotEquals(other Equatable) bool {
	return !m.Equals(other)
}

// Reference implements Map.
func (m MapStringComplex64) Reference() Pointer {
	return &m
}

// Set implements Map.
func (m MapStringComplex64) Set(k, v interface{}) {
	m[k.(string)] = v.(complex64)
}
