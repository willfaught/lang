package goo

var _ Map = MapUint16String(nil)

var _ Pointer = &MapUint16String{}

// MapUint16String is a map from uint16 to string.
type MapUint16String map[uint16]string

// Delete implements Map.
func (m MapUint16String) Delete(k interface{}) {
	delete(m, k.(uint16))
}

// Dereference implements Map.
func (m *MapUint16String) Dereference() Value {
	return *m
}

// Equals implements Map.
func (m MapUint16String) Equals(other Equatable) bool {
	var n = other.(MapUint16String)

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
func (m MapUint16String) Get(k interface{}) interface{} {
	return m[k.(uint16)]
}

// GetCheck implements Map.
func (m MapUint16String) GetCheck(k interface{}) (interface{}, bool) {
	var v, ok = m[k.(uint16)]

	return v, ok
}

// KeyValues implements Map.
func (m MapUint16String) KeyValues() [][2]interface{} {
	var kvs [][2]interface{}

	for k, v := range m {
		kvs = append(kvs, [2]interface{}{k, v})
	}

	return kvs
}

// Keys implements Map.
func (m MapUint16String) Keys() []interface{} {
	var ks []interface{}

	for k := range m {
		ks = append(ks, k)
	}

	return ks
}

// Len implements Map.
func (m MapUint16String) Len() int {
	return len(m)
}

// Make implements Map.
func (m MapUint16String) Make(c int) Map {
	return make(MapUint16String, c)
}

// NotEquals implements Map.
func (m MapUint16String) NotEquals(other Equatable) bool {
	return !m.Equals(other)
}

// Reference implements Map.
func (m MapUint16String) Reference() Pointer {
	return &m
}

// Set implements Map.
func (m MapUint16String) Set(k, v interface{}) {
	m[k.(uint16)] = v.(string)
}
